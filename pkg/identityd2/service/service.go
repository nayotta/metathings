package metathings_identityd2_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

type MetathingsIdentitydServiceOption struct {
	StorageDriver string
	StorageUri    string
}

type MetathingsIdentitydService struct {
	opt     MetathingsIdentitydServiceOption
	logger  log.FieldLogger
	storage storage.Storage
}

func (self *MetathingsIdentitydService) DeleteDomain(context.Context, *pb.DeleteDomainRequest) (*empty.Empty, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) PatchDomain(context.Context, *pb.PatchDomainRequest) (*pb.PatchDomainResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) GetPolicy(context.Context, *pb.GetPolicyRequest) (*pb.GetPolicyResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) ListPolicies(context.Context, *pb.ListPoliciesRequest) (*pb.ListPoliciesResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) ListPoliciesForRole(context.Context, *pb.ListPoliciesForRoleRequest) (*pb.ListPoliciesForRoleResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) DeleteRole(context.Context, *pb.DeleteRoleRequest) (*empty.Empty, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) PatchRole(context.Context, *pb.PatchRoleRequest) (*pb.PatchRoleResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) GetRole(context.Context, *pb.GetRoleRequest) (*pb.GetRoleResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) ListRoles(context.Context, *pb.ListRolesRequest) (*pb.ListRolesResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) ListRolesForEntity(context.Context, *pb.ListRolesForEntityRequest) (*pb.ListRolesForEntityResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) DeleteEntity(context.Context, *pb.DeleteEntityRequest) (*empty.Empty, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) PatchEntity(context.Context, *pb.PatchEntityRequest) (*pb.PatchEntityResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) GetEntity(context.Context, *pb.GetEntityRequest) (*pb.GetEntityResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) ListEntities(context.Context, *pb.ListEntitiesRequest) (*pb.ListEntitiesResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) ShowEntity(context.Context, *empty.Empty) (*pb.ShowEntityResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) DeleteGroup(context.Context, *pb.DeleteGroupRequest) (*empty.Empty, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) PatchGroup(context.Context, *pb.PatchGroupRequest) (*pb.PatchGroupResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) GetGroup(context.Context, *pb.GetGroupRequest) (*pb.GetGroupResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) ListGroups(context.Context, *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) ListGroupsForEntity(context.Context, *pb.ListGroupsForEntityRequest) (*pb.ListGroupsForEntityResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) ShowGroups(context.Context, *empty.Empty) (*pb.ShowGroupsResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) RemoveRoleFromGroup(context.Context, *pb.RemoveRoleFromGroupRequest) (*empty.Empty, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) AddEntityToGroup(context.Context, *pb.AddEntityToGroupRequest) (*empty.Empty, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) RemoveEntityFromGroup(context.Context, *pb.RemoveEntityFromGroupRequest) (*empty.Empty, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) CreateCredential(context.Context, *pb.CreateCredentialRequest) (*pb.CreateCredentialResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) DeleteCredential(context.Context, *pb.DeleteCredentialRequest) (*empty.Empty, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) PatchCredential(context.Context, *pb.PatchCredentialRequest) (*pb.PatchCredentialResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) GetCredential(context.Context, *pb.GetCredentialRequest) (*pb.GetCredentialResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) ListCredentials(context.Context, *pb.ListCredentialsRequest) (*pb.ListCredentialsResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) ListCredentialsForEntity(context.Context, *pb.ListCredentialsForEntityRequest) (*pb.ListCredentialsForEntityResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) IssueTokenByCredential(context.Context, *pb.IssueTokenByCredentialRequest) (*pb.IssueTokenByCredentialResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) IssueTokenByPassword(context.Context, *pb.IssueTokenByPasswordRequest) (*pb.IssueTokenByPasswordResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) RevokeToken(context.Context, *pb.RevokeTokenRequest) (*empty.Empty, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) ValidateToken(context.Context, *pb.ValidateTokenRequest) (*empty.Empty, error) {
	panic("unimplemented")
}

func (self *MetathingsIdentitydService) GetTokenByText(context.Context, *pb.GetTokenByTextRequest) (*pb.GetTokenByTextResponse, error) {
	panic("unimplemented")
}
