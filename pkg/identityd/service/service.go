package metathings_identityd_service

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

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	log_helper "github.com/nayotta/metathings/pkg/common/log"
	codec "github.com/nayotta/metathings/pkg/identityd/service/encode_decode"
	pb "github.com/nayotta/metathings/pkg/proto/identityd"
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

type metathingsIdentitydService struct {
	grpc_helper.AuthorizationTokenParser

	logger log.FieldLogger
	h      *helper
	opts   options
}

type helper struct {
	srv *metathingsIdentitydService
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

func (srv *metathingsIdentitydService) ignoreAuthMethods() []string {
	methods := []string{
		"IssueToken",
		"CheckToken",
		"ValidateToken",
	}
	return methods
}

func (srv *metathingsIdentitydService) validateTokenViaHTTP(token, subject_token string) (gorequest.Response, string, error) {
	url := srv.h.JoinURL("/v3/auth/tokens")

	http_res, http_body, errs := gorequest.New().Get(url).Set("X-Auth-Token", token).Set("X-Subject-Token", token).Query("nocatalog=1").End()

	if len(errs) > 0 {
		return nil, "", errs[0]
	}

	if http_res.StatusCode != 201 {
		return nil, "", status.Errorf(codes.Unauthenticated, fmt.Sprintf("%v", http_body))
	}

	return http_res, http_body, nil
}

func (srv *metathingsIdentitydService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
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
		srv.logger.WithField("error", err).Errorf("failed to validate token via http")
		return nil, err
	}

	if http_res.StatusCode != 200 {
		srv.logger.WithField("status_code", http_res.StatusCode).Errorf("unauthenticated")
		return nil, Unauthenticated
	}

	cred, err := codec.DecodeValidateTokenResponse(http_res, http_body)
	if err != nil {
		srv.logger.WithField("error", err).Errorf("failed to decode validate token response")
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
func (srv *metathingsIdentitydService) CreateRegion(context.Context, *pb.CreateRegionRequest) (*pb.CreateRegionResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/#delete-region
func (srv *metathingsIdentitydService) DeleteRegion(context.Context, *pb.DeleteRegionRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/#update-region
func (srv *metathingsIdentitydService) PatchRegion(context.Context, *pb.PatchRegionRequest) (*pb.PatchRegionResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/#show-region-details
func (srv *metathingsIdentitydService) GetRegion(context.Context, *pb.GetRegionRequest) (*pb.GetRegionResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/#list-regions
func (srv *metathingsIdentitydService) ListRegions(context.Context, *pb.ListRegionsRequest) (*pb.ListRegionsResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#create-domain
func (srv *metathingsIdentitydService) CreateDomain(context.Context, *pb.CreateDomainRequest) (*pb.CreateDomainResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#delete-domain
func (srv *metathingsIdentitydService) DeleteDomain(context.Context, *pb.DeleteDomainRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#update-domain
func (srv *metathingsIdentitydService) PatchDomain(context.Context, *pb.PatchDomainRequest) (*pb.PatchDomainResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#show-domain-details
func (srv *metathingsIdentitydService) GetDomain(context.Context, *pb.GetDomainRequest) (*pb.GetDomainResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-domains
func (srv *metathingsIdentitydService) ListDomains(context.Context, *pb.ListDomainsRequest) (*pb.ListDomainsResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#create-project
func (srv *metathingsIdentitydService) CreateProject(context.Context, *pb.CreateProjectRequest) (*pb.CreateProjectResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#delete-project
func (srv *metathingsIdentitydService) DeleteProject(context.Context, *pb.DeleteProjectRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#update-project
func (srv *metathingsIdentitydService) PatchProject(context.Context, *pb.PatchProjectRequest) (*pb.PatchProjectResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#show-project-details
func (srv *metathingsIdentitydService) GetProject(context.Context, *pb.GetProjectRequest) (*pb.GetProjectResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-projects
func (srv *metathingsIdentitydService) ListProjects(context.Context, *pb.ListProjectsRequest) (*pb.ListProjectsResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-projects-for-user
func (srv *metathingsIdentitydService) ListProjectsForUser(context.Context, *pb.ListProjectsForUserRequest) (*pb.ListProjectsForUserResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#create-user
func (srv *metathingsIdentitydService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	body, err := codec.EncodeCreateUser(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	url := srv.h.JoinURL("/v3/users")
	http_res, http_body, errs := gorequest.New().Post(url).Send(body).End()
	if len(errs) > 0 {
		srv.logger.WithError(errs[0]).Errorf("failed to keystone create user")
		return nil, status.Errorf(codes.Internal, errs[0].Error())
	}

	if http_res.StatusCode != 201 {
		srv.logger.WithField("status_code", http_res.StatusCode).Errorf("unexpected status code")
		return nil, status.Errorf(grpc_helper.HttpStatusCode2GrpcStatusCode(http_res.StatusCode), http_body)
	}

	res, err := codec.DecodeCreateUser(http_res, http_body)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to decode create user response")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	srv.logger.WithFields(log.Fields{
		"user_id":   res.User.Id,
		"user_name": res.User.Name,
	}).Infof("create user")

	return res, nil
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#delete-user
func (srv *metathingsIdentitydService) DeleteUser(context.Context, *pb.DeleteUserRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#update-user
func (srv *metathingsIdentitydService) PatchUser(context.Context, *pb.PatchUserRequest) (*pb.PatchUserResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#show-user-details
func (srv *metathingsIdentitydService) GetUser(context.Context, *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-users
func (srv *metathingsIdentitydService) ListUsers(context.Context, *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#change-password-for-user
func (srv *metathingsIdentitydService) ChangePassword(context.Context, *pb.ChangePasswordRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-users-in-group
func (srv *metathingsIdentitydService) ListUsersInGroup(context.Context, *pb.ListUsersInGroupRequest) (*pb.ListUsersInGroupResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#create-group
func (srv *metathingsIdentitydService) CreateGroup(context.Context, *pb.CreateGroupRequest) (*pb.CreateGroupResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#delete-group
func (srv *metathingsIdentitydService) DeleteGroup(context.Context, *pb.DeleteGroupRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#update-group
func (srv *metathingsIdentitydService) PatchGroup(context.Context, *pb.PatchGroupRequest) (*pb.PatchGroupResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#show-group-details
func (srv *metathingsIdentitydService) GetGroup(context.Context, *pb.GetGroupRequest) (*pb.GetGroupResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-groups
func (srv *metathingsIdentitydService) ListGroups(context.Context, *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#add-user-to-group
func (srv *metathingsIdentitydService) AddUserToGroup(context.Context, *pb.AddUserToGroupRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#remove-user-from-group
func (srv *metathingsIdentitydService) RemoveUserFromGroup(context.Context, *pb.RemoveUserFromGroupRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#check-whether-user-belongs-to-group
func (srv *metathingsIdentitydService) CheckUserInGroup(context.Context, *pb.CheckUserInGroupRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-groups-to-which-a-user-belongs
func (srv *metathingsIdentitydService) ListGroupsForUser(context.Context, *pb.ListGroupsForUserRequest) (*pb.ListGroupsForUserResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#create-role
func (srv *metathingsIdentitydService) CreateRole(context.Context, *pb.CreateRoleRequest) (*pb.CreateRoleResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#delete-role
func (srv *metathingsIdentitydService) DeleteRole(context.Context, *pb.DeleteRoleRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#update-role
func (srv *metathingsIdentitydService) PatchRole(context.Context, *pb.PatchRoleRequest) (*pb.PatchRoleResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#show-role-details
func (srv *metathingsIdentitydService) GetRole(context.Context, *pb.GetRoleRequest) (*pb.GetRoleResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-roles
func (srv *metathingsIdentitydService) ListRoles(context.Context, *pb.ListRolesRequest) (*pb.ListRolesResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#assign-role-to-group-on-domain
func (srv *metathingsIdentitydService) AddRoleToGroupOnDomain(context.Context, *pb.AddRoleToGroupOnDomainRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#unassign-role-from-group-on-domain
func (srv *metathingsIdentitydService) RemoveRoleFromGroupOnDomain(context.Context, *pb.RemoveRoleFromGroupOnDomainRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#check-whether-user-has-role-assignment-on-domain
func (srv *metathingsIdentitydService) CheckRoleInGroupOnDomain(context.Context, *pb.CheckRoleInGroupOnDomainRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-role-assignments-for-group-on-domain
func (srv *metathingsIdentitydService) ListRolesForGroupOnDomain(context.Context, *pb.ListRolesForGroupOnDomainRequest) (*pb.ListRolesForGroupOnDomainResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-role-assignments-for-group-on-domain
func (srv *metathingsIdentitydService) AddRoleToUserOnDomain(context.Context, *pb.AddRoleToUserOnDomainRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-role-assignments-for-group-on-domain
func (srv *metathingsIdentitydService) RemoveRoleFromUserOnDomain(context.Context, *pb.RemoveRoleFromUserOnDomainRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-role-assignments-for-group-on-domain
func (srv *metathingsIdentitydService) CheckRoleInUserOnDomain(context.Context, *pb.CheckRoleInUserOnDomainRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-role-assignments-for-group-on-domain
func (srv *metathingsIdentitydService) ListRolesForUserOnDomain(context.Context, *pb.ListRolesForUserOnDomainRequest) (*pb.ListRolesForUserOnDomainResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#assign-role-to-group-on-project
func (srv *metathingsIdentitydService) AddRoleToGroupOnProject(context.Context, *pb.AddRoleToGroupOnProjectRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#assign-role-to-group-on-project
func (srv *metathingsIdentitydService) RemoveRoleFromGroupOnProject(context.Context, *pb.RemoveRoleFromGroupOnProjectRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#check-whether-group-has-role-assignment-on-project
func (srv *metathingsIdentitydService) CheckRoleInGroupOnProject(context.Context, *pb.CheckRoleInGroupOnProjectRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-role-assignments-for-user-on-project
func (srv *metathingsIdentitydService) ListRolesForGroupOnProject(context.Context, *pb.ListRolesForGroupOnProjectRequest) (*pb.ListRolesForGroupOnProjectResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#assign-role-to-user-on-project
func (srv *metathingsIdentitydService) AddRoleToUserOnProject(context.Context, *pb.AddRoleToUserOnProjectRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#unassign-role-from-user-on-project
func (srv *metathingsIdentitydService) RemoveRoleFromUserOnProject(context.Context, *pb.RemoveRoleFromUserOnProjectRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#check-whether-user-has-role-assignment-on-project
func (srv *metathingsIdentitydService) CheckRoleInUserOnProject(context.Context, *pb.CheckRoleInUserOnProjectRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-role-assignments-for-user-on-project
func (srv *metathingsIdentitydService) ListRolesForUserOnProject(context.Context, *pb.ListRolesForUserOnProjectRequest) (*pb.ListRolesForUserOnProjectResponse, error) {
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
func (srv *metathingsIdentitydService) IssueToken(ctx context.Context, req *pb.IssueTokenRequest) (*pb.IssueTokenResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	body, err := codec.EncodeIssueTokenRequest(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	url := srv.h.JoinURL("/v3/auth/tokens")
	http_res, http_body, errs := gorequest.New().Post(url).Query("nocatalog=1").Send(body).End()
	if len(errs) > 0 {
		srv.logger.WithField("error", errs[0]).Errorf("failed to keystone issue token")
		return nil, status.Errorf(codes.Internal, errs[0].Error())
	}

	if http_res.StatusCode != 201 {
		srv.logger.WithField("status_code", http_res.StatusCode).Errorf("unexpected status code")
		return nil, status.Errorf(codes.Unauthenticated, http_body)
	}

	token_str := http_res.Header.Get("X-Subject-Token")
	srv.logger.WithField("token", token_str).Debugf("got token from keystone")
	err = srv.h.SendHeader(ctx, "authorization", fmt.Sprintf("mt %v", token_str))
	if err != nil {
		srv.logger.WithField("error", err).Warningf("failed to send headers")
	}

	res, err := codec.DecodeIssueTokenResponse(http_res, http_body)
	if err != nil {
		srv.logger.WithFields(log.Fields{}).Errorf("failed to decode issue token response")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	srv.logger.WithFields(log.Fields{
		"user_id": res.Token.User.Id,
		"user":    res.Token.User.Name,
	}).Infof("issue token")
	return res, nil
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#revoke-token
func (srv *metathingsIdentitydService) RevokeToken(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#check-token
func (srv *metathingsIdentitydService) CheckToken(ctx context.Context, _ *empty.Empty) (*empty.Empty, error) {
	token, err := srv.GetTokenFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	url := srv.h.JoinURL("/v3/auth/tokens")
	http_res, _, errs := gorequest.New().Head(url).Set("X-Auth-Token", token).Set("X-Subject-Token", token).End()
	if len(errs) > 0 {
		return nil, status.Errorf(codes.Internal, errs[0].Error())
	}

	if http_res.StatusCode != 200 {
		return nil, status.Errorf(codes.Unauthenticated, Unauthenticated.Error())
	}

	return &empty.Empty{}, nil
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#validate-and-show-information-for-token
func (srv *metathingsIdentitydService) ValidateToken(ctx context.Context, _ *empty.Empty) (*pb.ValidateTokenResponse, error) {
	token, err := srv.GetTokenFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	sub_token, err := srv.GetSubjectTokenFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	srv.logger.WithFields(log.Fields{"token": token, "sub_token": sub_token}).Debugf("validating token")

	url := srv.h.JoinURL("/v3/auth/tokens")
	http_res, http_body, errs := gorequest.New().Get(url).Set("X-Auth-Token", token).Set("X-Subject-Token", sub_token).End()
	if len(errs) > 0 {
		srv.logger.WithField("error", errs[0]).Errorf("failed to validate token via http")
		return nil, status.Errorf(codes.Internal, errs[0].Error())
	}

	if http_res.StatusCode != 200 {
		srv.logger.WithFields(log.Fields{
			"status_code": http_res.StatusCode,
			"http_body":   http_body,
		}).Errorf("unexpected status code")
		return nil, status.Errorf(grpc_helper.HttpStatusCode2GrpcStatusCode(http_res.StatusCode), http_body)
	}

	res, err := codec.DecodeValidateTokenResponse(http_res, http_body)
	if err != nil {
		srv.logger.WithFields(log.Fields{
			"error": err,
			"body":  http_body,
		}).Errorf("failed to decode validate token response")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return res, nil
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#create-application-credential
func (srv *metathingsIdentitydService) CreateApplicationCredential(ctx context.Context, req *pb.CreateApplicationCredentialRequest) (*pb.CreateApplicationCredentialResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	body, err := codec.EncodeCreateApplicationCredential(ctx, req)
	if err != nil {
		switch err {
		case codec.Unimplemented:
			return nil, status.Errorf(codes.Unimplemented, "unimplemented")
		default:
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	user_id := req.GetUserId().GetValue()
	url := srv.h.JoinURL("/v3/users/" + user_id + "/application_credentials")
	http_res, http_body, errs := gorequest.New().Post(url).Send(body).End()
	if len(errs) > 0 {
		srv.logger.WithError(errs[0]).Errorf("failed to keystone create application credential")
		return nil, status.Errorf(codes.Internal, errs[0].Error())
	}

	if http_res.StatusCode != 201 {
		srv.logger.WithField("status_code", http_res.StatusCode).Errorf("unexpected status code")
		return nil, status.Errorf(grpc_helper.HttpStatusCode2GrpcStatusCode(http_res.StatusCode), http_body)
	}

	res, err := codec.DecodeCreateApplicationCredential(http_res, http_body)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to decode create application credential response")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	srv.logger.WithFields(log.Fields{
		"user_id":                     user_id,
		"application_credential_id":   res.ApplicationCredential.Id,
		"application_credential_name": res.ApplicationCredential.Name,
	}).Infof("create application credential")

	return res, nil
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#delete-application-credential
func (srv *metathingsIdentitydService) DeleteApplicationCredential(ctx context.Context, req *pb.DeleteApplicationCredentialRequest) (*empty.Empty, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	user_id := req.GetUserId().GetValue()
	app_cred_id := req.GetApplicationCredentialId().GetValue()
	url := srv.h.JoinURL("/v3/users/" + user_id + "/application_credentials/" + app_cred_id)

	http_res, http_body, errs := gorequest.New().Delete(url).End()

	if len(errs) > 0 {
		srv.logger.WithError(errs[0]).Errorf("failed to delete application credential via http")
		return nil, status.Errorf(codes.Internal, errs[0].Error())
	}

	if http_res.StatusCode != 204 {
		srv.logger.WithFields(log.Fields{
			"status_code": http_res.StatusCode,
			"http_body":   http_body,
		}).Errorf("unexpected status code")
		return nil, status.Errorf(grpc_helper.HttpStatusCode2GrpcStatusCode(http_res.StatusCode), http_body)
	}

	return &empty.Empty{}, nil
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#show-application-credential-details
func (srv *metathingsIdentitydService) GetApplicationCredential(ctx context.Context, req *pb.GetApplicationCredentialRequest) (*pb.GetApplicationCredentialResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	user_id := req.GetUserId().GetValue()
	app_cred_id := req.GetApplicationCredentialId().GetValue()
	url := srv.h.JoinURL("/v3/users/" + user_id + "/application_credentials/" + app_cred_id)

	http_res, http_body, errs := gorequest.New().Get(url).End()
	if len(errs) > 0 {
		srv.logger.WithError(errs[0]).Errorf("failed to get application credential via http")
		return nil, status.Errorf(codes.Internal, errs[0].Error())
	}

	if http_res.StatusCode != 200 {
		srv.logger.WithFields(log.Fields{
			"status_code": http_res.StatusCode,
			"http_body":   http_body,
		}).Errorf("unexpected status code")
		return nil, status.Errorf(grpc_helper.HttpStatusCode2GrpcStatusCode(http_res.StatusCode), http_body)
	}

	res, err := codec.DecodeGetApplicationCredential(http_res, http_body)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to decode get application credential response")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return res, nil
}

// https://developer.openstack.org/api-ref/identity/v3/index.html#list-application-credentials
func (srv *metathingsIdentitydService) ListApplicationCredentials(ctx context.Context, req *pb.ListApplicationCredentialsRequest) (*pb.ListApplicationCredentialsResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	user_id := req.GetUserId().GetValue()
	url := srv.h.JoinURL("/v3/users/" + user_id + "/application_credentials")

	http_res, http_body, errs := gorequest.New().Get(url).End()
	if len(errs) > 0 {
		srv.logger.WithError(errs[0]).Errorf("failed to list application credential via http")
		return nil, status.Errorf(codes.Internal, errs[0].Error())
	}

	if http_res.StatusCode != 200 {
		srv.logger.WithFields(log.Fields{
			"status_code": http_res.StatusCode,
			"http_body":   http_body,
		}).Errorf("unexpected status code")
		return nil, status.Errorf(grpc_helper.HttpStatusCode2GrpcStatusCode(http_res.StatusCode), http_body)
	}

	res, err := codec.DecodeListApplicationCredential(http_res, http_body)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to decode list application credential response")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return res, nil
}

func NewIdentitydService(opt ...ServiceOptions) (*metathingsIdentitydService, error) {
	opts := defaultServiceOptions
	for _, o := range opt {
		o(&opts)
	}

	logger, err := log_helper.NewLogger("identityd", opts.logLevel)
	if err != nil {
		log.WithError(err).Errorf("failed to new logger")
		return nil, err
	}

	srv := &metathingsIdentitydService{
		opts:   opts,
		logger: logger,
	}
	srv.h = &helper{srv}

	return srv, nil
}
