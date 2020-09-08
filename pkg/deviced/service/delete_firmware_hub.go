package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateDeleteFirmwareHub(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, firmware_hub_getter) {
				req := in.(*pb.DeleteFirmwareHubRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_firmware_hub_id,
		},
	)
}

func (self *MetathingsDevicedService) AuthorizeDeleteFirmwareHub(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.DeleteFirmwareHubRequest).GetFirmwareHub().GetId().GetValue(), "deviced:delete_firmware_hub")
}

func (self *MetathingsDevicedService) DeleteFirmwareHub(ctx context.Context, req *pb.DeleteFirmwareHubRequest) (*empty.Empty, error) {
	var fh *storage.FirmwareHub
	var err error

	fh_id_str := req.GetFirmwareHub().GetId().GetValue()
	logger := self.get_logger().WithField("id", fh_id_str)

	if fh, err = self.storage.GetFirmwareHub(ctx, fh_id_str); err != nil {
		logger.WithError(err).Errorf("failed to get firmware hub in storage")
		return nil, self.ParseError(err)
	}

	if err = self.storage.RemoveAllDevicesInFirmwareHub(ctx, fh_id_str); err != nil {
		logger.WithError(err).Errorf("failed to remove all devices in firmware hub in storage")
		return nil, self.ParseError(err)
	}

	for _, frm_desc := range fh.FirmwareDescriptors {
		if _, err = self.RemoveFirmwareDescriptorFromFirmwareHub(ctx, &pb.RemoveFirmwareDescriptorFromFirmwareHubRequest{
			FirmwareHub: &pb.OpFirmwareHub{
				Id: &wrappers.StringValue{Value: fh_id_str},
			},
			FirmwareDescriptor: &pb.OpFirmwareDescriptor{
				Id: &wrappers.StringValue{Value: *frm_desc.Id},
			},
		}); err != nil {
			return nil, err
		}
	}

	if err = self.storage.DeleteFirmwareHub(ctx, fh_id_str); err != nil {
		logger.WithError(err).Errorf("failed to delete firmware hub in storage")
		return nil, self.ParseError(err)
	}

	logger.Infof("delete firmware hub")

	return &empty.Empty{}, nil
}
