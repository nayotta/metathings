package metathings_identityd2_storage

import (
	"time"
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
)

type StorageImpl struct {
	db     *gorm.DB
	logger log.FieldLogger
}

func (self *StorageImpl) list_view_children_domains_by_domain_id(id string) ([]*Domain, error) {
	var doms []*Domain
	var err error

	if err = self.db.Select("id").Find(&doms, "parent_id = ?", id).Error; err != nil {
		return nil, err
	}

	return doms, nil
}

func (self *StorageImpl) get_domain(id string) (*Domain, error) {
	var err error
	var dom *Domain

	if err = self.db.First(dom, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if dom.ParentId != nil && *dom.ParentId != "" {
		dom.Parent = &Domain{
			Id: dom.ParentId,
		}
	}

	if dom.Children, err = self.list_view_children_domains_by_domain_id(id); err != nil {
		return nil, err
	}

	return dom, nil
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
	var err error
	var dom *Domain

	tx := self.db.Begin()

	if err = self.db.First(&dom, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if domain.Name != nil && dom.Name != domain.Name {
		tx.Model(&dom).Update("Name", domain.Name)
	}
	if domain.Alias != nil && dom.Alias != domain.Alias {
		tx.Model(&dom).Update("Alias", domain.Alias)
	}
	if domain.ParentId != nil && dom.ParentId != domain.ParentId {
		tx.Model(&dom).Update("ParentId", domain.ParentId)
	}
	if domain.Extra != nil && dom.Extra != domain.Extra {
		tx.Model(&dom).Update("Extra", domain.Extra)
	}

	tx.Model(&dom).Update("UpdatedAt", time.Now())

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		self.logger.WithError(err).Debugf("failed to patch domain")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("patch domain")

	return dom, nil
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
	var err error

	if err = self.db.First(&role, "id = ?", id).Error; err != nil {
		return nil, err
	}

	role.Domain = &Domain{
		Id: role.DomainId,
	}

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

	if err = self.db.Delete(&Role{}, "id = ?", id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to delete role")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete role")

	return nil
}

func (self *StorageImpl) PatchRole(id string, role *Role) (*Role, error) {
	var err error
	var rol *Role

	tx := self.db.Begin()

	if err = self.db.First(&rol, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if role.DomainId != nil && rol.DomainId != role.DomainId {
		tx.Model(&rol).Update("DomainId", role.DomainId)
	}
	if role.Name != nil && rol.Name != role.Name {
		tx.Model(&rol).Update("Name", role.Name)
	}
	if role.Alias != nil && rol.Alias != role.Alias {
		tx.Model(&rol).Update("Alia", role.Alias)
	}
	if role.Description != nil && rol.Description != role.Description {
		tx.Model(&rol).Update("Description", role.Description)
	}
	if role.Extra != nil && rol.Extra != role.Extra {
		tx.Model(&rol).Update("Extra", role.Extra)
	}

	tx.Model(&rol).Update("UpdatedAt", time.Now())

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		self.logger.WithError(err).Debugf("failed to patch role")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("patch role")

	return rol, nil
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

func (self *StorageImpl) list_view_domains_by_entity_id(id string) ([]*Domain, error) {
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

func (self *StorageImpl) list_view_groups_by_entity_id(id string) ([]*Group, error) {
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

func (self *StorageImpl) list_view_roles_by_entity_id(id string) ([]*Role, error) {
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

	if ent.Domains, err = self.list_view_domains_by_entity_id(id); err != nil {
		return nil, err
	}

	if ent.Groups, err = self.list_view_groups_by_entity_id(id); err != nil {
		return nil, err
	}

	if ent.Roles, err = self.list_view_roles_by_entity_id(id); err != nil {
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
	var err error
	var ent *Entity

	tx := self.db.Begin()

	if err = self.db.First(&ent, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if entity.Name != nil && ent.Name != entity.Name {
		tx.Model(&ent).Update("Name", entity.Name)
	}
	if entity.Alias != nil && ent.Alias != entity.Name {
		tx.Model(&ent).Update("Alias", entity.Alias)
	}
	if entity.Password != nil && ent.Password != entity.Password {
		tx.Model(&ent).Update("Password", entity.Password)
	}
	if entity.Extra != nil && ent.Extra != entity.Extra {
		tx.Model(&ent).Update("Extra", entity.Extra)
	}

	tx.Model(&ent).Update("UpdatedAt", time.Now())

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		self.logger.WithError(err).Debugf("failed to patch entity")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("patch entity")

	return ent, nil
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

func (self *StorageImpl) list_view_all_roles_by_entity_id(id string) ([]*Role, error) {
	var roles []*Role
	var grps []*Group
	var err error
	role_ids_set := map[string]bool{}

	if roles, err = self.list_view_roles_by_entity_id(id); err != nil {
		return nil, err
	}

	if grps, err = self.list_view_groups_by_entity_id(id); err != nil {
		return nil, err
	}

	var grp_role_maps []*GroupRoleMapping
	var grps_str []string
	for _, g := range grps {
		grps_str = append(grps_str, *g.Id)
	}
	if err = self.db.Find(&grp_role_maps, "group_id in (?)", grps_str).Error; err != nil {
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
	var err error
	var grp *Group

	tx := self.db.Begin()

	if err = self.db.First(&grp, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if group.DomainId != nil && grp.DomainId != group.DomainId {
		tx.Model(&grp).Update("DomainId", group.DomainId)
	}
	if group.Alias != nil && grp.Alias != group.Alias {
		tx.Model(&grp).Update("Alias", group.Alias)
	}
	if group.Description != nil && grp.Description != group.Description {
		tx.Model(&grp).Update("Description", group.Description)
	}
	if group.Extra != nil && grp.Extra != group.Extra {
		tx.Model(&grp).Update("Extra", group.Extra)
	}

	tx.Model(&grp).Update("UpdatedAt", time.Now())

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		self.logger.WithError(err).Debugf("failed to patch group")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("patch group")

	return grp, nil
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

func (self *StorageImpl) get_credential(id string) (*Credential, error) {
	var cred *Credential
	var err error

	if err = self.db.First(&cred, "id = ?", id).Error; err != nil {
		return nil, err
	}

	cred.Domain = &Domain{Id: cred.DomainId}
	cred.Entity = &Entity{Id: cred.EntityId}

	if cred.Roles, err = self.internal_list_view_credential_roles(cred); err != nil {
		return nil, err
	}

	return cred, nil
}

func (self *StorageImpl) list_view_credential_roles(id string) ([]*Role, error) {
	var cred *Credential
	var err error

	if err = self.db.Select("id", "entity_id").First(&cred, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return self.internal_list_view_credential_roles(cred)
}

func (self *StorageImpl) internal_list_view_credential_roles(cred *Credential) ([]*Role, error) {
	var ent_roles []*Role
	var roles []*Role
	var cred_role_maps []*CredentialRoleMapping
	var err error

	if err = self.db.Find(&cred_role_maps, "credential_id = ?", *cred.Id).Error; err != nil {
		return nil, err
	}

	if ent_roles, err = self.list_view_all_roles_by_entity_id(*cred.EntityId); err != nil {
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

func (self *StorageImpl) list_credentials(cred *Credential) ([]*Credential, error) {
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
		c.DomainId = cred.DomainId
	}
	if cred.Name != nil {
		c.Name = cred.Name
	}
	if cred.Alias != nil {
		c.Alias = cred.Alias
	}

	if err = self.db.Find(&creds_t, c).Error; err != nil {
		return nil, err
	}

	for _, c = range creds_t {
		if cred, err = self.get_credential(*c.Id); err != nil {
			return nil, err
		}

		creds = append(creds, cred)
	}

	return creds, nil
}

func (self *StorageImpl) CreateCredential(cred *Credential) (*Credential, error) {
	var err error

	tx := self.db.Begin()
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

	if cred, err = self.get_credential(*cred.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get credential")
		return nil, err
	}

	self.logger.WithField("id", *cred.Id).Debugf("create credential")

	return cred, nil
}

func (self *StorageImpl) DeleteCredential(id string) error {
	var err error

	tx := self.db.Begin()
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

func (self *StorageImpl) PatchCredential(id string, credential *Credential) (*Credential, error) {
	var err error
	var cred *Credential

	tx := self.db.Begin()

	if err = self.db.First(&cred, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if credential.DomainId != nil && credential.DomainId != cred.DomainId {
		tx.Model(&cred).Update("DomainId", credential.DomainId)
	}
	if credential.EntityId != nil && credential.EntityId != cred.EntityId {
		tx.Model(&cred).Update("EntityId", credential.EntityId)
	}
	if credential.Name != nil && credential.Name != cred.Name {
		tx.Model(&cred).Update("Name", credential.Name)
	}
	if credential.Alias != nil && credential.Alias != cred.Alias {
		tx.Model(&cred).Update("Alias", credential.Alias)
	}
	if credential.Secret != nil && credential.Secret != cred.Secret {
		tx.Model(&cred).Update("Secret", credential.Secret)
	}
	if credential.Description != nil && credential.Description != cred.Description {
		tx.Model(&cred).Update("Description", credential.Description)
	}
	if credential.ExpiresAt != nil && credential.ExpiresAt != cred.ExpiresAt {
		tx.Model(&cred).Update("ExpiresAt", credential.ExpiresAt)
	}

	tx.Model(&cred).Update("UpdatedAt", time.Now())

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		self.logger.WithError(err).Debugf("failed to patch credential")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("patch credential")

	return cred, nil
}

func (self *StorageImpl) GetCredential(id string) (*Credential, error) {
	var cred *Credential
	var err error

	if cred, err = self.get_credential(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get credential")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("get credential")

	return cred, nil
}

func (self *StorageImpl) ListCredentials(cred *Credential) ([]*Credential, error) {
	var creds []*Credential
	var err error

	if creds, err = self.list_credentials(cred); err != nil {
		self.logger.WithError(err).Debugf("failed to list credentials")
		return nil, err
	}

	self.logger.Debugf("list credentials")

	return creds, nil
}

func (self *StorageImpl) get_token(id string) (*Token, error) {
	var tkn *Token
	var err error

	if err = self.db.First(&tkn, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if tkn, err = self.internal_get_token(tkn); err != nil {
		return nil, err
	}

	return tkn, nil
}

func (self *StorageImpl) get_token_by_text(text string) (*Token, error) {
	var tkn *Token
	var err error

	if err = self.db.First(&tkn, "text = ?", text).Error; err != nil {
		return nil, err
	}

	if tkn, err = self.internal_get_token(tkn); err != nil {
		return nil, err
	}

	return tkn, nil
}

func (self *StorageImpl) list_tokens(tkn *Token) ([]*Token, error) {
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

	if err = self.db.Find(&tkns_t, t).Error; err != nil {
		return nil, err
	}

	for _, t = range tkns_t {
		if tkn, err = self.get_token(*t.Id); err != nil {
			return nil, err
		}
		tkns = append(tkns, tkn)
	}

	return tkns, nil
}

func (self *StorageImpl) internal_get_token(tkn *Token) (*Token, error) {
	var err error

	tkn.Domain = &Domain{Id: tkn.DomainId}
	tkn.Entity = &Entity{Id: tkn.EntityId}
	tkn.Credential = &Credential{Id: tkn.CredentialId}
	if tkn.Roles, err = self.list_view_credential_roles(*tkn.CredentialId); err != nil {
		return nil, err
	}
	if tkn.Groups, err = self.list_view_groups_by_entity_id(*tkn.EntityId); err != nil {
		return nil, err
	}

	return tkn, nil
}

func (self *StorageImpl) CreateToken(tkn *Token) (*Token, error) {
	var err error

	if err = self.db.Create(tkn).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to create token")
		return nil, err
	}

	if tkn, err = self.get_token(*tkn.Id); err != nil {
		self.logger.WithError(err).Debugf("failed to get token by id")
		return nil, err
	}

	self.logger.WithField("id", *tkn.Id).Debugf("create token")

	return tkn, nil
}

func (self *StorageImpl) DeleteToken(id string) error {
	var err error

	if err = self.db.Delete(&Token{}, "id = ?", id).Error; err != nil {
		self.logger.WithError(err).Debugf("failed to delete token")
		return err
	}

	self.logger.WithField("id", id).Debugf("delete token")

	return nil
}

func (self *StorageImpl) GetTokenByText(text string) (*Token, error) {
	var tkn *Token
	var err error

	if tkn, err = self.get_token_by_text(text); err != nil {
		self.logger.WithError(err).Debugf("failed to get token by text")
		return nil, err
	}

	self.logger.WithField("text", text).Debugf("get token by text")

	return tkn, nil
}

func (self *StorageImpl) GetToken(id string) (*Token, error) {
	var tkn *Token
	var err error

	if tkn, err = self.get_token(id); err != nil {
		self.logger.WithError(err).Debugf("failed to get token by id")
		return nil, err
	}

	self.logger.WithField("id", id).Debugf("get token by id")

	return tkn, nil
}

func (self *StorageImpl) ListTokens(tkn *Token) ([]*Token, error) {
	var tkns []*Token
	var err error

	if tkns, err = self.list_tokens(tkn); err != nil {
		self.logger.WithError(err).Debugf("failed to list tokens")
		return nil, err
	}

	self.logger.Debugf("list tokens")

	return tkns, nil
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
