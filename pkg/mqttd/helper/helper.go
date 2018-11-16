package metathingsmqttdhelper

import (
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	pb_kind "github.com/nayotta/metathings/pkg/proto/constant/kind"
	pb_state "github.com/nayotta/metathings/pkg/proto/constant/state"
)

// DeviceStateEnumer DeviceStateEnumer
type DeviceStateEnumer struct {
	enumer pb_helper.Enumer
}

// ToString ToString
func (that DeviceStateEnumer) ToString(x pb_state.DeviceState) string {
	return that.enumer.ToString(int32(x))
}

// ToValue ToValue
func (that DeviceStateEnumer) ToValue(x string) pb_state.DeviceState {
	return pb_state.DeviceState(that.enumer.ToValue(x))
}

// DEVICESTATEENUMER DEVICE_STATE_ENUMER
var (
	DEVICESTATEENUMER = DeviceStateEnumer{
		enumer: pb_helper.NewEnumer("device_state", pb_state.DeviceState_name, pb_state.DeviceState_value),
	}
)

// DeviceKindEnumer DeviceKindEnumer
type DeviceKindEnumer struct {
	enumer pb_helper.Enumer
}

// ToString ToString
func (that DeviceKindEnumer) ToString(x pb_kind.DeviceKind) string {
	return that.enumer.ToString(int32(x))
}

// ToValue ToValue
func (that DeviceKindEnumer) ToValue(x string) pb_kind.DeviceKind {
	return pb_kind.DeviceKind(that.enumer.ToValue(x))
}

// DEVICEKINDENUMER DEVICE_KIND_ENUMER
var (
	DEVICEKINDENUMER = DeviceKindEnumer{
		enumer: pb_helper.NewEnumer("device_kind", pb_kind.DeviceKind_name, pb_kind.DeviceKind_value),
	}
)

