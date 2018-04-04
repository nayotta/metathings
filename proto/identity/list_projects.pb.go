// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_projects.proto

package identity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ListProjectsRequest struct {
	DomainId *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	ParentId *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=parent_id,json=parentId" json:"parent_id,omitempty"`
	Enabled  *google_protobuf.BoolValue   `protobuf:"bytes,3,opt,name=enabled" json:"enabled,omitempty"`
	Name     *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
}

func (m *ListProjectsRequest) Reset()                    { *m = ListProjectsRequest{} }
func (m *ListProjectsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListProjectsRequest) ProtoMessage()               {}
func (*ListProjectsRequest) Descriptor() ([]byte, []int) { return fileDescriptor34, []int{0} }

func (m *ListProjectsRequest) GetDomainId() *google_protobuf.StringValue {
	if m != nil {
		return m.DomainId
	}
	return nil
}

func (m *ListProjectsRequest) GetParentId() *google_protobuf.StringValue {
	if m != nil {
		return m.ParentId
	}
	return nil
}

func (m *ListProjectsRequest) GetEnabled() *google_protobuf.BoolValue {
	if m != nil {
		return m.Enabled
	}
	return nil
}

func (m *ListProjectsRequest) GetName() *google_protobuf.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

type ListProjectsResponse struct {
	Projects []*Project `protobuf:"bytes,1,rep,name=projects" json:"projects,omitempty"`
}

func (m *ListProjectsResponse) Reset()                    { *m = ListProjectsResponse{} }
func (m *ListProjectsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListProjectsResponse) ProtoMessage()               {}
func (*ListProjectsResponse) Descriptor() ([]byte, []int) { return fileDescriptor34, []int{1} }

func (m *ListProjectsResponse) GetProjects() []*Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

func init() {
	proto.RegisterType((*ListProjectsRequest)(nil), "ai.metathings.service.identity.ListProjectsRequest")
	proto.RegisterType((*ListProjectsResponse)(nil), "ai.metathings.service.identity.ListProjectsResponse")
}

func init() { proto.RegisterFile("list_projects.proto", fileDescriptor34) }

var fileDescriptor34 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4b, 0xfb, 0x30,
	0x14, 0x80, 0xe9, 0x6f, 0xe3, 0xe7, 0x96, 0xe1, 0xa5, 0xf3, 0x50, 0x86, 0x8c, 0xb1, 0x8b, 0xbb,
	0x2c, 0x95, 0x29, 0x82, 0x57, 0x3d, 0x0d, 0x3c, 0x48, 0x05, 0x2f, 0x1e, 0x46, 0xba, 0x3c, 0xbb,
	0xa7, 0x6d, 0x5e, 0x4d, 0x5e, 0x37, 0xfc, 0xe7, 0x45, 0xd6, 0xb4, 0x03, 0x11, 0x64, 0xb7, 0x84,
	0x7c, 0x5f, 0xf2, 0x85, 0x27, 0x86, 0x39, 0x3a, 0x5e, 0x95, 0x96, 0xde, 0x60, 0xcd, 0x4e, 0x96,
	0x96, 0x98, 0xc2, 0xb1, 0x42, 0x59, 0x00, 0x2b, 0xde, 0xa0, 0xc9, 0x9c, 0x74, 0x60, 0xb7, 0xb8,
	0x06, 0x89, 0x1a, 0x0c, 0x23, 0x7f, 0x8e, 0xc6, 0x19, 0x51, 0x96, 0x43, 0x5c, 0xd3, 0x69, 0xf5,
	0x1a, 0xef, 0xac, 0x2a, 0x4b, 0xb0, 0x8d, 0x3f, 0xba, 0xc9, 0x90, 0x37, 0x55, 0x2a, 0xd7, 0x54,
	0xc4, 0xc5, 0x0e, 0xf9, 0x9d, 0x76, 0x71, 0x46, 0xf3, 0xfa, 0x70, 0xbe, 0x55, 0x39, 0x6a, 0xc5,
	0x64, 0x5d, 0x7c, 0x58, 0x36, 0xde, 0x69, 0xd3, 0xe1, 0xb7, 0xd3, 0xaf, 0x40, 0x0c, 0x1f, 0xd0,
	0xf1, 0x63, 0x53, 0x97, 0xc0, 0x47, 0x05, 0x8e, 0xc3, 0x5b, 0xd1, 0xd7, 0x54, 0x28, 0x34, 0x2b,
	0xd4, 0x51, 0x30, 0x09, 0x66, 0x83, 0xc5, 0xb9, 0xf4, 0x49, 0xb2, 0x4d, 0x92, 0x4f, 0x6c, 0xd1,
	0x64, 0xcf, 0x2a, 0xaf, 0x20, 0xe9, 0x79, 0x7c, 0xa9, 0xf7, 0x6a, 0xa9, 0x2c, 0x18, 0xde, 0xab,
	0xff, 0x8e, 0x51, 0x3d, 0xbe, 0xd4, 0xe1, 0xb5, 0x38, 0x01, 0xa3, 0xd2, 0x1c, 0x74, 0xd4, 0xa9,
	0xc5, 0xd1, 0x2f, 0xf1, 0x8e, 0x28, 0xf7, 0x5a, 0x8b, 0x86, 0x97, 0xa2, 0x6b, 0x54, 0x01, 0x51,
	0xf7, 0x88, 0xb7, 0x6a, 0x32, 0xe9, 0xa3, 0x5b, 0xf9, 0xde, 0xe9, 0x8b, 0x38, 0xfb, 0xf9, 0x7f,
	0x57, 0x92, 0x71, 0x10, 0xde, 0x8b, 0x5e, 0x3b, 0xb1, 0x28, 0x98, 0x74, 0x66, 0x83, 0xc5, 0x85,
	0xfc, 0x7b, 0x64, 0xb2, 0xb9, 0x23, 0x39, 0x88, 0xe9, 0xff, 0xba, 0xe1, 0xea, 0x3b, 0x00, 0x00,
	0xff, 0xff, 0x60, 0x3a, 0x44, 0x08, 0x02, 0x02, 0x00, 0x00,
}
