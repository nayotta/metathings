// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_entity.proto

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
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x4f, 0x2d, 0x89,
	0x4f, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x48, 0xcc,
	0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0xcb, 0x4c, 0x81, 0xa8, 0x4a, 0x31, 0x92, 0x12, 0x2f, 0x4b, 0xcc, 0xc9,
	0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0x20, 0x5a, 0xa5, 0xb8, 0x73, 0xf3, 0x53, 0x52, 0x73,
	0x20, 0x1c, 0xa5, 0x04, 0x2e, 0x01, 0xf7, 0xd4, 0x12, 0x57, 0xb0, 0xa6, 0xa0, 0xd4, 0xc2, 0xd2,
	0xd4, 0xe2, 0x12, 0x21, 0x1f, 0x2e, 0x36, 0x88, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46,
	0x5a, 0x7a, 0x84, 0x2c, 0xd3, 0xf3, 0x2f, 0x80, 0x18, 0xe1, 0xc4, 0xf1, 0xcb, 0x89, 0xb5, 0x8b,
	0x91, 0x49, 0x80, 0x31, 0x08, 0x6a, 0x86, 0x52, 0x28, 0x97, 0x20, 0x92, 0x0d, 0xc5, 0x05, 0xf9,
	0x79, 0xc5, 0xa9, 0x42, 0x0e, 0x68, 0x56, 0x68, 0x10, 0xb6, 0x02, 0x6a, 0x02, 0x54, 0x5f, 0x12,
	0x1b, 0xd8, 0xfd, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0xc9, 0xc3, 0x59, 0x1b, 0x01,
	0x00, 0x00,
}
