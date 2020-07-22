package metathings_deviced_session_storage

import (
	"time"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type SessionStorage interface {
	GetStartupSession(id string) (int32, error)
	SetStartupSessionIfNotExists(id string, sess int32, expire time.Duration) error
	UnsetStartupSession(id string) error
	RefreshStartupSession(id string, expire time.Duration) error
}

type SessionStorageFactory func(...interface{}) (SessionStorage, error)

var session_storage_factories map[string]SessionStorageFactory

func register_session_storage_factory(drv string, fty SessionStorageFactory) {
	if session_storage_factories == nil {
		session_storage_factories = map[string]SessionStorageFactory{}
	}

	session_storage_factories[drv] = fty
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

func ToSessionStorage(y *SessionStorage) func(string, interface{}) error {
	return func(k string, v interface{}) error {
		var ok bool

		if *y, ok = v.(SessionStorage); !ok {
			return opt_helper.InvalidArgument(k)
		}

		return nil
	}
}
