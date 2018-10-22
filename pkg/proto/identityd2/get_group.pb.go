// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_group.proto

package identityd2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

type GetGroupRequest struct {
	Group                *OpGroup `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupRequest) Reset()         { *m = GetGroupRequest{} }
func (m *GetGroupRequest) String() string { return proto.CompactTextString(m) }
func (*GetGroupRequest) ProtoMessage()    {}
func (*GetGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_get_group_5613d59bef6bb959, []int{0}
}
func (m *GetGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupRequest.Unmarshal(m, b)
}
func (m *GetGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupRequest.Marshal(b, m, deterministic)
}
func (dst *GetGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupRequest.Merge(dst, src)
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
	Group                *Group   `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupResponse) Reset()         { *m = GetGroupResponse{} }
func (m *GetGroupResponse) String() string { return proto.CompactTextString(m) }
func (*GetGroupResponse) ProtoMessage()    {}
func (*GetGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_get_group_5613d59bef6bb959, []int{1}
}
func (m *GetGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupResponse.Unmarshal(m, b)
}
func (m *GetGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupResponse.Marshal(b, m, deterministic)
}
func (dst *GetGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupResponse.Merge(dst, src)
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

func init() { proto.RegisterFile("get_group.proto", fileDescriptor_get_group_5613d59bef6bb959) }

var fileDescriptor_get_group_5613d59bef6bb959 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x4f, 0x2d, 0x89,
	0x4f, 0x2f, 0xca, 0x2f, 0x2d, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x48, 0xcc, 0xd4,
	0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0xd5, 0xcb, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0x4c, 0x31, 0x92, 0x32, 0x4b,
	0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x2d, 0xcf, 0x2c, 0xc9, 0xce,
	0x2f, 0xd7, 0x4f, 0xcf, 0xd7, 0x05, 0x6b, 0xd7, 0x2d, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9,
	0x2f, 0x2a, 0xd6, 0x87, 0x33, 0x21, 0x26, 0x4b, 0x71, 0xe7, 0xe6, 0xa7, 0xa4, 0xe6, 0x40, 0x38,
	0x4a, 0x51, 0x5c, 0xfc, 0xee, 0xa9, 0x25, 0xee, 0x20, 0x8b, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b,
	0x4b, 0x84, 0xdc, 0xb9, 0x58, 0xc1, 0x0e, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0xd2, 0xd4,
	0x23, 0xe4, 0x12, 0x3d, 0xff, 0x02, 0xb0, 0x01, 0x4e, 0x6c, 0x8f, 0xee, 0xcb, 0x33, 0x29, 0x30,
	0x06, 0x41, 0xf4, 0x2b, 0x05, 0x72, 0x09, 0x20, 0xcc, 0x2e, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15,
	0xb2, 0x45, 0x35, 0x5c, 0x9d, 0xb0, 0xe1, 0x10, 0xfd, 0x10, 0x5d, 0x49, 0x6c, 0x60, 0x57, 0x1b,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x93, 0xba, 0xe0, 0x2f, 0x01, 0x00, 0x00,
}
