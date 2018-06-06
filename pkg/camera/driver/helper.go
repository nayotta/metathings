package metathings_camera_driver

import (
	state_helper "github.com/nayotta/metathings/pkg/camera/state"
	pb "github.com/nayotta/metathings/pkg/proto/camera"
)

var _camera_st_psr = state_helper.NewCameraStateParser()

func (s CameraState) ToString() string {
	return _camera_st_psr.ToString(pb.CameraState(s))
}

func FromValue(x int32) CameraState {
	if x >= int32(STATE_OVERFLOW) || x < int32(STATE_UNKNOWN) {
		return STATE_UNKNOWN
	}

	return CameraState(x)
}
