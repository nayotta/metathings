package metathings_deviced_service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	session_helper "github.com/nayotta/metathings/pkg/common/session"
	connection "github.com/nayotta/metathings/pkg/deviced/connection"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
	log "github.com/sirupsen/logrus"
)

func (self *MetathingsDevicedService) Connect(stream pb.DevicedService_ConnectServer) error {
	var dev_s *storage.Device
	var conn connection.Connection
	var err error

	ctx := stream.Context()
	if dev_s, err = self.get_device_by_context(ctx); err != nil {
		self.logger.WithError(err).Errorf("failed to get device by context in storage")
		return status.Errorf(codes.Internal, err.Error())
	}

	sess := grpc_helper.GetSessionFromContext(ctx)
	defer func() {
		if session_helper.IsMajorSession(sess) {
			self.offline_device(*dev_s.Id)
		}
	}()

	if conn, err = self.cc.BuildConnection(dev_s, stream); err != nil {
		self.logger.WithError(err).Errorf("failed to build connection")
		return status.Errorf(codes.Internal, err.Error())
	}
	defer conn.Cleanup()

	self.logger.WithFields(log.Fields{
		"session":   sess,
		"device_id": *dev_s.Id,
		"kind":      *dev_s.Kind,
		"name":      *dev_s.Name,
	}).Debugf("build connection")

	<-conn.Wait()
	self.logger.WithField("device_id", *dev_s.Id).Debugf("connection closed")

	if err = conn.Err(); err != nil {
		self.logger.WithError(err).Errorf("connect error")
		return status.Errorf(codes.Internal, err.Error())
	}

	return nil
}
