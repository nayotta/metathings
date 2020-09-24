// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_flow_set.proto

package deviced

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetFlowSetRequest struct {
	FlowSet              *OpFlowSet `protobuf:"bytes,1,opt,name=flow_set,json=flowSet,proto3" json:"flow_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetFlowSetRequest) Reset()         { *m = GetFlowSetRequest{} }
func (m *GetFlowSetRequest) String() string { return proto.CompactTextString(m) }
func (*GetFlowSetRequest) ProtoMessage()    {}
func (*GetFlowSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf816d140cde1d0d, []int{0}
}

func (m *GetFlowSetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFlowSetRequest.Unmarshal(m, b)
}
func (m *GetFlowSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFlowSetRequest.Marshal(b, m, deterministic)
}
func (m *GetFlowSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFlowSetRequest.Merge(m, src)
}
func (m *GetFlowSetRequest) XXX_Size() int {
	return xxx_messageInfo_GetFlowSetRequest.Size(m)
}
func (m *GetFlowSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFlowSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFlowSetRequest proto.InternalMessageInfo

func (m *GetFlowSetRequest) GetFlowSet() *OpFlowSet {
	if m != nil {
		return m.FlowSet
	}
	return nil
}

type GetFlowSetResponse struct {
	FlowSet              *FlowSet `protobuf:"bytes,1,opt,name=flow_set,json=flowSet,proto3" json:"flow_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFlowSetResponse) Reset()         { *m = GetFlowSetResponse{} }
func (m *GetFlowSetResponse) String() string { return proto.CompactTextString(m) }
func (*GetFlowSetResponse) ProtoMessage()    {}
func (*GetFlowSetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf816d140cde1d0d, []int{1}
}

func (m *GetFlowSetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFlowSetResponse.Unmarshal(m, b)
}
func (m *GetFlowSetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFlowSetResponse.Marshal(b, m, deterministic)
}
func (m *GetFlowSetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFlowSetResponse.Merge(m, src)
}
func (m *GetFlowSetResponse) XXX_Size() int {
	return xxx_messageInfo_GetFlowSetResponse.Size(m)
}
func (m *GetFlowSetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFlowSetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFlowSetResponse proto.InternalMessageInfo

func (m *GetFlowSetResponse) GetFlowSet() *FlowSet {
	if m != nil {
		return m.FlowSet
	}
	return nil
}

func init() {
	proto.RegisterType((*GetFlowSetRequest)(nil), "ai.metathings.service.deviced.GetFlowSetRequest")
	proto.RegisterType((*GetFlowSetResponse)(nil), "ai.metathings.service.deviced.GetFlowSetResponse")
}

func init() { proto.RegisterFile("get_flow_set.proto", fileDescriptor_cf816d140cde1d0d) }

var fileDescriptor_cf816d140cde1d0d = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x4f, 0x2d, 0x89,
	0x4f, 0xcb, 0xc9, 0x2f, 0x8f, 0x2f, 0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92,
	0x4d, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e,
	0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4b, 0x49, 0x05, 0x51, 0x29, 0x52, 0xe2, 0x65, 0x89, 0x39,
	0x99, 0x29, 0x89, 0x25, 0xa9, 0xfa, 0x30, 0x06, 0x44, 0x9f, 0x14, 0x77, 0x6e, 0x7e, 0x4a, 0x6a,
	0x0e, 0x84, 0xa3, 0x94, 0xc4, 0x25, 0xe8, 0x9e, 0x5a, 0xe2, 0x96, 0x93, 0x5f, 0x1e, 0x9c, 0x5a,
	0x12, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0xe4, 0xcb, 0xc5, 0x01, 0xb3, 0x4b, 0x82, 0x51,
	0x81, 0x51, 0x83, 0xdb, 0x48, 0x43, 0x0f, 0xaf, 0x65, 0x7a, 0xfe, 0x05, 0x50, 0x23, 0x9c, 0x38,
	0x7e, 0x39, 0xb1, 0x76, 0x31, 0x32, 0x09, 0x30, 0x06, 0xb1, 0xa7, 0x41, 0x84, 0x94, 0xc2, 0xb9,
	0x84, 0x90, 0xed, 0x28, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x72, 0xc4, 0xb0, 0x44, 0x8d, 0x80,
	0x25, 0x30, 0x13, 0x60, 0x06, 0x27, 0xb1, 0x81, 0xfd, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0xef, 0x40, 0x77, 0x9e, 0x1e, 0x01, 0x00, 0x00,
}
