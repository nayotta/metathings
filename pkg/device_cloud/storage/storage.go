package metathings_device_cloud_storage

import (
	"errors"
	"sync"
)

var (
	ErrInvalidStorageDriver = errors.New("invalid storage driver")
)

var (
	ErrConnectedByOtherDeviceCloud = errors.New("connected by other device cloud")
	ErrNotConnected                = errors.New("not connected")
)

type Storage interface {
	Heartbeat(mdl_id string) error
	IsConnected(sess string, dev_id string) error
	ConnectDevice(sess string, dev_id string) error
	UnconnectDevice(sess string, dev_id string) error
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
		return nil, ErrInvalidStorageDriver
	}

	return fty.New(args...)
}
