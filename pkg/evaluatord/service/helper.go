package metathings_evaluatord_service

import (
	"context"
	"errors"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	stpb "github.com/golang/protobuf/ptypes/struct"

	storage "github.com/nayotta/metathings/pkg/evaluatord/storage"
	evaluator_plugin "github.com/nayotta/metathings/pkg/plugin/evaluator"
	pb "github.com/nayotta/metathings/pkg/proto/evaluatord"
)

type evaluator_getter interface {
	GetEvaluator() *pb.OpEvaluator
}

type source_getter interface {
	GetSource() *pb.OpResource
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
