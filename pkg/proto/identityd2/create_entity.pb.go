// Code generated by protoc-gen-go. DO NOT EDIT.
// source: create_entity.proto

package identityd2

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

type CreateEntityRequest struct {
	Id                   *wrappers.StringValue            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Domains              []*OpDomain                      `protobuf:"bytes,2,rep,name=domains,proto3" json:"domains,omitempty"`
	Groups               []*OpGroup                       `protobuf:"bytes,3,rep,name=groups,proto3" json:"groups,omitempty"`
	Roles                []*OpRole                        `protobuf:"bytes,4,rep,name=roles,proto3" json:"roles,omitempty"`
	Name                 *wrappers.StringValue            `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Alias                *wrappers.StringValue            `protobuf:"bytes,6,opt,name=alias,proto3" json:"alias,omitempty"`
	Password             *wrappers.StringValue            `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
	Extra                map[string]*wrappers.StringValue `protobuf:"bytes,8,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *CreateEntityRequest) Reset()         { *m = CreateEntityRequest{} }
func (m *CreateEntityRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEntityRequest) ProtoMessage()    {}
func (*CreateEntityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7096d1c9da9b69b6, []int{0}
}

func (m *CreateEntityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEntityRequest.Unmarshal(m, b)
}
func (m *CreateEntityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEntityRequest.Marshal(b, m, deterministic)
}
func (m *CreateEntityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEntityRequest.Merge(m, src)
}
func (m *CreateEntityRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEntityRequest.Size(m)
}
func (m *CreateEntityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEntityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEntityRequest proto.InternalMessageInfo

func (m *CreateEntityRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *CreateEntityRequest) GetDomains() []*OpDomain {
	if m != nil {
		return m.Domains
	}
	return nil
}

func (m *CreateEntityRequest) GetGroups() []*OpGroup {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *CreateEntityRequest) GetRoles() []*OpRole {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *CreateEntityRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CreateEntityRequest) GetAlias() *wrappers.StringValue {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *CreateEntityRequest) GetPassword() *wrappers.StringValue {
	if m != nil {
		return m.Password
	}
	return nil
}

func (m *CreateEntityRequest) GetExtra() map[string]*wrappers.StringValue {
	if m != nil {
		return m.Extra
	}
	return nil
}

type CreateEntityResponse struct {
	Entity               *Entity  `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEntityResponse) Reset()         { *m = CreateEntityResponse{} }
func (m *CreateEntityResponse) String() string { return proto.CompactTextString(m) }
func (*CreateEntityResponse) ProtoMessage()    {}
func (*CreateEntityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7096d1c9da9b69b6, []int{1}
}

func (m *CreateEntityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEntityResponse.Unmarshal(m, b)
}
func (m *CreateEntityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEntityResponse.Marshal(b, m, deterministic)
}
func (m *CreateEntityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEntityResponse.Merge(m, src)
}
func (m *CreateEntityResponse) XXX_Size() int {
	return xxx_messageInfo_CreateEntityResponse.Size(m)
}
func (m *CreateEntityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEntityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEntityResponse proto.InternalMessageInfo

func (m *CreateEntityResponse) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateEntityRequest)(nil), "ai.metathings.service.identityd2.CreateEntityRequest")
	proto.RegisterMapType((map[string]*wrappers.StringValue)(nil), "ai.metathings.service.identityd2.CreateEntityRequest.ExtraEntry")
	proto.RegisterType((*CreateEntityResponse)(nil), "ai.metathings.service.identityd2.CreateEntityResponse")
}

func init() { proto.RegisterFile("create_entity.proto", fileDescriptor_7096d1c9da9b69b6) }

var fileDescriptor_7096d1c9da9b69b6 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x8b, 0xd4, 0x30,
	0x18, 0x86, 0x99, 0xce, 0xb4, 0xbb, 0x7e, 0x73, 0x91, 0xac, 0x87, 0x30, 0x88, 0x96, 0x3d, 0x8d,
	0xe2, 0xa6, 0x50, 0x41, 0x86, 0x3d, 0xc8, 0xaa, 0x3b, 0x78, 0x14, 0x22, 0x0c, 0xde, 0x24, 0x33,
	0xf9, 0xec, 0x86, 0x6d, 0x9b, 0x9a, 0xa4, 0x53, 0xe7, 0xd7, 0x0a, 0x9e, 0xfd, 0x11, 0xd2, 0xa4,
	0xbb, 0x22, 0x08, 0xd3, 0xbd, 0xa5, 0xed, 0xf7, 0x3c, 0x79, 0xdf, 0x26, 0x70, 0xb6, 0x33, 0x28,
	0x1c, 0x7e, 0xc5, 0xda, 0x29, 0x77, 0x60, 0x8d, 0xd1, 0x4e, 0x93, 0x54, 0x28, 0x56, 0xa1, 0x13,
	0xee, 0x46, 0xd5, 0x85, 0x65, 0x16, 0xcd, 0x5e, 0xed, 0x90, 0x29, 0x19, 0xa6, 0x64, 0xbe, 0x78,
	0x56, 0x68, 0x5d, 0x94, 0x98, 0xf9, 0xf9, 0x6d, 0xfb, 0x2d, 0xeb, 0x8c, 0x68, 0x1a, 0x34, 0x36,
	0x18, 0x16, 0x6f, 0x0a, 0xe5, 0x6e, 0xda, 0x2d, 0xdb, 0xe9, 0x2a, 0xab, 0x3a, 0xe5, 0x6e, 0x75,
	0x97, 0x15, 0xfa, 0xc2, 0x7f, 0xbc, 0xd8, 0x8b, 0x52, 0x49, 0xe1, 0xb4, 0xb1, 0xd9, 0xfd, 0x72,
	0xe0, 0xe6, 0x95, 0x96, 0x58, 0x86, 0x87, 0xf3, 0xdf, 0x33, 0x38, 0xfb, 0xe0, 0xe3, 0xad, 0xfd,
	0xbe, 0x1c, 0xbf, 0xb7, 0x68, 0x1d, 0x79, 0x05, 0x91, 0x92, 0x74, 0x92, 0x4e, 0x96, 0xf3, 0xfc,
	0x29, 0x0b, 0x49, 0xd8, 0x5d, 0x12, 0xf6, 0xd9, 0x19, 0x55, 0x17, 0x1b, 0x51, 0xb6, 0xc8, 0x23,
	0x25, 0xc9, 0x35, 0x9c, 0x48, 0x5d, 0x09, 0x55, 0x5b, 0x1a, 0xa5, 0xd3, 0xe5, 0x3c, 0x7f, 0xc9,
	0x8e, 0xd5, 0x63, 0x9f, 0x9a, 0x6b, 0x8f, 0xf0, 0x3b, 0x94, 0xbc, 0x83, 0xa4, 0x30, 0xba, 0x6d,
	0x2c, 0x9d, 0x7a, 0xc9, 0x8b, 0x31, 0x92, 0x8f, 0x3d, 0xc1, 0x07, 0x90, 0xbc, 0x85, 0xd8, 0xe8,
	0x12, 0x2d, 0x9d, 0x79, 0xc3, 0x72, 0x8c, 0x81, 0xeb, 0x12, 0x79, 0xc0, 0xc8, 0x0a, 0x66, 0xb5,
	0xa8, 0x90, 0xc6, 0xc7, 0x8b, 0xbf, 0x4f, 0x7e, 0xfd, 0x7c, 0x1e, 0xa5, 0x13, 0xee, 0x09, 0x72,
	0x09, 0xb1, 0x28, 0x95, 0xb0, 0x34, 0x79, 0x00, 0x1a, 0x10, 0xb2, 0x82, 0xd3, 0x46, 0x58, 0xdb,
	0x69, 0x23, 0xe9, 0xc9, 0x88, 0x5f, 0x7e, 0x3f, 0x4d, 0x36, 0x10, 0xe3, 0x0f, 0x67, 0x04, 0x3d,
	0xf5, 0x7d, 0xaf, 0x8e, 0xf7, 0xfd, 0xcf, 0x61, 0xb3, 0x75, 0xaf, 0x58, 0xd7, 0xce, 0x1c, 0x78,
	0xd0, 0x2d, 0x36, 0x00, 0x7f, 0x5f, 0x92, 0xc7, 0x30, 0xbd, 0xc5, 0x83, 0xbf, 0x0d, 0x8f, 0x78,
	0xbf, 0x24, 0x39, 0xc4, 0xfb, 0x3e, 0x0a, 0x8d, 0x46, 0xc4, 0x0d, 0xa3, 0x97, 0xd1, 0x6a, 0x72,
	0xfe, 0x05, 0x9e, 0xfc, 0x1b, 0xc0, 0x36, 0xba, 0xb6, 0x48, 0xae, 0x20, 0x09, 0x09, 0x87, 0x2b,
	0x37, 0xe2, 0xe0, 0x06, 0xc3, 0xc0, 0x6d, 0x13, 0xbf, 0xf5, 0xeb, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xfa, 0x87, 0x2f, 0x63, 0x6d, 0x03, 0x00, 0x00,
}
