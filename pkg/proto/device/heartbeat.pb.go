// Code generated by protoc-gen-go. DO NOT EDIT.
// source: heartbeat.proto

package ai_metathings_service_device

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	deviced "github.com/nayotta/metathings/pkg/proto/deviced"
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
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x48, 0x4d, 0x2c,
	0x2a, 0x49, 0x4a, 0x4d, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x49, 0xcc, 0xd4,
	0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0xd5, 0x4b, 0x49, 0x05, 0x51, 0x52, 0x66, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a,
	0xc9, 0xf9, 0xb9, 0xfa, 0xb9, 0xe5, 0x99, 0x25, 0xd9, 0xf9, 0xe5, 0xfa, 0xe9, 0xf9, 0xba, 0x60,
	0xad, 0xba, 0x65, 0x89, 0x39, 0x99, 0x29, 0x89, 0x25, 0xf9, 0x45, 0xc5, 0xfa, 0x70, 0x26, 0xc4,
	0x54, 0x29, 0x6b, 0x24, 0x7d, 0x79, 0x89, 0x95, 0xf9, 0x25, 0x25, 0x89, 0xfa, 0x08, 0x5b, 0xf4,
	0x0b, 0xb2, 0xd3, 0xf5, 0xc1, 0x0a, 0xf5, 0x21, 0xf6, 0xa4, 0xe8, 0xe7, 0xe6, 0xa7, 0xa4, 0xe6,
	0x40, 0x34, 0x2b, 0x45, 0x73, 0x09, 0x78, 0xc0, 0x5c, 0x19, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c,
	0x22, 0xe4, 0xce, 0xc5, 0x96, 0x9b, 0x9f, 0x52, 0x9a, 0x93, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x6d, 0xa4, 0xae, 0x87, 0xcf, 0xdd, 0x29, 0x7a, 0xfe, 0x05, 0xbe, 0x60, 0xe5, 0x4e, 0x6c, 0x8f,
	0xee, 0xcb, 0x33, 0x29, 0x30, 0x06, 0x41, 0xb5, 0x27, 0xb1, 0x81, 0xed, 0x30, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x4c, 0x3c, 0xec, 0xf4, 0x09, 0x01, 0x00, 0x00,
}
