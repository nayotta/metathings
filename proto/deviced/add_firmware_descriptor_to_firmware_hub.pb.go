// Code generated by protoc-gen-go. DO NOT EDIT.
// source: add_firmware_descriptor_to_firmware_hub.proto

package deviced

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type AddFirmwareDescriptorToFirmwareHubRequest struct {
	FirmwareHub          *OpFirmwareHub        `protobuf:"bytes,1,opt,name=firmware_hub,json=firmwareHub,proto3" json:"firmware_hub,omitempty"`
	FirmwareDescriptor   *OpFirmwareDescriptor `protobuf:"bytes,2,opt,name=firmware_descriptor,json=firmwareDescriptor,proto3" json:"firmware_descriptor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AddFirmwareDescriptorToFirmwareHubRequest) Reset() {
	*m = AddFirmwareDescriptorToFirmwareHubRequest{}
}
func (m *AddFirmwareDescriptorToFirmwareHubRequest) String() string { return proto.CompactTextString(m) }
func (*AddFirmwareDescriptorToFirmwareHubRequest) ProtoMessage()    {}
func (*AddFirmwareDescriptorToFirmwareHubRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ceebc584abc20cfe, []int{0}
}

func (m *AddFirmwareDescriptorToFirmwareHubRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddFirmwareDescriptorToFirmwareHubRequest.Unmarshal(m, b)
}
func (m *AddFirmwareDescriptorToFirmwareHubRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddFirmwareDescriptorToFirmwareHubRequest.Marshal(b, m, deterministic)
}
func (m *AddFirmwareDescriptorToFirmwareHubRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddFirmwareDescriptorToFirmwareHubRequest.Merge(m, src)
}
func (m *AddFirmwareDescriptorToFirmwareHubRequest) XXX_Size() int {
	return xxx_messageInfo_AddFirmwareDescriptorToFirmwareHubRequest.Size(m)
}
func (m *AddFirmwareDescriptorToFirmwareHubRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddFirmwareDescriptorToFirmwareHubRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddFirmwareDescriptorToFirmwareHubRequest proto.InternalMessageInfo

func (m *AddFirmwareDescriptorToFirmwareHubRequest) GetFirmwareHub() *OpFirmwareHub {
	if m != nil {
		return m.FirmwareHub
	}
	return nil
}

func (m *AddFirmwareDescriptorToFirmwareHubRequest) GetFirmwareDescriptor() *OpFirmwareDescriptor {
	if m != nil {
		return m.FirmwareDescriptor
	}
	return nil
}

func init() {
	proto.RegisterType((*AddFirmwareDescriptorToFirmwareHubRequest)(nil), "ai.metathings.service.deviced.AddFirmwareDescriptorToFirmwareHubRequest")
}

func init() {
	proto.RegisterFile("add_firmware_descriptor_to_firmware_hub.proto", fileDescriptor_ceebc584abc20cfe)
}

var fileDescriptor_ceebc584abc20cfe = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4d, 0x4c, 0x49, 0x89,
	0x4f, 0xcb, 0x2c, 0xca, 0x2d, 0x4f, 0x2c, 0x4a, 0x8d, 0x4f, 0x49, 0x2d, 0x4e, 0x2e, 0xca, 0x2c,
	0x28, 0xc9, 0x2f, 0x8a, 0x2f, 0xc9, 0x47, 0x08, 0x67, 0x94, 0x26, 0xe9, 0x15, 0x14, 0xe5, 0x97,
	0xe4, 0x0b, 0xc9, 0x26, 0x66, 0xea, 0xe5, 0xa6, 0x96, 0x24, 0x96, 0x64, 0x64, 0xe6, 0xa5, 0x17,
	0xeb, 0x15, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0xea, 0xa5, 0xa4, 0x82, 0xa8, 0x14, 0x29, 0xf1,
	0xb2, 0xc4, 0x9c, 0xcc, 0x94, 0xc4, 0x92, 0x54, 0x7d, 0x18, 0x03, 0xa2, 0x4f, 0x8a, 0x3b, 0x37,
	0x3f, 0x25, 0x35, 0x07, 0xc2, 0x51, 0xfa, 0xc6, 0xc8, 0xa5, 0xe9, 0x98, 0x92, 0xe2, 0x06, 0x35,
	0xde, 0x05, 0x6e, 0x69, 0x48, 0x3e, 0x4c, 0xcc, 0xa3, 0x34, 0x29, 0x28, 0xb5, 0xb0, 0x34, 0xb5,
	0xb8, 0x44, 0x28, 0x92, 0x8b, 0x07, 0xd9, 0x21, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x3a,
	0x7a, 0x78, 0x5d, 0xa2, 0xe7, 0x5f, 0x80, 0x64, 0x94, 0x13, 0xc7, 0x2f, 0x27, 0xd6, 0x2e, 0x46,
	0x26, 0x01, 0xc6, 0x20, 0xee, 0x34, 0x84, 0xb0, 0x50, 0x1e, 0x97, 0x30, 0x16, 0xaf, 0x4b, 0x30,
	0x81, 0x6d, 0x30, 0x26, 0xda, 0x06, 0x84, 0x07, 0x90, 0x2c, 0x12, 0x4a, 0xc3, 0x90, 0x4d, 0x62,
	0x03, 0xfb, 0xdf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x07, 0xcf, 0xcc, 0xa3, 0x75, 0x01, 0x00,
	0x00,
}
