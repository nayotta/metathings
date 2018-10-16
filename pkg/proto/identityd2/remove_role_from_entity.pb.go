// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remove_role_from_entity.proto

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

type RemoveRoleFromEntityRequest struct {
	Entity               *OpEntity `protobuf:"bytes,1,opt,name=entity" json:"entity,omitempty"`
	Role                 *OpRole   `protobuf:"bytes,2,opt,name=role" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RemoveRoleFromEntityRequest) Reset()         { *m = RemoveRoleFromEntityRequest{} }
func (m *RemoveRoleFromEntityRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRoleFromEntityRequest) ProtoMessage()    {}
func (*RemoveRoleFromEntityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_remove_role_from_entity_263c38fd4883b8c7, []int{0}
}
func (m *RemoveRoleFromEntityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRoleFromEntityRequest.Unmarshal(m, b)
}
func (m *RemoveRoleFromEntityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRoleFromEntityRequest.Marshal(b, m, deterministic)
}
func (dst *RemoveRoleFromEntityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRoleFromEntityRequest.Merge(dst, src)
}
func (m *RemoveRoleFromEntityRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRoleFromEntityRequest.Size(m)
}
func (m *RemoveRoleFromEntityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRoleFromEntityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRoleFromEntityRequest proto.InternalMessageInfo

func (m *RemoveRoleFromEntityRequest) GetEntity() *OpEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *RemoveRoleFromEntityRequest) GetRole() *OpRole {
	if m != nil {
		return m.Role
	}
	return nil
}

func init() {
	proto.RegisterType((*RemoveRoleFromEntityRequest)(nil), "ai.metathings.service.identityd2.RemoveRoleFromEntityRequest")
}

func init() {
	proto.RegisterFile("remove_role_from_entity.proto", fileDescriptor_remove_role_from_entity_263c38fd4883b8c7)
}

var fileDescriptor_remove_role_from_entity_263c38fd4883b8c7 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x4a, 0xcd, 0xcd,
	0x2f, 0x4b, 0x8d, 0x2f, 0xca, 0xcf, 0x49, 0x8d, 0x4f, 0x2b, 0xca, 0xcf, 0x8d, 0x4f, 0xcd, 0x2b,
	0xc9, 0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x48, 0xcc, 0xd4, 0xcb, 0x4d,
	0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e,
	0xd5, 0xcb, 0x4c, 0x81, 0xa8, 0x4a, 0x31, 0x92, 0x32, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2,
	0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x2d, 0xcf, 0x2c, 0xc9, 0xce, 0x2f, 0xd7, 0x4f, 0xcf, 0xd7, 0x05,
	0x6b, 0xd7, 0x2d, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0x2a, 0xd6, 0x87, 0x33, 0x21,
	0x26, 0x4b, 0x71, 0xe7, 0xe6, 0xa7, 0xa4, 0xe6, 0x40, 0x38, 0x4a, 0xeb, 0x19, 0xb9, 0xa4, 0x83,
	0xc0, 0x0e, 0x09, 0xca, 0xcf, 0x49, 0x75, 0x2b, 0xca, 0xcf, 0x75, 0x05, 0x9b, 0x1f, 0x94, 0x5a,
	0x58, 0x9a, 0x5a, 0x5c, 0x22, 0xe4, 0xc5, 0xc5, 0x06, 0xb1, 0x50, 0x82, 0x51, 0x81, 0x51, 0x83,
	0xdb, 0x48, 0x4b, 0x8f, 0x90, 0xbb, 0xf4, 0xfc, 0x0b, 0x20, 0x46, 0x38, 0xb1, 0x3d, 0xba, 0x2f,
	0xcf, 0xa4, 0xc0, 0x18, 0x04, 0x35, 0x41, 0xc8, 0x85, 0x8b, 0x05, 0xe4, 0x59, 0x09, 0x26, 0xb0,
	0x49, 0x1a, 0xc4, 0x98, 0x04, 0x72, 0x14, 0xdc, 0x1c, 0xb0, 0xee, 0x24, 0x36, 0xb0, 0xc3, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x4c, 0xbd, 0x73, 0x40, 0x01, 0x00, 0x00,
}
