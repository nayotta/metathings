package metathings_deviced_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	deviced_helper "github.com/nayotta/metathings/pkg/deviced/helper"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	pb_state "github.com/nayotta/metathings/pkg/proto/constant/state"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) CreateDevice(ctx context.Context, req *pb.CreateDeviceRequest) (*pb.CreateDeviceResponse, error) {
	var err error

	dev := req.GetDevice()

	dev_id_str := id_helper.NewId()
	if dev.GetId() != nil {
		dev_id_str = dev.GetId().GetValue()
	}
	dev_kind_str := deviced_helper.DEVICE_KIND_ENUMER.ToString(dev.GetKind())
	dev_state_str := deviced_helper.DEVICE_STATE_ENUMER.ToString(pb_state.DeviceState_DEVICE_STATE_OFFLINE)
	dev_name_str := dev.GetName().GetValue()
	dev_alias_str := dev.GetAlias().GetValue()

	dev_s := &storage.Device{
		Id:    &dev_id_str,
		Kind:  &dev_kind_str,
		State: &dev_state_str,
		Name:  &dev_name_str,
		Alias: &dev_alias_str,
	}

	for _, mdl := range dev.GetModules() {
		mdl_id_str := id_helper.NewId()
		if mdl.GetId() != nil {
			mdl_id_str = mdl.GetId().GetValue()
		}
		mdl_state_str := deviced_helper.MODULE_STATE_ENUMER.ToString(pb_state.ModuleState_MODULE_STATE_OFFLINE)
		mdl_name_str := mdl.GetName().GetValue()
		mdl_alias_str := mdl.GetAlias().GetValue()

		mdl_s := &storage.Module{
			Id:    &mdl_id_str,
			State: &mdl_state_str,
			Name:  &mdl_name_str,
			Alias: &mdl_alias_str,
		}

		if _, err = self.storage.CreateModule(mdl_s); err != nil {
			self.logger.WithError(err).Errorf("failed to create module in storage")
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	if dev_s, err = self.storage.CreateDevice(dev_s); err != nil {
		self.logger.WithError(err).Errorf("failed to create device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CreateDeviceResponse{
		Device: copy_device(dev_s),
	}

	self.logger.WithField("id", dev_id_str).Infof("create device")

	return res, nil
}
