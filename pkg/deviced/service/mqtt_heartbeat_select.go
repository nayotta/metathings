package metathings_deviced_service

import (
	"context"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateMqttHeartbeatSelect(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, device_getter) {
				req := in.(*pb.MqttHeartbeatSelectRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_device_id},
	)
}

func (self *MetathingsDevicedService) MqttHeartbeatSelect(ctx context.Context, req *pb.MqttHeartbeatSelectRequest) (*pb.MqttHeartbeatSelectResponse, error) {
	var err error

	devID := req.GetDevice().GetId().GetValue()
	sessionID, err := self.mqttBr.HeartBeatSelect(devID)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to MqttHeartBeatSelect")
		return nil, err
	}

	res := &pb.MqttHeartbeatSelectResponse{
		SessionId: (int32)(sessionID),
	}

	self.logger.WithField("id", devID).Infof("MqttHeartbeatSelect")

	return res, nil
}
