package store

import (
	"github.com/allegro/bigcache/v3"
	"github.com/gookit/goutil/stdutil"
	"time"
)

type BigCacheStore struct {
}

var bigClient *bigcache.BigCache

func NewBigCacheStore() *BigCacheStore {
	return &BigCacheStore{}
}

func (store *BigCacheStore) connect() {
	if bigClient == nil {
		bigClient, _ = bigcache.NewBigCache(bigcache.DefaultConfig(time.Duration(3600) * time.Second))
	}
}

func (store *BigCacheStore) close() {
	_ = bigClient.Close()
}

func (store *BigCacheStore) Set(key string, value interface{}, second int64) error {
	store.connect()
	return bigClient.Set(key, []byte(stdutil.ToString(value)))
}

func (store *BigCacheStore) Get(key string) (string, error) {
	store.connect()
	result, err := bigClient.Get(key)
	return string(result), err
}

func (store *BigCacheStore) Remove(key string) error {
	store.connect()
	return bigClient.Delete(key)
}

func (store *BigCacheStore) Flush() error {
	store.connect()
	return bigClient.Reset()
}
