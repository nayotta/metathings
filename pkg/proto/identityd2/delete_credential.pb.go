// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delete_credential.proto

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

type DeleteCredentialRequest struct {
	Credential           *OpCredential `protobuf:"bytes,1,opt,name=credential" json:"credential,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DeleteCredentialRequest) Reset()         { *m = DeleteCredentialRequest{} }
func (m *DeleteCredentialRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCredentialRequest) ProtoMessage()    {}
func (*DeleteCredentialRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_delete_credential_5724c2d6df7d6ac0, []int{0}
}
func (m *DeleteCredentialRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCredentialRequest.Unmarshal(m, b)
}
func (m *DeleteCredentialRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCredentialRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteCredentialRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCredentialRequest.Merge(dst, src)
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

func init() {
	proto.RegisterFile("delete_credential.proto", fileDescriptor_delete_credential_5724c2d6df7d6ac0)
}

var fileDescriptor_delete_credential_5724c2d6df7d6ac0 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x49, 0xcd, 0x49,
	0x2d, 0x49, 0x8d, 0x4f, 0x2e, 0x4a, 0x4d, 0x49, 0xcd, 0x2b, 0xc9, 0x4c, 0xcc, 0xd1, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x52, 0x48, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc,
	0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0xcb, 0x04, 0xab, 0x2a, 0xa9,
	0x4c, 0x31, 0x92, 0x32, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf,
	0x2d, 0xcf, 0x2c, 0xc9, 0xce, 0x2f, 0xd7, 0x4f, 0xcf, 0xd7, 0x05, 0x6b, 0xd7, 0x2d, 0x4b, 0xcc,
	0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0x2a, 0xd6, 0x87, 0x33, 0x21, 0x26, 0x4b, 0x71, 0xe7, 0xe6,
	0xa7, 0xa4, 0x42, 0xad, 0x51, 0x2a, 0xe4, 0x12, 0x77, 0x01, 0xbb, 0xc0, 0x19, 0xee, 0x80, 0xa0,
	0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0xa1, 0x30, 0x2e, 0x2e, 0x84, 0xab, 0x24, 0x18, 0x15, 0x18,
	0x35, 0xb8, 0x8d, 0xf4, 0xf4, 0x08, 0x39, 0x4b, 0xcf, 0xbf, 0x00, 0x61, 0x94, 0x13, 0xdb, 0xa3,
	0xfb, 0xf2, 0x4c, 0x0a, 0x8c, 0x41, 0x48, 0x26, 0x25, 0xb1, 0x81, 0x6d, 0x36, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xd1, 0x96, 0x3d, 0xcf, 0xfb, 0x00, 0x00, 0x00,
}
