package metathings_deviced_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	deviced_helper "github.com/nayotta/metathings/pkg/deviced/helper"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	pb_kind "github.com/nayotta/metathings/proto/constant/kind"
	pb_state "github.com/nayotta/metathings/proto/constant/state"
	pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDevicedService) ListDevices(ctx context.Context, req *pb.ListDevicesRequest) (*pb.ListDevicesResponse, error) {
	var devs_s []*storage.Device
	var err error

	dev := req.GetDevice()
	dev_s := &storage.Device{}

	logger := self.get_logger()

	id := dev.GetId()
	if id != nil {
		dev_s.Id = &id.Value
	}

	kind := dev.GetKind()
	if kind != pb_kind.DeviceKind_DEVICE_KIND_UNKNOWN {
		kind_str := deviced_helper.DEVICE_KIND_ENUMER.ToString(kind)
		dev_s.Kind = &kind_str
	}

	state := dev.GetState()
	if state != pb_state.DeviceState_DEVICE_STATE_UNKNOWN {
		state_str := deviced_helper.DEVICE_STATE_ENUMER.ToString(state)
		dev_s.State = &state_str
	}

	name := dev.GetName()
	if name != nil {
		dev_s.Name = &name.Value
	}

	alias := dev.GetAlias()
	if alias != nil {
		dev_s.Alias = &alias.Value
	}

	if devs_s, err = self.storage.ListDevices(ctx, dev_s); err != nil {
		logger.WithError(err).Errorf("failed to list devices in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListDevicesResponse{
		Devices: copy_devices(devs_s),
	}

	logger.Debugf("list devices")

	return res, nil
}
