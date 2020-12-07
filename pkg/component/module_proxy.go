package metathings_component

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/any"
	"google.golang.org/grpc"

	deviced_pb "github.com/nayotta/metathings/proto/deviced"
)

var (
	ErrUnknownModuleProxyDriver = errors.New("unknown module proxy driver")
)

type ModuleProxyStream interface {
	Send(*any.Any) error
	Recv() (*any.Any, error)
	grpc.ClientStream
}

type moduleProxyStream struct {
	deviced_pb.DevicedService_ConnectClient
	session int64
}

func (self *moduleProxyStream) Send(val *any.Any) error {
	msg := &deviced_pb.ConnectResponse{
		SessionId: self.session,
		Kind:      deviced_pb.ConnectMessageKind_CONNECT_MESSAGE_KIND_USER,
		Union: &deviced_pb.ConnectResponse_StreamCall{
			StreamCall: &deviced_pb.StreamCallValue{
				Union: &deviced_pb.StreamCallValue_Value{
					Value: val,
				},
			},
		},
	}

	err := self.DevicedService_ConnectClient.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (self *moduleProxyStream) Recv() (*any.Any, error) {
	msg, err := self.DevicedService_ConnectClient.Recv()
	if err != nil {
		return nil, err
	}

	return msg.GetStreamCall().GetValue(), nil
}

func NewModuleProxyStream(stm deviced_pb.DevicedService_ConnectClient, session int64) ModuleProxyStream {
	return &moduleProxyStream{
		DevicedService_ConnectClient: stm,
		session:                      session,
	}
}

type ModuleProxy interface {
	UnaryCall(ctx context.Context, method string, req *any.Any) (*any.Any, error)
	StreamCall(ctx context.Context, method string, stm ModuleProxyStream) error
	Close() error
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
