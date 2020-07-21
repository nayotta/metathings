package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) ShowDeviceFirmwareDescriptor(ctx context.Context, _ *empty.Empty) (*pb.ShowDeviceFirmwareDescriptorResponse, error) {
	dev_id_str := self.get_device_id_from_context(ctx)
	logger := self.logger.WithField("device", dev_id_str)

	fd_s, err := self.get_device_firmware_descriptor(ctx, dev_id_str)
	if err != nil {
		return nil, err
	}

	res := &pb.ShowDeviceFirmwareDescriptorResponse{
		FirmwareDescriptor: copy_firmware_descriptor(fd_s),
	}

	logger.Debugf("show device firmware descriptor")

	return res, nil
}
