package disrest

import (
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

func (c *Client) GatewayBot(opts ...rest.RequestOpt) (res *distype.GatewayBotGetResponse, err error) {
	err = c.client.Do(rest.GetGatewayBot.Compile(nil), nil, &res, opts...)
	if err != nil {
		return
	}

	return
}
