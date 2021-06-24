package metathings_deviced_service

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	session_helper "github.com/nayotta/metathings/pkg/common/session"
	connection "github.com/nayotta/metathings/pkg/deviced/connection"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDevicedService) Connect(stream pb.DevicedService_ConnectServer) error {
	var dev_s *storage.Device
	var conn connection.Connection
	var err error

	logger := self.get_logger().WithFields(log.Fields{
		"method": "Connect",
	})

	ctx := stream.Context()
	if dev_s, err = self.get_device_by_context(ctx); err != nil {
		logger.WithError(err).Errorf("failed to get device by context in storage")
		return status.Errorf(codes.Internal, err.Error())
	}

	sess := grpc_helper.GetSessionFromContext(ctx)
	defer func() {
		cnt, err := self.cc.CountConnections(stream.Context(), dev_s)
		if err != nil {
			logger.WithError(err).Warningf("failed to count connections")
		}
		if session_helper.IsMajorSession(sess) && cnt == 0 {
			self.offline_device(ctx, *dev_s.Id)
		}
	}()

	logger = logger.WithFields(log.Fields{
		"session": sess,
		"device":  *dev_s.Id,
		"kind":    *dev_s.Kind,
	})

	// TODO: limit max connections
	if conn, err = self.cc.BuildConnection(stream.Context(), dev_s, stream); err != nil {
		logger.WithError(err).Errorf("failed to build connection")
		return status.Errorf(codes.Internal, err.Error())
	}
	defer conn.Cleanup()

	logger.Debugf("build connection")

	<-conn.Wait()
	logger.Debugf("connection closed")

	if err = conn.Err(); err != nil {
		logger.WithError(err).Errorf("connect error")
		return status.Errorf(codes.Internal, err.Error())
	}

	return nil
}
