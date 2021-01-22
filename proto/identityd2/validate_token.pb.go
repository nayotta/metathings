// Code generated by protoc-gen-go. DO NOT EDIT.
// source: validate_token.proto

package identityd2

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

type ValidateTokenRequest struct {
	Token                *OpToken `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateTokenRequest) Reset()         { *m = ValidateTokenRequest{} }
func (m *ValidateTokenRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateTokenRequest) ProtoMessage()    {}
func (*ValidateTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b5b3bf2c98d17ba, []int{0}
}

func (m *ValidateTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateTokenRequest.Unmarshal(m, b)
}
func (m *ValidateTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateTokenRequest.Marshal(b, m, deterministic)
}
func (m *ValidateTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateTokenRequest.Merge(m, src)
}
func (m *ValidateTokenRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateTokenRequest.Size(m)
}
func (m *ValidateTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateTokenRequest proto.InternalMessageInfo

func (m *ValidateTokenRequest) GetToken() *OpToken {
	if m != nil {
		return m.Token
	}
	return nil
}

type ValidateTokenResponse struct {
	Token                *Token   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateTokenResponse) Reset()         { *m = ValidateTokenResponse{} }
func (m *ValidateTokenResponse) String() string { return proto.CompactTextString(m) }
func (*ValidateTokenResponse) ProtoMessage()    {}
func (*ValidateTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b5b3bf2c98d17ba, []int{1}
}

func (m *ValidateTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateTokenResponse.Unmarshal(m, b)
}
func (m *ValidateTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateTokenResponse.Marshal(b, m, deterministic)
}
func (m *ValidateTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateTokenResponse.Merge(m, src)
}
func (m *ValidateTokenResponse) XXX_Size() int {
	return xxx_messageInfo_ValidateTokenResponse.Size(m)
}
func (m *ValidateTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateTokenResponse proto.InternalMessageInfo

func (m *ValidateTokenResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func init() {
	proto.RegisterType((*ValidateTokenRequest)(nil), "ai.metathings.service.identityd2.ValidateTokenRequest")
	proto.RegisterType((*ValidateTokenResponse)(nil), "ai.metathings.service.identityd2.ValidateTokenResponse")
}

func init() { proto.RegisterFile("validate_token.proto", fileDescriptor_2b5b3bf2c98d17ba) }

var fileDescriptor_2b5b3bf2c98d17ba = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x4b, 0xcc, 0xc9,
	0x4c, 0x49, 0x2c, 0x49, 0x8d, 0x2f, 0xc9, 0xcf, 0x4e, 0xcd, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x52, 0x48, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6,
	0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0xcb, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9,
	0x4c, 0x31, 0x92, 0x12, 0x87, 0xe9, 0xd3, 0x87, 0x31, 0x20, 0x5a, 0xa5, 0xb8, 0x73, 0xf3, 0x53,
	0x52, 0x73, 0x20, 0x1c, 0xa5, 0x44, 0x2e, 0x91, 0x30, 0xa8, 0x74, 0x08, 0xc8, 0xf8, 0xa0, 0xd4,
	0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x4f, 0x2e, 0x56, 0xb0, 0x75, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0xdc, 0x46, 0x9a, 0x7a, 0x84, 0xec, 0xd3, 0xf3, 0x2f, 0x00, 0x1b, 0xe0, 0xc4, 0xf1, 0xcb, 0x89,
	0xb5, 0x8b, 0x91, 0x49, 0x80, 0x31, 0x08, 0x62, 0x82, 0x52, 0x18, 0x97, 0x28, 0x9a, 0x15, 0xc5,
	0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0xb6, 0xa8, 0x76, 0xa8, 0x13, 0xb6, 0x03, 0xa2, 0x1f, 0xa2,
	0x2b, 0x89, 0x0d, 0xec, 0x03, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xb2, 0x26, 0x3b,
	0x21, 0x01, 0x00, 0x00,
}