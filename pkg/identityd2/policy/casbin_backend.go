package metathings_identityd2_policy

import storage "github.com/nayotta/metathings/pkg/identityd2/storage"

type CasbinBackend struct{}

func (cb *CasbinBackend) CreateGroup(*storage.Group) error {
	panic("unimplemented")
}

func (cb *CasbinBackend) DeleteGroup(*storage.Group) error {
	panic("unimplemented")
}

func (cb *CasbinBackend) AddSubjectToGroup(*storage.Group, *storage.Entity) error {
	panic("unimplemented")
}

func (cb *CasbinBackend) RemoveSubjectFromGroup(*storage.Group, *storage.Entity) error {
	panic("unimplemented")
}

func (cb *CasbinBackend) AddObjectToGroup(*storage.Group, *storage.Entity) error {
	panic("unimplemented")
}

func (cb *CasbinBackend) RemoveObjectFromGroup(*storage.Group, *storage.Entity) error {
	panic("unimplemented")
}

func casbin_backend_factory(args ...interface{}) (Backend, error) {
	panic("unimplemented")
}

func init() {
	register_backend_factory("casbin", casbin_backend_factory)
}
