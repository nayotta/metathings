package metathings_device_service

import (
	"context"
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/any"
	component "github.com/nayotta/metathings/pkg/component"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

var (
	ErrInvalidEndpoint          = errors.New("invalid endpoint")
	ErrInvalidModuleProxyDriver = errors.New("invalid module proxy driver")
)

type Module interface {
	Id() string
	Heartbeat()
	IsAlive() bool
	HeartbeatAt() time.Time

	UnaryCall(context.Context, *deviced_pb.OpUnaryCallValue) (*deviced_pb.UnaryCallValue, error)
}

type ModuleImpl struct {
	module        *deviced_pb.Module
	proxy         component.ModuleProxy
	heartbeat_at  time.Time
	alive_timeout time.Duration
}

func (self *ModuleImpl) Id() string {
	return self.module.GetId()
}

func (self *ModuleImpl) Heartbeat() {
	self.heartbeat_at = time.Now()
}

func (self *ModuleImpl) HeartbeatAt() time.Time {
	return self.heartbeat_at
}

func (self *ModuleImpl) IsAlive() bool {
	return time.Now().Sub(self.heartbeat_at) < self.alive_timeout
}

func (self *ModuleImpl) init_proxy() error {
	var err error

	if self.proxy == nil {
		self.proxy, err = new_module_proxy_by_endpoint(self.module.Endpoint)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *ModuleImpl) UnaryCall(ctx context.Context, req *deviced_pb.OpUnaryCallValue) (*deviced_pb.UnaryCallValue, error) {
	var val *any.Any
	var err error

	if err = self.init_proxy(); err != nil {
		return nil, err
	}

	if val, err = self.proxy.UnaryCall(ctx, req.GetMethod().GetValue(), req.GetValue()); err != nil {
		return nil, err
	}

	return &deviced_pb.UnaryCallValue{
		Name:      req.GetName().GetValue(),
		Component: req.GetComponent().GetValue(),
		Method:    req.GetMethod().GetValue(),
		Value:     val,
	}, nil
}

func NewModule(module *deviced_pb.Module, alive_timeout time.Duration) Module {
	return &ModuleImpl{
		module:        module,
		heartbeat_at:  time.Time{},
		alive_timeout: alive_timeout,
	}
}

func new_module_proxy_by_endpoint(ep string) (component.ModuleProxy, error) {
	u, err := url.Parse(ep)
	if err != nil {
		return nil, err
	}

	scheme := strings.ToLower(u.Scheme)
	if !strings.HasPrefix(scheme, "mtp") {
		return nil, ErrInvalidEndpoint
	}

	proxy_driver := "grpc"
	if strings.HasPrefix(scheme, "mtp+") {
		proxy_driver = strings.TrimPrefix(scheme, "mtp+")
	}

	switch proxy_driver {
	case "grpc":
		return component.NewModuleProxy(proxy_driver, component.NewGrpcModuleServiceClientFactory(u.Host))
	}

	return nil, ErrInvalidModuleProxyDriver
}
