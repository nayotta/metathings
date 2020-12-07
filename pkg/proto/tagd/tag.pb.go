// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tag.proto

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TagRequest struct {
	Id                   *wrappers.StringValue   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tags                 []*wrappers.StringValue `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	Namespace            *wrappers.StringValue   `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *TagRequest) Reset()         { *m = TagRequest{} }
func (m *TagRequest) String() string { return proto.CompactTextString(m) }
func (*TagRequest) ProtoMessage()    {}
func (*TagRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27f545bcde37ecb5, []int{0}
}

func (m *TagRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagRequest.Unmarshal(m, b)
}
func (m *TagRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagRequest.Marshal(b, m, deterministic)
}
func (m *TagRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagRequest.Merge(m, src)
}
func (m *TagRequest) XXX_Size() int {
	return xxx_messageInfo_TagRequest.Size(m)
}
func (m *TagRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TagRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TagRequest proto.InternalMessageInfo

func (m *TagRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *TagRequest) GetTags() []*wrappers.StringValue {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *TagRequest) GetNamespace() *wrappers.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func init() {
	proto.RegisterType((*TagRequest)(nil), "ai.metathings.service.tagd.TagRequest")
}

func init() { proto.RegisterFile("tag.proto", fileDescriptor_27f545bcde37ecb5) }

var fileDescriptor_27f545bcde37ecb5 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8d, 0x41, 0x4a, 0xc4, 0x40,
	0x10, 0x45, 0x49, 0x46, 0x06, 0xa6, 0xdd, 0x65, 0x15, 0x82, 0x68, 0x70, 0xe5, 0x66, 0xba, 0x45,
	0xc5, 0x03, 0xcc, 0x11, 0xa2, 0xb8, 0xaf, 0x24, 0x65, 0x4d, 0x61, 0x92, 0x8a, 0xdd, 0x95, 0xe4,
	0x6a, 0xde, 0x46, 0xf0, 0x24, 0x42, 0x07, 0x75, 0x29, 0xb3, 0xfb, 0xf0, 0xdf, 0xe3, 0x99, 0x9d,
	0x02, 0xd9, 0xd1, 0x8b, 0x4a, 0x56, 0x00, 0xdb, 0x1e, 0x15, 0xf4, 0xc8, 0x03, 0x05, 0x1b, 0xd0,
	0xcf, 0xdc, 0xa0, 0x55, 0xa0, 0xb6, 0xb8, 0x24, 0x11, 0xea, 0xd0, 0x45, 0xb2, 0x9e, 0x5e, 0xdd,
	0xe2, 0x61, 0x1c, 0xd1, 0x87, 0xd5, 0x2d, 0x1e, 0x89, 0xf5, 0x38, 0xd5, 0xb6, 0x91, 0xde, 0xf5,
	0x0b, 0xeb, 0x9b, 0x2c, 0x8e, 0x64, 0x1f, 0xcf, 0xfd, 0x0c, 0x1d, 0xb7, 0xa0, 0xe2, 0x83, 0xfb,
	0x9d, 0xab, 0x77, 0xfd, 0x91, 0x18, 0xf3, 0x0c, 0x54, 0xe1, 0xfb, 0x84, 0x41, 0xb3, 0x07, 0x93,
	0x72, 0x9b, 0x27, 0x65, 0x72, 0x73, 0x7e, 0x77, 0x61, 0xd7, 0xa6, 0xfd, 0x69, 0xda, 0x27, 0xf5,
	0x3c, 0xd0, 0x0b, 0x74, 0x13, 0x1e, 0xb6, 0x5f, 0x9f, 0x57, 0x69, 0x99, 0x54, 0x29, 0xb7, 0xd9,
	0xad, 0x39, 0x53, 0xa0, 0x90, 0xa7, 0xe5, 0xe6, 0x3f, 0xaf, 0x8a, 0x64, 0x76, 0x30, 0xbb, 0x01,
	0x7a, 0x0c, 0x23, 0x34, 0x98, 0x6f, 0x4e, 0xc8, 0xfd, 0x69, 0xf5, 0x36, 0x82, 0xf7, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x9d, 0xfa, 0xe8, 0x1b, 0x42, 0x01, 0x00, 0x00,
}