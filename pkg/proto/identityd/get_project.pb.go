// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_project.proto

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

type GetProjectRequest struct {
	ProjectId            *wrappers.StringValue `protobuf:"bytes,1,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetProjectRequest) Reset()         { *m = GetProjectRequest{} }
func (m *GetProjectRequest) String() string { return proto.CompactTextString(m) }
func (*GetProjectRequest) ProtoMessage()    {}
func (*GetProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_get_project_62d969eb28fe4ae9, []int{0}
}
func (m *GetProjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProjectRequest.Unmarshal(m, b)
}
func (m *GetProjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProjectRequest.Marshal(b, m, deterministic)
}
func (dst *GetProjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProjectRequest.Merge(dst, src)
}
func (m *GetProjectRequest) XXX_Size() int {
	return xxx_messageInfo_GetProjectRequest.Size(m)
}
func (m *GetProjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProjectRequest proto.InternalMessageInfo

func (m *GetProjectRequest) GetProjectId() *wrappers.StringValue {
	if m != nil {
		return m.ProjectId
	}
	return nil
}

type GetProjectResponse struct {
	Project              *Project `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProjectResponse) Reset()         { *m = GetProjectResponse{} }
func (m *GetProjectResponse) String() string { return proto.CompactTextString(m) }
func (*GetProjectResponse) ProtoMessage()    {}
func (*GetProjectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_get_project_62d969eb28fe4ae9, []int{1}
}
func (m *GetProjectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProjectResponse.Unmarshal(m, b)
}
func (m *GetProjectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProjectResponse.Marshal(b, m, deterministic)
}
func (dst *GetProjectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProjectResponse.Merge(dst, src)
}
func (m *GetProjectResponse) XXX_Size() int {
	return xxx_messageInfo_GetProjectResponse.Size(m)
}
func (m *GetProjectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProjectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProjectResponse proto.InternalMessageInfo

func (m *GetProjectResponse) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

func init() {
	proto.RegisterType((*GetProjectRequest)(nil), "ai.metathings.service.identityd.GetProjectRequest")
	proto.RegisterType((*GetProjectResponse)(nil), "ai.metathings.service.identityd.GetProjectResponse")
}

func init() { proto.RegisterFile("get_project.proto", fileDescriptor_get_project_62d969eb28fe4ae9) }

var fileDescriptor_get_project_62d969eb28fe4ae9 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0xc9, 0x77, 0xe8, 0x87, 0x2b, 0x1e, 0x9a, 0x93, 0x14, 0xb1, 0xa5, 0xa7, 0x5e, 0x3a,
	0x0b, 0x0a, 0xfe, 0x80, 0x7a, 0x10, 0x6f, 0x52, 0x41, 0x7a, 0x2b, 0x9b, 0xec, 0xb8, 0x1d, 0x4d,
	0x32, 0xeb, 0xee, 0xa4, 0xc1, 0x5f, 0x2b, 0xf8, 0x4b, 0x84, 0x64, 0x23, 0x7a, 0xf2, 0x36, 0xcc,
	0xbc, 0xef, 0xc3, 0xc3, 0xa8, 0xa9, 0x43, 0xd9, 0xfb, 0xc0, 0x2f, 0x58, 0x0a, 0xf8, 0xc0, 0xc2,
	0xf9, 0xdc, 0x10, 0xd4, 0x28, 0x46, 0x0e, 0xd4, 0xb8, 0x08, 0x11, 0xc3, 0x91, 0x4a, 0x04, 0xb2,
	0xd8, 0x08, 0xc9, 0xbb, 0x9d, 0x5d, 0x3a, 0x66, 0x57, 0xa1, 0xee, 0xe3, 0x45, 0xfb, 0xac, 0xbb,
	0x60, 0xbc, 0xc7, 0x10, 0x07, 0xc0, 0xec, 0xc6, 0x91, 0x1c, 0xda, 0x02, 0x4a, 0xae, 0x75, 0xdd,
	0x91, 0xbc, 0x72, 0xa7, 0x1d, 0xaf, 0xfb, 0xe3, 0xfa, 0x68, 0x2a, 0xb2, 0x46, 0x38, 0x44, 0xfd,
	0x3d, 0xa6, 0xde, 0xd9, 0x2f, 0x8f, 0xe5, 0x4e, 0x4d, 0xef, 0x50, 0x1e, 0x86, 0xdd, 0x16, 0xdf,
	0x5a, 0x8c, 0x92, 0xdf, 0x2a, 0x95, 0x52, 0x7b, 0xb2, 0xe7, 0xd9, 0x22, 0x5b, 0x9d, 0x5e, 0x5d,
	0xc0, 0x20, 0x04, 0xa3, 0x10, 0x3c, 0x4a, 0xa0, 0xc6, 0x3d, 0x99, 0xaa, 0xc5, 0xcd, 0xe4, 0xf3,
	0x63, 0xfe, 0x6f, 0x91, 0x6d, 0x4f, 0x52, 0xef, 0xde, 0x2e, 0x77, 0x2a, 0xff, 0x49, 0x8e, 0x9e,
	0x9b, 0x88, 0xf9, 0x46, 0xfd, 0x4f, 0x91, 0xc4, 0x5d, 0xc1, 0x1f, 0x9f, 0x80, 0x11, 0x31, 0x16,
	0x8b, 0x49, 0xaf, 0x70, 0xfd, 0x15, 0x00, 0x00, 0xff, 0xff, 0x94, 0x86, 0x5c, 0x23, 0x57, 0x01,
	0x00, 0x00,
}