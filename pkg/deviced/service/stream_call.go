package metathings_deviced_service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	connection "github.com/nayotta/metathings/pkg/deviced/connection"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) StreamCall(stream pb.DevicedService_StreamCallServer) error {
	var conn connection.StreamConnection
	var err error

	if conn, err = self.cc.StreamCall(stream); err != nil {
		self.logger.WithError(err).Errorf("failed to stream call")
		return status.Errorf(codes.Internal, err.Error())
	}

	self.logger.Debugf("stream call")

	<-conn.Wait()
	self.logger.Debugf("stream call closed")

	if err = conn.Err(); err != nil {
		self.logger.WithError(err).Errorf("stream call error")
		return status.Errorf(codes.Internal, err.Error())
	}

	return nil
}
