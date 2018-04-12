// Code generated by protoc-gen-go. DO NOT EDIT.
// source: send_unary_call.proto

package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SendUnaryCallRequest struct {
	Payload *UnaryCallPayload `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
}

func (m *SendUnaryCallRequest) Reset()                    { *m = SendUnaryCallRequest{} }
func (m *SendUnaryCallRequest) String() string            { return proto.CompactTextString(m) }
func (*SendUnaryCallRequest) ProtoMessage()               {}
func (*SendUnaryCallRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *SendUnaryCallRequest) GetPayload() *UnaryCallPayload {
	if m != nil {
		return m.Payload
	}
	return nil
}

type SendUnaryCallResponse struct {
	Payload *UnaryCallPayload `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
}

func (m *SendUnaryCallResponse) Reset()                    { *m = SendUnaryCallResponse{} }
func (m *SendUnaryCallResponse) String() string            { return proto.CompactTextString(m) }
func (*SendUnaryCallResponse) ProtoMessage()               {}
func (*SendUnaryCallResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *SendUnaryCallResponse) GetPayload() *UnaryCallPayload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*SendUnaryCallRequest)(nil), "ai.metathings.service.core.SendUnaryCallRequest")
	proto.RegisterType((*SendUnaryCallResponse)(nil), "ai.metathings.service.core.SendUnaryCallResponse")
}

func init() { proto.RegisterFile("send_unary_call.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4e, 0xcd, 0x4b,
	0x89, 0x2f, 0xcd, 0x4b, 0x2c, 0xaa, 0x8c, 0x4f, 0x4e, 0xcc, 0xc9, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0x4a, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f,
	0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4b, 0xce, 0x2f, 0x4a, 0x95, 0x32, 0x4b,
	0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x2d, 0xcf, 0x2c, 0xc9, 0xce,
	0x2f, 0xd7, 0x4f, 0xcf, 0xd7, 0x05, 0x6b, 0xd4, 0x2d, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9,
	0x2f, 0x2a, 0xd6, 0x87, 0x33, 0x21, 0x66, 0x4a, 0x09, 0xa0, 0xdb, 0xa2, 0x94, 0xc6, 0x25, 0x12,
	0x9c, 0x9a, 0x97, 0x12, 0x0a, 0x12, 0x77, 0x4e, 0xcc, 0xc9, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x11, 0xf2, 0xe3, 0x62, 0x2f, 0x48, 0xac, 0xcc, 0xc9, 0x4f, 0x4c, 0x91, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x36, 0xd2, 0xd1, 0xc3, 0xed, 0x1e, 0x3d, 0xb8, 0xf6, 0x00, 0x88, 0x1e, 0x27, 0xb6,
	0x47, 0xf7, 0xe5, 0x99, 0x14, 0x18, 0x83, 0x60, 0x86, 0x28, 0xc5, 0x73, 0x89, 0xa2, 0xd9, 0x53,
	0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0xe4, 0x46, 0x91, 0x45, 0x70, 0x0b, 0x92, 0xd8, 0xc0, 0xfe,
	0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x97, 0xe4, 0x1c, 0x17, 0x4e, 0x01, 0x00, 0x00,
}