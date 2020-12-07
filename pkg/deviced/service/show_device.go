package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDevicedService) ShowDevice(ctx context.Context, _ *empty.Empty) (*pb.ShowDeviceResponse, error) {
	var dev_s *storage.Device
	var err error

	logger := self.get_logger()

	if dev_s, err = self.get_device_by_context(ctx); err != nil {
		logger.WithError(err).Errorf("failed to get device by context in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	logger = logger.WithField("device", *dev_s.Id)

	res := &pb.ShowDeviceResponse{
		Device: copy_device(dev_s),
	}

	logger.Debugf("show device")

	return res, nil
}
