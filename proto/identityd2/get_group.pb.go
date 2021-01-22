// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_group.proto

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

type GetGroupRequest struct {
	Group                *OpGroup `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupRequest) Reset()         { *m = GetGroupRequest{} }
func (m *GetGroupRequest) String() string { return proto.CompactTextString(m) }
func (*GetGroupRequest) ProtoMessage()    {}
func (*GetGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7be4ddf7f0956146, []int{0}
}

func (m *GetGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupRequest.Unmarshal(m, b)
}
func (m *GetGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupRequest.Marshal(b, m, deterministic)
}
func (m *GetGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupRequest.Merge(m, src)
}
func (m *GetGroupRequest) XXX_Size() int {
	return xxx_messageInfo_GetGroupRequest.Size(m)
}
func (m *GetGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupRequest proto.InternalMessageInfo

func (m *GetGroupRequest) GetGroup() *OpGroup {
	if m != nil {
		return m.Group
	}
	return nil
}

type GetGroupResponse struct {
	Group                *Group   `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupResponse) Reset()         { *m = GetGroupResponse{} }
func (m *GetGroupResponse) String() string { return proto.CompactTextString(m) }
func (*GetGroupResponse) ProtoMessage()    {}
func (*GetGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7be4ddf7f0956146, []int{1}
}

func (m *GetGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupResponse.Unmarshal(m, b)
}
func (m *GetGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupResponse.Marshal(b, m, deterministic)
}
func (m *GetGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupResponse.Merge(m, src)
}
func (m *GetGroupResponse) XXX_Size() int {
	return xxx_messageInfo_GetGroupResponse.Size(m)
}
func (m *GetGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupResponse proto.InternalMessageInfo

func (m *GetGroupResponse) GetGroup() *Group {
	if m != nil {
		return m.Group
	}
	return nil
}

func init() {
	proto.RegisterType((*GetGroupRequest)(nil), "ai.metathings.service.identityd2.GetGroupRequest")
	proto.RegisterType((*GetGroupResponse)(nil), "ai.metathings.service.identityd2.GetGroupResponse")
}

func init() { proto.RegisterFile("get_group.proto", fileDescriptor_7be4ddf7f0956146) }

var fileDescriptor_7be4ddf7f0956146 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x4f, 0x2d, 0x89,
	0x4f, 0x2f, 0xca, 0x2f, 0x2d, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x48, 0xcc, 0xd4,
	0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0xd5, 0xcb, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0x4c, 0x31, 0x92, 0x12, 0x2f,
	0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0x20, 0x5a, 0xa5, 0xb8, 0x73, 0xf3,
	0x53, 0x52, 0x73, 0x20, 0x1c, 0xa5, 0x18, 0x2e, 0x7e, 0xf7, 0xd4, 0x12, 0x77, 0x90, 0xc9, 0x41,
	0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x9e, 0x5c, 0xac, 0x60, 0x9b, 0x24, 0x18, 0x15, 0x18,
	0x35, 0xb8, 0x8d, 0x34, 0xf5, 0x08, 0x59, 0xa5, 0xe7, 0x5f, 0x00, 0x36, 0xc0, 0x89, 0xe3, 0x97,
	0x13, 0x6b, 0x17, 0x23, 0x93, 0x00, 0x63, 0x10, 0xc4, 0x04, 0xa5, 0x40, 0x2e, 0x01, 0x84, 0xe9,
	0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0xb6, 0xa8, 0xc6, 0xab, 0x13, 0x36, 0x1e, 0xa2, 0x1f,
	0xa2, 0x2b, 0x89, 0x0d, 0xec, 0x6e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x9b, 0x31,
	0x1d, 0x12, 0x01, 0x00, 0x00,
}