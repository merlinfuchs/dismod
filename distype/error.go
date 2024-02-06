package distype

import "time"

type RestErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type RestTooManyRequests struct {
	Bucket     string        `json:"bucket"`
	Message    string        `json:"message"`
	RetryAfter time.Duration `json:"retry_after"`
}
