package metathings_servo_driver

import (
	pb "github.com/nayotta/metathings/pkg/proto/servo"
	state_helper "github.com/nayotta/metathings/pkg/servo/state"
)

var _servo_st_pst = state_helper.SERVO_STATE_PARSER

func (s ServoState) ToString() string {
	return _servo_st_pst.ToString(pb.ServoState(s))
}

func StateFromValue(x int32) ServoState {
	if x >= int32(STATE_OVERFLOW) || x < int32(STATE_UNKNOWN) {
		return STATE_UNKNOWN
	}

	return ServoState(x)
}

func IsValidAngle(ang float32) bool {
	return ANGLE_MIN <= ang && ang <= ANGLE_MAX
}
