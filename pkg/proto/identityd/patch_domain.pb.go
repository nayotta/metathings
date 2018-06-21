// Code generated by protoc-gen-go. DO NOT EDIT.
// source: patch_domain.proto

package identityd

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type PatchDomainRequest struct {
	DomainId             *wrappers.StringValue `protobuf:"bytes,1,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description          *wrappers.StringValue `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Enabled              *wrappers.BoolValue   `protobuf:"bytes,4,opt,name=enabled" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PatchDomainRequest) Reset()         { *m = PatchDomainRequest{} }
func (m *PatchDomainRequest) String() string { return proto.CompactTextString(m) }
func (*PatchDomainRequest) ProtoMessage()    {}
func (*PatchDomainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_patch_domain_398a490be480359b, []int{0}
}
func (m *PatchDomainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchDomainRequest.Unmarshal(m, b)
}
func (m *PatchDomainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchDomainRequest.Marshal(b, m, deterministic)
}
func (dst *PatchDomainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchDomainRequest.Merge(dst, src)
}
func (m *PatchDomainRequest) XXX_Size() int {
	return xxx_messageInfo_PatchDomainRequest.Size(m)
}
func (m *PatchDomainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchDomainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PatchDomainRequest proto.InternalMessageInfo

func (m *PatchDomainRequest) GetDomainId() *wrappers.StringValue {
	if m != nil {
		return m.DomainId
	}
	return nil
}

func (m *PatchDomainRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *PatchDomainRequest) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *PatchDomainRequest) GetEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.Enabled
	}
	return nil
}

type PatchDomainResponse struct {
	Domain               *Domain  `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PatchDomainResponse) Reset()         { *m = PatchDomainResponse{} }
func (m *PatchDomainResponse) String() string { return proto.CompactTextString(m) }
func (*PatchDomainResponse) ProtoMessage()    {}
func (*PatchDomainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_patch_domain_398a490be480359b, []int{1}
}
func (m *PatchDomainResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchDomainResponse.Unmarshal(m, b)
}
func (m *PatchDomainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchDomainResponse.Marshal(b, m, deterministic)
}
func (dst *PatchDomainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchDomainResponse.Merge(dst, src)
}
func (m *PatchDomainResponse) XXX_Size() int {
	return xxx_messageInfo_PatchDomainResponse.Size(m)
}
func (m *PatchDomainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchDomainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PatchDomainResponse proto.InternalMessageInfo

func (m *PatchDomainResponse) GetDomain() *Domain {
	if m != nil {
		return m.Domain
	}
	return nil
}

func init() {
	proto.RegisterType((*PatchDomainRequest)(nil), "ai.metathings.service.identityd.PatchDomainRequest")
	proto.RegisterType((*PatchDomainResponse)(nil), "ai.metathings.service.identityd.PatchDomainResponse")
}

func init() { proto.RegisterFile("patch_domain.proto", fileDescriptor_patch_domain_398a490be480359b) }

var fileDescriptor_patch_domain_398a490be480359b = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0x33, 0x31,
	0x10, 0x86, 0xd9, 0x7e, 0xa5, 0x9f, 0xa6, 0x9e, 0xe2, 0x65, 0x29, 0x62, 0x4b, 0x2f, 0x7a, 0x69,
	0x56, 0x54, 0x3c, 0x2a, 0x16, 0x2f, 0xde, 0xa4, 0x42, 0xaf, 0x25, 0xbb, 0x19, 0xb7, 0x83, 0xbb,
	0x99, 0x35, 0x99, 0x6d, 0xf1, 0xd7, 0x0a, 0xfe, 0x11, 0xc5, 0xcd, 0x56, 0x2a, 0x1e, 0xea, 0x2d,
	0x24, 0xef, 0x33, 0x79, 0xde, 0x44, 0xc8, 0x4a, 0x73, 0xb6, 0x5c, 0x18, 0x2a, 0x35, 0x5a, 0x55,
	0x39, 0x62, 0x92, 0x43, 0x8d, 0xaa, 0x04, 0xd6, 0xbc, 0x44, 0x9b, 0x7b, 0xe5, 0xc1, 0xad, 0x30,
	0x03, 0x85, 0x06, 0x2c, 0x23, 0xbf, 0x9a, 0xc1, 0x71, 0x4e, 0x94, 0x17, 0x90, 0x34, 0xf1, 0xb4,
	0x7e, 0x4a, 0xd6, 0x4e, 0x57, 0x15, 0x38, 0x1f, 0x06, 0x0c, 0xae, 0x72, 0xe4, 0x65, 0x9d, 0xaa,
	0x8c, 0xca, 0xa4, 0x5c, 0x23, 0x3f, 0xd3, 0x3a, 0xc9, 0x69, 0xd2, 0x1c, 0x4e, 0x56, 0xba, 0x40,
	0xa3, 0x99, 0x9c, 0x4f, 0xbe, 0x97, 0x2d, 0x77, 0xb0, 0xad, 0x31, 0xfe, 0x88, 0x84, 0x7c, 0xf8,
	0xb2, 0xbb, 0x6b, 0x76, 0x67, 0xf0, 0x52, 0x83, 0x67, 0x79, 0x2b, 0xf6, 0x43, 0x6c, 0x81, 0x26,
	0x8e, 0x46, 0xd1, 0x69, 0xff, 0xfc, 0x48, 0x05, 0x21, 0xb5, 0x11, 0x52, 0x8f, 0xec, 0xd0, 0xe6,
	0x73, 0x5d, 0xd4, 0x30, 0xed, 0xbd, 0xbf, 0x0d, 0x3b, 0xa3, 0x68, 0xb6, 0x17, 0xb0, 0x7b, 0x23,
	0xcf, 0x44, 0xd7, 0xea, 0x12, 0xe2, 0xce, 0x6e, 0x7a, 0xd6, 0x24, 0xe5, 0xb5, 0xe8, 0x1b, 0xf0,
	0x99, 0xc3, 0x8a, 0x91, 0x6c, 0xfc, 0xef, 0x0f, 0xe0, 0x36, 0x20, 0x2f, 0xc5, 0x7f, 0xb0, 0x3a,
	0x2d, 0xc0, 0xc4, 0xdd, 0x86, 0x1d, 0xfc, 0x62, 0xa7, 0x44, 0x45, 0x20, 0x37, 0xd1, 0xf1, 0x5c,
	0x1c, 0xfe, 0x78, 0x00, 0x5f, 0x91, 0xf5, 0x20, 0x6f, 0x44, 0x2f, 0x54, 0x69, 0xeb, 0x9f, 0xa8,
	0x1d, 0x1f, 0xa6, 0xda, 0x01, 0x2d, 0x96, 0xf6, 0x9a, 0x4b, 0x2f, 0x3e, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x8f, 0x20, 0x3d, 0xd2, 0xfd, 0x01, 0x00, 0x00,
}