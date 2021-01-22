// Code generated by protoc-gen-go. DO NOT EDIT.
// source: show_device.proto

package deviced

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ShowDeviceResponse struct {
	Device               *Device  `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowDeviceResponse) Reset()         { *m = ShowDeviceResponse{} }
func (m *ShowDeviceResponse) String() string { return proto.CompactTextString(m) }
func (*ShowDeviceResponse) ProtoMessage()    {}
func (*ShowDeviceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfb5e9d344683f4a, []int{0}
}

func (m *ShowDeviceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowDeviceResponse.Unmarshal(m, b)
}
func (m *ShowDeviceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowDeviceResponse.Marshal(b, m, deterministic)
}
func (m *ShowDeviceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowDeviceResponse.Merge(m, src)
}
func (m *ShowDeviceResponse) XXX_Size() int {
	return xxx_messageInfo_ShowDeviceResponse.Size(m)
}
func (m *ShowDeviceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowDeviceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShowDeviceResponse proto.InternalMessageInfo

func (m *ShowDeviceResponse) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func init() {
	proto.RegisterType((*ShowDeviceResponse)(nil), "ai.metathings.service.deviced.ShowDeviceResponse")
}

func init() { proto.RegisterFile("show_device.proto", fileDescriptor_dfb5e9d344683f4a) }

var fileDescriptor_dfb5e9d344683f4a = []byte{
	// 122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xce, 0xc8, 0x2f,
	0x8f, 0x4f, 0x49, 0x2d, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4d,
	0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d,
	0x02, 0x4b, 0x42, 0xd4, 0xa4, 0x48, 0x71, 0xe7, 0xe6, 0xa7, 0xa4, 0xe6, 0x40, 0xd4, 0x2a, 0x05,
	0x73, 0x09, 0x05, 0x67, 0xe4, 0x97, 0xbb, 0x80, 0xe5, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a,
	0x53, 0x85, 0x6c, 0xb9, 0xd8, 0x20, 0xaa, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x54, 0xf5,
	0xf0, 0x1a, 0xa9, 0x07, 0xd5, 0x0e, 0xd5, 0x94, 0xc4, 0x06, 0x36, 0xdb, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0xbe, 0xe1, 0x82, 0x30, 0x9c, 0x00, 0x00, 0x00,
}