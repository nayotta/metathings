package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateAddDevicesToFirmwareHub(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, firmware_hub_getter) {
				req := in.(*pb.AddDevicesToFirmwareHubRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_firmware_hub_id},
	)
}

func (self *MetathingsDevicedService) AuthorizeAddDevicesToFirmwareHub(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.AddDevicesToFirmwareHubRequest).GetFirmwareHub().GetId().GetValue(), "deviced:add_devices_to_firmware_hub")
}

func (self *MetathingsDevicedService) AddDevicesToFirmwareHub(ctx context.Context, req *pb.AddDevicesToFirmwareHubRequest) (*empty.Empty, error) {
	var err error
	var devs_s []*storage.Device

	fh := req.GetFirmwareHub()
	fh_id_str := fh.GetId().GetValue()
	logger := self.get_logger().WithField("firmware_hub", fh_id_str)

	var dev_ids_str []string
	for _, dev := range req.GetDevices() {
		dev_ids_str = append(dev_ids_str, dev.GetId().GetValue())
	}

	if devs_s, err = self.storage.ListViewDevicesByFirmwareHubId(ctx, fh_id_str); err != nil {
		logger.WithError(err).Errorf("failed to list view devices by firmware hub id")
		return nil, self.ParseError(err)
	}

	var dev_ids_expect []string
	for _, dev_id_str := range dev_ids_str {
		exists := false
		for _, dev := range devs_s {
			if *dev.Id == dev_id_str {
				exists = true
				break
			}
		}

		if !exists {
			if err = self.storage.AddDeviceToFirmwareHub(ctx, fh_id_str, dev_id_str); err != nil {
				logger.WithError(err).Errorf("failed to add device to device in storage")
				return nil, self.ParseError(err)
			}
			dev_ids_expect = append(dev_ids_expect, dev_id_str)
		}
	}

	logger.WithFields(log.Fields{
		"devices": dev_ids_expect,
	}).Infof("add devices to firmware hub")

	return &empty.Empty{}, nil
}
