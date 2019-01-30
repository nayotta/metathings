// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get.proto

package tagd

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

type GetRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

type GetResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tags                 []string `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{1}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetResponse) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "ai.metathings.service.tagd.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "ai.metathings.service.tagd.GetResponse")
}

func init() { proto.RegisterFile("get.proto", fileDescriptor_21b2a6be0e6d8388) }

var fileDescriptor_21b2a6be0e6d8388 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8d, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x49, 0x94, 0x42, 0xb6, 0xe0, 0x61, 0x4f, 0x25, 0x88, 0x86, 0x9e, 0x7a, 0xe9, 0x2e,
	0xfe, 0xc1, 0x0f, 0xd0, 0x8b, 0xf7, 0x08, 0xde, 0x37, 0xcd, 0x38, 0x1d, 0x4c, 0x32, 0x71, 0x67,
	0xd2, 0x7c, 0x5c, 0xc1, 0x4f, 0x22, 0x6e, 0xb0, 0xb7, 0x07, 0xef, 0xbd, 0xdf, 0xcf, 0x14, 0x08,
	0xea, 0xc6, 0xc8, 0xca, 0xb6, 0x0c, 0xe4, 0x7a, 0xd0, 0xa0, 0x27, 0x1a, 0x50, 0x9c, 0x40, 0x3c,
	0xd3, 0x11, 0x9c, 0x06, 0x6c, 0xcb, 0x3b, 0x64, 0xc6, 0x0e, 0x7c, 0x5a, 0x36, 0xd3, 0x87, 0x9f,
	0x63, 0x18, 0x47, 0x88, 0xb2, 0x7c, 0xcb, 0x17, 0x24, 0x3d, 0x4d, 0x8d, 0x3b, 0x72, 0xef, 0xfb,
	0x99, 0xf4, 0x93, 0x67, 0x8f, 0xbc, 0x4f, 0xe5, 0xfe, 0x1c, 0x3a, 0x6a, 0x83, 0x72, 0x14, 0x7f,
	0x89, 0xcb, 0x6f, 0x7b, 0x30, 0xe6, 0x15, 0xb4, 0x86, 0xaf, 0x09, 0x44, 0xed, 0xb3, 0xc9, 0xa9,
	0xdd, 0x64, 0x55, 0xb6, 0x5b, 0x3f, 0xde, 0xba, 0x45, 0xe9, 0xfe, 0x95, 0xee, 0x4d, 0x23, 0x0d,
	0xf8, 0x1e, 0xba, 0x09, 0x0e, 0xab, 0x9f, 0xef, 0xfb, 0xbc, 0xca, 0xea, 0x9c, 0xda, 0xed, 0x83,
	0x59, 0x27, 0x86, 0x8c, 0x3c, 0x08, 0xd8, 0x9b, 0x0b, 0xa4, 0xf8, 0xab, 0xad, 0x35, 0xd7, 0x1a,
	0x50, 0x36, 0x79, 0x75, 0xb5, 0x2b, 0xea, 0x94, 0x9b, 0x55, 0x82, 0x3e, 0xfd, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x4d, 0x98, 0x2f, 0xbe, 0xfe, 0x00, 0x00, 0x00,
}
