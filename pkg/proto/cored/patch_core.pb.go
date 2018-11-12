// Code generated by protoc-gen-go. DO NOT EDIT.
// source: patch_core.proto

package cored

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/mwitkow/go-proto-validators"
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

type PatchCoreRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	State                state.CoreState       `protobuf:"varint,3,opt,name=state,proto3,enum=ai.metathings.state.CoreState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PatchCoreRequest) Reset()         { *m = PatchCoreRequest{} }
func (m *PatchCoreRequest) String() string { return proto.CompactTextString(m) }
func (*PatchCoreRequest) ProtoMessage()    {}
func (*PatchCoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5e24b2c43c37899, []int{0}
}

func (m *PatchCoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchCoreRequest.Unmarshal(m, b)
}
func (m *PatchCoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchCoreRequest.Marshal(b, m, deterministic)
}
func (m *PatchCoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchCoreRequest.Merge(m, src)
}
func (m *PatchCoreRequest) XXX_Size() int {
	return xxx_messageInfo_PatchCoreRequest.Size(m)
}
func (m *PatchCoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchCoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PatchCoreRequest proto.InternalMessageInfo

func (m *PatchCoreRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PatchCoreRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *PatchCoreRequest) GetState() state.CoreState {
	if m != nil {
		return m.State
	}
	return state.CoreState_CORE_STATE_UNKNOWN
}

type PatchCoreResponse struct {
	Core                 *Core    `protobuf:"bytes,1,opt,name=core,proto3" json:"core,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PatchCoreResponse) Reset()         { *m = PatchCoreResponse{} }
func (m *PatchCoreResponse) String() string { return proto.CompactTextString(m) }
func (*PatchCoreResponse) ProtoMessage()    {}
func (*PatchCoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5e24b2c43c37899, []int{1}
}

func (m *PatchCoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchCoreResponse.Unmarshal(m, b)
}
func (m *PatchCoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchCoreResponse.Marshal(b, m, deterministic)
}
func (m *PatchCoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchCoreResponse.Merge(m, src)
}
func (m *PatchCoreResponse) XXX_Size() int {
	return xxx_messageInfo_PatchCoreResponse.Size(m)
}
func (m *PatchCoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchCoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PatchCoreResponse proto.InternalMessageInfo

func (m *PatchCoreResponse) GetCore() *Core {
	if m != nil {
		return m.Core
	}
	return nil
}

func init() {
	proto.RegisterType((*PatchCoreRequest)(nil), "ai.metathings.service.cored.PatchCoreRequest")
	proto.RegisterType((*PatchCoreResponse)(nil), "ai.metathings.service.cored.PatchCoreResponse")
}

func init() { proto.RegisterFile("patch_core.proto", fileDescriptor_d5e24b2c43c37899) }

var fileDescriptor_d5e24b2c43c37899 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4e, 0x2a, 0x31,
	0x14, 0x86, 0x33, 0x73, 0xb9, 0x2c, 0x6a, 0x62, 0x70, 0x56, 0x04, 0x0d, 0x22, 0x2b, 0x36, 0xb4,
	0x06, 0xd1, 0x07, 0xd0, 0xb8, 0x71, 0x65, 0x86, 0xc4, 0x2d, 0x29, 0x33, 0xc7, 0xd2, 0xc0, 0xf4,
	0xd4, 0xf6, 0x0c, 0xc4, 0xb7, 0xf2, 0x8d, 0x4c, 0x7c, 0x12, 0xd3, 0x0e, 0x0a, 0x71, 0xa1, 0xbb,
	0x93, 0xf4, 0xff, 0x4f, 0xbf, 0xaf, 0x65, 0x1d, 0x2b, 0xa9, 0x58, 0xce, 0x0b, 0x74, 0xc0, 0xad,
	0x43, 0xc2, 0xec, 0x54, 0x6a, 0x5e, 0x01, 0x49, 0x5a, 0x6a, 0xa3, 0x3c, 0xf7, 0xe0, 0x36, 0xba,
	0x00, 0x1e, 0x12, 0x65, 0xaf, 0xaf, 0x10, 0xd5, 0x1a, 0x44, 0x8c, 0x2e, 0xea, 0x67, 0xb1, 0x75,
	0xd2, 0x5a, 0x70, 0xbe, 0x29, 0xf7, 0x6e, 0x94, 0xa6, 0x65, 0xbd, 0xe0, 0x05, 0x56, 0xa2, 0xda,
	0x6a, 0x5a, 0xe1, 0x56, 0x28, 0x1c, 0xc7, 0xc3, 0xf1, 0x46, 0xae, 0x75, 0x29, 0x09, 0x9d, 0x17,
	0xdf, 0xe3, 0xae, 0x77, 0x7f, 0xd0, 0x33, 0xf2, 0x15, 0x89, 0xa4, 0xd8, 0x43, 0x08, 0xbb, 0x52,
	0xcd, 0x95, 0xa2, 0xc0, 0xaa, 0x42, 0x23, 0x3c, 0x49, 0x02, 0x11, 0x98, 0xe6, 0x71, 0xdc, 0xad,
	0x61, 0x7b, 0x8f, 0xe1, 0x5b, 0xc2, 0x3a, 0x8f, 0x41, 0xee, 0x0e, 0x1d, 0xe4, 0xf0, 0x52, 0x83,
	0xa7, 0x6c, 0xca, 0x52, 0x5d, 0x76, 0x93, 0x41, 0x32, 0x3a, 0x9a, 0x9c, 0xf1, 0x46, 0x86, 0x7f,
	0xc9, 0xf0, 0x19, 0x39, 0x6d, 0xd4, 0x93, 0x5c, 0xd7, 0x70, 0xdb, 0xfe, 0x78, 0x3f, 0x4f, 0x07,
	0x49, 0x9e, 0xea, 0x32, 0xbb, 0x64, 0x2d, 0x23, 0x2b, 0xe8, 0xa6, 0x7f, 0xf7, 0xf2, 0x98, 0xcc,
	0xa6, 0xec, 0x7f, 0xe4, 0xea, 0xfe, 0x1b, 0x24, 0xa3, 0xe3, 0x49, 0x9f, 0xff, 0x78, 0xd4, 0xc8,
	0x1c, 0xc0, 0x66, 0x61, 0xca, 0x9b, 0xf0, 0xf0, 0x81, 0x9d, 0x1c, 0x10, 0x7b, 0x8b, 0xc6, 0x43,
	0x76, 0xcd, 0x5a, 0xc1, 0x6a, 0x07, 0x7d, 0xc1, 0x7f, 0xf9, 0x9e, 0xb8, 0x31, 0x8f, 0xf1, 0x45,
	0x3b, 0xd2, 0x5d, 0x7d, 0x06, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x22, 0x56, 0xee, 0xe1, 0x01, 0x00,
	0x00,
}
