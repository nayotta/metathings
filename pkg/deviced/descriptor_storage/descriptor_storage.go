package metathings_deviced_descriptor_storage

import (
	"sync"
)

type DescriptorStorage interface {
	SetDescriptor(sha1 string, body []byte) error
	GetDescriptor(sha1 string) ([]byte, error)
}

type DescriptorStorageFactory func(...interface{}) (DescriptorStorage, error)

var descriptor_storage_factories map[string]DescriptorStorageFactory
var descriptor_storage_factories_once sync.Once

func register_descriptor_storage_factory(name string, fty DescriptorStorageFactory) {
	descriptor_storage_factories_once.Do(func() {
		descriptor_storage_factories = map[string]DescriptorStorageFactory{}
	})

	descriptor_storage_factories[name] = fty
}

func NewDescriptorStorage(name string, args ...interface{}) (DescriptorStorage, error) {
	fty, ok := descriptor_storage_factories[name]
	if !ok {
		return nil, ErrUnknownDescriptorStorageDriver
	}

	return fty(args...)
}
