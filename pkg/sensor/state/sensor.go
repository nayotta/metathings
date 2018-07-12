package sensor_state_helper

import (
	state_helper "github.com/nayotta/metathings/pkg/common/state"
	pb "github.com/nayotta/metathings/pkg/proto/sensor"
)

type SensorStateParser struct {
	parser state_helper.StateParser
}

func (p SensorStateParser) ToString(s pb.SensorState) string {
	return p.parser.ToString(int32(s))
}

func (p SensorStateParser) ToValue(x string) pb.SensorState {
	return pb.SensorState(p.parser.ToValue(x))
}

func NewSensorStateParser() SensorStateParser {
	return SensorStateParser{
		parser: state_helper.NewStateParser("sensor_state", pb.SensorState_name, pb.SensorState_value),
	}
}

var (
	SENSOR_STATE_PARSER = NewSensorStateParser()
)
