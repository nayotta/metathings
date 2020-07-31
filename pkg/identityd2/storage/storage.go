package metathings_identityd2_storage

import (
	"context"
	"time"
)

type Domain struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	Name     *string `gorm:"column:name"`
	Alias    *string `gorm:"column:alias"`
	ParentId *string `gorm:"column:parent_id"`
	Extra    *string `gorm:"extra"`

	Parent      *Domain           `gorm:"-"`
	Children    []*Domain         `gorm:"-"`
	ExtraHelper map[string]string `gorm:"-"`
}

type Action struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	Name        *string `gorm:"column:name"`
	Alias       *string `gorm:"column:alias"`
	Description *string `gorm:"column:description"`
	Extra       *string `gorm:"column:extra"`

	ExtraHelper map[string]string `gorm:"-"`
}

type Role struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	Name        *string `gorm:"column:name"`
	Alias       *string `gorm:"column:alias"`
	Description *string `gorm:"column:description"`
	Extra       *string `gorm:"column:extra"`

	Actions     []*Action         `gorm:"-"`
	ExtraHelper map[string]string `gorm:"-"`
}

type ActionRoleMapping struct {
	CreatedAt time.Time

	ActionId *string `gorm:"action_id"`
	RoleId   *string `gorm:"role_id"`
}

type Entity struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	Name     *string `gorm:"column:name"`
	Alias    *string `gorm:"column:alias"`
	Password *string `gorm:"column:password"`
	Extra    *string `gorm:"column:extra"`

	Domains     []*Domain         `gorm:"-"`
	Groups      []*Group          `gorm:"-"`
	Roles       []*Role           `gorm:"-"`
	ExtraHelper map[string]string `gorm:"-"`
}

type EntityRoleMapping struct {
	CreatedAt time.Time

	EntityId *string `gorm:"entity_id"`
	RoleId   *string `gorm:"role_id"`
}

type EntityDomainMapping struct {
	CreatedAt time.Time

	EntityId *string `gorm:"entity_id"`
	DomainId *string `gorm:"domain_id"`
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

	Domain      *Domain           `gorm:"-"`
	Subjects    []*Entity         `gorm:"-"`
	Objects     []*Entity         `gorm:"-"`
	Roles       []*Role           `gorm:"-"`
	ExtraHelper map[string]string `gorm:"-"`
}

type SubjectGroupMapping struct {
	CreatedAt time.Time

	SubjectId *string `gorm:"subject_id"`
	GroupId   *string `gorm:"group_id"`
}

type ObjectGroupMapping struct {
	CreatedAt time.Time

	ObjectId *string `gorm:"object_id"`
	GroupId  *string `gorm:"group_id"`
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

	DomainId      *string    `gorm:"column:domain_id"`
	EntityId      *string    `gorm:"column:entity_id"`
	CredentialId  *string    `gorm:"column:credential_id"`
	IssuedAt      *time.Time `gorm:"column:issued_at"`
	ExpiresAt     *time.Time `gorm:"column:expires_at"`
	ExpiresPeriod *int64     `gorm:"column:expires_period;default:0"`
	Text          *string    `gorm:"column:text"`

	Domain     *Domain     `gorm:"-"`
	Entity     *Entity     `gorm:"-"`
	Credential *Credential `gorm:"-"`
	Roles      []*Role     `gorm:"-"`
	Groups     []*Group    `gorm:"-"`
}

type SystemConfig struct {
	CreatedAt time.Time
	UpdatedAt time.Time

	Key   *string `gorm:"column:key"`
	Value *string `gorm:"column:value"`
}

type Storage interface {
	IsInitialized(context.Context) (bool, error)
	Initialize(context.Context) error

	CreateDomain(context.Context, *Domain) (*Domain, error)
	DeleteDomain(ctx context.Context, id string) error
	PatchDomain(ctx context.Context, id string, domain *Domain) (*Domain, error)
	GetDomain(ctx context.Context, id string) (*Domain, error)
	ListDomains(context.Context, *Domain) ([]*Domain, error)
	AddEntityToDomain(ctx context.Context, domain_id, entity_id string) error
	RemoveEntityFromDomain(ctx context.Context, domain_id, entity_id string) error

	CreateAction(context.Context, *Action) (*Action, error)
	DeleteAction(ctx context.Context, id string) error
	PatchAction(ctx context.Context, id string, action *Action) (*Action, error)
	GetAction(ctx context.Context, id string) (*Action, error)
	ListActions(context.Context, *Action) ([]*Action, error)

	CreateRole(context.Context, *Role) (*Role, error)
	DeleteRole(ctx context.Context, id string) error
	PatchRole(ctx context.Context, id string, role *Role) (*Role, error)
	GetRole(ctx context.Context, id string) (*Role, error)
	GetRoleWithFullActions(ctx context.Context, id string) (*Role, error)
	ListRoles(context.Context, *Role) ([]*Role, error)
	AddActionToRole(ctx context.Context, role_id, action_id string) error
	RemoveActionFromRole(ctx context.Context, role_id, action_id string) error

	CreateEntity(context.Context, *Entity) (*Entity, error)
	DeleteEntity(ctx context.Context, id string) error
	PatchEntity(ctx context.Context, id string, entity *Entity) (*Entity, error)
	GetEntity(ctx context.Context, id string) (*Entity, error)
	GetEntityByName(ctx context.Context, name string) (*Entity, error)
	ListEntities(context.Context, *Entity) ([]*Entity, error)
	ListEntitiesByDomainId(ctx context.Context, dom_id string) ([]*Entity, error)
	AddRoleToEntity(ctx context.Context, entity_id, role_id string) error
	RemoveRoleFromEntity(ctx context.Context, entity_id, role_id string) error

	CreateGroup(context.Context, *Group) (*Group, error)
	DeleteGroup(ctx context.Context, id string) error
	PatchGroup(ctx context.Context, id string, group *Group) (*Group, error)
	GetGroup(ctx context.Context, id string) (*Group, error)
	ExistGroup(ctx context.Context, id string) (bool, error)
	ListGroups(context.Context, *Group) ([]*Group, error)
	AddRoleToGroup(ctx context.Context, group_id, role_id string) error
	RemoveRoleFromGroup(ctx context.Context, group_id, role_id string) error
	AddSubjectToGroup(ctx context.Context, group_id, subject_id string) error
	RemoveSubjectFromGroup(ctx context.Context, group_id, subject_id string) error
	SubjectExistsInGroup(ctx context.Context, subject_id, group_id string) (bool, error)
	AddObjectToGroup(ctx context.Context, group_id, object_id string) error
	RemoveObjectFromGroup(ctx context.Context, group_id, object_id string) error
	ObjectExistsInGroup(ctx context.Context, object_id, group_id string) (bool, error)
	ListGroupsForSubject(ctx context.Context, subject_id string) ([]*Group, error)
	ListGroupsForObject(ctx context.Context, object_id string) ([]*Group, error)

	CreateCredential(context.Context, *Credential) (*Credential, error)
	DeleteCredential(ctx context.Context, id string) error
	PatchCredential(ctx context.Context, id string, credential *Credential) (*Credential, error)
	GetCredential(ctx context.Context, id string) (*Credential, error)
	ListCredentials(context.Context, *Credential) ([]*Credential, error)

	CreateToken(context.Context, *Token) (*Token, error)
	DeleteToken(ctx context.Context, id string) error
	RefreshToken(ctx context.Context, id string, expires_at time.Time) error
	GetTokenByText(ctx context.Context, text string) (*Token, error)
	GetToken(ctx context.Context, id string) (*Token, error)
	ListTokens(context.Context, *Token) ([]*Token, error)
}

func NewStorage(driver, uri string, args ...interface{}) (Storage, error) {
	return NewStorageImpl(driver, uri, args...)
}
