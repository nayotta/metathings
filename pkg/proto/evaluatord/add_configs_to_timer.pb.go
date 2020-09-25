// Code generated by protoc-gen-go. DO NOT EDIT.
// source: add_configs_to_timer.proto

package evaluatord

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

type AddConfigsToTimerRequest struct {
	Timer                *OpTimer            `protobuf:"bytes,1,opt,name=timer,proto3" json:"timer,omitempty"`
	Configs              []*deviced.OpConfig `protobuf:"bytes,2,rep,name=configs,proto3" json:"configs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *AddConfigsToTimerRequest) Reset()         { *m = AddConfigsToTimerRequest{} }
func (m *AddConfigsToTimerRequest) String() string { return proto.CompactTextString(m) }
func (*AddConfigsToTimerRequest) ProtoMessage()    {}
func (*AddConfigsToTimerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed8f7be63e059bb0, []int{0}
}

func (m *AddConfigsToTimerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddConfigsToTimerRequest.Unmarshal(m, b)
}
func (m *AddConfigsToTimerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddConfigsToTimerRequest.Marshal(b, m, deterministic)
}
func (m *AddConfigsToTimerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddConfigsToTimerRequest.Merge(m, src)
}
func (m *AddConfigsToTimerRequest) XXX_Size() int {
	return xxx_messageInfo_AddConfigsToTimerRequest.Size(m)
}
func (m *AddConfigsToTimerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddConfigsToTimerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddConfigsToTimerRequest proto.InternalMessageInfo

func (m *AddConfigsToTimerRequest) GetTimer() *OpTimer {
	if m != nil {
		return m.Timer
	}
	return nil
}

func (m *AddConfigsToTimerRequest) GetConfigs() []*deviced.OpConfig {
	if m != nil {
		return m.Configs
	}
	return nil
}

func init() {
	proto.RegisterType((*AddConfigsToTimerRequest)(nil), "ai.metathings.service.evaluatord.AddConfigsToTimerRequest")
}

func init() { proto.RegisterFile("add_configs_to_timer.proto", fileDescriptor_ed8f7be63e059bb0) }

var fileDescriptor_ed8f7be63e059bb0 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xb1, 0x4e, 0x84, 0x40,
	0x10, 0x40, 0xb3, 0x67, 0x4e, 0xcd, 0x62, 0x61, 0x68, 0x24, 0x54, 0xc4, 0xc6, 0xb3, 0xd9, 0x4d,
	0xce, 0xc2, 0xfa, 0xb0, 0xb2, 0xba, 0x84, 0x5c, 0x65, 0x43, 0xe6, 0x98, 0x91, 0xdb, 0x04, 0x18,
	0x84, 0x81, 0xc4, 0x5f, 0xf0, 0x4b, 0xfc, 0x46, 0x2b, 0x73, 0xc0, 0x05, 0x1b, 0x73, 0xdd, 0x6e,
	0x32, 0xef, 0xcd, 0xcb, 0xe8, 0x10, 0x10, 0xd3, 0x8c, 0xab, 0x77, 0x97, 0xb7, 0xa9, 0x70, 0x2a,
	0xae, 0xa4, 0xc6, 0xd4, 0x0d, 0x0b, 0xfb, 0x11, 0x38, 0x53, 0x92, 0x80, 0x1c, 0x5c, 0x95, 0xb7,
	0xa6, 0xa5, 0xa6, 0x77, 0x19, 0x19, 0xea, 0xa1, 0xe8, 0x40, 0xb8, 0xc1, 0xf0, 0xae, 0x87, 0xc2,
	0x21, 0x08, 0xd9, 0xd3, 0x63, 0x44, 0xc3, 0xe7, 0xdc, 0xc9, 0xa1, 0xdb, 0x9b, 0x8c, 0x4b, 0x5b,
	0xc1, 0x27, 0x8b, 0x80, 0x9d, 0x55, 0x76, 0x18, 0xb2, 0x48, 0x47, 0x1f, 0xda, 0x92, 0x91, 0x8a,
	0x09, 0xf4, 0xfe, 0x7c, 0xee, 0xbf, 0x95, 0x0e, 0x36, 0x88, 0x2f, 0x63, 0xde, 0x8e, 0x77, 0xc7,
	0xb8, 0x84, 0x3e, 0x3a, 0x6a, 0xc5, 0x7f, 0xd5, 0xcb, 0x21, 0x36, 0x50, 0x91, 0x5a, 0x79, 0xeb,
	0x47, 0x73, 0xae, 0xd6, 0x6c, 0xeb, 0x41, 0x10, 0x5f, 0xff, 0xc4, 0xcb, 0x2f, 0xb5, 0xb8, 0x55,
	0xc9, 0x68, 0xf0, 0x37, 0xfa, 0x6a, 0x3a, 0x41, 0xb0, 0x88, 0x2e, 0x56, 0xde, 0xfa, 0xe1, 0x1f,
	0xd9, 0x54, 0x6c, 0xb6, 0xf5, 0xd8, 0x94, 0x9c, 0xb8, 0xf8, 0xe6, 0x4d, 0xcf, 0x9b, 0xf6, 0x97,
	0x43, 0xff, 0xd3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x3b, 0x1e, 0xa4, 0x5e, 0x01, 0x00,
	0x00,
}
