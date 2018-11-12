// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core.proto

package cored

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	state "github.com/nayotta/metathings/pkg/proto/common/state"
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

type Core struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ProjectId            string          `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	OwnerId              string          `protobuf:"bytes,4,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	State                state.CoreState `protobuf:"varint,5,opt,name=state,proto3,enum=ai.metathings.state.CoreState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Core) Reset()         { *m = Core{} }
func (m *Core) String() string { return proto.CompactTextString(m) }
func (*Core) ProtoMessage()    {}
func (*Core) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{0}
}

func (m *Core) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Core.Unmarshal(m, b)
}
func (m *Core) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Core.Marshal(b, m, deterministic)
}
func (m *Core) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Core.Merge(m, src)
}
func (m *Core) XXX_Size() int {
	return xxx_messageInfo_Core.Size(m)
}
func (m *Core) XXX_DiscardUnknown() {
	xxx_messageInfo_Core.DiscardUnknown(m)
}

var xxx_messageInfo_Core proto.InternalMessageInfo

func (m *Core) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Core) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Core) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *Core) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func (m *Core) GetState() state.CoreState {
	if m != nil {
		return m.State
	}
	return state.CoreState_CORE_STATE_UNKNOWN
}

func init() {
	proto.RegisterType((*Core)(nil), "ai.metathings.service.cored.Core")
}

func init() { proto.RegisterFile("core.proto", fileDescriptor_f7e43720d1edc0fe) }

var fileDescriptor_f7e43720d1edc0fe = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x4f, 0x41, 0x4e, 0x03, 0x31,
	0x0c, 0x54, 0x96, 0x2d, 0x50, 0x1f, 0x7a, 0xc8, 0x29, 0x80, 0x40, 0x15, 0xa7, 0x9e, 0x12, 0x09,
	0xf8, 0x01, 0xe2, 0xd0, 0x6b, 0x79, 0x40, 0x95, 0x26, 0xd6, 0x36, 0xa0, 0xc4, 0xab, 0xac, 0x01,
	0xf1, 0x15, 0x5e, 0x8b, 0xe2, 0x45, 0x42, 0xea, 0x6d, 0x3c, 0x33, 0xf6, 0x8c, 0x01, 0x02, 0x55,
	0xb4, 0x63, 0x25, 0x26, 0x7d, 0xe3, 0x93, 0xcd, 0xc8, 0x9e, 0x8f, 0xa9, 0x0c, 0x93, 0x9d, 0xb0,
	0x7e, 0xa6, 0x80, 0xb6, 0x39, 0xe2, 0xf5, 0xcb, 0x90, 0xf8, 0xf8, 0x71, 0xb0, 0x81, 0xb2, 0x2b,
	0xfe, 0x9b, 0x98, 0xbd, 0xfb, 0x37, 0xbb, 0xf1, 0x7d, 0x70, 0x72, 0xc5, 0x05, 0xca, 0x99, 0x8a,
	0x9b, 0xd8, 0x33, 0xba, 0xb6, 0xbb, 0x17, 0x38, 0x67, 0xdc, 0xff, 0x28, 0xe8, 0x9f, 0xa9, 0xa2,
	0x5e, 0x41, 0x97, 0xa2, 0x51, 0x6b, 0xb5, 0x59, 0xee, 0xba, 0x14, 0xb5, 0x86, 0xbe, 0xf8, 0x8c,
	0xa6, 0x13, 0x46, 0xb0, 0xbe, 0x05, 0x18, 0x2b, 0xbd, 0x61, 0xe0, 0x7d, 0x8a, 0xe6, 0x4c, 0x94,
	0xe5, 0x1f, 0xb3, 0x8d, 0xfa, 0x0a, 0x2e, 0xe9, 0xab, 0x60, 0x6d, 0x62, 0x2f, 0xe2, 0x85, 0xcc,
	0xdb, 0xa8, 0x9f, 0x60, 0x21, 0xa9, 0x66, 0xb1, 0x56, 0x9b, 0xd5, 0xc3, 0x9d, 0x3d, 0x79, 0x4d,
	0x1a, 0xb5, 0x1e, 0xaf, 0x0d, 0xed, 0x66, 0xf3, 0xe1, 0x5c, 0x3a, 0x3e, 0xfe, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x1a, 0x96, 0xab, 0x45, 0x15, 0x01, 0x00, 0x00,
}
