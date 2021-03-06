// Code generated by protoc-gen-go. DO NOT EDIT.
// source: add_devices_to_firmware_hub.proto

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
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0x4c, 0x49, 0x89,
	0x4f, 0x49, 0x2d, 0xcb, 0x4c, 0x4e, 0x2d, 0x8e, 0x2f, 0xc9, 0x8f, 0x4f, 0xcb, 0x2c, 0xca, 0x2d,
	0x4f, 0x2c, 0x4a, 0x8d, 0xcf, 0x28, 0x4d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4d,
	0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d,
	0x02, 0xa9, 0xd6, 0x83, 0x68, 0x4a, 0x91, 0x12, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49,
	0xd5, 0x87, 0x31, 0x20, 0xfa, 0xa4, 0xb8, 0x73, 0xf3, 0x53, 0x52, 0x73, 0x20, 0x1c, 0xa5, 0x13,
	0x8c, 0x5c, 0x72, 0x8e, 0x29, 0x29, 0x2e, 0x10, 0x9b, 0x42, 0xf2, 0xdd, 0xa0, 0xf6, 0x78, 0x94,
	0x26, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x45, 0x72, 0xf1, 0x20, 0xdb, 0x2e, 0xc1,
	0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xa4, 0xa3, 0x87, 0xd7, 0x7a, 0x3d, 0xff, 0x02, 0x24, 0xa3, 0x9c,
	0x38, 0x7e, 0x39, 0xb1, 0x76, 0x31, 0x32, 0x09, 0x30, 0x06, 0x71, 0xa7, 0x21, 0x84, 0x85, 0xbc,
	0xb9, 0xd8, 0xa1, 0x7e, 0x94, 0x60, 0x52, 0x60, 0xd6, 0xe0, 0x36, 0x52, 0x27, 0x68, 0x2a, 0xc4,
	0xa5, 0x60, 0x03, 0x27, 0x31, 0x32, 0x71, 0x30, 0x06, 0xc1, 0x4c, 0x48, 0x62, 0x03, 0xfb, 0xc8,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x16, 0x7a, 0x4b, 0x3b, 0x01, 0x00, 0x00,
}
