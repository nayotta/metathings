package metathings_deviced_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) ValidatePatchDevice(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, device_getter) {
				req := in.(*pb.PatchDeviceRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_device_id},
	)
}

func (self *MetathingsDevicedService) AuthorizePatchDevice(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.PatchDeviceRequest).GetDevice().GetId().GetValue(), "deviced:patch_device")
}

func (self *MetathingsDevicedService) PatchDevice(ctx context.Context, req *pb.PatchDeviceRequest) (*pb.PatchDeviceResponse, error) {
	dev_s := &storage.Device{}
	var err error

	dev := req.GetDevice()
	dev_id_str := dev.GetId().GetValue()

	alias := dev.GetAlias()
	if alias != nil {
		dev_s.Alias = &alias.Value
	}

	if dev_s, err = self.storage.PatchDevice(dev_id_str, dev_s); err != nil {
		self.logger.WithError(err).Errorf("failed to patch device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.PatchDeviceResponse{
		Device: copy_device(dev_s),
	}

	self.logger.WithField("id", dev_id_str).Infof("patch device")

	return res, nil
}
