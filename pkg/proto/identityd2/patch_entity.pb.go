// Code generated by protoc-gen-go. DO NOT EDIT.
// source: patch_entity.proto

package identityd2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
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

type PatchEntityRequest struct {
	Entity               *OpEntity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PatchEntityRequest) Reset()         { *m = PatchEntityRequest{} }
func (m *PatchEntityRequest) String() string { return proto.CompactTextString(m) }
func (*PatchEntityRequest) ProtoMessage()    {}
func (*PatchEntityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1bd139a219a3419, []int{0}
}

func (m *PatchEntityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchEntityRequest.Unmarshal(m, b)
}
func (m *PatchEntityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchEntityRequest.Marshal(b, m, deterministic)
}
func (m *PatchEntityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchEntityRequest.Merge(m, src)
}
func (m *PatchEntityRequest) XXX_Size() int {
	return xxx_messageInfo_PatchEntityRequest.Size(m)
}
func (m *PatchEntityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchEntityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PatchEntityRequest proto.InternalMessageInfo

func (m *PatchEntityRequest) GetEntity() *OpEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type PatchEntityResponse struct {
	Entity               *Entity  `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PatchEntityResponse) Reset()         { *m = PatchEntityResponse{} }
func (m *PatchEntityResponse) String() string { return proto.CompactTextString(m) }
func (*PatchEntityResponse) ProtoMessage()    {}
func (*PatchEntityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1bd139a219a3419, []int{1}
}

func (m *PatchEntityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchEntityResponse.Unmarshal(m, b)
}
func (m *PatchEntityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchEntityResponse.Marshal(b, m, deterministic)
}
func (m *PatchEntityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchEntityResponse.Merge(m, src)
}
func (m *PatchEntityResponse) XXX_Size() int {
	return xxx_messageInfo_PatchEntityResponse.Size(m)
}
func (m *PatchEntityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchEntityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PatchEntityResponse proto.InternalMessageInfo

func (m *PatchEntityResponse) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func init() {
	proto.RegisterType((*PatchEntityRequest)(nil), "ai.metathings.service.identityd2.PatchEntityRequest")
	proto.RegisterType((*PatchEntityResponse)(nil), "ai.metathings.service.identityd2.PatchEntityResponse")
}

func init() { proto.RegisterFile("patch_entity.proto", fileDescriptor_c1bd139a219a3419) }

var fileDescriptor_c1bd139a219a3419 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xb1, 0x4b, 0xc4, 0x30,
	0x14, 0xc6, 0xa9, 0x43, 0x87, 0xde, 0x16, 0x17, 0xb9, 0x41, 0xcb, 0x4d, 0x87, 0x70, 0x09, 0x9c,
	0xe0, 0x2c, 0x82, 0x8b, 0x8b, 0x72, 0x8b, 0xa3, 0xa6, 0xed, 0x33, 0x7d, 0xd8, 0xf4, 0xc5, 0xe4,
	0xf5, 0x8a, 0x7f, 0xad, 0xe0, 0x5f, 0x22, 0x26, 0x41, 0xd0, 0xa5, 0x5b, 0x42, 0xbe, 0xef, 0xf7,
	0xfb, 0x52, 0x09, 0xa7, 0xb9, 0xed, 0x9f, 0x61, 0x64, 0xe4, 0x0f, 0xe9, 0x3c, 0x31, 0x89, 0x5a,
	0xa3, 0xb4, 0xc0, 0x9a, 0x7b, 0x1c, 0x4d, 0x90, 0x01, 0xfc, 0x11, 0x5b, 0x90, 0xd8, 0xa5, 0x54,
	0xb7, 0x5f, 0x9f, 0x1b, 0x22, 0x33, 0x80, 0x8a, 0xf9, 0x66, 0x7a, 0x55, 0xb3, 0xd7, 0xce, 0x81,
	0x0f, 0x89, 0xb0, 0xbe, 0x36, 0xc8, 0xfd, 0xd4, 0xc8, 0x96, 0xac, 0xb2, 0x33, 0xf2, 0x1b, 0xcd,
	0xca, 0xd0, 0x2e, 0x3e, 0xee, 0x8e, 0x7a, 0xc0, 0x4e, 0x33, 0xf9, 0xa0, 0x7e, 0x8f, 0xb9, 0xb7,
	0xb2, 0xd4, 0xc1, 0x90, 0x2e, 0x9b, 0x97, 0x4a, 0x3c, 0xfe, 0x8c, 0xbb, 0x8b, 0xd6, 0x03, 0xbc,
	0x4f, 0x10, 0x58, 0xdc, 0x57, 0x65, 0x9a, 0x71, 0x56, 0xd4, 0xc5, 0x76, 0xb5, 0xbf, 0x94, 0x4b,
	0x6b, 0xe5, 0x83, 0x4b, 0x88, 0xdb, 0xf2, 0xeb, 0xf3, 0xe2, 0xa4, 0x2e, 0x0e, 0x99, 0xb0, 0x79,
	0xaa, 0x4e, 0xff, 0x18, 0x82, 0xa3, 0x31, 0x80, 0xb8, 0xf9, 0xa7, 0xd8, 0x2e, 0x2b, 0x32, 0x21,
	0xf7, 0x9a, 0x32, 0xfe, 0xe0, 0xea, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xae, 0xe8, 0x55, 0x7f, 0x5e,
	0x01, 0x00, 0x00,
}
