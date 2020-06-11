package metathings_evaluatord_storage

import "context"

type DummyTaskStorage struct{}

func (*DummyTaskStorage) ListTasksBySource(ctx context.Context, src *Resource, opts ...ListTasksBySourceOption) ([]*Task, error) {
	panic("unimplemented")
}

func (*DummyTaskStorage) GetTask(context.Context, string) (*Task, error) {
	panic("unimplemented")
}

func (*DummyTaskStorage) PatchTask(context.Context, *Task, *TaskState) error {
	panic("unimplemented")
}

func NewDummyTaskStorage(...interface{}) (TaskStorage, error) {
	return &DummyTaskStorage{}, nil
}

func init() {
	register_task_storage_factory("dummy", NewDummyTaskStorage)
}
