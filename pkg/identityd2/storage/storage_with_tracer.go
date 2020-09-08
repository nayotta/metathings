package metathings_identityd2_storage

import (
	"context"
	"time"

	opentracing_storage_helper "github.com/nayotta/metathings/pkg/common/opentracing/storage"
)

type TracedStorage struct {
	*opentracing_storage_helper.BaseTracedStorage
	*StorageImpl
}

func (s *TracedStorage) IsInitialized(ctx context.Context) (bool, error) {
	span, ctx := s.TraceWrapper(ctx, "IsInitialized")
	defer span.Finish()

	return s.StorageImpl.IsInitialized(ctx)
}

func (s *TracedStorage) Initialize(ctx context.Context) error {
	span, ctx := s.TraceWrapper(ctx, "Initialize")
	defer span.Finish()

	return s.StorageImpl.Initialize(ctx)
}

func (s *TracedStorage) CreateDomain(ctx context.Context, dom *Domain) (*Domain, error) {
	span, ctx := s.TraceWrapper(ctx, "CreateDomain")
	defer span.Finish()

	return s.StorageImpl.CreateDomain(ctx, dom)
}

func (s *TracedStorage) DeleteDomain(ctx context.Context, id string) error {
	span, ctx := s.TraceWrapper(ctx, "DeleteDomain")
	defer span.Finish()

	return s.StorageImpl.DeleteDomain(ctx, id)
}

func (s *TracedStorage) PatchDomain(ctx context.Context, id string, domain *Domain) (*Domain, error) {
	span, ctx := s.TraceWrapper(ctx, "PatchDomain")
	defer span.Finish()

	return s.StorageImpl.PatchDomain(ctx, id, domain)
}

func (s *TracedStorage) GetDomain(ctx context.Context, id string) (*Domain, error) {
	span, ctx := s.TraceWrapper(ctx, "GetDomain")
	defer span.Finish()

	return s.StorageImpl.GetDomain(ctx, id)
}

func (s *TracedStorage) ListDomains(ctx context.Context, dom *Domain) ([]*Domain, error) {
	span, ctx := s.TraceWrapper(ctx, "ListDomains")
	defer span.Finish()

	return s.StorageImpl.ListDomains(ctx, dom)
}

func (s *TracedStorage) AddEntityToDomain(ctx context.Context, domain_id, entity_id string) error {
	span, ctx := s.TraceWrapper(ctx, "AddEntityToDomain")
	defer span.Finish()

	return s.StorageImpl.AddEntityToDomain(ctx, domain_id, entity_id)
}

func (s *TracedStorage) RemoveEntityFromDomain(ctx context.Context, domain_id, entity_id string) error {
	span, ctx := s.TraceWrapper(ctx, "RemoveEntityFromDomain")
	defer span.Finish()

	return s.StorageImpl.RemoveEntityFromDomain(ctx, domain_id, entity_id)
}

func (s *TracedStorage) CreateAction(ctx context.Context, act *Action) (*Action, error) {
	span, ctx := s.TraceWrapper(ctx, "CreateAction")
	defer span.Finish()

	return s.StorageImpl.CreateAction(ctx, act)
}

func (s *TracedStorage) DeleteAction(ctx context.Context, id string) error {
	span, ctx := s.TraceWrapper(ctx, "DeleteAction")
	defer span.Finish()

	return s.StorageImpl.DeleteAction(ctx, id)
}

func (s *TracedStorage) PatchAction(ctx context.Context, id string, action *Action) (*Action, error) {
	span, ctx := s.TraceWrapper(ctx, "PatchAction")
	defer span.Finish()

	return s.StorageImpl.PatchAction(ctx, id, action)
}

func (s *TracedStorage) GetAction(ctx context.Context, id string) (*Action, error) {
	span, ctx := s.TraceWrapper(ctx, "GetAction")
	defer span.Finish()

	return s.StorageImpl.GetAction(ctx, id)
}

func (s *TracedStorage) ListActions(ctx context.Context, act *Action) ([]*Action, error) {
	span, ctx := s.TraceWrapper(ctx, "ListActions")
	defer span.Finish()

	return s.StorageImpl.ListActions(ctx, act)
}

func (s *TracedStorage) CreateRole(ctx context.Context, rol *Role) (*Role, error) {
	span, ctx := s.TraceWrapper(ctx, "CreateRole")
	defer span.Finish()

	return s.StorageImpl.CreateRole(ctx, rol)
}

func (s *TracedStorage) DeleteRole(ctx context.Context, id string) error {
	span, ctx := s.TraceWrapper(ctx, "DeleteRole")
	defer span.Finish()

	return s.StorageImpl.DeleteRole(ctx, id)
}

func (s *TracedStorage) PatchRole(ctx context.Context, id string, role *Role) (*Role, error) {
	span, ctx := s.TraceWrapper(ctx, "PatchRole")
	defer span.Finish()

	return s.StorageImpl.PatchRole(ctx, id, role)
}

func (s *TracedStorage) GetRole(ctx context.Context, id string) (*Role, error) {
	span, ctx := s.TraceWrapper(ctx, "GetRole")
	defer span.Finish()

	return s.StorageImpl.GetRole(ctx, id)
}

func (s *TracedStorage) GetRoleWithFullActions(ctx context.Context, id string) (*Role, error) {
	span, ctx := s.TraceWrapper(ctx, "GetRoleWithFullActions")
	defer span.Finish()

	return s.StorageImpl.GetRoleWithFullActions(ctx, id)
}

func (s *TracedStorage) ListRoles(ctx context.Context, rol *Role) ([]*Role, error) {
	span, ctx := s.TraceWrapper(ctx, "ListRoles")
	defer span.Finish()

	return s.StorageImpl.ListRoles(ctx, rol)
}

func (s *TracedStorage) AddActionToRole(ctx context.Context, role_id, action_id string) error {
	span, ctx := s.TraceWrapper(ctx, "AddActionToRole")
	defer span.Finish()

	return s.StorageImpl.AddActionToRole(ctx, role_id, action_id)
}

func (s *TracedStorage) RemoveActionFromRole(ctx context.Context, role_id, action_id string) error {
	span, ctx := s.TraceWrapper(ctx, "RemoveActionFromRole")
	defer span.Finish()

	return s.StorageImpl.RemoveActionFromRole(ctx, role_id, action_id)
}

func (s *TracedStorage) CreateEntity(ctx context.Context, ent *Entity) (*Entity, error) {
	span, ctx := s.TraceWrapper(ctx, "CreateEntity")
	defer span.Finish()

	return s.StorageImpl.CreateEntity(ctx, ent)
}

func (s *TracedStorage) DeleteEntity(ctx context.Context, id string) error {
	span, ctx := s.TraceWrapper(ctx, "DeleteEntity")
	defer span.Finish()

	return s.StorageImpl.DeleteEntity(ctx, id)
}

func (s *TracedStorage) PatchEntity(ctx context.Context, id string, entity *Entity) (*Entity, error) {
	span, ctx := s.TraceWrapper(ctx, "PatchEntity")
	defer span.Finish()

	return s.StorageImpl.PatchEntity(ctx, id, entity)
}

func (s *TracedStorage) GetEntity(ctx context.Context, id string) (*Entity, error) {
	span, ctx := s.TraceWrapper(ctx, "GetEntity")
	defer span.Finish()

	return s.StorageImpl.GetEntity(ctx, id)
}

func (s *TracedStorage) GetEntityByName(ctx context.Context, name string) (*Entity, error) {
	span, ctx := s.TraceWrapper(ctx, "GetEntityByName")
	defer span.Finish()

	return s.StorageImpl.GetEntityByName(ctx, name)
}

func (s *TracedStorage) ListEntities(ctx context.Context, ent *Entity) ([]*Entity, error) {
	span, ctx := s.TraceWrapper(ctx, "ListEntities")
	defer span.Finish()

	return s.StorageImpl.ListEntities(ctx, ent)
}

func (s *TracedStorage) ListEntitiesByDomainId(ctx context.Context, dom_id string) ([]*Entity, error) {
	span, ctx := s.TraceWrapper(ctx, "ListEntitiesByDomainId")
	defer span.Finish()

	return s.StorageImpl.ListEntitiesByDomainId(ctx, dom_id)
}

func (s *TracedStorage) AddRoleToEntity(ctx context.Context, entity_id, role_id string) error {
	span, ctx := s.TraceWrapper(ctx, "AddRoleToEntity")
	defer span.Finish()

	return s.StorageImpl.AddRoleToEntity(ctx, entity_id, role_id)
}

func (s *TracedStorage) RemoveRoleFromEntity(ctx context.Context, entity_id, role_id string) error {
	span, ctx := s.TraceWrapper(ctx, "RemoveRoleFromEntity")
	defer span.Finish()

	return s.StorageImpl.RemoveRoleFromEntity(ctx, entity_id, role_id)
}

func (s *TracedStorage) CreateGroup(ctx context.Context, grp *Group) (*Group, error) {
	span, ctx := s.TraceWrapper(ctx, "CreateGroup")
	defer span.Finish()

	return s.StorageImpl.CreateGroup(ctx, grp)
}

func (s *TracedStorage) DeleteGroup(ctx context.Context, id string) error {
	span, ctx := s.TraceWrapper(ctx, "DeleteGroup")
	defer span.Finish()

	return s.StorageImpl.DeleteGroup(ctx, id)
}

func (s *TracedStorage) PatchGroup(ctx context.Context, id string, group *Group) (*Group, error) {
	span, ctx := s.TraceWrapper(ctx, "PatchGroup")
	defer span.Finish()

	return s.StorageImpl.PatchGroup(ctx, id, group)
}

func (s *TracedStorage) GetGroup(ctx context.Context, id string) (*Group, error) {
	span, ctx := s.TraceWrapper(ctx, "GetGroup")
	defer span.Finish()

	return s.StorageImpl.GetGroup(ctx, id)
}

func (s *TracedStorage) ExistGroup(ctx context.Context, id string) (bool, error) {
	span, ctx := s.TraceWrapper(ctx, "ExistGroup")
	defer span.Finish()

	return s.StorageImpl.ExistGroup(ctx, id)
}

func (s *TracedStorage) ListGroups(ctx context.Context, grp *Group) ([]*Group, error) {
	span, ctx := s.TraceWrapper(ctx, "ListGroups")
	defer span.Finish()

	return s.StorageImpl.ListGroups(ctx, grp)
}

func (s *TracedStorage) AddRoleToGroup(ctx context.Context, group_id, role_id string) error {
	span, ctx := s.TraceWrapper(ctx, "AddRoleToGroup")
	defer span.Finish()

	return s.StorageImpl.AddRoleToGroup(ctx, group_id, role_id)
}

func (s *TracedStorage) RemoveRoleFromGroup(ctx context.Context, group_id, role_id string) error {
	span, ctx := s.TraceWrapper(ctx, "RemoveRoleFromGroup")
	defer span.Finish()

	return s.StorageImpl.RemoveRoleFromGroup(ctx, group_id, role_id)
}

func (s *TracedStorage) AddSubjectToGroup(ctx context.Context, group_id, subject_id string) error {
	span, ctx := s.TraceWrapper(ctx, "AddSubjectToGroup")
	defer span.Finish()

	return s.StorageImpl.AddSubjectToGroup(ctx, group_id, subject_id)
}

func (s *TracedStorage) RemoveSubjectFromGroup(ctx context.Context, group_id, subject_id string) error {
	span, ctx := s.TraceWrapper(ctx, "RemoveSubjectFromGroup")
	defer span.Finish()

	return s.StorageImpl.RemoveSubjectFromGroup(ctx, group_id, subject_id)
}

func (s *TracedStorage) SubjectExistsInGroup(ctx context.Context, subject_id, group_id string) (bool, error) {
	span, ctx := s.TraceWrapper(ctx, "SubjectExistsInGroup")
	defer span.Finish()

	return s.StorageImpl.SubjectExistsInGroup(ctx, subject_id, group_id)
}

func (s *TracedStorage) AddObjectToGroup(ctx context.Context, group_id, object_id string) error {
	span, ctx := s.TraceWrapper(ctx, "AddObjectToGroup")
	defer span.Finish()

	return s.StorageImpl.AddObjectToGroup(ctx, group_id, object_id)
}

func (s *TracedStorage) RemoveObjectFromGroup(ctx context.Context, group_id, object_id string) error {
	span, ctx := s.TraceWrapper(ctx, "RemoveObjectFromGroup")
	defer span.Finish()

	return s.StorageImpl.RemoveObjectFromGroup(ctx, group_id, object_id)
}

func (s *TracedStorage) ObjectExistsInGroup(ctx context.Context, object_id, group_id string) (bool, error) {
	span, ctx := s.TraceWrapper(ctx, "ObjectExistsInGroup")
	defer span.Finish()

	return s.StorageImpl.ObjectExistsInGroup(ctx, object_id, group_id)
}

func (s *TracedStorage) ListGroupsForSubject(ctx context.Context, subject_id string) ([]*Group, error) {
	span, ctx := s.TraceWrapper(ctx, "ListGroupsForSubject")
	defer span.Finish()

	return s.StorageImpl.ListGroupsForSubject(ctx, subject_id)
}

func (s *TracedStorage) ListGroupsForObject(ctx context.Context, object_id string) ([]*Group, error) {
	span, ctx := s.TraceWrapper(ctx, "ListGroupsForObject")
	defer span.Finish()

	return s.StorageImpl.ListGroupsForObject(ctx, object_id)
}

func (s *TracedStorage) CreateCredential(ctx context.Context, cred *Credential) (*Credential, error) {
	span, ctx := s.TraceWrapper(ctx, "CreateCredential")
	defer span.Finish()

	return s.StorageImpl.CreateCredential(ctx, cred)
}

func (s *TracedStorage) DeleteCredential(ctx context.Context, id string) error {
	span, ctx := s.TraceWrapper(ctx, "DeleteCredential")
	defer span.Finish()

	return s.StorageImpl.DeleteCredential(ctx, id)
}

func (s *TracedStorage) PatchCredential(ctx context.Context, id string, cred *Credential) (*Credential, error) {
	span, ctx := s.TraceWrapper(ctx, "PatchCredential")
	defer span.Finish()

	return s.StorageImpl.PatchCredential(ctx, id, cred)
}

func (s *TracedStorage) GetCredential(ctx context.Context, id string) (*Credential, error) {
	span, ctx := s.TraceWrapper(ctx, "GetCredential")
	defer span.Finish()

	return s.StorageImpl.GetCredential(ctx, id)
}

func (s *TracedStorage) ListCredentials(ctx context.Context, cred *Credential) ([]*Credential, error) {
	span, ctx := s.TraceWrapper(ctx, "ListCredentials")
	defer span.Finish()

	return s.StorageImpl.ListCredentials(ctx, cred)
}

func (s *TracedStorage) CreateToken(ctx context.Context, tkn *Token) (*Token, error) {
	span, ctx := s.TraceWrapper(ctx, "CreateToken")
	defer span.Finish()

	return s.StorageImpl.CreateToken(ctx, tkn)
}

func (s *TracedStorage) DeleteToken(ctx context.Context, id string) error {
	span, ctx := s.TraceWrapper(ctx, "DeleteToken")
	defer span.Finish()

	return s.StorageImpl.DeleteToken(ctx, id)
}

func (s *TracedStorage) RefreshToken(ctx context.Context, id string, expires_at time.Time) error {
	span, ctx := s.TraceWrapper(ctx, "RefreshToken")
	defer span.Finish()

	return s.StorageImpl.RefreshToken(ctx, id, expires_at)
}

func (s *TracedStorage) GetTokenByText(ctx context.Context, text string) (*Token, error) {
	span, ctx := s.TraceWrapper(ctx, "GetTokenByText")
	defer span.Finish()

	return s.StorageImpl.GetTokenByText(ctx, text)
}

func (s *TracedStorage) GetViewTokenByText(ctx context.Context, text string) (*Token, error) {
	span, ctx := s.TraceWrapper(ctx, "GetViewTokenByText")
	defer span.Finish()

	return s.StorageImpl.GetViewTokenByText(ctx, text)
}

func (s *TracedStorage) GetToken(ctx context.Context, id string) (*Token, error) {
	span, ctx := s.TraceWrapper(ctx, "GetToken")
	defer span.Finish()

	return s.StorageImpl.GetToken(ctx, id)
}

func (s *TracedStorage) ListTokens(ctx context.Context, tkn *Token) ([]*Token, error) {
	span, ctx := s.TraceWrapper(ctx, "ListTokens")
	defer span.Finish()

	return s.StorageImpl.ListTokens(ctx, tkn)
}

func NewTracedStorage(s *StorageImpl) (Storage, error) {
	return &TracedStorage{
		BaseTracedStorage: opentracing_storage_helper.NewBaseTracedStorage(s),
		StorageImpl:       s,
	}, nil
}
