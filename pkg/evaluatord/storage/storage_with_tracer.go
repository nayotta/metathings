package metathings_evaluatord_storage

import (
	"context"

	opentracing_storage_helper "github.com/nayotta/metathings/pkg/common/opentracing/storage"
)

type TracedStorage struct {
	*opentracing_storage_helper.BaseTracedStorage
	Storage
}

func (s *TracedStorage) CreateEvaluator(ctx context.Context, e *Evaluator) (*Evaluator, error) {
	span, ctx := s.TraceWrapper(ctx, "CreateEvaluator")
	defer span.Finish()

	return s.Storage.CreateEvaluator(ctx, e)
}

func (s *TracedStorage) DeleteEvaluator(ctx context.Context, id string) error {
	span, ctx := s.TraceWrapper(ctx, "DeleteEvaluator")
	defer span.Finish()

	return s.Storage.DeleteEvaluator(ctx, id)
}

func (s *TracedStorage) PatchEvaluator(ctx context.Context, id string, e *Evaluator) (*Evaluator, error) {
	span, ctx := s.TraceWrapper(ctx, "PatchEvaluator")
	defer span.Finish()

	return s.Storage.PatchEvaluator(ctx, id, e)
}

func (s *TracedStorage) GetEvaluator(ctx context.Context, id string) (*Evaluator, error) {
	span, ctx := s.TraceWrapper(ctx, "GetEvaluator")
	defer span.Finish()

	return s.Storage.GetEvaluator(ctx, id)
}

func (s *TracedStorage) ListEvaluators(ctx context.Context, e *Evaluator) ([]*Evaluator, error) {
	span, ctx := s.TraceWrapper(ctx, "ListEvaluators")
	defer span.Finish()

	return s.Storage.ListEvaluators(ctx, e)
}

func (s *TracedStorage) AddSourcesToEvaluator(ctx context.Context, evaluator_id string, sources []*Resource) error {
	span, ctx := s.TraceWrapper(ctx, "AddSourcesToEvaluator")
	defer span.Finish()

	return s.Storage.AddSourcesToEvaluator(ctx, evaluator_id, sources)
}

func (s *TracedStorage) RemoveSourcesFromEvaluator(ctx context.Context, evaluator_id string, sources []*Resource) error {
	span, ctx := s.TraceWrapper(ctx, "RemoveSourcesFromEvaluator")
	defer span.Finish()

	return s.Storage.RemoveSourcesFromEvaluator(ctx, evaluator_id, sources)
}

func (s *TracedStorage) ExistEvaluator(ctx context.Context, e *Evaluator) (bool, error) {
	span, ctx := s.TraceWrapper(ctx, "ExistEvaluator")
	defer span.Finish()

	return s.Storage.ExistEvaluator(ctx, e)
}

func (s *TracedStorage) ExistOperator(ctx context.Context, o *Operator) (bool, error) {
	span, ctx := s.TraceWrapper(ctx, "ExistOperator")
	defer span.Finish()

	return s.Storage.ExistOperator(ctx, o)
}

func NewTracedStorage(s Storage, getter opentracing_storage_helper.RootDBConnGetter) (Storage, error) {
	return &TracedStorage{
		BaseTracedStorage: opentracing_storage_helper.NewBaseTracedStorage(getter),
		Storage:           s,
	}, nil
}
