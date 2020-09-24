// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delete_firmware_hub.proto

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

type DeleteFirmwareHubRequest struct {
	FirmwareHub          *OpFirmwareHub `protobuf:"bytes,1,opt,name=firmware_hub,json=firmwareHub,proto3" json:"firmware_hub,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeleteFirmwareHubRequest) Reset()         { *m = DeleteFirmwareHubRequest{} }
func (m *DeleteFirmwareHubRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteFirmwareHubRequest) ProtoMessage()    {}
func (*DeleteFirmwareHubRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ea4ff006bea90aa, []int{0}
}

func (m *DeleteFirmwareHubRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteFirmwareHubRequest.Unmarshal(m, b)
}
func (m *DeleteFirmwareHubRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteFirmwareHubRequest.Marshal(b, m, deterministic)
}
func (m *DeleteFirmwareHubRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteFirmwareHubRequest.Merge(m, src)
}
func (m *DeleteFirmwareHubRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteFirmwareHubRequest.Size(m)
}
func (m *DeleteFirmwareHubRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteFirmwareHubRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteFirmwareHubRequest proto.InternalMessageInfo

func (m *DeleteFirmwareHubRequest) GetFirmwareHub() *OpFirmwareHub {
	if m != nil {
		return m.FirmwareHub
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteFirmwareHubRequest)(nil), "ai.metathings.service.deviced.DeleteFirmwareHubRequest")
}

func init() { proto.RegisterFile("delete_firmware_hub.proto", fileDescriptor_5ea4ff006bea90aa) }

var fileDescriptor_5ea4ff006bea90aa = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x49, 0xcd, 0x49,
	0x2d, 0x49, 0x8d, 0x4f, 0xcb, 0x2c, 0xca, 0x2d, 0x4f, 0x2c, 0x4a, 0x8d, 0xcf, 0x28, 0x4d, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4d, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9,
	0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4b, 0x49, 0x05,
	0x51, 0x29, 0x52, 0xe2, 0x65, 0x89, 0x39, 0x99, 0x29, 0x89, 0x25, 0xa9, 0xfa, 0x30, 0x06, 0x44,
	0x9f, 0x14, 0x77, 0x6e, 0x7e, 0x4a, 0x6a, 0x0e, 0x84, 0xa3, 0x54, 0xca, 0x25, 0xe1, 0x02, 0xb6,
	0xc1, 0x0d, 0x6a, 0x81, 0x47, 0x69, 0x52, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x50, 0x24,
	0x17, 0x0f, 0xb2, 0xb5, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x3a, 0x7a, 0x78, 0xed, 0xd5,
	0xf3, 0x2f, 0x40, 0x32, 0xca, 0x89, 0xe3, 0x97, 0x13, 0x6b, 0x17, 0x23, 0x93, 0x00, 0x63, 0x10,
	0x77, 0x1a, 0x42, 0x38, 0x89, 0x0d, 0x6c, 0xbb, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x02, 0x86,
	0x85, 0x6f, 0xdf, 0x00, 0x00, 0x00,
}
