package metathings_deviced_service

import (
	"context"

	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateUnaryCall(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, device_getter) {
				req := in.(*pb.UnaryCallRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_device_id,
			ensure_device_online(self.storage),
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
	if dev_s, err = self.storage.GetDevice(dev_id_str); err != nil {
		self.logger.WithError(err).Debugf("failed to get device in storage")
		return nil, status.Convert(err).Err()
	}

	if val, err = self.cc.UnaryCall(dev_s, req.GetValue()); err != nil {
		self.logger.WithError(err).Debugf("failed to unray call")
		return nil, status.Convert(err).Err()
	}

	res := &pb.UnaryCallResponse{
		Device: &pb.Device{Id: dev_id_str},
		Value:  val,
	}

	self.logger.WithField("id", dev_id_str).Debugf("unary call")

	return res, nil
}
