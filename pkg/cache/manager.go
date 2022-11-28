package cache

import (
	"encoding/json"
	"github.com/ixfan/gofan/pkg/cache/store"
	"os"
	"strconv"
)

type Manager struct {
	client ClientInterface
	result string
	err    error
}

type ClientInterface interface {
	Set(key string, value interface{}, second int64) error
	Get(key string) (string, error)
	Remove(key string) error
	Flush() error
}

func NewManager(client ClientInterface) *Manager {
	return &Manager{
		client: client,
	}
}

func Default() *Manager {
	if os.Getenv("cache.default") == "redis" {
		return &Manager{
			client: store.NewRedisStore(),
		}
	} else {
		return &Manager{
			client: store.NewBigCacheStore(),
		}
	}
}

func (manager *Manager) Set(key string, value interface{}, second ...int64) error {
	duration := int64(-1)
	if len(second) > 0 {
		duration = second[0]
	}
	mBytes, err := json.Marshal(value)
	if err != nil {
		return manager.client.Set(key, value, duration)
	} else {
		return manager.client.Set(key, string(mBytes), duration)
	}
}

func (manager *Manager) Get(key string) *Manager {
	manager.result, manager.err = manager.client.Get(key)
	return manager
}

func (manager *Manager) Remove(key string) error {
	return manager.client.Remove(key)
}

func (manager *Manager) Flush() error {
	return manager.client.Flush()
}

func (manager *Manager) ToString() (string, error) {
	return manager.result, manager.err
}

func (manager *Manager) ToInt64() (int64, error) {
	result, err := manager.ToInt()
	return int64(result), err
}

func (manager *Manager) ToInt() (int, error) {
	if manager.err != nil {
		return 0, manager.err
	}
	result, err := strconv.Atoi(manager.result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (manager *Manager) ToFloat64() (float64, error) {
	if manager.err != nil {
		return 0, manager.err
	}
	result, err := strconv.ParseFloat(manager.result, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (manager *Manager) ToStruct(value interface{}) error {
	if manager.err != nil {
		return manager.err
	}
	err := json.Unmarshal([]byte(manager.result), value)
	return err
}
