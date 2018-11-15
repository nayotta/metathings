package metathings_deviced_service

import (
	"errors"

	deviced_helper "github.com/nayotta/metathings/pkg/deviced/helper"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func copy_device(x *storage.Device) *pb.Device {
	y := &pb.Device{
		Id:      *x.Id,
		Kind:    deviced_helper.DEVICE_KIND_ENUMER.ToValue(*x.Kind),
		State:   deviced_helper.DEVICE_STATE_ENUMER.ToValue(*x.State),
		Name:    *x.Name,
		Alias:   *x.Alias,
		Modules: copy_modules(x.Modules),
	}

	return y
}

func copy_devices(xs []*storage.Device) []*pb.Device {
	var ys []*pb.Device

	for _, x := range xs {
		ys = append(ys, copy_device(x))
	}

	return ys
}

func copy_module(x *storage.Module) *pb.Module {
	y := &pb.Module{
		Id:       *x.Id,
		State:    deviced_helper.MODULE_STATE_ENUMER.ToValue(*x.State),
		DeviceId: *x.DeviceId,
		Endpoint: *x.Endpoint,
		Name:     *x.Name,
		Alias:    *x.Alias,
	}

	return y
}

func copy_modules(xs []*storage.Module) []*pb.Module {
	var ys []*pb.Module

	for _, x := range xs {
		ys = append(ys, copy_module(x))
	}

	return ys
}

type get_devicer interface {
	GetDevice() *pb.OpDevice
}

func ensure_get_device_id(x get_devicer) error {
	if x.GetDevice().GetId() == nil {
		return errors.New("device.id is empty")
	}
	return nil
}
