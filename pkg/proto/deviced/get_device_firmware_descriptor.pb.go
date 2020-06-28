// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_device_firmware_descriptor.proto

package deviced

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type GetDeviceFirmwareDescriptorRequest struct {
	Device               *OpDevice `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetDeviceFirmwareDescriptorRequest) Reset()         { *m = GetDeviceFirmwareDescriptorRequest{} }
func (m *GetDeviceFirmwareDescriptorRequest) String() string { return proto.CompactTextString(m) }
func (*GetDeviceFirmwareDescriptorRequest) ProtoMessage()    {}
func (*GetDeviceFirmwareDescriptorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_06ae3a6bf7ca5278, []int{0}
}

func (m *GetDeviceFirmwareDescriptorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceFirmwareDescriptorRequest.Unmarshal(m, b)
}
func (m *GetDeviceFirmwareDescriptorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceFirmwareDescriptorRequest.Marshal(b, m, deterministic)
}
func (m *GetDeviceFirmwareDescriptorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceFirmwareDescriptorRequest.Merge(m, src)
}
func (m *GetDeviceFirmwareDescriptorRequest) XXX_Size() int {
	return xxx_messageInfo_GetDeviceFirmwareDescriptorRequest.Size(m)
}
func (m *GetDeviceFirmwareDescriptorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceFirmwareDescriptorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceFirmwareDescriptorRequest proto.InternalMessageInfo

func (m *GetDeviceFirmwareDescriptorRequest) GetDevice() *OpDevice {
	if m != nil {
		return m.Device
	}
	return nil
}

type GetDeviceFirmwareDescriptorResponse struct {
	FirmwareDescriptor   *OpFirmwareDescriptor `protobuf:"bytes,1,opt,name=firmware_descriptor,json=firmwareDescriptor,proto3" json:"firmware_descriptor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetDeviceFirmwareDescriptorResponse) Reset()         { *m = GetDeviceFirmwareDescriptorResponse{} }
func (m *GetDeviceFirmwareDescriptorResponse) String() string { return proto.CompactTextString(m) }
func (*GetDeviceFirmwareDescriptorResponse) ProtoMessage()    {}
func (*GetDeviceFirmwareDescriptorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_06ae3a6bf7ca5278, []int{1}
}

func (m *GetDeviceFirmwareDescriptorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceFirmwareDescriptorResponse.Unmarshal(m, b)
}
func (m *GetDeviceFirmwareDescriptorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceFirmwareDescriptorResponse.Marshal(b, m, deterministic)
}
func (m *GetDeviceFirmwareDescriptorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceFirmwareDescriptorResponse.Merge(m, src)
}
func (m *GetDeviceFirmwareDescriptorResponse) XXX_Size() int {
	return xxx_messageInfo_GetDeviceFirmwareDescriptorResponse.Size(m)
}
func (m *GetDeviceFirmwareDescriptorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceFirmwareDescriptorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceFirmwareDescriptorResponse proto.InternalMessageInfo

func (m *GetDeviceFirmwareDescriptorResponse) GetFirmwareDescriptor() *OpFirmwareDescriptor {
	if m != nil {
		return m.FirmwareDescriptor
	}
	return nil
}

func init() {
	proto.RegisterType((*GetDeviceFirmwareDescriptorRequest)(nil), "ai.metathings.service.deviced.GetDeviceFirmwareDescriptorRequest")
	proto.RegisterType((*GetDeviceFirmwareDescriptorResponse)(nil), "ai.metathings.service.deviced.GetDeviceFirmwareDescriptorResponse")
}

func init() {
	proto.RegisterFile("get_device_firmware_descriptor.proto", fileDescriptor_06ae3a6bf7ca5278)
}

var fileDescriptor_06ae3a6bf7ca5278 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4a, 0x03, 0x31,
	0x14, 0x86, 0x89, 0x8b, 0x59, 0xa4, 0xbb, 0xb8, 0x91, 0x82, 0x58, 0x46, 0x41, 0x37, 0xcd, 0x80,
	0x05, 0x0f, 0x20, 0xc5, 0x2e, 0x85, 0x5e, 0x60, 0x48, 0x27, 0xaf, 0xe9, 0xc3, 0xa6, 0x2f, 0x26,
	0xaf, 0x9d, 0x43, 0x78, 0x48, 0xc1, 0x93, 0x88, 0x93, 0xa1, 0x1b, 0x87, 0xe9, 0x2e, 0x79, 0xf0,
	0x7f, 0xdf, 0xcf, 0x2f, 0x1f, 0x1c, 0x70, 0x6d, 0xe1, 0x84, 0x0d, 0xd4, 0x5b, 0x8c, 0xbe, 0x35,
	0x11, 0x6a, 0x0b, 0xa9, 0x89, 0x18, 0x98, 0xa2, 0x0e, 0x91, 0x98, 0xd4, 0xad, 0x41, 0xed, 0x81,
	0x0d, 0xef, 0xf0, 0xe0, 0x92, 0x4e, 0x10, 0xff, 0x02, 0x3a, 0xe7, 0xec, 0xf4, 0xc5, 0x21, 0xef,
	0x8e, 0x1b, 0xdd, 0x90, 0xaf, 0x7c, 0x8b, 0xfc, 0x41, 0x6d, 0xe5, 0x68, 0xde, 0x65, 0xe7, 0x27,
	0xb3, 0x47, 0x6b, 0x98, 0x62, 0xaa, 0xce, 0xcf, 0x8c, 0x9d, 0x4e, 0x3c, 0x59, 0xd8, 0xe7, 0x4f,
	0xe9, 0x65, 0xb9, 0x02, 0x5e, 0x76, 0xc8, 0xb7, 0xbe, 0xc9, 0xf2, 0x5c, 0x64, 0x0d, 0x9f, 0x47,
	0x48, 0xac, 0x56, 0xb2, 0xc8, 0xd6, 0x1b, 0x31, 0x13, 0x4f, 0x93, 0xe7, 0x47, 0x3d, 0x5a, 0x4d,
	0xbf, 0x87, 0x4c, 0x7c, 0x2d, 0x7e, 0xbe, 0xef, 0xae, 0x66, 0x62, 0xdd, 0xc7, 0xcb, 0x2f, 0x21,
	0xef, 0x47, 0x7d, 0x29, 0xd0, 0x21, 0x81, 0xb2, 0xf2, 0x7a, 0x60, 0x97, 0xde, 0xbe, 0xb8, 0x68,
	0x1f, 0x20, 0xab, 0xed, 0xbf, 0xdb, 0xa6, 0xe8, 0x36, 0x58, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x76, 0xa4, 0x37, 0xff, 0x8f, 0x01, 0x00, 0x00,
}
