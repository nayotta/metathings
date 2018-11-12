// Code generated by protoc-gen-go. DO NOT EDIT.
// source: check_role_in_group_on_project.proto

package identityd

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CheckRoleInGroupOnProjectRequest struct {
	ProjectId            *wrappers.StringValue `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	GroupId              *wrappers.StringValue `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	RoleId               *wrappers.StringValue `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CheckRoleInGroupOnProjectRequest) Reset()         { *m = CheckRoleInGroupOnProjectRequest{} }
func (m *CheckRoleInGroupOnProjectRequest) String() string { return proto.CompactTextString(m) }
func (*CheckRoleInGroupOnProjectRequest) ProtoMessage()    {}
func (*CheckRoleInGroupOnProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_43beeb027ce84c64, []int{0}
}

func (m *CheckRoleInGroupOnProjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckRoleInGroupOnProjectRequest.Unmarshal(m, b)
}
func (m *CheckRoleInGroupOnProjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckRoleInGroupOnProjectRequest.Marshal(b, m, deterministic)
}
func (m *CheckRoleInGroupOnProjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckRoleInGroupOnProjectRequest.Merge(m, src)
}
func (m *CheckRoleInGroupOnProjectRequest) XXX_Size() int {
	return xxx_messageInfo_CheckRoleInGroupOnProjectRequest.Size(m)
}
func (m *CheckRoleInGroupOnProjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckRoleInGroupOnProjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckRoleInGroupOnProjectRequest proto.InternalMessageInfo

func (m *CheckRoleInGroupOnProjectRequest) GetProjectId() *wrappers.StringValue {
	if m != nil {
		return m.ProjectId
	}
	return nil
}

func (m *CheckRoleInGroupOnProjectRequest) GetGroupId() *wrappers.StringValue {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func (m *CheckRoleInGroupOnProjectRequest) GetRoleId() *wrappers.StringValue {
	if m != nil {
		return m.RoleId
	}
	return nil
}

func init() {
	proto.RegisterType((*CheckRoleInGroupOnProjectRequest)(nil), "ai.metathings.service.identityd.CheckRoleInGroupOnProjectRequest")
}

func init() {
	proto.RegisterFile("check_role_in_group_on_project.proto", fileDescriptor_43beeb027ce84c64)
}

var fileDescriptor_43beeb027ce84c64 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x85, 0x56, 0xd7, 0x5b, 0x4e, 0xa5, 0x88, 0x2d, 0xe2, 0xc1, 0x4b, 0x37, 0xa0,
	0xe0, 0x4d, 0x04, 0x7b, 0x90, 0x9c, 0x94, 0x08, 0x5e, 0x97, 0x4d, 0x76, 0xdc, 0x8c, 0x4d, 0x76,
	0xe2, 0x66, 0xd2, 0xe0, 0xd3, 0x0a, 0x5e, 0x7c, 0x0d, 0x49, 0xb6, 0x7a, 0xef, 0x6d, 0x60, 0xfe,
	0xef, 0xe7, 0x9b, 0x11, 0x97, 0x45, 0x09, 0xc5, 0x56, 0x79, 0xaa, 0x40, 0xa1, 0x53, 0xd6, 0x53,
	0xd7, 0x28, 0x72, 0xaa, 0xf1, 0xf4, 0x0e, 0x05, 0xcb, 0xc6, 0x13, 0x53, 0xbc, 0xd4, 0x28, 0x6b,
	0x60, 0xcd, 0x25, 0x3a, 0xdb, 0xca, 0x16, 0xfc, 0x0e, 0x0b, 0x90, 0x68, 0xc0, 0x31, 0xf2, 0xa7,
	0x59, 0x9c, 0x5b, 0x22, 0x5b, 0x41, 0x32, 0xc6, 0xf3, 0xee, 0x2d, 0xe9, 0xbd, 0x6e, 0x1a, 0xf0,
	0x6d, 0x28, 0x58, 0xdc, 0x5a, 0xe4, 0xb2, 0xcb, 0x65, 0x41, 0x75, 0x52, 0xf7, 0xc8, 0x5b, 0xea,
	0x13, 0x4b, 0xeb, 0x71, 0xb9, 0xde, 0xe9, 0x0a, 0x8d, 0x66, 0xf2, 0x6d, 0xf2, 0x3f, 0x06, 0xee,
	0xe2, 0x27, 0x12, 0xab, 0xcd, 0x60, 0x98, 0x51, 0x05, 0xa9, 0x7b, 0x1c, 0xf4, 0x9e, 0xdc, 0x73,
	0x90, 0xcb, 0xe0, 0xa3, 0x83, 0x96, 0xe3, 0x8d, 0x10, 0x7b, 0x5d, 0x85, 0x66, 0x1e, 0xad, 0xa2,
	0xab, 0xd3, 0xeb, 0x33, 0x19, 0x8c, 0xe4, 0x9f, 0x91, 0x7c, 0x61, 0x8f, 0xce, 0xbe, 0xea, 0xaa,
	0x83, 0x87, 0xe9, 0xf7, 0xd7, 0x72, 0xb2, 0x8a, 0xb2, 0x93, 0x3d, 0x97, 0x9a, 0xf8, 0x5e, 0x1c,
	0x87, 0xe3, 0xd1, 0xcc, 0x27, 0x07, 0x54, 0xcc, 0x46, 0x2a, 0x35, 0xf1, 0x9d, 0x98, 0x85, 0x2f,
	0x9a, 0xf9, 0xd1, 0x01, 0xfc, 0x74, 0x80, 0x52, 0x93, 0x4f, 0xc7, 0xd4, 0xcd, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x48, 0xca, 0x60, 0x72, 0x91, 0x01, 0x00, 0x00,
}
