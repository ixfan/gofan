package queue

import (
	"encoding/json"
	"github.com/ixfan/gofan/pkg/queue/connectors"
)

type Manager struct {
	client ClientInterface
}

type ClientInterface interface {
	Publish(key string, value interface{}) error
	Subscribe(key string, handle connectors.SubscribeInterface)
	POP(key string, handle connectors.SubscribeInterface)
}

func Default() *Manager {
	return &Manager{
		client: connectors.NewRedisStore(),
	}
}

func (manager *Manager) Publish(key string, value interface{}) error {
	mBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return manager.client.Publish(key, string(mBytes))
}

func (manager *Manager) Subscribe(key string, handle connectors.SubscribeInterface) {
	manager.client.Subscribe(key, handle)
}

func (manager *Manager) POP(key string, handle connectors.SubscribeInterface) {
	manager.client.POP(key, handle)
}
