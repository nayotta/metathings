package metathings_identityd2_storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
)

type StorageImpl struct {
	db     *gorm.DB
	logger log.FieldLogger
}

func (self *StorageImpl) get_domain(id string) (*Domain, error) {
	var err error
	var dom Domain

	if err = self.db.First(&dom, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if dom.ParentId != nil && *dom.ParentId != "" {
		dom.Parent = &Domain{
			Id: dom.ParentId,
		}
	}

	children := []*Domain{}
	if err = self.db.Select("id").Where("parent_id = ?", id).Find(&children).Error; err != nil {
		return nil, err
	}

	return &dom, nil
}

func (self *StorageImpl) list_domains(dom *Domain) ([]*Domain, error) {
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

	if err = self.db.Select("id").Find(&doms_t, d).Error; err != nil {
		return nil, err
	}

	doms := []*Domain{}
	for _, d = range doms_t {
		if d, err = self.get_domain(*d.Id); err != nil {
			return nil, err
		}
		doms = append(doms, d)
	}

	return doms, nil
}

func (self *StorageImpl) CreateDomain(dom *Domain) (*Domain, error) {
	var err error

	if err = self.db.Create(dom).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to create domain")
		return nil, err
	}

	if dom, err = self.get_domain(*dom.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get domain")
		return nil, err
	}

	self.logger.WithField("id", *dom.Id).Debugf("create domain")

	return dom, nil
}

func (self *StorageImpl) DeleteDomain(id string) error {
	if err := self.db.Delete(&Domain{}, "id = ?", id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to delete domain")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete domain")

	return nil
}

func (self *StorageImpl) PatchDomain(id string, domain *Domain) (*Domain, error) {
	panic("unimplemented")
}

func (self *StorageImpl) GetDomain(id string) (*Domain, error) {
	var err error
	var dom *Domain

	if dom, err = self.get_domain(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get domain")
		return nil, err
	}

	self.logger.WithField("id", *dom.Id).Debugf("get domain")

	return dom, nil
}

func (self *StorageImpl) ListDomains(dom *Domain) ([]*Domain, error) {
	var doms []*Domain
	var err error

	if doms, err = self.list_domains(dom); err != nil {
		self.logger.WithError(err).Debugf("failed to list domains")
		return nil, err
	}

	self.logger.Debugf("list domains")

	return doms, nil
}

func (self *StorageImpl) AddEntityToDomain(domain_id, entity_id string) error {
	m := &EntityDomainMapping{
		DomainId: &domain_id,
		EntityId: &entity_id,
	}

	if err := self.db.Create(m).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to add entity to domain")
		return err
	}

	self.logger.WithFields(log.Fields{
		"entity_id": entity_id,
		"domain_id": domain_id,
	}).Debugf("add entity to domain")

	return nil
}

func (self *StorageImpl) RemoveEntityFromDomain(domain_id, entity_id string) error {
	if err := self.db.Delete(&EntityDomainMapping{}, "domain_id = ? and entity_id = ?", domain_id, entity_id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to remove entity from domain")
	}

	self.logger.WithFields(log.Fields{
		"entity_id": entity_id,
		"domain_id": domain_id,
	}).Debugf("remove entity from domain")

	return nil
}

func (self *StorageImpl) get_role(id string) (*Role, error) {
	var role *Role
	var policies []*Policy
	var err error

	if err = self.db.First(&role, "id = ?", id).Error; err != nil {
		return nil, err
	}

	role.Domain = &Domain{
		Id: role.DomainId,
	}

	if err = self.db.Find(&policies, "role_id = ?", id).Error; err != nil {
		return nil, err
	}

	role.Policies = policies

	return role, nil
}

func (self *StorageImpl) list_roles(role *Role) ([]*Role, error) {
	var err error
	var roles_t []*Role

	r := &Role{}
	if role.Id != nil {
		r.Id = role.Id
	}
	if role.DomainId != nil {
		r.DomainId = role.DomainId
	}
	if role.Name != nil {
		r.Name = role.Name
	}
	if role.Alias != nil {
		r.Alias = role.Alias
	}

	if err = self.db.Select("id").Find(&roles_t, r).Error; err != nil {
		return nil, err
	}

	var roles []*Role
	for _, r = range roles_t {
		if r, err = self.get_role(*r.Id); err != nil {
			return nil, err
		}
		roles = append(roles, r)
	}

	return roles, nil
}

func (self *StorageImpl) CreateRole(role *Role) (*Role, error) {
	var err error

	if err = self.db.Create(role).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to create role")
		return nil, err
	}

	if role, err = self.get_role(*role.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get role")
		return nil, err
	}

	self.logger.WithFields(log.Fields{}).Debugf("create role")

	return role, nil
}

func (self *StorageImpl) DeleteRole(id string) error {
	var err error

	tx := self.db.Begin()
	tx.Delete(&Policy{}, "role_id = ?", id)
	tx.Delete(&Role{}, "id = ?", id)
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		self.logger.WithError(err).Debugf("failed to delete role")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete role")

	return nil
}

func (self *StorageImpl) PatchRole(id string, role *Role) (*Role, error) {
	panic("unimplemented")
}

func (self *StorageImpl) GetRole(id string) (*Role, error) {
	var role *Role
	var err error

	if role, err = self.get_role(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get role")
		return nil, err
	}

	return role, nil
}

func (self *StorageImpl) ListRoles(role *Role) ([]*Role, error) {
	var roles []*Role
	var err error

	if roles, err = self.list_roles(role); err != nil {
		self.logger.WithError(err).Debugf("failed to list roles")
		return nil, err
	}

	self.logger.Debugf("list roles")

	return roles, nil
}

func (self *StorageImpl) get_policy(id string) (*Policy, error) {
	var plc *Policy
	var err error

	if err = self.db.First(&plc, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return plc, nil
}

func (self *StorageImpl) GetPolicy(id string) (*Policy, error) {
	var plc *Policy
	var err error

	if plc, err = self.get_policy(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get policy")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("get policy")

	return plc, nil
}

func (self *StorageImpl) CreatePolicy(plc *Policy) (*Policy, error) {
	var err error

	if err = self.db.Create(plc).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to create policy")
		return nil, err
	}

	if plc, err = self.get_policy(*plc.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get policy")
		return nil, err
	}

	self.logger.WithFields(log.Fields{
		"id":      *plc.Id,
		"role_id": *plc.RoleId,
		"rule":    *plc.Rule,
	}).Debugf("create policy")

	return plc, err
}

func (self *StorageImpl) DeletePolicy(id string) error {
	var err error

	if err = self.db.Delete(&Policy{}, "id = ?", id).Error; err != nil {
		self.logger.WithField("id", id).Debugf("failed to delete policy")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete policy")

	return nil
}

func (self *StorageImpl) ListPoliciesForEntity(id string) ([]*Policy, error) {
	panic("unimplemented")
}

func (self *StorageImpl) list_domains_by_entity_id(id string) ([]*Domain, error) {
	var err error

	var ent_dom_maps []*EntityDomainMapping
	if err = self.db.Find(&ent_dom_maps, "entity_id = ?", id).Error; err != nil {
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

func (self *StorageImpl) list_groups_by_entity_id(id string) ([]*Group, error) {
	var err error

	var ent_grp_maps []*EntityGroupMapping
	if err = self.db.Find(&ent_grp_maps, "entity_id = ?", id).Error; err != nil {
		return nil, err
	}

	var grps []*Group
	for _, m := range ent_grp_maps {
		grps = append(grps, &Group{Id: m.GroupId})
	}

	return grps, nil
}

func (self *StorageImpl) list_roles_by_entity_id(id string) ([]*Role, error) {
	var err error

	var ent_role_maps []*EntityRoleMapping
	if err = self.db.Find(&ent_role_maps, "entity_id = ?", id).Error; err != nil {
		return nil, err
	}

	var roles []*Role
	for _, m := range ent_role_maps {
		roles = append(roles, &Role{Id: m.RoleId})
	}

	return roles, nil
}

func (self *StorageImpl) get_entity(id string) (*Entity, error) {
	var ent *Entity
	var err error

	if err = self.db.First(&ent, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if ent.Domains, err = self.list_domains_by_entity_id(id); err != nil {
		return nil, err
	}

	if ent.Groups, err = self.list_groups_by_entity_id(id); err != nil {
		return nil, err
	}

	if ent.Roles, err = self.list_roles_by_entity_id(id); err != nil {
		return nil, err
	}

	return ent, nil
}

func (self *StorageImpl) list_entities(ent *Entity) ([]*Entity, error) {
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

	if err = self.db.Select("id").Find(&ents_t, e).Error; err != nil {
		return nil, err
	}

	var ents []*Entity
	for _, e := range ents_t {
		if ent, err = self.get_entity(*e.Id); err != nil {
			return nil, err
		}
		ents = append(ents, ent)
	}

	return ents, nil
}

func (self *StorageImpl) CreateEntity(ent *Entity) (*Entity, error) {
	var err error

	if err = self.db.Create(ent).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to create entity")
		return nil, err
	}

	if ent, err = self.get_entity(*ent.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get entity")
		return nil, err
	}

	self.logger.WithFields(log.Fields{
		"entity_id": *ent.Id,
	}).Debugf("create entity")

	return ent, nil
}

func (self *StorageImpl) DeleteEntity(id string) error {
	var err error

	if err = self.db.Delete(&Entity{}, "id = ?", id).Error; err != nil {
		self.logger.WithField("id", id).Debugf("failed to delete entity")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete entity")

	return nil
}

func (self *StorageImpl) PatchEntity(id string, entity *Entity) (*Entity, error) {
	panic("unimplemented")
}

func (self *StorageImpl) GetEntity(id string) (*Entity, error) {
	var ent *Entity
	var err error

	if ent, err = self.get_entity(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get entity")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("get entity")

	return ent, nil
}

func (self *StorageImpl) GetEntityByName(name string) (*Entity, error) {
	panic("unimplemented")
}

func (self *StorageImpl) ListEntities(ent *Entity) ([]*Entity, error) {
	var ents []*Entity
	var err error

	if ents, err = self.list_entities(ent); err != nil {
		self.logger.WithError(err).Debugf("failed to list entities")
		return nil, err
	}

	self.logger.Debugf("list entities")

	return ents, nil
}

func (self *StorageImpl) ListEntitiesByDomainId(id string) ([]*Entity, error) {
	var ent_dom_maps []*EntityDomainMapping
	var err error

	if err = self.db.Find(&ent_dom_maps, "domain_id = ?", id).Error; err != nil {
		self.logger.WithField("domain_id", id).WithError(err).Debugf("failed to list entity and domain mapping")
		return nil, err
	}

	var ent *Entity
	var ents []*Entity
	for _, m := range ent_dom_maps {
		if ent, err = self.get_entity(*m.EntityId); err != nil {
			self.logger.WithField("entity_id", *m.EntityId).WithError(err).Debugf("failed to get entity")
			return nil, err
		}
		ents = append(ents, ent)
	}

	self.logger.WithField("domain_id", id).Debugf("list entities by domain id")

	return ents, nil
}

func (self *StorageImpl) AddRoleToEntity(entity_id, role_id string) error {
	var err error

	m := &EntityRoleMapping{
		EntityId: &entity_id,
		RoleId:   &role_id,
	}

	if err = self.db.Create(m).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to add role to entity")
		return err
	}

	self.logger.WithFields(log.Fields{
		"entity_id": entity_id,
		"role_id":   role_id,
	}).Debugf("add role to entity")

	return nil
}

func (self *StorageImpl) RemoveRoleFromEntity(entity_id, role_id string) error {
	var err error

	if err = self.db.Delete(&EntityRoleMapping{}, "entity_id = ? and role_id = ?", entity_id, role_id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to remove role from entity")
		return err
	}

	self.logger.WithFields(log.Fields{
		"entity_id": entity_id,
		"role_id":   role_id,
	}).Debugf("remove role from entity")

	return nil
}

func (self *StorageImpl) get_group(id string) (*Group, error) {
	var grp *Group
	var err error

	if err = self.db.First(&grp, "id = ?", id).Error; err != nil {
		return nil, err
	}

	grp.Domain = &Domain{Id: grp.DomainId}

	var ent_grp_maps []*EntityGroupMapping
	if err = self.db.Find(&ent_grp_maps, "group_id = ?", id).Error; err != nil {
		return nil, err
	}
	var entities []*Entity
	for _, m := range ent_grp_maps {
		entities = append(entities, &Entity{Id: m.EntityId})
	}
	grp.Entities = entities

	var grp_role_maps []*GroupRoleMapping
	if err = self.db.Find(&grp_role_maps, "group_id = ?", id).Error; err != nil {
		return nil, err
	}
	var roles []*Role
	for _, m := range grp_role_maps {
		roles = append(roles, &Role{Id: m.RoleId})
	}
	grp.Roles = roles

	return grp, nil
}

func (self *StorageImpl) list_groups(grp *Group) ([]*Group, error) {
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

	if err = self.db.Select("id").Find(&grps_t, g).Error; err != nil {
		return nil, err
	}

	var grps []*Group
	for _, g = range grps_t {
		if g, err = self.get_group(*g.Id); err != nil {
			return nil, err
		}
		grps = append(grps, g)
	}

	return grps, nil
}

func (self *StorageImpl) CreateGroup(grp *Group) (*Group, error) {
	var err error

	if err = self.db.Create(grp).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to create group")
		return nil, err
	}

	if grp, err = self.get_group(*grp.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get group")
		return nil, err
	}

	self.logger.WithField("id", *grp.Id).Debugf("create group")

	return grp, nil
}

func (self *StorageImpl) DeleteGroup(id string) error {
	var err error

	if err = self.db.Delete(&Group{}, "id = ?", id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to delete group")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete group")

	return nil
}

func (self *StorageImpl) PatchGroup(id string, group *Group) (*Group, error) {
	panic("unimplemented")
}

func (self *StorageImpl) GetGroup(id string) (*Group, error) {
	var grp *Group
	var err error

	if grp, err = self.get_group(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get group")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("get group")

	return grp, nil
}

func (self *StorageImpl) ListGroups(grp *Group) ([]*Group, error) {
	var grps []*Group
	var err error

	if grps, err = self.list_groups(grp); err != nil {
		self.logger.WithError(err).Debugf("failed to list groups")
		return nil, err
	}

	self.logger.Debugf("list groups")

	return grps, nil
}

func (self *StorageImpl) AddRoleToGroup(group_id, role_id string) error {
	var err error

	m := &GroupRoleMapping{
		GroupId: &group_id,
		RoleId:  &role_id,
	}

	if err = self.db.Create(m).Error; err != nil {
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

func (self *StorageImpl) RemoveRoleFromGroup(group_id, role_id string) error {
	var err error

	if err = self.db.Delete(&GroupRoleMapping{}, "group_id = ? and role_id = ?", group_id, role_id).Error; err != nil {
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

func (self *StorageImpl) AddEntityToGroup(entity_id, group_id string) error {
	var err error

	m := &EntityGroupMapping{
		EntityId: &entity_id,
		GroupId:  &group_id,
	}

	if err = self.db.Create(m).Error; err != nil {
		self.logger.WithFields(log.Fields{
			"entity_id": entity_id,
			"group_id":  group_id,
		}).Debugf("failed to add entity to group")
	}

	self.logger.WithFields(log.Fields{
		"group_id":  group_id,
		"entity_id": entity_id,
	}).Debugf("add entity to group")

	return nil
}

func (self *StorageImpl) RemoveEntityFromGroup(entity_id, group_id string) error {
	var err error

	if err = self.db.Delete(&EntityGroupMapping{}, "entity_id = ? and group_id = ?", entity_id, group_id).Error; err != nil {
		self.logger.WithFields(log.Fields{
			"entity_id": entity_id,
			"group_id":  group_id,
		}).Debugf("failed to remove entity from group")
	}

	self.logger.WithFields(log.Fields{
		"entity_id": entity_id,
		"group_id":  group_id,
	}).Debugf("remove entity from group")

	return nil
}

func (self *StorageImpl) CreateCredential(*Credential) (*Credential, error) {
	panic("unimplemented")
}

func (self *StorageImpl) DeleteCredential(id string) error {
	panic("unimplemented")
}

func (self *StorageImpl) PatchCredential(id string, credential *Credential) (*Credential, error) {
	panic("unimplemented")
}

func (self *StorageImpl) GetCredential(id string) (*Credential, error) {
	panic("unimplemented")
}

func (self *StorageImpl) ListCredentials(*Credential) ([]*Credential, error) {
	panic("unimplemented")
}

func (self *StorageImpl) CreateToken(*Token) (*Token, error) {
	panic("unimplemented")
}

func (self *StorageImpl) DeleteToken(id string) error {
	panic("unimplemented")
}

func (self *StorageImpl) GetTokenByText(text string) (*Token, error) {
	panic("unimplemented")
}

func (self *StorageImpl) GetToken(id string) (*Token, error) {
	panic("unimplemented")
}

func (self *StorageImpl) ListTokens(*Token) ([]*Token, error) {
	panic("unimplemented")
}

func init_args(s *StorageImpl, args ...interface{}) error {
	var key string
	var ok bool

	if len(args)%2 != 0 {
		return BadArgument
	}

	for i := 0; i < len(args); i += 2 {
		key, ok = args[i].(string)
		if !ok {
			return BadArgument
		}

		switch key {
		case "logger":
			s.logger, ok = args[i+1].(log.FieldLogger)
			if !ok {
				return BadArgument
			}
		}
	}

	return nil
}

func new_db(s *StorageImpl, driver, uri string) error {
	var db *gorm.DB
	var err error

	if db, err = gorm.Open(driver, uri); err != nil {
		return err
	}
	s.db = db

	return nil
}

func init_db(s *StorageImpl) error {
	s.db.AutoMigrate(
		&Domain{},
		&Policy{},
		&Role{},
		&Entity{},
		&Group{},
		&Credential{},
		&Token{},

		&EntityRoleMapping{},
		&EntityDomainMapping{},
		&EntityGroupMapping{},
		&GroupRoleMapping{},
		&CredentialRoleMapping{},
		&TokenRoleMapping{},
	)

	return nil
}

func NewStorageImpl(driver, uri string, args ...interface{}) (*StorageImpl, error) {
	var err error

	s := &StorageImpl{}
	if err = init_args(s, args...); err != nil {
		return nil, err
	}
	if err = new_db(s, driver, uri); err != nil {
		return nil, err
	}
	if err = init_db(s); err != nil {
		return nil, err
	}

	return s, nil
}
