package metathings_identityd2_policy

import (
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
)

type Backend interface {
	Enforce(sub, obj *storage.Entity, act *storage.Action) error
	CreateGroup(*storage.Group) error
	DeleteGroup(*storage.Group) error
	AddSubjectToGroup(*storage.Group, *storage.Entity) error
	RemoveSubjectFromGroup(*storage.Group, *storage.Entity) error
	AddObjectToGroup(*storage.Group, *storage.Entity) error
	RemoveObjectFromGroup(*storage.Group, *storage.Entity) error
	AddRoleToGroup(*storage.Group, *storage.Role) error
	RemoveRoleFromGroup(*storage.Group, *storage.Role) error
}

type BackendFactory func(...interface{}) (Backend, error)

var backend_factories map[string]BackendFactory

func register_backend_factory(name string, fty BackendFactory) {
	backend_factories[name] = fty
}

func NewBackend(name string, args ...interface{}) (Backend, error) {
	var fty BackendFactory
	var ok bool

	if fty, ok = backend_factories[name]; !ok {
		return nil, ErrInvalidBackendDriver
	}

	return fty(args...)
}

func init() {
	backend_factories = make(map[string]BackendFactory)
}
