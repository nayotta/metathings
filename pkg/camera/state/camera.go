package camera_state_helper

import (
	state_helper "github.com/nayotta/metathings/pkg/common/state"
	pb "github.com/nayotta/metathings/pkg/proto/camera"
)

type CameraStateParser struct {
	parser state_helper.StateParser
}

func (p CameraStateParser) ToString(s pb.CameraState) string {
	return p.parser.ToString(int32(s))
}

func (p CameraStateParser) ToValue(x string) pb.CameraState {
	return pb.CameraState(p.parser.ToValue(x))
}

func NewCameraStateParser() CameraStateParser {
	return CameraStateParser{
		parser: state_helper.NewStateParser("camera_state", pb.CameraState_name, pb.CameraState_value),
	}
}
