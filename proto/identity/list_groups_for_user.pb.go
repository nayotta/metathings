// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_groups_for_user.proto

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

type ListGroupsForUserRequest struct {
	UserId *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *ListGroupsForUserRequest) Reset()                    { *m = ListGroupsForUserRequest{} }
func (m *ListGroupsForUserRequest) String() string            { return proto.CompactTextString(m) }
func (*ListGroupsForUserRequest) ProtoMessage()               {}
func (*ListGroupsForUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor34, []int{0} }

func (m *ListGroupsForUserRequest) GetUserId() *google_protobuf.StringValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

type ListGroupsForUserResponse struct {
	Groups []*Group `protobuf:"bytes,1,rep,name=groups" json:"groups,omitempty"`
}

func (m *ListGroupsForUserResponse) Reset()                    { *m = ListGroupsForUserResponse{} }
func (m *ListGroupsForUserResponse) String() string            { return proto.CompactTextString(m) }
func (*ListGroupsForUserResponse) ProtoMessage()               {}
func (*ListGroupsForUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor34, []int{1} }

func (m *ListGroupsForUserResponse) GetGroups() []*Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

func init() {
	proto.RegisterType((*ListGroupsForUserRequest)(nil), "ai.metathings.service.identity.ListGroupsForUserRequest")
	proto.RegisterType((*ListGroupsForUserResponse)(nil), "ai.metathings.service.identity.ListGroupsForUserResponse")
}

func init() { proto.RegisterFile("list_groups_for_user.proto", fileDescriptor34) }

var fileDescriptor34 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0xc9, 0xf7, 0x41, 0x84, 0xe4, 0x96, 0x53, 0x0c, 0x52, 0x43, 0x41, 0xe8, 0xa5, 0x1b,
	0xa8, 0xe0, 0xad, 0x17, 0x0f, 0x8a, 0xe0, 0x29, 0xa2, 0xa0, 0x97, 0xb0, 0x69, 0xa6, 0xdb, 0xc1,
	0x24, 0x13, 0x77, 0x66, 0x1b, 0xfc, 0xb5, 0x82, 0xbf, 0x44, 0xba, 0xa9, 0x9e, 0xc4, 0xdb, 0x0c,
	0x33, 0xcf, 0xcb, 0xc3, 0x1b, 0x65, 0x2d, 0xb2, 0x54, 0xc6, 0x92, 0x1b, 0xb8, 0xda, 0x92, 0xad,
	0x1c, 0x83, 0x55, 0x83, 0x25, 0xa1, 0x64, 0xa6, 0x51, 0x75, 0x20, 0x5a, 0x76, 0xd8, 0x1b, 0x56,
	0x0c, 0x76, 0x8f, 0x1b, 0x50, 0xd8, 0x40, 0x2f, 0x28, 0xef, 0xd9, 0xcc, 0x10, 0x99, 0x16, 0x0a,
	0xff, 0x5d, 0xbb, 0x6d, 0x31, 0x5a, 0x3d, 0x0c, 0x60, 0x79, 0xe2, 0xb3, 0x2b, 0x83, 0xb2, 0x73,
	0xb5, 0xda, 0x50, 0x57, 0x74, 0x23, 0xca, 0x2b, 0x8d, 0x85, 0xa1, 0xa5, 0x3f, 0x2e, 0xf7, 0xba,
	0xc5, 0x46, 0x0b, 0x59, 0x2e, 0x7e, 0xc6, 0x23, 0x17, 0x7b, 0x9d, 0x69, 0x99, 0x3f, 0x47, 0xe9,
	0x3d, 0xb2, 0xdc, 0x7a, 0xc3, 0x1b, 0xb2, 0x8f, 0x0c, 0xb6, 0x84, 0x37, 0x07, 0x2c, 0xc9, 0x3a,
	0x3a, 0x39, 0xe8, 0x56, 0xd8, 0xa4, 0x41, 0x1e, 0x2c, 0xe2, 0xd5, 0x99, 0x9a, 0x94, 0xd4, 0xb7,
	0x92, 0x7a, 0x10, 0x8b, 0xbd, 0x79, 0xd2, 0xad, 0x83, 0xeb, 0xf0, 0xf3, 0xe3, 0xfc, 0x5f, 0x1e,
	0x94, 0xe1, 0x01, 0xba, 0x6b, 0xe6, 0x2f, 0xd1, 0xe9, 0x2f, 0xd1, 0x3c, 0x50, 0xcf, 0x90, 0xac,
	0xa3, 0x70, 0x6a, 0x25, 0x0d, 0xf2, 0xff, 0x8b, 0x78, 0x75, 0xa1, 0xfe, 0x6e, 0x43, 0xf9, 0x98,
	0xf2, 0x08, 0xd5, 0xa1, 0x37, 0xb8, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x24, 0xcc, 0x0c, 0x56,
	0x60, 0x01, 0x00, 0x00,
}
