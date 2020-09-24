// Code generated by protoc-gen-go. DO NOT EDIT.
// source: heartbeat.proto

package deviced

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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
	Device               *OpDevice            `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	StartupSession       *wrappers.Int32Value `protobuf:"bytes,2,opt,name=startup_session,json=startupSession,proto3" json:"startup_session,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
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

func (m *HeartbeatRequest) GetDevice() *OpDevice {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *HeartbeatRequest) GetStartupSession() *wrappers.Int32Value {
	if m != nil {
		return m.StartupSession
	}
	return nil
}

func init() {
	proto.RegisterType((*HeartbeatRequest)(nil), "ai.metathings.service.deviced.HeartbeatRequest")
}

func init() { proto.RegisterFile("heartbeat.proto", fileDescriptor_3c667767fb9826a9) }

var fileDescriptor_3c667767fb9826a9 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x48, 0x4d, 0x2c,
	0x2a, 0x49, 0x4a, 0x4d, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4d, 0xcc, 0xd4,
	0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0xd5, 0x4b, 0x49, 0x05, 0x51, 0x29, 0x52, 0x72, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9,
	0xfa, 0x60, 0xc5, 0x49, 0xa5, 0x69, 0xfa, 0xe5, 0x45, 0x89, 0x05, 0x05, 0xa9, 0x45, 0xc5, 0x10,
	0xed, 0x52, 0xe2, 0x65, 0x89, 0x39, 0x99, 0x29, 0x89, 0x25, 0xa9, 0xfa, 0x30, 0x06, 0x54, 0x82,
	0x3b, 0x37, 0x3f, 0x25, 0x35, 0x07, 0xc2, 0x51, 0x5a, 0xcc, 0xc8, 0x25, 0xe0, 0x01, 0xb3, 0x38,
	0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0xc8, 0x93, 0x8b, 0x0d, 0x62, 0x8b, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0xb7, 0x91, 0xba, 0x1e, 0x5e, 0xa7, 0xe8, 0xf9, 0x17, 0xb8, 0x80, 0x59, 0x4e, 0x1c,
	0xbf, 0x9c, 0x58, 0xbb, 0x18, 0x99, 0x04, 0x18, 0x83, 0xa0, 0x06, 0x08, 0xb9, 0x70, 0xf1, 0x17,
	0x97, 0x24, 0x16, 0x95, 0x94, 0x16, 0xc4, 0x17, 0xa7, 0x16, 0x17, 0x67, 0xe6, 0xe7, 0x49, 0x30,
	0x81, 0xcd, 0x94, 0xd6, 0x83, 0xb8, 0x5f, 0x0f, 0xe6, 0x7e, 0x3d, 0xcf, 0xbc, 0x12, 0x63, 0xa3,
	0xb0, 0xc4, 0x9c, 0xd2, 0xd4, 0x20, 0x3e, 0xa8, 0x9e, 0x60, 0x88, 0x16, 0x27, 0xce, 0x28, 0x76,
	0xa8, 0x5d, 0x49, 0x6c, 0x60, 0xf5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xfd, 0x44,
	0xc2, 0x2f, 0x01, 0x00, 0x00,
}
