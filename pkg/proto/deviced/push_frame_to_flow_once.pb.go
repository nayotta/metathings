// Code generated by protoc-gen-go. DO NOT EDIT.
// source: push_frame_to_flow_once.proto

package deviced

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PushFrameToFlowOnceRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Device               *OpDevice             `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
	Frame                *OpFrame              `protobuf:"bytes,3,opt,name=frame,proto3" json:"frame,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PushFrameToFlowOnceRequest) Reset()         { *m = PushFrameToFlowOnceRequest{} }
func (m *PushFrameToFlowOnceRequest) String() string { return proto.CompactTextString(m) }
func (*PushFrameToFlowOnceRequest) ProtoMessage()    {}
func (*PushFrameToFlowOnceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3372753545565c1, []int{0}
}

func (m *PushFrameToFlowOnceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushFrameToFlowOnceRequest.Unmarshal(m, b)
}
func (m *PushFrameToFlowOnceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushFrameToFlowOnceRequest.Marshal(b, m, deterministic)
}
func (m *PushFrameToFlowOnceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushFrameToFlowOnceRequest.Merge(m, src)
}
func (m *PushFrameToFlowOnceRequest) XXX_Size() int {
	return xxx_messageInfo_PushFrameToFlowOnceRequest.Size(m)
}
func (m *PushFrameToFlowOnceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushFrameToFlowOnceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushFrameToFlowOnceRequest proto.InternalMessageInfo

func (m *PushFrameToFlowOnceRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PushFrameToFlowOnceRequest) GetDevice() *OpDevice {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *PushFrameToFlowOnceRequest) GetFrame() *OpFrame {
	if m != nil {
		return m.Frame
	}
	return nil
}

func init() {
	proto.RegisterType((*PushFrameToFlowOnceRequest)(nil), "ai.metathings.service.deviced.PushFrameToFlowOnceRequest")
}

func init() { proto.RegisterFile("push_frame_to_flow_once.proto", fileDescriptor_e3372753545565c1) }

var fileDescriptor_e3372753545565c1 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0x69, 0x65, 0x46, 0xcd, 0x6c, 0xa4, 0x1b, 0x4b, 0x71, 0x44, 0x5c, 0xa8, 0xab, 0x14,
	0x14, 0x7c, 0x80, 0x22, 0x05, 0x57, 0x23, 0xa3, 0xb8, 0x70, 0x53, 0x32, 0xcd, 0x6d, 0x1b, 0x48,
	0x7b, 0x63, 0x7e, 0xa6, 0xef, 0xe0, 0x8b, 0xfa, 0x0e, 0xae, 0x64, 0x92, 0x0e, 0xcc, 0x4a, 0x57,
	0xb9, 0x21, 0xf7, 0x7c, 0xe7, 0x9c, 0x90, 0xa5, 0x72, 0xa6, 0xab, 0x1a, 0xcd, 0x7a, 0xa8, 0x2c,
	0x56, 0x8d, 0xc4, 0xb1, 0xc2, 0xa1, 0x06, 0xaa, 0x34, 0x5a, 0x4c, 0x96, 0x4c, 0xd0, 0x1e, 0x2c,
	0xb3, 0x9d, 0x18, 0x5a, 0x43, 0x0d, 0xe8, 0xad, 0xa8, 0x81, 0x72, 0xd8, 0x1d, 0x3c, 0xbb, 0x6c,
	0x11, 0x5b, 0x09, 0xb9, 0x5f, 0xde, 0xb8, 0x26, 0x1f, 0x35, 0x53, 0x0a, 0xb4, 0x09, 0xf2, 0xec,
	0x7c, 0xcb, 0xa4, 0xe0, 0xcc, 0x42, 0xbe, 0x1f, 0xa6, 0x87, 0x45, 0x8f, 0x1c, 0x64, 0xb8, 0x5c,
	0x7f, 0x47, 0x24, 0x7b, 0x71, 0xa6, 0x2b, 0x77, 0x29, 0xde, 0xb0, 0x94, 0x38, 0xae, 0x86, 0x1a,
	0xd6, 0xf0, 0xe9, 0xc0, 0xd8, 0xe4, 0x91, 0xc4, 0x82, 0xa7, 0xd1, 0x55, 0x74, 0xb7, 0xb8, 0xbf,
	0xa0, 0xc1, 0x91, 0xee, 0x1d, 0xe9, 0xab, 0xd5, 0x62, 0x68, 0xdf, 0x99, 0x74, 0x50, 0x9c, 0xfc,
	0x14, 0xb3, 0xaf, 0x28, 0x3e, 0x8b, 0xd6, 0xb1, 0xe0, 0xc9, 0x33, 0x99, 0x87, 0x9c, 0x69, 0xec,
	0xb5, 0xb7, 0xf4, 0xcf, 0x32, 0x74, 0xa5, 0x9e, 0xfc, 0x74, 0x80, 0x99, 0x00, 0x49, 0x49, 0x66,
	0xfe, 0x8b, 0xd2, 0x23, 0x4f, 0xba, 0xf9, 0x97, 0xe4, 0xab, 0x1c, 0x80, 0x82, 0xbc, 0x38, 0xfd,
	0x38, 0x9e, 0x76, 0x36, 0x73, 0xdf, 0xe0, 0xe1, 0x37, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x39, 0xaf,
	0x74, 0x81, 0x01, 0x00, 0x00,
}
