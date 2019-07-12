package metathings_device_service

import (
	"context"
	"errors"
	"time"

	"github.com/golang/protobuf/ptypes/any"
	log "github.com/sirupsen/logrus"

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

	UnaryCall(ctx context.Context, req *deviced_pb.OpUnaryCallValue) (*deviced_pb.UnaryCallValue, error)
	StreamCall(ctx context.Context, cfg_req *deviced_pb.ConnectRequest, stm deviced_pb.DevicedService_ConnectClient) error
}

type ModuleStream interface {
	Send(*deviced_pb.OpStreamCallValue) error
	Recv() (*deviced_pb.StreamCallValue, error)
}

type ModuleImpl struct {
	logger        log.FieldLogger
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
		self.proxy, err = self.new_module_proxy_by_endpoint(self.module.Endpoint)
		if err != nil {
			return err
		}
	}
	self.logger.WithField("endpoint", self.module.Endpoint).Debugf("init module proxy")

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

func (self *ModuleImpl) StreamCall(ctx context.Context, cfg_req *deviced_pb.ConnectRequest, upstm deviced_pb.DevicedService_ConnectClient) error {
	var err error

	if err = self.init_proxy(); err != nil {
		return err
	}

	sess := cfg_req.GetSessionId().GetValue()
	cfg := cfg_req.GetStreamCall().GetConfig()
	method := cfg.GetMethod().GetValue()
	mpstm := component.NewModuleProxyStream(upstm, sess)

	if err = self.proxy.StreamCall(ctx, method, mpstm); err != nil {
		return err
	}

	return nil
}

func (self *ModuleImpl) new_module_proxy_by_endpoint(endpoint string) (component.ModuleProxy, error) {
	ep, err := component.ParseEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

	if !ep.IsMetathingsProtocol() {
		return nil, ErrInvalidEndpoint
	}

	proxy_driver := ep.GetTransportProtocol("grpc")
	switch proxy_driver {
	case "grpc":
		return component.NewModuleProxy(proxy_driver,
			"logger", self.logger,
			"client_factory", component.NewGrpcModuleServiceClientFactory(ep.Host))
	}

	return nil, ErrInvalidModuleProxyDriver
}

func NewModule(logger log.FieldLogger, module *deviced_pb.Module, alive_timeout time.Duration) Module {
	return &ModuleImpl{
		logger:        logger,
		module:        module,
		heartbeat_at:  time.Time{},
		alive_timeout: alive_timeout,
	}
}
