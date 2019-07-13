// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remove_object.proto

package deviced

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RemoveObjectRequest struct {
	Object               *OpObject `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RemoveObjectRequest) Reset()         { *m = RemoveObjectRequest{} }
func (m *RemoveObjectRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveObjectRequest) ProtoMessage()    {}
func (*RemoveObjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_87a38042457eacbc, []int{0}
}

func (m *RemoveObjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveObjectRequest.Unmarshal(m, b)
}
func (m *RemoveObjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveObjectRequest.Marshal(b, m, deterministic)
}
func (m *RemoveObjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveObjectRequest.Merge(m, src)
}
func (m *RemoveObjectRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveObjectRequest.Size(m)
}
func (m *RemoveObjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveObjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveObjectRequest proto.InternalMessageInfo

func (m *RemoveObjectRequest) GetObject() *OpObject {
	if m != nil {
		return m.Object
	}
	return nil
}

func init() {
	proto.RegisterType((*RemoveObjectRequest)(nil), "ai.metathings.service.deviced.RemoveObjectRequest")
}

func init() { proto.RegisterFile("remove_object.proto", fileDescriptor_87a38042457eacbc) }

var fileDescriptor_87a38042457eacbc = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x4a, 0xcd, 0xcd,
	0x2f, 0x4b, 0x8d, 0xcf, 0x4f, 0xca, 0x4a, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x92, 0x4d, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4b, 0x49, 0x05, 0x51, 0x29, 0x52, 0x66, 0xe9, 0x99,
	0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xb9, 0xe5, 0x99, 0x25, 0xd9, 0xf9, 0xe5,
	0xfa, 0xe9, 0xf9, 0xba, 0x60, 0xbd, 0xba, 0x65, 0x89, 0x39, 0x99, 0x29, 0x89, 0x25, 0xf9, 0x45,
	0xc5, 0xfa, 0x70, 0x26, 0xc4, 0x58, 0x29, 0xee, 0xdc, 0xfc, 0x94, 0xd4, 0x1c, 0x08, 0x47, 0x29,
	0x8e, 0x4b, 0x38, 0x08, 0x6c, 0xb5, 0x3f, 0xd8, 0xe6, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12,
	0x21, 0x77, 0x2e, 0x36, 0x88, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0xd4, 0xf5, 0xf0,
	0xba, 0x45, 0xcf, 0xbf, 0x00, 0xa2, 0xdf, 0x89, 0xed, 0xd1, 0x7d, 0x79, 0x26, 0x05, 0xc6, 0x20,
	0xa8, 0xf6, 0x24, 0x36, 0xb0, 0x35, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xa1, 0x65,
	0x80, 0xe1, 0x00, 0x00, 0x00,
}
