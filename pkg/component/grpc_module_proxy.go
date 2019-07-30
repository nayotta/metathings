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

func (self *GrpcModuleProxy) Close() error { return nil }

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

	logger := self.logger.WithFields(log.Fields{
		"dir": "up->down",
	})

	defer close(wait)
	for epoch := uint64(0); ; epoch++ {
		logger := logger.WithFields(log.Fields{
			"epoch": epoch,
		})

		if val, err = upstm.Recv(); err != nil {
			logger.WithError(err).Debugf("failed to recv msg from upstm")
			return
		}
		logger.Debugf("recv msg from upstm")

		downreq := &pb.StreamCallRequest{
			Request: &pb.StreamCallRequest_Data{
				Data: &pb.StreamCallDataRequest{
					Value: val,
				},
			},
		}

		if err = downstm.Send(downreq); err != nil {
			logger.WithError(err).Debugf("failed to send msg to downstm")
			return
		}
		logger.Debugf("send msg to downstm")
	}
}

func (self *GrpcModuleProxy) stm_down2up(upstm ModuleProxyStream, downstm pb.ModuleService_StreamCallClient, wait chan bool) {
	var downres *pb.StreamCallResponse
	var err error

	logger := self.logger.WithFields(log.Fields{
		"dir": "up<-down",
	})

	defer close(wait)
	for epoch := uint64(0); ; epoch++ {
		logger := logger.WithFields(log.Fields{
			"epoch": epoch,
		})

		if downres, err = downstm.Recv(); err != nil {
			logger.WithError(err).Debugf("failed to recv msg from downstm")
			return
		}
		logger.Debugf("recv msg from downstm")

		if err = upstm.Send(downres.GetData().GetValue()); err != nil {
			logger.Debugf("failed to send msg to upstm")
			return
		}
		logger.Debugf("send msg to upstm")
	}
}

type GrpcModuleProxyFactory struct{}

func (self *GrpcModuleProxyFactory) NewModuleProxy(args ...interface{}) (ModuleProxy, error) {
	p := &GrpcModuleProxy{}

	if err := opt_helper.Setopt(map[string]func(key string, val interface{}) error{
		"logger": opt_helper.ToLogger(&p.logger),
		"client_factory": func(key string, val interface{}) error {
			var ok bool
			if p.cli_fty, ok = val.(GrpcModuleServiceClientFactory); !ok {
				return opt_helper.InvalidArgument("client_factory")
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
