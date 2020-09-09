package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateRemoveFirmwareDescriptorFromFirmwareHub(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, firmware_hub_getter) {
				req := in.(*pb.RemoveFirmwareDescriptorFromFirmwareHubRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_firmware_hub_id},
	)
}

func (self *MetathingsDevicedService) AuthorizeRemoveFirmwareDescriptorFromFirmwareHub(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.RemoveFirmwareDescriptorFromFirmwareHubRequest).GetFirmwareHub().GetId().GetValue(), "deviced:remove_firmware_descriptor_from_firmware_hub")
}

func (self *MetathingsDevicedService) RemoveFirmwareDescriptorFromFirmwareHub(ctx context.Context, req *pb.RemoveFirmwareDescriptorFromFirmwareHubRequest) (*empty.Empty, error) {
	var err error
	var fh_s *storage.FirmwareHub

	fh := req.GetFirmwareHub()
	fh_id_str := fh.GetId().GetValue()

	desc := req.GetFirmwareDescriptor()
	desc_id_str := desc.GetId().GetValue()

	logger := self.get_logger().WithFields(log.Fields{
		"firmware_hub":        fh_id_str,
		"firmware_descriptor": desc_id_str,
	})

	fh_s, err = self.storage.GetFirmwareHub(ctx, fh_id_str)

	exists := false
	for _, desc := range fh_s.FirmwareDescriptors {
		if desc_id_str == *desc.Id {
			exists = true
			break
		}
	}

	if exists {
		if err = self.storage.DeleteFirmwareDescriptor(ctx, desc_id_str); err != nil {
			logger.WithError(err).Errorf("failed to remove firmware descriptor from firmware hub in storage")
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	logger.Infof("remove firmware descriptor in firmware hub")

	return &empty.Empty{}, nil
}
