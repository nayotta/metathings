// Code generated by protoc-gen-go. DO NOT EDIT.
// source: create_action.proto

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateActionRequest struct {
	Action               *OpAction `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateActionRequest) Reset()         { *m = CreateActionRequest{} }
func (m *CreateActionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateActionRequest) ProtoMessage()    {}
func (*CreateActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d960462c74a0d765, []int{0}
}

func (m *CreateActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateActionRequest.Unmarshal(m, b)
}
func (m *CreateActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateActionRequest.Marshal(b, m, deterministic)
}
func (m *CreateActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateActionRequest.Merge(m, src)
}
func (m *CreateActionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateActionRequest.Size(m)
}
func (m *CreateActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateActionRequest proto.InternalMessageInfo

func (m *CreateActionRequest) GetAction() *OpAction {
	if m != nil {
		return m.Action
	}
	return nil
}

type CreateActionResponse struct {
	Action               *Action  `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateActionResponse) Reset()         { *m = CreateActionResponse{} }
func (m *CreateActionResponse) String() string { return proto.CompactTextString(m) }
func (*CreateActionResponse) ProtoMessage()    {}
func (*CreateActionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d960462c74a0d765, []int{1}
}

func (m *CreateActionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateActionResponse.Unmarshal(m, b)
}
func (m *CreateActionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateActionResponse.Marshal(b, m, deterministic)
}
func (m *CreateActionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateActionResponse.Merge(m, src)
}
func (m *CreateActionResponse) XXX_Size() int {
	return xxx_messageInfo_CreateActionResponse.Size(m)
}
func (m *CreateActionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateActionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateActionResponse proto.InternalMessageInfo

func (m *CreateActionResponse) GetAction() *Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateActionRequest)(nil), "ai.metathings.service.identityd2.CreateActionRequest")
	proto.RegisterType((*CreateActionResponse)(nil), "ai.metathings.service.identityd2.CreateActionResponse")
}

func init() { proto.RegisterFile("create_action.proto", fileDescriptor_d960462c74a0d765) }

var fileDescriptor_d960462c74a0d765 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x2e, 0x4a, 0x4d,
	0x2c, 0x49, 0x8d, 0x4f, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x52, 0x48, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0xcb, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0x4c,
	0x31, 0x92, 0x32, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x2d,
	0xcf, 0x2c, 0xc9, 0xce, 0x2f, 0xd7, 0x4f, 0xcf, 0xd7, 0x05, 0x6b, 0xd7, 0x2d, 0x4b, 0xcc, 0xc9,
	0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0x2a, 0xd6, 0x87, 0x33, 0x21, 0x26, 0x4b, 0x71, 0xe7, 0xe6, 0xa7,
	0xa4, 0xe6, 0x40, 0x38, 0x4a, 0x89, 0x5c, 0xc2, 0xce, 0x60, 0xdb, 0x1d, 0xc1, 0x96, 0x07, 0xa5,
	0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x79, 0x71, 0xb1, 0x41, 0x5c, 0x23, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x6d, 0xa4, 0xa5, 0x47, 0xc8, 0x39, 0x7a, 0xfe, 0x05, 0x10, 0x23, 0x9c, 0xd8, 0x1e, 0xdd,
	0x97, 0x67, 0x52, 0x60, 0x0c, 0x82, 0x9a, 0xa0, 0x14, 0xc1, 0x25, 0x82, 0x6a, 0x45, 0x71, 0x41,
	0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x03, 0x9a, 0x1d, 0x1a, 0x84, 0xed, 0x80, 0x9a, 0x00, 0xd5, 0x97,
	0xc4, 0x06, 0xf6, 0x83, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x19, 0x8b, 0xef, 0x45, 0x41, 0x01,
	0x00, 0x00,
}
