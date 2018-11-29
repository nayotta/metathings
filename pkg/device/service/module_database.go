package metathings_device_service

import (
	"errors"
	"time"

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

func NewModuleDatabase(modules []*deviced_pb.Module, alive_timeout time.Duration) ModuleDatabase {
	db := &ModuleDatabaseImpl{
		modules: make(map[string]map[string]Module),
	}

	for _, m := range modules {
		component := m.GetComponent()
		name := m.GetName()

		_, ok := db.modules[component]
		if !ok {
			db.modules[component] = make(map[string]Module)
		}

		db.modules[component][name] = NewModule(m, alive_timeout)
	}

	return db
}
