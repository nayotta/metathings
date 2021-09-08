package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"

	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDevicedService) AuthorizeSetDeviceFirmwareDescriptor(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.SetDeviceFirmwareDescriptorRequest).GetDevice().GetId().GetValue(), "deviced:set_device_firmware_descriptor")
}

func (self *MetathingsDevicedService) ValidateSetDeviceFirmwareDescriptor(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (device_getter, firmware_descriptor_getter) {
				req := in.(*pb.SetDeviceFirmwareDescriptorRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_device_id,
			ensure_get_firmware_descriptor_id,
			ensure_firmware_hub_contains_device_and_firmware_descriptor(ctx, self.storage),
		},
	)
}

func (self *MetathingsDevicedService) SetDeviceFirmwareDescriptor(ctx context.Context, req *pb.SetDeviceFirmwareDescriptorRequest) (*empty.Empty, error) {
	var err error

	dev_id_str := req.GetDevice().GetId().GetValue()
	desc_id_str := req.GetFirmwareDescriptor().GetId().GetValue()
	logger := self.get_logger().WithFields(log.Fields{
		"device":              dev_id_str,
		"firmware_descriptor": desc_id_str,
	})

	if err = self.storage.UnsetDeviceFirmwareDescriptor(ctx, dev_id_str); err != nil {
		logger.WithError(err).Errorf("failed to cleanup device firmware descriptor")
		return nil, self.ParseError(err)
	}

	if err = self.storage.SetDeviceFirmwareDescriptor(ctx, dev_id_str, desc_id_str); err != nil {
		logger.WithError(err).Errorf("failed to set device firmware descriptor in storage")
		return nil, self.ParseError(err)
	}

	logger.Infof("set device firmware descriptor")

	return &empty.Empty{}, nil
}
