package metathings_deviced_service

import (
	"context"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateCreateMqttKey(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, device_getter) {
				req := in.(*pb.CreateMqttKeyRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_device_id},
	)
}

func (self *MetathingsDevicedService) CreateMqttKey(ctx context.Context, req *pb.CreateMqttKeyRequest) (*pb.CreateMqttKeyResponse, error) {
	var err error

	dev := req.GetDevice()

	dev_id_str := dev.GetId().GetValue()

	keyStr, err := self.mqttBr.KeyGenForDeviced(dev_id_str)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to create mqtt key")
		return nil, err
	}

	res := &pb.CreateMqttKeyResponse{
		Key: keyStr,
	}

	self.logger.WithField("id", dev_id_str).Infof("create mqtt key")

	return res, nil
}
