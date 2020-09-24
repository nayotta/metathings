// Code generated by protoc-gen-go. DO NOT EDIT.
// source: patch_device.proto

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

type PatchDeviceRequest struct {
	Device               *OpDevice `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PatchDeviceRequest) Reset()         { *m = PatchDeviceRequest{} }
func (m *PatchDeviceRequest) String() string { return proto.CompactTextString(m) }
func (*PatchDeviceRequest) ProtoMessage()    {}
func (*PatchDeviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c6fcaf6b9150132, []int{0}
}

func (m *PatchDeviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchDeviceRequest.Unmarshal(m, b)
}
func (m *PatchDeviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchDeviceRequest.Marshal(b, m, deterministic)
}
func (m *PatchDeviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchDeviceRequest.Merge(m, src)
}
func (m *PatchDeviceRequest) XXX_Size() int {
	return xxx_messageInfo_PatchDeviceRequest.Size(m)
}
func (m *PatchDeviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchDeviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PatchDeviceRequest proto.InternalMessageInfo

func (m *PatchDeviceRequest) GetDevice() *OpDevice {
	if m != nil {
		return m.Device
	}
	return nil
}

type PatchDeviceResponse struct {
	Device               *Device  `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PatchDeviceResponse) Reset()         { *m = PatchDeviceResponse{} }
func (m *PatchDeviceResponse) String() string { return proto.CompactTextString(m) }
func (*PatchDeviceResponse) ProtoMessage()    {}
func (*PatchDeviceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c6fcaf6b9150132, []int{1}
}

func (m *PatchDeviceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchDeviceResponse.Unmarshal(m, b)
}
func (m *PatchDeviceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchDeviceResponse.Marshal(b, m, deterministic)
}
func (m *PatchDeviceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchDeviceResponse.Merge(m, src)
}
func (m *PatchDeviceResponse) XXX_Size() int {
	return xxx_messageInfo_PatchDeviceResponse.Size(m)
}
func (m *PatchDeviceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchDeviceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PatchDeviceResponse proto.InternalMessageInfo

func (m *PatchDeviceResponse) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func init() {
	proto.RegisterType((*PatchDeviceRequest)(nil), "ai.metathings.service.deviced.PatchDeviceRequest")
	proto.RegisterType((*PatchDeviceResponse)(nil), "ai.metathings.service.deviced.PatchDeviceResponse")
}

func init() { proto.RegisterFile("patch_device.proto", fileDescriptor_7c6fcaf6b9150132) }

var fileDescriptor_7c6fcaf6b9150132 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x48, 0x2c, 0x49,
	0xce, 0x88, 0x4f, 0x49, 0x2d, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92,
	0x4d, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e,
	0x2d, 0x02, 0x4b, 0x42, 0xd4, 0xa4, 0x48, 0x89, 0x97, 0x25, 0xe6, 0x64, 0xa6, 0x24, 0x96, 0xa4,
	0xea, 0xc3, 0x18, 0x10, 0x7d, 0x52, 0xdc, 0xb9, 0xf9, 0x29, 0xa9, 0x39, 0x10, 0x8e, 0x52, 0x3c,
	0x97, 0x50, 0x00, 0xc8, 0x68, 0x17, 0xb0, 0xae, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21,
	0x4f, 0x2e, 0x36, 0x88, 0x31, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xea, 0x7a, 0x78, 0xed,
	0xd2, 0xf3, 0x2f, 0x80, 0xe8, 0x77, 0xe2, 0xf8, 0xe5, 0xc4, 0xda, 0xc5, 0xc8, 0x24, 0xc0, 0x18,
	0x04, 0x35, 0x40, 0x29, 0x84, 0x4b, 0x18, 0xc5, 0x82, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21,
	0x5b, 0x34, 0x1b, 0x54, 0x09, 0xd8, 0x00, 0xd5, 0x0e, 0xd5, 0x94, 0xc4, 0x06, 0x76, 0xbd, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x50, 0x52, 0xa6, 0x4e, 0x18, 0x01, 0x00, 0x00,
}
