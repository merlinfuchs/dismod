package distype

import (
	"strconv"
	"time"
)

type Snowflake string

func (s Snowflake) String() string {
	return string(s)
}

func (s Snowflake) Timestamp() time.Time {
	v, _ := strconv.ParseInt(string(s), 10, 64)
	ms := v >> 22
	return time.Unix((ms+1420070400000)/1000, (ms+1420070400000)%1000*1000000)
}
