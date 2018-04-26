// Code generated by protoc-gen-go. DO NOT EDIT.
// source: patch_entity.proto

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

type PatchEntityRequest struct {
	Id    *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	State EntityState                  `protobuf:"varint,2,opt,name=state,enum=ai.metathings.service.core.EntityState" json:"state,omitempty"`
}

func (m *PatchEntityRequest) Reset()                    { *m = PatchEntityRequest{} }
func (m *PatchEntityRequest) String() string            { return proto.CompactTextString(m) }
func (*PatchEntityRequest) ProtoMessage()               {}
func (*PatchEntityRequest) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *PatchEntityRequest) GetId() *google_protobuf.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PatchEntityRequest) GetState() EntityState {
	if m != nil {
		return m.State
	}
	return EntityState_ENTITY_STATE_UNKNOWN
}

type PatchEntityResponse struct {
	Entity *Entity `protobuf:"bytes,1,opt,name=entity" json:"entity,omitempty"`
}

func (m *PatchEntityResponse) Reset()                    { *m = PatchEntityResponse{} }
func (m *PatchEntityResponse) String() string            { return proto.CompactTextString(m) }
func (*PatchEntityResponse) ProtoMessage()               {}
func (*PatchEntityResponse) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

func (m *PatchEntityResponse) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func init() {
	proto.RegisterType((*PatchEntityRequest)(nil), "ai.metathings.service.core.PatchEntityRequest")
	proto.RegisterType((*PatchEntityResponse)(nil), "ai.metathings.service.core.PatchEntityResponse")
}

func init() { proto.RegisterFile("patch_entity.proto", fileDescriptor14) }

var fileDescriptor14 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xcf, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0xc0, 0x71, 0x5a, 0xb0, 0x43, 0x14, 0x87, 0xb8, 0x1c, 0x45, 0xb4, 0x74, 0xf1, 0x96, 0x4b,
	0xe1, 0x14, 0x07, 0xc1, 0x45, 0x70, 0xd7, 0x1e, 0xb8, 0x4a, 0xda, 0x3e, 0xd3, 0x87, 0x6d, 0x5f,
	0x4d, 0x5e, 0xaf, 0xf8, 0x11, 0xfc, 0x94, 0x82, 0x9f, 0x44, 0xae, 0xa9, 0xa2, 0x83, 0xb8, 0x85,
	0xe4, 0xfd, 0x93, 0x5f, 0x84, 0xec, 0x35, 0x97, 0xf5, 0x23, 0x74, 0x8c, 0xfc, 0xaa, 0x7a, 0x4b,
	0x4c, 0x32, 0xd6, 0xa8, 0x5a, 0x60, 0xcd, 0x35, 0x76, 0xc6, 0x29, 0x07, 0x76, 0x8b, 0x25, 0xa8,
	0x92, 0x2c, 0xc4, 0x27, 0x86, 0xc8, 0x34, 0x90, 0x4d, 0x93, 0xc5, 0xf0, 0x94, 0x8d, 0x56, 0xf7,
	0x3d, 0x58, 0xe7, 0xdb, 0xf8, 0xd2, 0x20, 0xd7, 0x43, 0xa1, 0x4a, 0x6a, 0xb3, 0x76, 0x44, 0x7e,
	0xa6, 0x31, 0x33, 0xb4, 0x9a, 0x0e, 0x57, 0x5b, 0xdd, 0x60, 0xa5, 0x99, 0xac, 0xcb, 0xbe, 0x97,
	0x73, 0x77, 0xf0, 0x53, 0x90, 0xbe, 0x05, 0x42, 0xde, 0xed, 0x60, 0xb7, 0xd3, 0x6e, 0x0e, 0x2f,
	0x03, 0x38, 0x96, 0x17, 0x22, 0xc4, 0x6a, 0x11, 0x24, 0xc1, 0x72, 0x7f, 0x7d, 0xac, 0xbc, 0x44,
	0x7d, 0x49, 0xd4, 0x86, 0x2d, 0x76, 0xe6, 0x41, 0x37, 0x03, 0xdc, 0x44, 0x1f, 0xef, 0xa7, 0x61,
	0x12, 0xe4, 0x21, 0x56, 0xf2, 0x5a, 0xec, 0x39, 0xd6, 0x0c, 0x8b, 0x30, 0x09, 0x96, 0x87, 0xeb,
	0x33, 0xf5, 0xf7, 0xf7, 0x94, 0x7f, 0x6f, 0xb3, 0x1b, 0xcf, 0x7d, 0x95, 0xde, 0x8b, 0xa3, 0x5f,
	0x14, 0xd7, 0x53, 0xe7, 0x40, 0x5e, 0x89, 0xc8, 0x93, 0x67, 0x4f, 0xfa, 0xff, 0xb5, 0xf9, 0x5c,
	0x14, 0xd1, 0x64, 0x3e, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x10, 0x47, 0x17, 0x5e, 0x7d, 0x01,
	0x00, 0x00,
}
