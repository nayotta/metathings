package metathings_deviced_service

import (
	"context"
	"errors"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	deviced_helper "github.com/nayotta/metathings/pkg/deviced/helper"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	state_pb "github.com/nayotta/metathings/pkg/proto/constant/state"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MqttHeartbeat MqttHeartbeat
func (self *MetathingsDevicedService) MqttHeartbeat(ctx context.Context, req *pb.MqttHeartbeatRequest) (*empty.Empty, error) {
	var dev_s *storage.Device
	var patch_dev_s *storage.Device
	var err error

	devID := req.GetDeviceId()
	if devID == "" {
		return nil, errors.New("device.id is empty")
	}

	if dev_s, err = self.storage.GetDevice(devID); err != nil {
		self.logger.WithError(err).Errorf("failed to get device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	now := time.Now()
	patch_dev_s = &storage.Device{
		HeartbeatAt: &now,
	}
	if deviced_helper.DEVICE_STATE_ENUMER.ToValue(*dev_s.State) == state_pb.DeviceState_DEVICE_STATE_OFFLINE {
		state_str := deviced_helper.DEVICE_STATE_ENUMER.ToString(state_pb.DeviceState_DEVICE_STATE_ONLINE)
		patch_dev_s.State = &state_str
	}
	if dev_s, err = self.storage.PatchDevice(devID, patch_dev_s); err != nil {
		self.logger.WithError(err).Errorf("failed to patch device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"device_id":    devID,
		"heartbeat_at": now,
		"state":        *dev_s.State,
	}).Debugf("device heartbeat")

	return &empty.Empty{}, nil
}
