package metathings_evaluatord_storage

import (
	"context"

	"github.com/jinzhu/gorm"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type DefaultTimerStorageOption struct {
	IsTraced bool
	Driver   string
	Uri      string
}

type DefaultTimerStorage struct {
	opt     *DefaultTimerStorageOption
	root_db *gorm.DB
	logger  logrus.FieldLogger
}

func (s *DefaultTimerStorage) GetRootDBConn() *gorm.DB {
	return s.root_db
}

func (s *DefaultTimerStorage) GetDBConn(ctx context.Context) *gorm.DB {
	if db := ctx.Value("dbconn"); db != nil {
		return db.(*gorm.DB)
	}

	return s.GetRootDBConn()
}

func (s *DefaultTimerStorage) get_logger() logrus.FieldLogger {
	return s.logger
}

func (s *DefaultTimerStorage) get_configs_by_timer_id(db *gorm.DB, id string) ([]string, error) {
	var err error
	var tcms []*TimerConfigMapping
	var cfgs []string

	if err = db.Find(&tcms, "timer_id = ?", id).Error; err != nil {
		return nil, err
	}

	for _, tcm := range tcms {
		cfgs = append(cfgs, *tcm.ConfigId)
	}

	return cfgs, nil
}

func (s *DefaultTimerStorage) get_timer(db *gorm.DB, id string) (*Timer, error) {
	var t Timer
	var err error

	if err = db.First(&t, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if t.Configs, err = s.get_configs_by_timer_id(db, id); err != nil {
		return nil, err
	}

	return &t, nil
}

func (s *DefaultTimerStorage) list_timers(db *gorm.DB, tmr *Timer) ([]*Timer, error) {
	var timers_t []*Timer
	var err error

	t := &Timer{}
	if tmr.Id != nil {
		t.Id = tmr.Id
	}
	if tmr.Alias != nil {
		t.Alias = tmr.Alias
	}
	if tmr.Schedule != nil {
		t.Schedule = tmr.Schedule
	}
	if tmr.Timezone != nil {
		t.Timezone = tmr.Timezone
	}
	if tmr.Enabled != nil {
		t.Enabled = tmr.Enabled
	}

	if err = db.Select("id").Find(&timers_t, t).Error; err != nil {
		return nil, err
	}

	var tmr_ids []string
	for _, tmr := range timers_t {
		tmr_ids = append(tmr_ids, *tmr.Id)
	}

	return s.list_timers_by_ids(db, tmr_ids)
}

func (s *DefaultTimerStorage) list_timers_by_ids(db *gorm.DB, ids []string) ([]*Timer, error) {
	var tmrs []*Timer

	for _, id := range ids {
		if tmr, err := s.get_timer(db, id); err != nil {
			return nil, err
		} else {
			tmrs = append(tmrs, tmr)
		}
	}

	return tmrs, nil
}

func (s *DefaultTimerStorage) disassociate_timer_to_configs(db *gorm.DB, id string) error {
	if err := db.Delete(&TimerConfigMapping{}, "timer_id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

func (s *DefaultTimerStorage) delete_timer(db *gorm.DB, id string) error {
	if err := db.Delete(&Timer{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

func (s *DefaultTimerStorage) add_configs_to_timer(db *gorm.DB, timer_id string, config_ids []string) error {
	for _, config_id := range config_ids {
		if err := db.Create(&TimerConfigMapping{
			TimerId:  &timer_id,
			ConfigId: &config_id,
		}).Error; err != nil {
			return err
		}
	}

	return nil
}

func (s *DefaultTimerStorage) remove_configs_from_timer(db *gorm.DB, timer_id string, config_ids []string) error {
	for _, config_id := range config_ids {
		if err := db.Delete(&TimerConfigMapping{}, "timer_id = ? and config_id = ?", timer_id, config_id).Error; err != nil {
			return err
		}
	}

	return nil
}

func (s *DefaultTimerStorage) CreateTimer(ctx context.Context, t *Timer) (*Timer, error) {
	id := *t.Id
	logger := s.get_logger().WithField("timer", id)
	db := s.GetDBConn(ctx)

	if err := db.Create(t).Error; err != nil {
		logger.WithError(err).Debugf("failed to create timer")
		return nil, err
	}

	t, err := s.get_timer(db, id)
	if err != nil {
		logger.WithError(err).Debugf("failed to get timer")
		return nil, err
	}

	logger.Debugf("create timer")

	return t, nil
}

func (s *DefaultTimerStorage) DeleteTimer(ctx context.Context, id string) error {
	logger := s.get_logger().WithField("timer", id)
	db := s.GetDBConn(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := s.disassociate_timer_to_configs(tx, id); err != nil {
			logger.WithError(err).Debugf("failed to disassociate timer to configs")
			return err
		}

		if err := s.delete_timer(tx, id); err != nil {
			logger.WithError(err).Debugf("failed to delete timer")
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	logger.Debugf("delete timer")

	return nil
}

func (s *DefaultTimerStorage) PatchTimer(ctx context.Context, id string, t *Timer) (*Timer, error) {
	pt := &Timer{}

	db := s.GetDBConn(ctx)
	logger := s.get_logger().WithField("timer", id)

	if t.Alias != nil {
		pt.Alias = t.Alias
	}
	if t.Description != nil {
		pt.Description = t.Description
	}
	if t.Schedule != nil {
		pt.Schedule = t.Schedule
	}
	if t.Timezone != nil {
		pt.Timezone = t.Timezone
	}
	if t.Enabled != nil {
		pt.Enabled = t.Enabled
	}

	if err := db.Model(&Timer{Id: &id}).Update(pt).Error; err != nil {
		logger.WithError(err).Debugf("failed to patch timer")
		return nil, err
	}

	t, err := s.get_timer(db, id)
	if err != nil {
		logger.WithError(err).Debugf("failed to get timer")
		return nil, err
	}

	logger.Debugf("patch timer")

	return t, nil
}

func (s *DefaultTimerStorage) GetTimer(ctx context.Context, id string) (*Timer, error) {
	db := s.GetDBConn(ctx)
	logger := s.get_logger().WithField("timer", id)

	t, err := s.get_timer(db, id)
	if err != nil {
		logger.WithError(err).Debugf("failed to get timer")
		return nil, err
	}

	logger.Debugf("get timer")

	return t, nil
}

func (s *DefaultTimerStorage) ListTimers(ctx context.Context, t *Timer) ([]*Timer, error) {
	db := s.GetDBConn(ctx)
	logger := s.get_logger()

	ts, err := s.list_timers(db, t)
	if err != nil {
		logger.WithError(err).Debugf("failed to list timers")
		return nil, err
	}

	logger.Debugf("list timers")

	return ts, nil
}

func (s *DefaultTimerStorage) AddConfigsToTimer(ctx context.Context, timer_id string, config_ids []string) error {
	logger := s.get_logger().WithField("timer", timer_id)

	err := s.GetDBConn(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.add_configs_to_timer(tx, timer_id, config_ids); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		logger.WithError(err).Debugf("failed to add configs to timer")
		return err
	}

	logger.Debugf("add configs to timer")

	return nil
}

func (s *DefaultTimerStorage) RemoveConfigsFromTimer(ctx context.Context, timer_id string, config_ids []string) error {
	logger := s.get_logger().WithField("timer", timer_id)

	err := s.GetDBConn(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.remove_configs_from_timer(tx, timer_id, config_ids); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		logger.WithError(err).Debugf("failed to remove configs from timer")
		return err
	}

	logger.Debugf("remove configs from timer")

	return nil
}

func (s *DefaultTimerStorage) ExistTimer(ctx context.Context, t *Timer) (bool, error) {
	var cnt int
	db := s.GetDBConn(ctx)
	logger := s.get_logger()

	if err := db.Find(&t).Count(&cnt).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		logger.WithError(err).Debugf("failed to exist timer")
		return false, err
	}

	return cnt > 0, nil
}

func NewDefaultTimerStorage(args ...interface{}) (TimerStorage, error) {
	var err error
	var ok bool
	var opt DefaultTimerStorageOption
	var logger logrus.FieldLogger

	if err = opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger":      opt_helper.ToLogger(&logger),
		"gorm_driver": opt_helper.ToString(&opt.Driver),
		"gorm_uri":    opt_helper.ToString(&opt.Uri),
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

	db, err := gorm.Open(opt.Driver, opt.Uri)
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(
		&Timer{},
		&TimerConfigMapping{},
	).Error; err != nil {
		return nil, err
	}

	s := &DefaultTimerStorage{
		opt:     &opt,
		logger:  logger,
		root_db: db,
	}

	if opt.IsTraced {
		return NewTracedTimerStorage(s, s)
	}

	return s, nil
}

func init() {
	register_timer_storage_factory("default", NewDefaultTimerStorage)
}
