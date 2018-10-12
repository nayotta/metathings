// Code generated by protoc-gen-go. DO NOT EDIT.
// source: validate_token.proto

package identityd2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

type ValidateTokenRequest struct {
	Token                *OpToken `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateTokenRequest) Reset()         { *m = ValidateTokenRequest{} }
func (m *ValidateTokenRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateTokenRequest) ProtoMessage()    {}
func (*ValidateTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_validate_token_817ad8c8502cfb63, []int{0}
}
func (m *ValidateTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateTokenRequest.Unmarshal(m, b)
}
func (m *ValidateTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateTokenRequest.Marshal(b, m, deterministic)
}
func (dst *ValidateTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateTokenRequest.Merge(dst, src)
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

func init() {
	proto.RegisterType((*ValidateTokenRequest)(nil), "ai.metathings.service.identityd2.ValidateTokenRequest")
}

func init() {
	proto.RegisterFile("validate_token.proto", fileDescriptor_validate_token_817ad8c8502cfb63)
}

var fileDescriptor_validate_token_817ad8c8502cfb63 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x4b, 0xcc, 0xc9,
	0x4c, 0x49, 0x2c, 0x49, 0x8d, 0x2f, 0xc9, 0xcf, 0x4e, 0xcd, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x52, 0x48, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6,
	0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0xcb, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9,
	0x4c, 0x31, 0x92, 0x32, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf,
	0x2d, 0xcf, 0x2c, 0xc9, 0xce, 0x2f, 0xd7, 0x4f, 0xcf, 0xd7, 0x05, 0x6b, 0xd7, 0x85, 0x9a, 0x99,
	0x5f, 0x54, 0xac, 0x0f, 0x67, 0x42, 0x4c, 0x96, 0xe2, 0xce, 0xcd, 0x4f, 0x49, 0xcd, 0x81, 0x70,
	0x94, 0xe2, 0xb9, 0x44, 0xc2, 0xa0, 0xd6, 0x87, 0x80, 0x6c, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x11, 0x72, 0xe7, 0x62, 0x05, 0xbb, 0x46, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x53,
	0x8f, 0x90, 0x73, 0xf4, 0xfc, 0x0b, 0xc0, 0x06, 0x38, 0xb1, 0x3d, 0xba, 0x2f, 0xcf, 0xa4, 0xc0,
	0x18, 0x04, 0xd1, 0x9f, 0xc4, 0x06, 0xb6, 0xc7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xee,
	0xab, 0x34, 0xe6, 0x00, 0x00, 0x00,
}