package metathings_identityd2_storage

import "time"

type Domain struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	Name     *string `gorm:"column:name"`
	Alias    *string `gorm:"column:alias"`
	ParentId *string `gorm:"column:domain_id"`
	Extra    *string `gorm:"extra"`

	Parent   *Domain   `gorm:"-"`
	Children []*Domain `gorm:"-"`
}

type Policy struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	RoleId      *string `gorm:"column:role_id"`
	Rule        *string `gorm:"column:rule"`
	Description *string `gorm:"column:description"`
}

type Role struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	DomainId    *string `gorm:"column:domain_id"`
	Name        *string `gorm:"column:name"`
	Alias       *string `gorm:"column:alias"`
	Description *string `gorm:"column:description"`
	Extra       *string `gorm:"column:extra"`

	Domain   *Domain   `gorm:"-"`
	Policies []*Policy `gorm:"-"`
}

type Entity struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	DomainId    *string `gorm:"column:domain_id"`
	Name        *string `gorm:"column:name"`
	Alias       *string `gorm:"column:alias"`
	Password    *string `gorm:"column:password"`
	Description *string `gorm:"column:description"`
	Extra       *string `gorm:"column:extra"`

	Domain *Domain  `gorm:"-"`
	Groups []*Group `gorm:"-"`
	Roles  []*Role  `gorm:"-"`
}

type EntityRoleMapping struct {
	CreatedAt time.Time

	EntityId *string `gorm:"entity_id"`
	RoleId   *string `gorm:"role_id"`
}

type Group struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	DomainId    *string `gorm:"column:domain_id"`
	Name        *string `gorm:"column:name"`
	Alias       *string `gorm:"column:alias"`
	Description *string `gorm:"column:description"`
	Extra       *string `gorm:"column:extra"`

	Domain   *Domain   `gorm:"-"`
	Entities []*Entity `gorm:"-"`
	Roles    []*Role   `gorm:"-"`
}

type EntityGroupMapping struct {
	CreatedAt time.Time

	EntityId *string `gorm:"entity_id"`
	RoleId   *string `gorm:"role_id"`
}

type GroupRoleMapping struct {
	CreatedAt time.Time

	GroupId *string `gorm:"group_id"`
	RoleId  *string `gorm:"role_id"`
}

type Credential struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	DomainId    *string    `gorm:"column:domain_id"`
	EntityId    *string    `gorm:"column:entity_id"`
	Name        *string    `gorm:"column:name"`
	Alias       *string    `gorm:"column:alias"`
	Secret      *string    `gorm:"column:secret"`
	Description *string    `gorm:"column:description"`
	ExpiresAt   *time.Time `gorm:"column:expires_at"`

	Domain *Domain `gorm:"-"`
	Entity *Entity `gorm:"-"`
	Roles  []*Role `gorm:"-"`
}

type CredentialRoleMapping struct {
	CreatedAt time.Time

	CredentialId *string `gorm:"credential_id"`
	RoleId       *string `gorm:"role_id"`
}

type Token struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	DomainId     *string    `gorm:"column:domain_id"`
	EntityId     *string    `gorm:"column:entity_id"`
	CredentialId *string    `gorm:"column:credential_id"`
	IssuedAt     *time.Time `gorm:"column:issued_at"`
	ExpiresAt    *time.Time `gorm:"column:expires_at"`
	Text         *string    `gorm:"column:text"`

	Domain     *Domain     `gorm:"-"`
	Entity     *Entity     `gorm:"-"`
	Credential *Credential `gorm:"-"`
	Roles      []*Role     `gorm:"-"`
}

type TokenRoleMapping struct {
	CreatedAt time.Time

	TokenId *string `gorm:"column:token_id"`
	RoleId  *string `gorm:"column:role_id"`
}

type Storage interface {
	CreateDomain(*Domain) (*Domain, error)
	DeleteDomain(id string) error
	PatchDomain(id string, domain *Domain) (*Domain, error)
	GetDomain(id string) (*Domain, error)
	ListDomains(*Domain) ([]*Domain, error)

	CreateRole(*Role) (*Role, error)
	DeleteRole(id string) error
	PatchRole(id string, role *Role) (*Role, error)
	GetRole(id string) (*Role, error)
	ListRoles(*Role) ([]*Role, error)
	GetPolicy(id string) (*Policy, error)
	CreatePolicy(*Policy) (*Policy, error)
	DeletePolicy(id string) error
	ListPoliciesForEntity(id string) ([]*Policy, error)

	CreateEntity(*Entity) (*Entity, error)
	DeleteEntity(id string) error
	PatchEntity(id string, entity *Entity) (*Entity, error)
	GetEntity(id string) (*Entity, error)
	ListEntities(*Entity) ([]*Entity, error)
	AddRoleToEntity(entity_id, role_id string) error
	RemoveRoleFromEntity(entity_id, role_id string) error

	CreateGroup(*Group) (*Group, error)
	DeleteGroup(id string) error
	PatchGroup(id string, group *Group) (*Group, error)
	GetGroup(id string) (*Group, error)
	ListGroups(*Group) ([]*Group, error)
	AddRoleToGroup(group_id, role_id string) error
	RemoveRoleFromGroup(group_id, role_id string) error
	AddEntityToGroup(entity_id, group_id string) error
	RemoveEntityFromGroup(entity_id, group_id string) error

	CreateCredential(*Credential) (*Credential, error)
	DeleteCredential(id string) error
	PatchCredential(id string, credential *Credential) (*Credential, error)
	GetCredential(id string) (*Credential, error)
	ListCredentials(*Credential) ([]*Credential, error)

	CreateToken(*Token) (*Token, error)
	DeleteToken(id string) error
	GetTokenByText(text string) (*Token, error)
	GetToken(id string) (*Token, error)
}

func NewStorage(driver, uri string, args ...interface{}) (Storage, error) {
	panic("unimplemented")
}
