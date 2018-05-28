package motor_state_helper

import (
	state_helper "github.com/nayotta/metathings/pkg/common/state"
	pb "github.com/nayotta/metathings/pkg/proto/motor"
)

type MotorStateParser struct {
	parser state_helper.StateParser
}

func (p MotorStateParser) ToString(s pb.MotorState) string {
	return p.parser.ToString(int32(s))
}

func (p MotorStateParser) ToValue(x string) pb.MotorState {
	return pb.MotorState(p.parser.ToValue(x))
}

func NewMotorStateParser() MotorStateParser {
	return MotorStateParser{
		parser: state_helper.NewStateParser("motor_state", pb.MotorState_name, pb.MotorState_value),
	}
}

type MotorDirectionParser struct {
	parser state_helper.StateParser
}

func (p MotorDirectionParser) ToString(d pb.MotorDirection) string {
	return p.parser.ToString(int32(d))
}

func (p MotorDirectionParser) ToValue(x string) pb.MotorDirection {
	return pb.MotorDirection(p.parser.ToValue(x))
}

func NewMotorDirectionParser() MotorDirectionParser {
	return MotorDirectionParser{
		parser: state_helper.NewStateParser("motor_direction", pb.MotorDirection_name, pb.MotorDirection_value),
	}
}

var (
	MOTOR_STATE_PARSER     = NewMotorStateParser()
	MOTOR_DIRECTION_PARSER = NewMotorDirectionParser()
)
