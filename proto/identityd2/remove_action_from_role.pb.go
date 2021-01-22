// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remove_action_from_role.proto

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

type RemoveActionFromRoleRequest struct {
	Action               *OpAction `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Role                 *OpRole   `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RemoveActionFromRoleRequest) Reset()         { *m = RemoveActionFromRoleRequest{} }
func (m *RemoveActionFromRoleRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveActionFromRoleRequest) ProtoMessage()    {}
func (*RemoveActionFromRoleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1023568fe613ac0e, []int{0}
}

func (m *RemoveActionFromRoleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveActionFromRoleRequest.Unmarshal(m, b)
}
func (m *RemoveActionFromRoleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveActionFromRoleRequest.Marshal(b, m, deterministic)
}
func (m *RemoveActionFromRoleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveActionFromRoleRequest.Merge(m, src)
}
func (m *RemoveActionFromRoleRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveActionFromRoleRequest.Size(m)
}
func (m *RemoveActionFromRoleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveActionFromRoleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveActionFromRoleRequest proto.InternalMessageInfo

func (m *RemoveActionFromRoleRequest) GetAction() *OpAction {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *RemoveActionFromRoleRequest) GetRole() *OpRole {
	if m != nil {
		return m.Role
	}
	return nil
}

func init() {
	proto.RegisterType((*RemoveActionFromRoleRequest)(nil), "ai.metathings.service.identityd2.RemoveActionFromRoleRequest")
}

func init() { proto.RegisterFile("remove_action_from_role.proto", fileDescriptor_1023568fe613ac0e) }

var fileDescriptor_1023568fe613ac0e = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x4a, 0xcd, 0xcd,
	0x2f, 0x4b, 0x8d, 0x4f, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0x8b, 0x4f, 0x2b, 0xca, 0xcf, 0x8d, 0x2f,
	0xca, 0xcf, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x48, 0xcc, 0xd4, 0xcb, 0x4d,
	0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e,
	0xd5, 0xcb, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0x4c, 0x31, 0x92, 0x12, 0x2f, 0x4b, 0xcc,
	0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0x20, 0x5a, 0xa5, 0xb8, 0x73, 0xf3, 0x53, 0x52,
	0x73, 0x20, 0x1c, 0xa5, 0xcd, 0x8c, 0x5c, 0xd2, 0x41, 0x60, 0x9b, 0x1c, 0xc1, 0x16, 0xb9, 0x15,
	0xe5, 0xe7, 0x06, 0xe5, 0xe7, 0xa4, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xf9, 0x70,
	0xb1, 0x41, 0x5c, 0x20, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xa4, 0xa5, 0x47, 0xc8, 0x62, 0x3d,
	0xff, 0x02, 0x88, 0x51, 0x4e, 0x1c, 0xbf, 0x9c, 0x58, 0xbb, 0x18, 0x99, 0x04, 0x18, 0x83, 0xa0,
	0x66, 0x08, 0xb9, 0x71, 0xb1, 0x80, 0xfc, 0x20, 0xc1, 0x04, 0x36, 0x4b, 0x83, 0x18, 0xb3, 0x40,
	0x8e, 0x41, 0x32, 0x09, 0xac, 0x3f, 0x89, 0x0d, 0xec, 0x78, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x8e, 0x41, 0x4d, 0xec, 0x25, 0x01, 0x00, 0x00,
}