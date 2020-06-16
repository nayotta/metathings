package metathings_identityd2_storage

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	otgorm "github.com/smacker/opentracing-gorm"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	storage_helper "github.com/nayotta/metathings/pkg/common/storage"
)

var (
	SYSTEM_CONFIG_INITIALIZE = "init"
)

type StorageImplOption struct {
	IsTraced bool
}

type StorageImpl struct {
	opt     *StorageImplOption
	root_db *gorm.DB
	logger  log.FieldLogger
}

func (self *StorageImpl) GetRootDBConn() *gorm.DB {
	return self.root_db
}

func (self *StorageImpl) GetDBConn(ctx context.Context) *gorm.DB {
	if db := ctx.Value("dbconn"); db != nil {
		return db.(*gorm.DB)
	}

	return self.GetRootDBConn()
}

func (self *StorageImpl) list_view_children_domains_by_domain_id(ctx context.Context, id string) ([]*Domain, error) {
	var doms []*Domain
	var err error

	if err = self.GetDBConn(ctx).Select("id").Where("parent_id = ?", id).Take(&doms).Error; err != nil {
		return nil, err
	}

	return doms, nil
}

func (self *StorageImpl) get_domain(ctx context.Context, id string) (*Domain, error) {
	var err error
	var dom Domain

	if err = self.GetDBConn(ctx).First(&dom, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if dom.ParentId != nil && *dom.ParentId != "" {
		dom.Parent = &Domain{
			Id: dom.ParentId,
		}
	}

	if dom.Children, err = self.list_view_children_domains_by_domain_id(ctx, id); err != nil {
		return nil, err
	}

	return &dom, nil
}

func (self *StorageImpl) list_domains(ctx context.Context, dom *Domain) ([]*Domain, error) {
	var err error
	var doms_t []*Domain

	d := &Domain{}
	if dom.Id != nil {
		d.Id = dom.Id
	}
	if dom.Name != nil {
		d.Name = dom.Name
	}
	if dom.Alias != nil {
		d.Alias = dom.Alias
	}
	if dom.ParentId != nil {
		d.ParentId = dom.ParentId
	}
	if dom.Parent != nil && dom.Parent.Id != nil {
		d.ParentId = dom.Parent.Id
	}

	if err = self.GetDBConn(ctx).Select("id").Find(&doms_t, d).Error; err != nil {
		return nil, err
	}

	doms := []*Domain{}
	for _, d = range doms_t {
		if d, err = self.get_domain(ctx, *d.Id); err != nil {
			return nil, err
		}
		doms = append(doms, d)
	}

	return doms, nil
}

func (self *StorageImpl) CreateDomain(ctx context.Context, dom *Domain) (*Domain, error) {
	var err error

	if err = self.GetDBConn(ctx).Create(dom).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to create domain")
		return nil, err
	}

	if dom.ExtraHelper != nil {
		if err = storage_helper.UpdateExtra(self.GetDBConn(ctx), &Domain{Id: dom.Id}, dom.ExtraHelper); err != nil {
			self.logger.WithError(err).Debugf("failed to update extra filed")
		}
	}

	if dom, err = self.get_domain(ctx, *dom.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get domain")
		return nil, err
	}

	self.logger.WithField("id", *dom.Id).Debugf("create domain")

	return dom, nil
}

func (self *StorageImpl) DeleteDomain(ctx context.Context, id string) error {
	if err := self.GetDBConn(ctx).Delete(&Domain{}, "id = ?", id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to delete domain")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete domain")

	return nil
}

func (self *StorageImpl) PatchDomain(ctx context.Context, id string, domain *Domain) (*Domain, error) {
	var err error
	var dom *Domain
	var domNew Domain

	if domain.Alias != nil {
		domNew.Alias = domain.Alias
	}

	if err = self.GetDBConn(ctx).Model(&Domain{Id: &id}).Update(domNew).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to patch domain")
		return nil, err
	}

	if domain.ExtraHelper != nil {
		if err = storage_helper.UpdateExtra(self.GetDBConn(ctx), &Domain{Id: &id}, domain.ExtraHelper); err != nil {
			self.logger.WithError(err).Debugf("failed to update extra filed")
		}
	}

	if dom, err = self.get_domain(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to get domain view")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("patch domain")

	return dom, nil
}

func (self *StorageImpl) GetDomain(ctx context.Context, id string) (*Domain, error) {
	var err error
	var dom *Domain

	if dom, err = self.get_domain(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to get domain")
		return nil, err
	}

	self.logger.WithField("id", *dom.Id).Debugf("get domain")

	return dom, nil
}

func (self *StorageImpl) ListDomains(ctx context.Context, dom *Domain) ([]*Domain, error) {
	var doms []*Domain
	var err error

	if doms, err = self.list_domains(ctx, dom); err != nil {
		self.logger.WithError(err).Debugf("failed to list domains")
		return nil, err
	}

	self.logger.Debugf("list domains")

	return doms, nil
}

func (self *StorageImpl) AddEntityToDomain(ctx context.Context, domain_id, entity_id string) error {
	m := &EntityDomainMapping{
		DomainId: &domain_id,
		EntityId: &entity_id,
	}

	if err := self.GetDBConn(ctx).Create(m).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to add entity to domain")
		return err
	}

	self.logger.WithFields(log.Fields{
		"entity_id": entity_id,
		"domain_id": domain_id,
	}).Debugf("add entity to domain")

	return nil
}

func (self *StorageImpl) RemoveEntityFromDomain(ctx context.Context, domain_id, entity_id string) error {
	if err := self.GetDBConn(ctx).Delete(&EntityDomainMapping{}, "domain_id = ? and entity_id = ?", domain_id, entity_id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to remove entity from domain")
		return err
	}

	self.logger.WithFields(log.Fields{
		"entity_id": entity_id,
		"domain_id": domain_id,
	}).Debugf("remove entity from domain")

	return nil
}

func (self *StorageImpl) get_action(ctx context.Context, id string) (*Action, error) {
	var err error
	var act Action

	if err = self.GetDBConn(ctx).First(&act, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &act, nil
}

func (self *StorageImpl) list_actions(ctx context.Context, act *Action) ([]*Action, error) {
	var err error
	var acts_t []*Action

	a := &Action{}
	if act.Id != nil {
		a.Id = act.Id
	}
	if act.Name != nil {
		a.Name = act.Name
	}
	if act.Alias != nil {
		a.Alias = act.Alias
	}

	if err = self.GetDBConn(ctx).Select("id").Find(&acts_t, a).Error; err != nil {
		return nil, err
	}

	acts := []*Action{}
	for _, a = range acts_t {
		if a, err = self.get_action(ctx, *a.Id); err != nil {
			return nil, err
		}

		acts = append(acts, a)
	}

	return acts, nil
}

func (self *StorageImpl) list_actions_by_view_actions(ctx context.Context, xs []*Action) ([]*Action, error) {
	var err error
	var y *Action
	var ys []*Action

	for _, x := range xs {
		if y, err = self.get_action(ctx, *x.Id); err != nil {
			return nil, err
		}
		ys = append(ys, y)
	}

	return ys, nil
}

func (self *StorageImpl) CreateAction(ctx context.Context, act *Action) (*Action, error) {
	var err error

	if err = self.GetDBConn(ctx).Create(act).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to create action")
		return nil, err
	}

	if act.ExtraHelper != nil {
		if err = storage_helper.UpdateExtra(self.GetDBConn(ctx), &Action{Id: act.Id}, act.ExtraHelper); err != nil {
			self.logger.WithError(err).Debugf("failed to update extra filed")
		}
	}

	if act, err = self.get_action(ctx, *act.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get action")
		return nil, err
	}

	return act, nil
}

func (self *StorageImpl) DeleteAction(ctx context.Context, id string) error {
	var err error

	if err = self.GetDBConn(ctx).Delete(&Action{}, "id = ?", id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to delete action")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete action")

	return nil
}

func (self *StorageImpl) PatchAction(ctx context.Context, id string, action *Action) (*Action, error) {
	var err error
	var act *Action
	var actNew Action

	if action.Alias != nil {
		actNew.Alias = action.Alias
	}

	if action.Description != nil {
		actNew.Description = action.Description
	}

	if err = self.GetDBConn(ctx).Model(&Action{Id: &id}).Update(actNew).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to patch action")
		return nil, err
	}

	if action.ExtraHelper != nil {
		if err = storage_helper.UpdateExtra(self.GetDBConn(ctx), &Action{Id: &id}, action.ExtraHelper); err != nil {
			self.logger.WithError(err).Debugf("failed to update extra filed")
		}
	}

	if act, err = self.get_action(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to get action view")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("patch action")

	return act, nil
}

func (self *StorageImpl) GetAction(ctx context.Context, id string) (*Action, error) {
	var err error
	var act *Action

	if act, err = self.get_action(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to get action")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("get action")

	return act, nil
}

func (self *StorageImpl) ListActions(ctx context.Context, act *Action) ([]*Action, error) {
	var err error
	var acts []*Action

	if acts, err = self.list_actions(ctx, act); err != nil {
		self.logger.WithError(err).Debugf("failed to list actions")
		return nil, err
	}

	self.logger.Debugf("list actions")

	return acts, nil
}

func (self *StorageImpl) get_role(ctx context.Context, id string) (*Role, error) {
	var role Role
	var act_role_maps []*ActionRoleMapping
	var err error

	if err = self.GetDBConn(ctx).First(&role, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if err = self.GetDBConn(ctx).Select("action_id").Find(&act_role_maps, "role_id = ?", id).Error; err != nil {
		return nil, err
	}
	for _, m := range act_role_maps {
		role.Actions = append(role.Actions, &Action{Id: m.ActionId})
	}

	return &role, nil
}

func (self *StorageImpl) list_roles(ctx context.Context, role *Role) ([]*Role, error) {
	var err error
	var roles_t []*Role

	r := &Role{}
	if role.Id != nil {
		r.Id = role.Id
	}
	if role.Name != nil {
		r.Name = role.Name
	}
	if role.Alias != nil {
		r.Alias = role.Alias
	}

	if err = self.GetDBConn(ctx).Select("id").Find(&roles_t, r).Error; err != nil {
		return nil, err
	}

	var roles []*Role
	for _, r = range roles_t {
		if r, err = self.get_role(ctx, *r.Id); err != nil {
			return nil, err
		}
		roles = append(roles, r)
	}

	return roles, nil
}

func (self *StorageImpl) CreateRole(ctx context.Context, role *Role) (*Role, error) {
	var err error

	if err = self.GetDBConn(ctx).Create(role).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to create role")
		return nil, err
	}

	if role, err = self.get_role(ctx, *role.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get role")
		return nil, err
	}

	if role.ExtraHelper != nil {
		if err = storage_helper.UpdateExtra(self.GetDBConn(ctx), &Role{Id: role.Id}, role.ExtraHelper); err != nil {
			self.logger.WithError(err).Debugf("failed to update extra filed")
		}
	}

	self.logger.WithFields(log.Fields{}).Debugf("create role")

	return role, nil
}

func (self *StorageImpl) DeleteRole(ctx context.Context, id string) error {
	var err error

	if err = self.GetDBConn(ctx).Delete(&Role{}, "id = ?", id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to delete role")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete role")

	return nil
}

func (self *StorageImpl) PatchRole(ctx context.Context, id string, role *Role) (*Role, error) {
	var err error
	var rol *Role
	var rolNew Role

	if role.Alias != nil {
		rolNew.Alias = role.Alias
	}
	if role.Description != nil {
		rolNew.Description = role.Description
	}

	if err = self.GetDBConn(ctx).Model(&Role{Id: &id}).Update(rolNew).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to patch role")
		return nil, err
	}

	if role.ExtraHelper != nil {
		if err = storage_helper.UpdateExtra(self.GetDBConn(ctx), &Role{Id: &id}, role.ExtraHelper); err != nil {
			self.logger.WithError(err).Debugf("failed to update extra filed")
		}
	}

	if rol, err = self.get_role(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to get role view")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("patch role")

	return rol, nil
}

func (self *StorageImpl) GetRole(ctx context.Context, id string) (*Role, error) {
	var role *Role
	var err error

	if role, err = self.get_role(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to get role")
		return nil, err
	}

	return role, nil
}

func (self *StorageImpl) GetRoleWithFullActions(ctx context.Context, id string) (*Role, error) {
	var role *Role
	var err error

	if role, err = self.get_role(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to get role")
		return nil, err
	}

	if role.Actions, err = self.list_actions_by_view_actions(ctx, role.Actions); err != nil {
		self.logger.WithError(err).Debugf("failed to list actions by view actions")
		return nil, err
	}

	return role, nil
}

func (self *StorageImpl) ListRoles(ctx context.Context, role *Role) ([]*Role, error) {
	var roles []*Role
	var err error

	if roles, err = self.list_roles(ctx, role); err != nil {
		self.logger.WithError(err).Debugf("failed to list roles")
		return nil, err
	}

	self.logger.Debugf("list roles")

	return roles, nil
}

func (self *StorageImpl) AddActionToRole(ctx context.Context, role_id, action_id string) error {
	var err error

	m := &ActionRoleMapping{
		ActionId: &action_id,
		RoleId:   &role_id,
	}

	if err = self.GetDBConn(ctx).Create(m).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to add action to role")
		return err
	}

	self.logger.WithFields(log.Fields{
		"action_id": action_id,
		"role_id":   role_id,
	}).Debugf("add action to role")

	return nil
}

func (self *StorageImpl) RemoveActionFromRole(ctx context.Context, role_id, action_id string) error {
	var err error

	if err = self.GetDBConn(ctx).Delete(&ActionRoleMapping{}, "role_id = ? and action_id = ?", role_id, action_id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to remove action from role")
		return err
	}

	self.logger.WithFields(log.Fields{
		"action_id": action_id,
		"role_id":   role_id,
	}).Debugf("remove action from role")

	return nil
}

func (self *StorageImpl) list_view_domains_by_entity_id(ctx context.Context, id string) ([]*Domain, error) {
	var err error

	var ent_dom_maps []*EntityDomainMapping
	if err = self.GetDBConn(ctx).Find(&ent_dom_maps, "entity_id = ?", id).Error; err != nil {
		return nil, err
	}

	var doms []*Domain
	for _, m := range ent_dom_maps {
		doms = append(doms, &Domain{
			Id: m.DomainId,
		})
	}

	return doms, nil
}

func (self *StorageImpl) list_view_groups_by_entity_id(ctx context.Context, id string) ([]*Group, error) {
	var err error

	grps_m := map[string]bool{}

	var sub_grp_maps []*SubjectGroupMapping
	if err = self.GetDBConn(ctx).Find(&sub_grp_maps, "subject_id = ?", id).Error; err != nil {
		return nil, err
	}
	for _, m := range sub_grp_maps {
		grps_m[*m.GroupId] = true
	}

	var obj_grp_maps []*ObjectGroupMapping
	if err = self.GetDBConn(ctx).Find(&obj_grp_maps, "object_id = ?", id).Error; err != nil {
		return nil, err
	}
	for _, m := range obj_grp_maps {
		grps_m[*m.GroupId] = true
	}

	var grps []*Group
	for grp_id, _ := range grps_m {
		grpIdStr := grp_id
		grps = append(grps, &Group{Id: &grpIdStr})
	}

	return grps, nil
}

func (self *StorageImpl) list_view_roles_by_entity_id(ctx context.Context, id string) ([]*Role, error) {
	var err error

	var ent_role_maps []*EntityRoleMapping
	if err = self.GetDBConn(ctx).Find(&ent_role_maps, "entity_id = ?", id).Error; err != nil {
		return nil, err
	}

	var roles []*Role
	for _, m := range ent_role_maps {
		roles = append(roles, &Role{Id: m.RoleId})
	}

	return roles, nil
}

func (self *StorageImpl) get_entity(ctx context.Context, id string) (*Entity, error) {
	var ent Entity
	var err error

	if err = self.GetDBConn(ctx).First(&ent, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if ent.Domains, err = self.list_view_domains_by_entity_id(ctx, id); err != nil {
		return nil, err
	}

	if ent.Groups, err = self.list_view_groups_by_entity_id(ctx, id); err != nil {
		return nil, err
	}

	if ent.Roles, err = self.list_view_roles_by_entity_id(ctx, id); err != nil {
		return nil, err
	}

	return &ent, nil
}

func (self *StorageImpl) get_entity_by_name(ctx context.Context, name string) (*Entity, error) {
	var ent Entity
	var err error

	if err = self.GetDBConn(ctx).First(&ent, "name = ?", name).Error; err != nil {
		return nil, err
	}

	if ent.Domains, err = self.list_view_domains_by_entity_id(ctx, *ent.Id); err != nil {
		return nil, err
	}

	if ent.Groups, err = self.list_view_groups_by_entity_id(ctx, *ent.Id); err != nil {
		return nil, err
	}

	if ent.Roles, err = self.list_view_roles_by_entity_id(ctx, *ent.Id); err != nil {
		return nil, err
	}

	return &ent, nil
}

func (self *StorageImpl) list_entities(ctx context.Context, ent *Entity) ([]*Entity, error) {
	var ents_t []*Entity
	var err error

	e := &Entity{}
	if ent.Id != nil {
		e.Id = ent.Id
	}
	if ent.Name != nil {
		e.Name = ent.Name
	}
	if ent.Alias != nil {
		e.Alias = ent.Alias
	}

	if err = self.GetDBConn(ctx).Select("id").Find(&ents_t, e).Error; err != nil {
		return nil, err
	}

	var ents []*Entity
	for _, e := range ents_t {
		if ent, err = self.get_entity(ctx, *e.Id); err != nil {
			return nil, err
		}
		ents = append(ents, ent)
	}

	return ents, nil
}

func (self *StorageImpl) CreateEntity(ctx context.Context, ent *Entity) (*Entity, error) {
	var err error

	if err = self.GetDBConn(ctx).Create(ent).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to create entity")
		return nil, err
	}

	if ent.ExtraHelper != nil {
		if err = storage_helper.UpdateExtra(self.GetDBConn(ctx), &Entity{Id: ent.Id}, ent.ExtraHelper); err != nil {
			self.logger.WithError(err).Debugf("failed to update extra filed")
		}
	}

	if ent, err = self.get_entity(ctx, *ent.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get entity")
		return nil, err
	}

	self.logger.WithFields(log.Fields{
		"entity_id": *ent.Id,
	}).Debugf("create entity")

	return ent, nil
}

func (self *StorageImpl) DeleteEntity(ctx context.Context, id string) error {
	var err error

	if err = self.GetDBConn(ctx).Delete(&Entity{}, "id = ?", id).Error; err != nil {
		self.logger.WithField("id", id).Debugf("failed to delete entity")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete entity")

	return nil
}

func (self *StorageImpl) PatchEntity(ctx context.Context, id string, entity *Entity) (*Entity, error) {
	var err error
	var ent *Entity
	var entNew Entity

	if entity.Alias != nil {
		entNew.Alias = entity.Alias
	}
	if entity.Password != nil {
		entNew.Password = entity.Password
	}

	if err = self.GetDBConn(ctx).Model(&Entity{Id: &id}).Update(entNew).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to patch entity")
		return nil, err
	}

	if entity.ExtraHelper != nil {
		if err = storage_helper.UpdateExtra(self.GetDBConn(ctx), &Entity{Id: &id}, entity.ExtraHelper); err != nil {
			self.logger.WithError(err).Debugf("failed to update extra filed")
		}
	}

	if ent, err = self.get_entity(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to get entity view")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("patch entity")

	return ent, nil
}

//todo remove password from return. zh
func (self *StorageImpl) GetEntity(ctx context.Context, id string) (*Entity, error) {
	var ent *Entity
	var err error

	if ent, err = self.get_entity(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to get entity")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("get entity")

	return ent, nil
}

func (self *StorageImpl) GetEntityByName(ctx context.Context, name string) (*Entity, error) {
	var ent *Entity
	var err error

	if ent, err = self.get_entity_by_name(ctx, name); err != nil {
		self.logger.WithError(err).Debugf("failed to get entity")
		return nil, err
	}

	self.logger.WithField("name", name).Debugf("get entity by name")

	return ent, nil
}

func (self *StorageImpl) ListEntities(ctx context.Context, ent *Entity) ([]*Entity, error) {
	var ents []*Entity
	var err error

	if ents, err = self.list_entities(ctx, ent); err != nil {
		self.logger.WithError(err).Debugf("failed to list entities")
		return nil, err
	}

	self.logger.Debugf("list entities")

	return ents, nil
}

func (self *StorageImpl) ListEntitiesByDomainId(ctx context.Context, id string) ([]*Entity, error) {
	var ent_dom_maps []*EntityDomainMapping
	var err error

	if err = self.GetDBConn(ctx).Find(&ent_dom_maps, "domain_id = ?", id).Error; err != nil {
		self.logger.WithField("domain_id", id).WithError(err).Debugf("failed to list entity and domain mapping")
		return nil, err
	}

	var ent *Entity
	var ents []*Entity
	for _, m := range ent_dom_maps {
		if ent, err = self.get_entity(ctx, *m.EntityId); err != nil {
			self.logger.WithField("entity_id", *m.EntityId).WithError(err).Debugf("failed to get entity")
			return nil, err
		}
		ents = append(ents, ent)
	}

	self.logger.WithField("domain_id", id).Debugf("list entities by domain id")

	return ents, nil
}

func (self *StorageImpl) AddRoleToEntity(ctx context.Context, entity_id, role_id string) error {
	var err error

	m := &EntityRoleMapping{
		EntityId: &entity_id,
		RoleId:   &role_id,
	}

	if err = self.GetDBConn(ctx).Create(m).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to add role to entity")
		return err
	}

	self.logger.WithFields(log.Fields{
		"entity_id": entity_id,
		"role_id":   role_id,
	}).Debugf("add role to entity")

	return nil
}

func (self *StorageImpl) RemoveRoleFromEntity(ctx context.Context, entity_id, role_id string) error {
	var err error

	if err = self.GetDBConn(ctx).Delete(&EntityRoleMapping{}, "entity_id = ? and role_id = ?", entity_id, role_id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to remove role from entity")
		return err
	}

	self.logger.WithFields(log.Fields{
		"entity_id": entity_id,
		"role_id":   role_id,
	}).Debugf("remove role from entity")

	return nil
}

func (self *StorageImpl) list_view_all_roles_by_entity_id(ctx context.Context, id string) ([]*Role, error) {
	var roles []*Role
	var grps []*Group
	var err error
	role_ids_set := map[string]bool{}

	if roles, err = self.list_view_roles_by_entity_id(ctx, id); err != nil {
		return nil, err
	}

	if grps, err = self.list_view_groups_by_entity_id(ctx, id); err != nil {
		return nil, err
	}

	var grp_role_maps []*GroupRoleMapping
	var grps_str []string
	for _, g := range grps {
		grps_str = append(grps_str, *g.Id)
	}
	if err = self.GetDBConn(ctx).Find(&grp_role_maps, "group_id in (?)", grps_str).Error; err != nil {
		return nil, err
	}

	for _, r := range roles {
		role_ids_set[*r.Id] = true
	}
	for _, m := range grp_role_maps {
		role_ids_set[*m.GroupId] = true
	}

	roles = nil
	for id, _ := range role_ids_set {
		roles = append(roles, &Role{Id: &id})
	}

	return roles, nil
}

func (self *StorageImpl) get_group(ctx context.Context, id string) (*Group, error) {
	var grp Group
	var err error

	if err = self.GetDBConn(ctx).First(&grp, "id = ?", id).Error; err != nil {
		return nil, err
	}

	grp.Domain = &Domain{Id: grp.DomainId}

	var sub_grp_maps []*SubjectGroupMapping
	if err = self.GetDBConn(ctx).Find(&sub_grp_maps, "group_id = ?", id).Error; err != nil {
		return nil, err
	}
	for _, m := range sub_grp_maps {
		grp.Subjects = append(grp.Subjects, &Entity{Id: m.SubjectId})
	}

	var obj_grp_maps []*ObjectGroupMapping
	if err = self.GetDBConn(ctx).Find(&obj_grp_maps, "group_id = ?", id).Error; err != nil {
		return nil, err
	}
	for _, m := range obj_grp_maps {
		grp.Objects = append(grp.Objects, &Entity{Id: m.ObjectId})
	}

	var grp_role_maps []*GroupRoleMapping
	if err = self.GetDBConn(ctx).Find(&grp_role_maps, "group_id = ?", id).Error; err != nil {
		return nil, err
	}

	// TODO(Peer): bad performance
	for _, m := range grp_role_maps {
		rol, err := self.get_role(ctx, *m.RoleId)
		if err != nil {
			return nil, err
		}
		grp.Roles = append(grp.Roles, rol)
	}

	return &grp, nil
}

func (self *StorageImpl) list_groups(ctx context.Context, grp *Group) ([]*Group, error) {
	var grps_t []*Group
	var err error

	g := &Group{}
	if grp.Id != nil {
		g.Id = grp.Id
	}
	if grp.DomainId != nil {
		g.DomainId = grp.DomainId
	}
	if grp.Name != nil {
		g.Name = grp.Name
	}
	if grp.Alias != nil {
		g.Alias = grp.Alias
	}

	if err = self.GetDBConn(ctx).Select("id").Find(&grps_t, g).Error; err != nil {
		return nil, err
	}

	var grps []*Group
	for _, g = range grps_t {
		if g, err = self.get_group(ctx, *g.Id); err != nil {
			return nil, err
		}
		grps = append(grps, g)
	}

	return grps, nil
}

func (self *StorageImpl) CreateGroup(ctx context.Context, grp *Group) (*Group, error) {
	var err error

	if err = self.GetDBConn(ctx).Create(grp).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to create group")
		return nil, err
	}

	if grp.ExtraHelper != nil {
		if err = storage_helper.UpdateExtra(self.GetDBConn(ctx), &Group{Id: grp.Id}, grp.ExtraHelper); err != nil {
			self.logger.WithError(err).Debugf("failed to update extra filed")
		}
	}

	if grp, err = self.get_group(ctx, *grp.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get group")
		return nil, err
	}

	self.logger.WithField("id", *grp.Id).Debugf("create group")

	return grp, nil
}

func (self *StorageImpl) DeleteGroup(ctx context.Context, id string) error {
	var err error

	if err = self.GetDBConn(ctx).Delete(&Group{}, "id = ?", id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to delete group")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete group")

	return nil
}

func (self *StorageImpl) PatchGroup(ctx context.Context, id string, group *Group) (*Group, error) {
	var err error
	var grp *Group
	var grpNew Group

	if group.Alias != nil {
		grpNew.Alias = group.Alias
	}
	if group.Description != nil {
		grpNew.Description = group.Description
	}

	if err = self.GetDBConn(ctx).Model(&Group{Id: &id}).Update(grpNew).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to patch group")
		return nil, err
	}

	if group.ExtraHelper != nil {
		if err = storage_helper.UpdateExtra(self.GetDBConn(ctx), &Group{Id: &id}, group.ExtraHelper); err != nil {
			self.logger.WithError(err).Debugf("failed to update extra filed")
		}
	}

	if grp, err = self.get_group(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to get group view")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("patch group")

	return grp, nil
}

func (self *StorageImpl) GetGroup(ctx context.Context, id string) (*Group, error) {
	var grp *Group
	var err error

	if grp, err = self.get_group(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to get group")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("get group")

	return grp, nil
}

func (self *StorageImpl) ExistGroup(ctx context.Context, id string) (bool, error) {
	var cnt int
	var err error

	if err = self.GetDBConn(ctx).Model(&Group{}).Where("id = ?", id).Count(&cnt).Error; err != nil {
		return false, err
	}

	return cnt > 0, nil
}

func (self *StorageImpl) ListGroups(ctx context.Context, grp *Group) ([]*Group, error) {
	var grps []*Group
	var err error

	if grps, err = self.list_groups(ctx, grp); err != nil {
		self.logger.WithError(err).Debugf("failed to list groups")
		return nil, err
	}

	self.logger.Debugf("list groups")

	return grps, nil
}

func (self *StorageImpl) AddRoleToGroup(ctx context.Context, group_id, role_id string) error {
	var err error

	m := &GroupRoleMapping{
		GroupId: &group_id,
		RoleId:  &role_id,
	}

	if err = self.GetDBConn(ctx).Create(m).Error; err != nil {
		self.logger.WithFields(log.Fields{
			"group_id": group_id,
			"role_id":  role_id,
		}).Debugf("failed to add role to group")
		return err
	}

	self.logger.WithFields(log.Fields{
		"group_id": group_id,
		"role_id":  role_id,
	}).Debugf("add role to group")

	return nil
}

func (self *StorageImpl) RemoveRoleFromGroup(ctx context.Context, group_id, role_id string) error {
	var err error

	if err = self.GetDBConn(ctx).Delete(&GroupRoleMapping{}, "group_id = ? and role_id = ?", group_id, role_id).Error; err != nil {
		self.logger.WithFields(log.Fields{
			"group_id": group_id,
			"role_id":  role_id,
		}).WithError(err).Debugf("failed to remove role from group")
		return err
	}

	self.logger.WithFields(log.Fields{
		"group_id": group_id,
		"role_id":  role_id,
	}).Debugf("remove role from group")

	return nil
}

func (self *StorageImpl) AddSubjectToGroup(ctx context.Context, group_id, subject_id string) error {
	var err error

	m := &SubjectGroupMapping{
		SubjectId: &subject_id,
		GroupId:   &group_id,
	}

	if err = self.GetDBConn(ctx).Create(m).Error; err != nil {
		self.logger.WithFields(log.Fields{
			"subject_id": subject_id,
			"group_id":   group_id,
		}).Debugf("failed to add subject to group")
	}

	self.logger.WithFields(log.Fields{
		"subject_id": subject_id,
		"group_id":   group_id,
	}).Debugf("add subject to group")

	return nil
}

func (self *StorageImpl) RemoveSubjectFromGroup(ctx context.Context, group_id, subject_id string) error {
	var err error

	if err = self.GetDBConn(ctx).Delete(&SubjectGroupMapping{}, "subject_id = ? and group_id = ?", subject_id, group_id).Error; err != nil {
		self.logger.WithFields(log.Fields{
			"subject_id": subject_id,
			"group_id":   group_id,
		}).Debugf("failed to remove subject from group")
	}

	self.logger.WithFields(log.Fields{
		"subject_id": subject_id,
		"group_id":   group_id,
	}).Debugf("remove subject from group")

	return nil
}

func (self *StorageImpl) SubjectExistsInGroup(ctx context.Context, subject_id, group_id string) (bool, error) {
	var cnt int
	var err error

	if err = self.GetDBConn(ctx).Model(&SubjectGroupMapping{}).Where("subject_id = ? and group_id = ?", subject_id, group_id).Count(&cnt).Error; err != nil {
		return false, err
	}

	return cnt > 0, nil
}

func (self *StorageImpl) AddObjectToGroup(ctx context.Context, group_id, object_id string) error {
	var err error

	m := &ObjectGroupMapping{
		ObjectId: &object_id,
		GroupId:  &group_id,
	}

	if err = self.GetDBConn(ctx).Create(m).Error; err != nil {
		self.logger.WithFields(log.Fields{
			"object_id": object_id,
			"group_id":  group_id,
		}).Debugf("failed to add object to group")
	}

	self.logger.WithFields(log.Fields{
		"object_id": object_id,
		"group_id":  group_id,
	}).Debugf("add object to group")

	return nil
}

func (self *StorageImpl) RemoveObjectFromGroup(ctx context.Context, group_id, object_id string) error {
	var err error

	if err = self.GetDBConn(ctx).Delete(&ObjectGroupMapping{}, "object_id = ? and group_id = ?", object_id, group_id).Error; err != nil {
		self.logger.WithFields(log.Fields{
			"object_id": object_id,
			"group_id":  group_id,
		}).Debugf("failed to remove object from group")
	}

	self.logger.WithFields(log.Fields{
		"object_id": object_id,
		"group_id":  group_id,
	}).Debugf("remove object from group")

	return nil
}

func (self *StorageImpl) ObjectExistsInGroup(ctx context.Context, object_id, group_id string) (bool, error) {
	var cnt int
	var err error

	if err = self.GetDBConn(ctx).Model(&ObjectGroupMapping{}).Where("object_id = ? and group_id = ?", object_id, group_id).Count(&cnt).Error; err != nil {
		return false, err
	}

	return cnt > 0, nil
}

func (self *StorageImpl) list_groups_by_group_ids(ctx context.Context, grp_ids []string) ([]*Group, error) {
	var err error
	var grps []*Group
	var grp *Group

	for _, grp_id := range grp_ids {
		if grp, err = self.get_group(ctx, grp_id); err != nil {
			return nil, err
		}
		grps = append(grps, grp)
	}

	return grps, nil
}

func (self *StorageImpl) list_groups_for_subject(ctx context.Context, subject_id string) ([]*Group, error) {
	var err error
	var group_ids []string
	var grps []*Group
	var sub_grp_maps []*SubjectGroupMapping

	if err = self.GetDBConn(ctx).Find(&sub_grp_maps, "subject_id = ?", subject_id).Error; err != nil {
		return nil, err
	}

	for _, m := range sub_grp_maps {
		group_ids = append(group_ids, *m.GroupId)
	}

	if grps, err = self.list_groups_by_group_ids(ctx, group_ids); err != nil {
		return nil, err
	}

	return grps, nil
}

func (self *StorageImpl) list_groups_for_object(ctx context.Context, object_id string) ([]*Group, error) {
	var err error
	var group_ids []string
	var grps []*Group
	var obj_grp_maps []*ObjectGroupMapping

	if err = self.GetDBConn(ctx).Find(&obj_grp_maps, "object_id = ?", object_id).Error; err != nil {
		return nil, err
	}

	for _, m := range obj_grp_maps {
		group_ids = append(group_ids, *m.GroupId)
	}

	if grps, err = self.list_groups_by_group_ids(ctx, group_ids); err != nil {
		return nil, err
	}

	return grps, nil
}

func (self *StorageImpl) ListGroupsForSubject(ctx context.Context, subject_id string) ([]*Group, error) {
	var err error
	var grps []*Group

	if grps, err = self.list_groups_for_subject(ctx, subject_id); err != nil {
		return nil, err
	}

	self.logger.Debugf("list groups for subject")

	return grps, nil
}

func (self *StorageImpl) ListGroupsForObject(ctx context.Context, object_id string) ([]*Group, error) {
	var err error
	var grps []*Group

	if grps, err = self.list_groups_for_object(ctx, object_id); err != nil {
		return nil, err
	}

	self.logger.Debugf("list groups for object")

	return grps, nil
}

func (self *StorageImpl) get_credential(ctx context.Context, id string) (*Credential, error) {
	var cred Credential
	var err error

	if err = self.GetDBConn(ctx).First(&cred, "id = ?", id).Error; err != nil {
		return nil, err
	}

	cred.Domain = &Domain{Id: cred.DomainId}
	cred.Entity = &Entity{Id: cred.EntityId}

	if cred.Roles, err = self.internal_list_view_credential_roles(ctx, &cred); err != nil {
		return nil, err
	}

	return &cred, nil
}

func (self *StorageImpl) list_view_credential_roles(ctx context.Context, id string) ([]*Role, error) {
	var cred Credential
	var err error

	if err = self.GetDBConn(ctx).Select("id, entity_id").First(&cred, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return self.internal_list_view_credential_roles(ctx, &cred)
}

func (self *StorageImpl) internal_list_view_credential_roles(ctx context.Context, cred *Credential) ([]*Role, error) {
	var ent_roles []*Role
	var roles []*Role
	var cred_role_maps []*CredentialRoleMapping
	var err error

	if err = self.GetDBConn(ctx).Find(&cred_role_maps, "credential_id = ?", *cred.Id).Error; err != nil {
		return nil, err
	}

	if ent_roles, err = self.list_view_all_roles_by_entity_id(ctx, *cred.EntityId); err != nil {
		return nil, err
	}

	if len(cred_role_maps) > 0 {
		ent_roles_set := map[string]bool{}
		for _, r := range ent_roles {
			ent_roles_set[*r.Id] = true
		}
		for _, m := range cred_role_maps {
			if _, ok := ent_roles_set[*m.RoleId]; ok {
				roles = append(roles, &Role{Id: m.RoleId})
			}
		}
	} else {
		roles = ent_roles
	}

	return roles, nil
}

func (self *StorageImpl) list_credentials(ctx context.Context, cred *Credential) ([]*Credential, error) {
	var creds_t []*Credential
	var creds []*Credential
	var err error

	c := &Credential{}
	if cred.Id != nil {
		c.Id = cred.Id
	}
	if cred.DomainId != nil {
		c.DomainId = cred.DomainId
	}
	if cred.EntityId != nil {
		c.EntityId = cred.EntityId
	}
	if cred.Name != nil {
		c.Name = cred.Name
	}
	if cred.Alias != nil {
		c.Alias = cred.Alias
	}

	if err = self.GetDBConn(ctx).Find(&creds_t, c).Error; err != nil {
		return nil, err
	}

	for _, c = range creds_t {
		if cred, err = self.get_credential(ctx, *c.Id); err != nil {
			return nil, err
		}

		creds = append(creds, cred)
	}

	return creds, nil
}

func (self *StorageImpl) CreateCredential(ctx context.Context, cred *Credential) (*Credential, error) {
	var err error

	tx := self.GetDBConn(ctx).Begin()
	tx.Create(cred)
	if len(cred.Roles) > 0 {
		var ms []*CredentialRoleMapping
		for _, r := range cred.Roles {
			ms = append(ms, &CredentialRoleMapping{
				CredentialId: cred.Id,
				RoleId:       r.Id,
			})
		}
		tx.Create(ms)
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		self.logger.WithError(err).Debugf("failed to create credential")
		return nil, err
	}

	if cred, err = self.get_credential(ctx, *cred.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get credential")
		return nil, err
	}

	self.logger.WithField("id", *cred.Id).Debugf("create credential")

	return cred, nil
}

func (self *StorageImpl) DeleteCredential(ctx context.Context, id string) error {
	var err error

	tx := self.GetDBConn(ctx).Begin()
	tx.Delete(&Credential{}, "id = ?", id)
	tx.Delete(&CredentialRoleMapping{}, "credential_id = ?", id)
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		self.logger.WithError(err).Debugf("failed to delete credential")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete credential")

	return nil
}

func (self *StorageImpl) PatchCredential(ctx context.Context, id string, credential *Credential) (*Credential, error) {
	var err error
	var cred *Credential
	var credNew Credential

	if credential.Alias != nil {
		credNew.Alias = credential.Alias
	}
	if credential.Description != nil {
		credNew.Description = credential.Description
	}

	if err = self.GetDBConn(ctx).Model(&Credential{Id: &id}).Update(credNew).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to patch credential")
		return nil, err
	}

	if cred, err = self.get_credential(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to get credential view")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("patch credential")

	return cred, nil
}

func (self *StorageImpl) GetCredential(ctx context.Context, id string) (*Credential, error) {
	var cred *Credential
	var err error

	if cred, err = self.get_credential(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to get credential")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("get credential")

	return cred, nil
}

func (self *StorageImpl) ListCredentials(ctx context.Context, cred *Credential) ([]*Credential, error) {
	var creds []*Credential
	var err error

	if creds, err = self.list_credentials(ctx, cred); err != nil {
		self.logger.WithError(err).Debugf("failed to list credentials")
		return nil, err
	}

	self.logger.Debugf("list credentials")

	return creds, nil
}

func (self *StorageImpl) get_token(ctx context.Context, id string) (*Token, error) {
	var tkn Token
	var tknp *Token
	var err error

	if err = self.GetDBConn(ctx).First(&tkn, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if tknp, err = self.internal_get_token(ctx, &tkn); err != nil {
		return nil, err
	}

	return tknp, nil
}

func (self *StorageImpl) get_token_by_text(ctx context.Context, text string) (*Token, error) {
	var tkn Token
	var tknp *Token
	var err error

	if err = self.GetDBConn(ctx).First(&tkn, "text = ?", text).Error; err != nil {
		return nil, err
	}

	if tknp, err = self.internal_get_token(ctx, &tkn); err != nil {
		return nil, err
	}

	return tknp, nil
}

func (self *StorageImpl) list_tokens(ctx context.Context, tkn *Token) ([]*Token, error) {
	var tkns_t []*Token
	var tkns []*Token
	var err error

	t := &Token{}
	if tkn.Id != nil {
		t.Id = tkn.Id
	}
	if tkn.DomainId != nil {
		t.DomainId = tkn.DomainId
	}
	if tkn.EntityId != nil {
		t.EntityId = tkn.EntityId
	}
	if tkn.CredentialId != nil {
		t.CredentialId = tkn.CredentialId
	}
	if tkn.Text != nil {
		t.Text = tkn.Text
	}

	if err = self.GetDBConn(ctx).Find(&tkns_t, t).Error; err != nil {
		return nil, err
	}

	for _, t = range tkns_t {
		if tkn, err = self.get_token(ctx, *t.Id); err != nil {
			return nil, err
		}
		tkns = append(tkns, tkn)
	}

	return tkns, nil
}

func (self *StorageImpl) internal_get_token(ctx context.Context, tkn *Token) (*Token, error) {
	var err error

	tkn.Domain = &Domain{Id: tkn.DomainId}
	tkn.Entity = &Entity{Id: tkn.EntityId}

	if tkn.CredentialId != nil {
		tkn.Credential = &Credential{Id: tkn.CredentialId}
		if tkn.Roles, err = self.list_view_credential_roles(ctx, *tkn.CredentialId); err != nil {
			return nil, err
		}
	} else {
		if tkn.Roles, err = self.list_view_roles_by_entity_id(ctx, *tkn.EntityId); err != nil {
			return nil, err
		}
	}

	if tkn.Groups, err = self.list_view_groups_by_entity_id(ctx, *tkn.EntityId); err != nil {
		return nil, err
	}

	return tkn, nil
}

func (self *StorageImpl) CreateToken(ctx context.Context, tkn *Token) (*Token, error) {
	var err error

	if err = self.GetDBConn(ctx).Create(tkn).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to create token")
		return nil, err
	}

	if tkn, err = self.get_token(ctx, *tkn.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get token by id")
		return nil, err
	}

	self.logger.WithField("id", *tkn.Id).Debugf("create token")

	return tkn, nil
}

func (self *StorageImpl) DeleteToken(ctx context.Context, id string) error {
	var err error

	if err = self.GetDBConn(ctx).Delete(&Token{}, "id = ?", id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to delete token")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete token")

	return nil
}

func (self *StorageImpl) RefreshToken(ctx context.Context, id string, expires_at time.Time) error {
	var err error

	if err = self.GetDBConn(ctx).Model(&Token{Id: &id}).Update(&Token{ExpiresAt: &expires_at}).Error; err != nil {
		return err
	}

	self.logger.WithFields(log.Fields{"id": id, "expires_at": expires_at}).Debugf("refresh token")

	return nil
}

func (self *StorageImpl) GetTokenByText(ctx context.Context, text string) (*Token, error) {
	var tkn *Token
	var err error

	if tkn, err = self.get_token_by_text(ctx, text); err != nil {
		self.logger.WithError(err).Debugf("failed to get token by text")
		return nil, err
	}

	self.logger.WithField("text", text).Debugf("get token by text")

	return tkn, nil
}

func (self *StorageImpl) GetToken(ctx context.Context, id string) (*Token, error) {
	var tkn *Token
	var err error

	if tkn, err = self.get_token(ctx, id); err != nil {
		self.logger.WithError(err).Debugf("failed to get token by id")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("get token by id")

	return tkn, nil
}

func (self *StorageImpl) ListTokens(ctx context.Context, tkn *Token) ([]*Token, error) {
	var tkns []*Token
	var err error

	if tkns, err = self.list_tokens(ctx, tkn); err != nil {
		self.logger.WithError(err).Debugf("failed to list tokens")
		return nil, err
	}

	self.logger.Debugf("list tokens")

	return tkns, nil
}

func (self *StorageImpl) Initialize(ctx context.Context) error {
	var err error
	var ok bool

	if ok, err = self.IsInitialized(ctx); err != nil {
		return err
	} else if ok {
		return ErrInitialized
	}

	val := ""
	cfg := &SystemConfig{
		Key:   &SYSTEM_CONFIG_INITIALIZE,
		Value: &val,
	}
	if err = self.GetDBConn(ctx).Create(cfg).Error; err != nil {
		return err
	}

	return nil
}

func (self *StorageImpl) IsInitialized(ctx context.Context) (bool, error) {
	var err error
	var cnt int

	if err = self.GetDBConn(ctx).Model(&SystemConfig{}).Where("key = ?", SYSTEM_CONFIG_INITIALIZE).Count(&cnt).Error; err != nil {
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
		&Domain{},
		&Action{},
		&Role{},
		&Entity{},
		&Group{},
		&Credential{},
		&Token{},
		&ActionRoleMapping{},
		&EntityRoleMapping{},
		&EntityDomainMapping{},
		&SubjectGroupMapping{},
		&ObjectGroupMapping{},
		&GroupRoleMapping{},
		&CredentialRoleMapping{},
		&SystemConfig{},
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

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
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
		return NewTracedStorage(s)
	} else {
		return s, nil
	}
}
