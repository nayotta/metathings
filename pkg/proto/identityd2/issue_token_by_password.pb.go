// Code generated by protoc-gen-go. DO NOT EDIT.
// source: issue_token_by_password.proto

package identityd2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type IssueTokenByPasswordRequest struct {
	Entity               *OpEntity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *IssueTokenByPasswordRequest) Reset()         { *m = IssueTokenByPasswordRequest{} }
func (m *IssueTokenByPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*IssueTokenByPasswordRequest) ProtoMessage()    {}
func (*IssueTokenByPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_186c4bb50bac5b83, []int{0}
}

func (m *IssueTokenByPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueTokenByPasswordRequest.Unmarshal(m, b)
}
func (m *IssueTokenByPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueTokenByPasswordRequest.Marshal(b, m, deterministic)
}
func (m *IssueTokenByPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueTokenByPasswordRequest.Merge(m, src)
}
func (m *IssueTokenByPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_IssueTokenByPasswordRequest.Size(m)
}
func (m *IssueTokenByPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueTokenByPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IssueTokenByPasswordRequest proto.InternalMessageInfo

func (m *IssueTokenByPasswordRequest) GetEntity() *OpEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type IssueTokenByPasswordResponse struct {
	Token                *Token   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IssueTokenByPasswordResponse) Reset()         { *m = IssueTokenByPasswordResponse{} }
func (m *IssueTokenByPasswordResponse) String() string { return proto.CompactTextString(m) }
func (*IssueTokenByPasswordResponse) ProtoMessage()    {}
func (*IssueTokenByPasswordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_186c4bb50bac5b83, []int{1}
}

func (m *IssueTokenByPasswordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueTokenByPasswordResponse.Unmarshal(m, b)
}
func (m *IssueTokenByPasswordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueTokenByPasswordResponse.Marshal(b, m, deterministic)
}
func (m *IssueTokenByPasswordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueTokenByPasswordResponse.Merge(m, src)
}
func (m *IssueTokenByPasswordResponse) XXX_Size() int {
	return xxx_messageInfo_IssueTokenByPasswordResponse.Size(m)
}
func (m *IssueTokenByPasswordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueTokenByPasswordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IssueTokenByPasswordResponse proto.InternalMessageInfo

func (m *IssueTokenByPasswordResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func init() {
	proto.RegisterType((*IssueTokenByPasswordRequest)(nil), "ai.metathings.service.identityd2.IssueTokenByPasswordRequest")
	proto.RegisterType((*IssueTokenByPasswordResponse)(nil), "ai.metathings.service.identityd2.IssueTokenByPasswordResponse")
}

func init() { proto.RegisterFile("issue_token_by_password.proto", fileDescriptor_186c4bb50bac5b83) }

var fileDescriptor_186c4bb50bac5b83 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x89, 0x60, 0x0e, 0xdb, 0x5b, 0x4e, 0x52, 0x15, 0x43, 0x2f, 0x8a, 0xd0, 0x0d, 0x54,
	0xf0, 0xe6, 0xa5, 0xe0, 0x41, 0x2f, 0x4a, 0xf0, 0x2a, 0x61, 0xd3, 0x1d, 0xd2, 0xa1, 0x4d, 0x66,
	0xdd, 0x99, 0x34, 0xe4, 0xd7, 0x0a, 0xfe, 0x12, 0xe9, 0x6e, 0xf0, 0x24, 0xf4, 0xf6, 0xf6, 0xf0,
	0xde, 0xf7, 0xed, 0xa8, 0x6b, 0x64, 0xee, 0xa1, 0x12, 0xda, 0x41, 0x57, 0xd5, 0x63, 0xe5, 0x0c,
	0xf3, 0x40, 0xde, 0x6a, 0xe7, 0x49, 0x28, 0xcb, 0x0d, 0xea, 0x16, 0xc4, 0xc8, 0x16, 0xbb, 0x86,
	0x35, 0x83, 0x3f, 0xe0, 0x06, 0x34, 0x5a, 0xe8, 0x04, 0x65, 0xb4, 0xab, 0xf9, 0x63, 0x83, 0xb2,
	0xed, 0x6b, 0xbd, 0xa1, 0xb6, 0x68, 0x07, 0x94, 0x1d, 0x0d, 0x45, 0x43, 0xcb, 0x50, 0x5f, 0x1e,
	0xcc, 0x1e, 0xad, 0x11, 0xf2, 0x5c, 0xfc, 0xc5, 0xb8, 0x3c, 0x9f, 0xb5, 0x64, 0x61, 0x1f, 0x1f,
	0x0b, 0x54, 0x97, 0x2f, 0x47, 0x8f, 0x8f, 0xa3, 0xc6, 0x7a, 0x7c, 0x9f, 0x24, 0x4a, 0xf8, 0xea,
	0x81, 0x25, 0x7b, 0x55, 0x69, 0xe4, 0x5d, 0x24, 0x79, 0x72, 0x37, 0x5b, 0xdd, 0xeb, 0x53, 0x5a,
	0xfa, 0xcd, 0x3d, 0x87, 0xb8, 0x4e, 0x7f, 0xbe, 0x6f, 0xce, 0xf2, 0xa4, 0x9c, 0x16, 0x16, 0x9f,
	0xea, 0xea, 0x7f, 0x14, 0x3b, 0xea, 0x18, 0xb2, 0x27, 0x75, 0x1e, 0x8e, 0x31, 0xa1, 0x6e, 0x4f,
	0xa3, 0xc2, 0x52, 0x19, 0x5b, 0x75, 0x1a, 0x3e, 0xf4, 0xf0, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xbf,
	0xf6, 0xd2, 0x1f, 0x58, 0x01, 0x00, 0x00,
}
