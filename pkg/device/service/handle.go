package metathings_device_service

import (
	"context"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
	log "github.com/sirupsen/logrus"
)

func (self *MetathingsDeviceServiceImpl) handle(req *deviced_pb.ConnectRequest) error {
	switch req.Kind {
	case deviced_pb.ConnectMessageKind_CONNECT_MESSAGE_KIND_SYSTEM:
		return self.handle_system_request(req)
	case deviced_pb.ConnectMessageKind_CONNECT_MESSAGE_KIND_USER:
		return self.handle_user_request(req)
	default:
		self.logger.Warningf("unexpected request data")
		return nil
	}
}

func (self *MetathingsDeviceServiceImpl) handle_system_request(req *deviced_pb.ConnectRequest) error {
	panic("unimplemented")
}

func (self *MetathingsDeviceServiceImpl) handle_user_request(req *deviced_pb.ConnectRequest) error {
	switch req.Union.(type) {
	case *deviced_pb.ConnectRequest_UnaryCall:
		return self.handle_user_unary_request(req)
	case *deviced_pb.ConnectRequest_StreamCall:
		return self.handle_user_stream_request(req)
	default:
		self.logger.WithField("union", req.Union).Debugf("unsupported union type")
		panic("unimplemented")
	}
}

func (self *MetathingsDeviceServiceImpl) handle_user_unary_request(req *deviced_pb.ConnectRequest) error {
	req_val := req.GetUnaryCall()
	sess := req.GetSessionId().GetValue()
	kind := req.GetKind()
	component := req_val.GetComponent().GetValue()
	name := req_val.GetName().GetValue()
	method := req_val.GetName().GetValue()

	logger := self.logger.WithFields(log.Fields{
		"#session":   sess,
		"#component": component,
		"#name":      name,
		"#method":    method,
	})

	mdl, err := self.mdl_db.Lookup(component, name)
	if err != nil {
		logger.WithError(err).Debugf("failed to lookup module in database")
		return err
	}
	logger.Debugf("lookup module in storage")

	res_val, err := mdl.UnaryCall(context.Background(), req_val)
	if err != nil {
		logger.WithError(err).Debugf("failed to unary call in module")
		return err
	}
	logger.Debugf("unary call in module")

	res := &deviced_pb.ConnectResponse{
		SessionId: sess,
		Kind:      kind,
		Union: &deviced_pb.ConnectResponse_UnaryCall{
			UnaryCall: res_val,
		},
	}

	err = self.conn_stm.Send(res)
	if err != nil {
		logger.Debugf("failed to send msg")
		return err
	}
	logger.Debugf("send msg")

	return nil
}

func (self *MetathingsDeviceServiceImpl) handle_user_stream_request(req *deviced_pb.ConnectRequest) error {
	var cli deviced_pb.DevicedServiceClient
	var cfn client_helper.CloseFn
	var stm deviced_pb.DevicedService_ConnectClient
	var err error

	req_val := req.GetStreamCall()
	sess := req.GetSessionId().GetValue()

	if cli, cfn, err = self.cli_fty.NewDevicedServiceClient(); err != nil {
		self.logger.WithError(err).Debugf("failed to new deviced service client")
		return err
	}

	ctx := context_helper.NewOutgoingContext(
		context.Background(),
		context_helper.WithTokenOp(self.tknr.GetToken()),
		context_helper.WithSessionOp(sess),
	)

	if stm, err = cli.Connect(ctx); err != nil {
		self.logger.WithError(err).Debugf("failed to connect to deviced service")
		return err
	}

	panic("unimplemented")
}
