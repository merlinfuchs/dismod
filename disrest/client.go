package disrest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/merlinfuchs/dismod/distype"
	"github.com/merlinfuchs/dismod/disutil"
)

type Client struct {
	ratelimiter *RateLimiter
	Log         disutil.Logger

	Token     string
	UserAgent string

	Debug                  bool
	ShouldRetryOnRateLimit bool
	MaxRestRetries         int
	Client                 *http.Client
}

func NewClient(token string) *Client {
	return &Client{
		ratelimiter: NewRateLimiter(),
		Log:         disutil.DefaultLogger,
		Token:       token,

		Debug:                  false,
		ShouldRetryOnRateLimit: true,
		MaxRestRetries:         3,
		Client:                 http.DefaultClient,
	}
}

// Request is the same as RequestWithBucketID but the bucket id is the same as the urlStr
func (c *Client) Request(method, urlStr string, data interface{}, options ...RequestOption) (response []byte, err error) {
	return c.RequestWithBucketID(method, urlStr, data, strings.SplitN(urlStr, "?", 2)[0], options...)
}

// RequestWithBucketID makes a (GET/POST/...) Requests to Discord REST API with JSON data.
func (c *Client) RequestWithBucketID(method, urlStr string, data interface{}, bucketID string, options ...RequestOption) (response []byte, err error) {
	var body []byte
	if data != nil {
		body, err = json.Marshal(data)
		if err != nil {
			return
		}
	}

	return c.request(method, urlStr, "application/json", body, bucketID, 0, options...)
}

// request makes a (GET/POST/...) Requests to Discord REST API.
// Sequence is the sequence number, if it fails with a 502 it will
// retry with sequence+1 until it either succeeds or sequence >= session.MaxRestRetries
func (c *Client) request(method, urlStr, contentType string, b []byte, bucketID string, sequence int, options ...RequestOption) (response []byte, err error) {
	if bucketID == "" {
		bucketID = strings.SplitN(urlStr, "?", 2)[0]
	}
	return c.RequestWithLockedBucket(method, urlStr, contentType, b, c.ratelimiter.LockBucket(bucketID), sequence, options...)
}

// RequestWithLockedBucket makes a request using a bucket that's already been locked
func (c *Client) RequestWithLockedBucket(method, urlStr, contentType string, b []byte, bucket *Bucket, sequence int, options ...RequestOption) (response []byte, err error) {
	req, err := http.NewRequest(method, urlStr, bytes.NewBuffer(b))
	if err != nil {
		bucket.Release(nil)
		return
	}

	// Not used on initial login..
	if c.Token != "" {
		req.Header.Set("authorization", c.Token)
	}

	// Discord's API returns a 400 Bad Request is Content-Type is set, but the
	// request body is empty.
	if b != nil {
		req.Header.Set("Content-Type", contentType)
	}

	req.Header.Set("User-Agent", c.UserAgent)

	cfg := newRequestConfig(c, req)
	for _, opt := range options {
		opt(cfg)
	}
	req = cfg.Request

	if c.Debug {
		for k, v := range req.Header {
			log.Printf("API REQUEST   HEADER :: [%s] = %+v\n", k, v)
		}
	}

	resp, err := cfg.Client.Do(req)
	if err != nil {
		bucket.Release(nil)
		return
	}
	defer func() {
		err2 := resp.Body.Close()
		if c.Debug && err2 != nil {
			log.Println("error closing resp body")
		}
	}()

	err = bucket.Release(resp.Header)
	if err != nil {
		return
	}

	response, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if c.Debug {
		log.Printf("API RESPONSE  STATUS :: %s\n", resp.Status)
		for k, v := range resp.Header {
			log.Printf("API RESPONSE  HEADER :: [%s] = %+v\n", k, v)
		}
		log.Printf("API RESPONSE    BODY :: [%s]\n\n\n", response)
	}

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	case http.StatusNoContent:
	case http.StatusBadGateway:
		// Retry sending request if possible
		if sequence < cfg.MaxRestRetries {

			c.Log(disutil.LogInfo, "%s Failed (%s), Retrying...", urlStr, resp.Status)
			response, err = c.RequestWithLockedBucket(method, urlStr, contentType, b, c.ratelimiter.LockBucketObject(bucket), sequence+1, options...)
		} else {
			err = fmt.Errorf("Exceeded Max retries HTTP %s, %s", resp.Status, response)
		}
	case 429: // TOO MANY REQUESTS - Rate limiting
		rl := distype.RestTooManyRequests{}
		err = json.Unmarshal(response, &rl)
		if err != nil {
			c.Log(disutil.LogError, "rate limit unmarshal error, %s", err)
			return
		}

		if cfg.ShouldRetryOnRateLimit {
			c.Log(disutil.LogInfo, "Rate Limiting %s, retry in %v", urlStr, rl.RetryAfter)

			time.Sleep(rl.RetryAfter)
			// we can make the above smarter
			// this method can cause longer delays than required

			response, err = c.RequestWithLockedBucket(method, urlStr, contentType, b, c.ratelimiter.LockBucketObject(bucket), sequence, options...)
		} else {
			err = &RateLimitError{&RateLimit{RestTooManyRequests: &rl, URL: urlStr}}
		}
	case http.StatusUnauthorized:
		if strings.Index(c.Token, "Bot ") != 0 {
			c.Log(disutil.LogInfo, ErrUnauthorized.Error())
			err = ErrUnauthorized
		}
		fallthrough
	default: // Error condition
		err = newRestError(req, resp, response)
	}

	return
}
