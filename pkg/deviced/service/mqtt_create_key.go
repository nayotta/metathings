package metathings_deviced_service

import (
	"context"
	"errors"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	deviced_helper "github.com/nayotta/metathings/pkg/deviced/helper"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	pb_state "github.com/nayotta/metathings/pkg/proto/constant/state"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (self *MetathingsDevicedService) ValidateCreateMqttKey(ctx context.Context, in interface{}) error {
	return self.validate_chain(
		[]interface{}{
			func() (policy_helper.Validator, get_devicer) {
				req := in.(*pb.CreateMqttKeyRequest)
				return req, req
			},
		},
		[]interface{}{
			func(x get_devicer) error {
				dev := x.GetDevice()

				if dev.GetId() == nil {
					return errors.New("device.id is empty")
				}

				return nil
			},
		},
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


