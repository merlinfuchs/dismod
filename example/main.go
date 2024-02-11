package main

import (
	"fmt"
	"os"

	"github.com/merlinfuchs/dismod/disrest"
	"github.com/merlinfuchs/dismod/disway"
)

func main() {
	discordToken := os.Getenv("DISCORD_TOKEN")

	client := disrest.NewClient(discordToken)

	cluster := disway.NewCluster(discordToken, client)

	err := cluster.Open()
	if err != nil {
		panic(err)
	}

	fmt.Println(cluster.ShardCount, cluster.MaxConcurrency, len(cluster.Shards))

	select {}
}
