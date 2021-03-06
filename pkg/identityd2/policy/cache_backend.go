package metathings_identityd2_policy

import (
	"context"

	log "github.com/sirupsen/logrus"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
)

type CacheBackendOption struct {
	Mongo struct {
		Uri        string
		Database   string
		Collection string
	}
}

type CacheBackend struct {
	opt     *CacheBackendOption
	backend Backend
	cache   BackendCache
	logger  log.FieldLogger
}

func (cb *CacheBackend) get_logger() log.FieldLogger {
	return cb.logger
}

func (cb *CacheBackend) Enforce(ctx context.Context, sub, obj *storage.Entity, act *storage.Action) error {
	// try to get cache result in cache
	ret, err := cb.cache.Get(sub, obj, act)
	// failed get cache, do enforce and cache result.
	if err != nil {
		if err != ErrNoCached {
			cb.get_logger().Debugf("failed to get cache")
			return err
		}

		err = cb.backend.Enforce(ctx, sub, obj, act)
		if err != nil {
			if err != ErrPermissionDenied {
				return err
			}
			ret = false
		} else {
			ret = true
		}

		err = cb.cache.Set(sub, obj, act, ret)
		if err != nil {
			cb.get_logger().Debugf("failed to set cache")
			return err
		}
	}

	if ret {
		err = nil
	} else {
		err = ErrPermissionDenied
	}

	return err
}

func (cb *CacheBackend) CreateGroup(ctx context.Context, grp *storage.Group) error {
	return cb.backend.CreateGroup(ctx, grp)
}

func (cb *CacheBackend) DeleteGroup(ctx context.Context, grp *storage.Group) error {
	err := cb.backend.DeleteGroup(ctx, grp)
	if err != nil {
		return err
	}

	for _, sub := range grp.Subjects {
		if err = cb.cache.Remove("subject", sub); err != nil {
			cb.get_logger().WithField("subject", *sub.Id).Debugf("failed to remove cache")
		}
	}

	for _, obj := range grp.Objects {
		if err = cb.cache.Remove("object", obj); err != nil {
			cb.get_logger().WithField("object", *obj.Id).Debugf("failed to remove cache")
		}
	}

	rmacts := map[string]interface{}{}
	for _, rol := range grp.Roles {
		for _, act := range rol.Actions {
			if _, ok := rmacts[*act.Name]; !ok {
				if err = cb.cache.Remove("action", act); err != nil {
					cb.get_logger().WithField("action", *act.Name).Debugf("failed to remove cache")
				}
				rmacts[*act.Name] = nil
			}
		}
	}

	return nil
}

func (cb *CacheBackend) AddSubjectToGroup(ctx context.Context, grp *storage.Group, sub *storage.Entity) error {
	err := cb.backend.AddSubjectToGroup(ctx, grp, sub)
	if err != nil {
		return err
	}

	if err = cb.cache.Remove("subject", sub); err != nil {
		cb.get_logger().WithField("subject", *sub.Id).Debugf("failed to remove cache")
	}

	return nil
}

func (cb *CacheBackend) RemoveSubjectFromGroup(ctx context.Context, grp *storage.Group, sub *storage.Entity) error {
	err := cb.backend.RemoveSubjectFromGroup(ctx, grp, sub)
	if err != nil {
		return err
	}

	if err = cb.cache.Remove("subject", sub); err != nil {
		cb.get_logger().WithField("subject", *sub.Id).Debugf("failed to remove cache")
	}

	return nil
}

func (cb *CacheBackend) AddObjectToGroup(ctx context.Context, grp *storage.Group, obj *storage.Entity) error {
	err := cb.backend.AddObjectToGroup(ctx, grp, obj)
	if err != nil {
		return err
	}

	if err = cb.cache.Remove("object", obj); err != nil {
		cb.get_logger().WithField("object", *obj.Id).Debugf("failed to remove cache")
	}

	return nil
}

func (cb *CacheBackend) RemoveObjectFromGroup(ctx context.Context, grp *storage.Group, obj *storage.Entity) error {
	err := cb.backend.RemoveObjectFromGroup(ctx, grp, obj)
	if err != nil {
		return err
	}

	if err = cb.cache.Remove("object", obj); err != nil {
		cb.get_logger().WithField("object", *obj.Id).Debugf("failed to remove cache")
	}

	return nil
}

func (cb *CacheBackend) AddRoleToGroup(ctx context.Context, grp *storage.Group, rol *storage.Role) error {
	err := cb.backend.AddRoleToGroup(ctx, grp, rol)
	if err != nil {
		return err
	}

	for _, act := range rol.Actions {
		if err = cb.cache.Remove("action", act); err != nil {
			cb.get_logger().WithField("action", *act.Name).Debugf("failed to remove cache")
		}
	}

	return nil
}

func (cb *CacheBackend) RemoveRoleFromGroup(ctx context.Context, grp *storage.Group, rol *storage.Role) error {
	err := cb.backend.RemoveRoleFromGroup(ctx, grp, rol)
	if err != nil {
		return err
	}

	for _, act := range rol.Actions {
		if err = cb.cache.Remove("action", act); err != nil {
			cb.get_logger().WithField("action", *act.Name).Debugf("failed to remove cache")
		}
	}

	return nil
}

func (cb *CacheBackend) AddRoleToEntity(ctx context.Context, ent *storage.Entity, rol *storage.Role) error {
	return cb.backend.AddRoleToEntity(ctx, ent, rol)
}

func (cb *CacheBackend) RemoveRoleFromEntity(ctx context.Context, ent *storage.Entity, rol *storage.Role) error {
	return cb.backend.RemoveRoleFromEntity(ctx, ent, rol)
}

func cache_backend_factory(args ...interface{}) (Backend, error) {
	var b Backend
	var logger log.FieldLogger
	var ok bool

	opt := &CacheBackendOption{}
	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"mongo_uri":        opt_helper.ToString(&opt.Mongo.Uri),
		"mongo_database":   opt_helper.ToString(&opt.Mongo.Database),
		"mongo_collection": opt_helper.ToString(&opt.Mongo.Collection),
		"backend": func(key string, val interface{}) error {
			b, ok = val.(Backend)
			if !ok {
				return opt_helper.InvalidArgument("backend")
			}
			return nil
		},
		"logger": opt_helper.ToLogger(&logger),
	})(args...); err != nil {
		return nil, err
	}

	c, err := NewBackendCache(
		"mongo",
		"mongo_uri", opt.Mongo.Uri,
		"mongo_database", opt.Mongo.Database,
		"mongo_collection", opt.Mongo.Collection,
		"logger", logger,
	)
	if err != nil {
		return nil, err
	}

	cb := &CacheBackend{
		opt:     opt,
		cache:   c,
		backend: b,
		logger:  logger,
	}

	return cb, nil
}

func init() {
	register_backend_factory("cache", cache_backend_factory)
}
