// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_entity.proto

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

type GetEntityRequest struct {
	Entity               *OpEntity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetEntityRequest) Reset()         { *m = GetEntityRequest{} }
func (m *GetEntityRequest) String() string { return proto.CompactTextString(m) }
func (*GetEntityRequest) ProtoMessage()    {}
func (*GetEntityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1c9f97a060eb838, []int{0}
}

func (m *GetEntityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEntityRequest.Unmarshal(m, b)
}
func (m *GetEntityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEntityRequest.Marshal(b, m, deterministic)
}
func (m *GetEntityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEntityRequest.Merge(m, src)
}
func (m *GetEntityRequest) XXX_Size() int {
	return xxx_messageInfo_GetEntityRequest.Size(m)
}
func (m *GetEntityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEntityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEntityRequest proto.InternalMessageInfo

func (m *GetEntityRequest) GetEntity() *OpEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type GetEntityResponse struct {
	Entity               *Entity  `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEntityResponse) Reset()         { *m = GetEntityResponse{} }
func (m *GetEntityResponse) String() string { return proto.CompactTextString(m) }
func (*GetEntityResponse) ProtoMessage()    {}
func (*GetEntityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1c9f97a060eb838, []int{1}
}

func (m *GetEntityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEntityResponse.Unmarshal(m, b)
}
func (m *GetEntityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEntityResponse.Marshal(b, m, deterministic)
}
func (m *GetEntityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEntityResponse.Merge(m, src)
}
func (m *GetEntityResponse) XXX_Size() int {
	return xxx_messageInfo_GetEntityResponse.Size(m)
}
func (m *GetEntityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEntityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetEntityResponse proto.InternalMessageInfo

func (m *GetEntityResponse) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func init() {
	proto.RegisterType((*GetEntityRequest)(nil), "ai.metathings.service.identityd2.GetEntityRequest")
	proto.RegisterType((*GetEntityResponse)(nil), "ai.metathings.service.identityd2.GetEntityResponse")
}

func init() { proto.RegisterFile("get_entity.proto", fileDescriptor_e1c9f97a060eb838) }

var fileDescriptor_e1c9f97a060eb838 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x4f, 0x2d, 0x89,
	0x4f, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x48, 0xcc,
	0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0xcb, 0x4c, 0x81, 0xa8, 0x4a, 0x31, 0x92, 0x32, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x2d, 0xcf, 0x2c, 0xc9, 0xce, 0x2f, 0xd7, 0x4f,
	0xcf, 0xd7, 0x05, 0x6b, 0xd7, 0x2d, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0x2a, 0xd6,
	0x87, 0x33, 0x21, 0x26, 0x4b, 0x71, 0xe7, 0xe6, 0xa7, 0xa4, 0xe6, 0x40, 0x38, 0x4a, 0x71, 0x5c,
	0x02, 0xee, 0xa9, 0x25, 0xae, 0x60, 0x33, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xbc,
	0xb8, 0xd8, 0x20, 0x96, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x69, 0xe9, 0x11, 0x72, 0x8b,
	0x9e, 0x7f, 0x01, 0xc4, 0x08, 0x27, 0xb6, 0x47, 0xf7, 0xe5, 0x99, 0x14, 0x18, 0x83, 0xa0, 0x26,
	0x28, 0x85, 0x72, 0x09, 0x22, 0x99, 0x5f, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0xe4, 0x80, 0x66,
	0x81, 0x06, 0x61, 0x0b, 0xa0, 0x26, 0x40, 0xf5, 0x25, 0xb1, 0x81, 0x5d, 0x6f, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x78, 0xda, 0xb5, 0x2c, 0x38, 0x01, 0x00, 0x00,
}
