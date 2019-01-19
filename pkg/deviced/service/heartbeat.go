package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	session_helper "github.com/nayotta/metathings/pkg/common/session"
	deviced_helper "github.com/nayotta/metathings/pkg/deviced/helper"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	state_pb "github.com/nayotta/metathings/pkg/proto/constant/state"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateHeartbeat(ctx context.Context, in interface{}) error {
	return self.validate_chain(
		[]interface{}{
			func() (policy_helper.Validator, get_devicer) {
				req := in.(*pb.HeartbeatRequest)
				return req, req
			},
		},
		[]interface{}{ensure_get_device_id},
	)
}

func (self *MetathingsDevicedService) AuthorizeHeartbeat(ctx context.Context, in interface{}) error {
	return self.enforce(ctx, in.(*pb.HeartbeatRequest).GetDevice().GetId().GetValue(), "heartbeat")
}

func (self *MetathingsDevicedService) Heartbeat(ctx context.Context, req *pb.HeartbeatRequest) (*empty.Empty, error) {
	var dev_s *storage.Device
	var patch_dev_s *storage.Device
	var patch_mdl_s *storage.Module
	var err error

	dev := req.GetDevice()
	dev_id_str := dev.GetId().GetValue()
	sess := req.GetStartupSession().GetValue()

	cur_sess, err := self.session_storage.GetStartupSession(dev_id_str)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get startup session")
		return nil, err
	}

	if cur_sess != sess {
		err = ErrDuplicatedDeviceInstance
		self.logger.WithError(err).Errorf("current startup session not equal heartbeat startup session")
		return nil, err
	}

	if err = self.session_storage.RefreshStartupSession(dev_id_str, session_helper.STARTUP_SESSION_EXPIRE); err != nil {
		self.logger.WithError(err).Errorf("failed to refresh startup session")
		return nil, err
	}

	if dev_s, err = self.storage.GetDevice(dev_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	for _, mdl_from_dev := range dev.GetModules() {
		for _, mdl_from_stor := range dev_s.Modules {
			if mdl_from_dev.GetId().GetValue() == *mdl_from_stor.Id {
				mdl_id_str := *mdl_from_stor.Id

				heartbeat_at := pb_helper.ToTime(*mdl_from_dev.GetHeartbeatAt())
				patch_mdl_s = &storage.Module{
					HeartbeatAt: &heartbeat_at,
				}

				mdl_state_from_stor_str := *mdl_from_stor.State
				mdl_state_from_dev_str := deviced_helper.MODULE_STATE_ENUMER.ToString(mdl_from_dev.GetState())
				if mdl_state_from_dev_str != mdl_state_from_stor_str {
					patch_mdl_s.State = &mdl_state_from_dev_str
				}
				if _, err = self.storage.PatchModule(mdl_id_str, patch_mdl_s); err != nil {
					self.logger.WithError(err).Errorf("failed to patch module in storage")
					return nil, status.Errorf(codes.Internal, err.Error())
				}

				self.logger.WithFields(log.Fields{
					"device_id":    dev_id_str,
					"module_id":    mdl_id_str,
					"heartbeat_at": heartbeat_at,
					"state":        mdl_state_from_dev_str,
				}).Debugf("module heartbeat")
			}
		}
	}

	heartbeat_at := pb_helper.ToTime(*dev.GetHeartbeatAt())
	patch_dev_s = &storage.Device{
		HeartbeatAt: &heartbeat_at,
	}
	if deviced_helper.DEVICE_STATE_ENUMER.ToValue(*dev_s.State) == state_pb.DeviceState_DEVICE_STATE_OFFLINE {
		state_str := deviced_helper.DEVICE_STATE_ENUMER.ToString(state_pb.DeviceState_DEVICE_STATE_ONLINE)
		patch_dev_s.State = &state_str
	}
	if dev_s, err = self.storage.PatchDevice(dev_id_str, patch_dev_s); err != nil {
		self.logger.WithError(err).Errorf("failed to patch device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"device_id":    dev_id_str,
		"heartbeat_at": heartbeat_at,
		"state":        *dev_s.State,
	}).Debugf("device heartbeat")

	return &empty.Empty{}, nil
}
