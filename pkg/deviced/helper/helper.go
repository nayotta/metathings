package metathings_deviced_helper

import (
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	pb_kind "github.com/nayotta/metathings/pkg/proto/constant/kind"
	pb_state "github.com/nayotta/metathings/pkg/proto/constant/state"
)

type DeviceStateEnumer struct {
	enumer pb_helper.Enumer
}

func (self DeviceStateEnumer) ToString(x pb_state.DeviceState) string {
	return self.enumer.ToString(int32(x))
}

func (self DeviceStateEnumer) ToValue(x string) pb_state.DeviceState {
	return pb_state.DeviceState(self.enumer.ToValue(x))
}

var (
	DEVICE_STATE_ENUMER = DeviceStateEnumer{
		enumer: pb_helper.NewEnumer("device_state", pb_state.DeviceState_name, pb_state.DeviceState_value),
	}
)

type DeviceKindEnumer struct {
	enumer pb_helper.Enumer
}

func (self DeviceKindEnumer) ToString(x pb_kind.DeviceKind) string {
	return self.enumer.ToString(int32(x))
}

func (self DeviceKindEnumer) ToValue(x string) pb_kind.DeviceKind {
	return pb_kind.DeviceKind(self.enumer.ToValue(x))
}

var (
	DEVICE_KIND_ENUMER = DeviceKindEnumer{
		enumer: pb_helper.NewEnumer("device_kind", pb_kind.DeviceKind_name, pb_kind.DeviceKind_value),
	}
)

type ModuleStateEnumer struct {
	enumer pb_helper.Enumer
}

func (self ModuleStateEnumer) ToString(x pb_state.ModuleState) string {
	return self.enumer.ToString(int32(x))
}

func (self ModuleStateEnumer) ToValue(x string) pb_state.ModuleState {
	return pb_state.ModuleState(self.enumer.ToValue(x))
}

var (
	MODULE_STATE_ENUMER = ModuleStateEnumer{
		enumer: pb_helper.NewEnumer("module_state", pb_state.ModuleState_name, pb_state.ModuleState_value),
	}
)

const (
	DEVICE_CONFIG_DESCRIPTOR string = "_sys_descriptor"
)
