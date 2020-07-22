// Code generated by protoc-gen-go. DO NOT EDIT.
// source: add_devices_to_firmware_hub.proto

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

type AddDevicesToFirmwareHubRequest struct {
	FirmwareHub          *OpFirmwareHub `protobuf:"bytes,1,opt,name=firmware_hub,json=firmwareHub,proto3" json:"firmware_hub,omitempty"`
	Devices              []*OpDevice    `protobuf:"bytes,2,rep,name=devices,proto3" json:"devices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AddDevicesToFirmwareHubRequest) Reset()         { *m = AddDevicesToFirmwareHubRequest{} }
func (m *AddDevicesToFirmwareHubRequest) String() string { return proto.CompactTextString(m) }
func (*AddDevicesToFirmwareHubRequest) ProtoMessage()    {}
func (*AddDevicesToFirmwareHubRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0d00c7c7c4a4cff, []int{0}
}

func (m *AddDevicesToFirmwareHubRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddDevicesToFirmwareHubRequest.Unmarshal(m, b)
}
func (m *AddDevicesToFirmwareHubRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddDevicesToFirmwareHubRequest.Marshal(b, m, deterministic)
}
func (m *AddDevicesToFirmwareHubRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddDevicesToFirmwareHubRequest.Merge(m, src)
}
func (m *AddDevicesToFirmwareHubRequest) XXX_Size() int {
	return xxx_messageInfo_AddDevicesToFirmwareHubRequest.Size(m)
}
func (m *AddDevicesToFirmwareHubRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddDevicesToFirmwareHubRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddDevicesToFirmwareHubRequest proto.InternalMessageInfo

func (m *AddDevicesToFirmwareHubRequest) GetFirmwareHub() *OpFirmwareHub {
	if m != nil {
		return m.FirmwareHub
	}
	return nil
}

func (m *AddDevicesToFirmwareHubRequest) GetDevices() []*OpDevice {
	if m != nil {
		return m.Devices
	}
	return nil
}

func init() {
	proto.RegisterType((*AddDevicesToFirmwareHubRequest)(nil), "ai.metathings.service.deviced.AddDevicesToFirmwareHubRequest")
}

func init() { proto.RegisterFile("add_devices_to_firmware_hub.proto", fileDescriptor_d0d00c7c7c4a4cff) }

var fileDescriptor_d0d00c7c7c4a4cff = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0x4c, 0x49, 0x89,
	0x4f, 0x49, 0x2d, 0xcb, 0x4c, 0x4e, 0x2d, 0x8e, 0x2f, 0xc9, 0x8f, 0x4f, 0xcb, 0x2c, 0xca, 0x2d,
	0x4f, 0x2c, 0x4a, 0x8d, 0xcf, 0x28, 0x4d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4d,
	0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d,
	0x02, 0xa9, 0xd6, 0x83, 0x68, 0x4a, 0x91, 0x32, 0x4b, 0xcf, 0x2c, 0x01, 0x29, 0x4e, 0xce, 0xcf,
	0xd5, 0xcf, 0x2d, 0xcf, 0x2c, 0xc9, 0xce, 0x2f, 0xd7, 0x4f, 0xcf, 0xd7, 0x05, 0xeb, 0xd5, 0x2d,
	0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0x2a, 0xd6, 0x87, 0x33, 0x21, 0xc6, 0x4a, 0x71,
	0xe7, 0xe6, 0xa7, 0xa4, 0xe6, 0x40, 0x38, 0x4a, 0x47, 0x18, 0xb9, 0xe4, 0x1c, 0x53, 0x52, 0x5c,
	0x20, 0x0e, 0x09, 0xc9, 0x77, 0x83, 0x3a, 0xc3, 0xa3, 0x34, 0x29, 0x28, 0xb5, 0xb0, 0x34, 0xb5,
	0xb8, 0x44, 0x28, 0x9c, 0x8b, 0x07, 0xd9, 0x71, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x3a,
	0x7a, 0x78, 0x5d, 0xa7, 0xe7, 0x5f, 0x80, 0x64, 0x94, 0x13, 0xdb, 0xa3, 0xfb, 0xf2, 0x4c, 0x0a,
	0x8c, 0x41, 0xdc, 0x69, 0x08, 0x41, 0x21, 0x4f, 0x2e, 0x76, 0x68, 0x00, 0x48, 0x30, 0x29, 0x30,
	0x6b, 0x70, 0x1b, 0xa9, 0x13, 0x34, 0x13, 0xe2, 0x4e, 0xb8, 0x71, 0x30, 0xfd, 0x49, 0x6c, 0x60,
	0xdf, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xef, 0xe6, 0xf1, 0xab, 0x56, 0x01, 0x00, 0x00,
}