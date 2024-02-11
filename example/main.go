package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/disgo/sharding"
	"github.com/merlinfuchs/dismod/disrest"
	"github.com/merlinfuchs/dismod/distype"
	"github.com/merlinfuchs/dismod/disway"
)

func main() {
	discordToken := os.Getenv("DISCORD_TOKEN")

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	client := disrest.NewClient(discordToken, logger)

	gatewayInfo, err := client.GatewayBot()
	if err != nil {
		logger.Error("failed to get gateway bot", "error", err)
		return
	}

	shardIDs := make([]int, gatewayInfo.Shards)
	for i := range shardIDs {
		shardIDs[i] = i
	}

	cluster := disway.NewCluster(
		discordToken,
		logger,
		sharding.WithShardCount(gatewayInfo.Shards),
		sharding.WithShardIDs(shardIDs...),
		sharding.WithAutoScaling(false),
		sharding.WithGatewayConfigOpts(
			gateway.WithIntents(gateway.IntentGuilds),
		),
	)

	cluster.AddEventListener(distype.EventTypeGuildCreate, func(s int, e any) {
		guild := e.(*distype.GuildCreateEvent)
		logger.Info("received guild create event", "guild", guild.ID)
	})

	cluster.Open(context.Background())

	select {}
}
