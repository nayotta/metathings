// Code generated by protoc-gen-go. DO NOT EDIT.
// source: add_role_to_group_on_domain.proto

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

type AddRoleToGroupOnDomainRequest struct {
	DomainId             *wrappers.StringValue `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	GroupId              *wrappers.StringValue `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	RoleId               *wrappers.StringValue `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AddRoleToGroupOnDomainRequest) Reset()         { *m = AddRoleToGroupOnDomainRequest{} }
func (m *AddRoleToGroupOnDomainRequest) String() string { return proto.CompactTextString(m) }
func (*AddRoleToGroupOnDomainRequest) ProtoMessage()    {}
func (*AddRoleToGroupOnDomainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6af50eccb3e006a0, []int{0}
}

func (m *AddRoleToGroupOnDomainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRoleToGroupOnDomainRequest.Unmarshal(m, b)
}
func (m *AddRoleToGroupOnDomainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRoleToGroupOnDomainRequest.Marshal(b, m, deterministic)
}
func (m *AddRoleToGroupOnDomainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRoleToGroupOnDomainRequest.Merge(m, src)
}
func (m *AddRoleToGroupOnDomainRequest) XXX_Size() int {
	return xxx_messageInfo_AddRoleToGroupOnDomainRequest.Size(m)
}
func (m *AddRoleToGroupOnDomainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRoleToGroupOnDomainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRoleToGroupOnDomainRequest proto.InternalMessageInfo

func (m *AddRoleToGroupOnDomainRequest) GetDomainId() *wrappers.StringValue {
	if m != nil {
		return m.DomainId
	}
	return nil
}

func (m *AddRoleToGroupOnDomainRequest) GetGroupId() *wrappers.StringValue {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func (m *AddRoleToGroupOnDomainRequest) GetRoleId() *wrappers.StringValue {
	if m != nil {
		return m.RoleId
	}
	return nil
}

func init() {
	proto.RegisterType((*AddRoleToGroupOnDomainRequest)(nil), "ai.metathings.service.identityd.AddRoleToGroupOnDomainRequest")
}

func init() { proto.RegisterFile("add_role_to_group_on_domain.proto", fileDescriptor_6af50eccb3e006a0) }

var fileDescriptor_6af50eccb3e006a0 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0x49, 0x3f, 0x68, 0xfb, 0xc5, 0x5b, 0x4e, 0xa5, 0xa8, 0xad, 0x9e, 0xbc, 0x74, 0x03,
	0x0a, 0xde, 0x44, 0x2a, 0x82, 0xe4, 0x24, 0x44, 0xf1, 0xba, 0x6c, 0x3a, 0xe3, 0x76, 0x30, 0xd9,
	0x89, 0x9b, 0x49, 0x83, 0xbf, 0x56, 0xe8, 0x2f, 0x91, 0xee, 0xa2, 0xf7, 0xde, 0x06, 0xe6, 0x7d,
	0x5e, 0x9e, 0x99, 0xf4, 0xc2, 0x00, 0x68, 0xcf, 0x35, 0x6a, 0x61, 0x6d, 0x3d, 0xf7, 0xad, 0x66,
	0xa7, 0x81, 0x1b, 0x43, 0x4e, 0xb5, 0x9e, 0x85, 0xb3, 0x85, 0x21, 0xd5, 0xa0, 0x18, 0xd9, 0x92,
	0xb3, 0x9d, 0xea, 0xd0, 0xef, 0x68, 0x83, 0x8a, 0x00, 0x9d, 0x90, 0x7c, 0xc1, 0xfc, 0xdc, 0x32,
	0xdb, 0x1a, 0xf3, 0x10, 0xaf, 0xfa, 0xf7, 0x7c, 0xf0, 0xa6, 0x6d, 0xd1, 0x77, 0xb1, 0x60, 0x7e,
	0x6b, 0x49, 0xb6, 0x7d, 0xa5, 0x36, 0xdc, 0xe4, 0xcd, 0x40, 0xf2, 0xc1, 0x43, 0x6e, 0x79, 0x15,
	0x96, 0xab, 0x9d, 0xa9, 0x09, 0x8c, 0xb0, 0xef, 0xf2, 0xbf, 0x31, 0x72, 0x97, 0xfb, 0x24, 0x3d,
	0x5b, 0x03, 0x94, 0x5c, 0xe3, 0x2b, 0x3f, 0x1d, 0xdc, 0x9e, 0xdd, 0x63, 0x30, 0x2b, 0xf1, 0xb3,
	0xc7, 0x4e, 0xb2, 0x75, 0xfa, 0x3f, 0xaa, 0x6a, 0x82, 0x59, 0xb2, 0x4c, 0xae, 0x4e, 0xae, 0x4f,
	0x55, 0xb4, 0x51, 0xbf, 0x36, 0xea, 0x45, 0x3c, 0x39, 0xfb, 0x66, 0xea, 0x1e, 0x1f, 0xc6, 0xfb,
	0xef, 0xc5, 0x68, 0x99, 0x94, 0xd3, 0x88, 0x15, 0x90, 0xdd, 0xa7, 0xd3, 0x78, 0x36, 0xc1, 0x6c,
	0x74, 0x44, 0xc3, 0x24, 0x50, 0x05, 0x64, 0x77, 0xe9, 0x24, 0xfc, 0x8f, 0x60, 0xf6, 0xef, 0x08,
	0x7e, 0x7c, 0x80, 0x0a, 0xa8, 0xc6, 0x21, 0x75, 0xf3, 0x13, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xe4,
	0x3c, 0x93, 0x89, 0x01, 0x00, 0x00,
}
