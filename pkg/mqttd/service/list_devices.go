package metathingsmqttdservice

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mqttd_helper "github.com/nayotta/metathings/pkg/mqttd/helper"
	storage "github.com/nayotta/metathings/pkg/mqttd/storage"
	pb_state "github.com/nayotta/metathings/pkg/proto/constant/state"
	pb "github.com/nayotta/metathings/pkg/proto/mqttd"
)

// ListDevices ListDevices
func (serv *MetathingsMqttdService) ListDevices(ctx context.Context, req *pb.ListDevicesRequest) (*pb.ListDevicesResponse, error) {
	var devsS []*storage.Device
	var err error

	dev := req.GetDevice()
	devS := &storage.Device{}

	id := dev.GetId()
	if id != nil {
		devS.ID = &id.Value
	}

	state := dev.GetState()
	if state != pb_state.DeviceState_DEVICE_STATE_UNKNOWN {
		stateStr := mqttd_helper.DEVICESTATEENUMER.ToString(state)
		devS.State = &stateStr
	}

	name := dev.GetName()
	if name != nil {
		devS.Name = &name.Value
	}

	alias := dev.GetAlias()
	if alias != nil {
		devS.Alias = &alias.Value
	}

	if devsS, err = serv.storage.ListDevices(devS); err != nil {
		serv.logger.WithError(err).Errorf("failed to list devices in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListDevicesResponse{
		Devices: copyDevices(devsS),
	}

	serv.logger.Debugf("list devices")

	return res, nil
}
