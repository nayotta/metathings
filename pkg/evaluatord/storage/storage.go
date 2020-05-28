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

type Timer struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	Alias       *string  `gorm:"column:alias"`
	Description *string  `gorm:"column:description"`
	Schedule    *string  `gorm:"column:schedule"`
	Timezone    *string  `gorm:"column:timezone"`
	Enabled     *bool    `gorm:"column:enabled"`
	Configs     []string `gorm:"-"`
}

type TimerConfigMapping struct {
	CreatedAt time.Time

	TimerId  *string `gorm:"column:timer_id"`
	ConfigId *string `gorm:"column:config_id"`
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
	ListTasksBySource(ctx context.Context, src *Resource, opts ...ListTasksBySourceOption) ([]*Task, error)
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

type TimerStorageFactory func(...interface{}) (TimerStorage, error)

var timer_storage_factories map[string]TimerStorageFactory
var timer_storage_factories_once sync.Once

func register_timer_storage_factory(name string, fty TimerStorageFactory) {
	timer_storage_factories_once.Do(func() {
		timer_storage_factories = map[string]TimerStorageFactory{}
	})
	timer_storage_factories[name] = fty
}

type TimerStorage interface {
	CreateTimer(context.Context, *Timer) (*Timer, error)
	DeleteTimer(context.Context, string) error
	PatchTimer(context.Context, string, *Timer) (*Timer, error)
	GetTimer(context.Context, string) (*Timer, error)
	ListTimers(context.Context, *Timer) ([]*Timer, error)
	AddConfigsToTimer(ctx context.Context, timer_id string, config_ids []string) error
	RemoveConfigsFromTimer(ctx context.Context, timer_id string, config_ids []string) error
}

func NewTimerStorage(driver string, args ...interface{}) (TimerStorage, error) {
	fty, ok := timer_storage_factories[driver]
	if !ok {
		return nil, ErrUnknownTimerStorageDriver
	}

	return fty(args...)
}
