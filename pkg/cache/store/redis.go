package store

import (
	"fmt"
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

func (store *RedisStore) Set(key string, value interface{}, second int64) error {
	store.connect()
	defer store.close()
	result, err := store.redisClient.Set(key, value, time.Duration(second)*time.Second).Result()
	fmt.Println(result, err)
	return err
}

func (store *RedisStore) Get(key string) (string, error) {
	store.connect()
	defer store.close()
	result, err := store.redisClient.Get(key).Result()
	return result, err
}

func (store *RedisStore) Remove(key string) error {
	store.connect()
	defer store.close()
	return store.redisClient.Del(key)
}

func (store *RedisStore) Flush() error {
	store.connect()
	defer store.close()
	return store.redisClient.FlushDB()
}

func (store *RedisStore) close() {
	_ = store.redisClient.Close()
}
