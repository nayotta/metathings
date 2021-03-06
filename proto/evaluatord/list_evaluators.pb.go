// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_evaluators.proto

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

type ListEvaluatorsRequest struct {
	Evaluator            *OpEvaluator `protobuf:"bytes,1,opt,name=evaluator,proto3" json:"evaluator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListEvaluatorsRequest) Reset()         { *m = ListEvaluatorsRequest{} }
func (m *ListEvaluatorsRequest) String() string { return proto.CompactTextString(m) }
func (*ListEvaluatorsRequest) ProtoMessage()    {}
func (*ListEvaluatorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_465e3bbdbf652446, []int{0}
}

func (m *ListEvaluatorsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEvaluatorsRequest.Unmarshal(m, b)
}
func (m *ListEvaluatorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEvaluatorsRequest.Marshal(b, m, deterministic)
}
func (m *ListEvaluatorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEvaluatorsRequest.Merge(m, src)
}
func (m *ListEvaluatorsRequest) XXX_Size() int {
	return xxx_messageInfo_ListEvaluatorsRequest.Size(m)
}
func (m *ListEvaluatorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEvaluatorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListEvaluatorsRequest proto.InternalMessageInfo

func (m *ListEvaluatorsRequest) GetEvaluator() *OpEvaluator {
	if m != nil {
		return m.Evaluator
	}
	return nil
}

type ListEvaluatorsResponse struct {
	Evaluators           []*Evaluator `protobuf:"bytes,1,rep,name=evaluators,proto3" json:"evaluators,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListEvaluatorsResponse) Reset()         { *m = ListEvaluatorsResponse{} }
func (m *ListEvaluatorsResponse) String() string { return proto.CompactTextString(m) }
func (*ListEvaluatorsResponse) ProtoMessage()    {}
func (*ListEvaluatorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_465e3bbdbf652446, []int{1}
}

func (m *ListEvaluatorsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEvaluatorsResponse.Unmarshal(m, b)
}
func (m *ListEvaluatorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEvaluatorsResponse.Marshal(b, m, deterministic)
}
func (m *ListEvaluatorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEvaluatorsResponse.Merge(m, src)
}
func (m *ListEvaluatorsResponse) XXX_Size() int {
	return xxx_messageInfo_ListEvaluatorsResponse.Size(m)
}
func (m *ListEvaluatorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEvaluatorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListEvaluatorsResponse proto.InternalMessageInfo

func (m *ListEvaluatorsResponse) GetEvaluators() []*Evaluator {
	if m != nil {
		return m.Evaluators
	}
	return nil
}

func init() {
	proto.RegisterType((*ListEvaluatorsRequest)(nil), "ai.metathings.service.evaluatord.ListEvaluatorsRequest")
	proto.RegisterType((*ListEvaluatorsResponse)(nil), "ai.metathings.service.evaluatord.ListEvaluatorsResponse")
}

func init() { proto.RegisterFile("list_evaluators.proto", fileDescriptor_465e3bbdbf652446) }

var fileDescriptor_465e3bbdbf652446 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0xc9, 0x2c, 0x2e,
	0x89, 0x4f, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x2c, 0xc9, 0x2f, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x52, 0x48, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f,
	0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x83, 0xab, 0x4b, 0x91, 0x12, 0x2f, 0x4b,
	0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0x20, 0x5a, 0xa5, 0xb8, 0x73, 0xf3, 0x53,
	0x52, 0x73, 0x20, 0x1c, 0xa5, 0x3c, 0x2e, 0x51, 0x9f, 0xcc, 0xe2, 0x12, 0x57, 0xb8, 0xf9, 0x41,
	0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0xa1, 0x5c, 0x9c, 0x70, 0xc3, 0x24, 0x18, 0x15, 0x18,
	0x35, 0xb8, 0x8d, 0x74, 0xf5, 0x08, 0x59, 0xaa, 0xe7, 0x5f, 0x00, 0x37, 0xc9, 0x89, 0xe3, 0x97,
	0x13, 0x6b, 0x17, 0x23, 0x93, 0x00, 0x63, 0x10, 0xc2, 0x24, 0xa5, 0x54, 0x2e, 0x31, 0x74, 0xfb,
	0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xbc, 0xb9, 0xb8, 0x10, 0xbe, 0x94, 0x60, 0x54, 0x60,
	0xd6, 0xe0, 0x36, 0xd2, 0x26, 0x6c, 0x23, 0xdc, 0xa4, 0x20, 0x24, 0xed, 0x4e, 0x3c, 0x51, 0x08,
	0x5e, 0x4a, 0x12, 0x1b, 0xd8, 0xaf, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x31, 0x55, 0x31,
	0x6e, 0x4c, 0x01, 0x00, 0x00,
}
