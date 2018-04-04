// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: add_role_to_group_on_domain.proto

/*
Package identity is a generated protocol buffer package.

It is generated from these files:
	add_role_to_group_on_domain.proto
	add_role_to_group_on_project.proto
	add_role_to_user_on_domain.proto
	add_role_to_user_on_project.proto
	add_user_to_group.proto
	change_password.proto
	check_role_in_group_on_domain.proto
	check_role_in_group_on_project.proto
	check_role_in_user_on_domain.proto
	check_role_in_user_on_project.proto
	check_user_in_group.proto
	create_domain.proto
	create_group.proto
	create_project.proto
	create_role.proto
	create_user.proto
	delete_domain.proto
	delete_group.proto
	delete_project.proto
	delete_role.proto
	delete_user.proto
	domain.proto
	get_domain.proto
	get_group.proto
	get_project.proto
	get_role.proto
	get_user.proto
	group.proto
	list_domains.proto
	list_groups.proto
	list_groups_for_user.proto
	list_projects.proto
	list_projects_for_user.proto
	list_role_in_group_on_domain.proto
	list_roles.proto
	list_roles_for_group_on_domain.proto
	list_roles_for_group_on_project.proto
	list_roles_for_user_on_domain.proto
	list_roles_for_user_on_project.proto
	list_users.proto
	list_users_in_group.proto
	patch_domain.proto
	patch_group.proto
	patch_project.proto
	patch_role.proto
	patch_user.proto
	project.proto
	remove_role_from_group_on_domain.proto
	remove_role_from_group_on_project.proto
	remove_role_from_user_on_domain.proto
	remove_role_from_user_on_project.proto
	remove_user_from_group.proto
	role.proto
	service.proto
	user.proto

It has these top-level messages:
	AddRoleToGroupOnDomainRequest
	AddRoleToGroupOnProjectRequest
	AddRoleToUserOnDomainRequest
	AddRoleToUserOnProjectRequest
	AddUserToGroupRequest
	ChangePasswordRequest
	CheckRoleInGroupOnDomainRequest
	CheckRoleInGroupOnProjectRequest
	CheckRoleInUserOnDomainRequest
	CheckRoleInUserOnProjectRequest
	CheckUserInGroupRequest
	CreateDomainRequest
	CreateDomainResponse
	CreateGroupRequest
	CreateGroupResponse
	CreateProjectRequest
	CreateProjectResponse
	CreateRoleRequest
	CreateRoleResponse
	CreateUserRequest
	CreateUserResponse
	DeleteDomainRequest
	DeleteGroupRequest
	DeleteProjectRequest
	DeleteRoleRequest
	DeleteUserRequest
	Domain
	GetDomainRequest
	GetDomainResponse
	GetGroupRequest
	GetGroupResponse
	GetProjectRequest
	GetProjectResponse
	GetRoleRequest
	GetRoleResponse
	GetUserRequest
	GetUserResponse
	Group
	ListDomainsRequest
	ListDomainsResponse
	ListGroupsRequest
	ListGroupsResponse
	ListGroupsForUserRequest
	ListGroupsForUserResponse
	ListProjectsRequest
	ListProjectsResponse
	ListProjectsForUserRequest
	ListProjectsForUserResponse
	ListRoleInGroupOnDomainRequest
	ListRoleInGroupOnDomainResponse
	ListRolesRequest
	ListRolesResponse
	ListRolesForGroupOnDomainRequest
	ListRolesForGroupOnDomainResponse
	ListRolesForGroupOnProjectRequest
	ListRolesForGroupOnProjectResponse
	ListRolesForUserOnDomainRequest
	ListRolesForUserOnDomainResponse
	ListRolesForUserOnProjectRequest
	ListRolesForUserOnProjectResponse
	ListUsersRequest
	ListUsersResponse
	ListUsersInGroupRequest
	ListUsersInGroupResponse
	PatchDomainRequest
	PatchDomainResponse
	PatchGroupRequest
	PatchGroupResponse
	PatchProjectRequest
	PatchProjectResponse
	PatchRoleRequest
	PatchRoleResponse
	PatchUserRequest
	PatchUserResponse
	Project
	RemoveRoleFromGroupOnDomainRequest
	RemoveRoleFromGroupOnProjectRequest
	RemoveRoleFromUserOnDomainRequest
	RemoveRoleFromUserOnProjectRequest
	RemoveUserFromGroupRequest
	Role
	User
*/
package identity

import fmt "fmt"
import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *AddRoleToGroupOnDomainRequest) Validate() error {
	if nil == this.DomainId {
		return go_proto_validators.FieldError("DomainId", fmt.Errorf("message must exist"))
	}
	if this.DomainId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.DomainId); err != nil {
			return go_proto_validators.FieldError("DomainId", err)
		}
	}
	if nil == this.GroupId {
		return go_proto_validators.FieldError("GroupId", fmt.Errorf("message must exist"))
	}
	if this.GroupId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.GroupId); err != nil {
			return go_proto_validators.FieldError("GroupId", err)
		}
	}
	if nil == this.RoleId {
		return go_proto_validators.FieldError("RoleId", fmt.Errorf("message must exist"))
	}
	if this.RoleId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.RoleId); err != nil {
			return go_proto_validators.FieldError("RoleId", err)
		}
	}
	return nil
}
