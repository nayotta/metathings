// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_evaluators_by_source.proto

package evaluatord

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

type ListEvaluatorsBySourceRequest struct {
	Source               *OpResource `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListEvaluatorsBySourceRequest) Reset()         { *m = ListEvaluatorsBySourceRequest{} }
func (m *ListEvaluatorsBySourceRequest) String() string { return proto.CompactTextString(m) }
func (*ListEvaluatorsBySourceRequest) ProtoMessage()    {}
func (*ListEvaluatorsBySourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e8648dd5493859c, []int{0}
}

func (m *ListEvaluatorsBySourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEvaluatorsBySourceRequest.Unmarshal(m, b)
}
func (m *ListEvaluatorsBySourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEvaluatorsBySourceRequest.Marshal(b, m, deterministic)
}
func (m *ListEvaluatorsBySourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEvaluatorsBySourceRequest.Merge(m, src)
}
func (m *ListEvaluatorsBySourceRequest) XXX_Size() int {
	return xxx_messageInfo_ListEvaluatorsBySourceRequest.Size(m)
}
func (m *ListEvaluatorsBySourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEvaluatorsBySourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListEvaluatorsBySourceRequest proto.InternalMessageInfo

func (m *ListEvaluatorsBySourceRequest) GetSource() *OpResource {
	if m != nil {
		return m.Source
	}
	return nil
}

type ListEvaluatorsBySourceResponse struct {
	Evaluators           []*Evaluator `protobuf:"bytes,1,rep,name=evaluators,proto3" json:"evaluators,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListEvaluatorsBySourceResponse) Reset()         { *m = ListEvaluatorsBySourceResponse{} }
func (m *ListEvaluatorsBySourceResponse) String() string { return proto.CompactTextString(m) }
func (*ListEvaluatorsBySourceResponse) ProtoMessage()    {}
func (*ListEvaluatorsBySourceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e8648dd5493859c, []int{1}
}

func (m *ListEvaluatorsBySourceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEvaluatorsBySourceResponse.Unmarshal(m, b)
}
func (m *ListEvaluatorsBySourceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEvaluatorsBySourceResponse.Marshal(b, m, deterministic)
}
func (m *ListEvaluatorsBySourceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEvaluatorsBySourceResponse.Merge(m, src)
}
func (m *ListEvaluatorsBySourceResponse) XXX_Size() int {
	return xxx_messageInfo_ListEvaluatorsBySourceResponse.Size(m)
}
func (m *ListEvaluatorsBySourceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEvaluatorsBySourceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListEvaluatorsBySourceResponse proto.InternalMessageInfo

func (m *ListEvaluatorsBySourceResponse) GetEvaluators() []*Evaluator {
	if m != nil {
		return m.Evaluators
	}
	return nil
}

func init() {
	proto.RegisterType((*ListEvaluatorsBySourceRequest)(nil), "ai.metathings.service.evaluatord.ListEvaluatorsBySourceRequest")
	proto.RegisterType((*ListEvaluatorsBySourceResponse)(nil), "ai.metathings.service.evaluatord.ListEvaluatorsBySourceResponse")
}

func init() { proto.RegisterFile("list_evaluators_by_source.proto", fileDescriptor_0e8648dd5493859c) }

var fileDescriptor_0e8648dd5493859c = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcf, 0xc9, 0x2c, 0x2e,
	0x89, 0x4f, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x2c, 0xc9, 0x2f, 0x2a, 0x8e, 0x4f, 0xaa, 0x8c, 0x2f,
	0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x48, 0xcc, 0xd4,
	0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0xd5, 0x83, 0xeb, 0x48, 0x91, 0x12, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49,
	0xd5, 0x87, 0x31, 0x20, 0x5a, 0xa5, 0xb8, 0x73, 0xf3, 0x53, 0x52, 0x73, 0x20, 0x1c, 0xa5, 0x7c,
	0x2e, 0x59, 0x9f, 0xcc, 0xe2, 0x12, 0x57, 0xb8, 0x4d, 0x4e, 0x95, 0xc1, 0x60, 0x7b, 0x82, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xfc, 0xb8, 0xd8, 0x20, 0x16, 0x4b, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x1b, 0xe9, 0xe8, 0x11, 0xb2, 0x59, 0xcf, 0xbf, 0x20, 0x28, 0x15, 0xa2, 0xc7, 0x89, 0xe3,
	0x97, 0x13, 0x6b, 0x17, 0x23, 0x93, 0x00, 0x63, 0x10, 0xd4, 0x14, 0xa5, 0x5c, 0x2e, 0x39, 0x5c,
	0x16, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x79, 0x73, 0x71, 0x21, 0x3c, 0x2e, 0xc1, 0xa8,
	0xc0, 0xac, 0xc1, 0x6d, 0xa4, 0x4d, 0xd8, 0x56, 0xb8, 0x89, 0x41, 0x48, 0xda, 0x9d, 0x78, 0xa2,
	0x10, 0xbc, 0x94, 0x24, 0x36, 0xb0, 0xa7, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7c, 0xaa,
	0xad, 0xf0, 0x5f, 0x01, 0x00, 0x00,
}