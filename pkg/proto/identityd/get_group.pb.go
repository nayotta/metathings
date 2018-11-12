// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_group.proto

package identityd

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type GetGroupRequest struct {
	GroupId              *wrappers.StringValue `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
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

func (m *GetGroupRequest) GetGroupId() *wrappers.StringValue {
	if m != nil {
		return m.GroupId
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
	proto.RegisterType((*GetGroupRequest)(nil), "ai.metathings.service.identityd.GetGroupRequest")
	proto.RegisterType((*GetGroupResponse)(nil), "ai.metathings.service.identityd.GetGroupResponse")
}

func init() { proto.RegisterFile("get_group.proto", fileDescriptor_7be4ddf7f0956146) }

var fileDescriptor_7be4ddf7f0956146 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xbf, 0x4b, 0x3b, 0x41,
	0x10, 0xc5, 0xb9, 0x2f, 0x7c, 0xa3, 0x6c, 0x8a, 0xc8, 0x55, 0x12, 0xc4, 0x84, 0x14, 0x62, 0x93,
	0x59, 0x50, 0xb0, 0x12, 0x04, 0x9b, 0x60, 0x27, 0x27, 0xd8, 0x86, 0xbd, 0xec, 0xb8, 0x19, 0xbc,
	0xbb, 0x59, 0x77, 0x67, 0x73, 0xf8, 0xd7, 0x0a, 0xfe, 0x25, 0xe2, 0x5e, 0xfc, 0xd1, 0xd9, 0x0d,
	0xf3, 0xde, 0xe7, 0xcd, 0x1b, 0x35, 0x71, 0x28, 0x6b, 0x17, 0x38, 0x79, 0xf0, 0x81, 0x85, 0xcb,
	0x99, 0x21, 0x68, 0x51, 0x8c, 0x6c, 0xa9, 0x73, 0x11, 0x22, 0x86, 0x1d, 0x6d, 0x10, 0xc8, 0x62,
	0x27, 0x24, 0xaf, 0x76, 0x7a, 0xea, 0x98, 0x5d, 0x83, 0x3a, 0xdb, 0xeb, 0xf4, 0xa4, 0xfb, 0x60,
	0xbc, 0xc7, 0x10, 0x87, 0x80, 0xe9, 0x95, 0x23, 0xd9, 0xa6, 0x1a, 0x36, 0xdc, 0xea, 0xb6, 0x27,
	0x79, 0xe6, 0x5e, 0x3b, 0x5e, 0x66, 0x71, 0xb9, 0x33, 0x0d, 0x59, 0x23, 0x1c, 0xa2, 0xfe, 0x1e,
	0xf7, 0xdc, 0xf8, 0x57, 0x8b, 0x45, 0xa5, 0x26, 0x2b, 0x94, 0xd5, 0xe7, 0xa6, 0xc2, 0x97, 0x84,
	0x51, 0xca, 0x1b, 0x75, 0x98, 0x1d, 0x6b, 0xb2, 0xc7, 0xc5, 0xbc, 0x38, 0x1f, 0x5f, 0x9c, 0xc0,
	0x50, 0x05, 0xbe, 0xaa, 0xc0, 0x83, 0x04, 0xea, 0xdc, 0xa3, 0x69, 0x12, 0xde, 0x8e, 0xde, 0xdf,
	0x66, 0xff, 0xe6, 0x45, 0x75, 0x90, 0xa9, 0x3b, 0xbb, 0xb8, 0x57, 0x47, 0x3f, 0x99, 0xd1, 0x73,
	0x17, 0xb1, 0xbc, 0x56, 0xff, 0xb3, 0xbc, 0x4f, 0x3c, 0x83, 0x3f, 0xbe, 0x87, 0x01, 0x1f, 0xa0,
	0x7a, 0x94, 0x0f, 0x5f, 0x7e, 0x04, 0x00, 0x00, 0xff, 0xff, 0x60, 0xfb, 0xe5, 0xbf, 0x45, 0x01,
	0x00, 0x00,
}
