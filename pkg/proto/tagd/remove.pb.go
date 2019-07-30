// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remove.proto

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

type RemoveRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RemoveRequest) Reset()         { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()    {}
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd927fd793157ac6, []int{0}
}

func (m *RemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRequest.Unmarshal(m, b)
}
func (m *RemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRequest.Marshal(b, m, deterministic)
}
func (m *RemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRequest.Merge(m, src)
}
func (m *RemoveRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRequest.Size(m)
}
func (m *RemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRequest proto.InternalMessageInfo

func (m *RemoveRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func init() {
	proto.RegisterType((*RemoveRequest)(nil), "ai.metathings.service.tagd.RemoveRequest")
}

func init() { proto.RegisterFile("remove.proto", fileDescriptor_dd927fd793157ac6) }

var fileDescriptor_dd927fd793157ac6 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x69, 0x0f, 0x3d, 0x54, 0xbd, 0xf4, 0x24, 0x45, 0xb4, 0x78, 0xf2, 0xd2, 0x09, 0xa8,
	0xf8, 0x00, 0x82, 0x2f, 0x50, 0xc1, 0x7b, 0xda, 0x8c, 0xe9, 0xb0, 0x4d, 0xa7, 0x9b, 0x4c, 0xda,
	0xc7, 0x5d, 0xd8, 0x27, 0x59, 0x48, 0xd9, 0xbd, 0x0d, 0xcc, 0xf7, 0xf1, 0xfd, 0xe5, 0xbd, 0x47,
	0xc7, 0x2b, 0xc2, 0xe2, 0x59, 0xb8, 0xaa, 0x35, 0x81, 0x43, 0xd1, 0x32, 0xd2, 0x6c, 0x03, 0x04,
	0xf4, 0x2b, 0x0d, 0x08, 0xa2, 0xad, 0xa9, 0x9f, 0x2d, 0xb3, 0x9d, 0x50, 0x25, 0xb2, 0x8f, 0xff,
	0x6a, 0xf3, 0x7a, 0x59, 0xd0, 0x87, 0xdd, 0xad, 0xbf, 0x2c, 0xc9, 0x18, 0x7b, 0x18, 0xd8, 0x29,
	0xb7, 0x91, 0x1c, 0x78, 0x53, 0x96, 0xdb, 0xf4, 0x6c, 0x57, 0x3d, 0x91, 0xd1, 0xc2, 0x3e, 0xa8,
	0xdb, 0xb9, 0x7b, 0xaf, 0x3f, 0xe5, 0x43, 0x97, 0x36, 0x74, 0x78, 0x8c, 0x18, 0xa4, 0xfa, 0x2c,
	0x73, 0x32, 0x8f, 0x59, 0x93, 0xbd, 0xdd, 0xbd, 0x3f, 0xc1, 0x5e, 0x85, 0x6b, 0x15, 0x7e, 0xc5,
	0xd3, 0x6c, 0xff, 0xf4, 0x14, 0xf1, 0xbb, 0x38, 0x9f, 0x5e, 0xf2, 0x26, 0xeb, 0x72, 0x32, 0x7d,
	0x91, 0x88, 0x8f, 0x4b, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xaf, 0x82, 0x9e, 0xd1, 0x00, 0x00,
	0x00,
}
