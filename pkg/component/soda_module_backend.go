package metathings_component

import "sync"

type SodaModuleBackend interface {
	Start() error
	Stop() error
	Done() <-chan struct{}
}

type SodaModuleBackendFactory func(*Module) (SodaModuleBackend, error)

var soda_module_backend_factories map[string]SodaModuleBackendFactory
var soda_module_backend_factories_once sync.Once

func register_soda_module_backend_factory(name string, fty SodaModuleBackendFactory) {
	soda_module_backend_factories_once.Do(func() {
		soda_module_backend_factories = make(map[string]SodaModuleBackendFactory)
	})

	soda_module_backend_factories[name] = fty
}

func NewSodaModuleBackend(name string, mdl *Module) (SodaModuleBackend, error) {
	fty, ok := soda_module_backend_factories[name]
	if !ok {
		return nil, ErrUnknownSodaModuleBackendDriver
	}

	return fty(mdl)
}
