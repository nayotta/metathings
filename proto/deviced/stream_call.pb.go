// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stream_call.proto

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
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x2e, 0x29, 0x4a,
	0x4d, 0xcc, 0x8d, 0x4f, 0x4e, 0xcc, 0xc9, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4d,
	0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d,
	0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4b, 0x49, 0x05, 0x51, 0x29, 0x52, 0xe2, 0x65, 0x89, 0x39, 0x99,
	0x29, 0x89, 0x25, 0xa9, 0xfa, 0x30, 0x06, 0x44, 0x9f, 0x14, 0x77, 0x6e, 0x7e, 0x4a, 0x2a, 0xd4,
	0x10, 0xa5, 0x0d, 0x8c, 0x5c, 0x82, 0xc1, 0x60, 0xa3, 0x9d, 0x13, 0x73, 0x72, 0x82, 0x52, 0x0b,
	0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x3c, 0xb9, 0xd8, 0x20, 0xc6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x1b, 0xa9, 0xeb, 0xe1, 0xb5, 0x4b, 0xcf, 0xbf, 0xc0, 0x05, 0xcc, 0x72, 0xe2, 0xf8, 0xe5, 0xc4,
	0xda, 0xc5, 0xc8, 0x24, 0xc0, 0x18, 0x04, 0x35, 0x40, 0x28, 0x80, 0x8b, 0xb5, 0x2c, 0x31, 0xa7,
	0x34, 0x55, 0x82, 0x09, 0x6c, 0x92, 0x01, 0x41, 0x93, 0x10, 0xae, 0x09, 0x03, 0xe9, 0x43, 0x32,
	0x12, 0x62, 0x90, 0xd2, 0x4c, 0x46, 0x2e, 0x21, 0x64, 0x27, 0x17, 0x17, 0xe4, 0xe7, 0x15, 0xa7,
	0x0a, 0xd9, 0xa2, 0xb9, 0x59, 0x95, 0x80, 0x4d, 0x10, 0x17, 0xc3, 0xdd, 0xe9, 0x82, 0xea, 0x4e,
	0x3d, 0x02, 0xba, 0xd1, 0x5c, 0x09, 0x75, 0x9b, 0x13, 0x67, 0x14, 0x3b, 0x54, 0x45, 0x12, 0x1b,
	0x38, 0x80, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x19, 0x5d, 0x39, 0x4d, 0xba, 0x01, 0x00,
	0x00,
}
