// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_credential.proto

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

type GetCredentialRequest struct {
	Credential           *OpCredential `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetCredentialRequest) Reset()         { *m = GetCredentialRequest{} }
func (m *GetCredentialRequest) String() string { return proto.CompactTextString(m) }
func (*GetCredentialRequest) ProtoMessage()    {}
func (*GetCredentialRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cfd402ca0433cf9, []int{0}
}

func (m *GetCredentialRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCredentialRequest.Unmarshal(m, b)
}
func (m *GetCredentialRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCredentialRequest.Marshal(b, m, deterministic)
}
func (m *GetCredentialRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCredentialRequest.Merge(m, src)
}
func (m *GetCredentialRequest) XXX_Size() int {
	return xxx_messageInfo_GetCredentialRequest.Size(m)
}
func (m *GetCredentialRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCredentialRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCredentialRequest proto.InternalMessageInfo

func (m *GetCredentialRequest) GetCredential() *OpCredential {
	if m != nil {
		return m.Credential
	}
	return nil
}

type GetCredentialResponse struct {
	Credential           *Credential `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetCredentialResponse) Reset()         { *m = GetCredentialResponse{} }
func (m *GetCredentialResponse) String() string { return proto.CompactTextString(m) }
func (*GetCredentialResponse) ProtoMessage()    {}
func (*GetCredentialResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cfd402ca0433cf9, []int{1}
}

func (m *GetCredentialResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCredentialResponse.Unmarshal(m, b)
}
func (m *GetCredentialResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCredentialResponse.Marshal(b, m, deterministic)
}
func (m *GetCredentialResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCredentialResponse.Merge(m, src)
}
func (m *GetCredentialResponse) XXX_Size() int {
	return xxx_messageInfo_GetCredentialResponse.Size(m)
}
func (m *GetCredentialResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCredentialResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCredentialResponse proto.InternalMessageInfo

func (m *GetCredentialResponse) GetCredential() *Credential {
	if m != nil {
		return m.Credential
	}
	return nil
}

func init() {
	proto.RegisterType((*GetCredentialRequest)(nil), "ai.metathings.service.identityd2.GetCredentialRequest")
	proto.RegisterType((*GetCredentialResponse)(nil), "ai.metathings.service.identityd2.GetCredentialResponse")
}

func init() { proto.RegisterFile("get_credential.proto", fileDescriptor_0cfd402ca0433cf9) }

var fileDescriptor_0cfd402ca0433cf9 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x4f, 0x2d, 0x89,
	0x4f, 0x2e, 0x4a, 0x4d, 0x49, 0xcd, 0x2b, 0xc9, 0x4c, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x52, 0x48, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6,
	0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0xcb, 0x04, 0xab, 0x2a, 0xa9, 0x4c, 0x31, 0x92,
	0x12, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0x20, 0x5a, 0xa5, 0xb8,
	0x73, 0xf3, 0x53, 0x52, 0xa1, 0xe6, 0x28, 0x15, 0x70, 0x89, 0xb8, 0xa7, 0x96, 0x38, 0xc3, 0x8d,
	0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x8a, 0xe0, 0xe2, 0x42, 0xd8, 0x29, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0x6d, 0xa4, 0xa7, 0x47, 0xc8, 0x52, 0x3d, 0xff, 0x02, 0x84, 0x51, 0x4e, 0x1c,
	0xbf, 0x9c, 0x58, 0xbb, 0x18, 0x99, 0x04, 0x18, 0x83, 0x90, 0xcc, 0x52, 0x4a, 0xe5, 0x12, 0x45,
	0xb3, 0xb1, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0xc8, 0x07, 0x8b, 0x95, 0x3a, 0x84, 0xad, 0x44,
	0x32, 0x09, 0x49, 0x7f, 0x12, 0x1b, 0xd8, 0x7f, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69,
	0xad, 0xee, 0xe9, 0x3f, 0x01, 0x00, 0x00,
}
