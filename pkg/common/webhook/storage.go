package webhook_helper

import (
	"errors"
	"sync"
)

var (
	ErrUnknownStorageDriver = errors.New("unknown storage driver")
	ErrWebhookNotFound      = errors.New("webhook not found")
)

type Storage interface {
	CreateWebhook(wh *Webhook) (*Webhook, error)
	DeleteWebhook(id string) error
	ListWebhooks(wh *Webhook) ([]*Webhook, error)
	GetWebhook(id string) (*Webhook, error)
	UpdateWebhook(id string, wh *Webhook) (*Webhook, error)
}

type StorageFactory interface {
	New(args ...interface{}) (Storage, error)
}

var storage_factories map[string]StorageFactory
var storage_factories_once sync.Once

func register_storage_factory(name string, fty StorageFactory) {
	storage_factories_once.Do(func() {
		storage_factories = make(map[string]StorageFactory)
	})

	storage_factories[name] = fty
}

func NewStorage(name string, args ...interface{}) (Storage, error) {
	fty, ok := storage_factories[name]
	if !ok {
		return nil, ErrUnknownStorageDriver
	}

	return fty.New(args...)
}
