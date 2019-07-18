// Code generated by protoc-gen-go. DO NOT EDIT.
// source: show_groups.proto

package identityd2

import (
	fmt "fmt"
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

type ShowGroupsResponse struct {
	Groups               []*Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowGroupsResponse) Reset()         { *m = ShowGroupsResponse{} }
func (m *ShowGroupsResponse) String() string { return proto.CompactTextString(m) }
func (*ShowGroupsResponse) ProtoMessage()    {}
func (*ShowGroupsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_775e18fd48520bfe, []int{0}
}

func (m *ShowGroupsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowGroupsResponse.Unmarshal(m, b)
}
func (m *ShowGroupsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowGroupsResponse.Marshal(b, m, deterministic)
}
func (m *ShowGroupsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowGroupsResponse.Merge(m, src)
}
func (m *ShowGroupsResponse) XXX_Size() int {
	return xxx_messageInfo_ShowGroupsResponse.Size(m)
}
func (m *ShowGroupsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowGroupsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShowGroupsResponse proto.InternalMessageInfo

func (m *ShowGroupsResponse) GetGroups() []*Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

func init() {
	proto.RegisterType((*ShowGroupsResponse)(nil), "ai.metathings.service.identityd2.ShowGroupsResponse")
}

func init() { proto.RegisterFile("show_groups.proto", fileDescriptor_775e18fd48520bfe) }

var fileDescriptor_775e18fd48520bfe = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xce, 0xc8, 0x2f,
	0x8f, 0x4f, 0x2f, 0xca, 0x2f, 0x2d, 0x28, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x48,
	0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d,
	0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0xcb, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0x4c, 0x31, 0x92,
	0xe2, 0xce, 0xcd, 0x4f, 0x49, 0xcd, 0x81, 0x28, 0x57, 0x0a, 0xe5, 0x12, 0x0a, 0xce, 0xc8, 0x2f,
	0x77, 0x07, 0x1b, 0x11, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x64, 0xcf, 0xc5, 0x06,
	0x31, 0x54, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x5d, 0x8f, 0x90, 0xa9, 0x7a, 0x60, 0x13,
	0x82, 0xa0, 0xda, 0x92, 0xd8, 0xc0, 0xa6, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x0b,
	0x9d, 0x6f, 0xa1, 0x00, 0x00, 0x00,
}
