package metathings_identityd2_policy

import (
	"sync"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
)

type BackendCache interface {
	Get(sub, obj *storage.Entity, act *storage.Action) (ret bool, err error)
	Set(sub, obj *storage.Entity, act *storage.Action, ret bool) (err error)
	// Remove("subject", &subject) or
	// Remove("object", &object) or
	// Remove("action", &action) or
	// Remove("subject", &subject, "object", &object) etc.
	Remove(vals ...interface{}) (err error)
}

type BackendCacheFactory interface {
	New(...interface{}) (BackendCache, error)
}

var backend_cache_factories_once sync.Once
var backend_cache_factories map[string]BackendCacheFactory

func register_backend_cache_factory(name string, fty BackendCacheFactory) {
	backend_cache_factories_once.Do(func() {
		backend_cache_factories = make(map[string]BackendCacheFactory)
	})
	backend_cache_factories[name] = fty
}

func NewBackendCache(name string, args ...interface{}) (BackendCache, error) {
	fty, ok := backend_cache_factories[name]
	if !ok {
		return nil, ErrInvalidBackendCacheDriver
	}

	return fty.New(args...)
}
