package metathings_device_service

import (
	"context"

	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type Module interface {
	UnaryCall(context.Context, *deviced_pb.OpUnaryCallValue) (*deviced_pb.UnaryCallValue, error)
}

type ModuleDatabase interface {
	Lookup(component, module string) (Module, error)
}
