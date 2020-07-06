package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) AuthorizeSetDeviceFirmwareDescriptor(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.SetDeviceFirmwareDescriptorRequest).GetDevice().GetId().GetValue(), "deviced:set_device_firmware_descriptor")
}

func (self *MetathingsDevicedService) ValidateSetDeviceFirmwareDescriptor(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, device_getter) {
				req := in.(*pb.SetDeviceFirmwareDescriptorRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_device_id},
	)
}

func (self *MetathingsDevicedService) SetDeviceFirmwareDescriptor(ctx context.Context, req *pb.SetDeviceFirmwareDescriptorRequest) (*empty.Empty, error) {
	var err error

	dev_id_str := req.GetDevice().GetId().GetValue()
	desc_id_str := req.GetFirmwareDescriptor().GetId().GetValue()
	logger := self.logger.WithFields(log.Fields{
		"device":              dev_id_str,
		"firmware_descriptor": desc_id_str,
	})

	if err = self.storage.SetDeviceFirmwareDescriptor(ctx, dev_id_str, desc_id_str); err != nil {
		logger.WithError(err).Errorf("failed to set device firmware descriptor in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	logger.Infof("set device firmware descriptor")

	return &empty.Empty{}, nil
}
