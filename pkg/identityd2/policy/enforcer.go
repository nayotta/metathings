package metathings_identityd2_policy

import (
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
)

type Enforcer interface {
	Initialize() error
	Enforce(domain, group, subject, object, action interface{}) error
	AddGroup(domain, group string) error
	RemoveGroup(domain, group string) error
	AddSubjectToRole(subject, role string) error
	RemoveSubjectFromRole(subject, role string) error
	AddObjectToKind(object, kind string) error
	RemoveObjectFromKind(object, kind string) error
}

type Backend interface {
	CreateGroup(*storage.Group) error
	DeleteGroup(*storage.Group) error
	AddSubjectToGroup(*storage.Group, *storage.Entity) error
	RemoveSubjectFromGroup(*storage.Group, *storage.Entity) error
	AddObjectToGroup(*storage.Group, *storage.Entity) error
	RemoveObjectFromGroup(*storage.Group, *storage.Entity) error
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
