package disway

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/merlinfuchs/dismod/disrest"
	"github.com/merlinfuchs/dismod/distype"
	"github.com/merlinfuchs/dismod/disutil"
)

type Cluster struct {
	sync.RWMutex
	Log    disutil.Logger
	Client *disrest.Client

	Token          string
	ShardCount     int
	FirstShardID   int
	LastShardID    int
	MaxConcurrency int
	Gateway        string

	Shards map[int]*Shard
}

func NewCluster(token string, client *disrest.Client) *Cluster {
	c := &Cluster{
		Log:    disutil.DefaultLogger,
		Client: client,
		Token:  token,
		Shards: make(map[int]*Shard),
	}

	return c
}

func (c *Cluster) prepare() error {
	gateway, err := c.Client.GatewayBot()
	if err != nil {
		return fmt.Errorf("failed to get gateway: %w", err)
	}

	c.Gateway = gateway.URL
	if !strings.HasSuffix(c.Gateway, "/") {
		c.Gateway += "/"
	}
	c.Gateway += "?v=" + apiVersion + "&encoding=json"
	fmt.Println(c.Gateway)

	if c.ShardCount == 0 {
		c.ShardCount = gateway.Shards
		c.MaxConcurrency = gateway.MaxConcurrency
	}

	if c.LastShardID == 0 {
		c.LastShardID = c.ShardCount - 1
	}

	if c.MaxConcurrency == 0 {
		c.MaxConcurrency = 1
	}

	for i := c.FirstShardID; i <= c.LastShardID; i++ {
		shard := NewShard(c.Token)
		shard.ShardID = i
		shard.ShardCount = c.ShardCount
		shard.Gateway = c.Gateway
		shard.Dispatch = func(t distype.EventType, d interface{}) {
			fmt.Println("dispatch", t)
			// TODO: dispatch to event handlers
		}
		c.Shards[i] = shard
	}

	return nil
}

func (c *Cluster) Open() error {
	err := c.prepare()
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}

	for d := 0; d < c.MaxConcurrency; d++ {
		wg.Add(1)

		go func(divider int) {
			defer wg.Done()

			for id, shard := range c.Shards {
				if id%c.MaxConcurrency != divider {
					continue
				}

				err := shard.Open()
				if err != nil {
					c.Log(disutil.LogError, "failed to open shard %d: %w", id, err)
					return
				}

				// This does some unnecessary waiting, but it's fine for now.
				if id != len(c.Shards)-1 {
					time.Sleep(5 * time.Second)
				}
			}
		}(d)
	}

	wg.Wait()
	return nil
}
