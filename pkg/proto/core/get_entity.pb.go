// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_entity.proto

package core

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

type GetEntityRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetEntityRequest) Reset()         { *m = GetEntityRequest{} }
func (m *GetEntityRequest) String() string { return proto.CompactTextString(m) }
func (*GetEntityRequest) ProtoMessage()    {}
func (*GetEntityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_get_entity_8cba45390b1fa08e, []int{0}
}
func (m *GetEntityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEntityRequest.Unmarshal(m, b)
}
func (m *GetEntityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEntityRequest.Marshal(b, m, deterministic)
}
func (dst *GetEntityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEntityRequest.Merge(dst, src)
}
func (m *GetEntityRequest) XXX_Size() int {
	return xxx_messageInfo_GetEntityRequest.Size(m)
}
func (m *GetEntityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEntityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEntityRequest proto.InternalMessageInfo

func (m *GetEntityRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

type GetEntityResponse struct {
	Entity               *Entity  `protobuf:"bytes,1,opt,name=entity" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEntityResponse) Reset()         { *m = GetEntityResponse{} }
func (m *GetEntityResponse) String() string { return proto.CompactTextString(m) }
func (*GetEntityResponse) ProtoMessage()    {}
func (*GetEntityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_get_entity_8cba45390b1fa08e, []int{1}
}
func (m *GetEntityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEntityResponse.Unmarshal(m, b)
}
func (m *GetEntityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEntityResponse.Marshal(b, m, deterministic)
}
func (dst *GetEntityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEntityResponse.Merge(dst, src)
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
	proto.RegisterType((*GetEntityRequest)(nil), "ai.metathings.service.core.GetEntityRequest")
	proto.RegisterType((*GetEntityResponse)(nil), "ai.metathings.service.core.GetEntityResponse")
}

func init() { proto.RegisterFile("get_entity.proto", fileDescriptor_get_entity_8cba45390b1fa08e) }

var fileDescriptor_get_entity_8cba45390b1fa08e = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0x86, 0x49, 0x8a, 0x14, 0xab, 0xc5, 0x99, 0x4a, 0x82, 0xe8, 0x91, 0xca, 0xe6, 0x66, 0x41,
	0xc5, 0xc2, 0x52, 0x10, 0xed, 0x84, 0x13, 0x6c, 0x65, 0x93, 0x8c, 0x7b, 0x83, 0x49, 0x26, 0xee,
	0x4e, 0x2e, 0xf8, 0xb4, 0x82, 0x4f, 0x22, 0x6c, 0x56, 0x49, 0x73, 0xdd, 0xb2, 0x33, 0xdf, 0xff,
	0x7f, 0xa3, 0x56, 0x16, 0xe5, 0x0d, 0x7b, 0x21, 0xf9, 0x82, 0xc1, 0xb1, 0x70, 0x5e, 0x18, 0x82,
	0x0e, 0xc5, 0xc8, 0x8e, 0x7a, 0xeb, 0xc1, 0xa3, 0xdb, 0x53, 0x8d, 0x50, 0xb3, 0xc3, 0xe2, 0xdc,
	0x32, 0xdb, 0x16, 0x75, 0xd8, 0xac, 0xc6, 0x77, 0x3d, 0x39, 0x33, 0x0c, 0xe8, 0xfc, 0xcc, 0x16,
	0xb7, 0x96, 0x64, 0x37, 0x56, 0x50, 0x73, 0xa7, 0xbb, 0x89, 0xe4, 0x83, 0x27, 0x6d, 0x79, 0x13,
	0x86, 0x9b, 0xbd, 0x69, 0xa9, 0x31, 0xc2, 0xce, 0xeb, 0xff, 0x67, 0xe4, 0x8e, 0x97, 0x06, 0xe5,
	0x93, 0x5a, 0x3d, 0xa2, 0x3c, 0x84, 0xaf, 0x2d, 0x7e, 0x8e, 0xe8, 0x25, 0xbf, 0x51, 0x29, 0x35,
	0xa7, 0xc9, 0x3a, 0xb9, 0x3c, 0xba, 0x3a, 0x83, 0x59, 0x03, 0xfe, 0x34, 0xe0, 0x45, 0x1c, 0xf5,
	0xf6, 0xd5, 0xb4, 0x23, 0xde, 0x67, 0x3f, 0xdf, 0x17, 0xe9, 0x3a, 0xd9, 0xa6, 0xd4, 0x94, 0xcf,
	0xea, 0x64, 0x91, 0xe4, 0x07, 0xee, 0x3d, 0xe6, 0x77, 0x2a, 0x9b, 0xeb, 0x62, 0x5c, 0x09, 0x87,
	0x2f, 0x86, 0xc8, 0x46, 0xa2, 0xca, 0x42, 0xe5, 0xf5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74,
	0x52, 0xca, 0xf4, 0x37, 0x01, 0x00, 0x00,
}