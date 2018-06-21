// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list.proto

package camerad

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ListRequest struct {
	Core                 *OpCore   `protobuf:"bytes,1,opt,name=core" json:"core,omitempty"`
	Entity               *OpEntity `protobuf:"bytes,2,opt,name=entity" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_a29afbd89e1136b5, []int{0}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (dst *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(dst, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetCore() *OpCore {
	if m != nil {
		return m.Core
	}
	return nil
}

func (m *ListRequest) GetEntity() *OpEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type ListResponse struct {
	Camera               []*Camera `protobuf:"bytes,1,rep,name=camera" json:"camera,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_a29afbd89e1136b5, []int{1}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (dst *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(dst, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetCamera() []*Camera {
	if m != nil {
		return m.Camera
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "ai.metathings.service.camerad.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "ai.metathings.service.camerad.ListResponse")
}

func init() { proto.RegisterFile("list.proto", fileDescriptor_list_a29afbd89e1136b5) }

var fileDescriptor_list_a29afbd89e1136b5 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8f, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x89, 0xca, 0x1e, 0xd2, 0x9e, 0x72, 0x5a, 0x0a, 0x42, 0x29, 0x88, 0x5e, 0x9a, 0x05,
	0x05, 0xc1, 0x83, 0x78, 0x28, 0xde, 0x14, 0x21, 0x6f, 0x90, 0xa6, 0xc3, 0x76, 0xb0, 0xd9, 0x59,
	0x33, 0xb3, 0xbb, 0xf8, 0x08, 0xbe, 0xb5, 0x90, 0x2c, 0x1e, 0xd5, 0xd3, 0xcc, 0xc0, 0xff, 0xfd,
	0x7c, 0xa3, 0xf5, 0x09, 0x59, 0x6c, 0x9f, 0x48, 0xc8, 0x5c, 0x7a, 0xb4, 0x11, 0xc4, 0xcb, 0x11,
	0xbb, 0x96, 0x2d, 0x43, 0x1a, 0x31, 0x80, 0x0d, 0x3e, 0x42, 0xf2, 0x87, 0xd5, 0x7d, 0x8b, 0x72,
	0x1c, 0xf6, 0x36, 0x50, 0x6c, 0xe2, 0x84, 0xf2, 0x4e, 0x53, 0xd3, 0xd2, 0x36, 0xb3, 0xdb, 0xd1,
	0x9f, 0xf0, 0xe0, 0x85, 0x12, 0x37, 0x3f, 0x6b, 0xa9, 0x5d, 0x2d, 0x4b, 0x41, 0xb9, 0x36, 0x5f,
	0x4a, 0x2f, 0x5e, 0x90, 0xc5, 0xc1, 0xc7, 0x00, 0x2c, 0xe6, 0x41, 0x5f, 0x04, 0x4a, 0x50, 0xab,
	0xb5, 0xba, 0x59, 0xdc, 0x5e, 0xd9, 0x5f, 0x1d, 0xec, 0x5b, 0xbf, 0xa3, 0x04, 0x2e, 0x23, 0xe6,
	0x49, 0x57, 0xd0, 0x09, 0xca, 0x67, 0x7d, 0x96, 0xe1, 0xeb, 0x3f, 0xe1, 0xe7, 0x1c, 0x77, 0x33,
	0xb6, 0x79, 0xd5, 0xcb, 0xa2, 0xc2, 0x3d, 0x75, 0x0c, 0xe6, 0x51, 0x57, 0x25, 0x5b, 0xab, 0xf5,
	0xf9, 0x3f, 0x6c, 0x76, 0x79, 0xba, 0x19, 0xda, 0x57, 0xf9, 0xc3, 0xbb, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xb7, 0x43, 0xd3, 0x0a, 0x54, 0x01, 0x00, 0x00,
}