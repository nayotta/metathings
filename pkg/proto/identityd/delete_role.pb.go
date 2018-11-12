// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delete_role.proto

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

type DeleteRoleRequest struct {
	RoleId               *wrappers.StringValue `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DeleteRoleRequest) Reset()         { *m = DeleteRoleRequest{} }
func (m *DeleteRoleRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRoleRequest) ProtoMessage()    {}
func (*DeleteRoleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a89897a4ed4a653, []int{0}
}

func (m *DeleteRoleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRoleRequest.Unmarshal(m, b)
}
func (m *DeleteRoleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRoleRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRoleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRoleRequest.Merge(m, src)
}
func (m *DeleteRoleRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRoleRequest.Size(m)
}
func (m *DeleteRoleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRoleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRoleRequest proto.InternalMessageInfo

func (m *DeleteRoleRequest) GetRoleId() *wrappers.StringValue {
	if m != nil {
		return m.RoleId
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteRoleRequest)(nil), "ai.metathings.service.identityd.DeleteRoleRequest")
}

func init() { proto.RegisterFile("delete_role.proto", fileDescriptor_1a89897a4ed4a653) }

var fileDescriptor_1a89897a4ed4a653 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0x31, 0x4b, 0xc4, 0x40,
	0x10, 0x85, 0x89, 0x45, 0x84, 0x58, 0xdd, 0x55, 0x72, 0x88, 0x77, 0x58, 0xd9, 0x64, 0x16, 0x14,
	0xec, 0x6c, 0xc4, 0xc6, 0x76, 0x05, 0xdb, 0xb0, 0xc9, 0x8e, 0x9b, 0xc1, 0x4d, 0x26, 0xee, 0xce,
	0x26, 0xf8, 0x6b, 0x05, 0x7f, 0x89, 0x64, 0x83, 0x76, 0x03, 0xf3, 0xde, 0xfb, 0xbe, 0x6a, 0x67,
	0xd1, 0xa3, 0x60, 0x13, 0xd8, 0x23, 0x4c, 0x81, 0x85, 0xf7, 0x47, 0x43, 0x30, 0xa0, 0x18, 0xe9,
	0x69, 0x74, 0x11, 0x22, 0x86, 0x99, 0x3a, 0x04, 0xb2, 0x38, 0x0a, 0xc9, 0x97, 0x3d, 0x5c, 0x3b,
	0x66, 0xe7, 0x51, 0xe5, 0x78, 0x9b, 0xde, 0xd5, 0x12, 0xcc, 0x34, 0x61, 0x88, 0xdb, 0xc0, 0xe1,
	0xc1, 0x91, 0xf4, 0xa9, 0x85, 0x8e, 0x07, 0x35, 0x2c, 0x24, 0x1f, 0xbc, 0x28, 0xc7, 0x75, 0x7e,
	0xd6, 0xb3, 0xf1, 0x64, 0x8d, 0x70, 0x88, 0xea, 0xff, 0xdc, 0x7a, 0x37, 0xba, 0xda, 0x3d, 0x67,
	0x1b, 0xcd, 0x1e, 0x35, 0x7e, 0x26, 0x8c, 0xb2, 0x7f, 0xac, 0xce, 0x57, 0xb7, 0x86, 0xec, 0x65,
	0x71, 0x2a, 0x6e, 0x2f, 0xee, 0xae, 0x60, 0xc3, 0xc3, 0x1f, 0x1e, 0x5e, 0x25, 0xd0, 0xe8, 0xde,
	0x8c, 0x4f, 0xf8, 0x54, 0xfe, 0x7c, 0x1f, 0xcf, 0x4e, 0x85, 0x2e, 0xd7, 0xd2, 0x8b, 0x6d, 0xcb,
	0x9c, 0xba, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x12, 0x80, 0x96, 0x93, 0xe8, 0x00, 0x00, 0x00,
}
