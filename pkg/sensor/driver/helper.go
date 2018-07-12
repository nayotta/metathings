package metathings_sensor_driver

import (
	pb "github.com/nayotta/metathings/pkg/proto/sensor"
	state_helper "github.com/nayotta/metathings/pkg/sensor/state"
)

var _sensor_st_psr = state_helper.SENSOR_STATE_PARSER

func (s SensorState) ToString() string {
	return _sensor_st_psr.ToString(pb.SensorState(s))
}

func StateFromValue(x int32) SensorState {
	if x >= int32(STATE_OVERFLOW) || x < int32(STATE_UNKNOWN) {
		return STATE_UNKNOWN
	}

	return SensorState(x)
}
