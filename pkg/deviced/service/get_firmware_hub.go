package metathings_deviced_service

import (
	"context"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateGetFirmwareHub(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, firmware_hub_getter) {
				req := in.(*pb.GetFirmwareHubRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_firmware_hub_id},
	)
}

func (self *MetathingsDevicedService) AuthorizeGetFirmwareHub(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.GetFirmwareHubRequest).GetFirmwareHub().GetId().GetValue(), "deviced:get_firmware_hub")
}

func (self *MetathingsDevicedService) GetFirmwareHub(ctx context.Context, req *pb.GetFirmwareHubRequest) (*pb.GetFirmwareHubResponse, error) {
	var frm_hub_s *storage.FirmwareHub
	var err error

	frm_hub_id_str := req.GetFirmwareHub().GetId().GetValue()
	logger := self.get_logger().WithField("firmware_hub", frm_hub_id_str)

	if frm_hub_s, err = self.storage.GetFirmwareHub(ctx, frm_hub_id_str); err != nil {
		logger.WithError(err).Errorf("failed to get firmware hub in storage")
		return nil, self.ParseError(err)
	}

	res := &pb.GetFirmwareHubResponse{
		FirmwareHub: copy_firmware_hub(frm_hub_s),
	}

	logger.Debugf("get firmware hub")

	return res, nil
}
