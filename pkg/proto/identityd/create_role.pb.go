// Code generated by protoc-gen-go. DO NOT EDIT.
// source: create_role.proto

package identityd

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type CreateRoleRequest struct {
	Name                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	DomainId             *wrappers.StringValue `protobuf:"bytes,2,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateRoleRequest) Reset()         { *m = CreateRoleRequest{} }
func (m *CreateRoleRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRoleRequest) ProtoMessage()    {}
func (*CreateRoleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_create_role_8a58f60f0c53edf3, []int{0}
}
func (m *CreateRoleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoleRequest.Unmarshal(m, b)
}
func (m *CreateRoleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoleRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRoleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoleRequest.Merge(dst, src)
}
func (m *CreateRoleRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRoleRequest.Size(m)
}
func (m *CreateRoleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoleRequest proto.InternalMessageInfo

func (m *CreateRoleRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CreateRoleRequest) GetDomainId() *wrappers.StringValue {
	if m != nil {
		return m.DomainId
	}
	return nil
}

type CreateRoleResponse struct {
	Role                 *Role    `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoleResponse) Reset()         { *m = CreateRoleResponse{} }
func (m *CreateRoleResponse) String() string { return proto.CompactTextString(m) }
func (*CreateRoleResponse) ProtoMessage()    {}
func (*CreateRoleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_create_role_8a58f60f0c53edf3, []int{1}
}
func (m *CreateRoleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoleResponse.Unmarshal(m, b)
}
func (m *CreateRoleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoleResponse.Marshal(b, m, deterministic)
}
func (dst *CreateRoleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoleResponse.Merge(dst, src)
}
func (m *CreateRoleResponse) XXX_Size() int {
	return xxx_messageInfo_CreateRoleResponse.Size(m)
}
func (m *CreateRoleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoleResponse proto.InternalMessageInfo

func (m *CreateRoleResponse) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRoleRequest)(nil), "ai.metathings.service.identityd.CreateRoleRequest")
	proto.RegisterType((*CreateRoleResponse)(nil), "ai.metathings.service.identityd.CreateRoleResponse")
}

func init() { proto.RegisterFile("create_role.proto", fileDescriptor_create_role_8a58f60f0c53edf3) }

var fileDescriptor_create_role_8a58f60f0c53edf3 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x49, 0x29, 0x45, 0xd7, 0x53, 0x73, 0x2a, 0x45, 0x6c, 0x29, 0x08, 0x5e, 0x3a, 0x01,
	0x05, 0xd1, 0xab, 0x9e, 0x3c, 0x09, 0x11, 0xbc, 0x96, 0x6d, 0x76, 0xdc, 0x0e, 0x66, 0x77, 0xe2,
	0xee, 0xa4, 0xc1, 0x37, 0xf0, 0x2d, 0x05, 0x9f, 0x44, 0xba, 0xa9, 0x7f, 0x6e, 0xbd, 0x0d, 0xcc,
	0xf7, 0xfb, 0xe6, 0xc7, 0xa8, 0x71, 0x15, 0x50, 0x0b, 0xae, 0x02, 0xd7, 0x08, 0x4d, 0x60, 0xe1,
	0x7c, 0xa6, 0x09, 0x1c, 0x8a, 0x96, 0x0d, 0x79, 0x1b, 0x21, 0x62, 0xd8, 0x52, 0x85, 0x40, 0x06,
	0xbd, 0x90, 0xbc, 0x9b, 0xe9, 0x99, 0x65, 0xb6, 0x35, 0x16, 0x29, 0xbe, 0x6e, 0x5f, 0x8a, 0x2e,
	0xe8, 0xa6, 0xc1, 0x10, 0xfb, 0x82, 0xe9, 0xb5, 0x25, 0xd9, 0xb4, 0x6b, 0xa8, 0xd8, 0x15, 0xae,
	0x23, 0x79, 0xe5, 0xae, 0xb0, 0xbc, 0x4c, 0xcb, 0xe5, 0x56, 0xd7, 0x64, 0xb4, 0x70, 0x88, 0xc5,
	0xef, 0xb8, 0xe7, 0xd4, 0x9f, 0xc4, 0xe2, 0x23, 0x53, 0xe3, 0xfb, 0xa4, 0x56, 0x72, 0x8d, 0x25,
	0xbe, 0xb5, 0x18, 0x25, 0xbf, 0x51, 0x43, 0xaf, 0x1d, 0x4e, 0xb2, 0x79, 0x76, 0x71, 0x72, 0x79,
	0x0a, 0xbd, 0x08, 0xfc, 0x88, 0xc0, 0x93, 0x04, 0xf2, 0xf6, 0x59, 0xd7, 0x2d, 0xde, 0x8d, 0xbe,
	0x3e, 0x67, 0x83, 0x79, 0x56, 0x26, 0x22, 0xbf, 0x55, 0xc7, 0x86, 0x9d, 0x26, 0xbf, 0x22, 0x33,
	0x19, 0x1c, 0xc6, 0xcb, 0xa3, 0x3e, 0xfe, 0x60, 0x16, 0x8f, 0x2a, 0xff, 0x6f, 0x12, 0x1b, 0xf6,
	0x71, 0x57, 0x38, 0xdc, 0xe9, 0xee, 0x55, 0xce, 0xe1, 0xc0, 0xd3, 0x20, 0xc1, 0x09, 0x59, 0x8f,
	0xd2, 0xc1, 0xab, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x4c, 0x9a, 0xa0, 0x7c, 0x01, 0x00,
	0x00,
}