package metathings_component

import (
	"context"

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

func (self *GrpcModuleProxy) UnaryCall(ctx context.Context, req *deviced_pb.OpUnaryCallValue) (*deviced_pb.UnaryCallValue, error) {
	method := req.GetMethod().GetValue()
	value := req.GetValue()

	cli, cfn, err := self.cli_fty.NewModuleServiceClient()
	if err != nil {
		return nil, err
	}
	defer cfn()

	mdl_req := &pb.UnaryCallRequest{
		Method: &wrappers.StringValue{Value: method},
		Value:  value,
	}
	mdl_res, err := cli.UnaryCall(ctx, mdl_req)
	if err != nil {
		return nil, err
	}

	res := &deviced_pb.UnaryCallValue{
		Component: req.GetComponent().GetValue(),
		Name:      req.GetName().GetValue(),
		Method:    req.GetMethod().GetValue(),
		Value:     mdl_res.GetValue(),
	}

	return res, nil
}

func (self *GrpcModuleProxy) StreamCall(ctx context.Context, upstm deviced_pb.DevicedService_ConnectClient) error {
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

	cfg_req, err := self.recv_cfg_msg(upstm)
	if err != nil {
		self.logger.WithError(err).Debugf("failed to recv config msg")
		return err
	}

	err = self.init_downstream(downstm, cfg_req.GetStreamCall().GetConfig())
	if err != nil {
		self.logger.WithError(err).Debugf("failed to initial downstream")
		return err
	}

	up2down_wait := make(chan bool)
	down2up_wait := make(chan bool)
	go self.stm_up2down(upstm, downstm, up2down_wait)
	go self.stm_down2up(upstm, downstm, cfg_req, down2up_wait)

	select {
	case <-up2down_wait:
	case <-down2up_wait:
	}

	panic("unimplemented")
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

func (self *GrpcModuleProxy) init_downstream(stm pb.ModuleService_StreamCallClient, cfg *deviced_pb.OpStreamCallConfig) error {
	var err error

	cfg_req := &pb.StreamCallRequest{
		Request: &pb.StreamCallRequest_Config{
			Config: &pb.StreamCallConfigRequest{
				Method: cfg.GetMethod(),
			},
		},
	}

	if err = stm.Send(cfg_req); err != nil {
		return err
	}

	return nil
}

func (self *GrpcModuleProxy) stm_up2down(upstm deviced_pb.DevicedService_ConnectClient, downstm pb.ModuleService_StreamCallClient, wait chan bool) {
	var upreq *deviced_pb.ConnectRequest
	var err error

	defer close(wait)
	for {
		if upreq, err = upstm.Recv(); err != nil {
			return
		}

		downreq := &pb.StreamCallRequest{
			Request: &pb.StreamCallRequest_Data{
				Data: &pb.StreamCallDataRequest{
					Value: upreq.GetStreamCall().GetValue(),
				},
			},
		}

		if err = downstm.Send(downreq); err != nil {
			return
		}
	}
}

func (self *GrpcModuleProxy) stm_down2up(upstm deviced_pb.DevicedService_ConnectClient, downstm pb.ModuleService_StreamCallClient, cfg_req *deviced_pb.ConnectRequest, wait chan bool) {
	var downres *pb.StreamCallResponse
	var err error

	kind := cfg_req.GetKind()
	sess := cfg_req.GetSessionId().GetValue()

	defer close(wait)
	for {
		if downres, err = downstm.Recv(); err != nil {
			return
		}

		upres := &deviced_pb.ConnectResponse{
			Kind:      kind,
			SessionId: sess,
			Union: &deviced_pb.ConnectResponse_StreamCall{
				StreamCall: &deviced_pb.StreamCallValue{
					Union: &deviced_pb.StreamCallValue_Value{
						Value: downres.GetData().GetValue(),
					},
				},
			},
		}

		if err = upstm.Send(upres); err != nil {
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
