package metathings_identity_service

import (
	"context"
	"fmt"
	"net/url"
	"path"

	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/parnurzeal/gorequest"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	grpc_helper "github.com/bigdatagz/metathings/pkg/common/grpc"
	"github.com/bigdatagz/metathings/pkg/common/log"
	codec "github.com/bigdatagz/metathings/pkg/identity/service/encode_decode"
	pb "github.com/bigdatagz/metathings/pkg/proto/identity"
)

type options struct {
	keystoneBaseURL string
	logLevel        string
}

var defaultServiceOptions = options{
	logLevel: "info",
}

type ServiceOptions func(*options)

func SetKeystoneBaseURL(url string) ServiceOptions {
	return func(o *options) {
		o.keystoneBaseURL = url
	}
}

func SetLogLevel(lvl string) ServiceOptions {
	return func(o *options) {
		o.logLevel = lvl
	}
}

type metathingsIdentityService struct {
	grpc_helper.AuthorizationTokenParser

	logger log.FieldLogger
	h      *helper
	opts   options
}

type helper struct {
	srv *metathingsIdentityService
}

func (h *helper) JoinURL(p string) string {
	url_str := h.srv.opts.keystoneBaseURL
	u, err := url.Parse(url_str)
	if err != nil {
		h.srv.logger.Errorf("bad keystone base url: %v, error: %v\n", url_str, err)
		return ""
	}
	u.Path = path.Join(u.Path, p)
	return u.String()
}

func (h *helper) SendHeader(ctx context.Context, pairs ...string) error {
	return grpc.SendHeader(ctx, metadata.Pairs(pairs...))
}

func (srv *metathingsIdentityService) ignoreAuthMethods() []string {
	methods := []string{
		"IssueToken",
		"CheckToken",
		"ValidateToken",
	}
	return methods
}

func (srv *metathingsIdentityService) validateTokenViaHTTP(token, subject_token string) (gorequest.Response, string, error) {
	url := srv.h.JoinURL("/v3/auth/tokens")

	http_res, http_body, errs := gorequest.
		New().Get(url).
		Set("X-Auth-Token", token).
		Set("X-Subject-Token", token).
		Query("nocatalog=1").End()

	if len(errs) > 0 {
		return nil, "", errs[0]
	}

	if http_res.StatusCode != 201 {
		return nil, "", status.Errorf(codes.Unauthenticated, fmt.Sprintf("%v", http_body))
	}

	return http_res, http_body, nil
}

func (srv *metathingsIdentityService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	methDesc, _ := grpc_helper.ParseMethodDescription(fullMethodName)
	for _, m := range srv.ignoreAuthMethods() {
		if m == methDesc.Method {
			return ctx, nil
		}
	}

	token, err := srv.GetTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}

	subject_token, err := srv.GetSubjectTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}

	http_res, http_body, err := srv.validateTokenViaHTTP(token, subject_token)
	if err != nil {
		srv.logger.
			WithField("error", err).
			Errorf("failed to validate token via http")
		return nil, err
	}

	if http_res.StatusCode != 200 {
		srv.logger.
			WithField("status_code", http_res.StatusCode).
			Errorf("unauthenticated")
		return nil, Unauthenticated
	}

	cred, err := codec.DecodeValidateTokenResponse(http_res, http_body)
	if err != nil {
		srv.logger.
			WithField("error", err).
			Errorf("failed to decode validate token response")
		return nil, err
	}

	ctx = context.WithValue(ctx, "token", token)
	ctx = context.WithValue(ctx, "credential", cred.Token)

	srv.logger.WithFields(log.Fields{
		"package":    methDesc.Package,
		"service":    methDesc.Service,
		"method":     methDesc.Method,
		"token":      token,
		"credential": cred.Token,
	}).Debugf("validate token")

	return ctx, nil
}

// https://developer.openstack.org/api-ref/identity/v3/#create-region
func (srv *metathingsIdentityService) CreateRegion(context.Context, *pb.CreateRegionRequest) (*pb.CreateRegionResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/#delete-region
func (srv *metathingsIdentityService) DeleteRegion(context.Context, *pb.DeleteRegionRequest) (*empty.Empty, error) {
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
func (srv *metathingsIdentityService) DeleteDomain(context.Context, *pb.DeleteDomainRequest) (*empty.Empty, error) {
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
func (srv *metathingsIdentityService) DeleteProject(context.Context, *pb.DeleteProjectRequest) (*empty.Empty, error) {
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
func (srv *metathingsIdentityService) DeleteUser(context.Context, *pb.DeleteUserRequest) (*empty.Empty, error) {
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
func (srv *metathingsIdentityService) ChangePassword(context.Context, *pb.ChangePasswordRequest) (*empty.Empty, error) {
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
func (srv *metathingsIdentityService) DeleteGroup(context.Context, *pb.DeleteGroupRequest) (*empty.Empty, error) {
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
func (srv *metathingsIdentityService) AddUserToGroup(context.Context, *pb.AddUserToGroupRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#remove-user-from-group
func (srv *metathingsIdentityService) RemoveUserFromGroup(context.Context, *pb.RemoveUserFromGroupRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#check-whether-user-belongs-to-group
func (srv *metathingsIdentityService) CheckUserInGroup(context.Context, *pb.CheckUserInGroupRequest) (*empty.Empty, error) {
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
func (srv *metathingsIdentityService) DeleteRole(context.Context, *pb.DeleteRoleRequest) (*empty.Empty, error) {
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
func (srv *metathingsIdentityService) AddRoleToGroupOnDomain(context.Context, *pb.AddRoleToGroupOnDomainRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#unassign-role-from-group-on-domain
func (srv *metathingsIdentityService) RemoveRoleFromGroupOnDomain(context.Context, *pb.RemoveRoleFromGroupOnDomainRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#check-whether-user-has-role-assignment-on-domain
func (srv *metathingsIdentityService) CheckRoleInGroupOnDomain(context.Context, *pb.CheckRoleInGroupOnDomainRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-role-assignments-for-group-on-domain
func (srv *metathingsIdentityService) ListRolesForGroupOnDomain(context.Context, *pb.ListRolesForGroupOnDomainRequest) (*pb.ListRolesForGroupOnDomainResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-role-assignments-for-group-on-domain
func (srv *metathingsIdentityService) AddRoleToUserOnDomain(context.Context, *pb.AddRoleToUserOnDomainRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-role-assignments-for-group-on-domain
func (srv *metathingsIdentityService) RemoveRoleFromUserOnDomain(context.Context, *pb.RemoveRoleFromUserOnDomainRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-role-assignments-for-group-on-domain
func (srv *metathingsIdentityService) CheckRoleInUserOnDomain(context.Context, *pb.CheckRoleInUserOnDomainRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-role-assignments-for-group-on-domain
func (srv *metathingsIdentityService) ListRolesForUserOnDomain(context.Context, *pb.ListRolesForUserOnDomainRequest) (*pb.ListRolesForUserOnDomainResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#assign-role-to-group-on-project
func (srv *metathingsIdentityService) AddRoleToGroupOnProject(context.Context, *pb.AddRoleToGroupOnProjectRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#assign-role-to-group-on-project
func (srv *metathingsIdentityService) RemoveRoleFromGroupOnProject(context.Context, *pb.RemoveRoleFromGroupOnProjectRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#check-whether-group-has-role-assignment-on-project
func (srv *metathingsIdentityService) CheckRoleInGroupOnProject(context.Context, *pb.CheckRoleInGroupOnProjectRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-role-assignments-for-user-on-project
func (srv *metathingsIdentityService) ListRolesForGroupOnProject(context.Context, *pb.ListRolesForGroupOnProjectRequest) (*pb.ListRolesForGroupOnProjectResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#assign-role-to-user-on-project
func (srv *metathingsIdentityService) AddRoleToUserOnProject(context.Context, *pb.AddRoleToUserOnProjectRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#unassign-role-from-user-on-project
func (srv *metathingsIdentityService) RemoveRoleFromUserOnProject(context.Context, *pb.RemoveRoleFromUserOnProjectRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#check-whether-user-has-role-assignment-on-project
func (srv *metathingsIdentityService) CheckRoleInUserOnProject(context.Context, *pb.CheckRoleInUserOnProjectRequest) (*empty.Empty, error) {
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
	body, err := codec.EncodeIssueTokenRequest(ctx, req)
	if err != nil {
		switch err {
		case codec.Unimplemented:
			return nil, status.Errorf(codes.Unimplemented, "unimplement")
		default:
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}
	url := srv.h.JoinURL("/v3/auth/tokens")
	http_res, http_body, errs := gorequest.
		New().Post(url).
		Query("nocatalog=1").
		Send(&body).End()
	if len(errs) > 0 {
		srv.logger.
			WithField("error", errs[0]).
			Errorf("failed to keystone issue token")
		return nil, status.Errorf(codes.Internal, errs[0].Error())
	}

	if http_res.StatusCode != 201 {
		srv.logger.
			WithField("status_code", http_res.StatusCode).
			Errorf("unexpected status code")
		return nil, status.Errorf(codes.Unauthenticated, http_body)
	}

	token_str := http_res.Header.Get("X-Subject-Token")
	srv.logger.WithField("token", token_str).Debugf("got token from keystone")
	err = srv.h.SendHeader(ctx, "authorization", fmt.Sprintf("mt %v", token_str))
	if err != nil {
		srv.logger.
			WithField("error", err).
			Warningf("failed to send headers")
	}

	res, err := codec.DecodeIssueTokenResponse(http_res, http_body)
	if err != nil {
		srv.logger.
			WithFields(log.Fields{}).
			Errorf("failed to decode issue token response")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	srv.logger.WithFields(log.Fields{
		"user_id": res.Token.User.Id,
		"user":    res.Token.User.Name,
	}).Infof("issue token")
	return res, nil
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#revoke-token
func (srv *metathingsIdentityService) RevokeToken(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#check-token
func (srv *metathingsIdentityService) CheckToken(ctx context.Context, _ *empty.Empty) (*empty.Empty, error) {
	token, err := srv.GetTokenFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	url := srv.h.JoinURL("/v3/auth/tokens")
	http_res, _, errs := gorequest.
		New().Head(url).
		Set("X-Auth-Token", token).
		Set("X-Subject-Token", token).
		End()
	if len(errs) > 0 {
		return nil, status.Errorf(codes.Internal, errs[0].Error())
	}

	if http_res.StatusCode != 200 {
		return nil, status.Errorf(codes.Unauthenticated, Unauthenticated.Error())
	}

	return &empty.Empty{}, nil
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#validate-and-show-information-for-token
func (srv *metathingsIdentityService) ValidateToken(ctx context.Context, _ *empty.Empty) (*pb.ValidateTokenResponse, error) {
	token, err := srv.GetTokenFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	sub_token, err := srv.GetSubjectTokenFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	srv.logger.
		WithFields(log.Fields{"token": token, "sub_token": sub_token}).
		Debugf("validating token")

	url := srv.h.JoinURL("/v3/auth/tokens")
	http_res, http_body, errs := gorequest.
		New().Get(url).
		Set("X-Auth-Token", token).
		Set("X-Subject-Token", sub_token).
		End()
	if len(errs) > 0 {
		srv.logger.
			WithField("error", errs[0]).
			Errorf("failed to validate token via http")
		return nil, status.Errorf(codes.Internal, errs[0].Error())
	}

	if http_res.StatusCode != 200 {
		srv.logger.
			WithField("status_code", http_res.StatusCode).
			WithField("http_body", http_body).
			Errorf("unexpected status code")
		return nil, status.Errorf(codes.Unauthenticated, Unauthenticated.Error())
	}

	res, err := codec.DecodeValidateTokenResponse(http_res, http_body)
	if err != nil {
		srv.logger.
			WithFields(log.Fields{
				"error": err,
				"body":  http_body,
			}).
			Errorf("failed to decode validate token response")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return res, nil
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#create-application-credential
func (srv *metathingsIdentityService) CreateApplicationCredential(context.Context, *pb.CreateApplicationCredentialRequest) (*pb.CreateApplicationCredentialResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#delete-application-credential
func (srv *metathingsIdentityService) DeleteApplicationCredential(context.Context, *pb.DeleteApplicationCredentialRequest) (*empty.Empty, error) {
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

func NewIdentityService(opt ...ServiceOptions) (*metathingsIdentityService, error) {
	opts := defaultServiceOptions
	for _, o := range opt {
		o(&opts)
	}

	logger, err := log_helper.NewLogger("identityd", opts.logLevel)
	if err != nil {
		log.WithError(err).Errorf("failed to new logger")
		return nil, err
	}

	srv := &metathingsIdentityService{
		opts:   opts,
		logger: logger,
	}
	srv.h = &helper{srv}

	return srv, nil
}
