// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

package datad

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type Bundle struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Flows                []*Flow  `protobuf:"bytes,2,rep,name=flows,proto3" json:"flows,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bundle) Reset()         { *m = Bundle{} }
func (m *Bundle) String() string { return proto.CompactTextString(m) }
func (*Bundle) ProtoMessage()    {}
func (*Bundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{0}
}

func (m *Bundle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bundle.Unmarshal(m, b)
}
func (m *Bundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bundle.Marshal(b, m, deterministic)
}
func (m *Bundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bundle.Merge(m, src)
}
func (m *Bundle) XXX_Size() int {
	return xxx_messageInfo_Bundle.Size(m)
}
func (m *Bundle) XXX_DiscardUnknown() {
	xxx_messageInfo_Bundle.DiscardUnknown(m)
}

var xxx_messageInfo_Bundle proto.InternalMessageInfo

func (m *Bundle) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Bundle) GetFlows() []*Flow {
	if m != nil {
		return m.Flows
	}
	return nil
}

type OpBundle struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Flows                []*Flow               `protobuf:"bytes,2,rep,name=flows,proto3" json:"flows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OpBundle) Reset()         { *m = OpBundle{} }
func (m *OpBundle) String() string { return proto.CompactTextString(m) }
func (*OpBundle) ProtoMessage()    {}
func (*OpBundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{1}
}

func (m *OpBundle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpBundle.Unmarshal(m, b)
}
func (m *OpBundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpBundle.Marshal(b, m, deterministic)
}
func (m *OpBundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpBundle.Merge(m, src)
}
func (m *OpBundle) XXX_Size() int {
	return xxx_messageInfo_OpBundle.Size(m)
}
func (m *OpBundle) XXX_DiscardUnknown() {
	xxx_messageInfo_OpBundle.DiscardUnknown(m)
}

var xxx_messageInfo_OpBundle proto.InternalMessageInfo

func (m *OpBundle) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *OpBundle) GetFlows() []*Flow {
	if m != nil {
		return m.Flows
	}
	return nil
}

type Flow struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Flow) Reset()         { *m = Flow{} }
func (m *Flow) String() string { return proto.CompactTextString(m) }
func (*Flow) ProtoMessage()    {}
func (*Flow) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{2}
}

func (m *Flow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Flow.Unmarshal(m, b)
}
func (m *Flow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Flow.Marshal(b, m, deterministic)
}
func (m *Flow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Flow.Merge(m, src)
}
func (m *Flow) XXX_Size() int {
	return xxx_messageInfo_Flow.Size(m)
}
func (m *Flow) XXX_DiscardUnknown() {
	xxx_messageInfo_Flow.DiscardUnknown(m)
}

var xxx_messageInfo_Flow proto.InternalMessageInfo

func (m *Flow) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Flow) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type OpFlow struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OpFlow) Reset()         { *m = OpFlow{} }
func (m *OpFlow) String() string { return proto.CompactTextString(m) }
func (*OpFlow) ProtoMessage()    {}
func (*OpFlow) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{3}
}

func (m *OpFlow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpFlow.Unmarshal(m, b)
}
func (m *OpFlow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpFlow.Marshal(b, m, deterministic)
}
func (m *OpFlow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpFlow.Merge(m, src)
}
func (m *OpFlow) XXX_Size() int {
	return xxx_messageInfo_OpFlow.Size(m)
}
func (m *OpFlow) XXX_DiscardUnknown() {
	xxx_messageInfo_OpFlow.DiscardUnknown(m)
}

var xxx_messageInfo_OpFlow proto.InternalMessageInfo

func (m *OpFlow) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *OpFlow) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func init() {
	proto.RegisterType((*Bundle)(nil), "ai.metathings.service.datad.Bundle")
	proto.RegisterType((*OpBundle)(nil), "ai.metathings.service.datad.OpBundle")
	proto.RegisterType((*Flow)(nil), "ai.metathings.service.datad.Flow")
	proto.RegisterType((*OpFlow)(nil), "ai.metathings.service.datad.OpFlow")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_4c16552f9fdb66d8) }

var fileDescriptor_4c16552f9fdb66d8 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xcd, 0x4f, 0x49,
	0xcd, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4e, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49,
	0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4b,
	0x49, 0x2c, 0x49, 0x4c, 0x91, 0x92, 0x4b, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x2b, 0x4d,
	0x2a, 0x4d, 0xd3, 0x2f, 0x2f, 0x4a, 0x2c, 0x28, 0x48, 0x2d, 0x2a, 0x86, 0x68, 0x56, 0x0a, 0xe4,
	0x62, 0x73, 0x2a, 0xcd, 0x4b, 0xc9, 0x49, 0x15, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0x32, 0xe7, 0x62, 0x4d, 0xcb, 0xc9, 0x2f, 0x2f,
	0x96, 0x60, 0x52, 0x60, 0xd6, 0xe0, 0x36, 0x52, 0xd4, 0xc3, 0x63, 0x8d, 0x9e, 0x5b, 0x4e, 0x7e,
	0x79, 0x10, 0x44, 0xbd, 0x52, 0x21, 0x17, 0x87, 0x7f, 0x01, 0xd4, 0x50, 0x1d, 0xb8, 0xa1, 0xdc,
	0x46, 0x32, 0x7a, 0x10, 0xb7, 0xe8, 0xc1, 0xdc, 0xa2, 0x17, 0x5c, 0x52, 0x94, 0x99, 0x97, 0x1e,
	0x96, 0x98, 0x53, 0x9a, 0x4a, 0x99, 0x95, 0x5a, 0x5c, 0x2c, 0x20, 0x2e, 0x86, 0x1f, 0x84, 0xb8,
	0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x98, 0xc0, 0x22, 0x60, 0xb6, 0x52, 0x06, 0x17, 0x9b, 0x7f,
	0x01, 0x58, 0x35, 0x69, 0x8e, 0x33, 0x40, 0x32, 0x8b, 0x90, 0x7a, 0xb0, 0x4a, 0x27, 0xf6, 0x28,
	0x56, 0xb0, 0x53, 0x93, 0xd8, 0xc0, 0x8a, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xae,
	0x07, 0x7c, 0xb7, 0x01, 0x00, 0x00,
}
