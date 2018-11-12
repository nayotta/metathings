// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delete.proto

package streamd

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

type DeleteRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_600d681a62b3a9a7, []int{0}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteRequest)(nil), "ai.metathings.service.streamd.DeleteRequest")
}

func init() { proto.RegisterFile("delete.proto", fileDescriptor_600d681a62b3a9a7) }

var fileDescriptor_600d681a62b3a9a7 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0x4d, 0x4a, 0xc7, 0x30,
	0x10, 0x47, 0x69, 0x17, 0x5d, 0x54, 0xdd, 0x74, 0x25, 0xc5, 0x8f, 0xe2, 0xca, 0x4d, 0x27, 0xa0,
	0xe2, 0x01, 0x44, 0x2f, 0x50, 0xc1, 0x7d, 0xda, 0x8c, 0xe9, 0x60, 0xd2, 0xa9, 0xc9, 0xa4, 0x3d,
	0xae, 0xe0, 0x49, 0x84, 0x14, 0xff, 0xbb, 0x81, 0x79, 0x8f, 0xf7, 0xab, 0xcf, 0x0d, 0x3a, 0x14,
	0x84, 0x35, 0xb0, 0x70, 0x73, 0xad, 0x09, 0x3c, 0x8a, 0x96, 0x99, 0x16, 0x1b, 0x21, 0x62, 0xd8,
	0x68, 0x42, 0x88, 0x12, 0x50, 0x7b, 0xd3, 0xde, 0x58, 0x66, 0xeb, 0x50, 0x65, 0x78, 0x4c, 0x9f,
	0x6a, 0x0f, 0x7a, 0x5d, 0x31, 0xc4, 0x43, 0x6f, 0x9f, 0x2d, 0xc9, 0x9c, 0x46, 0x98, 0xd8, 0x2b,
	0xbf, 0x93, 0x7c, 0xf1, 0xae, 0x2c, 0xf7, 0xf9, 0xd9, 0x6f, 0xda, 0x91, 0xd1, 0xc2, 0x21, 0xaa,
	0xd3, 0x79, 0x78, 0x77, 0x6f, 0xf5, 0xc5, 0x6b, 0x9e, 0x31, 0xe0, 0x77, 0xc2, 0x28, 0xcd, 0x53,
	0x5d, 0x92, 0xb9, 0x2c, 0xba, 0xe2, 0xfe, 0xec, 0xe1, 0x0a, 0x8e, 0x2a, 0xfc, 0x57, 0xe1, 0x5d,
	0x02, 0x2d, 0xf6, 0x43, 0xbb, 0x84, 0x2f, 0xd5, 0xef, 0xcf, 0x6d, 0xd9, 0x15, 0x43, 0x49, 0x66,
	0xac, 0x32, 0xf1, 0xf8, 0x17, 0x00, 0x00, 0xff, 0xff, 0xb2, 0xc2, 0xe8, 0x7c, 0xd4, 0x00, 0x00,
	0x00,
}
