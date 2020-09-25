package metathings_deviced_service

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDevicedService) StreamCall(stm pb.DevicedService_StreamCallServer) error {
	var req *pb.StreamCallRequest
	var dev_s *storage.Device
	var err error

	logger := self.get_logger()

	if req, err = stm.Recv(); err != nil {
		logger.WithError(err).Errorf("failed to recv config msg")
		return status.Errorf(codes.Internal, err.Error())
	}

	cfg := req.GetValue().GetConfig()
	dev_id_str := req.GetDevice().GetId().GetValue()
	mdl_name_str := cfg.GetName().GetValue()
	meth_str := cfg.GetMethod().GetValue()
	logger = self.get_logger().WithFields(log.Fields{
		"device": dev_id_str,
		"module": mdl_name_str,
		"method": meth_str,
	})

	logger.Debugf("recv config msg")

	if dev_s, err = self.storage.GetDevice(stm.Context(), dev_id_str); err != nil {
		logger.WithError(err).Errorf("failed to get device in storage")
		return status.Errorf(codes.Internal, err.Error())
	}

	if err = self.cc.StreamCall(stm.Context(), dev_s, cfg, stm); err != nil {
		logger.WithError(err).Errorf("failed to stream call")
		return status.Errorf(codes.Internal, err.Error())
	}

	logger.Debugf("stream call done")

	return nil
}
