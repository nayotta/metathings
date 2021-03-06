// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_actions.proto

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

type ListActionsRequest struct {
	Action               *OpAction `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListActionsRequest) Reset()         { *m = ListActionsRequest{} }
func (m *ListActionsRequest) String() string { return proto.CompactTextString(m) }
func (*ListActionsRequest) ProtoMessage()    {}
func (*ListActionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f26ddb2c99b463f, []int{0}
}

func (m *ListActionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListActionsRequest.Unmarshal(m, b)
}
func (m *ListActionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListActionsRequest.Marshal(b, m, deterministic)
}
func (m *ListActionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListActionsRequest.Merge(m, src)
}
func (m *ListActionsRequest) XXX_Size() int {
	return xxx_messageInfo_ListActionsRequest.Size(m)
}
func (m *ListActionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListActionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListActionsRequest proto.InternalMessageInfo

func (m *ListActionsRequest) GetAction() *OpAction {
	if m != nil {
		return m.Action
	}
	return nil
}

type ListActionsResponse struct {
	Actions              []*Action `protobuf:"bytes,1,rep,name=actions,proto3" json:"actions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListActionsResponse) Reset()         { *m = ListActionsResponse{} }
func (m *ListActionsResponse) String() string { return proto.CompactTextString(m) }
func (*ListActionsResponse) ProtoMessage()    {}
func (*ListActionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f26ddb2c99b463f, []int{1}
}

func (m *ListActionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListActionsResponse.Unmarshal(m, b)
}
func (m *ListActionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListActionsResponse.Marshal(b, m, deterministic)
}
func (m *ListActionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListActionsResponse.Merge(m, src)
}
func (m *ListActionsResponse) XXX_Size() int {
	return xxx_messageInfo_ListActionsResponse.Size(m)
}
func (m *ListActionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListActionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListActionsResponse proto.InternalMessageInfo

func (m *ListActionsResponse) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

func init() {
	proto.RegisterType((*ListActionsRequest)(nil), "ai.metathings.service.identityd2.ListActionsRequest")
	proto.RegisterType((*ListActionsResponse)(nil), "ai.metathings.service.identityd2.ListActionsResponse")
}

func init() { proto.RegisterFile("list_actions.proto", fileDescriptor_1f26ddb2c99b463f) }

var fileDescriptor_1f26ddb2c99b463f = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0xc9, 0x2c, 0x2e,
	0x89, 0x4f, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52,
	0x48, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e,
	0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0xcb, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0x4c, 0x31,
	0x92, 0xe2, 0xce, 0xcd, 0x4f, 0x49, 0xcd, 0x81, 0x28, 0x57, 0x8a, 0xe0, 0x12, 0xf2, 0xc9, 0x2c,
	0x2e, 0x71, 0x84, 0x98, 0x11, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0xe4, 0xc4, 0xc5, 0x06,
	0x31, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x4b, 0x8f, 0x90, 0xa9, 0x7a, 0xfe, 0x05,
	0x10, 0x33, 0x82, 0xa0, 0x3a, 0x95, 0x22, 0xb9, 0x84, 0x51, 0x4c, 0x2e, 0x2e, 0xc8, 0xcf, 0x2b,
	0x4e, 0x15, 0x72, 0xe2, 0x62, 0x87, 0x3a, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x83,
	0xb0, 0xd9, 0x50, 0x93, 0x61, 0x1a, 0x93, 0xd8, 0xc0, 0x6e, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x0f, 0x2e, 0x88, 0x97, 0x00, 0x01, 0x00, 0x00,
}
