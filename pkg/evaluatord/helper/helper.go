package metathings_evaluatord_helper

import (
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	pb_state "github.com/nayotta/metathings/pkg/proto/constant/state"
)

type TaskStateEnumer struct {
	enumer pb_helper.Enumer
}

func (e TaskStateEnumer) ToStringP(x pb_state.TaskState) *string {
	s := e.enumer.ToString(int32(x))
	return &s
}

func (e TaskStateEnumer) ToValue(x string) pb_state.TaskState {
	return pb_state.TaskState(e.enumer.ToValue(x))
}

var (
	TASK_STATE_ENUMER = TaskStateEnumer{
		enumer: pb_helper.NewEnumer("task_state", pb_state.TaskState_name, pb_state.TaskState_value),
	}
)
