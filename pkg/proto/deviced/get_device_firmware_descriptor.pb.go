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
	FirmwareDescriptor   *FirmwareDescriptor `protobuf:"bytes,1,opt,name=firmware_descriptor,json=firmwareDescriptor,proto3" json:"firmware_descriptor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
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

func (m *GetDeviceFirmwareDescriptorResponse) GetFirmwareDescriptor() *FirmwareDescriptor {
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
	0x14, 0x86, 0x89, 0x8b, 0x59, 0xa4, 0xbb, 0xb8, 0x91, 0x82, 0x58, 0x46, 0x41, 0x37, 0xcd, 0xa0,
	0x82, 0x07, 0x90, 0x62, 0x97, 0x42, 0x2f, 0x30, 0x64, 0x26, 0xaf, 0xe9, 0xc3, 0xa6, 0x2f, 0x26,
	0xaf, 0x9d, 0x33, 0x78, 0x4a, 0xc1, 0x93, 0x88, 0x93, 0xa1, 0x1b, 0xcb, 0x74, 0x97, 0x3c, 0xf8,
	0xbf, 0xef, 0xe7, 0x97, 0x77, 0x0e, 0xb8, 0xb6, 0x70, 0xc0, 0x16, 0xea, 0x35, 0x46, 0xdf, 0x99,
	0x08, 0xb5, 0x85, 0xd4, 0x46, 0x0c, 0x4c, 0x51, 0x87, 0x48, 0x4c, 0xea, 0xda, 0xa0, 0xf6, 0xc0,
	0x86, 0x37, 0xb8, 0x73, 0x49, 0x27, 0x88, 0x7f, 0x01, 0x9d, 0x73, 0x76, 0xfa, 0xe2, 0x90, 0x37,
	0xfb, 0x46, 0xb7, 0xe4, 0x2b, 0xdf, 0x21, 0x7f, 0x50, 0x57, 0x39, 0x9a, 0xf7, 0xd9, 0xf9, 0xc1,
	0x6c, 0xd1, 0x1a, 0xa6, 0x98, 0xaa, 0xe3, 0x33, 0x63, 0xa7, 0x13, 0x4f, 0x16, 0xb6, 0xf9, 0x53,
	0x7a, 0x59, 0x2e, 0x81, 0x17, 0x3d, 0xf2, 0x6d, 0x68, 0xb2, 0x38, 0x16, 0x59, 0xc1, 0xe7, 0x1e,
	0x12, 0xab, 0xa5, 0x2c, 0xb2, 0xf5, 0x4a, 0xcc, 0xc4, 0xc3, 0xe4, 0xe9, 0x5e, 0x8f, 0x56, 0xd3,
	0xef, 0x21, 0x13, 0x5f, 0x8b, 0x9f, 0xef, 0x9b, 0x8b, 0x99, 0x58, 0x0d, 0xf1, 0xf2, 0x4b, 0xc8,
	0xdb, 0x51, 0x5f, 0x0a, 0xb4, 0x4b, 0xa0, 0x1a, 0x79, 0x79, 0x62, 0x97, 0xc1, 0xfe, 0x78, 0xc6,
	0x7e, 0x82, 0xab, 0xd6, 0xff, 0x6e, 0x4d, 0xd1, 0x2f, 0xf0, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0xee, 0x6a, 0x9b, 0xd6, 0x8d, 0x01, 0x00, 0x00,
}
