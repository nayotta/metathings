// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remove_configs_from_device.proto

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

type RemoveConfigsFromDeviceRequest struct {
	Device               *OpDevice   `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	Configs              []*OpConfig `protobuf:"bytes,2,rep,name=configs,proto3" json:"configs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RemoveConfigsFromDeviceRequest) Reset()         { *m = RemoveConfigsFromDeviceRequest{} }
func (m *RemoveConfigsFromDeviceRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveConfigsFromDeviceRequest) ProtoMessage()    {}
func (*RemoveConfigsFromDeviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5a322a3cf41a846, []int{0}
}

func (m *RemoveConfigsFromDeviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveConfigsFromDeviceRequest.Unmarshal(m, b)
}
func (m *RemoveConfigsFromDeviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveConfigsFromDeviceRequest.Marshal(b, m, deterministic)
}
func (m *RemoveConfigsFromDeviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveConfigsFromDeviceRequest.Merge(m, src)
}
func (m *RemoveConfigsFromDeviceRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveConfigsFromDeviceRequest.Size(m)
}
func (m *RemoveConfigsFromDeviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveConfigsFromDeviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveConfigsFromDeviceRequest proto.InternalMessageInfo

func (m *RemoveConfigsFromDeviceRequest) GetDevice() *OpDevice {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *RemoveConfigsFromDeviceRequest) GetConfigs() []*OpConfig {
	if m != nil {
		return m.Configs
	}
	return nil
}

func init() {
	proto.RegisterType((*RemoveConfigsFromDeviceRequest)(nil), "ai.metathings.service.deviced.RemoveConfigsFromDeviceRequest")
}

func init() { proto.RegisterFile("remove_configs_from_device.proto", fileDescriptor_a5a322a3cf41a846) }

var fileDescriptor_a5a322a3cf41a846 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x4a, 0xcd, 0xcd,
	0x2f, 0x4b, 0x8d, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0x2f, 0x8e, 0x4f, 0x2b, 0xca, 0xcf, 0x8d,
	0x4f, 0x49, 0x2d, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4d, 0xcc,
	0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x02,
	0x4b, 0x42, 0xd4, 0xa4, 0x48, 0x99, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7,
	0xea, 0xe7, 0x96, 0x67, 0x96, 0x64, 0xe7, 0x97, 0xeb, 0xa7, 0xe7, 0xeb, 0x82, 0xf5, 0xea, 0x96,
	0x25, 0xe6, 0x64, 0xa6, 0x24, 0x96, 0xe4, 0x17, 0x15, 0xeb, 0xc3, 0x99, 0x10, 0x63, 0xa5, 0xb8,
	0x73, 0xf3, 0x53, 0x52, 0x73, 0x20, 0x1c, 0xa5, 0x35, 0x8c, 0x5c, 0x72, 0x41, 0x60, 0x87, 0x38,
	0x43, 0xdc, 0xe1, 0x56, 0x94, 0x9f, 0xeb, 0x02, 0xb6, 0x21, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8,
	0x44, 0xc8, 0x9d, 0x8b, 0x0d, 0x62, 0xa5, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0xba, 0x1e,
	0x5e, 0x77, 0xe9, 0xf9, 0x17, 0x40, 0xf4, 0x3b, 0xb1, 0x3d, 0xba, 0x2f, 0xcf, 0xa4, 0xc0, 0x18,
	0x04, 0xd5, 0x2e, 0xe4, 0xc8, 0xc5, 0x0e, 0xf5, 0xac, 0x04, 0x93, 0x02, 0x33, 0x51, 0x26, 0x41,
	0x1c, 0x15, 0x04, 0xd3, 0x97, 0xc4, 0x06, 0x76, 0xb5, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x8f,
	0x80, 0x46, 0x02, 0x3d, 0x01, 0x00, 0x00,
}
