package metathings_evaluatord_storage

import (
	"context"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	otgorm "github.com/smacker/opentracing-gorm"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type StorageImplOption struct {
	IsTraced bool
}

type StorageImpl struct {
	opt     *StorageImplOption
	root_db *gorm.DB
	logger  log.FieldLogger
}

func (stor *StorageImpl) GetRootDBConn() *gorm.DB {
	return stor.root_db
}

func (stor *StorageImpl) GetDBConn(ctx context.Context) *gorm.DB {
	if db := ctx.Value("dbconn"); db != nil {
		return db.(*gorm.DB)
	}

	return stor.GetRootDBConn()
}

func (stor *StorageImpl) get_logger() log.FieldLogger {
	return stor.logger
}

func (stor *StorageImpl) get_lua_descriptor_by_operator_id(db *gorm.DB, operator_id string) (*LuaDescriptor, error) {
	var err error
	var desc LuaDescriptor

	if err = db.First(&desc, "operator_id = ?", operator_id).Error; err != nil {
		return nil, err
	}

	return &desc, nil
}

func (stor *StorageImpl) get_operator_by_evaluator_id(db *gorm.DB, evaluator_id string) (*Operator, error) {
	var err error
	var op Operator

	if err = db.First(&op, "evaluator_id = ?", evaluator_id).Error; err != nil {
		return nil, err
	}

	// SYM:REFACTOR:lua_operator
	switch *op.Driver {
	case "lua":
		fallthrough
	case "default":
		if op.LuaDescriptor, err = stor.get_lua_descriptor_by_operator_id(db, *op.Id); err != nil {
			return nil, err
		}
	}

	return &op, nil
}

func (stor *StorageImpl) get_sources_by_evaluator_id(db *gorm.DB, evaluator_id string) ([]*Resource, error) {
	var err error
	var srcs []*Resource

	esms := []*EvaluatorSourceMapping{}
	if err = db.Find(&esms, "evaluator_id = ?", evaluator_id).Error; err != nil {
		return nil, err
	}

	for _, esm := range esms {
		src := &Resource{
			Id:   esm.SourceId,
			Type: esm.SourceType,
		}
		srcs = append(srcs, src)
	}

	return srcs, nil
}

func (stor *StorageImpl) get_evaluator(db *gorm.DB, id string) (*Evaluator, error) {
	var e Evaluator
	var err error

	if err = db.First(&e, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if e.Operator, err = stor.get_operator_by_evaluator_id(db, id); err != nil {
		return nil, err
	}

	if e.Sources, err = stor.get_sources_by_evaluator_id(db, id); err != nil {
		return nil, err
	}

	return &e, nil
}

func (stor *StorageImpl) list_evaluators(db *gorm.DB, evltr *Evaluator) ([]*Evaluator, error) {
	var evltrs_t []*Evaluator
	var err error

	e := &Evaluator{}
	if evltr.Id != nil {
		e.Id = evltr.Id
	}
	if evltr.Alias != nil {
		e.Alias = evltr.Alias
	}

	if err = db.Select("id").Find(&evltrs_t, e).Error; err != nil {
		return nil, err
	}

	evltr_ids := []string{}
	for _, e := range evltrs_t {
		evltr_ids = append(evltr_ids, *e.Id)
	}

	return stor.list_evaluators_by_ids(db, evltr_ids)
}

func (stor *StorageImpl) list_evaluators_by_source(db *gorm.DB, src *Resource) ([]*Evaluator, error) {
	var esms_t []*EvaluatorSourceMapping
	var err error

	if err = db.Select("evaluator_id").Find(&esms_t, "source_id = ? and source_type = ?", src.Id, src.Type).Error; err != nil {
		return nil, err
	}

	evltr_ids := []string{}
	for _, esm := range esms_t {
		evltr_ids = append(evltr_ids, *esm.EvaluatorId)
	}

	return stor.list_evaluators_by_ids(db, evltr_ids)
}

func (stor *StorageImpl) list_evaluators_by_ids(db *gorm.DB, ids []string) ([]*Evaluator, error) {
	var evltrs []*Evaluator
	for _, id := range ids {
		if evltr, err := stor.get_evaluator(db, id); err != nil {
			return nil, err
		} else {
			evltrs = append(evltrs, evltr)
		}
	}

	return evltrs, nil
}

func (stor *StorageImpl) disassociate_evaluator_to_resources(db *gorm.DB, id string) error {
	if err := db.Delete(&EvaluatorSourceMapping{}, "evaluator_id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

func (stor *StorageImpl) add_sources_to_evaluator(db *gorm.DB, evaluator_id string, sources []*Resource) error {
	for _, src := range sources {
		if err := db.Create(&EvaluatorSourceMapping{
			EvaluatorId: &evaluator_id,
			SourceId:    src.Id,
			SourceType:  src.Type,
		}).Error; err != nil {
			return err
		}
	}

	return nil
}

func (stor *StorageImpl) remove_sources_from_evaluator(db *gorm.DB, evaluator_id string, sources []*Resource) error {
	for _, src := range sources {
		if err := db.Delete(&EvaluatorSourceMapping{}, "evaluator_id = ? and source_id = ? and source_type = ?", evaluator_id, *src.Id, *src.Type).Error; err != nil {
			return err
		}
	}

	return nil
}

func (stor *StorageImpl) delete_operator(db *gorm.DB, id string) error {
	if err := db.Delete(&Operator{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

func (stor *StorageImpl) delete_evalator(db *gorm.DB, id string) error {
	if err := db.Delete(&Evaluator{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

func (stor *StorageImpl) CreateEvaluator(ctx context.Context, e *Evaluator) (*Evaluator, error) {
	logger := stor.get_logger().WithField("evaluator", *e.Id)
	db := stor.GetDBConn(ctx)

	err := db.Transaction(func(tx *gorm.DB) error {
		// SYM:REFACTOR:lua_operator
		switch *e.Operator.Driver {
		case "lua":
			fallthrough
		case "default":
			e.Operator.LuaDescriptor.OperatorId = e.Operator.Id
			if err := tx.Create(e.Operator.LuaDescriptor).Error; err != nil {
				logger.WithError(err).Debugf("failed to create lua descriptor")
				return err
			}
		}

		e.Operator.EvaluatorId = e.Id
		if err := tx.Create(e.Operator).Error; err != nil {
			logger.WithError(err).Debugf("failed to create operator")
			return err
		}

		if err := tx.Create(e).Error; err != nil {
			logger.WithError(err).Debugf("failed to create evaluator")
			return err
		}

		for _, src := range e.Sources {
			if err := tx.Create(&EvaluatorSourceMapping{
				EvaluatorId: e.Id,
				SourceId:    src.Id,
				SourceType:  src.Type,
			}).Error; err != nil {
				logger.WithError(err).Debugf("failed to add source to evaluator")
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	if e, err = stor.get_evaluator(db, *e.Id); err != nil {
		logger.WithError(err).Debugf("failed to get evaluator")
		return nil, err
	}

	logger.Debugf("create evaluator")

	return e, nil
}

func (stor *StorageImpl) DeleteEvaluator(ctx context.Context, id string) error {
	logger := stor.get_logger().WithField("evaluator", id)

	err := stor.GetDBConn(ctx).Transaction(func(tx *gorm.DB) error {
		e, err := stor.get_evaluator(tx, id)
		if err != nil {
			logger.WithError(err).Debugf("failed to get evaluator")
			return err
		}

		if err = stor.disassociate_evaluator_to_resources(tx, id); err != nil {
			logger.WithError(err).Debugf("failed to disassociate evaluator to resources")
			return err
		}

		if err = stor.delete_operator(tx, *e.Operator.Id); err != nil {
			logger.WithError(err).Debugf("failed to delete operator")
			return nil
		}

		if err = stor.delete_evalator(tx, id); err != nil {
			logger.WithError(err).Debugf("failed to delete evaluator")
			return nil
		}

		return nil
	})

	if err != nil {
		return err
	}

	logger.Debugf("delete evaluator")

	return nil
}

func (stor *StorageImpl) PatchEvaluator(ctx context.Context, id string, e *Evaluator) (*Evaluator, error) {
	var err error
	pe := &Evaluator{}
	po := &Operator{}

	db := stor.GetDBConn(ctx)

	logger := stor.get_logger().WithField("evaluator", id)

	se, err := stor.get_evaluator(db, id)
	if err != nil {
		logger.WithError(err).Debugf("failed to get evaluator")
		return nil, err
	}

	if e.Alias != nil {
		pe.Alias = e.Alias
	}
	if e.Description != nil {
		pe.Description = e.Description
	}
	if e.Config != nil {
		pe.Config = e.Config
	}

	o := e.Operator
	if o != nil {
		if o.Alias != nil {
			po.Alias = o.Alias
		}
		if o.Description != nil {
			po.Description = o.Description
		}

		// SYM:REFACTOR:lua_operator
		switch *se.Operator.Driver {
		case "lua":
			fallthrough
		case "default":
			desc := o.LuaDescriptor
			if desc != nil {
				po.LuaDescriptor = &LuaDescriptor{}
				if desc.Code != nil {
					po.LuaDescriptor.Code = desc.Code
				}
			}
		}
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Model(&Evaluator{Id: &id}).Update(pe).Error; err != nil {
			logger.WithError(err).Debugf("failed to patch evaluator")
			return err
		}

		if e, err = stor.get_evaluator(tx, id); err != nil {
			logger.WithError(err).Debugf("failed to get evaluator")
			return err
		}

		if err = tx.Model(&Operator{Id: se.Operator.Id}).Update(po).Error; err != nil {
			logger.WithError(err).Debugf("failed to patch operator")
			return err
		}

		// SYM:REFACTOR:lua_operator
		switch *se.Operator.Driver {
		case "lua":
			fallthrough
		case "default":
			if po.LuaDescriptor != nil {
				if err = tx.Model(&LuaDescriptor{OperatorId: se.Operator.Id}).Update(po.LuaDescriptor).Error; err != nil {
					logger.WithError(err).Debugf("failed to patch lua descriptor")
					return err
				}
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	if e, err = stor.get_evaluator(db, id); err != nil {
		logger.WithError(err).Debugf("failed to get evaluator")
		return nil, err
	}

	logger.Debugf("patch evaluator")

	return e, nil
}

func (stor *StorageImpl) GetEvaluator(ctx context.Context, id string) (*Evaluator, error) {
	logger := stor.get_logger().WithField("evaluator", id)

	e, err := stor.get_evaluator(stor.GetDBConn(ctx), id)
	if err != nil {
		logger.WithError(err).Debugf("failed to get evaluator")
		return nil, err
	}

	logger.Debugf("get evaluator")

	return e, nil
}

func (stor *StorageImpl) ListEvaluators(ctx context.Context, e *Evaluator) ([]*Evaluator, error) {
	var es []*Evaluator
	var err error

	db := stor.GetDBConn(ctx)
	logger := stor.get_logger()

	if es, err = stor.list_evaluators(db, e); err != nil {
		logger.WithError(err).Debugf("failed to list evaluators")
		return nil, err
	}

	logger.Debugf("list evaluators")

	return es, nil
}

func (stor *StorageImpl) ListEvaluatorsBySource(ctx context.Context, src *Resource) ([]*Evaluator, error) {
	var es []*Evaluator
	var err error

	db := stor.GetDBConn(ctx)
	logger := stor.get_logger().WithFields(log.Fields{
		"source_id":   *src.Id,
		"source_type": *src.Type,
	})

	if es, err = stor.list_evaluators_by_source(db, src); err != nil {
		logger.WithError(err).Debugf("failed to list evaluators by source")
		return nil, err
	}

	logger.Debugf("list evaluators by source")

	return es, nil
}

func (stor *StorageImpl) AddSourcesToEvaluator(ctx context.Context, evaluator_id string, sources []*Resource) error {
	logger := stor.get_logger().WithField("evaluator", evaluator_id)

	err := stor.GetDBConn(ctx).Transaction(func(tx *gorm.DB) error {
		if err := stor.add_sources_to_evaluator(tx, evaluator_id, sources); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		logger.WithError(err).Debugf("failed to add sources to evaluator")
		return err
	}

	logger.Debugf("add sources to evaluator")

	return nil
}

func (stor *StorageImpl) RemoveSourcesFromEvaluator(ctx context.Context, evaluator_id string, sources []*Resource) error {
	logger := stor.get_logger().WithField("evaluator", evaluator_id)

	err := stor.GetDBConn(ctx).Transaction(func(tx *gorm.DB) error {
		if err := stor.remove_sources_from_evaluator(tx, evaluator_id, sources); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		logger.WithError(err).Debugf("failed to remove sources from evaluator")
		return err
	}

	logger.Debugf("remove sources from evaluator")

	return nil
}

func (stor *StorageImpl) ExistEvaluator(ctx context.Context, e *Evaluator) (bool, error) {
	var cnt int
	db := stor.GetDBConn(ctx)
	logger := stor.get_logger()

	if err := db.Find(&e).Count(&cnt).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		logger.WithError(err).Debugf("failed to exist evaluator")
		return false, err
	}

	return cnt > 0, nil
}

func (stor *StorageImpl) ExistOperator(ctx context.Context, o *Operator) (bool, error) {
	var cnt int
	db := stor.GetDBConn(ctx)
	logger := stor.get_logger()

	if err := db.Find(&o).Count(&cnt).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		logger.WithError(err).Debugf("failed to exist operator")
		return false, err
	}

	return cnt > 0, nil
}

func new_db(s *StorageImpl, driver, uri string) error {
	var db *gorm.DB
	var err error

	if db, err = gorm.Open(driver, uri); err != nil {
		return err
	}

	if s.opt.IsTraced {
		otgorm.AddGormCallbacks(db)
	}

	s.root_db = db

	return nil
}

func init_db(s *StorageImpl) error {
	if err := s.GetRootDBConn().AutoMigrate(
		&Evaluator{},
		&Operator{},
		&LuaDescriptor{},
		&EvaluatorSourceMapping{},
	).Error; err != nil {
		return err
	}

	return nil
}

func NewStorageImpl(driver, uri string, args ...interface{}) (Storage, error) {
	var err error
	var ok bool
	var opt StorageImplOption
	var logger log.FieldLogger

	if err = opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger": opt_helper.ToLogger(&logger),
		"tracer": func(key string, val interface{}) error {
			if _, ok = val.(opentracing.Tracer); !ok {
				return nil
			}

			opt.IsTraced = true

			return nil
		},
	})(args...); err != nil {
		return nil, err
	}

	s := &StorageImpl{
		opt:    &opt,
		logger: logger,
	}

	if err = new_db(s, driver, uri); err != nil {
		return nil, err
	}

	if err = init_db(s); err != nil {
		return nil, err
	}

	if s.opt.IsTraced {
		return NewTracedStorage(s, s)
	}

	return s, nil
}
