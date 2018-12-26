package metathings_deviced_service

import (
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	kind "github.com/nayotta/metathings/pkg/proto/constant/kind"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	// dispatch mqtt
	if *dev_s.Kind == kind.DeviceKind_DEVICE_KIND_SIMPLE.String() {
		if err = self.mqttBr.StreamCallForDeviced(*dev_s.Id, stm); err != nil {
			self.logger.WithError(err).Errorf("failed to simple kind stream call")
			return status.Errorf(codes.Internal, err.Error())
		}
	} else {
		if err = self.cc.StreamCall(dev_s, cfg, stm); err != nil {
			self.logger.WithError(err).Errorf("failed to stream call")
			return status.Errorf(codes.Internal, err.Error())
		}
	}

	self.logger.Debugf("stream call done")

	return nil
}
