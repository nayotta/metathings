// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_roles_for_group_on_project.proto

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

type ListRolesForGroupOnProjectRequest struct {
	ProjectId *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	GroupId   *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
}

func (m *ListRolesForGroupOnProjectRequest) Reset()         { *m = ListRolesForGroupOnProjectRequest{} }
func (m *ListRolesForGroupOnProjectRequest) String() string { return proto.CompactTextString(m) }
func (*ListRolesForGroupOnProjectRequest) ProtoMessage()    {}
func (*ListRolesForGroupOnProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor41, []int{0}
}

func (m *ListRolesForGroupOnProjectRequest) GetProjectId() *google_protobuf.StringValue {
	if m != nil {
		return m.ProjectId
	}
	return nil
}

func (m *ListRolesForGroupOnProjectRequest) GetGroupId() *google_protobuf.StringValue {
	if m != nil {
		return m.GroupId
	}
	return nil
}

type ListRolesForGroupOnProjectResponse struct {
	Roles []*Role `protobuf:"bytes,1,rep,name=roles" json:"roles,omitempty"`
}

func (m *ListRolesForGroupOnProjectResponse) Reset()         { *m = ListRolesForGroupOnProjectResponse{} }
func (m *ListRolesForGroupOnProjectResponse) String() string { return proto.CompactTextString(m) }
func (*ListRolesForGroupOnProjectResponse) ProtoMessage()    {}
func (*ListRolesForGroupOnProjectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor41, []int{1}
}

func (m *ListRolesForGroupOnProjectResponse) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRolesForGroupOnProjectRequest)(nil), "ai.metathings.service.identity.ListRolesForGroupOnProjectRequest")
	proto.RegisterType((*ListRolesForGroupOnProjectResponse)(nil), "ai.metathings.service.identity.ListRolesForGroupOnProjectResponse")
}

func init() { proto.RegisterFile("list_roles_for_group_on_project.proto", fileDescriptor41) }

var fileDescriptor41 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd0, 0xc1, 0x4a, 0x33, 0x31,
	0x10, 0xc0, 0x71, 0xb6, 0x1f, 0x5f, 0xd5, 0xf4, 0xb6, 0xa7, 0x52, 0xa4, 0xd6, 0xa2, 0xd0, 0x4b,
	0xb3, 0x50, 0xc1, 0x83, 0x17, 0x41, 0x41, 0x29, 0x08, 0xca, 0x0a, 0x5e, 0xd7, 0xb4, 0x99, 0xa6,
	0xa3, 0xdb, 0x4c, 0x4c, 0x66, 0x5b, 0x7c, 0x24, 0x9f, 0x4a, 0xf0, 0x49, 0x64, 0x37, 0xd5, 0x63,
	0xc1, 0xdb, 0x40, 0xf2, 0x1f, 0x7e, 0x8c, 0x38, 0x2d, 0x31, 0x70, 0xe1, 0xa9, 0x84, 0x50, 0x2c,
	0xc8, 0x17, 0xc6, 0x53, 0xe5, 0x0a, 0xb2, 0x85, 0xf3, 0xf4, 0x02, 0x73, 0x96, 0xce, 0x13, 0x53,
	0xda, 0x57, 0x28, 0x57, 0xc0, 0x8a, 0x97, 0x68, 0x4d, 0x90, 0x01, 0xfc, 0x1a, 0xe7, 0x20, 0x51,
	0x83, 0x65, 0xe4, 0xf7, 0x5e, 0xdf, 0x10, 0x99, 0x12, 0xb2, 0xe6, 0xf7, 0xac, 0x5a, 0x64, 0x1b,
	0xaf, 0x9c, 0x03, 0x1f, 0x62, 0xdf, 0x3b, 0x37, 0xc8, 0xcb, 0x6a, 0x26, 0xe7, 0xb4, 0xca, 0x56,
	0x1b, 0xe4, 0x57, 0xda, 0x64, 0x86, 0xc6, 0xcd, 0xe3, 0x78, 0xad, 0x4a, 0xd4, 0x8a, 0xc9, 0x87,
	0xec, 0x77, 0xdc, 0x76, 0xa2, 0x96, 0xc5, 0x79, 0xf8, 0x91, 0x88, 0xe3, 0x3b, 0x0c, 0x9c, 0xd7,
	0xd8, 0x1b, 0xf2, 0xb7, 0x35, 0xf5, 0xde, 0x3e, 0x44, 0x68, 0x0e, 0x6f, 0x15, 0x04, 0x4e, 0xaf,
	0x85, 0xd8, 0xd2, 0x0b, 0xd4, 0xdd, 0x64, 0x90, 0x8c, 0x3a, 0x93, 0x43, 0x19, 0x79, 0xf2, 0x87,
	0x27, 0x1f, 0xd9, 0xa3, 0x35, 0x4f, 0xaa, 0xac, 0xe0, 0xaa, 0xfd, 0xf5, 0x79, 0xd4, 0x1a, 0x24,
	0xf9, 0xc1, 0xb6, 0x9b, 0xea, 0xf4, 0x52, 0xec, 0xc7, 0x43, 0xa0, 0xee, 0xb6, 0xfe, 0xb0, 0x62,
	0xaf, 0xa9, 0xa6, 0x7a, 0xf8, 0x2c, 0x86, 0xbb, 0xa8, 0xc1, 0x91, 0x0d, 0x90, 0x5e, 0x88, 0xff,
	0xcd, 0xe5, 0xbb, 0xc9, 0xe0, 0xdf, 0xa8, 0x33, 0x39, 0x91, 0xbb, 0xaf, 0x2c, 0xeb, 0x75, 0x79,
	0x4c, 0x66, 0xed, 0x06, 0x72, 0xf6, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x8c, 0x0d, 0xa6, 0xc1,
	0x01, 0x00, 0x00,
}
