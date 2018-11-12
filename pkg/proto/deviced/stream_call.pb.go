// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stream_call.proto

package deviced

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/wrappers"
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

type StreamCallRequest struct {
	Device               *OpDevice          `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	Value                *OpStreamCallValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *StreamCallRequest) Reset()         { *m = StreamCallRequest{} }
func (m *StreamCallRequest) String() string { return proto.CompactTextString(m) }
func (*StreamCallRequest) ProtoMessage()    {}
func (*StreamCallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61c43dc63f14d203, []int{0}
}

func (m *StreamCallRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamCallRequest.Unmarshal(m, b)
}
func (m *StreamCallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamCallRequest.Marshal(b, m, deterministic)
}
func (m *StreamCallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamCallRequest.Merge(m, src)
}
func (m *StreamCallRequest) XXX_Size() int {
	return xxx_messageInfo_StreamCallRequest.Size(m)
}
func (m *StreamCallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamCallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamCallRequest proto.InternalMessageInfo

func (m *StreamCallRequest) GetDevice() *OpDevice {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *StreamCallRequest) GetValue() *OpStreamCallValue {
	if m != nil {
		return m.Value
	}
	return nil
}

type StreamCallResponse struct {
	Device               *Device          `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	Value                *StreamCallValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *StreamCallResponse) Reset()         { *m = StreamCallResponse{} }
func (m *StreamCallResponse) String() string { return proto.CompactTextString(m) }
func (*StreamCallResponse) ProtoMessage()    {}
func (*StreamCallResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61c43dc63f14d203, []int{1}
}

func (m *StreamCallResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamCallResponse.Unmarshal(m, b)
}
func (m *StreamCallResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamCallResponse.Marshal(b, m, deterministic)
}
func (m *StreamCallResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamCallResponse.Merge(m, src)
}
func (m *StreamCallResponse) XXX_Size() int {
	return xxx_messageInfo_StreamCallResponse.Size(m)
}
func (m *StreamCallResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamCallResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamCallResponse proto.InternalMessageInfo

func (m *StreamCallResponse) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *StreamCallResponse) GetValue() *StreamCallValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamCallRequest)(nil), "ai.metathings.service.deviced.StreamCallRequest")
	proto.RegisterType((*StreamCallResponse)(nil), "ai.metathings.service.deviced.StreamCallResponse")
}

func init() { proto.RegisterFile("stream_call.proto", fileDescriptor_61c43dc63f14d203) }

var fileDescriptor_61c43dc63f14d203 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd0, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x07, 0x70, 0x3a, 0x70, 0x87, 0xee, 0xb4, 0x9e, 0xe6, 0x40, 0x1d, 0x03, 0xd1, 0xcb, 0x52,
	0x51, 0xf0, 0xe6, 0x45, 0x07, 0xde, 0x14, 0x26, 0x78, 0x95, 0xd7, 0xf6, 0x99, 0x05, 0x93, 0xbe,
	0x98, 0xa4, 0x2d, 0x7e, 0x14, 0xbf, 0x83, 0xdf, 0x49, 0xf0, 0x93, 0xc8, 0xd2, 0x50, 0xb5, 0x07,
	0x7b, 0xea, 0x83, 0xd7, 0xff, 0x3f, 0xbf, 0x24, 0x9e, 0x5a, 0x67, 0x10, 0xd4, 0x53, 0x0e, 0x52,
	0x32, 0x6d, 0xc8, 0x51, 0x72, 0x00, 0x82, 0x29, 0x74, 0xe0, 0xb6, 0xa2, 0xe4, 0x96, 0x59, 0x34,
	0xb5, 0xc8, 0x91, 0x15, 0xb8, 0xfb, 0x14, 0xf3, 0x43, 0x4e, 0xc4, 0x25, 0xa6, 0xfe, 0xe7, 0xac,
	0x7a, 0x4e, 0x1b, 0x03, 0x5a, 0xa3, 0xb1, 0x6d, 0x7c, 0xbe, 0xdf, 0xdf, 0x43, 0xf9, 0x16, 0x56,
	0x97, 0x5c, 0xb8, 0x6d, 0x95, 0xb1, 0x9c, 0x54, 0xaa, 0x1a, 0xe1, 0x5e, 0xa8, 0x49, 0x39, 0xad,
	0xfc, 0x72, 0x55, 0x83, 0x14, 0x05, 0x38, 0x32, 0x36, 0xed, 0xc6, 0x90, 0x9b, 0x28, 0x2a, 0x30,
	0xf0, 0x96, 0x1f, 0x51, 0x3c, 0x7d, 0xf0, 0xe8, 0x1b, 0x90, 0x72, 0x83, 0xaf, 0x15, 0x5a, 0x97,
	0xdc, 0xc6, 0xe3, 0x16, 0x38, 0x8b, 0x16, 0xd1, 0xe9, 0xe4, 0xfc, 0x84, 0xfd, 0x7b, 0x0b, 0x76,
	0xaf, 0xd7, 0x7e, 0xba, 0x1e, 0x7f, 0x7d, 0x1e, 0x8d, 0x16, 0xd1, 0x26, 0xc4, 0x93, 0xbb, 0x78,
	0xaf, 0x06, 0x59, 0xe1, 0x6c, 0xe4, 0x7b, 0xce, 0x06, 0x7b, 0x7e, 0x2c, 0x8f, 0xbb, 0x5c, 0x57,
	0xd8, 0xd6, 0x2c, 0xdf, 0xa3, 0x38, 0xf9, 0xcd, 0xb5, 0x9a, 0x4a, 0x8b, 0xc9, 0x55, 0xcf, 0x7b,
	0x3c, 0x70, 0x4e, 0xab, 0xed, 0x94, 0xeb, 0xbf, 0x4a, 0x36, 0x90, 0xee, 0x19, 0x83, 0x2d, 0x1b,
	0xfb, 0x17, 0xbd, 0xf8, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x2a, 0x43, 0xe0, 0x05, 0x02, 0x00,
	0x00,
}
