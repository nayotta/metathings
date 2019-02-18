// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delete_bundle.proto

package datad

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

type DeleteBundleRequest struct {
	Bundle               *OpBundle `protobuf:"bytes,1,opt,name=bundle,proto3" json:"bundle,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DeleteBundleRequest) Reset()         { *m = DeleteBundleRequest{} }
func (m *DeleteBundleRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBundleRequest) ProtoMessage()    {}
func (*DeleteBundleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3165fa33e0dddef, []int{0}
}

func (m *DeleteBundleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBundleRequest.Unmarshal(m, b)
}
func (m *DeleteBundleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBundleRequest.Marshal(b, m, deterministic)
}
func (m *DeleteBundleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBundleRequest.Merge(m, src)
}
func (m *DeleteBundleRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBundleRequest.Size(m)
}
func (m *DeleteBundleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBundleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBundleRequest proto.InternalMessageInfo

func (m *DeleteBundleRequest) GetBundle() *OpBundle {
	if m != nil {
		return m.Bundle
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteBundleRequest)(nil), "ai.metathings.service.datad.DeleteBundleRequest")
}

func init() { proto.RegisterFile("delete_bundle.proto", fileDescriptor_a3165fa33e0dddef) }

var fileDescriptor_a3165fa33e0dddef = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x49, 0xcd, 0x49,
	0x2d, 0x49, 0x8d, 0x4f, 0x2a, 0xcd, 0x4b, 0xc9, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x92, 0x4e, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4b, 0x49, 0x2c, 0x49, 0x4c, 0x91, 0x32, 0x4b, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x2d, 0xcf, 0x2c, 0xc9, 0xce, 0x2f,
	0xd7, 0x4f, 0xcf, 0xd7, 0x05, 0xeb, 0xd4, 0x2d, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f,
	0x2a, 0xd6, 0x87, 0x33, 0x21, 0x86, 0x4a, 0x71, 0xe7, 0xe6, 0xa7, 0xa4, 0xe6, 0x40, 0x38, 0x4a,
	0x31, 0x5c, 0xc2, 0x2e, 0x60, 0x8b, 0x9d, 0xc0, 0xf6, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97,
	0x08, 0xb9, 0x72, 0xb1, 0x41, 0x1c, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xa4, 0xaa, 0x87,
	0xc7, 0x25, 0x7a, 0xfe, 0x05, 0x10, 0xdd, 0x4e, 0x6c, 0x8f, 0xee, 0xcb, 0x33, 0x29, 0x30, 0x06,
	0x41, 0x35, 0x27, 0xb1, 0x81, 0x2d, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x40, 0x88,
	0x98, 0xdd, 0x00, 0x00, 0x00,
}
