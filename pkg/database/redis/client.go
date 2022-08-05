package redis

import (
	"time"
)

type ClientInterface interface {
	Connect() ClientInterface
	Set(key string, value interface{}, duration time.Duration) CmdResultInterface
	Get(key string) CmdResultInterface
	Del(key string) error
	FlushDB() error
	Close() error
	LPop(key string) CmdResultInterface
	RPush(key string, value interface{}) CmdResultInterface
}

type CmdResultInterface interface {
	Result() (string, error)
	String() string
}

func Default() ClientInterface {
	return &GoRedisClient{}
}
