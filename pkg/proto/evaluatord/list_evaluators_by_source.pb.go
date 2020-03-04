// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_evaluators_by_source.proto

package evaluatord

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
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
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x40, 0x89, 0x42, 0x0e, 0x1b, 0x4f, 0x39, 0x95, 0x82, 0x36, 0xf4, 0x54, 0xd0, 0x6e, 0xa0,
	0x82, 0x1f, 0x10, 0xf0, 0x64, 0x41, 0x88, 0x37, 0x2f, 0x61, 0x93, 0x0c, 0xe9, 0x60, 0xb6, 0x13,
	0x33, 0x93, 0x94, 0x7e, 0xad, 0xe0, 0x97, 0x08, 0x59, 0xd9, 0x7a, 0x91, 0xdc, 0x76, 0x0e, 0xef,
	0xcd, 0xdb, 0x51, 0xab, 0x16, 0x59, 0x0a, 0x18, 0x4d, 0x3b, 0x18, 0xa1, 0x9e, 0x8b, 0xf2, 0x5c,
	0x30, 0x0d, 0x7d, 0x05, 0xba, 0xeb, 0x49, 0x28, 0x4e, 0x0c, 0x6a, 0x0b, 0x62, 0xe4, 0x80, 0xc7,
	0x86, 0x35, 0x43, 0x3f, 0x62, 0x05, 0xda, 0x13, 0xf5, 0xf2, 0xa9, 0x41, 0x39, 0x0c, 0xa5, 0xae,
	0xc8, 0xa6, 0xf6, 0x84, 0xf2, 0x41, 0xa7, 0xb4, 0xa1, 0xed, 0x84, 0x6f, 0x47, 0xd3, 0x62, 0x3d,
	0x99, 0x53, 0xff, 0x74, 0xe6, 0x65, 0x64, 0xa9, 0x86, 0xd6, 0x0d, 0x6b, 0xab, 0x6e, 0xf7, 0xc8,
	0xf2, 0xec, 0x43, 0xb2, 0xf3, 0xdb, 0x94, 0x91, 0xc3, 0xe7, 0x00, 0x2c, 0xf1, 0x5e, 0x85, 0xae,
	0x6b, 0x11, 0x24, 0xc1, 0x26, 0xda, 0x3d, 0xe8, 0xb9, 0x30, 0xfd, 0xda, 0xe5, 0xe0, 0x98, 0x2c,
	0xfc, 0xfe, 0x5a, 0x5d, 0x25, 0x41, 0xfe, 0xeb, 0x58, 0x5b, 0x75, 0xf7, 0xdf, 0x3a, 0xee, 0xe8,
	0xc8, 0x10, 0xbf, 0x28, 0x75, 0xb9, 0xca, 0x22, 0x48, 0xae, 0x37, 0xd1, 0xee, 0x7e, 0x7e, 0xa7,
	0x37, 0xe6, 0x7f, 0xf0, 0xec, 0xe6, 0xfd, 0x32, 0xd5, 0x65, 0x38, 0x7d, 0xf9, 0xf1, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x6c, 0x7c, 0x47, 0xb5, 0x7c, 0x01, 0x00, 0x00,
}
