package metathings_component

import (
	"sync"

	pb "github.com/nayotta/metathings/proto/component"
)

type SodaModuleWrapper interface {
	pb.ModuleServiceServer
}

type SodaModuleWrapperFactory interface {
	NewModuleWrapper(*Module) (SodaModuleWrapper, error)
}

var soda_module_wrapper_factories_once sync.Once
var soda_module_wrapper_factories map[string]SodaModuleWrapperFactory

func NewSodaModuleWrapper(m *Module) (SodaModuleWrapper, error) {
	cfg := m.Kernel().Config()
	name := cfg.GetString("backend.name")

	fty, ok := soda_module_wrapper_factories[name]
	if !ok {
		return nil, ErrUnknownSodaModuleWrapperDriver
	}

	return fty.NewModuleWrapper(m)
}

func register_soda_module_wrapper_factory(name string, fty SodaModuleWrapperFactory) {
	soda_module_wrapper_factories_once.Do(func() {
		soda_module_wrapper_factories = make(map[string]SodaModuleWrapperFactory)
	})
	soda_module_wrapper_factories[name] = fty
}
