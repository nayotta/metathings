package metathings_device_cloud_storage

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrInvalidStorageDriver        = errors.New("invalid storage driver")
	ErrConnectedByOtherDeviceCloud = errors.New("connected by other device cloud")
	ErrNotConnected                = errors.New("not connected")

	NOTIME = time.Unix(0, 0)
)

type Storage interface {
	Heartbeat(mdl_id string) error
	GetHeartbeatAt(mdl_id string) (time.Time, error)
	SetModuleSession(mdl_id string, sess int64) error
	UnsetModuleSession(mdl_id string) error
	GetModuleSession(mdl_id string) (int64, error)
	SetDeviceConnectSession(dev_id string, sess string) error
	UnsetDeviceConnectSession(dev_id string, sess string) error
	GetDeviceConnectSession(dev_id string) (string, error)
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
