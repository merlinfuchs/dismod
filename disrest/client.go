package disrest

import (
	"context"
	"log/slog"

	"github.com/disgoorg/disgo/rest"
	"github.com/merlinfuchs/dismod/distype"
)

type Client struct {
	log    *slog.Logger
	client rest.Client
}

func NewClient(token string, logger *slog.Logger, opts ...rest.ConfigOpt) *Client {
	opts = append(opts,
		rest.WithLogger(logger),
		rest.WithRateLimiterConfigOpts(
			rest.WithRateLimiterLogger(logger),
		),
	)

	client := rest.NewClient(token, opts...)

	return &Client{
		log:    logger,
		client: client,
	}
}

func (c *Client) Close(ctx context.Context) {
	c.client.Close(ctx)
}

func (c *Client) Request(endpoint *rest.CompiledEndpoint, rqBody interface{}, rsBody interface{}, opts ...rest.RequestOpt) error {
	return c.client.Do(endpoint, rqBody, rsBody, opts...)
}

func (c *Client) GatewayBot(opts ...rest.RequestOpt) (res *distype.GatewayBotGetResponse, err error) {
	err = c.Request(rest.GetGatewayBot.Compile(nil), nil, &res, opts...)
	if err != nil {
		return
	}

	return
}
