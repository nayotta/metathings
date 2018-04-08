package metathings_identity_service

import (
	"context"
	"fmt"

	google_protobuf3 "github.com/golang/protobuf/ptypes/empty"
	"github.com/parnurzeal/gorequest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/bigdatagz/metathings/identity/service/encode_decode"
	pb "github.com/bigdatagz/metathings/proto/identity"
)

type options struct {
	keystoneAdminAddr  string
	keystonePublicAddr string
}

var defaultServiceOptions = options{}

type ServiceOptions func(*options)

func SetKeystoneAdminAddress(addr string) func(*options) {
	return func(o *options) {
		o.keystoneAdminAddr = addr
	}
}

func SetKeystonePublicAddress(addr string) func(*options) {
	return func(o *options) {
		o.keystonePublicAddr = addr
	}
}

type metathingsIdentityService struct {
	opts options
}

// https://developer.openstack.org/api-ref/identity/v3/#create-region
func (srv *metathingsIdentityService) CreateRegion(context.Context, *pb.CreateRegionRequest) (*pb.CreateRegionResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/#delete-region
func (srv *metathingsIdentityService) DeleteRegion(context.Context, *pb.DeleteRegionRequest) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/#update-region
func (srv *metathingsIdentityService) PatchRegion(context.Context, *pb.PatchRegionRequest) (*pb.PatchRegionResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/#show-region-details
func (srv *metathingsIdentityService) GetRegion(context.Context, *pb.GetRegionRequest) (*pb.GetRegionResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/#list-regions
func (srv *metathingsIdentityService) ListRegions(context.Context, *pb.ListRegionsRequest) (*pb.ListRegionsResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#create-domain
func (srv *metathingsIdentityService) CreateDomain(context.Context, *pb.CreateDomainRequest) (*pb.CreateDomainResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#delete-domain
func (srv *metathingsIdentityService) DeleteDomain(context.Context, *pb.DeleteDomainRequest) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#update-domain
func (srv *metathingsIdentityService) PatchDomain(context.Context, *pb.PatchDomainRequest) (*pb.PatchDomainResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#show-domain-details
func (srv *metathingsIdentityService) GetDomain(context.Context, *pb.GetDomainRequest) (*pb.GetDomainResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-domains
func (srv *metathingsIdentityService) ListDomains(context.Context, *pb.ListDomainsRequest) (*pb.ListDomainsResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#create-project
func (srv *metathingsIdentityService) CreateProject(context.Context, *pb.CreateProjectRequest) (*pb.CreateProjectResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#delete-project
func (srv *metathingsIdentityService) DeleteProject(context.Context, *pb.DeleteProjectRequest) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#update-project
func (srv *metathingsIdentityService) PatchProject(context.Context, *pb.PatchProjectRequest) (*pb.PatchProjectResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#show-project-details
func (srv *metathingsIdentityService) GetProject(context.Context, *pb.GetProjectRequest) (*pb.GetProjectResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-projects
func (srv *metathingsIdentityService) ListProjects(context.Context, *pb.ListProjectsRequest) (*pb.ListProjectsResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-projects-for-user
func (srv *metathingsIdentityService) ListProjectsForUser(context.Context, *pb.ListProjectsForUserRequest) (*pb.ListProjectsForUserResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#create-user
func (srv *metathingsIdentityService) CreateUser(context.Context, *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#delete-user
func (srv *metathingsIdentityService) DeleteUser(context.Context, *pb.DeleteUserRequest) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#update-user
func (srv *metathingsIdentityService) PatchUser(context.Context, *pb.PatchUserRequest) (*pb.PatchUserResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#show-user-details
func (srv *metathingsIdentityService) GetUser(context.Context, *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-users
func (srv *metathingsIdentityService) ListUsers(context.Context, *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#change-password-for-user
func (srv *metathingsIdentityService) ChangePassword(context.Context, *pb.ChangePasswordRequest) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-users-in-group
func (srv *metathingsIdentityService) ListUsersInGroup(context.Context, *pb.ListUsersInGroupRequest) (*pb.ListUsersInGroupResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#create-group
func (srv *metathingsIdentityService) CreateGroup(context.Context, *pb.CreateGroupRequest) (*pb.CreateGroupResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#delete-group
func (srv *metathingsIdentityService) DeleteGroup(context.Context, *pb.DeleteGroupRequest) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#update-group
func (srv *metathingsIdentityService) PatchGroup(context.Context, *pb.PatchGroupRequest) (*pb.PatchGroupResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#show-group-details
func (srv *metathingsIdentityService) GetGroup(context.Context, *pb.GetGroupRequest) (*pb.GetGroupResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-groups
func (srv *metathingsIdentityService) ListGroups(context.Context, *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#add-user-to-group
func (srv *metathingsIdentityService) AddUserToGroup(context.Context, *pb.AddUserToGroupRequest) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#remove-user-from-group
func (srv *metathingsIdentityService) RemoveUserFromGroup(context.Context, *pb.RemoveUserFromGroupRequest) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#check-whether-user-belongs-to-group
func (srv *metathingsIdentityService) CheckUserInGroup(context.Context, *pb.CheckUserInGroupRequest) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-groups-to-which-a-user-belongs
func (srv *metathingsIdentityService) ListGroupsForUser(context.Context, *pb.ListGroupsForUserRequest) (*pb.ListGroupsForUserResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#create-role
func (srv *metathingsIdentityService) CreateRole(context.Context, *pb.CreateRoleRequest) (*pb.CreateRoleResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#delete-role
func (srv *metathingsIdentityService) DeleteRole(context.Context, *pb.DeleteRoleRequest) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#update-role
func (srv *metathingsIdentityService) PatchRole(context.Context, *pb.PatchRoleRequest) (*pb.PatchRoleResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#show-role-details
func (srv *metathingsIdentityService) GetRole(context.Context, *pb.GetRoleRequest) (*pb.GetRoleResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-roles
func (srv *metathingsIdentityService) ListRoles(context.Context, *pb.ListRolesRequest) (*pb.ListRolesResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#assign-role-to-group-on-domain
func (srv *metathingsIdentityService) AddRoleToGroupOnDomain(context.Context, *pb.AddRoleToGroupOnDomainRequest) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#unassign-role-from-group-on-domain
func (srv *metathingsIdentityService) RemoveRoleFromGroupOnDomain(context.Context, *pb.RemoveRoleFromGroupOnDomainRequest) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#check-whether-user-has-role-assignment-on-domain
func (srv *metathingsIdentityService) CheckRoleInGroupOnDomain(context.Context, *pb.CheckRoleInGroupOnDomainRequest) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-role-assignments-for-group-on-domain
func (srv *metathingsIdentityService) ListRolesForGroupOnDomain(context.Context, *pb.ListRolesForGroupOnDomainRequest) (*pb.ListRolesForGroupOnDomainResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-role-assignments-for-group-on-domain
func (srv *metathingsIdentityService) AddRoleToUserOnDomain(context.Context, *pb.AddRoleToUserOnDomainRequest) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-role-assignments-for-group-on-domain
func (srv *metathingsIdentityService) RemoveRoleFromUserOnDomain(context.Context, *pb.RemoveRoleFromUserOnDomainRequest) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-role-assignments-for-group-on-domain
func (srv *metathingsIdentityService) CheckRoleInUserOnDomain(context.Context, *pb.CheckRoleInUserOnDomainRequest) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-role-assignments-for-group-on-domain
func (srv *metathingsIdentityService) ListRolesForUserOnDomain(context.Context, *pb.ListRolesForUserOnDomainRequest) (*pb.ListRolesForUserOnDomainResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#assign-role-to-group-on-project
func (srv *metathingsIdentityService) AddRoleToGroupOnProject(context.Context, *pb.AddRoleToGroupOnProjectRequest) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#assign-role-to-group-on-project
func (srv *metathingsIdentityService) RemoveRoleFromGroupOnProject(context.Context, *pb.RemoveRoleFromGroupOnProjectRequest) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#check-whether-group-has-role-assignment-on-project
func (srv *metathingsIdentityService) CheckRoleInGroupOnProject(context.Context, *pb.CheckRoleInGroupOnProjectRequest) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-role-assignments-for-user-on-project
func (srv *metathingsIdentityService) ListRolesForGroupOnProject(context.Context, *pb.ListRolesForGroupOnProjectRequest) (*pb.ListRolesForGroupOnProjectResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#assign-role-to-user-on-project
func (srv *metathingsIdentityService) AddRoleToUserOnProject(context.Context, *pb.AddRoleToUserOnProjectRequest) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#unassign-role-from-user-on-project
func (srv *metathingsIdentityService) RemoveRoleFromUserOnProject(context.Context, *pb.RemoveRoleFromUserOnProjectRequest) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#check-whether-user-has-role-assignment-on-project
func (srv *metathingsIdentityService) CheckRoleInUserOnProject(context.Context, *pb.CheckRoleInUserOnProjectRequest) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-role-assignments-for-user-on-project
func (srv *metathingsIdentityService) ListRolesForUserOnProject(context.Context, *pb.ListRolesForUserOnProjectRequest) (*pb.ListRolesForUserOnProjectResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// password authentication with unsnscoped authorization
// https://developer.openstack.org/api-ref/identity/v3/index.html#password-authentication-with-unscoped-authorization
// password authentication with scoped authorization
// https://developer.openstack.org/api-ref/identity/v3/index.html#password-authentication-with-scoped-authorization
// password authentication with unscoped authorization
// https://developer.openstack.org/api-ref/identity/v3/index.html#password-authentication-with-explicit-unscoped-authorization
// token authentication with unscoped authorization
// https://developer.openstack.org/api-ref/identity/v3/index.html#token-authentication-with-unscoped-authorization
// token authentication with scoped authorization
// https://developer.openstack.org/api-ref/identity/v3/index.html#token-authentication-with-scoped-authorization
// token authentication with explicit unscoped authorization
// https://developer.openstack.org/api-ref/identity/v3/index.html#token-authentication-with-explicit-unscoped-authorization
// application credential authorization
// https://developer.openstack.org/api-ref/identity/v3/index.html#authenticating-with-an-application-credential
func (srv *metathingsIdentityService) IssueToken(ctx context.Context, req *pb.IssueTokenRequest) (*pb.IssueTokenResponse, error) {
	body, err := encode_decode.EncodeIssueTokenRequest(ctx, req)
	if err != nil {
		switch err {
		case encode_decode.Unimplemented:
			return nil, grpc.Errorf(codes.Unauthenticated, "unimplement")
		default:
			return nil, grpc.Errorf(codes.Internal, "internal error")
		}
	}
	_, rbody, _ := gorequest.New().Post("http://fls-vps:35357/v3/auth/tokens").Send(&body).End()
	fmt.Println(rbody)

	return nil, nil
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#revoke-token
func (srv *metathingsIdentityService) RevokeToken(context.Context, *google_protobuf3.Empty) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#create-application-credential
func (srv *metathingsIdentityService) CreateApplicationCredential(context.Context, *pb.CreateApplicationCredentialRequest) (*pb.CreateApplicationCredentialResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#delete-application-credential
func (srv *metathingsIdentityService) DeleteApplicationCredential(context.Context, *pb.DeleteApplicationCredentialRequest) (*google_protobuf3.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#show-application-credential-details
func (srv *metathingsIdentityService) GetApplicationCredential(context.Context, *pb.GetApplicationCredentialRequest) (*pb.GetApplicationCredentialResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-application-credentials
func (srv *metathingsIdentityService) ListApplicationCredentials(context.Context, *pb.ListApplicationCredentialsRequest) (*pb.ListApplicationCredentialsResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

func NewIdentityService(opt ...ServiceOptions) *metathingsIdentityService {
	opts := defaultServiceOptions
	for _, o := range opt {
		o(&opts)
	}

	return &metathingsIdentityService{
		opts: opts,
	}
}
