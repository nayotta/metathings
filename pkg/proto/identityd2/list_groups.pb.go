// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_groups.proto

package identityd2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
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

type ListGroupsRequest struct {
	Group                *OpGroup `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListGroupsRequest) Reset()         { *m = ListGroupsRequest{} }
func (m *ListGroupsRequest) String() string { return proto.CompactTextString(m) }
func (*ListGroupsRequest) ProtoMessage()    {}
func (*ListGroupsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e73ec0602ce3afc8, []int{0}
}

func (m *ListGroupsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGroupsRequest.Unmarshal(m, b)
}
func (m *ListGroupsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGroupsRequest.Marshal(b, m, deterministic)
}
func (m *ListGroupsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGroupsRequest.Merge(m, src)
}
func (m *ListGroupsRequest) XXX_Size() int {
	return xxx_messageInfo_ListGroupsRequest.Size(m)
}
func (m *ListGroupsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGroupsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListGroupsRequest proto.InternalMessageInfo

func (m *ListGroupsRequest) GetGroup() *OpGroup {
	if m != nil {
		return m.Group
	}
	return nil
}

type ListGroupsResponse struct {
	Groups               []*Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListGroupsResponse) Reset()         { *m = ListGroupsResponse{} }
func (m *ListGroupsResponse) String() string { return proto.CompactTextString(m) }
func (*ListGroupsResponse) ProtoMessage()    {}
func (*ListGroupsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e73ec0602ce3afc8, []int{1}
}

func (m *ListGroupsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGroupsResponse.Unmarshal(m, b)
}
func (m *ListGroupsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGroupsResponse.Marshal(b, m, deterministic)
}
func (m *ListGroupsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGroupsResponse.Merge(m, src)
}
func (m *ListGroupsResponse) XXX_Size() int {
	return xxx_messageInfo_ListGroupsResponse.Size(m)
}
func (m *ListGroupsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGroupsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListGroupsResponse proto.InternalMessageInfo

func (m *ListGroupsResponse) GetGroups() []*Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

func init() {
	proto.RegisterType((*ListGroupsRequest)(nil), "ai.metathings.service.identityd2.ListGroupsRequest")
	proto.RegisterType((*ListGroupsResponse)(nil), "ai.metathings.service.identityd2.ListGroupsResponse")
}

func init() { proto.RegisterFile("list_groups.proto", fileDescriptor_e73ec0602ce3afc8) }

var fileDescriptor_e73ec0602ce3afc8 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xb1, 0x6a, 0xc3, 0x30,
	0x10, 0x40, 0x31, 0xa5, 0x19, 0xe4, 0x29, 0x9e, 0x42, 0x86, 0x62, 0xb2, 0x34, 0x1d, 0x22, 0x81,
	0x0b, 0x5d, 0x33, 0x76, 0x29, 0x14, 0x42, 0x3b, 0x17, 0x39, 0xbe, 0x2a, 0x47, 0x6d, 0x9f, 0xaa,
	0x3b, 0xc5, 0xf4, 0xef, 0x0b, 0x92, 0x29, 0xdd, 0xbc, 0x49, 0xdc, 0xbd, 0xc7, 0xe3, 0xd4, 0xba,
	0x47, 0x96, 0x0f, 0x17, 0x28, 0x7a, 0xd6, 0x3e, 0x90, 0x50, 0x55, 0x5b, 0xd4, 0x03, 0x88, 0x95,
	0x0b, 0x8e, 0x8e, 0x35, 0x43, 0xb8, 0xe2, 0x19, 0x34, 0x76, 0x30, 0x0a, 0xca, 0x4f, 0xd7, 0x6c,
	0xef, 0x1c, 0x91, 0xeb, 0xc1, 0xa4, 0xfd, 0x36, 0x7e, 0x9a, 0x29, 0x58, 0xef, 0x21, 0xcc, 0x86,
	0xed, 0x93, 0x43, 0xb9, 0xc4, 0x56, 0x9f, 0x69, 0x30, 0xc3, 0x84, 0xf2, 0x45, 0x93, 0x71, 0x74,
	0x48, 0xc3, 0xc3, 0xd5, 0xf6, 0xd8, 0x59, 0xa1, 0xc0, 0xe6, 0xef, 0x39, 0x73, 0xe5, 0x40, 0x1d,
	0xf4, 0xf9, 0xb3, 0x7b, 0x53, 0xeb, 0x17, 0x64, 0x79, 0x4e, 0x69, 0x27, 0xf8, 0x8e, 0xc0, 0x52,
	0x1d, 0xd5, 0x6d, 0x6a, 0xdd, 0x14, 0x75, 0xb1, 0x2f, 0x9b, 0x07, 0xbd, 0xd4, 0xaa, 0x5f, 0x7d,
	0x32, 0x9c, 0x32, 0xb7, 0x7b, 0x57, 0xd5, 0x7f, 0x2b, 0x7b, 0x1a, 0x19, 0xaa, 0xa3, 0x5a, 0xe5,
	0x13, 0x6c, 0x8a, 0xfa, 0x66, 0x5f, 0x36, 0xf7, 0xcb, 0xde, 0x6c, 0x9d, 0xb1, 0x76, 0x95, 0x9a,
	0x1f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x09, 0x0b, 0xdf, 0xbe, 0x4f, 0x01, 0x00, 0x00,
}
