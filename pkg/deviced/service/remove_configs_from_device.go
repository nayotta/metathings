package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) AuthorizeRemoveConfigsFromDevice(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.RemoveConfigsFromDeviceRequest).GetDevice().GetId().GetValue(), "deviced:remove_configs_from_device")
}

func (self *MetathingsDevicedService) RemoveConfigsFromDevice(ctx context.Context, req *pb.RemoveConfigsFromDeviceRequest) (*empty.Empty, error) {
	var err error

	dev := req.GetDevice()
	dev_id_str := dev.GetId().GetValue()

	cfgs := req.GetConfigs()

	logger := self.get_logger().WithField("device", dev_id_str)

	var cfg_ids_str []string
	for _, cfg := range cfgs {
		cfg_id_str := cfg.GetId().GetValue()
		if err = self.storage.RemoveConfigFromDevice(ctx, dev_id_str, cfg_id_str); err != nil {
			logger.WithError(err).WithField("config", cfg_id_str).Errorf("failed to remove config from device")
			return nil, status.Errorf(codes.Internal, err.Error())
		}
		cfg_ids_str = append(cfg_ids_str, cfg_id_str)
	}

	logger.WithField("configs", cfg_ids_str).Infof("remove configs from device")

	return &empty.Empty{}, nil
}
