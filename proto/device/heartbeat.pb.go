// Code generated by protoc-gen-go. DO NOT EDIT.
// source: heartbeat.proto

package device

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	deviced "github.com/nayotta/metathings/proto/deviced"
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

type HeartbeatRequest struct {
	Module               *deviced.OpModule `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HeartbeatRequest) Reset()         { *m = HeartbeatRequest{} }
func (m *HeartbeatRequest) String() string { return proto.CompactTextString(m) }
func (*HeartbeatRequest) ProtoMessage()    {}
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c667767fb9826a9, []int{0}
}

func (m *HeartbeatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatRequest.Unmarshal(m, b)
}
func (m *HeartbeatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatRequest.Marshal(b, m, deterministic)
}
func (m *HeartbeatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatRequest.Merge(m, src)
}
func (m *HeartbeatRequest) XXX_Size() int {
	return xxx_messageInfo_HeartbeatRequest.Size(m)
}
func (m *HeartbeatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatRequest proto.InternalMessageInfo

func (m *HeartbeatRequest) GetModule() *deviced.OpModule {
	if m != nil {
		return m.Module
	}
	return nil
}

func init() {
	proto.RegisterType((*HeartbeatRequest)(nil), "ai.metathings.service.device.HeartbeatRequest")
}

func init() { proto.RegisterFile("heartbeat.proto", fileDescriptor_3c667767fb9826a9) }

var fileDescriptor_3c667767fb9826a9 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x48, 0x4d, 0x2c,
	0x2a, 0x49, 0x4a, 0x4d, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x49, 0xcc, 0xd4,
	0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0xd5, 0x4b, 0x49, 0x05, 0x51, 0x52, 0xe2, 0x65, 0x89, 0x39, 0x99, 0x29, 0x89, 0x25,
	0xa9, 0xfa, 0x30, 0x06, 0x44, 0x9b, 0x94, 0x79, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72,
	0x7e, 0xae, 0x7e, 0x5e, 0x62, 0x65, 0x7e, 0x49, 0x49, 0xa2, 0x3e, 0xc2, 0x18, 0x7d, 0xb0, 0x22,
	0x7d, 0x88, 0x21, 0x29, 0xfa, 0xb9, 0xf9, 0x29, 0xa9, 0x39, 0x10, 0x8d, 0x4a, 0xb1, 0x5c, 0x02,
	0x1e, 0x30, 0x27, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x79, 0x72, 0xb1, 0xe5, 0xe6,
	0xa7, 0x94, 0xe6, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0xa9, 0xeb, 0xe1, 0x73, 0x54,
	0x8a, 0x9e, 0x7f, 0x81, 0x2f, 0x58, 0xb9, 0x13, 0xc7, 0x2f, 0x27, 0xd6, 0x2e, 0x46, 0x26, 0x01,
	0xc6, 0x20, 0xa8, 0x01, 0x4e, 0x1c, 0x51, 0x6c, 0x10, 0x55, 0x49, 0x6c, 0x60, 0xfb, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x7b, 0x42, 0x3a, 0xf2, 0x00, 0x00, 0x00,
}
