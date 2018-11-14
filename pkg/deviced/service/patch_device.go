package metathings_deviced_service

import (
	"context"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (self *MetathingsDevicedService) ValidatePatchDevice(ctx context.Context, in interface{}) error {
	return self.validate_chain(
		[]interface{}{
			func() (policy_helper.Validator, get_devicer) {
				req := in.(*pb.PatchDeviceRequest)
				return req, req
			},
		},
		[]interface{}{ensure_get_device_id},
	)
}

func (self *MetathingsDevicedService) AuthorizePatchDevice(ctx context.Context, in interface{}) error {
	return self.enforce(ctx, in.(*pb.PatchDeviceRequest).GetDevice().GetId().GetValue(), "patch_device")
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
