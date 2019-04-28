// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_object_content.proto

package ai_metathings_service_device

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
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
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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

func (m *GetObjectContentResponse) GetContent() []byte {
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
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x19, 0x17, 0x23, 0x44, 0x57, 0xb3, 0x71, 0x28, 0xa2, 0xa5, 0x1b, 0xdd, 0x34, 0x01,
	0x15, 0x37, 0xee, 0x74, 0xd1, 0x65, 0x61, 0x5e, 0xa0, 0x24, 0x93, 0x6b, 0x1a, 0x3b, 0x93, 0x1b,
	0x93, 0x3b, 0x1d, 0x7c, 0x5a, 0xc1, 0x27, 0x11, 0x93, 0xf1, 0x07, 0x84, 0xae, 0x92, 0x90, 0x7c,
	0x39, 0xe7, 0xbb, 0xac, 0x36, 0x40, 0x1b, 0x54, 0x2f, 0xd0, 0xd2, 0xa6, 0x45, 0x47, 0xe0, 0x88,
	0xfb, 0x80, 0x84, 0xd5, 0xb9, 0xb4, 0xbc, 0x07, 0x92, 0xb4, 0xb5, 0xce, 0x44, 0x1e, 0x21, 0xec,
	0x6d, 0x0b, 0x5c, 0xc3, 0xd7, 0x32, 0xbb, 0x30, 0x88, 0xa6, 0x03, 0x91, 0xde, 0xaa, 0xe1, 0x59,
	0x8c, 0x41, 0x7a, 0x0f, 0x21, 0x66, 0x7a, 0x76, 0x6f, 0x2c, 0x6d, 0x07, 0xc5, 0x5b, 0xec, 0x45,
	0x3f, 0x5a, 0xda, 0xe1, 0x28, 0x0c, 0x2e, 0xd3, 0xe5, 0x72, 0x2f, 0x3b, 0xab, 0x25, 0x61, 0x88,
	0xe2, 0x67, 0x3b, 0x71, 0x0f, 0x7f, 0x38, 0x27, 0xdf, 0x90, 0x48, 0x8a, 0xdf, 0x16, 0xc2, 0xef,
	0x4c, 0x8e, 0x14, 0xb9, 0x87, 0x16, 0x3d, 0x6a, 0xe8, 0x32, 0xbc, 0x50, 0xec, 0x6c, 0x05, 0xb4,
	0x4e, 0x36, 0x4f, 0x59, 0xa6, 0x81, 0xd7, 0x01, 0x22, 0x55, 0x2b, 0x56, 0x66, 0xcb, 0xba, 0x98,
	0x17, 0xd7, 0x27, 0x37, 0x57, 0xfc, 0x90, 0x9e, 0xe6, 0x6b, 0x9f, 0xbf, 0x79, 0x2c, 0x3f, 0xde,
	0x2f, 0x8f, 0xe6, 0x45, 0x33, 0xe1, 0x8b, 0x3b, 0x56, 0xff, 0xcf, 0x88, 0x1e, 0x5d, 0x84, 0xaa,
	0x66, 0xc7, 0xd3, 0x0c, 0x53, 0xca, 0x69, 0xf3, 0x7d, 0x54, 0x65, 0x2a, 0x78, 0xfb, 0x19, 0x00,
	0x00, 0xff, 0xff, 0x90, 0x38, 0x60, 0x70, 0x6f, 0x01, 0x00, 0x00,
}
