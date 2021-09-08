package metathings_deviced_service

import (
	"context"

	log "github.com/sirupsen/logrus"

	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateUnaryCall(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() device_getter {
				req := in.(*pb.UnaryCallRequest)
				return req
			},
		},
		identityd_validator.Invokers{
			ensure_get_device_id,
			ensure_device_online(ctx, self.storage),
		},
	)
}

func (self *MetathingsDevicedService) AuthorizeUnaryCall(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.UnaryCallRequest).GetDevice().GetId().GetValue(), "deviced:unary_call")
}

func (self *MetathingsDevicedService) UnaryCall(ctx context.Context, req *pb.UnaryCallRequest) (*pb.UnaryCallResponse, error) {
	var dev_s *storage.Device
	var val *pb.UnaryCallValue
	var err error

	dev_id_str := req.GetDevice().GetId().GetValue()

	logger := self.get_logger().WithFields(log.Fields{
		"device": dev_id_str,
	})
	if dev_s, err = self.storage.GetDevice(ctx, dev_id_str); err != nil {
		logger.WithError(err).Errorf("failed to get device in storage")
		return nil, self.ParseError(err)
	}

	mdl_name_str := req.GetValue().GetName().GetValue()
	meth_str := req.GetValue().GetMethod().GetValue()
	logger = logger.WithFields(log.Fields{
		"module": mdl_name_str,
		"method": meth_str,
	})

	if val, err = self.cc.UnaryCall(ctx, dev_s, req.GetValue()); err != nil {
		logger.WithError(err).Errorf("failed to unray call")
		return nil, self.ParseError(err)
	}

	res := &pb.UnaryCallResponse{
		Device: &pb.Device{Id: dev_id_str},
		Value:  val,
	}

	logger.Debugf("unary call")

	return res, nil
}
