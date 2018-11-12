// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remove_role_from_group_on_domain.proto

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

type RemoveRoleFromGroupOnDomainRequest struct {
	DomainId             *wrappers.StringValue `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	GroupId              *wrappers.StringValue `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	RoleId               *wrappers.StringValue `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RemoveRoleFromGroupOnDomainRequest) Reset()         { *m = RemoveRoleFromGroupOnDomainRequest{} }
func (m *RemoveRoleFromGroupOnDomainRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRoleFromGroupOnDomainRequest) ProtoMessage()    {}
func (*RemoveRoleFromGroupOnDomainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5399e8618b5fd43, []int{0}
}

func (m *RemoveRoleFromGroupOnDomainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRoleFromGroupOnDomainRequest.Unmarshal(m, b)
}
func (m *RemoveRoleFromGroupOnDomainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRoleFromGroupOnDomainRequest.Marshal(b, m, deterministic)
}
func (m *RemoveRoleFromGroupOnDomainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRoleFromGroupOnDomainRequest.Merge(m, src)
}
func (m *RemoveRoleFromGroupOnDomainRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRoleFromGroupOnDomainRequest.Size(m)
}
func (m *RemoveRoleFromGroupOnDomainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRoleFromGroupOnDomainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRoleFromGroupOnDomainRequest proto.InternalMessageInfo

func (m *RemoveRoleFromGroupOnDomainRequest) GetDomainId() *wrappers.StringValue {
	if m != nil {
		return m.DomainId
	}
	return nil
}

func (m *RemoveRoleFromGroupOnDomainRequest) GetGroupId() *wrappers.StringValue {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func (m *RemoveRoleFromGroupOnDomainRequest) GetRoleId() *wrappers.StringValue {
	if m != nil {
		return m.RoleId
	}
	return nil
}

func init() {
	proto.RegisterType((*RemoveRoleFromGroupOnDomainRequest)(nil), "ai.metathings.service.identityd.RemoveRoleFromGroupOnDomainRequest")
}

func init() {
	proto.RegisterFile("remove_role_from_group_on_domain.proto", fileDescriptor_f5399e8618b5fd43)
}

var fileDescriptor_f5399e8618b5fd43 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0x49, 0x3f, 0x68, 0xfb, 0xc5, 0x5b, 0x4e, 0xa1, 0x88, 0x2d, 0x3d, 0x88, 0x97, 0x6e,
	0x40, 0xc1, 0x9b, 0x88, 0x22, 0x4a, 0x4e, 0x42, 0x04, 0xaf, 0xcb, 0xa6, 0x3b, 0xdd, 0x0e, 0x66,
	0x77, 0xe2, 0x64, 0x93, 0xe0, 0xaf, 0x15, 0xbc, 0xf8, 0x37, 0xa4, 0xbb, 0xe8, 0xbd, 0xb7, 0x81,
	0x79, 0x9f, 0x97, 0x67, 0x26, 0x3d, 0x67, 0xb0, 0x34, 0x80, 0x64, 0x6a, 0x40, 0xee, 0x98, 0xac,
	0x34, 0x4c, 0x7d, 0x2b, 0xc9, 0x49, 0x4d, 0x56, 0xa1, 0x13, 0x2d, 0x93, 0xa7, 0x6c, 0xa9, 0x50,
	0x58, 0xf0, 0xca, 0xef, 0xd1, 0x99, 0x4e, 0x74, 0xc0, 0x03, 0x6e, 0x41, 0xa0, 0x06, 0xe7, 0xd1,
	0x7f, 0xe8, 0xc5, 0x99, 0x21, 0x32, 0x0d, 0x14, 0x21, 0x5e, 0xf7, 0xbb, 0x62, 0x64, 0xd5, 0xb6,
	0xc0, 0x5d, 0x2c, 0x58, 0x5c, 0x1b, 0xf4, 0xfb, 0xbe, 0x16, 0x5b, 0xb2, 0x85, 0x1d, 0xd1, 0xbf,
	0xd1, 0x58, 0x18, 0xda, 0x84, 0xe5, 0x66, 0x50, 0x0d, 0x6a, 0xe5, 0x89, 0xbb, 0xe2, 0x6f, 0x8c,
	0xdc, 0xfa, 0x3b, 0x49, 0xd7, 0x55, 0x70, 0xac, 0xa8, 0x81, 0x47, 0x26, 0xfb, 0x74, 0x10, 0x7c,
	0x76, 0x0f, 0x41, 0xaf, 0x82, 0xf7, 0x1e, 0x3a, 0x9f, 0xdd, 0xa5, 0xff, 0xa3, 0xaf, 0x44, 0x9d,
	0x27, 0xab, 0xe4, 0xe2, 0xe4, 0xf2, 0x54, 0x44, 0x25, 0xf1, 0xab, 0x24, 0x5e, 0x3c, 0xa3, 0x33,
	0xaf, 0xaa, 0xe9, 0xe1, 0x7e, 0xfa, 0xf5, 0xb9, 0x9c, 0xac, 0x92, 0x6a, 0x1e, 0xb1, 0x52, 0x67,
	0xb7, 0xe9, 0x3c, 0xde, 0x8e, 0x3a, 0x9f, 0x1c, 0xd1, 0x30, 0x0b, 0x54, 0xa9, 0xb3, 0x9b, 0x74,
	0x16, 0xde, 0x88, 0x3a, 0xff, 0x77, 0x04, 0x3f, 0x3d, 0x40, 0xa5, 0xae, 0xa7, 0x21, 0x75, 0xf5,
	0x13, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x9e, 0x38, 0xa0, 0x93, 0x01, 0x00, 0x00,
}
