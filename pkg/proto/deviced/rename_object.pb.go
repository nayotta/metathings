// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rename_object.proto

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RenameObjectRequest struct {
	Source               *OpObject `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Destination          *OpObject `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RenameObjectRequest) Reset()         { *m = RenameObjectRequest{} }
func (m *RenameObjectRequest) String() string { return proto.CompactTextString(m) }
func (*RenameObjectRequest) ProtoMessage()    {}
func (*RenameObjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a65dde1df189b9d, []int{0}
}

func (m *RenameObjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenameObjectRequest.Unmarshal(m, b)
}
func (m *RenameObjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenameObjectRequest.Marshal(b, m, deterministic)
}
func (m *RenameObjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenameObjectRequest.Merge(m, src)
}
func (m *RenameObjectRequest) XXX_Size() int {
	return xxx_messageInfo_RenameObjectRequest.Size(m)
}
func (m *RenameObjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RenameObjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RenameObjectRequest proto.InternalMessageInfo

func (m *RenameObjectRequest) GetSource() *OpObject {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *RenameObjectRequest) GetDestination() *OpObject {
	if m != nil {
		return m.Destination
	}
	return nil
}

func init() {
	proto.RegisterType((*RenameObjectRequest)(nil), "ai.metathings.service.deviced.RenameObjectRequest")
}

func init() { proto.RegisterFile("rename_object.proto", fileDescriptor_2a65dde1df189b9d) }

var fileDescriptor_2a65dde1df189b9d = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x4a, 0xcd, 0x4b,
	0xcc, 0x4d, 0x8d, 0xcf, 0x4f, 0xca, 0x4a, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x92, 0x4d, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4b, 0x49, 0x05, 0x51, 0x29, 0x52, 0x66, 0xe9, 0x99,
	0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xb9, 0xe5, 0x99, 0x25, 0xd9, 0xf9, 0xe5,
	0xfa, 0xe9, 0xf9, 0xba, 0x60, 0xbd, 0xba, 0x65, 0x89, 0x39, 0x99, 0x29, 0x89, 0x25, 0xf9, 0x45,
	0xc5, 0xfa, 0x70, 0x26, 0xc4, 0x58, 0x29, 0xee, 0xdc, 0xfc, 0x94, 0xd4, 0x1c, 0x08, 0x47, 0x69,
	0x23, 0x23, 0x97, 0x70, 0x10, 0xd8, 0x6e, 0x7f, 0xb0, 0xd5, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5,
	0x25, 0x42, 0xee, 0x5c, 0x6c, 0xc5, 0xf9, 0xa5, 0x45, 0xc9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0xdc, 0x46, 0xea, 0x7a, 0x78, 0x1d, 0xa3, 0xe7, 0x5f, 0x00, 0xd1, 0xef, 0xc4, 0xf6, 0xe8, 0xbe,
	0x3c, 0x93, 0x02, 0x63, 0x10, 0x54, 0xbb, 0x50, 0x20, 0x17, 0x77, 0x4a, 0x6a, 0x71, 0x49, 0x66,
	0x5e, 0x62, 0x49, 0x66, 0x7e, 0x9e, 0x04, 0x13, 0x79, 0xa6, 0x21, 0x9b, 0x91, 0xc4, 0x06, 0x76,
	0xba, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xef, 0x90, 0x03, 0xee, 0x35, 0x01, 0x00, 0x00,
}
