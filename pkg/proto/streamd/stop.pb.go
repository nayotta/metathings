// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stop.proto

package streamd

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

type StopRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *StopRequest) Reset()         { *m = StopRequest{} }
func (m *StopRequest) String() string { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()    {}
func (*StopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_stop_b108dae1357d5c4d, []int{0}
}
func (m *StopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopRequest.Unmarshal(m, b)
}
func (m *StopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopRequest.Marshal(b, m, deterministic)
}
func (dst *StopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopRequest.Merge(dst, src)
}
func (m *StopRequest) XXX_Size() int {
	return xxx_messageInfo_StopRequest.Size(m)
}
func (m *StopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopRequest proto.InternalMessageInfo

func (m *StopRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

type StopResponse struct {
	Stream               *Stream  `protobuf:"bytes,1,opt,name=stream" json:"stream,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopResponse) Reset()         { *m = StopResponse{} }
func (m *StopResponse) String() string { return proto.CompactTextString(m) }
func (*StopResponse) ProtoMessage()    {}
func (*StopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_stop_b108dae1357d5c4d, []int{1}
}
func (m *StopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopResponse.Unmarshal(m, b)
}
func (m *StopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopResponse.Marshal(b, m, deterministic)
}
func (dst *StopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopResponse.Merge(dst, src)
}
func (m *StopResponse) XXX_Size() int {
	return xxx_messageInfo_StopResponse.Size(m)
}
func (m *StopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StopResponse proto.InternalMessageInfo

func (m *StopResponse) GetStream() *Stream {
	if m != nil {
		return m.Stream
	}
	return nil
}

func init() {
	proto.RegisterType((*StopRequest)(nil), "ai.metathings.service.streamd.StopRequest")
	proto.RegisterType((*StopResponse)(nil), "ai.metathings.service.streamd.StopResponse")
}

func init() { proto.RegisterFile("stop.proto", fileDescriptor_stop_b108dae1357d5c4d) }

var fileDescriptor_stop_b108dae1357d5c4d = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x86, 0xd5, 0x0c, 0x19, 0xdc, 0x4e, 0x99, 0x50, 0xc5, 0x47, 0x55, 0x09, 0x89, 0xa5, 0x67,
	0x09, 0x10, 0x1b, 0x0b, 0xcc, 0x2c, 0xa9, 0xc4, 0xee, 0x34, 0x87, 0x7b, 0x22, 0xce, 0x19, 0xdf,
	0xa5, 0xf9, 0xb9, 0x48, 0xfc, 0x12, 0x44, 0x1c, 0x18, 0xd9, 0x6c, 0xdd, 0xbd, 0xcf, 0xf3, 0x9e,
	0x31, 0xa2, 0x1c, 0x21, 0x26, 0x56, 0xae, 0x2e, 0x1c, 0x41, 0x40, 0x75, 0x7a, 0xa4, 0xde, 0x0b,
	0x08, 0xa6, 0x13, 0x1d, 0x10, 0x44, 0x13, 0xba, 0xd0, 0xae, 0x2f, 0x3d, 0xb3, 0xef, 0xd0, 0x4e,
	0xcb, 0xcd, 0xf0, 0x66, 0xc7, 0xe4, 0x62, 0xc4, 0x24, 0x39, 0xbe, 0x7e, 0xf0, 0xa4, 0xc7, 0xa1,
	0x81, 0x03, 0x07, 0x1b, 0x46, 0xd2, 0x77, 0x1e, 0xad, 0xe7, 0xdd, 0x34, 0xdc, 0x9d, 0x5c, 0x47,
	0xad, 0x53, 0x4e, 0x62, 0xff, 0x9e, 0x73, 0x6e, 0x95, 0x05, 0xf9, 0xb7, 0x7d, 0x36, 0xcb, 0xbd,
	0x72, 0xac, 0xf1, 0x63, 0x40, 0xd1, 0xea, 0xde, 0x14, 0xd4, 0x9e, 0x2d, 0x36, 0x8b, 0x9b, 0xe5,
	0xed, 0x39, 0xe4, 0x06, 0xf0, 0xdb, 0x00, 0xf6, 0x9a, 0xa8, 0xf7, 0xaf, 0xae, 0x1b, 0xf0, 0xa9,
	0xfc, 0xfa, 0xbc, 0x2a, 0x36, 0x8b, 0xba, 0xa0, 0x76, 0xfb, 0x62, 0x56, 0x19, 0x22, 0x91, 0x7b,
	0xc1, 0xea, 0xd1, 0x94, 0x59, 0x32, 0x93, 0xae, 0xe1, 0xdf, 0x53, 0x7f, 0xb8, 0xe8, 0x42, 0x3d,
	0x87, 0x9a, 0x72, 0x12, 0xde, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x96, 0xbb, 0x39, 0x5f, 0x2d,
	0x01, 0x00, 0x00,
}