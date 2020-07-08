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

func (self *MetathingsDevicedService) ValidatePatchFirmwareHub(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, firmware_hub_getter) {
				req := in.(*pb.PatchFirmwareHubRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_firmware_hub_id},
	)
}

func (self *MetathingsDevicedService) AuthorizePatchFirmwareHub(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.PatchFirmwareHubRequest).GetFirmwareHub().GetId().GetValue(), "deviced:patch_firmware_hub")
}

func (self *MetathingsDevicedService) PatchFirmwareHub(ctx context.Context, req *pb.PatchFirmwareHubRequest) (*pb.PatchFirmwareHubResponse, error) {
	var err error
	frm_hub_s := &storage.FirmwareHub{}

	frm_hub := req.GetFirmwareHub()
	frm_hub_id_str := frm_hub.GetId().GetValue()
	logger := self.logger.WithField("firmware_hub", frm_hub_id_str)

	if alias := frm_hub.GetAlias(); alias != nil {
		frm_hub_s.Alias = &alias.Value
	}

	if description := frm_hub.GetDescription(); description != nil {
		frm_hub_s.Description = &description.Value
	}

	if frm_hub_s, err = self.storage.PatchFirmwareHub(ctx, frm_hub_id_str, frm_hub_s); err != nil {
		logger.WithError(err).Errorf("failed to patch firmware hub in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.PatchFirmwareHubResponse{
		FirmwareHub: copy_firmware_hub(frm_hub_s),
	}

	logger.Infof("patch firmware hub")

	return res, nil
}
