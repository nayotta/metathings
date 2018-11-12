// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delete_device.proto

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DeleteDeviceRequest struct {
	Device               *OpDevice `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DeleteDeviceRequest) Reset()         { *m = DeleteDeviceRequest{} }
func (m *DeleteDeviceRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDeviceRequest) ProtoMessage()    {}
func (*DeleteDeviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_20c7bd6751f415b5, []int{0}
}

func (m *DeleteDeviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDeviceRequest.Unmarshal(m, b)
}
func (m *DeleteDeviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDeviceRequest.Marshal(b, m, deterministic)
}
func (m *DeleteDeviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDeviceRequest.Merge(m, src)
}
func (m *DeleteDeviceRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteDeviceRequest.Size(m)
}
func (m *DeleteDeviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDeviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDeviceRequest proto.InternalMessageInfo

func (m *DeleteDeviceRequest) GetDevice() *OpDevice {
	if m != nil {
		return m.Device
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteDeviceRequest)(nil), "ai.metathings.service.deviced.DeleteDeviceRequest")
}

func init() { proto.RegisterFile("delete_device.proto", fileDescriptor_20c7bd6751f415b5) }

var fileDescriptor_20c7bd6751f415b5 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x49, 0xcd, 0x49,
	0x2d, 0x49, 0x8d, 0x4f, 0x49, 0x2d, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x92, 0x4d, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b,
	0x4e, 0x2d, 0x02, 0x4b, 0x42, 0xd4, 0xa4, 0x48, 0x99, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9,
	0x25, 0xe7, 0xe7, 0xea, 0xe7, 0x96, 0x67, 0x96, 0x64, 0xe7, 0x97, 0xeb, 0xa7, 0xe7, 0xeb, 0x82,
	0xf5, 0xea, 0x96, 0x25, 0xe6, 0x64, 0xa6, 0x24, 0x96, 0xe4, 0x17, 0x15, 0xeb, 0xc3, 0x99, 0x10,
	0x63, 0xa5, 0xb8, 0x73, 0xf3, 0x53, 0x52, 0x73, 0x20, 0x1c, 0xa5, 0x38, 0x2e, 0x61, 0x17, 0xb0,
	0xd5, 0x2e, 0x60, 0x53, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xdc, 0xb9, 0xd8, 0x20,
	0xd6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0xa9, 0xeb, 0xe1, 0x75, 0x8b, 0x9e, 0x7f, 0x01,
	0x44, 0xbf, 0x13, 0xdb, 0xa3, 0xfb, 0xf2, 0x4c, 0x0a, 0x8c, 0x41, 0x50, 0xed, 0x49, 0x6c, 0x60,
	0x6b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe1, 0xe7, 0x83, 0x6e, 0xe1, 0x00, 0x00, 0x00,
}
