// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delete_entity.proto

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

type DeleteEntityRequest struct {
	Entity               *OpEntity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DeleteEntityRequest) Reset()         { *m = DeleteEntityRequest{} }
func (m *DeleteEntityRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteEntityRequest) ProtoMessage()    {}
func (*DeleteEntityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8647a4975befbd14, []int{0}
}

func (m *DeleteEntityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEntityRequest.Unmarshal(m, b)
}
func (m *DeleteEntityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEntityRequest.Marshal(b, m, deterministic)
}
func (m *DeleteEntityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEntityRequest.Merge(m, src)
}
func (m *DeleteEntityRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteEntityRequest.Size(m)
}
func (m *DeleteEntityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEntityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEntityRequest proto.InternalMessageInfo

func (m *DeleteEntityRequest) GetEntity() *OpEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteEntityRequest)(nil), "ai.metathings.service.identityd2.DeleteEntityRequest")
}

func init() { proto.RegisterFile("delete_entity.proto", fileDescriptor_8647a4975befbd14) }

var fileDescriptor_8647a4975befbd14 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x49, 0xcd, 0x49,
	0x2d, 0x49, 0x8d, 0x4f, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x52, 0x48, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0xcb, 0x4c, 0x81, 0xa8, 0x4a, 0x31, 0x92, 0x12, 0x2f,
	0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0x20, 0x5a, 0xa5, 0xb8, 0x73, 0xf3,
	0x53, 0x52, 0x73, 0x20, 0x1c, 0xa5, 0x64, 0x2e, 0x61, 0x17, 0xb0, 0xf1, 0xae, 0x60, 0x7d, 0x41,
	0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x3e, 0x5c, 0x6c, 0x10, 0x83, 0x24, 0x18, 0x15, 0x18,
	0x35, 0xb8, 0x8d, 0xb4, 0xf4, 0x08, 0xd9, 0xa7, 0xe7, 0x5f, 0x00, 0x31, 0xc2, 0x89, 0xe3, 0x97,
	0x13, 0x6b, 0x17, 0x23, 0x93, 0x00, 0x63, 0x10, 0xd4, 0x8c, 0x24, 0x36, 0xb0, 0x5d, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x1d, 0x90, 0x2f, 0xca, 0x00, 0x00, 0x00,
}