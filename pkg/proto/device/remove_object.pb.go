// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remove_object.proto

package device

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	deviced "github.com/nayotta/metathings/proto/deviced"
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
	Object               *deviced.OpObject `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
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

func (m *RemoveObjectRequest) GetObject() *deviced.OpObject {
	if m != nil {
		return m.Object
	}
	return nil
}

func init() {
	proto.RegisterType((*RemoveObjectRequest)(nil), "ai.metathings.service.device.RemoveObjectRequest")
}

func init() { proto.RegisterFile("remove_object.proto", fileDescriptor_87a38042457eacbc) }

var fileDescriptor_87a38042457eacbc = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x4a, 0xcd, 0xcd,
	0x2f, 0x4b, 0x8d, 0xcf, 0x4f, 0xca, 0x4a, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x92, 0x49, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4b, 0x49, 0x05, 0x51, 0x52, 0xe2, 0x65, 0x89, 0x39,
	0x99, 0x29, 0x89, 0x25, 0xa9, 0xfa, 0x30, 0x06, 0x44, 0x9b, 0x94, 0x79, 0x7a, 0x66, 0x49, 0x46,
	0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x5e, 0x62, 0x65, 0x7e, 0x49, 0x49, 0xa2, 0x3e, 0xc2,
	0x18, 0x7d, 0xb0, 0x22, 0x7d, 0x88, 0x21, 0x29, 0xfa, 0xb9, 0xf9, 0x29, 0xa9, 0x39, 0x10, 0x8d,
	0x4a, 0x09, 0x5c, 0xc2, 0x41, 0x60, 0x67, 0xf8, 0x83, 0x5d, 0x11, 0x94, 0x5a, 0x58, 0x9a, 0x5a,
	0x5c, 0x22, 0xe4, 0xc9, 0xc5, 0x06, 0x71, 0x96, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0xba,
	0x1e, 0x3e, 0x77, 0xa5, 0xe8, 0xf9, 0x17, 0x40, 0xf4, 0x3b, 0x71, 0xfc, 0x72, 0x62, 0xed, 0x62,
	0x64, 0x12, 0x60, 0x0c, 0x82, 0x1a, 0xe0, 0xc4, 0x11, 0xc5, 0x06, 0x51, 0x95, 0xc4, 0x06, 0xb6,
	0xd2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x18, 0xfb, 0x16, 0x5a, 0xf9, 0x00, 0x00, 0x00,
}
