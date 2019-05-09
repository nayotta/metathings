// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_groups_for_object.proto

package identityd2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ListGroupsForObjectRequest struct {
	Object               *OpEntity `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListGroupsForObjectRequest) Reset()         { *m = ListGroupsForObjectRequest{} }
func (m *ListGroupsForObjectRequest) String() string { return proto.CompactTextString(m) }
func (*ListGroupsForObjectRequest) ProtoMessage()    {}
func (*ListGroupsForObjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_53596400b56b7ca9, []int{0}
}

func (m *ListGroupsForObjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGroupsForObjectRequest.Unmarshal(m, b)
}
func (m *ListGroupsForObjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGroupsForObjectRequest.Marshal(b, m, deterministic)
}
func (m *ListGroupsForObjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGroupsForObjectRequest.Merge(m, src)
}
func (m *ListGroupsForObjectRequest) XXX_Size() int {
	return xxx_messageInfo_ListGroupsForObjectRequest.Size(m)
}
func (m *ListGroupsForObjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGroupsForObjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListGroupsForObjectRequest proto.InternalMessageInfo

func (m *ListGroupsForObjectRequest) GetObject() *OpEntity {
	if m != nil {
		return m.Object
	}
	return nil
}

type ListGroupsForObjectResponse struct {
	Groups               []*Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListGroupsForObjectResponse) Reset()         { *m = ListGroupsForObjectResponse{} }
func (m *ListGroupsForObjectResponse) String() string { return proto.CompactTextString(m) }
func (*ListGroupsForObjectResponse) ProtoMessage()    {}
func (*ListGroupsForObjectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_53596400b56b7ca9, []int{1}
}

func (m *ListGroupsForObjectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGroupsForObjectResponse.Unmarshal(m, b)
}
func (m *ListGroupsForObjectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGroupsForObjectResponse.Marshal(b, m, deterministic)
}
func (m *ListGroupsForObjectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGroupsForObjectResponse.Merge(m, src)
}
func (m *ListGroupsForObjectResponse) XXX_Size() int {
	return xxx_messageInfo_ListGroupsForObjectResponse.Size(m)
}
func (m *ListGroupsForObjectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGroupsForObjectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListGroupsForObjectResponse proto.InternalMessageInfo

func (m *ListGroupsForObjectResponse) GetGroups() []*Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

func init() {
	proto.RegisterType((*ListGroupsForObjectRequest)(nil), "ai.metathings.service.identityd2.ListGroupsForObjectRequest")
	proto.RegisterType((*ListGroupsForObjectResponse)(nil), "ai.metathings.service.identityd2.ListGroupsForObjectResponse")
}

func init() { proto.RegisterFile("list_groups_for_object.proto", fileDescriptor_53596400b56b7ca9) }

var fileDescriptor_53596400b56b7ca9 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x89, 0x42, 0x0e, 0x9b, 0x5b, 0x4e, 0x25, 0x0a, 0x86, 0x5e, 0x2c, 0x42, 0x37, 0x10,
	0xc1, 0xab, 0x20, 0xa8, 0x20, 0x42, 0x21, 0x0f, 0x60, 0xc8, 0x9f, 0x31, 0x19, 0x4d, 0x32, 0xeb,
	0xce, 0xa4, 0xc5, 0xa7, 0x15, 0x7c, 0x12, 0x61, 0xb7, 0x78, 0x2a, 0xe4, 0xf6, 0xed, 0xe1, 0xfb,
	0x7d, 0xbf, 0x1d, 0x75, 0x39, 0x20, 0x4b, 0xd9, 0x59, 0x9a, 0x0d, 0x97, 0xef, 0x64, 0x4b, 0xaa,
	0x3f, 0xa0, 0x11, 0x6d, 0x2c, 0x09, 0xc5, 0x69, 0x85, 0x7a, 0x04, 0xa9, 0xa4, 0xc7, 0xa9, 0x63,
	0xcd, 0x60, 0xf7, 0xd8, 0x80, 0xc6, 0x16, 0x26, 0x41, 0xf9, 0x6e, 0xf3, 0xe4, 0xae, 0x43, 0xe9,
	0xe7, 0x5a, 0x37, 0x34, 0x66, 0xe3, 0x01, 0xe5, 0x93, 0x0e, 0x59, 0x47, 0x5b, 0x57, 0xdf, 0xee,
	0xab, 0x01, 0xdb, 0x4a, 0xc8, 0x72, 0xf6, 0x1f, 0x3d, 0x39, 0x89, 0x46, 0x6a, 0x61, 0xf0, 0x8f,
	0x75, 0xaf, 0x92, 0x57, 0x64, 0x79, 0x76, 0x16, 0x4f, 0x64, 0x77, 0xce, 0xa1, 0x80, 0xaf, 0x19,
	0x58, 0xe2, 0x17, 0x15, 0x7a, 0xa9, 0x55, 0x90, 0x06, 0x9b, 0x28, 0xbf, 0xd1, 0x4b, 0x56, 0x7a,
	0x67, 0x1e, 0x5d, 0x7c, 0x08, 0x7f, 0x7f, 0xae, 0xce, 0xd2, 0xa0, 0x38, 0x12, 0xd6, 0x6f, 0xea,
	0xe2, 0xe4, 0x12, 0x1b, 0x9a, 0x18, 0xe2, 0x7b, 0x15, 0xfa, 0x53, 0xac, 0x82, 0xf4, 0x7c, 0x13,
	0xe5, 0xd7, 0xcb, 0x53, 0x0e, 0x55, 0x1c, 0x6b, 0x75, 0xe8, 0x3e, 0x74, 0xfb, 0x17, 0x00, 0x00,
	0xff, 0xff, 0x36, 0x08, 0x6f, 0x72, 0x57, 0x01, 0x00, 0x00,
}