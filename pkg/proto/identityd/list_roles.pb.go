// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_roles.proto

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

type ListRolesRequest struct {
	Name                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DomainId             *wrappers.StringValue `protobuf:"bytes,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListRolesRequest) Reset()         { *m = ListRolesRequest{} }
func (m *ListRolesRequest) String() string { return proto.CompactTextString(m) }
func (*ListRolesRequest) ProtoMessage()    {}
func (*ListRolesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_39632ff8fb9178eb, []int{0}
}

func (m *ListRolesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRolesRequest.Unmarshal(m, b)
}
func (m *ListRolesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRolesRequest.Marshal(b, m, deterministic)
}
func (m *ListRolesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRolesRequest.Merge(m, src)
}
func (m *ListRolesRequest) XXX_Size() int {
	return xxx_messageInfo_ListRolesRequest.Size(m)
}
func (m *ListRolesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRolesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRolesRequest proto.InternalMessageInfo

func (m *ListRolesRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ListRolesRequest) GetDomainId() *wrappers.StringValue {
	if m != nil {
		return m.DomainId
	}
	return nil
}

type ListRolesResponse struct {
	Roles                []*Role  `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRolesResponse) Reset()         { *m = ListRolesResponse{} }
func (m *ListRolesResponse) String() string { return proto.CompactTextString(m) }
func (*ListRolesResponse) ProtoMessage()    {}
func (*ListRolesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_39632ff8fb9178eb, []int{1}
}

func (m *ListRolesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRolesResponse.Unmarshal(m, b)
}
func (m *ListRolesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRolesResponse.Marshal(b, m, deterministic)
}
func (m *ListRolesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRolesResponse.Merge(m, src)
}
func (m *ListRolesResponse) XXX_Size() int {
	return xxx_messageInfo_ListRolesResponse.Size(m)
}
func (m *ListRolesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRolesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRolesResponse proto.InternalMessageInfo

func (m *ListRolesResponse) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRolesRequest)(nil), "ai.metathings.service.identityd.ListRolesRequest")
	proto.RegisterType((*ListRolesResponse)(nil), "ai.metathings.service.identityd.ListRolesResponse")
}

func init() { proto.RegisterFile("list_roles.proto", fileDescriptor_39632ff8fb9178eb) }

var fileDescriptor_39632ff8fb9178eb = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x89, 0xff, 0xd0, 0xed, 0xa5, 0xe6, 0x14, 0x8a, 0x68, 0x29, 0x08, 0xbd, 0x74, 0x22,
	0x15, 0x04, 0xf1, 0x13, 0x08, 0x1e, 0x24, 0x82, 0xd7, 0xb2, 0xe9, 0x8e, 0xdb, 0xc1, 0x64, 0x27,
	0xee, 0x4e, 0x1a, 0x3c, 0xf9, 0xd5, 0x25, 0x9b, 0x2a, 0xde, 0xec, 0xed, 0xc1, 0xbe, 0xdf, 0xce,
	0x8f, 0xa7, 0xc6, 0x15, 0x05, 0x59, 0x79, 0xae, 0x30, 0x40, 0xe3, 0x59, 0x38, 0xbd, 0xd2, 0x04,
	0x35, 0x8a, 0x96, 0x0d, 0x39, 0x1b, 0x20, 0xa0, 0xdf, 0xd2, 0x1a, 0x81, 0x0c, 0x3a, 0x21, 0xf9,
	0x34, 0x93, 0x4b, 0xcb, 0x6c, 0x2b, 0xcc, 0x63, 0xbd, 0x6c, 0xdf, 0xf2, 0xce, 0xeb, 0xa6, 0x41,
	0xbf, 0xfb, 0x60, 0x72, 0x67, 0x49, 0x36, 0x6d, 0x09, 0x6b, 0xae, 0xf3, 0xba, 0x23, 0x79, 0xe7,
	0x2e, 0xb7, 0xbc, 0x88, 0x8f, 0x8b, 0xad, 0xae, 0xc8, 0x68, 0x61, 0x1f, 0xf2, 0xdf, 0xb8, 0xe3,
	0x54, 0x6f, 0x31, 0xe4, 0xd9, 0x97, 0x1a, 0x3f, 0x51, 0x90, 0xa2, 0xf7, 0x2a, 0xf0, 0xa3, 0xc5,
	0x20, 0xe9, 0x8d, 0x3a, 0x72, 0xba, 0xc6, 0x2c, 0x99, 0x26, 0xf3, 0xd1, 0xf2, 0x02, 0x06, 0x0d,
	0xf8, 0xd1, 0x80, 0x17, 0xf1, 0xe4, 0xec, 0xab, 0xae, 0x5a, 0x2c, 0x62, 0x33, 0xbd, 0x57, 0x67,
	0x86, 0x6b, 0x4d, 0x6e, 0x45, 0x26, 0x3b, 0xd8, 0x03, 0x3b, 0x1d, 0xea, 0x8f, 0x66, 0xf6, 0xac,
	0xce, 0xff, 0x08, 0x84, 0x86, 0x5d, 0xc0, 0xf4, 0x41, 0x1d, 0xc7, 0xa5, 0xb2, 0x64, 0x7a, 0x38,
	0x1f, 0x2d, 0xaf, 0xe1, 0x9f, 0xa9, 0xa0, 0xc7, 0x8b, 0x81, 0x29, 0x4f, 0xe2, 0xc5, 0xdb, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x11, 0x66, 0x14, 0x4d, 0x72, 0x01, 0x00, 0x00,
}
