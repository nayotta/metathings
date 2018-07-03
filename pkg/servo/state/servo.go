package servo_state_helper

import (
	state_helper "github.com/nayotta/metathings/pkg/common/state"
	pb "github.com/nayotta/metathings/pkg/proto/servo"
)

type ServoStateParser struct {
	parser state_helper.StateParser
}

func (p ServoStateParser) ToString(s pb.ServoState) string {
	return p.parser.ToString(int32(s))
}

func (p ServoStateParser) ToValue(x string) pb.ServoState {
	return pb.ServoState(p.parser.ToValue(x))
}

func NewServoStateParser() ServoStateParser {
	return ServoStateParser{
		parser: state_helper.NewStateParser("servo_state", pb.ServoState_name, pb.ServoState_value),
	}
}

var (
	SERVO_STATE_PARSER = NewServoStateParser()
)
