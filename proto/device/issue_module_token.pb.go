// Code generated by protoc-gen-go. DO NOT EDIT.
// source: issue_module_token.proto

package device

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	identityd2 "github.com/nayotta/metathings/proto/identityd2"
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

type IssueModuleTokenRequest struct {
	Credential           *identityd2.OpCredential `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
	Timestamp            *timestamp.Timestamp     `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Nonce                *wrappers.Int64Value     `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Hmac                 *wrappers.StringValue    `protobuf:"bytes,4,opt,name=hmac,proto3" json:"hmac,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *IssueModuleTokenRequest) Reset()         { *m = IssueModuleTokenRequest{} }
func (m *IssueModuleTokenRequest) String() string { return proto.CompactTextString(m) }
func (*IssueModuleTokenRequest) ProtoMessage()    {}
func (*IssueModuleTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b99749586de46f6, []int{0}
}

func (m *IssueModuleTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueModuleTokenRequest.Unmarshal(m, b)
}
func (m *IssueModuleTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueModuleTokenRequest.Marshal(b, m, deterministic)
}
func (m *IssueModuleTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueModuleTokenRequest.Merge(m, src)
}
func (m *IssueModuleTokenRequest) XXX_Size() int {
	return xxx_messageInfo_IssueModuleTokenRequest.Size(m)
}
func (m *IssueModuleTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueModuleTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IssueModuleTokenRequest proto.InternalMessageInfo

func (m *IssueModuleTokenRequest) GetCredential() *identityd2.OpCredential {
	if m != nil {
		return m.Credential
	}
	return nil
}

func (m *IssueModuleTokenRequest) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *IssueModuleTokenRequest) GetNonce() *wrappers.Int64Value {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *IssueModuleTokenRequest) GetHmac() *wrappers.StringValue {
	if m != nil {
		return m.Hmac
	}
	return nil
}

type IssueModuleTokenResponse struct {
	Token                *identityd2.Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *IssueModuleTokenResponse) Reset()         { *m = IssueModuleTokenResponse{} }
func (m *IssueModuleTokenResponse) String() string { return proto.CompactTextString(m) }
func (*IssueModuleTokenResponse) ProtoMessage()    {}
func (*IssueModuleTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b99749586de46f6, []int{1}
}

func (m *IssueModuleTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueModuleTokenResponse.Unmarshal(m, b)
}
func (m *IssueModuleTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueModuleTokenResponse.Marshal(b, m, deterministic)
}
func (m *IssueModuleTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueModuleTokenResponse.Merge(m, src)
}
func (m *IssueModuleTokenResponse) XXX_Size() int {
	return xxx_messageInfo_IssueModuleTokenResponse.Size(m)
}
func (m *IssueModuleTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueModuleTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IssueModuleTokenResponse proto.InternalMessageInfo

func (m *IssueModuleTokenResponse) GetToken() *identityd2.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func init() {
	proto.RegisterType((*IssueModuleTokenRequest)(nil), "ai.metathings.service.device.IssueModuleTokenRequest")
	proto.RegisterType((*IssueModuleTokenResponse)(nil), "ai.metathings.service.device.IssueModuleTokenResponse")
}

func init() { proto.RegisterFile("issue_module_token.proto", fileDescriptor_4b99749586de46f6) }

var fileDescriptor_4b99749586de46f6 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x69, 0xdd, 0x64, 0xc6, 0x8b, 0xf4, 0xb2, 0x32, 0x87, 0xca, 0x2e, 0x7a, 0x4a, 0x60,
	0x8a, 0x87, 0x89, 0x97, 0x79, 0xda, 0x41, 0x84, 0x39, 0x44, 0x4f, 0x23, 0x6b, 0x9f, 0x5d, 0xb0,
	0x49, 0x6a, 0xf3, 0x32, 0xd9, 0xbf, 0xe0, 0x5f, 0x23, 0xfe, 0x79, 0x9e, 0x64, 0xc9, 0x7e, 0xb1,
	0x0a, 0x9e, 0x5a, 0x78, 0xef, 0xf3, 0xe5, 0x7d, 0xbe, 0x84, 0xc4, 0xc2, 0x18, 0x0b, 0x63, 0xa9,
	0x53, 0x9b, 0xc3, 0x18, 0xf5, 0x1b, 0x28, 0x5a, 0x94, 0x1a, 0x75, 0xd4, 0xe6, 0x82, 0x4a, 0x40,
	0x8e, 0x53, 0xa1, 0x32, 0x43, 0x0d, 0x94, 0x33, 0x91, 0x00, 0x4d, 0x61, 0xf1, 0x69, 0x9d, 0x66,
	0x5a, 0x67, 0x39, 0x30, 0xb7, 0x3b, 0xb1, 0xaf, 0x0c, 0x85, 0x04, 0x83, 0x5c, 0x16, 0x1e, 0x6f,
	0x9d, 0xec, 0x2e, 0x7c, 0x94, 0xbc, 0x28, 0xa0, 0x34, 0xcb, 0x79, 0x73, 0xc6, 0x73, 0x91, 0x72,
	0x04, 0xb6, 0xfa, 0x59, 0x0e, 0x7a, 0x99, 0xc0, 0xa9, 0x9d, 0xd0, 0x44, 0x4b, 0xa6, 0xf8, 0x5c,
	0x23, 0x72, 0xb6, 0xb9, 0xc3, 0xe7, 0x31, 0x91, 0x82, 0x42, 0x81, 0xf3, 0xb4, 0xcb, 0xa4, 0x4e,
	0x21, 0xf7, 0x6c, 0xe7, 0x2b, 0x24, 0xcd, 0xc1, 0x42, 0xe8, 0xde, 0xf9, 0x8c, 0x16, 0x3a, 0x43,
	0x78, 0xb7, 0x60, 0x30, 0x7a, 0x26, 0x24, 0x29, 0xc1, 0x71, 0x3c, 0x8f, 0x83, 0xb3, 0xe0, 0xe2,
	0xb0, 0x4b, 0xe9, 0xdf, 0x92, 0x9b, 0x78, 0xfa, 0x50, 0xdc, 0xad, 0xa9, 0x7e, 0xe3, 0xa7, 0x5f,
	0xff, 0x0c, 0xc2, 0xa3, 0x60, 0xb8, 0x95, 0x15, 0xf5, 0xc9, 0xc1, 0xda, 0x3e, 0x0e, 0x5d, 0x70,
	0x8b, 0x7a, 0x7d, 0xba, 0xd2, 0xa7, 0xa3, 0xd5, 0x86, 0x0b, 0xf9, 0x0e, 0xc2, 0x46, 0x30, 0xdc,
	0x60, 0xd1, 0x0d, 0xa9, 0x2b, 0xad, 0x12, 0x88, 0xf7, 0x1c, 0x7f, 0x5c, 0xe1, 0x07, 0x0a, 0xaf,
	0xaf, 0x9e, 0x78, 0x6e, 0x61, 0xeb, 0x0a, 0xcf, 0x44, 0x3d, 0x52, 0x9b, 0x4a, 0x9e, 0xc4, 0x35,
	0xc7, 0xb6, 0x2b, 0xec, 0x23, 0x96, 0x42, 0x65, 0xbb, 0xb0, 0x63, 0x3a, 0x2f, 0x24, 0xae, 0x36,
	0x66, 0x0a, 0xad, 0x0c, 0x44, 0xb7, 0xa4, 0xee, 0x5e, 0xc4, 0xb2, 0xad, 0xf3, 0xff, 0xdb, 0xf2,
	0xbc, 0xa7, 0x26, 0xfb, 0xee, 0x80, 0xcb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xc7, 0xa6,
	0x12, 0x64, 0x02, 0x00, 0x00,
}