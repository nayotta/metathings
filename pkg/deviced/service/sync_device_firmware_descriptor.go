package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateSyncDeviceFirmwareDescriptor(ctx context.Context, in interface{}) error {
	return self.validator.Validate(identityd_validator.Providers{
		func() (policy_helper.Validator, device_getter) {
			req := in.(*pb.SyncDeviceFirmwareDescriptorRequest)
			return req, req
		},
	}, identityd_validator.Invokers{
		ensure_get_device_id,
	})

}

func (self *MetathingsDevicedService) AuthorizeSyncDeviceFirmwareDescriptor(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.SyncDeviceFirmwareDescriptorRequest).GetDevice().GetId().GetValue(), "deviced:sync_device_firmware_descriptor")
}

func (self *MetathingsDevicedService) SyncDeviceFirmwareDescriptor(ctx context.Context, req *pb.SyncDeviceFirmwareDescriptorRequest) (*empty.Empty, error) {
	var dev_s *storage.Device
	var err error

	dev_id_str := req.GetDevice().GetId().GetValue()
	logger := self.get_logger().WithField("device", dev_id_str)

	if dev_s, err = self.storage.GetDevice(ctx, dev_id_str); err != nil {
		logger.WithError(err).Errorf("failed to get devicein storage")
		return nil, self.ParseError(err)
	}

	if err = self.cc.SyncFirmware(ctx, dev_s); err != nil {
		logger.WithError(err).Errorf("failed to sync firmware descriptor")
		return nil, self.ParseError(err)
	}

	logger.Infof("sync firmware descriptor")

	return &empty.Empty{}, nil
}
