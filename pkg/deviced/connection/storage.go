package metathings_deviced_connection

import (
	"errors"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

var (
	ErrUnknownStorageDriver = errors.New("unknown storage driver")
)

type Storage interface {
	AddBridgeToDevice(dev_id string, sess int32, br_id string) error
	RemoveBridgeFromDevice(dev_id string, sess int32, br_id string) error
	ListBridgesFromDevice(dev_id string, sess int32) ([]string, error)
}

type StorageFactory func(...interface{}) (Storage, error)

var storage_factories map[string]StorageFactory

func register_storage_factory(name string, fty StorageFactory) {
	if storage_factories == nil {
		storage_factories = map[string]StorageFactory{}
	}

	storage_factories[name] = fty
}

func NewStorage(name string, args ...interface{}) (Storage, error) {
	st_fty, ok := storage_factories[name]
	if !ok {
		return nil, ErrUnknownStorageDriver
	}
	stor, err := st_fty(args...)
	if err != nil {
		return nil, err
	}
	return stor, nil
}

func ToStorage(y *Storage) func(string, interface{}) error {
	return func(k string, v interface{}) error {
		var ok bool

		if *y, ok = v.(Storage); !ok {
			return opt_helper.InvalidArgument(k)
		}

		return nil
	}
}
