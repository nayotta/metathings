// Code generated by protoc-gen-go. DO NOT EDIT.
// source: application_credential.proto

package identityd

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ApplicationCredential struct {
	Id                   string                         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string                         `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Secret               string                         `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty"`
	Unrestricted         bool                           `protobuf:"varint,5,opt,name=unrestricted,proto3" json:"unrestricted,omitempty"`
	ProjectId            string                         `protobuf:"bytes,6,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Roles                []*ApplicationCredential__Role `protobuf:"bytes,7,rep,name=roles,proto3" json:"roles,omitempty"`
	ExpiresAt            *timestamp.Timestamp           `protobuf:"bytes,8,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *ApplicationCredential) Reset()         { *m = ApplicationCredential{} }
func (m *ApplicationCredential) String() string { return proto.CompactTextString(m) }
func (*ApplicationCredential) ProtoMessage()    {}
func (*ApplicationCredential) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d833bcee5bfb27f, []int{0}
}

func (m *ApplicationCredential) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationCredential.Unmarshal(m, b)
}
func (m *ApplicationCredential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationCredential.Marshal(b, m, deterministic)
}
func (m *ApplicationCredential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationCredential.Merge(m, src)
}
func (m *ApplicationCredential) XXX_Size() int {
	return xxx_messageInfo_ApplicationCredential.Size(m)
}
func (m *ApplicationCredential) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationCredential.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationCredential proto.InternalMessageInfo

func (m *ApplicationCredential) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ApplicationCredential) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ApplicationCredential) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ApplicationCredential) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *ApplicationCredential) GetUnrestricted() bool {
	if m != nil {
		return m.Unrestricted
	}
	return false
}

func (m *ApplicationCredential) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *ApplicationCredential) GetRoles() []*ApplicationCredential__Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *ApplicationCredential) GetExpiresAt() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiresAt
	}
	return nil
}

type ApplicationCredential__Role struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DomainId             string   `protobuf:"bytes,3,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplicationCredential__Role) Reset()         { *m = ApplicationCredential__Role{} }
func (m *ApplicationCredential__Role) String() string { return proto.CompactTextString(m) }
func (*ApplicationCredential__Role) ProtoMessage()    {}
func (*ApplicationCredential__Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d833bcee5bfb27f, []int{0, 0}
}

func (m *ApplicationCredential__Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationCredential__Role.Unmarshal(m, b)
}
func (m *ApplicationCredential__Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationCredential__Role.Marshal(b, m, deterministic)
}
func (m *ApplicationCredential__Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationCredential__Role.Merge(m, src)
}
func (m *ApplicationCredential__Role) XXX_Size() int {
	return xxx_messageInfo_ApplicationCredential__Role.Size(m)
}
func (m *ApplicationCredential__Role) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationCredential__Role.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationCredential__Role proto.InternalMessageInfo

func (m *ApplicationCredential__Role) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ApplicationCredential__Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ApplicationCredential__Role) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func init() {
	proto.RegisterType((*ApplicationCredential)(nil), "ai.metathings.service.identityd.ApplicationCredential")
	proto.RegisterType((*ApplicationCredential__Role)(nil), "ai.metathings.service.identityd.ApplicationCredential._Role")
}

func init() { proto.RegisterFile("application_credential.proto", fileDescriptor_9d833bcee5bfb27f) }

var fileDescriptor_9d833bcee5bfb27f = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcf, 0x4b, 0x3b, 0x31,
	0x10, 0xc5, 0xd9, 0xfe, 0xfa, 0x76, 0xa7, 0x5f, 0x3c, 0x04, 0x94, 0x50, 0x95, 0x2e, 0x3d, 0xed,
	0x29, 0x85, 0x7a, 0x12, 0xbc, 0x14, 0x2f, 0xf6, 0xba, 0x78, 0x5f, 0xd2, 0xcd, 0x58, 0x47, 0x76,
	0x37, 0x21, 0x99, 0x8a, 0xfe, 0x19, 0xfe, 0xc7, 0xd2, 0x6c, 0xeb, 0x0f, 0x10, 0xf4, 0x96, 0xbc,
	0x79, 0xf3, 0x1e, 0xf3, 0x81, 0x0b, 0xed, 0x5c, 0x4d, 0x95, 0x66, 0xb2, 0x6d, 0x59, 0x79, 0x34,
	0xd8, 0x32, 0xe9, 0x5a, 0x39, 0x6f, 0xd9, 0x8a, 0x99, 0x26, 0xd5, 0x20, 0x6b, 0x7e, 0xa4, 0x76,
	0x1b, 0x54, 0x40, 0xff, 0x4c, 0x15, 0x2a, 0x8a, 0x2e, 0x7e, 0x35, 0xd3, 0xd9, 0xd6, 0xda, 0x6d,
	0x8d, 0x8b, 0x68, 0xdf, 0xec, 0x1e, 0x16, 0x4c, 0x0d, 0x06, 0xd6, 0x8d, 0xeb, 0x12, 0xe6, 0x6f,
	0x7d, 0x38, 0x5d, 0x7d, 0x56, 0xdc, 0x7e, 0x34, 0x88, 0x13, 0xe8, 0x91, 0x91, 0x49, 0x96, 0xe4,
	0x69, 0xd1, 0x23, 0x23, 0x04, 0x0c, 0x5a, 0xdd, 0xa0, 0xec, 0x45, 0x25, 0xbe, 0x45, 0x06, 0x13,
	0x83, 0xa1, 0xf2, 0xe4, 0xf6, 0xcb, 0xb2, 0x1f, 0x47, 0x5f, 0x25, 0x71, 0x06, 0xa3, 0x80, 0x95,
	0x47, 0x96, 0x83, 0x38, 0x3c, 0xfc, 0xc4, 0x1c, 0xfe, 0xef, 0x5a, 0x8f, 0x81, 0x3d, 0x55, 0x8c,
	0x46, 0x0e, 0xb3, 0x24, 0x1f, 0x17, 0xdf, 0x34, 0x71, 0x09, 0xe0, 0xbc, 0x7d, 0xc2, 0x8a, 0x4b,
	0x32, 0x72, 0x14, 0xf7, 0xd3, 0x83, 0xb2, 0x36, 0xa2, 0x80, 0xa1, 0xb7, 0x35, 0x06, 0xf9, 0x2f,
	0xeb, 0xe7, 0x93, 0xe5, 0x8d, 0xfa, 0x05, 0x86, 0xfa, 0xf1, 0x4e, 0x55, 0x16, 0xb6, 0xc6, 0xa2,
	0x8b, 0x12, 0xd7, 0x00, 0xf8, 0xe2, 0xc8, 0x63, 0x28, 0x35, 0xcb, 0x71, 0x96, 0xe4, 0x93, 0xe5,
	0x54, 0x75, 0x10, 0xd5, 0x11, 0xa2, 0xba, 0x3f, 0x42, 0x2c, 0xd2, 0x83, 0x7b, 0xc5, 0xd3, 0x3b,
	0x18, 0xc6, 0xa8, 0x3f, 0x81, 0x3b, 0x87, 0xd4, 0xd8, 0x46, 0x53, 0xbb, 0xbf, 0xac, 0xc3, 0x36,
	0xee, 0x84, 0xb5, 0xd9, 0x8c, 0x62, 0xd1, 0xd5, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0xd3,
	0xb6, 0x02, 0xfc, 0x01, 0x00, 0x00,
}
