// Code generated by protoc-gen-go. DO NOT EDIT.
// source: show_core.proto

package cored

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ShowCoreResponse struct {
	Core                 *Core    `protobuf:"bytes,1,opt,name=core,proto3" json:"core,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowCoreResponse) Reset()         { *m = ShowCoreResponse{} }
func (m *ShowCoreResponse) String() string { return proto.CompactTextString(m) }
func (*ShowCoreResponse) ProtoMessage()    {}
func (*ShowCoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_51088005da6dbdcd, []int{0}
}

func (m *ShowCoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowCoreResponse.Unmarshal(m, b)
}
func (m *ShowCoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowCoreResponse.Marshal(b, m, deterministic)
}
func (m *ShowCoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowCoreResponse.Merge(m, src)
}
func (m *ShowCoreResponse) XXX_Size() int {
	return xxx_messageInfo_ShowCoreResponse.Size(m)
}
func (m *ShowCoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowCoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShowCoreResponse proto.InternalMessageInfo

func (m *ShowCoreResponse) GetCore() *Core {
	if m != nil {
		return m.Core
	}
	return nil
}

func init() {
	proto.RegisterType((*ShowCoreResponse)(nil), "ai.metathings.service.cored.ShowCoreResponse")
}

func init() { proto.RegisterFile("show_core.proto", fileDescriptor_51088005da6dbdcd) }

var fileDescriptor_51088005da6dbdcd = []byte{
	// 120 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xce, 0xc8, 0x2f,
	0x8f, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4e, 0xcc, 0xd4,
	0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0xd5, 0x03, 0xa9, 0x48, 0x91, 0xe2, 0x42, 0x28, 0x54, 0xf2, 0xe4, 0x12, 0x08, 0xce,
	0xc8, 0x2f, 0x77, 0xce, 0x2f, 0x4a, 0x0d, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x32,
	0xe5, 0x62, 0x01, 0xa9, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x52, 0xd4, 0xc3, 0x63, 0x96,
	0x1e, 0x58, 0x23, 0x58, 0x79, 0x12, 0x1b, 0xd8, 0x44, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x54, 0xf4, 0xf2, 0x85, 0x8d, 0x00, 0x00, 0x00,
}
