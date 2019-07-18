// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remove_entity_from_domain.proto

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RemoveEntityFromDomainRequest struct {
	Domain               *OpDomain `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Entity               *OpEntity `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RemoveEntityFromDomainRequest) Reset()         { *m = RemoveEntityFromDomainRequest{} }
func (m *RemoveEntityFromDomainRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveEntityFromDomainRequest) ProtoMessage()    {}
func (*RemoveEntityFromDomainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c802fdb451d3192f, []int{0}
}

func (m *RemoveEntityFromDomainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveEntityFromDomainRequest.Unmarshal(m, b)
}
func (m *RemoveEntityFromDomainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveEntityFromDomainRequest.Marshal(b, m, deterministic)
}
func (m *RemoveEntityFromDomainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveEntityFromDomainRequest.Merge(m, src)
}
func (m *RemoveEntityFromDomainRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveEntityFromDomainRequest.Size(m)
}
func (m *RemoveEntityFromDomainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveEntityFromDomainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveEntityFromDomainRequest proto.InternalMessageInfo

func (m *RemoveEntityFromDomainRequest) GetDomain() *OpDomain {
	if m != nil {
		return m.Domain
	}
	return nil
}

func (m *RemoveEntityFromDomainRequest) GetEntity() *OpEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func init() {
	proto.RegisterType((*RemoveEntityFromDomainRequest)(nil), "ai.metathings.service.identityd2.RemoveEntityFromDomainRequest")
}

func init() { proto.RegisterFile("remove_entity_from_domain.proto", fileDescriptor_c802fdb451d3192f) }

var fileDescriptor_c802fdb451d3192f = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x4a, 0xcd, 0xcd,
	0x2f, 0x4b, 0x8d, 0x4f, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0x8c, 0x4f, 0x2b, 0xca, 0xcf, 0x8d, 0x4f,
	0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x48, 0xcc, 0xd4,
	0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0xd5, 0xcb, 0x4c, 0x81, 0xe8, 0x48, 0x31, 0x92, 0x32, 0x4b, 0xcf, 0x2c, 0xc9, 0x28,
	0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x2d, 0xcf, 0x2c, 0xc9, 0xce, 0x2f, 0xd7, 0x4f, 0xcf,
	0xd7, 0x05, 0x6b, 0xd7, 0x2d, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0x2a, 0xd6, 0x87,
	0x33, 0x21, 0x26, 0x4b, 0x71, 0xe7, 0xe6, 0xa7, 0xa4, 0xe6, 0x40, 0x38, 0x4a, 0xdb, 0x19, 0xb9,
	0x64, 0x83, 0xc0, 0x4e, 0x71, 0x05, 0x9b, 0xeb, 0x56, 0x94, 0x9f, 0xeb, 0x02, 0x76, 0x47, 0x50,
	0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x17, 0x17, 0x1b, 0xc4, 0x61, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0xdc, 0x46, 0x5a, 0x7a, 0x84, 0x5c, 0xa6, 0xe7, 0x5f, 0x00, 0x31, 0xc2, 0x89, 0xed, 0xd1,
	0x7d, 0x79, 0x26, 0x05, 0xc6, 0x20, 0xa8, 0x09, 0x20, 0xb3, 0x20, 0x8a, 0x24, 0x98, 0x88, 0x37,
	0x0b, 0xe2, 0x30, 0x84, 0x59, 0x10, 0xa9, 0x24, 0x36, 0xb0, 0x07, 0x8c, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xfb, 0xab, 0x0e, 0x2f, 0x4a, 0x01, 0x00, 0x00,
}
