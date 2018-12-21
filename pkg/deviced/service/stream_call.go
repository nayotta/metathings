package metathings_deviced_service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) StreamCall(stm pb.DevicedService_StreamCallServer) error {
	var req *pb.StreamCallRequest
	var dev_s *storage.Device
	var err error

	if req, err = stm.Recv(); err != nil {
		self.logger.WithError(err).Errorf("failed to recv config msg")
		return status.Errorf(codes.Internal, err.Error())
	}
	self.logger.Debugf("recv config msg")

	cfg := req.GetValue().GetConfig()
	dev_id_str := req.GetDevice().GetId().GetValue()
	if dev_s, err = self.storage.GetDevice(dev_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get device in storage")
		return status.Errorf(codes.Internal, err.Error())
	}

	if err = self.cc.StreamCall(dev_s, cfg, stm); err != nil {
		self.logger.WithError(err).Errorf("failed to stream call")
		return status.Errorf(codes.Internal, err.Error())
	}

	self.logger.Debugf("stream call done")

	return nil
}
