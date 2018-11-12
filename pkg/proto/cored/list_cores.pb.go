// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_cores.proto

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

type ListCoresRequest struct {
	Name                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ProjectId            *wrappers.StringValue `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	OwnerId              *wrappers.StringValue `protobuf:"bytes,3,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	State                state.CoreState       `protobuf:"varint,4,opt,name=state,proto3,enum=ai.metathings.state.CoreState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListCoresRequest) Reset()         { *m = ListCoresRequest{} }
func (m *ListCoresRequest) String() string { return proto.CompactTextString(m) }
func (*ListCoresRequest) ProtoMessage()    {}
func (*ListCoresRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecd305d3b0bd9062, []int{0}
}

func (m *ListCoresRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCoresRequest.Unmarshal(m, b)
}
func (m *ListCoresRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCoresRequest.Marshal(b, m, deterministic)
}
func (m *ListCoresRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCoresRequest.Merge(m, src)
}
func (m *ListCoresRequest) XXX_Size() int {
	return xxx_messageInfo_ListCoresRequest.Size(m)
}
func (m *ListCoresRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCoresRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCoresRequest proto.InternalMessageInfo

func (m *ListCoresRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ListCoresRequest) GetProjectId() *wrappers.StringValue {
	if m != nil {
		return m.ProjectId
	}
	return nil
}

func (m *ListCoresRequest) GetOwnerId() *wrappers.StringValue {
	if m != nil {
		return m.OwnerId
	}
	return nil
}

func (m *ListCoresRequest) GetState() state.CoreState {
	if m != nil {
		return m.State
	}
	return state.CoreState_CORE_STATE_UNKNOWN
}

type ListCoresResponse struct {
	Cores                []*Core  `protobuf:"bytes,1,rep,name=cores,proto3" json:"cores,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCoresResponse) Reset()         { *m = ListCoresResponse{} }
func (m *ListCoresResponse) String() string { return proto.CompactTextString(m) }
func (*ListCoresResponse) ProtoMessage()    {}
func (*ListCoresResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecd305d3b0bd9062, []int{1}
}

func (m *ListCoresResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCoresResponse.Unmarshal(m, b)
}
func (m *ListCoresResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCoresResponse.Marshal(b, m, deterministic)
}
func (m *ListCoresResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCoresResponse.Merge(m, src)
}
func (m *ListCoresResponse) XXX_Size() int {
	return xxx_messageInfo_ListCoresResponse.Size(m)
}
func (m *ListCoresResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCoresResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCoresResponse proto.InternalMessageInfo

func (m *ListCoresResponse) GetCores() []*Core {
	if m != nil {
		return m.Cores
	}
	return nil
}

func init() {
	proto.RegisterType((*ListCoresRequest)(nil), "ai.metathings.service.cored.ListCoresRequest")
	proto.RegisterType((*ListCoresResponse)(nil), "ai.metathings.service.cored.ListCoresResponse")
}

func init() { proto.RegisterFile("list_cores.proto", fileDescriptor_ecd305d3b0bd9062) }

var fileDescriptor_ecd305d3b0bd9062 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4f, 0xf2, 0x40,
	0x10, 0xc6, 0xd3, 0x17, 0x78, 0xff, 0x2c, 0xc9, 0x1b, 0xec, 0xa9, 0x41, 0x43, 0x90, 0x13, 0x17,
	0x76, 0x0d, 0x1a, 0x39, 0x78, 0x34, 0x1e, 0x48, 0x38, 0x95, 0xc4, 0x2b, 0x59, 0xda, 0xb1, 0xac,
	0xb4, 0x3b, 0x75, 0x77, 0x4a, 0xe3, 0xb7, 0xf6, 0x23, 0x98, 0xdd, 0x56, 0x25, 0x1e, 0x0c, 0xb7,
	0x27, 0xed, 0xef, 0x99, 0x9d, 0x5f, 0x86, 0x0d, 0x72, 0x65, 0x69, 0x93, 0xa0, 0x01, 0xcb, 0x4b,
	0x83, 0x84, 0xe1, 0xb9, 0x54, 0xbc, 0x00, 0x92, 0xb4, 0x53, 0x3a, 0xb3, 0xdc, 0x82, 0x39, 0xa8,
	0x04, 0xb8, 0x43, 0xd2, 0xe1, 0x28, 0x43, 0xcc, 0x72, 0x10, 0x1e, 0xdd, 0x56, 0x4f, 0xa2, 0x36,
	0xb2, 0x2c, 0xc1, 0xb4, 0xe5, 0xe1, 0x6d, 0xa6, 0x68, 0x57, 0x6d, 0x79, 0x82, 0x85, 0x28, 0x6a,
	0x45, 0x7b, 0xac, 0x45, 0x86, 0x33, 0xff, 0x73, 0x76, 0x90, 0xb9, 0x4a, 0x25, 0xa1, 0xb1, 0xe2,
	0x33, 0xb6, 0xbd, 0x87, 0xa3, 0x9e, 0x96, 0xaf, 0x48, 0x24, 0xc5, 0xd7, 0x12, 0xa2, 0xdc, 0x67,
	0xcd, 0x93, 0x22, 0xc1, 0xa2, 0x40, 0x2d, 0x2c, 0x49, 0x02, 0xe1, 0x76, 0xda, 0xf8, 0xd8, 0x8e,
	0x61, 0xee, 0x4b, 0x93, 0x27, 0x6f, 0x01, 0x1b, 0xac, 0x94, 0xa5, 0x7b, 0xe7, 0x16, 0xc3, 0x4b,
	0x05, 0x96, 0xc2, 0x2b, 0xd6, 0xd5, 0xb2, 0x80, 0x28, 0x18, 0x07, 0xd3, 0xfe, 0xfc, 0x82, 0x37,
	0x3a, 0xfc, 0x43, 0x87, 0xaf, 0xc9, 0x28, 0x9d, 0x3d, 0xca, 0xbc, 0x82, 0xd8, 0x93, 0xe1, 0x1d,
	0x63, 0xa5, 0xc1, 0x67, 0x48, 0x68, 0xa3, 0xd2, 0xe8, 0xd7, 0x09, 0xbd, 0x7f, 0x2d, 0xbf, 0x4c,
	0xc3, 0x05, 0xfb, 0x8b, 0xb5, 0x06, 0xe3, 0xaa, 0x9d, 0x13, 0xaa, 0x7f, 0x3c, 0xbd, 0x4c, 0xc3,
	0x1b, 0xd6, 0xf3, 0x5e, 0x51, 0x77, 0x1c, 0x4c, 0xff, 0xcf, 0x47, 0xfc, 0xdb, 0x51, 0xbc, 0xb3,
	0x33, 0x5b, 0xbb, 0x14, 0x37, 0xf0, 0x64, 0xc5, 0xce, 0x8e, 0x8c, 0x6d, 0x89, 0xda, 0x42, 0xb8,
	0x60, 0x3d, 0x7f, 0xde, 0x28, 0x18, 0x77, 0xa6, 0xfd, 0xf9, 0x25, 0xff, 0xe1, 0xbe, 0x7e, 0x64,
	0xdc, 0xf0, 0xdb, 0xdf, 0x7e, 0xc5, 0xeb, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x01, 0xf0,
	0x97, 0x23, 0x02, 0x00, 0x00,
}
