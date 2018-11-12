// Code generated by protoc-gen-go. DO NOT EDIT.
// source: show.proto

package camera

import (
	fmt "fmt"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ShowResponse struct {
	Camera               *Camera  `protobuf:"bytes,1,opt,name=camera,proto3" json:"camera,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowResponse) Reset()         { *m = ShowResponse{} }
func (m *ShowResponse) String() string { return proto.CompactTextString(m) }
func (*ShowResponse) ProtoMessage()    {}
func (*ShowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_40dadfe47b432028, []int{0}
}

func (m *ShowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowResponse.Unmarshal(m, b)
}
func (m *ShowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowResponse.Marshal(b, m, deterministic)
}
func (m *ShowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowResponse.Merge(m, src)
}
func (m *ShowResponse) XXX_Size() int {
	return xxx_messageInfo_ShowResponse.Size(m)
}
func (m *ShowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShowResponse proto.InternalMessageInfo

func (m *ShowResponse) GetCamera() *Camera {
	if m != nil {
		return m.Camera
	}
	return nil
}

func init() {
	proto.RegisterType((*ShowResponse)(nil), "ai.metathings.service.camera.ShowResponse")
}

func init() { proto.RegisterFile("show.proto", fileDescriptor_40dadfe47b432028) }

var fileDescriptor_40dadfe47b432028 = []byte{
	// 117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xce, 0xc8, 0x2f,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x49, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c,
	0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4b, 0x4e,
	0xcc, 0x4d, 0x2d, 0x4a, 0x94, 0xe2, 0x81, 0xd0, 0x10, 0xb5, 0x4a, 0x3e, 0x5c, 0x3c, 0xc1, 0x19,
	0xf9, 0xe5, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x36, 0x5c, 0x6c, 0x10, 0x79,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x15, 0x3d, 0x7c, 0x86, 0xe9, 0x39, 0x83, 0xa9, 0x20,
	0xa8, 0x9e, 0x24, 0x36, 0xb0, 0xa1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xec, 0x6e, 0x23,
	0xad, 0x8e, 0x00, 0x00, 0x00,
}
