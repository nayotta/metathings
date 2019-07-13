// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_object.proto

package ai_metathings_service_device

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	deviced "github.com/nayotta/metathings/pkg/proto/deviced"
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

type GetObjectRequest struct {
	Object               *deviced.OpObject `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetObjectRequest) Reset()         { *m = GetObjectRequest{} }
func (m *GetObjectRequest) String() string { return proto.CompactTextString(m) }
func (*GetObjectRequest) ProtoMessage()    {}
func (*GetObjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fbb09ce0bdddc8b, []int{0}
}

func (m *GetObjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetObjectRequest.Unmarshal(m, b)
}
func (m *GetObjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetObjectRequest.Marshal(b, m, deterministic)
}
func (m *GetObjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetObjectRequest.Merge(m, src)
}
func (m *GetObjectRequest) XXX_Size() int {
	return xxx_messageInfo_GetObjectRequest.Size(m)
}
func (m *GetObjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetObjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetObjectRequest proto.InternalMessageInfo

func (m *GetObjectRequest) GetObject() *deviced.OpObject {
	if m != nil {
		return m.Object
	}
	return nil
}

type GetObjectResponse struct {
	Object               *deviced.Object `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetObjectResponse) Reset()         { *m = GetObjectResponse{} }
func (m *GetObjectResponse) String() string { return proto.CompactTextString(m) }
func (*GetObjectResponse) ProtoMessage()    {}
func (*GetObjectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fbb09ce0bdddc8b, []int{1}
}

func (m *GetObjectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetObjectResponse.Unmarshal(m, b)
}
func (m *GetObjectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetObjectResponse.Marshal(b, m, deterministic)
}
func (m *GetObjectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetObjectResponse.Merge(m, src)
}
func (m *GetObjectResponse) XXX_Size() int {
	return xxx_messageInfo_GetObjectResponse.Size(m)
}
func (m *GetObjectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetObjectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetObjectResponse proto.InternalMessageInfo

func (m *GetObjectResponse) GetObject() *deviced.Object {
	if m != nil {
		return m.Object
	}
	return nil
}

func init() {
	proto.RegisterType((*GetObjectRequest)(nil), "ai.metathings.service.device.GetObjectRequest")
	proto.RegisterType((*GetObjectResponse)(nil), "ai.metathings.service.device.GetObjectResponse")
}

func init() { proto.RegisterFile("get_object.proto", fileDescriptor_4fbb09ce0bdddc8b) }

var fileDescriptor_4fbb09ce0bdddc8b = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8e, 0xb1, 0x4b, 0xc4, 0x30,
	0x14, 0x87, 0xa9, 0x43, 0x87, 0xb8, 0x9c, 0x9d, 0xe4, 0x10, 0x3c, 0x0e, 0x44, 0x97, 0xcb, 0x03,
	0x05, 0x17, 0x71, 0x71, 0xb9, 0xf1, 0xa0, 0xab, 0x83, 0xa4, 0xcd, 0x23, 0x17, 0xaf, 0xe9, 0x8b,
	0xcd, 0x6b, 0x8b, 0x7f, 0xad, 0xe0, 0x5f, 0x22, 0x26, 0x45, 0x8b, 0x83, 0x37, 0x25, 0xc3, 0xfb,
	0xbe, 0xdf, 0x27, 0x16, 0x06, 0xf9, 0x85, 0xaa, 0x57, 0xac, 0x59, 0xfa, 0x8e, 0x98, 0x8a, 0x0b,
	0x65, 0xa5, 0x43, 0x56, 0xbc, 0xb7, 0xad, 0x09, 0x32, 0x60, 0x37, 0xd8, 0x1a, 0xa5, 0xc6, 0xef,
	0x67, 0x79, 0x6f, 0x2c, 0xef, 0xfb, 0x4a, 0xd6, 0xe4, 0xc0, 0x8d, 0x96, 0x0f, 0x34, 0x82, 0xa1,
	0x4d, 0x44, 0x37, 0x83, 0x6a, 0xac, 0x56, 0x4c, 0x5d, 0x80, 0x9f, 0x6f, 0xb2, 0x2e, 0x1f, 0x66,
	0x5c, 0xab, 0xde, 0x89, 0x59, 0xc1, 0xef, 0x0a, 0xf8, 0x83, 0x81, 0x78, 0x08, 0x69, 0x47, 0x83,
	0x23, 0x8d, 0x4d, 0x82, 0xd7, 0xcf, 0x62, 0xb1, 0x45, 0xde, 0xc5, 0xca, 0x12, 0xdf, 0x7a, 0x0c,
	0x5c, 0x6c, 0x45, 0x9e, 0xb2, 0xcf, 0xb3, 0x55, 0x76, 0x73, 0x7a, 0x7b, 0x2d, 0xff, 0xeb, 0xd6,
	0x72, 0xe7, 0x13, 0xff, 0x94, 0x7f, 0x7e, 0x5c, 0x9e, 0xac, 0xb2, 0x72, 0xc2, 0xd7, 0xa5, 0x38,
	0x9b, 0xc9, 0x83, 0xa7, 0x36, 0x60, 0xf1, 0xf8, 0xc7, 0x7e, 0x75, 0xcc, 0x9e, 0xf0, 0x09, 0xaa,
	0xf2, 0xd8, 0x7d, 0xf7, 0x15, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x7b, 0xb8, 0x0e, 0x5e, 0x01, 0x00,
	0x00,
}
