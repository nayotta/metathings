package metathingsmqttdservice

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/wrappers"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	mqttd_helper "github.com/nayotta/metathings/pkg/mqttd/helper"
	storage "github.com/nayotta/metathings/pkg/mqttd/storage"
	pb_state "github.com/nayotta/metathings/pkg/proto/constant/state"
	identityd_pb "github.com/nayotta/metathings/pkg/proto/identityd2"
	pb "github.com/nayotta/metathings/pkg/proto/mqttd"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ValidateCreateDevice ValidateCreateDevice
func (serv *MetathingsMqttdService) ValidateCreateDevice(ctx context.Context, in interface{}) error {
	return serv.validateChain(
		[]interface{}{
			func() (policy_helper.Validator, getDevicer) {
				req := in.(*pb.CreateDeviceRequest)
				return req, req
			},
		},
		[]interface{}{
			func(x getDevicer) error {
				dev := x.GetDevice()

				if dev.GetName() == nil {
					return errors.New("domain.name is empty")
				}

				return nil
			},
		},
	)
}

func (serv *MetathingsMqttdService) createEntity(entID, entName string) error {
	cli, cfn, err := serv.cliFty.NewIdentityd2ServiceClient()
	if err != nil {
		return err
	}
	cfn()

	ctx := context_helper.WithToken(context.Background(), serv.tknr.GetToken())
	req := &identityd_pb.CreateEntityRequest{
		Id:   &wrappers.StringValue{Value: entID},
		Name: &wrappers.StringValue{Value: entName},
	}
	if _, err = cli.CreateEntity(ctx, req); err != nil {
		return err
	}

	return nil
}

func (serv *MetathingsMqttdService) createDeviceEntity(dev *storage.Device) error {
	return serv.createEntity(*dev.ID, "/deviced/device/"+*dev.Name)
}

// CreateDevice create device
func (serv *MetathingsMqttdService) CreateDevice(ctx context.Context, req *pb.CreateDeviceRequest) (*pb.CreateDeviceResponse, error) {
	var err error

	dev := req.GetDevice()

	devIDStr := id_helper.NewId()
	if dev.GetId() != nil {
		devIDStr = dev.GetId().GetValue()
	}
	devStateStr := mqttd_helper.DEVICESTATEENUMER.ToString(pb_state.DeviceState_DEVICE_STATE_OFFLINE)
	devNameStr := dev.GetName().GetValue()
	devAliasStr := devNameStr
	if dev.GetAlias() != nil {
		devAliasStr = dev.GetAlias().GetValue()
	}

	devS := &storage.Device{
		ID:    &devIDStr,
		State: &devStateStr,
		Name:  &devNameStr,
		Alias: &devAliasStr,
	}

	if err = serv.createDeviceEntity(devS); err != nil {
		serv.logger.WithError(err).Errorf("failed to create entity for device")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err = serv.enforcer.AddObjectToKind(devIDStr, KINDDEVICE); err != nil {
		serv.logger.WithError(err).Errorf("failed to add device in enforcer")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if devS, err = serv.storage.CreateDevice(devS); err != nil {
		serv.logger.WithError(err).Errorf("failed to create device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CreateDeviceResponse{
		Device: copyDevice(devS),
	}

	serv.logger.WithField("id", devIDStr).Infof("create device")

	return res, nil
}
