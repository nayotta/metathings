// Code generated by protoc-gen-go. DO NOT EDIT.
// source: add_role_to_group.proto

package identityd2

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AddRoleToGroupRequest struct {
	Group                *OpGroup `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Role                 *OpRole  `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRoleToGroupRequest) Reset()         { *m = AddRoleToGroupRequest{} }
func (m *AddRoleToGroupRequest) String() string { return proto.CompactTextString(m) }
func (*AddRoleToGroupRequest) ProtoMessage()    {}
func (*AddRoleToGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6fcf59c0d5979f4, []int{0}
}

func (m *AddRoleToGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRoleToGroupRequest.Unmarshal(m, b)
}
func (m *AddRoleToGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRoleToGroupRequest.Marshal(b, m, deterministic)
}
func (m *AddRoleToGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRoleToGroupRequest.Merge(m, src)
}
func (m *AddRoleToGroupRequest) XXX_Size() int {
	return xxx_messageInfo_AddRoleToGroupRequest.Size(m)
}
func (m *AddRoleToGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRoleToGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRoleToGroupRequest proto.InternalMessageInfo

func (m *AddRoleToGroupRequest) GetGroup() *OpGroup {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *AddRoleToGroupRequest) GetRole() *OpRole {
	if m != nil {
		return m.Role
	}
	return nil
}

func init() {
	proto.RegisterType((*AddRoleToGroupRequest)(nil), "ai.metathings.service.identityd2.AddRoleToGroupRequest")
}

func init() { proto.RegisterFile("add_role_to_group.proto", fileDescriptor_d6fcf59c0d5979f4) }

var fileDescriptor_d6fcf59c0d5979f4 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x4c, 0x49, 0x89,
	0x2f, 0xca, 0xcf, 0x49, 0x8d, 0x2f, 0xc9, 0x8f, 0x4f, 0x2f, 0xca, 0x2f, 0x2d, 0xd0, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x52, 0x48, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc,
	0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0xcb, 0x4c, 0x49, 0xcd, 0x2b,
	0xc9, 0x2c, 0xa9, 0x4c, 0x31, 0x92, 0x12, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5,
	0x87, 0x31, 0x20, 0x5a, 0xa5, 0xb8, 0x73, 0xf3, 0x53, 0x52, 0x73, 0x20, 0x1c, 0xa5, 0x55, 0x8c,
	0x5c, 0xa2, 0x8e, 0x29, 0x29, 0x41, 0xf9, 0x39, 0xa9, 0x21, 0xf9, 0xee, 0x20, 0x0b, 0x82, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x3c, 0xb9, 0x58, 0xc1, 0x16, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x1b, 0x69, 0xea, 0x11, 0xb2, 0x51, 0xcf, 0xbf, 0x00, 0x6c, 0x80, 0x13, 0xc7, 0x2f, 0x27,
	0xd6, 0x2e, 0x46, 0x26, 0x01, 0xc6, 0x20, 0x88, 0x09, 0x42, 0x6e, 0x5c, 0x2c, 0x20, 0x3f, 0x48,
	0x30, 0x81, 0x4d, 0xd2, 0x20, 0xc6, 0x24, 0x90, 0x83, 0x90, 0x0c, 0x02, 0xeb, 0x4f, 0x62, 0x03,
	0xbb, 0xd9, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x26, 0x39, 0x18, 0xef, 0x16, 0x01, 0x00, 0x00,
}
