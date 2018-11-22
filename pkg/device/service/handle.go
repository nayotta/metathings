package metathings_device_service

import (
	"context"

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
		panic("unimplemented")
	default:
		panic("unimplemented")
	}
}

func (self *MetathingsDeviceServiceImpl) handle_user_unary_request(req *deviced_pb.ConnectRequest) error {
	req_val := req.GetUnaryCall()
	logger := self.logger.WithFields(log.Fields{
		"#session":   req.GetSessionId(),
		"#component": req_val.GetComponent(),
		"#name":      req_val.GetName(),
		"#method":    req_val.GetMethod(),
	})

	mdl, err := self.mdl_db.Lookup(req_val.GetComponent().GetValue(), req_val.GetName().GetValue())
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
		SessionId: req.GetSessionId().GetValue(),
		Kind:      req.GetKind(),
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
