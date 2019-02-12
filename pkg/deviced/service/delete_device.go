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
			func() (policy_helper.Validator, get_devicer) {
				req := in.(*pb.DeleteDeviceRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_device_id},
	)
}

func (self *MetathingsDevicedService) AuthorizeDeleteDevice(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.DeleteDeviceRequest).GetDevice().GetId().GetValue(), "delete_device")
}

func (self *MetathingsDevicedService) DeleteDevice(ctx context.Context, req *pb.DeleteDeviceRequest) (*empty.Empty, error) {
	var dev *storage.Device
	var err error

	dev_id_str := req.GetDevice().GetId().GetValue()
	if dev, err = self.storage.GetDevice(dev_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	for _, m := range dev.Modules {
		mdl_id_str := *m.Id
		if err = self.enforcer.RemoveObjectFromKind(mdl_id_str, KIND_MODULE); err != nil {
			self.logger.WithError(err).Warningf("failed to remove module from kind in enforcer")
		}
		if err = self.storage.DeleteModule(mdl_id_str); err != nil {
			self.logger.WithError(err).WithField("id", mdl_id_str).Warningf("failed to delete module in storage")
		}
	}

	if err = self.enforcer.RemoveObjectFromKind(dev_id_str, KIND_DEVICE); err != nil {
		self.logger.WithError(err).Warningf("failed to remove device from kind in enforcer")
	}
	if err = self.storage.DeleteDevice(dev_id_str); err != nil {
		self.logger.WithError(err).Debugf("failed to delete device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithField("id", dev_id_str).Infof("delete device")

	return &empty.Empty{}, nil
}
