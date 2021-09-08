package metathings_deviced_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateGetDevice(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() device_getter {
				req := in.(*pb.GetDeviceRequest)
				return req
			},
		},
		identityd_validator.Invokers{ensure_get_device_id},
	)
}

func (self *MetathingsDevicedService) AuthorizeGetDevice(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.GetDeviceRequest).GetDevice().GetId().GetValue(), "deviced:get_device")
}

func (self *MetathingsDevicedService) GetDevice(ctx context.Context, req *pb.GetDeviceRequest) (*pb.GetDeviceResponse, error) {
	var dev_s *storage.Device
	var err error

	dev_id_str := req.GetDevice().GetId().GetValue()
	logger := self.get_logger().WithField("device", dev_id_str)

	if dev_s, err = self.storage.GetDevice(ctx, dev_id_str); err != nil {
		logger.WithError(err).Errorf("failed to get device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.GetDeviceResponse{
		Device: copy_device(dev_s),
	}

	logger.Debugf("get device")

	return res, nil
}
