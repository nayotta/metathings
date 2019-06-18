package metathings_moqusitto_plugin_storage

import "sync"

type Permission struct {
	Topic *string
	Mask  *string
}

type User struct {
	Username    *string
	Password    *string
	Superuser   *bool
	Permissions []*Permission
}

type Storage interface {
	AddUser(*User) error
	RemoveUser(string) error
	AddPermission(string, *Permission) error
	RemovePermission(username string, topic string) error
	GetUser(string) (*User, error)
	ExistUser(string) (bool, error)
}

type StorageFactory interface {
	New(...interface{}) (Storage, error)
}

var (
	storage_factories      map[string]StorageFactory
	storage_factories_once sync.Once
)

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
