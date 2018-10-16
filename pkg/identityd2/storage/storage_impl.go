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

func (self *StorageImpl) CreateDomain(*Domain) (*Domain, error) {
	panic("unimplemented")
}
func (self *StorageImpl) DeleteDomain(id string) error {
	panic("unimplemented")
}
func (self *StorageImpl) PatchDomain(id string, domain *Domain) (*Domain, error) {
	panic("unimplemented")
}
func (self *StorageImpl) GetDomain(id string) (*Domain, error) {
	panic("unimplemented")
}
func (self *StorageImpl) ListDomains(*Domain) ([]*Domain, error) {
	panic("unimplemented")
}
func (self *StorageImpl) AddEntityToDomain(domain_id, entity_id string) error {
	panic("unimplemented")
}
func (self *StorageImpl) RemoveEntityFromDomain(domain_id, entity_id string) error {
	panic("unimplemented")
}

func (self *StorageImpl) CreateRole(*Role) (*Role, error) {
	panic("unimplemented")
}
func (self *StorageImpl) DeleteRole(id string) error {
	panic("unimplemented")
}
func (self *StorageImpl) PatchRole(id string, role *Role) (*Role, error) {
	panic("unimplemented")
}
func (self *StorageImpl) GetRole(id string) (*Role, error) {
	panic("unimplemented")
}
func (self *StorageImpl) ListRoles(*Role) ([]*Role, error) {
	panic("unimplemented")
}
func (self *StorageImpl) GetPolicy(id string) (*Policy, error) {
	panic("unimplemented")
}
func (self *StorageImpl) CreatePolicy(*Policy) (*Policy, error) {
	panic("unimplemented")
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
