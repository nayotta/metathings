package metathings_deviced_service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) StreamCall(stream pb.DevicedService_StreamCallServer) error {
	var err error

	if err = self.cc.StreamCall(stream); err != nil {
		self.logger.WithError(err).Errorf("failed to stream call")
		return status.Errorf(codes.Internal, err.Error())
	}

	return nil
}
