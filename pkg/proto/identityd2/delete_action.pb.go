// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delete_action.proto

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

type DeleteActionRequest struct {
	Action               *OpAction `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DeleteActionRequest) Reset()         { *m = DeleteActionRequest{} }
func (m *DeleteActionRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteActionRequest) ProtoMessage()    {}
func (*DeleteActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f2b3e5c939b0545, []int{0}
}

func (m *DeleteActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteActionRequest.Unmarshal(m, b)
}
func (m *DeleteActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteActionRequest.Marshal(b, m, deterministic)
}
func (m *DeleteActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteActionRequest.Merge(m, src)
}
func (m *DeleteActionRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteActionRequest.Size(m)
}
func (m *DeleteActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteActionRequest proto.InternalMessageInfo

func (m *DeleteActionRequest) GetAction() *OpAction {
	if m != nil {
		return m.Action
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteActionRequest)(nil), "ai.metathings.service.identityd2.DeleteActionRequest")
}

func init() { proto.RegisterFile("delete_action.proto", fileDescriptor_3f2b3e5c939b0545) }

var fileDescriptor_3f2b3e5c939b0545 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x49, 0xcd, 0x49,
	0x2d, 0x49, 0x8d, 0x4f, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x52, 0x48, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0xcb, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0x4c,
	0x31, 0x92, 0x32, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x2d,
	0xcf, 0x2c, 0xc9, 0xce, 0x2f, 0xd7, 0x4f, 0xcf, 0xd7, 0x05, 0x6b, 0xd7, 0x2d, 0x4b, 0xcc, 0xc9,
	0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0x2a, 0xd6, 0x87, 0x33, 0x21, 0x26, 0x4b, 0x71, 0xe7, 0xe6, 0xa7,
	0xa4, 0xe6, 0x40, 0x38, 0x4a, 0x89, 0x5c, 0xc2, 0x2e, 0x60, 0xdb, 0x1d, 0xc1, 0x96, 0x07, 0xa5,
	0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x79, 0x71, 0xb1, 0x41, 0x5c, 0x23, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x6d, 0xa4, 0xa5, 0x47, 0xc8, 0x39, 0x7a, 0xfe, 0x05, 0x10, 0x23, 0x9c, 0xd8, 0x1e, 0xdd,
	0x97, 0x67, 0x52, 0x60, 0x0c, 0x82, 0x9a, 0x90, 0xc4, 0x06, 0xb6, 0xc9, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x7d, 0x25, 0x9f, 0xe1, 0xe7, 0x00, 0x00, 0x00,
}
