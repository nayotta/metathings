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

	roles := []*Role{}
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
	panic("unimplemented")
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
	panic("unimplemented")
}

func (self *StorageImpl) ListPoliciesForEntity(id string) ([]*Policy, error) {
	panic("unimplemented")
}

func (self *StorageImpl) CreateEntity(*Entity) (*Entity, error) {
	panic("unimplemented")
}

func (self *StorageImpl) DeleteEntity(id string) error {
	panic("unimplemented")
}

func (self *StorageImpl) PatchEntity(id string, entity *Entity) (*Entity, error) {
	panic("unimplemented")
}

func (self *StorageImpl) GetEntity(id string) (*Entity, error) {
	panic("unimplemented")
}

func (self *StorageImpl) GetEntityByName(name string) (*Entity, error) {
	panic("unimplemented")
}

func (self *StorageImpl) ListEntities(*Entity) ([]*Entity, error) {
	panic("unimplemented")
}

func (self *StorageImpl) ListEntitiesByDomainId(dom_id string) ([]*Entity, error) {
	panic("unimplemented")
}

func (self *StorageImpl) AddRoleToEntity(entity_id, role_id string) error {
	panic("unimplemented")
}

func (self *StorageImpl) RemoveRoleFromEntity(entity_id, role_id string) error {
	panic("unimplemented")
}

func (self *StorageImpl) CreateGroup(*Group) (*Group, error) {
	panic("unimplemented")
}

func (self *StorageImpl) DeleteGroup(id string) error {
	panic("unimplemented")
}

func (self *StorageImpl) PatchGroup(id string, group *Group) (*Group, error) {
	panic("unimplemented")
}

func (self *StorageImpl) GetGroup(id string) (*Group, error) {
	panic("unimplemented")
}

func (self *StorageImpl) ListGroups(*Group) ([]*Group, error) {
	panic("unimplemented")
}

func (self *StorageImpl) AddRoleToGroup(group_id, role_id string) error {
	panic("unimplemented")
}

func (self *StorageImpl) RemoveRoleFromGroup(group_id, role_id string) error {
	panic("unimplemented")
}

func (self *StorageImpl) AddEntityToGroup(entity_id, group_id string) error {
	panic("unimplemented")
}

func (self *StorageImpl) RemoveEntityFromGroup(entity_id, group_id string) error {
	panic("unimplemented")
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
