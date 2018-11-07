package metathings_deviced_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) UnaryCall(ctx context.Context, req *pb.UnaryCallRequest) (*pb.UnaryCallResponse, error) {
	var dev_s *storage.Device
	var val *pb.UnaryCallValue
	var err error

	dev_id_str := req.GetDevice().GetId().GetValue()
	if dev_s, err = self.storage.GetDevice(dev_id_str); err != nil {
		self.logger.WithError(err).Debugf("failed to get device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if val, err = self.cc.UnaryCall(dev_s, req.GetValue()); err != nil {
		self.logger.WithError(err).Debugf("failed to unray call")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.UnaryCallResponse{
		Device: &pb.Device{Id: dev_id_str},
		Value:  val,
	}

	self.logger.WithField("id", dev_id_str).Debugf("unary call")

	return res, nil
}
