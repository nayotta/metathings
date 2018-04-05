// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_projects_for_user.proto

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

type ListProjectsForUserRequest struct {
	UserId *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *ListProjectsForUserRequest) Reset()                    { *m = ListProjectsForUserRequest{} }
func (m *ListProjectsForUserRequest) String() string            { return proto.CompactTextString(m) }
func (*ListProjectsForUserRequest) ProtoMessage()               {}
func (*ListProjectsForUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor36, []int{0} }

func (m *ListProjectsForUserRequest) GetUserId() *google_protobuf.StringValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

type ListProjectsForUserResponse struct {
	Projects []*Project `protobuf:"bytes,1,rep,name=projects" json:"projects,omitempty"`
}

func (m *ListProjectsForUserResponse) Reset()                    { *m = ListProjectsForUserResponse{} }
func (m *ListProjectsForUserResponse) String() string            { return proto.CompactTextString(m) }
func (*ListProjectsForUserResponse) ProtoMessage()               {}
func (*ListProjectsForUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor36, []int{1} }

func (m *ListProjectsForUserResponse) GetProjects() []*Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

func init() {
	proto.RegisterType((*ListProjectsForUserRequest)(nil), "ai.metathings.service.identity.ListProjectsForUserRequest")
	proto.RegisterType((*ListProjectsForUserResponse)(nil), "ai.metathings.service.identity.ListProjectsForUserResponse")
}

func init() { proto.RegisterFile("list_projects_for_user.proto", fileDescriptor36) }

var fileDescriptor36 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4b, 0xc3, 0x60,
	0x10, 0x86, 0x89, 0x42, 0x95, 0x14, 0x97, 0x4e, 0xa5, 0x96, 0x1a, 0xba, 0xd8, 0xa5, 0x5f, 0xa0,
	0x82, 0x9b, 0x8b, 0x82, 0x20, 0x38, 0x48, 0x44, 0x17, 0x87, 0xf0, 0x25, 0xb9, 0x7e, 0x3d, 0x4d,
	0x72, 0xf1, 0xee, 0xd2, 0xe0, 0xaf, 0x15, 0xfc, 0x25, 0xd2, 0x24, 0xed, 0x24, 0x6e, 0x77, 0xdc,
	0x3d, 0x2f, 0x0f, 0xaf, 0x3f, 0xcd, 0x51, 0x34, 0xae, 0x98, 0xde, 0x21, 0x55, 0x89, 0xd7, 0xc4,
	0x71, 0x2d, 0xc0, 0xa6, 0x62, 0x52, 0x1a, 0xcd, 0x2c, 0x9a, 0x02, 0xd4, 0xea, 0x06, 0x4b, 0x27,
	0x46, 0x80, 0xb7, 0x98, 0x82, 0xc1, 0x0c, 0x4a, 0x45, 0xfd, 0x9a, 0xcc, 0x1c, 0x91, 0xcb, 0x21,
	0x6c, 0xbf, 0x93, 0x7a, 0x1d, 0x36, 0x6c, 0xab, 0x0a, 0x58, 0x3a, 0x7e, 0x72, 0xed, 0x50, 0x37,
	0x75, 0x62, 0x52, 0x2a, 0xc2, 0xa2, 0x41, 0xfd, 0xa0, 0x26, 0x74, 0xb4, 0x6c, 0x8f, 0xcb, 0xad,
	0xcd, 0x31, 0xb3, 0x4a, 0x2c, 0xe1, 0x61, 0xec, 0xb9, 0xb3, 0x5e, 0xa8, 0x5b, 0xe7, 0x6f, 0xfe,
	0xe4, 0x11, 0x45, 0x9f, 0x7a, 0xcb, 0x7b, 0xe2, 0x17, 0x01, 0x8e, 0xe0, 0xb3, 0x06, 0xd1, 0xd1,
	0x8d, 0x7f, 0xb2, 0x53, 0x8e, 0x31, 0x1b, 0x7b, 0x81, 0xb7, 0x18, 0xae, 0xa6, 0xa6, 0xd3, 0x32,
	0x7b, 0x2d, 0xf3, 0xac, 0x8c, 0xa5, 0x7b, 0xb5, 0x79, 0x0d, 0xb7, 0x83, 0x9f, 0xef, 0x8b, 0xa3,
	0xc0, 0x8b, 0x06, 0x3b, 0xe8, 0x21, 0x9b, 0x27, 0xfe, 0xf9, 0x9f, 0xe1, 0x52, 0x51, 0x29, 0x30,
	0xba, 0xf3, 0x4f, 0xf7, 0xed, 0x8c, 0xbd, 0xe0, 0x78, 0x31, 0x5c, 0x5d, 0x9a, 0xff, 0x5b, 0x31,
	0x7d, 0x54, 0x74, 0x00, 0x93, 0x41, 0x6b, 0x72, 0xf5, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x92, 0x3c,
	0xb2, 0xeb, 0x6e, 0x01, 0x00, 0x00,
}
