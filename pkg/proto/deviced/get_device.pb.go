// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_device.proto

package deviced

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetDeviceRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetDeviceRequest) Reset()         { *m = GetDeviceRequest{} }
func (m *GetDeviceRequest) String() string { return proto.CompactTextString(m) }
func (*GetDeviceRequest) ProtoMessage()    {}
func (*GetDeviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_get_device_82f4a5fe1dca5b24, []int{0}
}
func (m *GetDeviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceRequest.Unmarshal(m, b)
}
func (m *GetDeviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceRequest.Marshal(b, m, deterministic)
}
func (dst *GetDeviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceRequest.Merge(dst, src)
}
func (m *GetDeviceRequest) XXX_Size() int {
	return xxx_messageInfo_GetDeviceRequest.Size(m)
}
func (m *GetDeviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceRequest proto.InternalMessageInfo

func (m *GetDeviceRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

type GetDeviceResponse struct {
	Device               *Device  `protobuf:"bytes,1,opt,name=device" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDeviceResponse) Reset()         { *m = GetDeviceResponse{} }
func (m *GetDeviceResponse) String() string { return proto.CompactTextString(m) }
func (*GetDeviceResponse) ProtoMessage()    {}
func (*GetDeviceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_get_device_82f4a5fe1dca5b24, []int{1}
}
func (m *GetDeviceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceResponse.Unmarshal(m, b)
}
func (m *GetDeviceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceResponse.Marshal(b, m, deterministic)
}
func (dst *GetDeviceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceResponse.Merge(dst, src)
}
func (m *GetDeviceResponse) XXX_Size() int {
	return xxx_messageInfo_GetDeviceResponse.Size(m)
}
func (m *GetDeviceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceResponse proto.InternalMessageInfo

func (m *GetDeviceResponse) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func init() {
	proto.RegisterType((*GetDeviceRequest)(nil), "ai.metathings.service.deviced.GetDeviceRequest")
	proto.RegisterType((*GetDeviceResponse)(nil), "ai.metathings.service.deviced.GetDeviceResponse")
}

func init() { proto.RegisterFile("get_device.proto", fileDescriptor_get_device_82f4a5fe1dca5b24) }

var fileDescriptor_get_device_82f4a5fe1dca5b24 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x3f, 0x4b, 0xc4, 0x40,
	0x10, 0xc5, 0x49, 0x8a, 0x14, 0xab, 0xc5, 0x99, 0x4a, 0x0e, 0xff, 0x1c, 0x07, 0x82, 0xcd, 0xcd,
	0x82, 0x8a, 0x9d, 0x8d, 0x08, 0x5a, 0x47, 0xb0, 0x95, 0xcd, 0xed, 0xb8, 0x37, 0x98, 0x64, 0xe2,
	0xee, 0xe4, 0xf2, 0x71, 0x05, 0x3f, 0x89, 0xb0, 0x59, 0x25, 0xd5, 0x75, 0xcb, 0xbe, 0xf7, 0x7b,
	0xef, 0x8d, 0x5a, 0x38, 0x94, 0x77, 0x8b, 0x7b, 0xda, 0x22, 0xf4, 0x9e, 0x85, 0xcb, 0x73, 0x43,
	0xd0, 0xa2, 0x18, 0xd9, 0x51, 0xe7, 0x02, 0x04, 0xf4, 0x51, 0x9c, 0x3c, 0x76, 0x79, 0xe1, 0x98,
	0x5d, 0x83, 0x3a, 0x9a, 0xeb, 0xe1, 0x43, 0x8f, 0xde, 0xf4, 0x3d, 0xfa, 0x30, 0xe1, 0xcb, 0x7b,
	0x47, 0xb2, 0x1b, 0x6a, 0xd8, 0x72, 0xab, 0xdb, 0x91, 0xe4, 0x93, 0x47, 0xed, 0x78, 0x13, 0xc5,
	0xcd, 0xde, 0x34, 0x64, 0x8d, 0xb0, 0x0f, 0xfa, 0xff, 0x99, 0xb8, 0xe3, 0xf9, 0x88, 0xf5, 0x8b,
	0x5a, 0x3c, 0xa3, 0x3c, 0xc5, 0xaf, 0x0a, 0xbf, 0x06, 0x0c, 0x52, 0xde, 0xa9, 0x9c, 0xec, 0x69,
	0xb6, 0xca, 0xae, 0x8f, 0x6e, 0xce, 0x60, 0x9a, 0x01, 0x7f, 0x33, 0xe0, 0x55, 0x3c, 0x75, 0xee,
	0xcd, 0x34, 0x03, 0x3e, 0x16, 0x3f, 0xdf, 0x97, 0xf9, 0x2a, 0xab, 0x72, 0xb2, 0xeb, 0x4a, 0x9d,
	0xcc, 0x92, 0x42, 0xcf, 0x5d, 0xc0, 0xf2, 0x41, 0x15, 0x53, 0x5d, 0x8a, 0xbb, 0x82, 0x83, 0x47,
	0x43, 0xc2, 0x13, 0x54, 0x17, 0xb1, 0xf5, 0xf6, 0x37, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x8a, 0xb2,
	0x6c, 0x3d, 0x01, 0x00, 0x00,
}
