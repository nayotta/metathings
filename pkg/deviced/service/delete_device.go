package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) DeleteDevice(ctx context.Context, req *pb.DeleteDeviceRequest) (*empty.Empty, error) {
	var dev *storage.Device
	var err error

	dev_id_str := req.GetDevice().GetId().GetValue()
	if dev, err = self.storage.GetDevice(dev_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	for _, m := range dev.Modules {
		if err = self.storage.DeleteModule(*m.Id); err != nil {
			self.logger.WithError(err).WithField("id", *m.Id).Warningf("failed to delete module in storage")
		}
	}

	if err = self.storage.DeleteDevice(dev_id_str); err != nil {
		self.logger.WithError(err).Debugf("failed to delete device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithField("id", dev_id_str).Infof("delete device")

	return &empty.Empty{}, nil
}
