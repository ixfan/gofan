package connectors

import (
	"github.com/ixfan/gofan/pkg/database/redis"
	"time"
)

type RedisStore struct {
	redisClient redis.ClientInterface
}

func NewRedisStore() *RedisStore {
	return &RedisStore{}
}

func (store *RedisStore) connect() {
	store.redisClient = redis.Default().Connect()
}

func (store *RedisStore) Publish(key string, value interface{}) error {
	store.connect()
	defer store.close()
	_, err := store.redisClient.RPush(key, value).Result()
	return err
}

func (store *RedisStore) Subscribe(key string, subscribe SubscribeInterface) {
	for {
		store.connect()
		result, err := store.redisClient.LPop(key).Result()
		if err == nil {
			subscribe.Handle(result)
		}
		store.close()
		time.Sleep(1 * time.Second)
	}
}

func (store *RedisStore) POP(key string, subscribe SubscribeInterface) {
	store.connect()
	defer store.close()
	result, _ := store.redisClient.LPop(key).Result()
	subscribe.Handle(result)
}

func (store *RedisStore) close() {
	_ = store.redisClient.Close()
}
