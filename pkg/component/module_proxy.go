package metathings_component

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/any"
)

var (
	ErrUnknownModuleProxyDriver = errors.New("unknown module proxy driver")
)

type ModuleProxyStream interface {
	Recv() (*any.Any, error)
	Send(*any.Any) error
}

type ModuleProxy interface {
	UnaryCall(ctx context.Context, method string, value *any.Any) (*any.Any, error)
	StreamCall(ctx context.Context, method string) (ModuleProxyStream, error)
}

type ModuleProxyFactory interface {
	NewModuleProxy(args ...interface{}) (ModuleProxy, error)
}

var module_proxy_factories map[string]ModuleProxyFactory

func NewModuleProxy(name string, args ...interface{}) (ModuleProxy, error) {
	fty, ok := module_proxy_factories[name]
	if !ok {
		return nil, ErrUnknownModuleProxyDriver
	}

	return fty.NewModuleProxy(args...)
}

func register_module_proxy_factory(name string, fty ModuleProxyFactory) {
	if module_proxy_factories == nil {
		module_proxy_factories = make(map[string]ModuleProxyFactory)
	}

	module_proxy_factories[name] = fty
}
