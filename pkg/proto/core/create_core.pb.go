// Code generated by protoc-gen-go. DO NOT EDIT.
// source: create_core.proto

package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateCoreRequest struct {
	Name *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CreateCoreRequest) Reset()                    { *m = CreateCoreRequest{} }
func (m *CreateCoreRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateCoreRequest) ProtoMessage()               {}
func (*CreateCoreRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CreateCoreRequest) GetName() *google_protobuf.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

type CreateCoreResponse struct {
	Core *Core `protobuf:"bytes,1,opt,name=core" json:"core,omitempty"`
}

func (m *CreateCoreResponse) Reset()                    { *m = CreateCoreResponse{} }
func (m *CreateCoreResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateCoreResponse) ProtoMessage()               {}
func (*CreateCoreResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *CreateCoreResponse) GetCore() *Core {
	if m != nil {
		return m.Core
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateCoreRequest)(nil), "ai.metathings.service.core.CreateCoreRequest")
	proto.RegisterType((*CreateCoreResponse)(nil), "ai.metathings.service.core.CreateCoreResponse")
}

func init() { proto.RegisterFile("create_core.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0x29, 0x14, 0x0f, 0xf1, 0xd4, 0x3d, 0xc9, 0x22, 0x52, 0x7a, 0xf2, 0xd2, 0x89, 0xa8,
	0xf8, 0x05, 0x8a, 0x17, 0x8f, 0x15, 0xbc, 0xca, 0x6c, 0x1c, 0xd3, 0xe0, 0x26, 0x13, 0x27, 0x93,
	0xee, 0xd7, 0x97, 0xa6, 0xf5, 0xcf, 0xc5, 0x5b, 0x20, 0xbf, 0xf7, 0xde, 0x6f, 0xcc, 0xc2, 0x09,
	0xa1, 0xd2, 0xab, 0x63, 0x21, 0xc8, 0xc2, 0xca, 0x5d, 0x8f, 0x01, 0x22, 0x29, 0xea, 0x2e, 0x24,
	0x5f, 0xa0, 0x90, 0xec, 0x83, 0x23, 0x38, 0x10, 0xfd, 0x95, 0x67, 0xf6, 0x23, 0xd9, 0x46, 0x0e,
	0xf5, 0xdd, 0x4e, 0x82, 0x39, 0x93, 0x94, 0x63, 0xb6, 0x7f, 0xf0, 0x41, 0x77, 0x75, 0x00, 0xc7,
	0xd1, 0xc6, 0x29, 0xe8, 0x07, 0x4f, 0xd6, 0xf3, 0xba, 0x7d, 0xae, 0xf7, 0x38, 0x86, 0x37, 0x54,
	0x96, 0x62, 0x7f, 0x9e, 0xa7, 0x9c, 0xf9, 0xdd, 0x5f, 0x3d, 0x9a, 0xc5, 0xa6, 0x49, 0x6d, 0x58,
	0x68, 0x4b, 0x9f, 0x95, 0x8a, 0x76, 0x37, 0x66, 0x9e, 0x30, 0xd2, 0xc5, 0x6c, 0x39, 0xbb, 0x3e,
	0xbf, 0xbd, 0x84, 0xa3, 0x07, 0x7c, 0x7b, 0xc0, 0xb3, 0x4a, 0x48, 0xfe, 0x05, 0xc7, 0x4a, 0xdb,
	0x46, 0xae, 0x9e, 0x4c, 0xf7, 0xb7, 0xa6, 0x64, 0x4e, 0x85, 0xba, 0x7b, 0x33, 0x3f, 0x4c, 0x9d,
	0x7a, 0x96, 0xf0, 0xff, 0xad, 0xd0, 0x72, 0x8d, 0x1e, 0xce, 0xda, 0xce, 0xdd, 0x57, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x50, 0xfe, 0xa0, 0xa8, 0x2e, 0x01, 0x00, 0x00,
}
