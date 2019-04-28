// Code generated by protoc-gen-go. DO NOT EDIT.
// source: create_credential.proto

package identityd2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type CreateCredentialRequest struct {
	Credential           *OpCredential        `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
	SecretSize           *wrappers.Int32Value `protobuf:"bytes,2,opt,name=secret_size,json=secretSize,proto3" json:"secret_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateCredentialRequest) Reset()         { *m = CreateCredentialRequest{} }
func (m *CreateCredentialRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCredentialRequest) ProtoMessage()    {}
func (*CreateCredentialRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f34767a917b63eb, []int{0}
}

func (m *CreateCredentialRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCredentialRequest.Unmarshal(m, b)
}
func (m *CreateCredentialRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCredentialRequest.Marshal(b, m, deterministic)
}
func (m *CreateCredentialRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCredentialRequest.Merge(m, src)
}
func (m *CreateCredentialRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCredentialRequest.Size(m)
}
func (m *CreateCredentialRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCredentialRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCredentialRequest proto.InternalMessageInfo

func (m *CreateCredentialRequest) GetCredential() *OpCredential {
	if m != nil {
		return m.Credential
	}
	return nil
}

func (m *CreateCredentialRequest) GetSecretSize() *wrappers.Int32Value {
	if m != nil {
		return m.SecretSize
	}
	return nil
}

type CreateCredentialResponse struct {
	Credential           *Credential `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateCredentialResponse) Reset()         { *m = CreateCredentialResponse{} }
func (m *CreateCredentialResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCredentialResponse) ProtoMessage()    {}
func (*CreateCredentialResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f34767a917b63eb, []int{1}
}

func (m *CreateCredentialResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCredentialResponse.Unmarshal(m, b)
}
func (m *CreateCredentialResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCredentialResponse.Marshal(b, m, deterministic)
}
func (m *CreateCredentialResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCredentialResponse.Merge(m, src)
}
func (m *CreateCredentialResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCredentialResponse.Size(m)
}
func (m *CreateCredentialResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCredentialResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCredentialResponse proto.InternalMessageInfo

func (m *CreateCredentialResponse) GetCredential() *Credential {
	if m != nil {
		return m.Credential
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateCredentialRequest)(nil), "ai.metathings.service.identityd2.CreateCredentialRequest")
	proto.RegisterType((*CreateCredentialResponse)(nil), "ai.metathings.service.identityd2.CreateCredentialResponse")
}

func init() { proto.RegisterFile("create_credential.proto", fileDescriptor_7f34767a917b63eb) }

var fileDescriptor_7f34767a917b63eb = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xcf, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x80, 0x61, 0xd2, 0x43, 0x0f, 0x9b, 0x5b, 0x2e, 0x0d, 0x15, 0x34, 0xf4, 0xe4, 0xc1, 0x4e,
	0x20, 0x05, 0x4f, 0x9e, 0xec, 0x49, 0x10, 0x84, 0x08, 0xbd, 0x96, 0x4d, 0x32, 0x26, 0x8b, 0x49,
	0x66, 0xdd, 0x9d, 0x34, 0xd8, 0x97, 0xf2, 0x91, 0x04, 0x9f, 0x44, 0xd8, 0xd4, 0x36, 0x60, 0xa1,
	0xb7, 0x85, 0x9d, 0xf9, 0xf9, 0x46, 0xcc, 0x72, 0x83, 0x92, 0x71, 0x9b, 0x1b, 0x2c, 0xb0, 0x65,
	0x25, 0x6b, 0xd0, 0x86, 0x98, 0x82, 0x48, 0x2a, 0x68, 0x90, 0x25, 0x57, 0xaa, 0x2d, 0x2d, 0x58,
	0x34, 0x3b, 0x95, 0x23, 0x28, 0x37, 0xc5, 0x9f, 0x45, 0x32, 0xbf, 0x2e, 0x89, 0xca, 0x1a, 0x63,
	0x37, 0x9f, 0x75, 0x6f, 0x71, 0x6f, 0xa4, 0xd6, 0x68, 0xec, 0x50, 0x98, 0xdf, 0x97, 0x8a, 0xab,
	0x2e, 0x83, 0x9c, 0x9a, 0xb8, 0xe9, 0x15, 0xbf, 0x53, 0x1f, 0x97, 0xb4, 0x74, 0x9f, 0xcb, 0x9d,
	0xac, 0x55, 0x21, 0x99, 0x8c, 0x8d, 0x8f, 0xcf, 0xc3, 0x9e, 0xdf, 0x50, 0x81, 0x07, 0xc6, 0xe2,
	0xcb, 0x13, 0xb3, 0xb5, 0x23, 0xae, 0x8f, 0xc2, 0x14, 0x3f, 0x3a, 0xb4, 0x1c, 0x6c, 0x84, 0x38,
	0xb1, 0x43, 0x2f, 0xf2, 0x6e, 0xfd, 0x04, 0xe0, 0x92, 0x1b, 0x5e, 0xf4, 0x29, 0xf5, 0x38, 0xfd,
	0xf9, 0xbe, 0x99, 0x44, 0x5e, 0x3a, 0x2a, 0x05, 0x0f, 0xc2, 0xb7, 0x98, 0x1b, 0xe4, 0xad, 0x55,
	0x7b, 0x0c, 0x27, 0x2e, 0x7c, 0x05, 0xc3, 0xb9, 0xf0, 0x77, 0x2e, 0x3c, 0xb5, 0xbc, 0x4a, 0x36,
	0xb2, 0xee, 0x30, 0x15, 0xc3, 0xfc, 0xab, 0xda, 0xe3, 0xa2, 0x12, 0xe1, 0x7f, 0xb0, 0xd5, 0xd4,
	0x5a, 0x0c, 0x9e, 0xcf, 0x88, 0xef, 0x2e, 0x8b, 0x47, 0xa5, 0xd1, 0x7e, 0x36, 0x75, 0x94, 0xd5,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x11, 0x97, 0x44, 0x7f, 0xc4, 0x01, 0x00, 0x00,
}
