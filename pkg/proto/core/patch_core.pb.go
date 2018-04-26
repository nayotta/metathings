// Code generated by protoc-gen-go. DO NOT EDIT.
// source: patch_core.proto

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

type PatchCoreRequest struct {
	Id   *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *PatchCoreRequest) Reset()                    { *m = PatchCoreRequest{} }
func (m *PatchCoreRequest) String() string            { return proto.CompactTextString(m) }
func (*PatchCoreRequest) ProtoMessage()               {}
func (*PatchCoreRequest) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *PatchCoreRequest) GetId() *google_protobuf.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PatchCoreRequest) GetName() *google_protobuf.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

type PatchCoreResponse struct {
	Core *Core `protobuf:"bytes,1,opt,name=core" json:"core,omitempty"`
}

func (m *PatchCoreResponse) Reset()                    { *m = PatchCoreResponse{} }
func (m *PatchCoreResponse) String() string            { return proto.CompactTextString(m) }
func (*PatchCoreResponse) ProtoMessage()               {}
func (*PatchCoreResponse) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *PatchCoreResponse) GetCore() *Core {
	if m != nil {
		return m.Core
	}
	return nil
}

func init() {
	proto.RegisterType((*PatchCoreRequest)(nil), "ai.metathings.service.core.PatchCoreRequest")
	proto.RegisterType((*PatchCoreResponse)(nil), "ai.metathings.service.core.PatchCoreResponse")
}

func init() { proto.RegisterFile("patch_core.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0x69, 0x28, 0x3d, 0xac, 0x97, 0x9a, 0x53, 0x09, 0xa2, 0xa1, 0x27, 0x2f, 0x9d, 0x15,
	0x2d, 0x3e, 0x80, 0x9e, 0xbc, 0x49, 0x04, 0xaf, 0xb2, 0x49, 0xc6, 0xcd, 0x60, 0x92, 0x59, 0x77,
	0x27, 0x0d, 0xf8, 0xb2, 0x82, 0x4f, 0x22, 0xd9, 0xd6, 0x3f, 0x17, 0xf1, 0xb6, 0xb0, 0xbf, 0x8f,
	0xef, 0x63, 0xd4, 0xd2, 0x19, 0xa9, 0x9a, 0xa7, 0x8a, 0x3d, 0x82, 0xf3, 0x2c, 0x9c, 0x66, 0x86,
	0xa0, 0x43, 0x31, 0xd2, 0x50, 0x6f, 0x03, 0x04, 0xf4, 0x3b, 0xaa, 0x10, 0xa6, 0x45, 0x76, 0x6a,
	0x99, 0x6d, 0x8b, 0x3a, 0x2e, 0xcb, 0xe1, 0x59, 0x8f, 0xde, 0x38, 0x87, 0x3e, 0xec, 0xd9, 0xec,
	0xda, 0x92, 0x34, 0x43, 0x09, 0x15, 0x77, 0xba, 0x1b, 0x49, 0x5e, 0x78, 0xd4, 0x96, 0x37, 0xf1,
	0x73, 0xb3, 0x33, 0x2d, 0xd5, 0x46, 0xd8, 0x07, 0xfd, 0xfd, 0x3c, 0x70, 0xea, 0xc7, 0xbf, 0x7e,
	0x53, 0xcb, 0xfb, 0xa9, 0xe9, 0x96, 0x3d, 0x16, 0xf8, 0x3a, 0x60, 0x90, 0x74, 0xab, 0x12, 0xaa,
	0x57, 0xb3, 0x7c, 0x76, 0x7e, 0x74, 0x79, 0x02, 0xfb, 0x08, 0xf8, 0x8a, 0x80, 0x07, 0xf1, 0xd4,
	0xdb, 0x47, 0xd3, 0x0e, 0x78, 0xb3, 0xf8, 0x78, 0x3f, 0x4b, 0xf2, 0x59, 0x91, 0x50, 0x9d, 0x5e,
	0xa8, 0x79, 0x6f, 0x3a, 0x5c, 0x25, 0xff, 0x73, 0x45, 0x5c, 0xae, 0xef, 0xd4, 0xf1, 0x2f, 0x77,
	0x70, 0xdc, 0x07, 0x4c, 0xb7, 0x6a, 0x3e, 0xe5, 0x1d, 0xf4, 0x39, 0xfc, 0x7d, 0x1f, 0x88, 0x5c,
	0x5c, 0x97, 0x8b, 0xa8, 0xb9, 0xfa, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x77, 0x85, 0x30, 0x61,
	0x01, 0x00, 0x00,
}
