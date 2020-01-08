package metathings_identityd2_policy

import (
	"context"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
)

type Backend interface {
	Enforce(ctx context.Context, sub, obj *storage.Entity, act *storage.Action) error
	CreateGroup(context.Context, *storage.Group) error
	DeleteGroup(context.Context, *storage.Group) error
	AddSubjectToGroup(context.Context, *storage.Group, *storage.Entity) error
	RemoveSubjectFromGroup(context.Context, *storage.Group, *storage.Entity) error
	AddObjectToGroup(context.Context, *storage.Group, *storage.Entity) error
	RemoveObjectFromGroup(context.Context, *storage.Group, *storage.Entity) error
	AddRoleToGroup(context.Context, *storage.Group, *storage.Role) error
	RemoveRoleFromGroup(context.Context, *storage.Group, *storage.Role) error

	AddRoleToEntity(ctx context.Context, ent *storage.Entity, rol *storage.Role) error
	RemoveRoleFromEntity(ctx context.Context, ent *storage.Entity, rol *storage.Role) error
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
