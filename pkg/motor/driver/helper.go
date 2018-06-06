package metathings_motor_driver

import (
	state_helper "github.com/nayotta/metathings/pkg/motor/state"
	pb "github.com/nayotta/metathings/pkg/proto/motor"
)

var _motor_st_psr = state_helper.MOTOR_STATE_PARSER

func (s MotorState) ToString() string {
	return _motor_st_psr.ToString(pb.MotorState(s))
}

func StateFromValue(x int32) MotorState {
	if x >= int32(STATE_OVERFLOW) || x < int32(STATE_UNKNOWN) {
		return STATE_UNKNOWN
	}

	return MotorState(x)
}

var _motor_dir_psr = state_helper.MOTOR_DIRECTION_PARSER

func (d MotorDirection) ToString() string {
	return _motor_dir_psr.ToString(pb.MotorDirection(d))
}

func DirectionFromValue(x int32) MotorDirection {
	if x >= int32(DIRECTION_OVERFLOW) || x < int32(STATE_UNKNOWN) {
		return DIRECTION_UNKNOWN
	}

	return MotorDirection(x)
}

func IsValidSpeed(spd float32) bool {
	return SPEED_MIN <= spd && spd <= SPEED_MAX
}
