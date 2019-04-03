// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_object_streaming_content.proto

package deviced

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

type GetObjectStreamingContentRequest struct {
	Object               *OpObject `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
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

func (m *GetObjectStreamingContentRequest) GetObject() *OpObject {
	if m != nil {
		return m.Object
	}
	return nil
}

type GetObjectStreamingContentResponse struct {
	Content              *wrappers.BytesValue `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
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

func (m *GetObjectStreamingContentResponse) GetContent() *wrappers.BytesValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*GetObjectStreamingContentRequest)(nil), "ai.metathings.service.deviced.GetObjectStreamingContentRequest")
	proto.RegisterType((*GetObjectStreamingContentResponse)(nil), "ai.metathings.service.deviced.GetObjectStreamingContentResponse")
}

func init() { proto.RegisterFile("get_object_streaming_content.proto", fileDescriptor_349f78370bb4aa14) }

var fileDescriptor_349f78370bb4aa14 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x15, 0x86, 0x20, 0xa5, 0x5b, 0x26, 0x54, 0x04, 0x84, 0x2c, 0xb0, 0xf4, 0x22, 0x81,
	0xe0, 0x07, 0x94, 0xa1, 0x63, 0xa5, 0x20, 0x31, 0xb0, 0x44, 0x4e, 0x72, 0xb8, 0xa6, 0xb1, 0xcf,
	0xd8, 0x97, 0x46, 0xfc, 0x5a, 0x24, 0x7e, 0x09, 0x22, 0x4e, 0x3a, 0x76, 0xb2, 0xad, 0x7b, 0xcf,
	0xef, 0xbb, 0x97, 0xe4, 0x12, 0xb9, 0xa2, 0xfa, 0x13, 0x1b, 0xae, 0x3c, 0x3b, 0x14, 0x5a, 0x19,
	0x59, 0x35, 0x64, 0x18, 0x0d, 0x83, 0x75, 0xc4, 0x94, 0x5e, 0x09, 0x05, 0x1a, 0x59, 0xf0, 0x4e,
	0x19, 0xe9, 0xc1, 0xa3, 0x3b, 0xa8, 0x06, 0xa1, 0xc5, 0xff, 0xa3, 0x5d, 0x5e, 0x4b, 0x22, 0xd9,
	0x61, 0x31, 0x8a, 0xeb, 0xfe, 0xa3, 0x18, 0x9c, 0xb0, 0x16, 0x9d, 0x0f, 0xf6, 0xe5, 0xb3, 0x54,
	0xbc, 0xeb, 0x6b, 0x68, 0x48, 0x17, 0x7a, 0x50, 0xbc, 0xa7, 0xa1, 0x90, 0xb4, 0x1a, 0x87, 0xab,
	0x83, 0xe8, 0x54, 0x2b, 0x98, 0x9c, 0x2f, 0x8e, 0xd7, 0xc9, 0xb7, 0xd0, 0xd4, 0x62, 0x17, 0x1e,
	0xf9, 0x3e, 0xc9, 0x36, 0xc8, 0xdb, 0x11, 0xf4, 0x75, 0xe6, 0x7c, 0x09, 0x98, 0x25, 0x7e, 0xf5,
	0xe8, 0x39, 0xdd, 0x24, 0x71, 0xd8, 0xe4, 0x22, 0xca, 0xa2, 0xfb, 0xc5, 0xc3, 0x1d, 0x9c, 0x04,
	0x87, 0xad, 0x0d, 0xff, 0xad, 0xe3, 0xdf, 0x9f, 0x9b, 0xb3, 0x2c, 0x2a, 0x27, 0x7b, 0xfe, 0x9e,
	0xdc, 0x9e, 0x08, 0xf3, 0x96, 0x8c, 0xc7, 0xf4, 0x29, 0x39, 0x9f, 0x6a, 0x9a, 0xe2, 0x2e, 0x21,
	0x14, 0x01, 0x73, 0x11, 0xb0, 0xfe, 0x66, 0xf4, 0x6f, 0xa2, 0xeb, 0xb1, 0x9c, 0xb5, 0x75, 0x3c,
	0x4e, 0x1f, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x57, 0x18, 0xc3, 0xfa, 0x79, 0x01, 0x00, 0x00,
}
