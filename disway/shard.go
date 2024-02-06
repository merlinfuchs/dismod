package disway

import "github.com/merlinfuchs/dismod/disutil"

type Shard struct {
	Log disutil.Logger

	Token      string
	ShardCount int
	ShardID    int
}

func NewShard(token string) *Shard {
	return &Shard{
		Log:   disutil.DefaultLogger,
		Token: token,
	}
}

func (s *Shard) Open() error {
	return nil
}
