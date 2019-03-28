// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_object_content.proto

package ai_metathings_service_device

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetObjectContentRequest struct {
	Object               *deviced.OpObject `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetObjectContentRequest) Reset()         { *m = GetObjectContentRequest{} }
func (m *GetObjectContentRequest) String() string { return proto.CompactTextString(m) }
func (*GetObjectContentRequest) ProtoMessage()    {}
func (*GetObjectContentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e38a252548b95948, []int{0}
}

func (m *GetObjectContentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetObjectContentRequest.Unmarshal(m, b)
}
func (m *GetObjectContentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetObjectContentRequest.Marshal(b, m, deterministic)
}
func (m *GetObjectContentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetObjectContentRequest.Merge(m, src)
}
func (m *GetObjectContentRequest) XXX_Size() int {
	return xxx_messageInfo_GetObjectContentRequest.Size(m)
}
func (m *GetObjectContentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetObjectContentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetObjectContentRequest proto.InternalMessageInfo

func (m *GetObjectContentRequest) GetObject() *deviced.OpObject {
	if m != nil {
		return m.Object
	}
	return nil
}

type GetObjectContentResponse struct {
	Content              *wrappers.BytesValue `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetObjectContentResponse) Reset()         { *m = GetObjectContentResponse{} }
func (m *GetObjectContentResponse) String() string { return proto.CompactTextString(m) }
func (*GetObjectContentResponse) ProtoMessage()    {}
func (*GetObjectContentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e38a252548b95948, []int{1}
}

func (m *GetObjectContentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetObjectContentResponse.Unmarshal(m, b)
}
func (m *GetObjectContentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetObjectContentResponse.Marshal(b, m, deterministic)
}
func (m *GetObjectContentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetObjectContentResponse.Merge(m, src)
}
func (m *GetObjectContentResponse) XXX_Size() int {
	return xxx_messageInfo_GetObjectContentResponse.Size(m)
}
func (m *GetObjectContentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetObjectContentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetObjectContentResponse proto.InternalMessageInfo

func (m *GetObjectContentResponse) GetContent() *wrappers.BytesValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*GetObjectContentRequest)(nil), "ai.metathings.service.device.GetObjectContentRequest")
	proto.RegisterType((*GetObjectContentResponse)(nil), "ai.metathings.service.device.GetObjectContentResponse")
}

func init() { proto.RegisterFile("get_object_content.proto", fileDescriptor_e38a252548b95948) }

var fileDescriptor_e38a252548b95948 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0xc7, 0x15, 0x86, 0x20, 0x85, 0x2d, 0x0b, 0x51, 0x41, 0x50, 0x75, 0x81, 0xa5, 0x67, 0x09,
	0x04, 0x0b, 0x5b, 0x19, 0x3a, 0x56, 0x64, 0x60, 0xad, 0x9c, 0xf8, 0x70, 0x4d, 0x13, 0x9f, 0xb1,
	0x2f, 0x8d, 0xfa, 0xb4, 0x48, 0x3c, 0x09, 0x22, 0x4e, 0x01, 0x09, 0x89, 0xc9, 0x96, 0xee, 0x7e,
	0xf7, 0xff, 0xc8, 0x0a, 0x8d, 0xbc, 0xa6, 0xea, 0x15, 0x6b, 0x5e, 0xd7, 0x64, 0x19, 0x2d, 0x83,
	0xf3, 0xc4, 0x94, 0x9f, 0x4b, 0x03, 0x2d, 0xb2, 0xe4, 0x8d, 0xb1, 0x3a, 0x40, 0x40, 0xbf, 0x33,
	0x35, 0x82, 0xc2, 0xaf, 0x67, 0x72, 0xa1, 0x89, 0x74, 0x83, 0x62, 0xd8, 0xad, 0xba, 0x17, 0xd1,
	0x7b, 0xe9, 0x1c, 0xfa, 0x10, 0xe9, 0xc9, 0xbd, 0x36, 0xbc, 0xe9, 0x2a, 0xa8, 0xa9, 0x15, 0x6d,
	0x6f, 0x78, 0x4b, 0xbd, 0xd0, 0x34, 0x1f, 0x86, 0xf3, 0x9d, 0x6c, 0x8c, 0x92, 0x4c, 0x3e, 0x88,
	0xef, 0xef, 0xc8, 0x3d, 0xfc, 0xe2, 0xac, 0xdc, 0x13, 0xb3, 0x14, 0x3f, 0x2e, 0x84, 0xdb, 0xea,
	0x28, 0x29, 0xa2, 0x0f, 0x25, 0x5a, 0x52, 0xd8, 0x44, 0x78, 0x56, 0x65, 0xa7, 0x4b, 0xe4, 0xd5,
	0x90, 0xe6, 0x31, 0x86, 0x29, 0xf1, 0xad, 0xc3, 0xc0, 0xf9, 0x32, 0x4b, 0x63, 0xca, 0x22, 0x99,
	0x26, 0xd7, 0x27, 0x37, 0x57, 0xf0, 0x5f, 0x3c, 0x05, 0x2b, 0x17, 0xcf, 0x2c, 0xd2, 0x8f, 0xf7,
	0xcb, 0xa3, 0x69, 0x52, 0x8e, 0xf8, 0xec, 0x29, 0x2b, 0xfe, 0x6a, 0x04, 0x47, 0x36, 0x60, 0x7e,
	0x97, 0x1d, 0x8f, 0x1d, 0x8e, 0x2a, 0x67, 0x10, 0x6b, 0x82, 0x43, 0x4d, 0xb0, 0xd8, 0x33, 0x86,
	0x67, 0xd9, 0x74, 0x58, 0x1e, 0x76, 0xab, 0x74, 0x98, 0xde, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xd4, 0x15, 0x75, 0x4f, 0x8c, 0x01, 0x00, 0x00,
}