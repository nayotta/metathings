package metathings_device_service

import (
	"errors"
	"time"

	log "github.com/sirupsen/logrus"

	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

var (
	ErrModuleNotFound = errors.New("module not found")
)

type ModuleDatabase interface {
	Lookup(component, name string) (Module, error)
	All() []Module
}

type ModuleDatabaseImpl struct {
	logger log.FieldLogger

	modules map[string]map[string]Module
}

func (self *ModuleDatabaseImpl) Lookup(component, name string) (Module, error) {
	t, ok := self.modules[component]
	if !ok {
		return nil, ErrModuleNotFound
	}

	m, ok := t[name]
	if !ok {
		return nil, ErrModuleNotFound
	}

	return m, nil
}

func (self *ModuleDatabaseImpl) All() []Module {
	var modules []Module

	for _, v := range self.modules {
		for _, m := range v {
			modules = append(modules, m)
		}
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
		modules: make(map[string]map[string]Module),
	}

	for _, m := range modules {
		component := m.GetComponent()
		name := m.GetName()

		_, ok := db.modules[component]
		if !ok {
			db.modules[component] = make(map[string]Module)
		}

		db.modules[component][name] = NewModule(db.logger, m, alive_timeout)
		logger.WithFields(log.Fields{
			"name":      name,
			"component": component,
		}).Debugf("register module to database")
	}

	return db
}
