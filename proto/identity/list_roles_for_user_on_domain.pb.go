// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_roles_for_user_on_domain.proto

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

type ListRolesForUserOnDomainRequest struct {
	DomainId *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	UserId   *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *ListRolesForUserOnDomainRequest) Reset()         { *m = ListRolesForUserOnDomainRequest{} }
func (m *ListRolesForUserOnDomainRequest) String() string { return proto.CompactTextString(m) }
func (*ListRolesForUserOnDomainRequest) ProtoMessage()    {}
func (*ListRolesForUserOnDomainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor42, []int{0}
}

func (m *ListRolesForUserOnDomainRequest) GetDomainId() *google_protobuf.StringValue {
	if m != nil {
		return m.DomainId
	}
	return nil
}

func (m *ListRolesForUserOnDomainRequest) GetUserId() *google_protobuf.StringValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

type ListRolesForUserOnDomainResponse struct {
	Roles []*Role `protobuf:"bytes,1,rep,name=roles" json:"roles,omitempty"`
}

func (m *ListRolesForUserOnDomainResponse) Reset()         { *m = ListRolesForUserOnDomainResponse{} }
func (m *ListRolesForUserOnDomainResponse) String() string { return proto.CompactTextString(m) }
func (*ListRolesForUserOnDomainResponse) ProtoMessage()    {}
func (*ListRolesForUserOnDomainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor42, []int{1}
}

func (m *ListRolesForUserOnDomainResponse) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRolesForUserOnDomainRequest)(nil), "ai.metathings.service.identity.ListRolesForUserOnDomainRequest")
	proto.RegisterType((*ListRolesForUserOnDomainResponse)(nil), "ai.metathings.service.identity.ListRolesForUserOnDomainResponse")
}

func init() { proto.RegisterFile("list_roles_for_user_on_domain.proto", fileDescriptor42) }

var fileDescriptor42 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0x49, 0x3f, 0xbe, 0xaa, 0xdb, 0x5b, 0x4e, 0xa5, 0x48, 0x1b, 0xaa, 0x87, 0x5e, 0xba,
	0x81, 0x0a, 0x1e, 0x04, 0x0f, 0x8a, 0x08, 0x05, 0x41, 0x88, 0xe8, 0xd1, 0xb0, 0xed, 0x4e, 0xd3,
	0xc1, 0x64, 0x27, 0xee, 0x4c, 0x1a, 0xfc, 0x3d, 0xfe, 0x30, 0xc1, 0x5f, 0x22, 0x49, 0xd4, 0x9b,
	0x82, 0xb7, 0x17, 0x66, 0x9e, 0x97, 0x87, 0x57, 0x1d, 0xe5, 0xc8, 0x92, 0x7a, 0xca, 0x81, 0xd3,
	0x0d, 0xf9, 0xb4, 0x62, 0xf0, 0x29, 0xb9, 0xd4, 0x52, 0x61, 0xd0, 0xe9, 0xd2, 0x93, 0x50, 0x38,
	0x36, 0xa8, 0x0b, 0x10, 0x23, 0x5b, 0x74, 0x19, 0x6b, 0x06, 0xbf, 0xc3, 0x35, 0x68, 0xb4, 0xe0,
	0x04, 0xe5, 0x65, 0x34, 0xce, 0x88, 0xb2, 0x1c, 0xe2, 0xf6, 0x7b, 0x55, 0x6d, 0xe2, 0xda, 0x9b,
	0xb2, 0x04, 0xcf, 0x1d, 0x3f, 0x3a, 0xcd, 0x50, 0xb6, 0xd5, 0x4a, 0xaf, 0xa9, 0x88, 0x8b, 0x1a,
	0xe5, 0x89, 0xea, 0x38, 0xa3, 0x79, 0x7b, 0x9c, 0xef, 0x4c, 0x8e, 0xd6, 0x08, 0x79, 0x8e, 0xbf,
	0xe3, 0x27, 0xa7, 0x1a, 0xaf, 0x2e, 0x4f, 0x5f, 0x03, 0x35, 0xb9, 0x41, 0x96, 0xa4, 0x51, 0xbd,
	0x26, 0x7f, 0xcf, 0xe0, 0x6f, 0xdd, 0x55, 0xab, 0x99, 0xc0, 0x73, 0x05, 0x2c, 0xe1, 0x85, 0x3a,
	0xe8, 0xbc, 0x53, 0xb4, 0xc3, 0x20, 0x0a, 0x66, 0x83, 0xc5, 0xa1, 0xee, 0xdc, 0xf4, 0x97, 0x9b,
	0xbe, 0x13, 0x8f, 0x2e, 0x7b, 0x30, 0x79, 0x05, 0x97, 0xfd, 0xf7, 0xb7, 0x49, 0x2f, 0x0a, 0x92,
	0xfd, 0x0e, 0x5b, 0xda, 0xf0, 0x5c, 0xed, 0xb5, 0x13, 0xa0, 0x1d, 0xf6, 0xfe, 0x50, 0xd0, 0x6f,
	0xa0, 0xa5, 0x9d, 0x3e, 0xaa, 0xe8, 0x67, 0x49, 0x2e, 0xc9, 0x31, 0x84, 0x67, 0xea, 0x7f, 0xbb,
	0xf7, 0x30, 0x88, 0xfe, 0xcd, 0x06, 0x8b, 0x63, 0xfd, 0xfb, 0xba, 0xba, 0x29, 0x4b, 0x3a, 0x64,
	0xd5, 0x6f, 0x2d, 0x4e, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x12, 0x17, 0xac, 0xb6, 0xb7, 0x01,
	0x00, 0x00,
}
