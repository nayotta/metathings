package metathings_deviced_service

import (
	"context"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/deviced"
	log "github.com/sirupsen/logrus"
)

func (self *MetathingsDevicedService) ValidateCreateFirmwareHub(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() firmware_hub_getter {
				req := in.(*pb.CreateFirmwareHubRequest)
				return req
			},
		},
		identityd_validator.Invokers{},
	)
}

func (self *MetathingsDevicedService) AuthorizeCreateFirmwareHub(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, self.opt.Domain, "deviced:create_firmware_hub")
}

func (self *MetathingsDevicedService) CreateFirmwareHub(ctx context.Context, req *pb.CreateFirmwareHubRequest) (*pb.CreateFirmwareHubResponse, error) {
	var err error

	fh := req.GetFirmwareHub()

	fh_id_str := id_helper.NewId()
	if fh.GetId() != nil {
		fh_id_str = fh.GetId().GetValue()
	}

	logger := self.get_logger().WithFields(log.Fields{
		"id": fh_id_str,
	})

	fh_alias_str := fh.GetAlias().GetValue()
	if fh_alias_str == "" {
		fh_alias_str = fh_id_str
	}

	fh_description_str := fh.GetDescription().GetValue()

	fh_s := &storage.FirmwareHub{
		Id:          &fh_id_str,
		Alias:       &fh_alias_str,
		Description: &fh_description_str,
	}

	if fh_s, err = self.storage.CreateFirmwareHub(ctx, fh_s); err != nil {
		logger.WithError(err).Errorf("failed to create firmware hub in storage")
		return nil, self.ParseError(err)
	}

	res := &pb.CreateFirmwareHubResponse{
		FirmwareHub: copy_firmware_hub(fh_s),
	}

	logger.Infof("create firmware hub")

	return res, nil
}
