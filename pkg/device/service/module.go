package metathings_device_service

import (
	"context"
	"time"

	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type Module interface {
	Id() string
	Heartbeat()
	IsAlive() bool

	UnaryCall(context.Context, *deviced_pb.OpUnaryCallValue) (*deviced_pb.UnaryCallValue, error)
}

type ModuleImpl struct {
	module        *deviced_pb.Module
	heartbeat_at  time.Time
	alive_timeout time.Duration
}

func (self *ModuleImpl) Id() string {
	return self.module.GetId()
}

func (self *ModuleImpl) Heartbeat() {
	self.heartbeat_at = time.Now()
}

func (self *ModuleImpl) IsAlive() bool {
	return time.Now().Sub(self.heartbeat_at) < self.alive_timeout
}

func (self *ModuleImpl) UnaryCall(ctx context.Context, req *deviced_pb.OpUnaryCallValue) (*deviced_pb.UnaryCallValue, error) {
	panic("unimplemented")
}

func NewModule(module *deviced_pb.Module, alive_timeout time.Duration) Module {
	return &ModuleImpl{
		module:        module,
		heartbeat_at:  time.Time{},
		alive_timeout: alive_timeout,
	}
}
