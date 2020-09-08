package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateRemoveDevicesFromFirmwareHub(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, firmware_hub_getter) {
				req := in.(*pb.RemoveDevicesFromFirmwareHubRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_firmware_hub_id},
	)
}

func (self *MetathingsDevicedService) AuthorizeRemoveDevicesFromFirmwareHub(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.RemoveDevicesFromFirmwareHubRequest).GetFirmwareHub().GetId().GetValue(), "deviced:remove_devices_from_firmware_hub")
}

func (self *MetathingsDevicedService) RemoveDevicesFromFirmwareHub(ctx context.Context, req *pb.RemoveDevicesFromFirmwareHubRequest) (*empty.Empty, error) {
	var err error

	fh := req.GetFirmwareHub()
	fh_id_str := fh.GetId().GetValue()

	devs := req.GetDevices()

	logger := self.get_logger().WithFields(log.Fields{
		"firmware_hub": fh_id_str,
	})

	var dev_ids_str []string
	for _, dev := range devs {
		dev_id_str := dev.GetId().GetValue()
		if err = self.storage.RemoveDeviceFromFirmwareHub(ctx, fh_id_str, dev_id_str); err != nil {
			logger.WithError(err).WithField("device", dev_id_str).Errorf("failed to remove device from firmware hub")
			return nil, self.ParseError(err)
		}
		dev_ids_str = append(dev_ids_str, dev_id_str)
	}

	logger.WithField("devices", dev_ids_str).Infof("remove devices from firmware hub")

	return &empty.Empty{}, nil
}
