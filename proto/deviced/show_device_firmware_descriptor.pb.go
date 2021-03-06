// Code generated by protoc-gen-go. DO NOT EDIT.
// source: show_device_firmware_descriptor.proto

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

type ShowDeviceFirmwareDescriptorResponse struct {
	FirmwareDescriptor   *FirmwareDescriptor `protobuf:"bytes,1,opt,name=firmware_descriptor,json=firmwareDescriptor,proto3" json:"firmware_descriptor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ShowDeviceFirmwareDescriptorResponse) Reset()         { *m = ShowDeviceFirmwareDescriptorResponse{} }
func (m *ShowDeviceFirmwareDescriptorResponse) String() string { return proto.CompactTextString(m) }
func (*ShowDeviceFirmwareDescriptorResponse) ProtoMessage()    {}
func (*ShowDeviceFirmwareDescriptorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_76b04338a28a6433, []int{0}
}

func (m *ShowDeviceFirmwareDescriptorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowDeviceFirmwareDescriptorResponse.Unmarshal(m, b)
}
func (m *ShowDeviceFirmwareDescriptorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowDeviceFirmwareDescriptorResponse.Marshal(b, m, deterministic)
}
func (m *ShowDeviceFirmwareDescriptorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowDeviceFirmwareDescriptorResponse.Merge(m, src)
}
func (m *ShowDeviceFirmwareDescriptorResponse) XXX_Size() int {
	return xxx_messageInfo_ShowDeviceFirmwareDescriptorResponse.Size(m)
}
func (m *ShowDeviceFirmwareDescriptorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowDeviceFirmwareDescriptorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShowDeviceFirmwareDescriptorResponse proto.InternalMessageInfo

func (m *ShowDeviceFirmwareDescriptorResponse) GetFirmwareDescriptor() *FirmwareDescriptor {
	if m != nil {
		return m.FirmwareDescriptor
	}
	return nil
}

func init() {
	proto.RegisterType((*ShowDeviceFirmwareDescriptorResponse)(nil), "ai.metathings.service.deviced.ShowDeviceFirmwareDescriptorResponse")
}

func init() {
	proto.RegisterFile("show_device_firmware_descriptor.proto", fileDescriptor_76b04338a28a6433)
}

var fileDescriptor_76b04338a28a6433 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0xce, 0xc8, 0x2f,
	0x8f, 0x4f, 0x49, 0x2d, 0xcb, 0x4c, 0x4e, 0x8d, 0x4f, 0xcb, 0x2c, 0xca, 0x2d, 0x4f, 0x2c, 0x4a,
	0x8d, 0x4f, 0x49, 0x2d, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0x4d, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f,
	0xd6, 0x2b, 0x4e, 0x2d, 0x02, 0x69, 0xd0, 0x83, 0xe8, 0x4b, 0x91, 0xe2, 0xce, 0xcd, 0x4f, 0x49,
	0xcd, 0x81, 0xa8, 0x55, 0xea, 0x62, 0xe4, 0x52, 0x09, 0xce, 0xc8, 0x2f, 0x77, 0x01, 0x4b, 0xba,
	0x41, 0xcd, 0x74, 0x81, 0x1b, 0x19, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x94, 0xc4,
	0x25, 0x8c, 0xc5, 0x46, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x43, 0x3d, 0xbc, 0x56, 0xea,
	0x61, 0x31, 0x57, 0x28, 0x0d, 0x43, 0xcc, 0x89, 0x33, 0x8a, 0x1d, 0xaa, 0x23, 0x89, 0x0d, 0xec,
	0x3c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x51, 0x00, 0x01, 0x50, 0xf3, 0x00, 0x00, 0x00,
}
