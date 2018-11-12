// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_projects.proto

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

type ListProjectsRequest struct {
	DomainId             *wrappers.StringValue `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	ParentId             *wrappers.StringValue `protobuf:"bytes,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Enabled              *wrappers.BoolValue   `protobuf:"bytes,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListProjectsRequest) Reset()         { *m = ListProjectsRequest{} }
func (m *ListProjectsRequest) String() string { return proto.CompactTextString(m) }
func (*ListProjectsRequest) ProtoMessage()    {}
func (*ListProjectsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_569f9848119c89e7, []int{0}
}

func (m *ListProjectsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProjectsRequest.Unmarshal(m, b)
}
func (m *ListProjectsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProjectsRequest.Marshal(b, m, deterministic)
}
func (m *ListProjectsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProjectsRequest.Merge(m, src)
}
func (m *ListProjectsRequest) XXX_Size() int {
	return xxx_messageInfo_ListProjectsRequest.Size(m)
}
func (m *ListProjectsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProjectsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListProjectsRequest proto.InternalMessageInfo

func (m *ListProjectsRequest) GetDomainId() *wrappers.StringValue {
	if m != nil {
		return m.DomainId
	}
	return nil
}

func (m *ListProjectsRequest) GetParentId() *wrappers.StringValue {
	if m != nil {
		return m.ParentId
	}
	return nil
}

func (m *ListProjectsRequest) GetEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.Enabled
	}
	return nil
}

func (m *ListProjectsRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

type ListProjectsResponse struct {
	Projects             []*Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListProjectsResponse) Reset()         { *m = ListProjectsResponse{} }
func (m *ListProjectsResponse) String() string { return proto.CompactTextString(m) }
func (*ListProjectsResponse) ProtoMessage()    {}
func (*ListProjectsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_569f9848119c89e7, []int{1}
}

func (m *ListProjectsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProjectsResponse.Unmarshal(m, b)
}
func (m *ListProjectsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProjectsResponse.Marshal(b, m, deterministic)
}
func (m *ListProjectsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProjectsResponse.Merge(m, src)
}
func (m *ListProjectsResponse) XXX_Size() int {
	return xxx_messageInfo_ListProjectsResponse.Size(m)
}
func (m *ListProjectsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProjectsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListProjectsResponse proto.InternalMessageInfo

func (m *ListProjectsResponse) GetProjects() []*Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

func init() {
	proto.RegisterType((*ListProjectsRequest)(nil), "ai.metathings.service.identityd.ListProjectsRequest")
	proto.RegisterType((*ListProjectsResponse)(nil), "ai.metathings.service.identityd.ListProjectsResponse")
}

func init() { proto.RegisterFile("list_projects.proto", fileDescriptor_569f9848119c89e7) }

var fileDescriptor_569f9848119c89e7 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x80, 0xa9, 0x1b, 0xba, 0x65, 0x78, 0xe9, 0x3c, 0x94, 0x21, 0x3a, 0x76, 0xda, 0x65, 0xa9,
	0x4c, 0x11, 0xbc, 0x8a, 0x97, 0x81, 0x07, 0xa9, 0xe0, 0x49, 0x18, 0xe9, 0xf2, 0xec, 0x9e, 0xb6,
	0x79, 0x35, 0x79, 0xdd, 0xf0, 0xcf, 0x8b, 0xac, 0x69, 0x07, 0xe2, 0xc1, 0xdd, 0x5a, 0xf2, 0x7d,
	0x2f, 0x5f, 0x78, 0x62, 0x98, 0xa3, 0xe3, 0x65, 0x69, 0xe9, 0x1d, 0x56, 0xec, 0x64, 0x69, 0x89,
	0x29, 0xbc, 0x54, 0x28, 0x0b, 0x60, 0xc5, 0x6b, 0x34, 0x99, 0x93, 0x0e, 0xec, 0x06, 0x57, 0x20,
	0x51, 0x83, 0x61, 0xe4, 0x2f, 0x3d, 0xba, 0xc8, 0x88, 0xb2, 0x1c, 0xe2, 0x1a, 0x4f, 0xab, 0xb7,
	0x78, 0x6b, 0x55, 0x59, 0x82, 0x6d, 0x06, 0x8c, 0x6e, 0x33, 0xe4, 0x75, 0x95, 0xca, 0x15, 0x15,
	0x71, 0xb1, 0x45, 0xfe, 0xa0, 0x6d, 0x9c, 0xd1, 0xac, 0x3e, 0x9c, 0x6d, 0x54, 0x8e, 0x5a, 0x31,
	0x59, 0x17, 0xef, 0x3f, 0x1b, 0xef, 0xb4, 0x09, 0xf1, 0xbf, 0x93, 0xef, 0x40, 0x0c, 0x1f, 0xd1,
	0xf1, 0x53, 0x93, 0x97, 0xc0, 0x67, 0x05, 0x8e, 0xc3, 0x3b, 0xd1, 0xd7, 0x54, 0x28, 0x34, 0x4b,
	0xd4, 0x51, 0x30, 0x0e, 0xa6, 0x83, 0xf9, 0xb9, 0xf4, 0x49, 0xb2, 0x4d, 0x92, 0xcf, 0x6c, 0xd1,
	0x64, 0x2f, 0x2a, 0xaf, 0x20, 0xe9, 0x79, 0x7c, 0xa1, 0x77, 0x6a, 0xa9, 0x2c, 0x18, 0xde, 0xa9,
	0x47, 0x87, 0xa8, 0x1e, 0x5f, 0xe8, 0xf0, 0x46, 0x9c, 0x80, 0x51, 0x69, 0x0e, 0x3a, 0xea, 0xd4,
	0xe2, 0xe8, 0x8f, 0x78, 0x4f, 0x94, 0x7b, 0xad, 0x45, 0xc3, 0x2b, 0xd1, 0x35, 0xaa, 0x80, 0xa8,
	0x7b, 0xc0, 0x5d, 0x35, 0x99, 0xf4, 0xd1, 0x2d, 0x7d, 0xef, 0xe4, 0x55, 0x9c, 0xfd, 0x7e, 0xbf,
	0x2b, 0xc9, 0x38, 0x08, 0x1f, 0x44, 0xaf, 0x5d, 0x59, 0x14, 0x8c, 0x3b, 0xd3, 0xc1, 0x7c, 0x2a,
	0xff, 0xd9, 0x99, 0x6c, 0x86, 0x24, 0x7b, 0x33, 0x3d, 0xae, 0x23, 0xae, 0x7f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xfd, 0x50, 0xe4, 0xfd, 0x04, 0x02, 0x00, 0x00,
}
