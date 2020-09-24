// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delete_credential.proto

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

type DeleteCredentialRequest struct {
	Credential           *OpCredential `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DeleteCredentialRequest) Reset()         { *m = DeleteCredentialRequest{} }
func (m *DeleteCredentialRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCredentialRequest) ProtoMessage()    {}
func (*DeleteCredentialRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7074b343416089b, []int{0}
}

func (m *DeleteCredentialRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCredentialRequest.Unmarshal(m, b)
}
func (m *DeleteCredentialRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCredentialRequest.Marshal(b, m, deterministic)
}
func (m *DeleteCredentialRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCredentialRequest.Merge(m, src)
}
func (m *DeleteCredentialRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCredentialRequest.Size(m)
}
func (m *DeleteCredentialRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCredentialRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCredentialRequest proto.InternalMessageInfo

func (m *DeleteCredentialRequest) GetCredential() *OpCredential {
	if m != nil {
		return m.Credential
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteCredentialRequest)(nil), "ai.metathings.service.identityd2.DeleteCredentialRequest")
}

func init() { proto.RegisterFile("delete_credential.proto", fileDescriptor_a7074b343416089b) }

var fileDescriptor_a7074b343416089b = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x49, 0xcd, 0x49,
	0x2d, 0x49, 0x8d, 0x4f, 0x2e, 0x4a, 0x4d, 0x49, 0xcd, 0x2b, 0xc9, 0x4c, 0xcc, 0xd1, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x52, 0x48, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc,
	0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0xcb, 0x04, 0xab, 0x2a, 0xa9,
	0x4c, 0x31, 0x92, 0x12, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0x20,
	0x5a, 0xa5, 0xb8, 0x73, 0xf3, 0x53, 0x52, 0xa1, 0xe6, 0x28, 0x15, 0x73, 0x89, 0xbb, 0x80, 0xad,
	0x70, 0x86, 0xdb, 0x10, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x14, 0xc1, 0xc5, 0x85, 0xb0,
	0x56, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x4f, 0x8f, 0x90, 0xbd, 0x7a, 0xfe, 0x05, 0x08,
	0xa3, 0x9c, 0x38, 0x7e, 0x39, 0xb1, 0x76, 0x31, 0x32, 0x09, 0x30, 0x06, 0x21, 0x99, 0x95, 0xc4,
	0x06, 0xb6, 0xdb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd9, 0xfb, 0x9e, 0x57, 0xde, 0x00, 0x00,
	0x00,
}
