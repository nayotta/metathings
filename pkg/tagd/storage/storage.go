package metathings_tagd_storage

type Storage interface {
	Tag(ns string, id string, tags []string) error
	Untag(ns string, id string, tags []string) error
	Remove(ns string, id string) error
	Get(ns string, id string) ([]string, error)
	Query(ns string, tags []string) ([]string, error)
}

var storage_factories map[string]func(...interface{}) (Storage, error)

func register_storage_factory(driver string, fty func(...interface{}) (Storage, error)) {
	if storage_factories == nil {
		storage_factories = make(map[string]func(...interface{}) (Storage, error))
	}

	storage_factories[driver] = fty
}

func NewStorage(driver string, args ...interface{}) (Storage, error) {
	fty, ok := storage_factories[driver]
	if !ok {
		return nil, ErrUnknownDriver
	}

	stor, err := fty(args...)
	if err != nil {
		return nil, err
	}

	return stor, nil
}
