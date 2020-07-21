package metathings_deviced_service

import (
	"context"
	"errors"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	stpb "github.com/golang/protobuf/ptypes/struct"
	tspb "github.com/golang/protobuf/ptypes/timestamp"

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
	length := x.GetLength().GetValue()

	y := &simple_storage.Object{
		Device: device,
		Prefix: prefix,
		Name:   name,
		Length: length,
	}

	return y
}

func copy_device(x *storage.Device) *pb.Device {
	var hbt *tspb.Timestamp
	if x.HeartbeatAt != nil {
		hbt, _ = ptypes.TimestampProto(*x.HeartbeatAt)
	}

	y := &pb.Device{
		Id:          *x.Id,
		Kind:        deviced_helper.DEVICE_KIND_ENUMER.ToValue(*x.Kind),
		State:       deviced_helper.DEVICE_STATE_ENUMER.ToValue(*x.State),
		HeartbeatAt: hbt,
		Name:        *x.Name,
		Alias:       *x.Alias,
		Extra:       pb_helper.CopyExtra(x.Extra),
		Modules:     copy_modules(x.Modules),
		Flows:       copy_flows(x.Flows),
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

func copy_flow_view(x *storage.Flow) *pb.Flow {
	return &pb.Flow{Id: *x.Id}
}

func copy_flows(xs []*storage.Flow) []*pb.Flow {
	var ys []*pb.Flow

	for _, x := range xs {
		ys = append(ys, copy_flow(x))
	}

	return ys
}

func copy_flows_view(xs []*storage.Flow) []*pb.Flow {
	var ys []*pb.Flow

	for _, x := range xs {
		ys = append(ys, copy_flow_view(x))
	}

	return ys
}

func copy_flow_set(x *storage.FlowSet) *pb.FlowSet {
	y := &pb.FlowSet{
		Id:    *x.Id,
		Name:  *x.Name,
		Alias: *x.Alias,
		Flows: copy_flows_view(x.Flows),
	}

	return y
}

func copy_flow_sets(xs []*storage.FlowSet) []*pb.FlowSet {
	var ys []*pb.FlowSet

	for _, x := range xs {
		ys = append(ys, copy_flow_set(x))
	}

	return ys
}

func copy_config(x *storage.Config) *pb.Config {
	y, _ := copy_config_error(x)
	return y
}

func copy_config_error(x *storage.Config) (*pb.Config, error) {
	var body stpb.Struct

	if err := jsonpb.UnmarshalString(*x.Body, &body); err != nil {
		return nil, err
	}

	y := &pb.Config{
		Id:    *x.Id,
		Alias: *x.Alias,
		Body:  &body,
	}

	return y, nil
}

func copy_configs(xs []*storage.Config) []*pb.Config {
	var ys []*pb.Config

	for _, x := range xs {
		ys = append(ys, copy_config(x))
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

func copy_firmware_descriptor(x *storage.FirmwareDescriptor) *pb.FirmwareDescriptor {
	created_at, _ := ptypes.TimestampProto(x.CreatedAt)

	var desc stpb.Struct
	jsonpb.Unmarshal(strings.NewReader(*x.Descriptor), &desc)

	y := &pb.FirmwareDescriptor{
		Id:          *x.Id,
		Name:        *x.Name,
		CreatedAt:   created_at,
		Descriptor_: &desc,
	}

	return y
}

func copy_firmware_descriptors(xs []*storage.FirmwareDescriptor) []*pb.FirmwareDescriptor {
	var ys []*pb.FirmwareDescriptor

	for _, x := range xs {
		ys = append(ys, copy_firmware_descriptor(x))
	}

	return ys
}

func copy_firmware_hub(x *storage.FirmwareHub) *pb.FirmwareHub {
	var devices []*pb.Device

	for _, dev_s := range x.Devices {
		devices = append(devices, &pb.Device{
			Id: *dev_s.Id,
		})
	}

	y := &pb.FirmwareHub{
		Id:                  *x.Id,
		Alias:               *x.Alias,
		Description:         *x.Description,
		Devices:             devices,
		FirmwareDescriptors: copy_firmware_descriptors(x.FirmwareDescriptors),
	}

	return y
}

func copy_firmware_hubs(xs []*storage.FirmwareHub) []*pb.FirmwareHub {
	var ys []*pb.FirmwareHub

	for _, x := range xs {
		ys = append(ys, copy_firmware_hub(x))
	}

	return ys
}

type descriptor_getter interface {
	GetDescriptor_() *pb.OpDescriptor
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

type flow_set_getter interface {
	GetFlowSet() *pb.OpFlowSet
}

type config_getter interface {
	GetConfig() *pb.OpConfig
}

type firmware_hub_getter interface {
	GetFirmwareHub() *pb.OpFirmwareHub
}

type firmware_descriptor_getter interface {
	GetFirmwareDescriptor() *pb.OpFirmwareDescriptor
}

func ensure_get_descriptor_body(x descriptor_getter) error {
	if x.GetDescriptor_().GetBody() == nil {
		return errors.New("descriptor.body is empty")
	}

	return nil
}

func ensure_get_descriptor_sha1(x descriptor_getter) error {
	if x.GetDescriptor_().GetSha1() == nil {
		return errors.New("descriptor.sha1 is empty")
	}

	return nil
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

func ensure_device_online(ctx context.Context, s storage.Storage) func(device_getter) error {
	return func(x device_getter) error {
		dev_id := x.GetDevice().GetId().GetValue()
		dev, err := s.GetDevice(ctx, dev_id)
		if err != nil {
			return err
		}
		if copy_device(dev).State != state_pb.DeviceState_DEVICE_STATE_ONLINE {
			return ErrDeviceOffline
		}

		return nil
	}
}

func ensure_get_flow_set_id(x flow_set_getter) error {
	fs := x.GetFlowSet()
	if fs == nil {
		return errors.New("flow_set.id is empty")
	}
	if fs.GetId() == nil {
		return errors.New("flow_set.id is empty")
	}

	return nil
}

func ensure_get_config_id(x config_getter) error {
	cfg := x.GetConfig()
	if cfg.GetId() == nil {
		return errors.New("config.id is empty")
	}

	return nil
}

func ensure_get_firmware_hub_id(x firmware_hub_getter) error {
	fh := x.GetFirmwareHub()
	if fh.GetId() == nil {
		return errors.New("firmware_hub.id is empty")
	}

	return nil
}

func ensure_get_firmware_descriptor_id(x firmware_descriptor_getter) error {
	fd := x.GetFirmwareDescriptor()
	if fd.GetId() == nil {
		return errors.New("frimware_descriptor.id is empty")
	}

	return nil
}

func ensure_firmware_hub_contains_device_and_firmware_descriptor(ctx context.Context, s storage.Storage) func(device_getter, firmware_descriptor_getter) error {
	return func(x device_getter, y firmware_descriptor_getter) error {
		dev_id := x.GetDevice().GetId().GetValue()
		desc_id := y.GetFirmwareDescriptor().GetId().GetValue()
		contained, err := s.FirmwareHubContainsDeviceAndFirmwareDescriptor(ctx, dev_id, desc_id)
		if err != nil {
			return err
		}

		if !contained {
			return errors.New("device and firmware descriptor not in the same firmware hub")
		}

		return nil
	}

}
