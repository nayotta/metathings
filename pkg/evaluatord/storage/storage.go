package metathings_evaluatord_storage

import (
	"context"
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
