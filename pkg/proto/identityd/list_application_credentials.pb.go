// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_application_credentials.proto

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

type ListApplicationCredentialsRequest struct {
	UserId               *wrappers.StringValue `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListApplicationCredentialsRequest) Reset()         { *m = ListApplicationCredentialsRequest{} }
func (m *ListApplicationCredentialsRequest) String() string { return proto.CompactTextString(m) }
func (*ListApplicationCredentialsRequest) ProtoMessage()    {}
func (*ListApplicationCredentialsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_application_credentials_a18122049cedcace, []int{0}
}
func (m *ListApplicationCredentialsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListApplicationCredentialsRequest.Unmarshal(m, b)
}
func (m *ListApplicationCredentialsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListApplicationCredentialsRequest.Marshal(b, m, deterministic)
}
func (dst *ListApplicationCredentialsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListApplicationCredentialsRequest.Merge(dst, src)
}
func (m *ListApplicationCredentialsRequest) XXX_Size() int {
	return xxx_messageInfo_ListApplicationCredentialsRequest.Size(m)
}
func (m *ListApplicationCredentialsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListApplicationCredentialsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListApplicationCredentialsRequest proto.InternalMessageInfo

func (m *ListApplicationCredentialsRequest) GetUserId() *wrappers.StringValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

type ListApplicationCredentialsResponse struct {
	ApplicationCredentials []*ApplicationCredential `protobuf:"bytes,1,rep,name=application_credentials,json=applicationCredentials" json:"application_credentials,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                 `json:"-"`
	XXX_unrecognized       []byte                   `json:"-"`
	XXX_sizecache          int32                    `json:"-"`
}

func (m *ListApplicationCredentialsResponse) Reset()         { *m = ListApplicationCredentialsResponse{} }
func (m *ListApplicationCredentialsResponse) String() string { return proto.CompactTextString(m) }
func (*ListApplicationCredentialsResponse) ProtoMessage()    {}
func (*ListApplicationCredentialsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_application_credentials_a18122049cedcace, []int{1}
}
func (m *ListApplicationCredentialsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListApplicationCredentialsResponse.Unmarshal(m, b)
}
func (m *ListApplicationCredentialsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListApplicationCredentialsResponse.Marshal(b, m, deterministic)
}
func (dst *ListApplicationCredentialsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListApplicationCredentialsResponse.Merge(dst, src)
}
func (m *ListApplicationCredentialsResponse) XXX_Size() int {
	return xxx_messageInfo_ListApplicationCredentialsResponse.Size(m)
}
func (m *ListApplicationCredentialsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListApplicationCredentialsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListApplicationCredentialsResponse proto.InternalMessageInfo

func (m *ListApplicationCredentialsResponse) GetApplicationCredentials() []*ApplicationCredential {
	if m != nil {
		return m.ApplicationCredentials
	}
	return nil
}

func init() {
	proto.RegisterType((*ListApplicationCredentialsRequest)(nil), "ai.metathings.service.identityd.ListApplicationCredentialsRequest")
	proto.RegisterType((*ListApplicationCredentialsResponse)(nil), "ai.metathings.service.identityd.ListApplicationCredentialsResponse")
}

func init() {
	proto.RegisterFile("list_application_credentials.proto", fileDescriptor_list_application_credentials_a18122049cedcace)
}

var fileDescriptor_list_application_credentials_a18122049cedcace = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x89, 0x42, 0x85, 0xf4, 0x96, 0x83, 0x96, 0x52, 0x6c, 0xcc, 0xa9, 0x97, 0x6e, 0xa0,
	0x42, 0x6f, 0x1e, 0xd4, 0x93, 0xe0, 0x29, 0x82, 0xd7, 0xb0, 0x49, 0xc6, 0xed, 0xe0, 0x26, 0xb3,
	0xee, 0x4c, 0x1a, 0xfc, 0x21, 0xfe, 0x3e, 0xc1, 0x5f, 0x22, 0xa6, 0xd5, 0x5e, 0xa2, 0xde, 0x06,
	0xde, 0x7b, 0xf3, 0x3e, 0x5e, 0x98, 0x58, 0x64, 0xc9, 0xb5, 0x73, 0x16, 0x4b, 0x2d, 0x48, 0x4d,
	0x5e, 0x7a, 0xa8, 0xa0, 0x11, 0xd4, 0x96, 0x95, 0xf3, 0x24, 0x14, 0xcd, 0x35, 0xaa, 0x1a, 0x44,
	0xcb, 0x06, 0x1b, 0xc3, 0x8a, 0xc1, 0x6f, 0xb1, 0x04, 0x85, 0xbd, 0x4d, 0x5e, 0xab, 0xe9, 0xb9,
	0x21, 0x32, 0x16, 0xd2, 0xde, 0x5e, 0xb4, 0x4f, 0x69, 0xe7, 0xb5, 0x73, 0xe0, 0xf7, 0x0f, 0xa6,
	0x6b, 0x83, 0xb2, 0x69, 0x0b, 0x55, 0x52, 0x9d, 0xd6, 0x1d, 0xca, 0x33, 0x75, 0xa9, 0xa1, 0x65,
	0x2f, 0x2e, 0xb7, 0xda, 0x62, 0xa5, 0x85, 0x3c, 0xa7, 0x3f, 0xe7, 0x3e, 0x37, 0x1b, 0xe6, 0xda,
	0xa9, 0x49, 0x11, 0x5e, 0xdc, 0x23, 0xcb, 0xf5, 0xc1, 0x73, 0x7b, 0x40, 0xcf, 0xe0, 0xa5, 0x05,
	0x96, 0xe8, 0x2a, 0x3c, 0x69, 0x19, 0x7c, 0x8e, 0xd5, 0x24, 0x88, 0x83, 0xc5, 0x78, 0x35, 0x53,
	0x3b, 0x58, 0xf5, 0x0d, 0xab, 0x1e, 0xc4, 0x63, 0x63, 0x1e, 0xb5, 0x6d, 0xe1, 0x66, 0xf4, 0xf1,
	0x3e, 0x3f, 0x8a, 0x83, 0x6c, 0xf4, 0x15, 0xba, 0xab, 0x92, 0xb7, 0x20, 0x4c, 0xfe, 0x2a, 0x61,
	0x47, 0x0d, 0x43, 0x44, 0xe1, 0xd9, 0x2f, 0x13, 0x4e, 0x82, 0xf8, 0x78, 0x31, 0x5e, 0xad, 0xd5,
	0x3f, 0x1b, 0xaa, 0xc1, 0x86, 0xec, 0x54, 0x0f, 0x16, 0x17, 0xa3, 0x9e, 0xfe, 0xf2, 0x33, 0x00,
	0x00, 0xff, 0xff, 0xea, 0xe8, 0x4f, 0xf2, 0xbf, 0x01, 0x00, 0x00,
}
