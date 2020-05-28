package metathings_evaluatord_storage

import (
	"context"

	opentracing_storage_helper "github.com/nayotta/metathings/pkg/common/opentracing/storage"
)

type TracedTimerStorage struct {
	*opentracing_storage_helper.BaseTracedStorage
	TimerStorage
}

func (s *TracedTimerStorage) CreateTimer(ctx context.Context, t *Timer) (*Timer, error) {
	span, ctx := s.TraceWrapper(ctx, "CreateTimer")
	defer span.Finish()

	return s.TimerStorage.CreateTimer(ctx, t)
}

func (s *TracedTimerStorage) DeleteTimer(ctx context.Context, id string) error {
	span, ctx := s.TraceWrapper(ctx, "DeleteTimer")
	defer span.Finish()

	return s.TimerStorage.DeleteTimer(ctx, id)
}

func (s *TracedTimerStorage) PatchTimer(ctx context.Context, id string, t *Timer) (*Timer, error) {
	span, ctx := s.TraceWrapper(ctx, "PatchTimer")
	defer span.Finish()

	return s.TimerStorage.PatchTimer(ctx, id, t)
}

func (s *TracedTimerStorage) GetTimer(ctx context.Context, id string) (*Timer, error) {
	span, ctx := s.TraceWrapper(ctx, "GetTimer")
	defer span.Finish()

	return s.TimerStorage.GetTimer(ctx, id)
}

func (s *TracedTimerStorage) ListTimers(ctx context.Context, t *Timer) ([]*Timer, error) {
	span, ctx := s.TraceWrapper(ctx, "ListTimers")
	defer span.Finish()

	return s.TimerStorage.ListTimers(ctx, t)
}

func (s *TracedTimerStorage) AddConfigsToTimer(ctx context.Context, timer_id string, config_ids []string) error {
	span, ctx := s.TraceWrapper(ctx, "AddConfigsToTimer")
	defer span.Finish()

	return s.TimerStorage.AddConfigsToTimer(ctx, timer_id, config_ids)
}

func (s *TracedTimerStorage) RemoveConfigsFromTimer(ctx context.Context, timer_id string, config_ids []string) error {
	span, ctx := s.TraceWrapper(ctx, "RemoveConfigsFromTimer")
	defer span.Finish()

	return s.TimerStorage.RemoveConfigsFromTimer(ctx, timer_id, config_ids)
}

func NewTracedTimerStorage(s TimerStorage, getter opentracing_storage_helper.RootDBConnGetter) (TimerStorage, error) {
	return &TracedTimerStorage{
		BaseTracedStorage: opentracing_storage_helper.NewBaseTracedStorage(getter),
		TimerStorage:      s,
	}, nil
}
