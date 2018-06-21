// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_projects_for_user.proto

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

type ListProjectsForUserRequest struct {
	UserId               *wrappers.StringValue `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListProjectsForUserRequest) Reset()         { *m = ListProjectsForUserRequest{} }
func (m *ListProjectsForUserRequest) String() string { return proto.CompactTextString(m) }
func (*ListProjectsForUserRequest) ProtoMessage()    {}
func (*ListProjectsForUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_projects_for_user_f37dc91d413774eb, []int{0}
}
func (m *ListProjectsForUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProjectsForUserRequest.Unmarshal(m, b)
}
func (m *ListProjectsForUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProjectsForUserRequest.Marshal(b, m, deterministic)
}
func (dst *ListProjectsForUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProjectsForUserRequest.Merge(dst, src)
}
func (m *ListProjectsForUserRequest) XXX_Size() int {
	return xxx_messageInfo_ListProjectsForUserRequest.Size(m)
}
func (m *ListProjectsForUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProjectsForUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListProjectsForUserRequest proto.InternalMessageInfo

func (m *ListProjectsForUserRequest) GetUserId() *wrappers.StringValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

type ListProjectsForUserResponse struct {
	Projects             []*Project `protobuf:"bytes,1,rep,name=projects" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListProjectsForUserResponse) Reset()         { *m = ListProjectsForUserResponse{} }
func (m *ListProjectsForUserResponse) String() string { return proto.CompactTextString(m) }
func (*ListProjectsForUserResponse) ProtoMessage()    {}
func (*ListProjectsForUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_projects_for_user_f37dc91d413774eb, []int{1}
}
func (m *ListProjectsForUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProjectsForUserResponse.Unmarshal(m, b)
}
func (m *ListProjectsForUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProjectsForUserResponse.Marshal(b, m, deterministic)
}
func (dst *ListProjectsForUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProjectsForUserResponse.Merge(dst, src)
}
func (m *ListProjectsForUserResponse) XXX_Size() int {
	return xxx_messageInfo_ListProjectsForUserResponse.Size(m)
}
func (m *ListProjectsForUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProjectsForUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListProjectsForUserResponse proto.InternalMessageInfo

func (m *ListProjectsForUserResponse) GetProjects() []*Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

func init() {
	proto.RegisterType((*ListProjectsForUserRequest)(nil), "ai.metathings.service.identityd.ListProjectsForUserRequest")
	proto.RegisterType((*ListProjectsForUserResponse)(nil), "ai.metathings.service.identityd.ListProjectsForUserResponse")
}

func init() {
	proto.RegisterFile("list_projects_for_user.proto", fileDescriptor_list_projects_for_user_f37dc91d413774eb)
}

var fileDescriptor_list_projects_for_user_f37dc91d413774eb = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4b, 0xc3, 0x50,
	0x10, 0xc7, 0x89, 0x42, 0x94, 0x14, 0x97, 0x4c, 0x25, 0x16, 0x1b, 0x3a, 0x65, 0xe9, 0x0b, 0x54,
	0x70, 0x73, 0x11, 0x11, 0x04, 0x07, 0x89, 0xe8, 0xe2, 0x10, 0x5e, 0x92, 0xeb, 0xeb, 0x69, 0x92,
	0x8b, 0xef, 0x2e, 0x0d, 0x7e, 0x5a, 0xc1, 0x4f, 0x22, 0x4d, 0xd2, 0x4e, 0x42, 0xb7, 0x3b, 0xee,
	0xfe, 0x3f, 0x7e, 0xfc, 0xbd, 0x59, 0x89, 0x2c, 0x69, 0x63, 0xe9, 0x03, 0x72, 0xe1, 0x74, 0x4d,
	0x36, 0x6d, 0x19, 0xac, 0x6a, 0x2c, 0x09, 0xf9, 0x73, 0x8d, 0xaa, 0x02, 0xd1, 0xb2, 0xc1, 0xda,
	0xb0, 0x62, 0xb0, 0x5b, 0xcc, 0x41, 0x61, 0x01, 0xb5, 0xa0, 0x7c, 0x17, 0xc1, 0x95, 0x21, 0x32,
	0x25, 0xc4, 0xfd, 0x7b, 0xd6, 0xae, 0xe3, 0xce, 0xea, 0xa6, 0x01, 0xcb, 0x03, 0x20, 0xb8, 0x31,
	0x28, 0x9b, 0x36, 0x53, 0x39, 0x55, 0x71, 0xd5, 0xa1, 0x7c, 0x52, 0x17, 0x1b, 0x5a, 0xf6, 0xc7,
	0xe5, 0x56, 0x97, 0x58, 0x68, 0x21, 0xcb, 0xf1, 0x61, 0x1c, 0x73, 0x17, 0xa3, 0xd1, 0xb0, 0x2e,
	0xde, 0xbd, 0xe0, 0x09, 0x59, 0x9e, 0x47, 0xcd, 0x07, 0xb2, 0xaf, 0x0c, 0x36, 0x81, 0xaf, 0x16,
	0x58, 0xfc, 0x5b, 0xef, 0x6c, 0xe7, 0x9c, 0x62, 0x31, 0x75, 0x42, 0x27, 0x9a, 0xac, 0x66, 0x6a,
	0xd0, 0x52, 0x7b, 0x2d, 0xf5, 0x22, 0x16, 0x6b, 0xf3, 0xa6, 0xcb, 0x16, 0xee, 0xdc, 0xdf, 0x9f,
	0xf9, 0x49, 0xe8, 0x24, 0xee, 0x2e, 0xf4, 0x58, 0x2c, 0x72, 0xef, 0xf2, 0x5f, 0x38, 0x37, 0x54,
	0x33, 0xf8, 0xf7, 0xde, 0xf9, 0xbe, 0x9e, 0xa9, 0x13, 0x9e, 0x46, 0x93, 0x55, 0xa4, 0x8e, 0xd4,
	0xa2, 0x46, 0x56, 0x72, 0x48, 0x66, 0x6e, 0xaf, 0x72, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0x87,
	0x24, 0xe1, 0x1a, 0x70, 0x01, 0x00, 0x00,
}