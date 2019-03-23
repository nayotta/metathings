// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_object.proto

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

type GetObjectRequest struct {
	Object               *OpObject `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
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

func (m *GetObjectRequest) GetObject() *OpObject {
	if m != nil {
		return m.Object
	}
	return nil
}

type GetObjectResponse struct {
	Object               *Object  `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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

func (m *GetObjectResponse) GetObject() *Object {
	if m != nil {
		return m.Object
	}
	return nil
}

func init() {
	proto.RegisterType((*GetObjectRequest)(nil), "ai.metathings.service.deviced.GetObjectRequest")
	proto.RegisterType((*GetObjectResponse)(nil), "ai.metathings.service.deviced.GetObjectResponse")
}

func init() { proto.RegisterFile("get_object.proto", fileDescriptor_4fbb09ce0bdddc8b) }

var fileDescriptor_4fbb09ce0bdddc8b = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x4f, 0x2d, 0x89,
	0xcf, 0x4f, 0xca, 0x4a, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4d, 0xcc,
	0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x4b, 0x49, 0x05, 0x51, 0x29, 0x52, 0x66, 0xe9, 0x99, 0x25, 0x19, 0xa5,
	0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xb9, 0xe5, 0x99, 0x25, 0xd9, 0xf9, 0xe5, 0xfa, 0xe9, 0xf9,
	0xba, 0x60, 0xbd, 0xba, 0x65, 0x89, 0x39, 0x99, 0x29, 0x89, 0x25, 0xf9, 0x45, 0xc5, 0xfa, 0x70,
	0x26, 0xc4, 0x58, 0x29, 0xee, 0xdc, 0xfc, 0x94, 0xd4, 0x1c, 0x08, 0x47, 0x29, 0x9a, 0x4b, 0xc0,
	0x3d, 0xb5, 0xc4, 0x1f, 0x6c, 0x6d, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x3b, 0x17,
	0x1b, 0xc4, 0x1d, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xea, 0x7a, 0x78, 0x1d, 0xa2, 0xe7,
	0x5f, 0x00, 0xd1, 0xef, 0xc4, 0xf6, 0xe8, 0xbe, 0x3c, 0x93, 0x02, 0x63, 0x10, 0x54, 0xbb, 0x52,
	0x10, 0x97, 0x20, 0x92, 0xe1, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0xb6, 0x68, 0xa6, 0xab,
	0x12, 0x32, 0x1d, 0xa2, 0x1d, 0xaa, 0x29, 0x89, 0x0d, 0xec, 0x6e, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x85, 0x62, 0x64, 0x0a, 0x2f, 0x01, 0x00, 0x00,
}