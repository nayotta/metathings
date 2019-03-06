// Code generated by protoc-gen-go. DO NOT EDIT.
// source: patch_group.proto

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

type PatchGroupRequest struct {
	Group                *OpGroup `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PatchGroupRequest) Reset()         { *m = PatchGroupRequest{} }
func (m *PatchGroupRequest) String() string { return proto.CompactTextString(m) }
func (*PatchGroupRequest) ProtoMessage()    {}
func (*PatchGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_74f1cd4f90fa4471, []int{0}
}

func (m *PatchGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchGroupRequest.Unmarshal(m, b)
}
func (m *PatchGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchGroupRequest.Marshal(b, m, deterministic)
}
func (m *PatchGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchGroupRequest.Merge(m, src)
}
func (m *PatchGroupRequest) XXX_Size() int {
	return xxx_messageInfo_PatchGroupRequest.Size(m)
}
func (m *PatchGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PatchGroupRequest proto.InternalMessageInfo

func (m *PatchGroupRequest) GetGroup() *OpGroup {
	if m != nil {
		return m.Group
	}
	return nil
}

type PatchGroupResponse struct {
	Group                *Group   `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PatchGroupResponse) Reset()         { *m = PatchGroupResponse{} }
func (m *PatchGroupResponse) String() string { return proto.CompactTextString(m) }
func (*PatchGroupResponse) ProtoMessage()    {}
func (*PatchGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_74f1cd4f90fa4471, []int{1}
}

func (m *PatchGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchGroupResponse.Unmarshal(m, b)
}
func (m *PatchGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchGroupResponse.Marshal(b, m, deterministic)
}
func (m *PatchGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchGroupResponse.Merge(m, src)
}
func (m *PatchGroupResponse) XXX_Size() int {
	return xxx_messageInfo_PatchGroupResponse.Size(m)
}
func (m *PatchGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PatchGroupResponse proto.InternalMessageInfo

func (m *PatchGroupResponse) GetGroup() *Group {
	if m != nil {
		return m.Group
	}
	return nil
}

func init() {
	proto.RegisterType((*PatchGroupRequest)(nil), "ai.metathings.service.identityd2.PatchGroupRequest")
	proto.RegisterType((*PatchGroupResponse)(nil), "ai.metathings.service.identityd2.PatchGroupResponse")
}

func init() { proto.RegisterFile("patch_group.proto", fileDescriptor_74f1cd4f90fa4471) }

var fileDescriptor_74f1cd4f90fa4471 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8e, 0xb1, 0x4a, 0x04, 0x31,
	0x10, 0x40, 0x59, 0xc1, 0x2b, 0xf6, 0xaa, 0x4b, 0x25, 0x57, 0xe8, 0x72, 0x8d, 0x5a, 0x5c, 0x02,
	0x27, 0xd8, 0xd9, 0xd8, 0x5c, 0xa9, 0x9c, 0xad, 0x20, 0xd9, 0xcd, 0x98, 0x1d, 0xdc, 0xec, 0xc4,
	0x64, 0x72, 0x8b, 0x5f, 0x2b, 0xf8, 0x25, 0x72, 0x89, 0x88, 0x56, 0xd7, 0xcd, 0x30, 0xf3, 0x1e,
	0xaf, 0x5e, 0x78, 0xcd, 0x5d, 0xff, 0x62, 0x03, 0x25, 0x2f, 0x7d, 0x20, 0x26, 0xd1, 0x68, 0x94,
	0x0e, 0x58, 0x73, 0x8f, 0xa3, 0x8d, 0x32, 0x42, 0xd8, 0x63, 0x07, 0x12, 0x0d, 0x8c, 0x8c, 0xfc,
	0x61, 0x36, 0xcb, 0x73, 0x4b, 0x64, 0x07, 0x50, 0xf9, 0xbf, 0x4d, 0xaf, 0x6a, 0x0a, 0xda, 0x7b,
	0x08, 0xb1, 0x18, 0x96, 0xb7, 0x16, 0xb9, 0x4f, 0xad, 0xec, 0xc8, 0x29, 0x37, 0x21, 0xbf, 0xd1,
	0xa4, 0x2c, 0xad, 0xf3, 0x71, 0xbd, 0xd7, 0x03, 0x1a, 0xcd, 0x14, 0xa2, 0xfa, 0x1d, 0x7f, 0xb8,
	0xb9, 0x23, 0x03, 0x43, 0x59, 0x56, 0xcf, 0xf5, 0xe2, 0xf1, 0xd0, 0xb6, 0x3d, 0xa4, 0xed, 0xe0,
	0x3d, 0x41, 0x64, 0xb1, 0xad, 0x4f, 0x73, 0xea, 0x59, 0xd5, 0x54, 0x57, 0xf3, 0xcd, 0xb5, 0x3c,
	0xd6, 0x2a, 0x1f, 0x7c, 0x16, 0xdc, 0xcf, 0xbe, 0x3e, 0x2f, 0x4e, 0x9a, 0x6a, 0x57, 0xf8, 0xd5,
	0x53, 0x2d, 0xfe, 0xda, 0xa3, 0xa7, 0x31, 0x82, 0xb8, 0xfb, 0xaf, 0xbf, 0x3c, 0xae, 0x2f, 0x7c,
	0xa1, 0xda, 0x59, 0x2e, 0xbf, 0xf9, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x16, 0x6a, 0x52, 0x8e, 0x55,
	0x01, 0x00, 0x00,
}
