// Code generated by protoc-gen-go. DO NOT EDIT.
// source: add_entity_to_domain.proto

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

type AddEntityToDomainRequest struct {
	Domain               *OpDomain `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Entity               *OpEntity `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddEntityToDomainRequest) Reset()         { *m = AddEntityToDomainRequest{} }
func (m *AddEntityToDomainRequest) String() string { return proto.CompactTextString(m) }
func (*AddEntityToDomainRequest) ProtoMessage()    {}
func (*AddEntityToDomainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3a988c0435a95df, []int{0}
}

func (m *AddEntityToDomainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddEntityToDomainRequest.Unmarshal(m, b)
}
func (m *AddEntityToDomainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddEntityToDomainRequest.Marshal(b, m, deterministic)
}
func (m *AddEntityToDomainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddEntityToDomainRequest.Merge(m, src)
}
func (m *AddEntityToDomainRequest) XXX_Size() int {
	return xxx_messageInfo_AddEntityToDomainRequest.Size(m)
}
func (m *AddEntityToDomainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddEntityToDomainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddEntityToDomainRequest proto.InternalMessageInfo

func (m *AddEntityToDomainRequest) GetDomain() *OpDomain {
	if m != nil {
		return m.Domain
	}
	return nil
}

func (m *AddEntityToDomainRequest) GetEntity() *OpEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func init() {
	proto.RegisterType((*AddEntityToDomainRequest)(nil), "ai.metathings.service.identityd2.AddEntityToDomainRequest")
}

func init() { proto.RegisterFile("add_entity_to_domain.proto", fileDescriptor_d3a988c0435a95df) }

var fileDescriptor_d3a988c0435a95df = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0x4c, 0x49, 0x89,
	0x4f, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0x8c, 0x2f, 0xc9, 0x8f, 0x4f, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc,
	0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x48, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c,
	0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0xcb, 0x4c,
	0x81, 0xa8, 0x4f, 0x31, 0x92, 0x12, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87,
	0x31, 0x20, 0x5a, 0xa5, 0xb8, 0x73, 0xf3, 0x53, 0x52, 0x73, 0x20, 0x1c, 0xa5, 0x6d, 0x8c, 0x5c,
	0x12, 0x8e, 0x29, 0x29, 0xae, 0x60, 0x5d, 0x21, 0xf9, 0x2e, 0x60, 0x3b, 0x82, 0x52, 0x0b, 0x4b,
	0x53, 0x8b, 0x4b, 0x84, 0x7c, 0xb8, 0xd8, 0x20, 0x96, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b,
	0x69, 0xe9, 0x11, 0xb2, 0x55, 0xcf, 0xbf, 0x00, 0x62, 0x84, 0x13, 0xc7, 0x2f, 0x27, 0xd6, 0x2e,
	0x46, 0x26, 0x01, 0xc6, 0x20, 0xa8, 0x19, 0x20, 0xd3, 0x20, 0xca, 0x24, 0x98, 0x88, 0x37, 0x0d,
	0xe2, 0x30, 0x64, 0xd3, 0x20, 0x92, 0x49, 0x6c, 0x60, 0xf7, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x97, 0x31, 0x0b, 0xc1, 0x25, 0x01, 0x00, 0x00,
}
