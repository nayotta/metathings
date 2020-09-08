package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateDeleteDevice(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, device_getter) {
				req := in.(*pb.DeleteDeviceRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_device_id},
	)
}

func (self *MetathingsDevicedService) AuthorizeDeleteDevice(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.DeleteDeviceRequest).GetDevice().GetId().GetValue(), "deviced:delete_device")
}

func (self *MetathingsDevicedService) DeleteDevice(ctx context.Context, req *pb.DeleteDeviceRequest) (*empty.Empty, error) {
	var dev *storage.Device
	var err error

	dev_id_str := req.GetDevice().GetId().GetValue()
	logger := self.get_logger().WithField("device", dev_id_str)

	if dev, err = self.storage.GetDevice(ctx, dev_id_str); err != nil {
		logger.WithError(err).Errorf("failed to get device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	for _, m := range dev.Modules {
		mdl_id_str := *m.Id
		if err = self.storage.DeleteModule(ctx, mdl_id_str); err != nil {
			logger.WithError(err).WithField("id", mdl_id_str).Warningf("failed to delete module in storage")
		}
	}

	for _, f := range dev.Flows {
		flw_id_str := *f.Id
		if err = self.storage.DeleteFlow(ctx, flw_id_str); err != nil {
			logger.WithError(err).WithField("id", flw_id_str).Warningf("failed to delete flow in storage")
		}
	}

	if err = self.storage.DeleteDevice(ctx, dev_id_str); err != nil {
		logger.WithError(err).Debugf("failed to delete device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	logger.WithField("id", dev_id_str).Infof("delete device")

	return &empty.Empty{}, nil
}
