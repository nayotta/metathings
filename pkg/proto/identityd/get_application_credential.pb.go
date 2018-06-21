// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_application_credential.proto

package identityd

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type GetApplicationCredentialRequest struct {
	UserId                  *wrappers.StringValue `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	ApplicationCredentialId *wrappers.StringValue `protobuf:"bytes,2,opt,name=application_credential_id,json=applicationCredentialId" json:"application_credential_id,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}              `json:"-"`
	XXX_unrecognized        []byte                `json:"-"`
	XXX_sizecache           int32                 `json:"-"`
}

func (m *GetApplicationCredentialRequest) Reset()         { *m = GetApplicationCredentialRequest{} }
func (m *GetApplicationCredentialRequest) String() string { return proto.CompactTextString(m) }
func (*GetApplicationCredentialRequest) ProtoMessage()    {}
func (*GetApplicationCredentialRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_get_application_credential_c69056171f9f27b7, []int{0}
}
func (m *GetApplicationCredentialRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetApplicationCredentialRequest.Unmarshal(m, b)
}
func (m *GetApplicationCredentialRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetApplicationCredentialRequest.Marshal(b, m, deterministic)
}
func (dst *GetApplicationCredentialRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetApplicationCredentialRequest.Merge(dst, src)
}
func (m *GetApplicationCredentialRequest) XXX_Size() int {
	return xxx_messageInfo_GetApplicationCredentialRequest.Size(m)
}
func (m *GetApplicationCredentialRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetApplicationCredentialRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetApplicationCredentialRequest proto.InternalMessageInfo

func (m *GetApplicationCredentialRequest) GetUserId() *wrappers.StringValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *GetApplicationCredentialRequest) GetApplicationCredentialId() *wrappers.StringValue {
	if m != nil {
		return m.ApplicationCredentialId
	}
	return nil
}

type GetApplicationCredentialResponse struct {
	ApplicationCredential *ApplicationCredential `protobuf:"bytes,1,opt,name=application_credential,json=applicationCredential" json:"application_credential,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *GetApplicationCredentialResponse) Reset()         { *m = GetApplicationCredentialResponse{} }
func (m *GetApplicationCredentialResponse) String() string { return proto.CompactTextString(m) }
func (*GetApplicationCredentialResponse) ProtoMessage()    {}
func (*GetApplicationCredentialResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_get_application_credential_c69056171f9f27b7, []int{1}
}
func (m *GetApplicationCredentialResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetApplicationCredentialResponse.Unmarshal(m, b)
}
func (m *GetApplicationCredentialResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetApplicationCredentialResponse.Marshal(b, m, deterministic)
}
func (dst *GetApplicationCredentialResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetApplicationCredentialResponse.Merge(dst, src)
}
func (m *GetApplicationCredentialResponse) XXX_Size() int {
	return xxx_messageInfo_GetApplicationCredentialResponse.Size(m)
}
func (m *GetApplicationCredentialResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetApplicationCredentialResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetApplicationCredentialResponse proto.InternalMessageInfo

func (m *GetApplicationCredentialResponse) GetApplicationCredential() *ApplicationCredential {
	if m != nil {
		return m.ApplicationCredential
	}
	return nil
}

func init() {
	proto.RegisterType((*GetApplicationCredentialRequest)(nil), "ai.metathings.service.identityd.GetApplicationCredentialRequest")
	proto.RegisterType((*GetApplicationCredentialResponse)(nil), "ai.metathings.service.identityd.GetApplicationCredentialResponse")
}

func init() {
	proto.RegisterFile("get_application_credential.proto", fileDescriptor_get_application_credential_c69056171f9f27b7)
}

var fileDescriptor_get_application_credential_c69056171f9f27b7 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x49, 0x17, 0x15, 0xc6, 0x5d, 0xc0, 0xbf, 0x52, 0x4c, 0xe8, 0xca, 0x4d, 0x27, 0xa0,
	0xd0, 0x9d, 0x0b, 0x75, 0x21, 0xdd, 0x56, 0x70, 0x1b, 0x27, 0x99, 0xeb, 0xf4, 0x62, 0x92, 0x3b,
	0xce, 0xdc, 0x34, 0xf8, 0x18, 0xbe, 0x92, 0x2f, 0x22, 0xf8, 0x24, 0xd2, 0xb4, 0xd6, 0xcd, 0xa8,
	0xb8, 0x1b, 0x98, 0x73, 0x3e, 0xbe, 0xc3, 0x15, 0xa9, 0x01, 0xce, 0x95, 0xb5, 0x15, 0x96, 0x8a,
	0x91, 0x9a, 0xbc, 0x74, 0xa0, 0xa1, 0x61, 0x54, 0x95, 0xb4, 0x8e, 0x98, 0xe2, 0x44, 0xa1, 0xac,
	0x81, 0x15, 0x2f, 0xb1, 0x31, 0x5e, 0x7a, 0x70, 0x2b, 0x2c, 0x41, 0x62, 0x9f, 0xe2, 0x17, 0x3d,
	0x3a, 0x35, 0x44, 0xa6, 0x82, 0xac, 0x8f, 0x17, 0xed, 0x63, 0xd6, 0x39, 0x65, 0x2d, 0x38, 0xbf,
	0x01, 0x8c, 0x66, 0x06, 0x79, 0xd9, 0x16, 0xb2, 0xa4, 0x3a, 0xab, 0x3b, 0xe4, 0x27, 0xea, 0x32,
	0x43, 0xd3, 0xfe, 0x73, 0xba, 0x52, 0x15, 0x6a, 0xc5, 0xe4, 0x7c, 0xb6, 0x7b, 0x6e, 0x7b, 0xe3,
	0xdf, 0xb4, 0x26, 0x6f, 0x91, 0x48, 0x6e, 0x81, 0xaf, 0xbe, 0x33, 0x37, 0xbb, 0xc8, 0x02, 0x9e,
	0x5b, 0xf0, 0x1c, 0x5f, 0x8a, 0xbd, 0xd6, 0x83, 0xcb, 0x51, 0x1f, 0x47, 0x69, 0x74, 0xb6, 0x7f,
	0x3e, 0x96, 0x1b, 0x57, 0xf9, 0xe5, 0x2a, 0xef, 0xd8, 0x61, 0x63, 0xee, 0x55, 0xd5, 0xc2, 0xf5,
	0xf0, 0xe3, 0x3d, 0x19, 0xa4, 0xd1, 0x62, 0xb8, 0x2e, 0xcd, 0x75, 0xfc, 0x20, 0x4e, 0xc2, 0x0a,
	0x6b, 0xe0, 0xe0, 0x1f, 0xc0, 0x23, 0x15, 0xb2, 0x9c, 0xeb, 0xc9, 0x6b, 0x24, 0xd2, 0x9f, 0x47,
	0x78, 0x4b, 0x8d, 0x87, 0xb8, 0x16, 0x87, 0x61, 0x8d, 0xed, 0xa8, 0x99, 0xfc, 0xe3, 0x42, 0x32,
	0xcc, 0x3f, 0x08, 0x5a, 0x15, 0xc3, 0x7e, 0xca, 0xc5, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xed,
	0x88, 0x21, 0x22, 0x1a, 0x02, 0x00, 0x00,
}