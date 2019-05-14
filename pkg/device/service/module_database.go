package metathings_device_service

import (
	"time"

	log "github.com/sirupsen/logrus"

	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type ModuleDatabase interface {
	Lookup(name string) (Module, error)
	All() []Module
}

type ModuleDatabaseImpl struct {
	logger log.FieldLogger

	modules map[string]Module
}

func (self *ModuleDatabaseImpl) Lookup(name string) (Module, error) {
	m, ok := self.modules[name]
	if !ok {
		return nil, ErrModuleNotFound
	}

	return m, nil
}

func (self *ModuleDatabaseImpl) All() []Module {
	var modules []Module

	for _, m := range self.modules {
		modules = append(modules, m)
	}

	return modules
}

func NewModuleDatabase(
	modules []*deviced_pb.Module,
	alive_timeout time.Duration,
	logger log.FieldLogger,
) ModuleDatabase {
	db := &ModuleDatabaseImpl{
		logger:  logger,
		modules: make(map[string]Module),
	}

	for _, m := range modules {
		name := m.GetName()
		db.modules[name] = NewModule(db.logger, m, alive_timeout)
		logger.WithFields(log.Fields{
			"name": name,
		}).Debugf("register module to database")
	}

	return db
}
