// Code generated by protoc-gen-go. DO NOT EDIT.
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
	create_region.proto
	create_role.proto
	create_user.proto
	delete_domain.proto
	delete_group.proto
	delete_project.proto
	delete_region.proto
	delete_role.proto
	delete_user.proto
	domain.proto
	get_domain.proto
	get_group.proto
	get_project.proto
	get_region.proto
	get_role.proto
	get_user.proto
	group.proto
	list_domains.proto
	list_groups.proto
	list_groups_for_user.proto
	list_projects.proto
	list_projects_for_user.proto
	list_regions.proto
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
	patch_region.proto
	patch_role.proto
	patch_user.proto
	project.proto
	region.proto
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
	CreateRegionRequest
	CreateRegionResponse
	CreateRoleRequest
	CreateRoleResponse
	CreateUserRequest
	CreateUserResponse
	DeleteDomainRequest
	DeleteGroupRequest
	DeleteProjectRequest
	DeleteRegionRequest
	DeleteRoleRequest
	DeleteUserRequest
	Domain
	GetDomainRequest
	GetDomainResponse
	GetGroupRequest
	GetGroupResponse
	GetProjectRequest
	GetProjectResponse
	GetRegionRequest
	GetRegionResponse
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
	ListRegionsRequest
	ListRegionsResponse
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
	PatchRegionRequest
	PatchRegionResponse
	PatchRoleRequest
	PatchRoleResponse
	PatchUserRequest
	PatchUserResponse
	Project
	Region
	RemoveRoleFromGroupOnDomainRequest
	RemoveRoleFromGroupOnProjectRequest
	RemoveRoleFromUserOnDomainRequest
	RemoveRoleFromUserOnProjectRequest
	RemoveUserFromGroupRequest
	Role
	User
*/
package identity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddRoleToGroupOnDomainRequest struct {
	DomainId *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	GroupId  *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	RoleId   *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=role_id,json=roleId" json:"role_id,omitempty"`
}

func (m *AddRoleToGroupOnDomainRequest) Reset()                    { *m = AddRoleToGroupOnDomainRequest{} }
func (m *AddRoleToGroupOnDomainRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRoleToGroupOnDomainRequest) ProtoMessage()               {}
func (*AddRoleToGroupOnDomainRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddRoleToGroupOnDomainRequest) GetDomainId() *google_protobuf.StringValue {
	if m != nil {
		return m.DomainId
	}
	return nil
}

func (m *AddRoleToGroupOnDomainRequest) GetGroupId() *google_protobuf.StringValue {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func (m *AddRoleToGroupOnDomainRequest) GetRoleId() *google_protobuf.StringValue {
	if m != nil {
		return m.RoleId
	}
	return nil
}

func init() {
	proto.RegisterType((*AddRoleToGroupOnDomainRequest)(nil), "ai.metathings.service.identity.AddRoleToGroupOnDomainRequest")
}

func init() { proto.RegisterFile("add_role_to_group_on_domain.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0x49, 0x3f, 0x68, 0xfb, 0xc5, 0x5b, 0x4e, 0xa5, 0x68, 0xad, 0x9e, 0xbc, 0x74, 0x03,
	0x0a, 0xde, 0x44, 0x2a, 0x82, 0xe4, 0x24, 0x44, 0xf1, 0xba, 0x6c, 0x3a, 0xe3, 0x76, 0x30, 0xd9,
	0x89, 0x9b, 0x49, 0x83, 0xbf, 0x56, 0xe8, 0x2f, 0x91, 0xee, 0xa2, 0xf7, 0xde, 0x06, 0xe6, 0x7d,
	0x86, 0x67, 0xde, 0xf4, 0xc2, 0x00, 0x68, 0xcf, 0x35, 0x6a, 0x61, 0x6d, 0x3d, 0xf7, 0xad, 0x66,
	0xa7, 0x81, 0x1b, 0x43, 0x4e, 0xb5, 0x9e, 0x85, 0xb3, 0x85, 0x21, 0xd5, 0xa0, 0x18, 0xd9, 0x92,
	0xb3, 0x9d, 0xea, 0xd0, 0xef, 0x68, 0x83, 0x8a, 0x00, 0x9d, 0x90, 0x7c, 0xcd, 0x17, 0x96, 0xd9,
	0xd6, 0x98, 0x87, 0x74, 0xd5, 0xbf, 0xe7, 0x83, 0x37, 0x6d, 0x8b, 0xbe, 0x8b, 0xfc, 0xfc, 0xd6,
	0x92, 0x6c, 0xfb, 0x4a, 0x6d, 0xb8, 0xc9, 0x9b, 0x81, 0xe4, 0x83, 0x87, 0xdc, 0xf2, 0x2a, 0x2c,
	0x57, 0x3b, 0x53, 0x13, 0x18, 0x61, 0xdf, 0xe5, 0x7f, 0x63, 0xe4, 0x2e, 0xf7, 0x49, 0x7a, 0xb6,
	0x06, 0x28, 0xb9, 0xc6, 0x57, 0x7e, 0x3a, 0xa8, 0x3d, 0xbb, 0xc7, 0x20, 0x56, 0xe2, 0x67, 0x8f,
	0x9d, 0x64, 0xeb, 0xf4, 0x7f, 0x34, 0xd5, 0x04, 0xb3, 0x64, 0x99, 0x5c, 0x9d, 0x5c, 0x9f, 0xaa,
	0x68, 0xa3, 0x7e, 0x6d, 0xd4, 0x8b, 0x78, 0x72, 0xf6, 0xcd, 0xd4, 0x3d, 0x3e, 0x8c, 0xf7, 0xdf,
	0xe7, 0xa3, 0x65, 0x52, 0x4e, 0x23, 0x56, 0x40, 0x76, 0x9f, 0x4e, 0xe3, 0xd7, 0x04, 0xb3, 0xd1,
	0x11, 0x17, 0x26, 0x81, 0x2a, 0x20, 0xbb, 0x4b, 0x27, 0xa1, 0x3e, 0x82, 0xd9, 0xbf, 0x23, 0xf8,
	0xf1, 0x01, 0x2a, 0xa0, 0x1a, 0x87, 0xd4, 0xcd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x4f,
	0xa2, 0xe5, 0x88, 0x01, 0x00, 0x00,
}
