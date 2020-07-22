// Code generated by protoc-gen-go. DO NOT EDIT.
// source: add_firmware_descriptor_to_firmware_hub.proto

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
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4d, 0x4c, 0x49, 0x89,
	0x4f, 0xcb, 0x2c, 0xca, 0x2d, 0x4f, 0x2c, 0x4a, 0x8d, 0x4f, 0x49, 0x2d, 0x4e, 0x2e, 0xca, 0x2c,
	0x28, 0xc9, 0x2f, 0x8a, 0x2f, 0xc9, 0x47, 0x08, 0x67, 0x94, 0x26, 0xe9, 0x15, 0x14, 0xe5, 0x97,
	0xe4, 0x0b, 0xc9, 0x26, 0x66, 0xea, 0xe5, 0xa6, 0x96, 0x24, 0x96, 0x64, 0x64, 0xe6, 0xa5, 0x17,
	0xeb, 0x15, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0xea, 0xa5, 0xa4, 0x82, 0xa8, 0x14, 0x29, 0xb3,
	0xf4, 0xcc, 0x12, 0x90, 0xe2, 0xe4, 0xfc, 0x5c, 0xfd, 0xdc, 0xf2, 0xcc, 0x92, 0xec, 0xfc, 0x72,
	0xfd, 0xf4, 0x7c, 0x5d, 0xb0, 0x5e, 0xdd, 0xb2, 0xc4, 0x9c, 0xcc, 0x94, 0xc4, 0x92, 0xfc, 0xa2,
	0x62, 0x7d, 0x38, 0x13, 0x62, 0xac, 0x14, 0x77, 0x6e, 0x7e, 0x4a, 0x6a, 0x0e, 0x84, 0xa3, 0xf4,
	0x89, 0x91, 0x4b, 0xd3, 0x31, 0x25, 0xc5, 0x0d, 0x6a, 0xbb, 0x0b, 0xdc, 0x4d, 0x21, 0xf9, 0x30,
	0x31, 0x8f, 0xd2, 0xa4, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0xa1, 0x70, 0x2e, 0x1e, 0x64,
	0x77, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0xe9, 0xe8, 0xe1, 0x75, 0xa8, 0x9e, 0x7f, 0x01,
	0x92, 0x51, 0x4e, 0x6c, 0x8f, 0xee, 0xcb, 0x33, 0x29, 0x30, 0x06, 0x71, 0xa7, 0x21, 0x04, 0x85,
	0x72, 0xb8, 0x84, 0xb1, 0x84, 0x8b, 0x04, 0x13, 0xd8, 0x7c, 0x63, 0xa2, 0xcd, 0x47, 0x38, 0x1f,
	0x6e, 0x8d, 0x50, 0x1a, 0x86, 0x5c, 0x12, 0x1b, 0xd8, 0xef, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x37, 0xe3, 0xca, 0xaf, 0x90, 0x01, 0x00, 0x00,
}