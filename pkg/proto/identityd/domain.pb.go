// Code generated by protoc-gen-go. DO NOT EDIT.
// source: domain.proto

package identityd

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Domain struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Enabled              bool     `protobuf:"varint,4,opt,name=enabled" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Domain) Reset()         { *m = Domain{} }
func (m *Domain) String() string { return proto.CompactTextString(m) }
func (*Domain) ProtoMessage()    {}
func (*Domain) Descriptor() ([]byte, []int) {
	return fileDescriptor_domain_1c148c38e1a19247, []int{0}
}
func (m *Domain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Domain.Unmarshal(m, b)
}
func (m *Domain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Domain.Marshal(b, m, deterministic)
}
func (dst *Domain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Domain.Merge(dst, src)
}
func (m *Domain) XXX_Size() int {
	return xxx_messageInfo_Domain.Size(m)
}
func (m *Domain) XXX_DiscardUnknown() {
	xxx_messageInfo_Domain.DiscardUnknown(m)
}

var xxx_messageInfo_Domain proto.InternalMessageInfo

func (m *Domain) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Domain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Domain) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Domain) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func init() {
	proto.RegisterType((*Domain)(nil), "ai.metathings.service.identityd.Domain")
}

func init() { proto.RegisterFile("domain.proto", fileDescriptor_domain_1c148c38e1a19247) }

var fileDescriptor_domain_1c148c38e1a19247 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0x41, 0x0a, 0xc2, 0x30,
	0x10, 0x45, 0x69, 0x2d, 0x55, 0x47, 0x71, 0x31, 0xab, 0xec, 0x2c, 0xae, 0xba, 0xca, 0xc6, 0x2b,
	0x78, 0x82, 0xde, 0x20, 0xed, 0x0c, 0x76, 0xc0, 0x4c, 0x4a, 0x32, 0x08, 0xde, 0x5e, 0x08, 0x08,
	0xee, 0xfe, 0x7f, 0x9f, 0x0f, 0x0f, 0xce, 0x94, 0x62, 0x10, 0xf5, 0x5b, 0x4e, 0x96, 0xf0, 0x1a,
	0xc4, 0x47, 0xb6, 0x60, 0xab, 0xe8, 0xb3, 0xf8, 0xc2, 0xf9, 0x2d, 0x0b, 0x7b, 0x21, 0x56, 0x13,
	0xfb, 0xd0, 0x6d, 0x85, 0xfe, 0x51, 0x0f, 0x78, 0x81, 0x56, 0xc8, 0x35, 0x43, 0x33, 0x1e, 0xa7,
	0x56, 0x08, 0x11, 0x3a, 0x0d, 0x91, 0x5d, 0x5b, 0x49, 0xcd, 0x38, 0xc0, 0x89, 0xb8, 0x2c, 0x59,
	0x36, 0x93, 0xa4, 0x6e, 0x57, 0xa7, 0x7f, 0x84, 0x0e, 0xf6, 0xac, 0x61, 0x7e, 0x31, 0xb9, 0x6e,
	0x68, 0xc6, 0xc3, 0xf4, 0xab, 0x73, 0x5f, 0x8d, 0xee, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x67,
	0xce, 0x1d, 0x73, 0xa1, 0x00, 0x00, 0x00,
}