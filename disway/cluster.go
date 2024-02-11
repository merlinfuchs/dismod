package disway

import (
	"context"
	"io"
	"log/slog"

	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/disgo/sharding"
	"github.com/merlinfuchs/dismod/distype"
)

type Cluster struct {
	log            *slog.Logger
	manager        sharding.ShardManager
	eventListeners map[distype.EventType][]func(s int, e any)
}

func NewCluster(token string, logger *slog.Logger, opts ...sharding.ConfigOpt) *Cluster {
	c := &Cluster{
		log:            logger,
		eventListeners: make(map[distype.EventType][]func(s int, e any)),
	}

	opts = append(opts,
		sharding.WithGatewayConfigOpts(
			gateway.WithEnableRawEvents(true),
			gateway.WithLogger(c.log),
			gateway.WithAutoReconnect(true),
		),
		sharding.WithRateLimiterConfigOpt(
			sharding.WithRateLimiterLogger(c.log),
		),
		sharding.WithLogger(c.log),
		sharding.WithAutoScaling(false),
	)

	c.manager = sharding.New(token, c.handleEvent, opts...)
	return c
}

func (c *Cluster) Open(ctx context.Context) {
	c.manager.Open(ctx)
}

func (c *Cluster) Close(ctx context.Context) {
	c.manager.Close(ctx)
}

func (c *Cluster) handleEvent(t gateway.EventType, _ int, s int, e gateway.EventData) {
	if t != gateway.EventTypeRaw {
		return
	}

	raw := e.(gateway.EventRaw)
	payload, err := io.ReadAll(raw.Payload)
	if err != nil {
		c.log.With("error", err).Error("failed to read raw event payload")
		return
	}

	event, err := distype.UnmarshalEvent(distype.EventType(raw.EventType), payload)
	if err != nil {
		c.log.With("error", err).Error("failed to unmarshal raw event")
		return
	}

	for _, f := range c.eventListeners[distype.EventType(raw.EventType)] {
		go f(s, event)
	}
}

func (c *Cluster) AddEventListener(t distype.EventType, f func(s int, e any)) {
	c.eventListeners[t] = append(c.eventListeners[t], f)
}
