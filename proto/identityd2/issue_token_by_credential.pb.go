// Code generated by protoc-gen-go. DO NOT EDIT.
// source: issue_token_by_credential.proto

package identityd2

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type IssueTokenByCredentialRequest struct {
	Credential           *OpCredential         `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
	Timestamp            *timestamp.Timestamp  `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Nonce                *wrappers.Int64Value  `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Hmac                 *wrappers.StringValue `protobuf:"bytes,4,opt,name=hmac,proto3" json:"hmac,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *IssueTokenByCredentialRequest) Reset()         { *m = IssueTokenByCredentialRequest{} }
func (m *IssueTokenByCredentialRequest) String() string { return proto.CompactTextString(m) }
func (*IssueTokenByCredentialRequest) ProtoMessage()    {}
func (*IssueTokenByCredentialRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73e82e40d8b6fecc, []int{0}
}

func (m *IssueTokenByCredentialRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueTokenByCredentialRequest.Unmarshal(m, b)
}
func (m *IssueTokenByCredentialRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueTokenByCredentialRequest.Marshal(b, m, deterministic)
}
func (m *IssueTokenByCredentialRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueTokenByCredentialRequest.Merge(m, src)
}
func (m *IssueTokenByCredentialRequest) XXX_Size() int {
	return xxx_messageInfo_IssueTokenByCredentialRequest.Size(m)
}
func (m *IssueTokenByCredentialRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueTokenByCredentialRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IssueTokenByCredentialRequest proto.InternalMessageInfo

func (m *IssueTokenByCredentialRequest) GetCredential() *OpCredential {
	if m != nil {
		return m.Credential
	}
	return nil
}

func (m *IssueTokenByCredentialRequest) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *IssueTokenByCredentialRequest) GetNonce() *wrappers.Int64Value {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *IssueTokenByCredentialRequest) GetHmac() *wrappers.StringValue {
	if m != nil {
		return m.Hmac
	}
	return nil
}

type IssueTokenByCredentialResponse struct {
	Token                *Token   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IssueTokenByCredentialResponse) Reset()         { *m = IssueTokenByCredentialResponse{} }
func (m *IssueTokenByCredentialResponse) String() string { return proto.CompactTextString(m) }
func (*IssueTokenByCredentialResponse) ProtoMessage()    {}
func (*IssueTokenByCredentialResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_73e82e40d8b6fecc, []int{1}
}

func (m *IssueTokenByCredentialResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueTokenByCredentialResponse.Unmarshal(m, b)
}
func (m *IssueTokenByCredentialResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueTokenByCredentialResponse.Marshal(b, m, deterministic)
}
func (m *IssueTokenByCredentialResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueTokenByCredentialResponse.Merge(m, src)
}
func (m *IssueTokenByCredentialResponse) XXX_Size() int {
	return xxx_messageInfo_IssueTokenByCredentialResponse.Size(m)
}
func (m *IssueTokenByCredentialResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueTokenByCredentialResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IssueTokenByCredentialResponse proto.InternalMessageInfo

func (m *IssueTokenByCredentialResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func init() {
	proto.RegisterType((*IssueTokenByCredentialRequest)(nil), "ai.metathings.service.identityd2.IssueTokenByCredentialRequest")
	proto.RegisterType((*IssueTokenByCredentialResponse)(nil), "ai.metathings.service.identityd2.IssueTokenByCredentialResponse")
}

func init() { proto.RegisterFile("issue_token_by_credential.proto", fileDescriptor_73e82e40d8b6fecc) }

var fileDescriptor_73e82e40d8b6fecc = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x86, 0x95, 0x7c, 0xed, 0xa7, 0xe2, 0x2e, 0xc8, 0x0b, 0x51, 0x81, 0xb6, 0xea, 0x02, 0x93,
	0x2b, 0x15, 0xc4, 0x00, 0x62, 0x09, 0x53, 0x27, 0xa4, 0x50, 0x21, 0xb6, 0xc8, 0x4d, 0x0e, 0xa9,
	0x45, 0x62, 0x1b, 0xfb, 0xa4, 0xa8, 0xb7, 0xc0, 0xe5, 0x30, 0x70, 0x71, 0x4c, 0xa8, 0x4e, 0xfa,
	0xa3, 0x56, 0xa8, 0x9b, 0xa5, 0xf3, 0x3e, 0xaf, 0x9e, 0x73, 0x64, 0xd2, 0x13, 0xd6, 0x96, 0x10,
	0xa3, 0x7a, 0x03, 0x19, 0x4f, 0x17, 0x71, 0x62, 0x20, 0x05, 0x89, 0x82, 0xe7, 0x4c, 0x1b, 0x85,
	0x8a, 0xf6, 0xb9, 0x60, 0x05, 0x20, 0xc7, 0x99, 0x90, 0x99, 0x65, 0x16, 0xcc, 0x5c, 0x24, 0xc0,
	0x84, 0x4b, 0xe1, 0x22, 0x1d, 0x75, 0xba, 0x99, 0x52, 0x59, 0x0e, 0x43, 0x97, 0x9f, 0x96, 0xaf,
	0xc3, 0x0f, 0xc3, 0xb5, 0x06, 0x63, 0xab, 0x86, 0x4e, 0x6f, 0x77, 0x8e, 0xa2, 0x00, 0x8b, 0xbc,
	0xd0, 0x75, 0xe0, 0x64, 0xce, 0x73, 0x91, 0x72, 0x84, 0xe1, 0xea, 0x51, 0x0f, 0xda, 0x85, 0x4a,
	0xa1, 0x16, 0x19, 0x7c, 0xfb, 0xe4, 0x7c, 0xbc, 0x94, 0x9d, 0x2c, 0x5d, 0xc3, 0xc5, 0xc3, 0xda,
	0x34, 0x82, 0xf7, 0x12, 0x2c, 0xd2, 0x17, 0x42, 0x36, 0xfa, 0x81, 0xd7, 0xf7, 0x2e, 0xdb, 0x23,
	0xc6, 0x0e, 0xf9, 0xb3, 0x47, 0xbd, 0xa9, 0x0a, 0x5b, 0x3f, 0x61, 0xf3, 0xd3, 0xf3, 0x8f, 0xbd,
	0x68, 0xab, 0x8b, 0x86, 0xe4, 0x68, 0x2d, 0x1d, 0xf8, 0xae, 0xb8, 0xc3, 0xaa, 0xb5, 0xd8, 0x6a,
	0x2d, 0x36, 0x59, 0x25, 0x5c, 0xc9, 0x97, 0xe7, 0xb7, 0xbc, 0x68, 0x83, 0xd1, 0x3b, 0xd2, 0x94,
	0x4a, 0x26, 0x10, 0xfc, 0x73, 0xfc, 0xe9, 0x1e, 0x3f, 0x96, 0x78, 0x73, 0xfd, 0xcc, 0xf3, 0x12,
	0xb6, 0x2c, 0x2a, 0x86, 0xde, 0x92, 0xc6, 0xac, 0xe0, 0x49, 0xd0, 0x70, 0xec, 0xd9, 0x1e, 0xfb,
	0x84, 0x46, 0xc8, 0x6c, 0x17, 0x76, 0xcc, 0x20, 0x26, 0xdd, 0xbf, 0xee, 0x66, 0xb5, 0x92, 0x16,
	0xe8, 0x3d, 0x69, 0xba, 0x0f, 0x50, 0xdf, 0xec, 0xe2, 0xf0, 0xcd, 0x5c, 0x57, 0x54, 0x51, 0xd3,
	0xff, 0x4e, 0xe3, 0xea, 0x37, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xc5, 0x64, 0x26, 0x4c, 0x02, 0x00,
	0x00,
}