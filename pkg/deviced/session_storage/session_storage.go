package metathings_deviced_session_storage

import "time"

type SessionStorage interface {
	GetStartupSession(id string) (int32, error)
	SetStartupSessionIfNotExists(id string, sess int32, expire time.Duration) error
	UnsetStartupSession(id string) error
	RefreshStartupSession(id string, expire time.Duration) error
}

type SessionStorageFactory func(...interface{}) (SessionStorage, error)

var session_storage_factories map[string]SessionStorageFactory

func register_session_storage_factory(driver string, fty SessionStorageFactory) {
	if session_storage_factories == nil {
		session_storage_factories = map[string]SessionStorageFactory{}
	}

	session_storage_factories[driver] = fty
}

func NewSessionStorage(driver string, args ...interface{}) (SessionStorage, error) {

	fty, ok := session_storage_factories[driver]
	if !ok {
		return nil, ErrUnknownSessionStorageDriver
	}

	stor, err := fty(args...)
	if err != nil {
		return nil, err
	}

	return stor, nil
}
