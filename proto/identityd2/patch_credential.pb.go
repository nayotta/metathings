// Code generated by protoc-gen-go. DO NOT EDIT.
// source: patch_credential.proto

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

type PatchCredentialRequest struct {
	Credential           *OpCredential `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PatchCredentialRequest) Reset()         { *m = PatchCredentialRequest{} }
func (m *PatchCredentialRequest) String() string { return proto.CompactTextString(m) }
func (*PatchCredentialRequest) ProtoMessage()    {}
func (*PatchCredentialRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed2cb3a75ddc78b6, []int{0}
}

func (m *PatchCredentialRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchCredentialRequest.Unmarshal(m, b)
}
func (m *PatchCredentialRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchCredentialRequest.Marshal(b, m, deterministic)
}
func (m *PatchCredentialRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchCredentialRequest.Merge(m, src)
}
func (m *PatchCredentialRequest) XXX_Size() int {
	return xxx_messageInfo_PatchCredentialRequest.Size(m)
}
func (m *PatchCredentialRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchCredentialRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PatchCredentialRequest proto.InternalMessageInfo

func (m *PatchCredentialRequest) GetCredential() *OpCredential {
	if m != nil {
		return m.Credential
	}
	return nil
}

type PatchCredentialResponse struct {
	Credential           *Credential `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PatchCredentialResponse) Reset()         { *m = PatchCredentialResponse{} }
func (m *PatchCredentialResponse) String() string { return proto.CompactTextString(m) }
func (*PatchCredentialResponse) ProtoMessage()    {}
func (*PatchCredentialResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed2cb3a75ddc78b6, []int{1}
}

func (m *PatchCredentialResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchCredentialResponse.Unmarshal(m, b)
}
func (m *PatchCredentialResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchCredentialResponse.Marshal(b, m, deterministic)
}
func (m *PatchCredentialResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchCredentialResponse.Merge(m, src)
}
func (m *PatchCredentialResponse) XXX_Size() int {
	return xxx_messageInfo_PatchCredentialResponse.Size(m)
}
func (m *PatchCredentialResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchCredentialResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PatchCredentialResponse proto.InternalMessageInfo

func (m *PatchCredentialResponse) GetCredential() *Credential {
	if m != nil {
		return m.Credential
	}
	return nil
}

func init() {
	proto.RegisterType((*PatchCredentialRequest)(nil), "ai.metathings.service.identityd2.PatchCredentialRequest")
	proto.RegisterType((*PatchCredentialResponse)(nil), "ai.metathings.service.identityd2.PatchCredentialResponse")
}

func init() { proto.RegisterFile("patch_credential.proto", fileDescriptor_ed2cb3a75ddc78b6) }

var fileDescriptor_ed2cb3a75ddc78b6 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x48, 0x2c, 0x49,
	0xce, 0x88, 0x4f, 0x2e, 0x4a, 0x4d, 0x49, 0xcd, 0x2b, 0xc9, 0x4c, 0xcc, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x52, 0x48, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b,
	0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0xcb, 0x04, 0xab, 0x2a, 0xa9, 0x4c,
	0x31, 0x92, 0x12, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0x20, 0x5a,
	0xa5, 0xb8, 0x73, 0xf3, 0x53, 0x52, 0xa1, 0xe6, 0x28, 0x15, 0x71, 0x89, 0x05, 0x80, 0x6c, 0x70,
	0x86, 0x5b, 0x10, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x14, 0xc1, 0xc5, 0x85, 0xb0, 0x55,
	0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x4f, 0x8f, 0x90, 0xb5, 0x7a, 0xfe, 0x05, 0x08, 0xa3,
	0x9c, 0x38, 0x7e, 0x39, 0xb1, 0x76, 0x31, 0x32, 0x09, 0x30, 0x06, 0x21, 0x99, 0xa5, 0x94, 0xce,
	0x25, 0x8e, 0x61, 0x67, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x0f, 0x16, 0x4b, 0x75, 0x08,
	0x5b, 0x8a, 0x64, 0x12, 0x92, 0xfe, 0x24, 0x36, 0xb0, 0x1f, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xd2, 0x93, 0x0f, 0x24, 0x45, 0x01, 0x00, 0x00,
}