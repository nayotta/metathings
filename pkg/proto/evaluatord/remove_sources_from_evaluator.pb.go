// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remove_sources_from_evaluator.proto

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

type RemoveSourcesFromEvaluatorRequest struct {
	Sources              []*OpResource `protobuf:"bytes,1,rep,name=sources,proto3" json:"sources,omitempty"`
	Evaluator            *OpEvaluator  `protobuf:"bytes,2,opt,name=evaluator,proto3" json:"evaluator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RemoveSourcesFromEvaluatorRequest) Reset()         { *m = RemoveSourcesFromEvaluatorRequest{} }
func (m *RemoveSourcesFromEvaluatorRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveSourcesFromEvaluatorRequest) ProtoMessage()    {}
func (*RemoveSourcesFromEvaluatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33efe53d7fd0aa6f, []int{0}
}

func (m *RemoveSourcesFromEvaluatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveSourcesFromEvaluatorRequest.Unmarshal(m, b)
}
func (m *RemoveSourcesFromEvaluatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveSourcesFromEvaluatorRequest.Marshal(b, m, deterministic)
}
func (m *RemoveSourcesFromEvaluatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveSourcesFromEvaluatorRequest.Merge(m, src)
}
func (m *RemoveSourcesFromEvaluatorRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveSourcesFromEvaluatorRequest.Size(m)
}
func (m *RemoveSourcesFromEvaluatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveSourcesFromEvaluatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveSourcesFromEvaluatorRequest proto.InternalMessageInfo

func (m *RemoveSourcesFromEvaluatorRequest) GetSources() []*OpResource {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *RemoveSourcesFromEvaluatorRequest) GetEvaluator() *OpEvaluator {
	if m != nil {
		return m.Evaluator
	}
	return nil
}

func init() {
	proto.RegisterType((*RemoveSourcesFromEvaluatorRequest)(nil), "ai.metathings.service.evaluatord.RemoveSourcesFromEvaluatorRequest")
}

func init() {
	proto.RegisterFile("remove_sources_from_evaluator.proto", fileDescriptor_33efe53d7fd0aa6f)
}

var fileDescriptor_33efe53d7fd0aa6f = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x4a, 0xcd, 0xcd,
	0x2f, 0x4b, 0x8d, 0x2f, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0x2d, 0x8e, 0x4f, 0x2b, 0xca, 0xcf, 0x8d,
	0x4f, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x52, 0x48, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x83, 0x2b, 0x4b, 0x91, 0x32, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x2d, 0xcf, 0x2c, 0xc9, 0xce, 0x2f, 0xd7, 0x4f,
	0xcf, 0xd7, 0x05, 0x6b, 0xd7, 0x2d, 0x4b, 0xcc, 0xc9, 0x4c, 0x01, 0xa9, 0x2b, 0xd6, 0x87, 0x33,
	0x21, 0x26, 0x4b, 0x71, 0xe7, 0xe6, 0xa7, 0xa4, 0xe6, 0x40, 0x38, 0x4a, 0x07, 0x18, 0xb9, 0x14,
	0x83, 0xc0, 0xce, 0x09, 0x86, 0xb8, 0xc6, 0xad, 0x28, 0x3f, 0xd7, 0x15, 0x66, 0x49, 0x50, 0x6a,
	0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x1b, 0x17, 0x3b, 0xd4, 0xb1, 0x12, 0x8c, 0x0a, 0xcc, 0x1a,
	0xdc, 0x46, 0x3a, 0x7a, 0x84, 0x9c, 0xa7, 0xe7, 0x5f, 0x10, 0x94, 0x0a, 0xd1, 0x14, 0x04, 0xd3,
	0x2c, 0x14, 0xcc, 0xc5, 0x09, 0x57, 0x21, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x6d, 0xa4, 0x4b, 0x8c,
	0x49, 0x70, 0x07, 0x39, 0xb1, 0x3d, 0xba, 0x2f, 0xcf, 0xa4, 0xc0, 0x18, 0x84, 0x30, 0xc7, 0x89,
	0x27, 0x8a, 0x0b, 0xa1, 0x38, 0x89, 0x0d, 0xec, 0x2f, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x9a, 0x78, 0xdc, 0x1e, 0x65, 0x01, 0x00, 0x00,
}
