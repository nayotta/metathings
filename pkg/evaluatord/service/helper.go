package metathings_evaluatord_service

import (
	"context"
	"errors"

	storage "github.com/nayotta/metathings/pkg/evaluatord/storage"
	evaluator_plugin "github.com/nayotta/metathings/pkg/plugin/evaluator"
	pb "github.com/nayotta/metathings/pkg/proto/evaluatord"
)

type evaluator_getter interface {
	GetEvaluator() *pb.OpEvaluator
}

func copy_evaluator(x *storage.Evaluator) *pb.Evaluator {
	panic("unimplemented")
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
