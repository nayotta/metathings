package metathings_component

import (
	"context"
	"errors"

	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

var (
	ErrUnknownModuleProxyDriver = errors.New("unknown module proxy driver")
)

type ModuleProxy interface {
	UnaryCall(ctx context.Context, req *deviced_pb.OpUnaryCallValue) (*deviced_pb.UnaryCallValue, error)
	StreamCall(ctx context.Context, stm deviced_pb.DevicedService_ConnectClient) error
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
