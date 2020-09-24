// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_entities.proto

package identityd2

import (
	fmt "fmt"
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

type ListEntitiesRequest struct {
	Entity               *OpEntity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListEntitiesRequest) Reset()         { *m = ListEntitiesRequest{} }
func (m *ListEntitiesRequest) String() string { return proto.CompactTextString(m) }
func (*ListEntitiesRequest) ProtoMessage()    {}
func (*ListEntitiesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_16fafe5a45308a83, []int{0}
}

func (m *ListEntitiesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEntitiesRequest.Unmarshal(m, b)
}
func (m *ListEntitiesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEntitiesRequest.Marshal(b, m, deterministic)
}
func (m *ListEntitiesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEntitiesRequest.Merge(m, src)
}
func (m *ListEntitiesRequest) XXX_Size() int {
	return xxx_messageInfo_ListEntitiesRequest.Size(m)
}
func (m *ListEntitiesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEntitiesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListEntitiesRequest proto.InternalMessageInfo

func (m *ListEntitiesRequest) GetEntity() *OpEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type ListEntitiesResponse struct {
	Entities             []*Entity `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListEntitiesResponse) Reset()         { *m = ListEntitiesResponse{} }
func (m *ListEntitiesResponse) String() string { return proto.CompactTextString(m) }
func (*ListEntitiesResponse) ProtoMessage()    {}
func (*ListEntitiesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_16fafe5a45308a83, []int{1}
}

func (m *ListEntitiesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEntitiesResponse.Unmarshal(m, b)
}
func (m *ListEntitiesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEntitiesResponse.Marshal(b, m, deterministic)
}
func (m *ListEntitiesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEntitiesResponse.Merge(m, src)
}
func (m *ListEntitiesResponse) XXX_Size() int {
	return xxx_messageInfo_ListEntitiesResponse.Size(m)
}
func (m *ListEntitiesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEntitiesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListEntitiesResponse proto.InternalMessageInfo

func (m *ListEntitiesResponse) GetEntities() []*Entity {
	if m != nil {
		return m.Entities
	}
	return nil
}

func init() {
	proto.RegisterType((*ListEntitiesRequest)(nil), "ai.metathings.service.identityd2.ListEntitiesRequest")
	proto.RegisterType((*ListEntitiesResponse)(nil), "ai.metathings.service.identityd2.ListEntitiesResponse")
}

func init() { proto.RegisterFile("list_entities.proto", fileDescriptor_16fafe5a45308a83) }

var fileDescriptor_16fafe5a45308a83 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0xc9, 0x2c, 0x2e,
	0x89, 0x4f, 0xcd, 0x2b, 0xc9, 0x2c, 0xc9, 0x4c, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x52, 0x48, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0xcb, 0x4c, 0x01, 0xab, 0xab, 0x4c, 0x31, 0x92, 0xe2,
	0xce, 0xcd, 0x4f, 0x49, 0xcd, 0x81, 0x28, 0x57, 0x8a, 0xe4, 0x12, 0xf6, 0xc9, 0x2c, 0x2e, 0x71,
	0x85, 0x1a, 0x12, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0xe4, 0xc4, 0xc5, 0x06, 0x51, 0x2f,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xa4, 0xa5, 0x47, 0xc8, 0x58, 0x3d, 0xff, 0x02, 0xb0, 0x21,
	0x95, 0x41, 0x50, 0x9d, 0x4a, 0x31, 0x5c, 0x22, 0xa8, 0x46, 0x17, 0x17, 0xe4, 0xe7, 0x15, 0xa7,
	0x0a, 0xb9, 0x70, 0x71, 0xc0, 0xdc, 0x2c, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0xa4, 0x41, 0xd8,
	0x74, 0xa8, 0xd9, 0x70, 0x9d, 0x49, 0x6c, 0x60, 0xf7, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xe7, 0xbf, 0x47, 0xad, 0x05, 0x01, 0x00, 0x00,
}
