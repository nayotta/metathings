package metathings_deviced_service

import (
	"context"
	"errors"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateMqttHeartbeatSelect(ctx context.Context, in interface{}) error {
	return self.validate_chain(
		[]interface{}{
			func() (policy_helper.Validator, get_devicer) {
				req := in.(*pb.MqttHeartbeatSelectRequest)
				return req, req
			},
		},
		[]interface{}{
			func(x get_devicer) error {
				dev := x.GetDevice()

				if dev.GetId() == nil {
					return errors.New("device.id is empty")
				}

				if dev.GetId().GetValue() == "" {
					return errors.New("device.id.value is empty")
				}

				return nil
			},
		},
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
