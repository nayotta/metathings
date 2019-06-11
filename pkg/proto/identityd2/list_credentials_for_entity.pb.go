// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_credentials_for_entity.proto

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ListCredentialsForEntityRequest struct {
	Entity               *OpEntity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListCredentialsForEntityRequest) Reset()         { *m = ListCredentialsForEntityRequest{} }
func (m *ListCredentialsForEntityRequest) String() string { return proto.CompactTextString(m) }
func (*ListCredentialsForEntityRequest) ProtoMessage()    {}
func (*ListCredentialsForEntityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_97760605265d9fcb, []int{0}
}

func (m *ListCredentialsForEntityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCredentialsForEntityRequest.Unmarshal(m, b)
}
func (m *ListCredentialsForEntityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCredentialsForEntityRequest.Marshal(b, m, deterministic)
}
func (m *ListCredentialsForEntityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCredentialsForEntityRequest.Merge(m, src)
}
func (m *ListCredentialsForEntityRequest) XXX_Size() int {
	return xxx_messageInfo_ListCredentialsForEntityRequest.Size(m)
}
func (m *ListCredentialsForEntityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCredentialsForEntityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCredentialsForEntityRequest proto.InternalMessageInfo

func (m *ListCredentialsForEntityRequest) GetEntity() *OpEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type ListCredentialsForEntityResponse struct {
	Credentials          []*Credential `protobuf:"bytes,1,rep,name=credentials,proto3" json:"credentials,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListCredentialsForEntityResponse) Reset()         { *m = ListCredentialsForEntityResponse{} }
func (m *ListCredentialsForEntityResponse) String() string { return proto.CompactTextString(m) }
func (*ListCredentialsForEntityResponse) ProtoMessage()    {}
func (*ListCredentialsForEntityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_97760605265d9fcb, []int{1}
}

func (m *ListCredentialsForEntityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCredentialsForEntityResponse.Unmarshal(m, b)
}
func (m *ListCredentialsForEntityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCredentialsForEntityResponse.Marshal(b, m, deterministic)
}
func (m *ListCredentialsForEntityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCredentialsForEntityResponse.Merge(m, src)
}
func (m *ListCredentialsForEntityResponse) XXX_Size() int {
	return xxx_messageInfo_ListCredentialsForEntityResponse.Size(m)
}
func (m *ListCredentialsForEntityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCredentialsForEntityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCredentialsForEntityResponse proto.InternalMessageInfo

func (m *ListCredentialsForEntityResponse) GetCredentials() []*Credential {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func init() {
	proto.RegisterType((*ListCredentialsForEntityRequest)(nil), "ai.metathings.service.identityd2.ListCredentialsForEntityRequest")
	proto.RegisterType((*ListCredentialsForEntityResponse)(nil), "ai.metathings.service.identityd2.ListCredentialsForEntityResponse")
}

func init() { proto.RegisterFile("list_credentials_for_entity.proto", fileDescriptor_97760605265d9fcb) }

var fileDescriptor_97760605265d9fcb = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x59, 0x84, 0x1e, 0x92, 0xdb, 0x9e, 0x4a, 0x2f, 0xc6, 0x9e, 0x8a, 0xd8, 0x2c, 0xac,
	0xe0, 0x03, 0x28, 0x7a, 0x12, 0x85, 0xbe, 0xc0, 0x92, 0x6e, 0xc6, 0xed, 0x60, 0xb2, 0xb3, 0x66,
	0xa6, 0x2d, 0xbe, 0xbd, 0xb8, 0x29, 0xb6, 0x17, 0xd9, 0xdb, 0x1f, 0xc8, 0xff, 0x7d, 0x7f, 0xa2,
	0x6e, 0x02, 0xb2, 0x34, 0x6d, 0x02, 0x0f, 0xbd, 0xa0, 0x0b, 0xdc, 0x7c, 0x50, 0x6a, 0x7e, 0xb3,
	0x7c, 0xdb, 0x21, 0x91, 0x50, 0x69, 0x1c, 0xda, 0x08, 0xe2, 0x64, 0x87, 0x7d, 0xc7, 0x96, 0x21,
	0x1d, 0xb0, 0x05, 0x8b, 0x3e, 0xdf, 0xf2, 0xf5, 0xe2, 0xa1, 0x43, 0xd9, 0xed, 0xb7, 0xb6, 0xa5,
	0x58, 0xc5, 0x23, 0xca, 0x27, 0x1d, 0xab, 0x8e, 0xd6, 0x63, 0x7d, 0x7d, 0x70, 0x01, 0xbd, 0x13,
	0x4a, 0x5c, 0xfd, 0xc5, 0x4c, 0x5e, 0xe8, 0x48, 0x1e, 0x42, 0x3e, 0x2c, 0x41, 0x5d, 0xbf, 0x22,
	0xcb, 0xd3, 0x79, 0xca, 0x0b, 0xa5, 0xe7, 0x51, 0xb1, 0x81, 0xaf, 0x3d, 0xb0, 0x94, 0x8f, 0x6a,
	0x96, 0x9d, 0xf3, 0xc2, 0x14, 0x2b, 0x5d, 0xdf, 0xda, 0xa9, 0x69, 0xf6, 0x7d, 0x38, 0x21, 0x4e,
	0xcd, 0x65, 0x52, 0xe6, 0x7f, 0x0d, 0x0f, 0xd4, 0x33, 0x94, 0x6f, 0x4a, 0x5f, 0xfc, 0xc8, 0xbc,
	0x30, 0x57, 0x2b, 0x5d, 0xdf, 0x4d, 0xcb, 0xce, 0xd0, 0xcd, 0x25, 0x60, 0x3b, 0x1b, 0x5f, 0x78,
	0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0xea, 0x42, 0xfb, 0x1f, 0x6d, 0x01, 0x00, 0x00,
}
