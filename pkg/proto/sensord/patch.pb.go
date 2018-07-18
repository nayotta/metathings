// Code generated by protoc-gen-go. DO NOT EDIT.
// source: patch.proto

package sensord

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PatchRequest struct {
	Id                   *wrappers.StringValue   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name                 *wrappers.StringValue   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Tags                 []*wrappers.StringValue `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PatchRequest) Reset()         { *m = PatchRequest{} }
func (m *PatchRequest) String() string { return proto.CompactTextString(m) }
func (*PatchRequest) ProtoMessage()    {}
func (*PatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_patch_1c1e9c0c8695a46c, []int{0}
}
func (m *PatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchRequest.Unmarshal(m, b)
}
func (m *PatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchRequest.Marshal(b, m, deterministic)
}
func (dst *PatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchRequest.Merge(dst, src)
}
func (m *PatchRequest) XXX_Size() int {
	return xxx_messageInfo_PatchRequest.Size(m)
}
func (m *PatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PatchRequest proto.InternalMessageInfo

func (m *PatchRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PatchRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *PatchRequest) GetTags() []*wrappers.StringValue {
	if m != nil {
		return m.Tags
	}
	return nil
}

type PatchResponse struct {
	Sensor               *Sensor  `protobuf:"bytes,1,opt,name=sensor" json:"sensor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PatchResponse) Reset()         { *m = PatchResponse{} }
func (m *PatchResponse) String() string { return proto.CompactTextString(m) }
func (*PatchResponse) ProtoMessage()    {}
func (*PatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_patch_1c1e9c0c8695a46c, []int{1}
}
func (m *PatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchResponse.Unmarshal(m, b)
}
func (m *PatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchResponse.Marshal(b, m, deterministic)
}
func (dst *PatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchResponse.Merge(dst, src)
}
func (m *PatchResponse) XXX_Size() int {
	return xxx_messageInfo_PatchResponse.Size(m)
}
func (m *PatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PatchResponse proto.InternalMessageInfo

func (m *PatchResponse) GetSensor() *Sensor {
	if m != nil {
		return m.Sensor
	}
	return nil
}

func init() {
	proto.RegisterType((*PatchRequest)(nil), "ai.metathings.service.sensord.PatchRequest")
	proto.RegisterType((*PatchResponse)(nil), "ai.metathings.service.sensord.PatchResponse")
}

func init() { proto.RegisterFile("patch.proto", fileDescriptor_patch_1c1e9c0c8695a46c) }

var fileDescriptor_patch_1c1e9c0c8695a46c = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x49, 0x2a, 0x3d, 0x6c, 0xeb, 0x25, 0xa7, 0x50, 0xfc, 0x13, 0x0a, 0x82, 0x97, 0x4e,
	0x44, 0xc5, 0x9b, 0x17, 0x1f, 0x40, 0x24, 0x05, 0xef, 0x9b, 0x64, 0xdc, 0x2c, 0x26, 0x99, 0xb8,
	0x33, 0x69, 0x5e, 0xc9, 0xb7, 0x12, 0x7c, 0x12, 0x49, 0x36, 0x7a, 0xb4, 0xa7, 0xdd, 0x65, 0xbf,
	0x1f, 0xdf, 0x6f, 0x46, 0xad, 0x3a, 0x2d, 0x45, 0x05, 0x9d, 0x23, 0xa1, 0xe8, 0x5c, 0x5b, 0x68,
	0x50, 0xb4, 0x54, 0xb6, 0x35, 0x0c, 0x8c, 0xee, 0x60, 0x0b, 0x04, 0xc6, 0x96, 0xc9, 0x95, 0x9b,
	0x0b, 0x43, 0x64, 0x6a, 0x4c, 0xa7, 0x70, 0xde, 0xbf, 0xa5, 0x83, 0xd3, 0x5d, 0x87, 0x8e, 0x3d,
	0xbe, 0x79, 0x30, 0x56, 0xaa, 0x3e, 0x87, 0x82, 0x9a, 0xb4, 0x19, 0xac, 0xbc, 0xd3, 0x90, 0x1a,
	0xda, 0x4d, 0x9f, 0xbb, 0x83, 0xae, 0x6d, 0xa9, 0x85, 0x1c, 0xa7, 0x7f, 0xd7, 0x99, 0x5b, 0xfb,
	0x02, 0xff, 0xda, 0x7e, 0x06, 0x6a, 0xfd, 0x32, 0x4a, 0x65, 0xf8, 0xd1, 0x23, 0x4b, 0x74, 0xaf,
	0x42, 0x5b, 0xc6, 0x41, 0x12, 0x5c, 0xaf, 0x6e, 0xcf, 0xc0, 0x3b, 0xc0, 0xaf, 0x03, 0xec, 0xc5,
	0xd9, 0xd6, 0xbc, 0xea, 0xba, 0xc7, 0xa7, 0xe5, 0xf7, 0xd7, 0x65, 0x98, 0x04, 0x59, 0x68, 0xcb,
	0xe8, 0x46, 0x9d, 0xb4, 0xba, 0xc1, 0x38, 0x3c, 0xce, 0x65, 0x53, 0x72, 0x24, 0x44, 0x1b, 0x8e,
	0x17, 0xc9, 0xe2, 0x38, 0x31, 0x26, 0xb7, 0xcf, 0xea, 0x74, 0x36, 0xe5, 0x8e, 0x5a, 0xc6, 0xe8,
	0x51, 0x2d, 0xfd, 0x2c, 0xb3, 0xee, 0x15, 0xfc, 0xbb, 0x51, 0xd8, 0x4f, 0x67, 0x36, 0x43, 0xf9,
	0x72, 0xea, 0xba, 0xfb, 0x09, 0x00, 0x00, 0xff, 0xff, 0x15, 0x15, 0x35, 0x08, 0x95, 0x01, 0x00,
	0x00,
}
