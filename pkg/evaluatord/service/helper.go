package metathings_evaluatord_service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	stpb "github.com/golang/protobuf/ptypes/struct"

	evaluatord_helper "github.com/nayotta/metathings/pkg/evaluatord/helper"
	storage "github.com/nayotta/metathings/pkg/evaluatord/storage"
	evaluator_plugin "github.com/nayotta/metathings/pkg/plugin/evaluator"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
	pb "github.com/nayotta/metathings/pkg/proto/evaluatord"
)

type evaluator_getter interface {
	GetEvaluator() *pb.OpEvaluator
}

type source_getter interface {
	GetSource() *pb.OpResource
}

type task_getter interface {
	GetTask() *pb.OpTask
}

type timer_getter interface {
	GetTimer() *pb.OpTimer
}

func copy_lua_descriptor(x *storage.LuaDescriptor) *pb.Operator_Lua {
	y := &pb.Operator_Lua{
		Lua: &pb.LuaDescriptor{
			Code: *x.Code,
		},
	}

	return y
}

func copy_operator(x *storage.Operator) *pb.Operator {
	y := &pb.Operator{
		Id:          *x.Id,
		Alias:       *x.Alias,
		Description: *x.Description,
		Driver:      *x.Driver,
	}

	// SYM:REFACTOR:lua_operator
	switch *x.Driver {
	case "lua":
		fallthrough
	case "default":
		y.Descriptor_ = copy_lua_descriptor(x.LuaDescriptor)
	}

	return y
}

func copy_resource(x *storage.Resource) *pb.Resource {
	y := &pb.Resource{
		Id:   *x.Id,
		Type: *x.Type,
	}

	return y
}

func copy_resources(xs []*storage.Resource) []*pb.Resource {
	ys := []*pb.Resource{}

	for _, x := range xs {
		ys = append(ys, copy_resource(x))
	}

	return ys
}

func copy_evaluator(x *storage.Evaluator) *pb.Evaluator {
	var cfg stpb.Struct

	// TODO(Peer): catch error
	new(jsonpb.Unmarshaler).Unmarshal(strings.NewReader(*x.Config), &cfg)

	y := &pb.Evaluator{
		Id:          *x.Id,
		Alias:       *x.Alias,
		Description: *x.Description,
		Sources:     copy_resources(x.Sources),
		Operator:    copy_operator(x.Operator),
		Config:      &cfg,
	}

	return y
}

func copy_evaluators(xs []*storage.Evaluator) []*pb.Evaluator {
	var ys []*pb.Evaluator

	for _, x := range xs {
		ys = append(ys, copy_evaluator(x))
	}

	return ys
}

func copy_task_state(x *storage.TaskState) *pb.TaskState {
	var tags stpb.Struct

	at, _ := ptypes.TimestampProto(*x.At)
	tags_buf, _ := json.Marshal(x.Tags)
	jsonpb.Unmarshal(bytes.NewReader(tags_buf), &tags)

	y := &pb.TaskState{
		At:    at,
		State: evaluatord_helper.TASK_STATE_ENUMER.ToValue(*x.State),
		Tags:  &tags,
	}

	return y
}

func copy_task_states(xs []*storage.TaskState) []*pb.TaskState {
	var ys []*pb.TaskState
	for _, x := range xs {
		ys = append(ys, copy_task_state(x))
	}
	return ys
}

func copy_task(x *storage.Task) *pb.Task {
	created_at, _ := ptypes.TimestampProto(*x.States[0].At)
	updated_at, _ := ptypes.TimestampProto(*x.CurrentState.At)

	y := &pb.Task{
		Id:           *x.Id,
		CreatedAt:    created_at,
		UpdatedAt:    updated_at,
		CurrentState: copy_task_state(x.CurrentState),
		Source: &pb.Resource{
			Id:   *x.Source.Id,
			Type: *x.Source.Type,
		},
		States: copy_task_states(x.States),
	}

	return y
}

func copy_tasks(xs []*storage.Task) []*pb.Task {
	var ys []*pb.Task
	for _, x := range xs {
		ys = append(ys, copy_task(x))
	}
	return ys
}

func copy_timer(x *storage.Timer) *pb.Timer {
	var cfgs []*deviced_pb.Config

	for _, cfg_id := range x.Configs {
		cfgs = append(cfgs, &deviced_pb.Config{
			Id: cfg_id,
		})
	}

	y := &pb.Timer{
		Id:          *x.Id,
		Alias:       *x.Alias,
		Description: *x.Description,
		Schedule:    *x.Schedule,
		Timezone:    *x.Timezone,
		Enabled:     *x.Enabled,
		Configs:     cfgs,
	}

	return y
}

func copy_timers(xs []*storage.Timer) []*pb.Timer {
	var ys []*pb.Timer
	for _, x := range xs {
		ys = append(ys, copy_timer(x))
	}
	return ys
}

func ensure_get_source(x source_getter) error {
	if x.GetSource() == nil {
		return errors.New("source is empty")
	}

	return nil
}

func ensure_get_source_id(x source_getter) error {
	if x.GetSource().GetId() == nil {
		return errors.New("source.id is empty")
	}

	return nil
}

func ensure_get_source_type(x source_getter) error {
	if x.GetSource().GetType() == nil {
		return errors.New("source.type is empty")
	}

	return nil
}

func ensure_get_evaluator(x evaluator_getter) error {
	if x.GetEvaluator() == nil {
		return errors.New("evaluator is empty")
	}
	return nil
}

func ensure_get_evaluator_id(x evaluator_getter) error {
	if x.GetEvaluator().GetId() == nil {
		return errors.New("evaluator.id is empty")
	}
	return nil
}

func ensure_get_operator(x evaluator_getter) error {
	e := x.GetEvaluator()
	if e.GetOperator() == nil {
		return errors.New("operator is empty")
	}
	return nil
}

func ensure_evaluator_id_not_exists(ctx context.Context, s storage.Storage) func(x evaluator_getter) error {
	return func(x evaluator_getter) error {
		e := x.GetEvaluator()
		eid := e.GetId()
		if eid == nil {
			return nil
		}
		eid_str := eid.GetValue()
		exists, err := s.ExistEvaluator(ctx, &storage.Evaluator{Id: &eid_str})
		if err != nil {
			return err
		}

		if exists {
			return errors.New("evaluator exists")
		}

		return nil
	}
}

func ensure_evaluator_id_exists(ctx context.Context, s storage.Storage) func(x evaluator_getter) error {
	return func(x evaluator_getter) error {
		e := x.GetEvaluator()
		eid := e.GetId()
		if eid == nil {
			return nil
		}

		eid_str := eid.GetValue()
		exists, err := s.ExistEvaluator(ctx, &storage.Evaluator{Id: &eid_str})
		if err != nil {
			return err
		}

		if !exists {
			return errors.New("evaluator not exists")
		}

		return nil
	}
}

func ensure_operator_id_not_exists(ctx context.Context, s storage.Storage) func(x evaluator_getter) error {
	return func(x evaluator_getter) error {
		e := x.GetEvaluator()
		o := e.GetOperator()
		oid := o.GetId()
		if oid == nil {
			return nil
		}

		oid_str := oid.GetValue()
		exists, err := s.ExistOperator(ctx, &storage.Operator{Id: &oid_str})
		if err != nil {
			return err
		}

		if exists {
			return errors.New("evaluator.operator exists")
		}

		return nil
	}
}

func ensure_timer_id_not_exists(ctx context.Context, ts storage.TimerStorage) func(x timer_getter) error {
	return func(x timer_getter) error {
		t := x.GetTimer()
		tid := t.GetId()
		if tid == nil {
			return nil
		}

		tid_str := tid.GetValue()
		exists, err := ts.ExistTimer(ctx, &storage.Timer{Id: &tid_str})
		if err != nil {
			return err
		}

		if exists {
			return errors.New("timer exists")
		}

		return nil
	}
}

func ensure_timer_id_exists(ctx context.Context, ts storage.TimerStorage) func(x timer_getter) error {
	return func(x timer_getter) error {
		t := x.GetTimer()
		tid := t.GetId()
		if tid == nil {
			return nil
		}

		tid_str := tid.GetValue()
		exists, err := ts.ExistTimer(ctx, &storage.Timer{Id: &tid_str})
		if err != nil {
			return err
		}

		if !exists {
			return errors.New("timer not exists")
		}

		return nil
	}
}

func ensure_valid_operator_driver(x evaluator_getter) error {
	drv := x.GetEvaluator().GetOperator().GetDriver()
	if drv == nil {
		return nil
	}

	drv_str := drv.GetValue()
	if !evaluator_plugin.IsValidOperatorName(drv_str) {
		return errors.New("evaluator.operator.driver is invalid")
	}

	return nil
}

func ensure_get_task(x task_getter) error {
	if tsk := x.GetTask(); tsk == nil {
		return errors.New("task is empty")
	}

	return nil
}

func ensure_get_task_id(x task_getter) error {
	if tsk_id := x.GetTask().GetId(); tsk_id == nil {
		return errors.New("task.id is empty")
	}

	return nil
}

func ensure_get_timer(x timer_getter) error {
	if tmr := x.GetTimer(); tmr == nil {
		return errors.New("timer is empty")
	}

	return nil
}

func ensure_get_timer_id(x timer_getter) error {
	if tmr_id := x.GetTimer().GetId(); tmr_id == nil {
		return errors.New("timer.id is empty")
	}

	return nil
}

// TODO(Peer): unimplemented
func ensure_valid_timer_timezone(x timer_getter) error {
	return nil
}

// TODO(Peer): unimplemented
func ensure_valid_timer_schedule(x timer_getter) error {
	return nil
}
