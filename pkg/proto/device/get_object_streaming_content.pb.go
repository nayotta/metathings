// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_object_streaming_content.proto

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

type GetObjectStreamingContentRequest struct {
	Object               *deviced.OpObject `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetObjectStreamingContentRequest) Reset()         { *m = GetObjectStreamingContentRequest{} }
func (m *GetObjectStreamingContentRequest) String() string { return proto.CompactTextString(m) }
func (*GetObjectStreamingContentRequest) ProtoMessage()    {}
func (*GetObjectStreamingContentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_349f78370bb4aa14, []int{0}
}

func (m *GetObjectStreamingContentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetObjectStreamingContentRequest.Unmarshal(m, b)
}
func (m *GetObjectStreamingContentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetObjectStreamingContentRequest.Marshal(b, m, deterministic)
}
func (m *GetObjectStreamingContentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetObjectStreamingContentRequest.Merge(m, src)
}
func (m *GetObjectStreamingContentRequest) XXX_Size() int {
	return xxx_messageInfo_GetObjectStreamingContentRequest.Size(m)
}
func (m *GetObjectStreamingContentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetObjectStreamingContentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetObjectStreamingContentRequest proto.InternalMessageInfo

func (m *GetObjectStreamingContentRequest) GetObject() *deviced.OpObject {
	if m != nil {
		return m.Object
	}
	return nil
}

type GetObjectStreamingContentResponse struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetObjectStreamingContentResponse) Reset()         { *m = GetObjectStreamingContentResponse{} }
func (m *GetObjectStreamingContentResponse) String() string { return proto.CompactTextString(m) }
func (*GetObjectStreamingContentResponse) ProtoMessage()    {}
func (*GetObjectStreamingContentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_349f78370bb4aa14, []int{1}
}

func (m *GetObjectStreamingContentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetObjectStreamingContentResponse.Unmarshal(m, b)
}
func (m *GetObjectStreamingContentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetObjectStreamingContentResponse.Marshal(b, m, deterministic)
}
func (m *GetObjectStreamingContentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetObjectStreamingContentResponse.Merge(m, src)
}
func (m *GetObjectStreamingContentResponse) XXX_Size() int {
	return xxx_messageInfo_GetObjectStreamingContentResponse.Size(m)
}
func (m *GetObjectStreamingContentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetObjectStreamingContentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetObjectStreamingContentResponse proto.InternalMessageInfo

func (m *GetObjectStreamingContentResponse) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*GetObjectStreamingContentRequest)(nil), "ai.metathings.service.device.GetObjectStreamingContentRequest")
	proto.RegisterType((*GetObjectStreamingContentResponse)(nil), "ai.metathings.service.device.GetObjectStreamingContentResponse")
}

func init() { proto.RegisterFile("get_object_streaming_content.proto", fileDescriptor_349f78370bb4aa14) }

var fileDescriptor_349f78370bb4aa14 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0xc7, 0x89, 0x60, 0x3d, 0xa2, 0x83, 0x74, 0xf1, 0x38, 0x1c, 0xce, 0x2e, 0x3a, 0x25, 0xa0,
	0x83, 0x93, 0x4b, 0x1d, 0xc4, 0xe9, 0xa0, 0x6e, 0x2e, 0x25, 0x4d, 0x1e, 0xb9, 0xc8, 0x25, 0xaf,
	0x36, 0xef, 0x0a, 0x7e, 0x05, 0x3f, 0xb2, 0x93, 0x98, 0xb4, 0xb8, 0x75, 0x4a, 0x1e, 0xbc, 0xff,
	0xef, 0xff, 0xe3, 0xf1, 0xca, 0x02, 0xb5, 0xd8, 0x7d, 0x80, 0xa6, 0x36, 0xd2, 0x00, 0xca, 0xbb,
	0x60, 0x5b, 0x8d, 0x81, 0x20, 0x90, 0xe8, 0x07, 0x24, 0x2c, 0xaf, 0x95, 0x13, 0x1e, 0x48, 0xd1,
	0xde, 0x05, 0x1b, 0x45, 0x84, 0x61, 0x74, 0x1a, 0x84, 0x81, 0xbf, 0x67, 0x73, 0x35, 0xaa, 0x83,
	0x33, 0x8a, 0x40, 0xce, 0x9f, 0x1c, 0xdb, 0x3c, 0x5a, 0x47, 0xfb, 0x63, 0x27, 0x34, 0x7a, 0x19,
	0xd4, 0x17, 0x12, 0x29, 0xf9, 0x8f, 0x91, 0x69, 0x49, 0x66, 0x88, 0x91, 0x1e, 0x0d, 0x1c, 0x72,
	0xb0, 0xf2, 0x7c, 0xfb, 0x02, 0xb4, 0x4b, 0x52, 0x6f, 0xb3, 0xd3, 0x73, 0x56, 0x6a, 0xe0, 0xf3,
	0x08, 0x91, 0xca, 0x57, 0x5e, 0x64, 0xeb, 0x35, 0xdb, 0xb2, 0xbb, 0xf3, 0xfb, 0x5b, 0xb1, 0x24,
	0x69, 0xc4, 0xae, 0xcf, 0xbc, 0x7a, 0xf5, 0x53, 0x9f, 0x7e, 0xb3, 0x93, 0x4b, 0xd6, 0x4c, 0x80,
	0xea, 0x89, 0xdf, 0x2c, 0xd4, 0xc5, 0x1e, 0x43, 0x84, 0x72, 0xcd, 0xcf, 0xa6, 0xa3, 0xa4, 0xc2,
	0x8b, 0x66, 0x1e, 0xeb, 0xd5, 0x7b, 0x91, 0x4b, 0xba, 0x22, 0xe9, 0x3f, 0xfc, 0x06, 0x00, 0x00,
	0xff, 0xff, 0xd8, 0x73, 0x10, 0xee, 0x54, 0x01, 0x00, 0x00,
}
