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
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	state_pb "github.com/nayotta/metathings/proto/constant/state"
	pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateHeartbeat(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, device_getter) {
				req := in.(*pb.HeartbeatRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_device_id},
	)
}

func (self *MetathingsDevicedService) AuthorizeHeartbeat(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.HeartbeatRequest).GetDevice().GetId().GetValue(), "deviced:heartbeat")
}

func (self *MetathingsDevicedService) Heartbeat(ctx context.Context, req *pb.HeartbeatRequest) (*empty.Empty, error) {
	var dev_s *storage.Device
	var mdls_s []*storage.Module
	var patch_dev_s *storage.Device
	var patch_mdl_s *storage.Module
	var err error

	dev := req.GetDevice()
	dev_id_str := dev.GetId().GetValue()
	sess := req.GetStartupSession().GetValue()

	logger := self.get_logger().WithFields(log.Fields{
		"device":          dev_id_str,
		"startup_session": sess,
	})

	cur_sess, err := self.session_storage.GetStartupSession(dev_id_str)
	if err != nil {
		logger.WithError(err).Errorf("failed to get startup session")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if cur_sess == 0 {
		err = ErrUnconnectedDevice
		logger.WithError(err).Warningf("device not connected")
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	if cur_sess != sess {
		err = ErrDuplicatedDevice
		logger.WithError(err).Errorf("current startup session not equal heartbeat startup session")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err = self.session_storage.RefreshStartupSession(dev_id_str, session_helper.STARTUP_SESSION_EXPIRE); err != nil {
		logger.WithError(err).Errorf("failed to refresh startup session")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if mdls_s, err = self.storage.ListModulesByDeviceId(
		ctx, dev_id_str,
		storage.SelectFieldsOption("module", "id", "state"),
	); err != nil {
		logger.WithError(err).Errorf("failed to list modules by device id in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	for _, mdl_from_dev := range dev.GetModules() {
		for _, mdl_from_stor := range mdls_s {
			if mdl_from_dev.GetId().GetValue() == *mdl_from_stor.Id {
				mdl_id_str := *mdl_from_stor.Id

				heartbeat_at := pb_helper.ToTime(mdl_from_dev.GetHeartbeatAt())
				patch_mdl_s = &storage.Module{
					HeartbeatAt: &heartbeat_at,
				}

				mdl_state_from_stor_str := *mdl_from_stor.State
				mdl_state_from_dev_str := deviced_helper.MODULE_STATE_ENUMER.ToString(mdl_from_dev.GetState())
				if mdl_state_from_dev_str != mdl_state_from_stor_str {
					patch_mdl_s.State = &mdl_state_from_dev_str
				}
				if err = self.storage.ModifyModule(ctx, mdl_id_str, patch_mdl_s); err != nil {
					logger.WithError(err).Errorf("failed to patch module in storage")
					return nil, status.Errorf(codes.Internal, err.Error())
				}

				logger.WithFields(log.Fields{
					"device_id":    dev_id_str,
					"module_id":    mdl_id_str,
					"heartbeat_at": heartbeat_at,
					"state":        mdl_state_from_dev_str,
				}).Debugf("module heartbeat")
			}
		}
	}

	dev_s, err = self.storage.GetDevice(
		ctx, dev_id_str,
		storage.SkipInternalQueryOption(true),
		storage.SelectFieldsOption("device", "state"),
	)
	heartbeat_at := pb_helper.ToTime(dev.GetHeartbeatAt())
	patch_dev_s = &storage.Device{
		HeartbeatAt: &heartbeat_at,
	}
	if deviced_helper.DEVICE_STATE_ENUMER.ToValue(*dev_s.State) == state_pb.DeviceState_DEVICE_STATE_OFFLINE {
		state_str := deviced_helper.DEVICE_STATE_ENUMER.ToString(state_pb.DeviceState_DEVICE_STATE_ONLINE)
		patch_dev_s.State = &state_str
	}
	if err = self.storage.ModifyDevice(ctx, dev_id_str, patch_dev_s); err != nil {
		logger.WithError(err).Errorf("failed to patch device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	logger.WithFields(log.Fields{
		"device_id":    dev_id_str,
		"heartbeat_at": heartbeat_at,
		"state":        *dev_s.State,
	}).Debugf("device heartbeat")

	return &empty.Empty{}, nil
}
