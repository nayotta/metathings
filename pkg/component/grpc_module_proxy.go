package metathings_component

import (
	"context"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	pb "github.com/nayotta/metathings/pkg/proto/component"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type GrpcModuleServiceClientFactory interface {
	NewModuleServiceClient(opts ...grpc.DialOption) (pb.ModuleServiceClient, client_helper.CloseFn, error)
}

type GrpcModuleServiceClientFactoryImpl struct {
	Address string
}

func (self *GrpcModuleServiceClientFactoryImpl) NewModuleServiceClient(opts ...grpc.DialOption) (pb.ModuleServiceClient, client_helper.CloseFn, error) {
	opts = append([]grpc.DialOption{grpc.WithInsecure()}, opts...)
	conn, err := grpc.Dial(self.Address, opts...)
	if err != nil {
		return nil, nil, err
	}

	return pb.NewModuleServiceClient(conn), conn.Close, nil
}

func NewGrpcModuleServiceClientFactory(addr string) GrpcModuleServiceClientFactory {
	return &GrpcModuleServiceClientFactoryImpl{Address: addr}
}

type GrpcModuleProxy struct {
	cli_fty GrpcModuleServiceClientFactory
	logger  log.FieldLogger
}

func (self *GrpcModuleProxy) UnaryCall(ctx context.Context, method string, value *any.Any) (*any.Any, error) {
	cli, cfn, err := self.cli_fty.NewModuleServiceClient()
	if err != nil {
		return nil, err
	}
	defer cfn()

	req := &pb.UnaryCallRequest{
		Method: &wrappers.StringValue{Value: method},
		Value:  value,
	}
	res, err := cli.UnaryCall(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.GetValue(), nil
}

func (self *GrpcModuleProxy) StreamCall(ctx context.Context, method string, upstm ModuleProxyStream) error {
	cli, cfn, err := self.cli_fty.NewModuleServiceClient()
	if err != nil {
		self.logger.WithError(err).Debugf("failed to new module service client")
		return err
	}
	defer cfn()

	downstm, err := cli.StreamCall(ctx)
	if err != nil {
		self.logger.WithError(err).Debugf("failed to start stream call")
		return err
	}
	self.logger.Debugf("connect to module service(downstream)")

	err = self.init_downstream(downstm, method)
	if err != nil {
		self.logger.WithError(err).Debugf("failed to initial downstream")
		return err
	}
	self.logger.Debugf("downstream initialized")

	up2down_wait := make(chan bool)
	down2up_wait := make(chan bool)
	go self.stm_up2down(upstm, downstm, up2down_wait)
	go self.stm_down2up(upstm, downstm, down2up_wait)

	self.logger.Debugf("stream call started")
	select {
	case <-up2down_wait:
	case <-down2up_wait:
	}

	self.logger.Debugf("stream call done")

	return nil
}

func (self *GrpcModuleProxy) recv_cfg_msg(stm deviced_pb.DevicedService_ConnectClient) (*deviced_pb.ConnectRequest, error) {
	req, err := stm.Recv()
	if err != nil {
		return nil, err
	}

	if req.GetKind() != deviced_pb.ConnectMessageKind_CONNECT_MESSAGE_KIND_USER {
		return nil, ErrInvalidArguments
	}

	if req.GetStreamCall() == nil {
		return nil, ErrInvalidArguments
	}

	return req, nil
}

func (self *GrpcModuleProxy) init_downstream(stm pb.ModuleService_StreamCallClient, method string) error {
	var err error

	cfg := &pb.StreamCallRequest{
		Request: &pb.StreamCallRequest_Config{
			Config: &pb.StreamCallConfigRequest{
				Method: &wrappers.StringValue{Value: method},
			},
		},
	}

	if err = stm.Send(cfg); err != nil {
		return err
	}

	return nil
}

func (self *GrpcModuleProxy) stm_up2down(upstm ModuleProxyStream, downstm pb.ModuleService_StreamCallClient, wait chan bool) {
	var val *any.Any
	var err error

	defer close(wait)
	for {
		if val, err = upstm.Recv(); err != nil {
			return
		}

		downreq := &pb.StreamCallRequest{
			Request: &pb.StreamCallRequest_Data{
				Data: &pb.StreamCallDataRequest{
					Value: val,
				},
			},
		}

		if err = downstm.Send(downreq); err != nil {
			return
		}
	}
}

func (self *GrpcModuleProxy) stm_down2up(upstm ModuleProxyStream, downstm pb.ModuleService_StreamCallClient, wait chan bool) {
	var downres *pb.StreamCallResponse
	var err error

	defer close(wait)
	for {
		if downres, err = downstm.Recv(); err != nil {
			return
		}

		if err = upstm.Send(downres.GetData().GetValue()); err != nil {
			return
		}
	}
}

type GrpcModuleProxyFactory struct{}

func (self *GrpcModuleProxyFactory) NewModuleProxy(args ...interface{}) (ModuleProxy, error) {
	p := &GrpcModuleProxy{}

	if err := opt_helper.Setopt(map[string]func(key string, val interface{}) error{
		"logger": func(key string, val interface{}) error {
			var ok bool
			if p.logger, ok = val.(log.FieldLogger); !ok {
				return opt_helper.ErrInvalidArguments
			}
			return nil
		},
		"client_factory": func(key string, val interface{}) error {
			var ok bool
			if p.cli_fty, ok = val.(GrpcModuleServiceClientFactory); !ok {
				return opt_helper.ErrInvalidArguments
			}
			return nil
		},
	})(args...); err != nil {
		return nil, err
	}

	return p, nil
}

func init() {
	register_module_proxy_factory("grpc", new(GrpcModuleProxyFactory))
}
