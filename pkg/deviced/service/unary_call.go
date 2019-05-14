package metathings_deviced_service

import (
	"context"

	deviced_helper "github.com/nayotta/metathings/pkg/deviced/helper"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	kind "github.com/nayotta/metathings/pkg/proto/constant/kind"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
	"google.golang.org/grpc/status"
)

func (self *MetathingsDevicedService) UnaryCall(ctx context.Context, req *pb.UnaryCallRequest) (*pb.UnaryCallResponse, error) {
	var dev_s *storage.Device
	var val *pb.UnaryCallValue
	var err error

	dev_id_str := req.GetDevice().GetId().GetValue()
	if dev_s, err = self.storage.GetDevice(dev_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get device in storage")
		return nil, status.Convert(err).Err()
	}

	// dispatch mqtt
	if *dev_s.Kind == deviced_helper.DEVICE_KIND_ENUMER.ToString(kind.DeviceKind_DEVICE_KIND_SIMPLE) {
		val, err = self.mqttBr.UnaryCallForDeviced(dev_s, req.GetValue())
		if err != nil {
			self.logger.WithError(err).Errorf("failed to simple kind unary call")
			return nil, status.Convert(err).Err()
		}
	} else {
		if val, err = self.cc.UnaryCall(dev_s, req.GetValue()); err != nil {
			self.logger.WithError(err).Errorf("failed to unray call")
			return nil, status.Convert(err).Err()
		}
	}

	res := &pb.UnaryCallResponse{
		Device: &pb.Device{Id: dev_id_str},
		Value:  val,
	}

	self.logger.WithField("id", dev_id_str).Debugf("unary call")

	return res, nil
}
