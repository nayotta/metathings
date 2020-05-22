package metathings_evaluatord_storage

import (
	"context"
	"sync"
	"time"
)

type Resource struct {
	Id   *string
	Type *string
}

type Evaluator struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	Alias       *string `gorm:"column:alias"`
	Description *string `gorm:"column:description"`
	Config      *string `gorm:"column:config"`

	Sources  []*Resource `gorm:"-"`
	Operator *Operator   `gorm:"-"`
}

type Operator struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	EvaluatorId *string `gorm:"column:evaluator_id"`
	Alias       *string `gorm:"column:alias"`
	Description *string `gorm:"column:description"`

	Driver *string `gorm:"column:driver"`

	LuaDescriptor *LuaDescriptor `gorm:"-"`
}

type LuaDescriptor struct {
	CreatedAt time.Time

	OperatorId *string `gorm:"column:operator_id"`

	Code *string `gorm:"column:code"`
}

type EvaluatorSourceMapping struct {
	CreatedAt time.Time

	EvaluatorId *string `gorm:"column:evaluator_id"`
	SourceId    *string `gorm:"column:source_id"`
	SourceType  *string `gorm:"column:source_type"`
}

type TaskState struct {
	At    *time.Time
	State *string
	Tags  map[string]interface{}
}

type Task struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	CurrentState *TaskState
	Source       *Resource

	States []*TaskState
}

type Storage interface {
	CreateEvaluator(context.Context, *Evaluator) (*Evaluator, error)
	DeleteEvaluator(ctx context.Context, id string) error
	PatchEvaluator(ctx context.Context, id string, evaluator *Evaluator) (*Evaluator, error)
	GetEvaluator(ctx context.Context, id string) (*Evaluator, error)
	ListEvaluators(context.Context, *Evaluator) ([]*Evaluator, error)
	ListEvaluatorsBySource(context.Context, *Resource) ([]*Evaluator, error)
	AddSourcesToEvaluator(ctx context.Context, evaluator_id string, sources []*Resource) error
	RemoveSourcesFromEvaluator(ctx context.Context, evaluator_id string, sources []*Resource) error

	ExistEvaluator(context.Context, *Evaluator) (bool, error)
	ExistOperator(context.Context, *Operator) (bool, error)
}

func NewStorage(driver, uri string, args ...interface{}) (Storage, error) {
	return NewStorageImpl(driver, uri, args...)
}

type TaskStorageFactory func(...interface{}) (TaskStorage, error)

var task_storage_factories map[string]TaskStorageFactory
var task_storage_factories_once sync.Once

func register_task_storage_factory(name string, fty TaskStorageFactory) {
	task_storage_factories_once.Do(func() {
		task_storage_factories = map[string]TaskStorageFactory{}
	})
	task_storage_factories[name] = fty
}

type TaskStorage interface {
	ListTasksBySource(ctx context.Context, src *Resource) ([]*Task, error)
	GetTask(context.Context, string) (*Task, error)
	PatchTask(context.Context, *Task, *TaskState) error
}

func NewTaskStorage(driver string, args ...interface{}) (TaskStorage, error) {
	fty, ok := task_storage_factories[driver]
	if !ok {
		return nil, ErrUnknownTaskStorageDriver
	}

	return fty(args...)
}
