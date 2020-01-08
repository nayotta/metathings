package opentracing_storage_helper

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"
	otgorm "github.com/smacker/opentracing-gorm"
)

type RootDBConnGetter interface {
	GetRootDBConn() *gorm.DB
}

type BaseTracedStorage struct {
	dbconn RootDBConnGetter
}

func (s *BaseTracedStorage) TraceWrapper(ctx context.Context, name string) (span opentracing.Span, new_ctx context.Context) {
	span, new_ctx = opentracing.StartSpanFromContext(ctx, name)

	db := otgorm.SetSpanToGorm(new_ctx, s.dbconn.GetRootDBConn())
	new_ctx = context.WithValue(new_ctx, "dbconn", db)

	return
}

func NewBaseTracedStorage(dbconn RootDBConnGetter) *BaseTracedStorage {
	return &BaseTracedStorage{dbconn: dbconn}
}
