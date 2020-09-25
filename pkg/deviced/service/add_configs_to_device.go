package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDevicedService) AuthorizeAddConfigsToDevice(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.AddConfigsToDeviceRequest).GetDevice().GetId().GetValue(), "deviced:add_configs_to_device")
}

func (self *MetathingsDevicedService) AddConfigsToDevice(ctx context.Context, req *pb.AddConfigsToDeviceRequest) (*empty.Empty, error) {
	var err error

	dev := req.GetDevice()
	dev_id_str := dev.GetId().GetValue()
	logger := self.get_logger().WithField("device", dev_id_str)

	var cfg_ids_str []string
	for _, cfg := range req.GetConfigs() {
		cfg_ids_str = append(cfg_ids_str, cfg.GetId().GetValue())
	}

	cfgs_s, err := self.storage.ListConfigsByDeviceId(ctx, dev_id_str)
	if err != nil {
		logger.WithError(err).Errorf("failed to list configs by device")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var cfg_ids_expect []string
	for _, cfg_id_str := range cfg_ids_str {
		exists := false
		for _, cfg := range cfgs_s {
			if *cfg.Id == cfg_id_str {
				exists = true
				break
			}
		}

		if !exists {
			if err = self.storage.AddConfigToDevice(ctx, dev_id_str, cfg_id_str); err != nil {
				logger.WithError(err).Errorf("failed to add config to device in storage")
				return nil, status.Errorf(codes.Internal, err.Error())
			}
			cfg_ids_expect = append(cfg_ids_expect, cfg_id_str)
		}
	}

	logger.WithFields(log.Fields{
		"configs": cfg_ids_expect,
	}).Infof("add configs to device")

	return &empty.Empty{}, nil
}
