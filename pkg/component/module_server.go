package metathings_component

import "errors"

var (
	ErrUnknownModuleServerAdapter = errors.New("unknown module server adapter")
)

type ModuleServer interface {
	Stop()
	Serve() error
}

type ModuleServerFactory func(*Module) (ModuleServer, error)

var module_server_factories map[string]ModuleServerFactory

func register_module_server_factory(name string, fty ModuleServerFactory) {
	if module_server_factories == nil {
		module_server_factories = make(map[string]ModuleServerFactory)
	}

	module_server_factories[name] = fty
}

func NewModuleServer(name string, mdl *Module) (ModuleServer, error) {
	fty, ok := module_server_factories[name]
	if !ok {
		return nil, ErrUnknownModuleServerAdapter
	}

	return fty(mdl)
}
