// Code generated by protoc-gen-go. DO NOT EDIT.
// source: upload_descriptor.proto

package deviced

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

type UploadDescriptorRequest struct {
	Descriptor_          *OpDescriptor `protobuf:"bytes,1,opt,name=descriptor,proto3" json:"descriptor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UploadDescriptorRequest) Reset()         { *m = UploadDescriptorRequest{} }
func (m *UploadDescriptorRequest) String() string { return proto.CompactTextString(m) }
func (*UploadDescriptorRequest) ProtoMessage()    {}
func (*UploadDescriptorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ec1c21ccd18c308, []int{0}
}

func (m *UploadDescriptorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadDescriptorRequest.Unmarshal(m, b)
}
func (m *UploadDescriptorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadDescriptorRequest.Marshal(b, m, deterministic)
}
func (m *UploadDescriptorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadDescriptorRequest.Merge(m, src)
}
func (m *UploadDescriptorRequest) XXX_Size() int {
	return xxx_messageInfo_UploadDescriptorRequest.Size(m)
}
func (m *UploadDescriptorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadDescriptorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadDescriptorRequest proto.InternalMessageInfo

func (m *UploadDescriptorRequest) GetDescriptor_() *OpDescriptor {
	if m != nil {
		return m.Descriptor_
	}
	return nil
}

type UploadDescriptorResponse struct {
	Descriptor_          *Descriptor `protobuf:"bytes,1,opt,name=descriptor,proto3" json:"descriptor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UploadDescriptorResponse) Reset()         { *m = UploadDescriptorResponse{} }
func (m *UploadDescriptorResponse) String() string { return proto.CompactTextString(m) }
func (*UploadDescriptorResponse) ProtoMessage()    {}
func (*UploadDescriptorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ec1c21ccd18c308, []int{1}
}

func (m *UploadDescriptorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadDescriptorResponse.Unmarshal(m, b)
}
func (m *UploadDescriptorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadDescriptorResponse.Marshal(b, m, deterministic)
}
func (m *UploadDescriptorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadDescriptorResponse.Merge(m, src)
}
func (m *UploadDescriptorResponse) XXX_Size() int {
	return xxx_messageInfo_UploadDescriptorResponse.Size(m)
}
func (m *UploadDescriptorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadDescriptorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadDescriptorResponse proto.InternalMessageInfo

func (m *UploadDescriptorResponse) GetDescriptor_() *Descriptor {
	if m != nil {
		return m.Descriptor_
	}
	return nil
}

func init() {
	proto.RegisterType((*UploadDescriptorRequest)(nil), "ai.metathings.service.deviced.UploadDescriptorRequest")
	proto.RegisterType((*UploadDescriptorResponse)(nil), "ai.metathings.service.deviced.UploadDescriptorResponse")
}

func init() { proto.RegisterFile("upload_descriptor.proto", fileDescriptor_7ec1c21ccd18c308) }

var fileDescriptor_7ec1c21ccd18c308 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x2d, 0xc8, 0xc9,
	0x4f, 0x4c, 0x89, 0x4f, 0x49, 0x2d, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0x2f, 0xd2, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4d, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc,
	0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4b, 0x49, 0x05, 0x51, 0x29,
	0x52, 0xe2, 0x65, 0x89, 0x39, 0x99, 0x29, 0x89, 0x25, 0xa9, 0xfa, 0x30, 0x06, 0x44, 0x9f, 0x14,
	0x77, 0x6e, 0x7e, 0x4a, 0x6a, 0x0e, 0x84, 0xa3, 0x54, 0xc0, 0x25, 0x1e, 0x0a, 0x36, 0xdf, 0x05,
	0x6e, 0x7c, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x50, 0x28, 0x17, 0x17, 0xc2, 0x4e, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x6d, 0x3d, 0xbc, 0x96, 0xea, 0xf9, 0x17, 0x20, 0xcc, 0x71,
	0xe2, 0xf8, 0xe5, 0xc4, 0xda, 0xc5, 0xc8, 0x24, 0xc0, 0x18, 0x84, 0x64, 0x90, 0x52, 0x2a, 0x97,
	0x04, 0xa6, 0x8d, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x9e, 0x58, 0xac, 0xd4, 0x24, 0x60,
	0x25, 0x92, 0x31, 0x48, 0x9a, 0x93, 0xd8, 0xc0, 0xfe, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x98, 0x02, 0xb6, 0x73, 0x3f, 0x01, 0x00, 0x00,
}
