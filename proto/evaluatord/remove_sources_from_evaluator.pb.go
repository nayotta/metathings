// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remove_sources_from_evaluator.proto

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
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x4a, 0xcd, 0xcd,
	0x2f, 0x4b, 0x8d, 0x2f, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0x2d, 0x8e, 0x4f, 0x2b, 0xca, 0xcf, 0x8d,
	0x4f, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x52, 0x48, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x83, 0x2b, 0x4b, 0x91, 0x12, 0x2f, 0x4b, 0xcc, 0xc9,
	0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0x20, 0x5a, 0xa5, 0xb8, 0x73, 0xf3, 0x53, 0x52, 0x73,
	0x20, 0x1c, 0xa5, 0x43, 0x8c, 0x5c, 0x8a, 0x41, 0x60, 0xfb, 0x82, 0x21, 0xd6, 0xb9, 0x15, 0xe5,
	0xe7, 0xba, 0xc2, 0x4c, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x72, 0xe3, 0x62, 0x87,
	0xba, 0x46, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x47, 0x8f, 0x90, 0xfd, 0x7a, 0xfe, 0x05,
	0x41, 0xa9, 0x10, 0x4d, 0x41, 0x30, 0xcd, 0x42, 0xa1, 0x5c, 0x9c, 0x70, 0x15, 0x12, 0x4c, 0x0a,
	0x8c, 0x1a, 0xdc, 0x46, 0xba, 0xc4, 0x98, 0x04, 0x77, 0x90, 0x13, 0xc7, 0x2f, 0x27, 0xd6, 0x2e,
	0x46, 0x26, 0x01, 0xc6, 0x20, 0x84, 0x49, 0x4e, 0x3c, 0x51, 0x5c, 0x08, 0xe5, 0x49, 0x6c, 0x60,
	0x9f, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x68, 0xc1, 0x9d, 0x75, 0x48, 0x01, 0x00, 0x00,
}