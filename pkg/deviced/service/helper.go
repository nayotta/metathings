package metathings_deviced_service

import (
	"errors"

	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	deviced_helper "github.com/nayotta/metathings/pkg/deviced/helper"
	simple_storage "github.com/nayotta/metathings/pkg/deviced/simple_storage"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	state_pb "github.com/nayotta/metathings/pkg/proto/constant/state"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func parse_object(x *pb.OpObject) *simple_storage.Object {
	device := x.GetDevice().GetId().GetValue()
	prefix := x.GetPrefix().GetValue()
	name := x.GetName().GetValue()
	return simple_storage.NewObject(device, prefix, name)
}

func copy_device(x *storage.Device) *pb.Device {
	y := &pb.Device{
		Id:      *x.Id,
		Kind:    deviced_helper.DEVICE_KIND_ENUMER.ToValue(*x.Kind),
		State:   deviced_helper.DEVICE_STATE_ENUMER.ToValue(*x.State),
		Name:    *x.Name,
		Alias:   *x.Alias,
		Modules: copy_modules(x.Modules),
		Flows:   copy_flows(x.Flows),
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
		Id:        *x.Id,
		State:     deviced_helper.MODULE_STATE_ENUMER.ToValue(*x.State),
		DeviceId:  *x.DeviceId,
		Endpoint:  *x.Endpoint,
		Component: *x.Component,
		Name:      *x.Name,
		Alias:     *x.Alias,
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

func copy_flow(x *storage.Flow) *pb.Flow {
	y := &pb.Flow{
		Id:       *x.Id,
		DeviceId: *x.DeviceId,
		Name:     *x.Name,
		Alias:    *x.Alias,
	}

	return y
}

func copy_flows(xs []*storage.Flow) []*pb.Flow {
	var ys []*pb.Flow

	for _, x := range xs {
		ys = append(ys, copy_flow(x))
	}

	return ys
}

func copy_object(x *simple_storage.Object) *pb.Object {
	mod := pb_helper.FromTime(x.LastModified)

	y := &pb.Object{
		Device:       &pb.Device{Id: x.Device},
		Prefix:       x.Prefix,
		Name:         x.Name,
		Length:       x.Length,
		Etag:         x.Etag,
		LastModified: &mod,
	}

	return y
}

func copy_objects(xs []*simple_storage.Object) []*pb.Object {
	var ys []*pb.Object

	for _, x := range xs {
		ys = append(ys, copy_object(x))
	}

	return ys
}

type device_getter interface {
	GetDevice() *pb.OpDevice
}

type module_getter interface {
	GetModule() *pb.OpModule
}

type object_getter interface {
	GetObject() *pb.OpObject
}

type source_getter interface {
	GetSource() *pb.OpObject
}

type destination_getter interface {
	GetDestination() *pb.OpObject
}

func ensure_get_device_id(x device_getter) error {
	if x.GetDevice().GetId() == nil {
		return errors.New("device.id is empty")
	}
	return nil
}

func ensure_get_module_id(x module_getter) error {
	if x.GetModule().GetId() == nil {
		return errors.New("module.id is empty")
	}
	return nil
}

func ensure_get_object_name(x object_getter) error {
	return _ensure_get_object_name(x.GetObject())
}

func _ensure_get_object_name(x *pb.OpObject) error {
	if x.GetName() == nil {
		return errors.New("object.name is empty")
	}
	return nil
}

func ensure_get_object_device_id(x object_getter) error {
	return _ensure_get_object_device_id(x.GetObject())
}

func _ensure_get_object_device_id(x *pb.OpObject) error {
	dev := x.GetDevice()
	if dev == nil {
		return errors.New("object.device.id is empty")
	}
	if dev.GetId() == nil {
		return errors.New("object.device.id is empty")
	}
	return nil
}

func ensure_device_online(s storage.Storage) func(device_getter) error {
	return func(x device_getter) error {
		dev_id := x.GetDevice().GetId().GetValue()
		dev, err := s.GetDevice(dev_id)
		if err != nil {
			return err
		}
		if copy_device(dev).State != state_pb.DeviceState_DEVICE_STATE_ONLINE {
			return ErrDeviceOffline
		}

		return nil
	}
}
