package metathings_data_storage_sdk

import (
	"context"
	"sync"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type DataStorage interface {
	Write(ctx context.Context, measurement string, tags map[string]string, data map[string]interface{}) error
}

type DataStorageFactory func(...interface{}) (DataStorage, error)

var data_storage_factories map[string]DataStorageFactory
var data_storage_factories_once sync.Once

func registry_data_storage_factory(name string, fty DataStorageFactory) {
	data_storage_factories_once.Do(func() {
		data_storage_factories = map[string]DataStorageFactory{}
	})
	data_storage_factories[name] = fty
}

func NewDataStorage(name string, args ...interface{}) (DataStorage, error) {
	fty, ok := data_storage_factories[name]
	if !ok {
		return nil, ErrUnsupportedDataStorageDriver
	}

	return fty(args...)
}

func ToDataStorage(ds *DataStorage) func(string, interface{}) error {
	return func(key string, val interface{}) error {
		var ok bool
		if *ds, ok = val.(DataStorage); !ok {
			return opt_helper.InvalidArgument(key)
		}
		return nil
	}
}
