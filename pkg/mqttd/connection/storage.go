package metathingsmqttdconnection

import "errors"

// ErrUnknownStorageDriver ErrUnknownStorageDriver
var (
	ErrUnknownStorageDriver = errors.New("unknown storage driver")
)

// Storage Storage
type Storage interface {
	AddBridgeToDevice(devID, brID string) error
	RemoveBridgeFromDevice(devID, brID string) error
	ListBridgesFromDevice(devID string) ([]string, error)
}

// StorageFactory StorageFactory
type StorageFactory func(...interface{}) (Storage, error)

var storageFactories map[string]StorageFactory

func registerStorageFactory(name string, fty StorageFactory) {
	if storageFactories == nil {
		storageFactories = map[string]StorageFactory{}
	}

	storageFactories[name] = fty
}

// NewStorage NewStorage
func NewStorage(name string, args ...interface{}) (Storage, error) {
	stFty, ok := storageFactories[name]
	if !ok {
		return nil, ErrUnknownStorageDriver
	}
	stor, err := stFty(args...)
	if err != nil {
		return nil, err
	}
	return stor, nil
}
