// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delete_group.proto

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

type DeleteGroupRequest struct {
	Group                *OpGroup `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteGroupRequest) Reset()         { *m = DeleteGroupRequest{} }
func (m *DeleteGroupRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteGroupRequest) ProtoMessage()    {}
func (*DeleteGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2d2818c2a690017, []int{0}
}

func (m *DeleteGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteGroupRequest.Unmarshal(m, b)
}
func (m *DeleteGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteGroupRequest.Marshal(b, m, deterministic)
}
func (m *DeleteGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteGroupRequest.Merge(m, src)
}
func (m *DeleteGroupRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteGroupRequest.Size(m)
}
func (m *DeleteGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteGroupRequest proto.InternalMessageInfo

func (m *DeleteGroupRequest) GetGroup() *OpGroup {
	if m != nil {
		return m.Group
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteGroupRequest)(nil), "ai.metathings.service.identityd2.DeleteGroupRequest")
}

func init() { proto.RegisterFile("delete_group.proto", fileDescriptor_a2d2818c2a690017) }

var fileDescriptor_a2d2818c2a690017 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x49, 0xcd, 0x49,
	0x2d, 0x49, 0x8d, 0x4f, 0x2f, 0xca, 0x2f, 0x2d, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52,
	0x48, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e,
	0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0xcb, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0x4c, 0x31,
	0x92, 0x12, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0x20, 0x5a, 0xa5,
	0xb8, 0x73, 0xf3, 0x53, 0x52, 0x73, 0x20, 0x1c, 0xa5, 0x78, 0x2e, 0x21, 0x17, 0xb0, 0xe9, 0xee,
	0x20, 0xc3, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x3c, 0xb9, 0x58, 0xc1, 0x96, 0x49,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x69, 0xea, 0x11, 0xb2, 0x4d, 0xcf, 0xbf, 0x00, 0x6c, 0x80,
	0x13, 0xc7, 0x2f, 0x27, 0xd6, 0x2e, 0x46, 0x26, 0x01, 0xc6, 0x20, 0x88, 0x09, 0x49, 0x6c, 0x60,
	0x7b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x20, 0x60, 0xa9, 0xc5, 0x00, 0x00, 0x00,
}
