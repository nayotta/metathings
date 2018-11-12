// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get.proto

package sensor

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetRequest struct {
	Name                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

type GetResponse struct {
	Sensor               *Sensor  `protobuf:"bytes,1,opt,name=sensor,proto3" json:"sensor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{1}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetSensor() *Sensor {
	if m != nil {
		return m.Sensor
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "ai.metathings.service.sensor.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "ai.metathings.service.sensor.GetResponse")
}

func init() { proto.RegisterFile("get.proto", fileDescriptor_21b2a6be0e6d8388) }

var fileDescriptor_21b2a6be0e6d8388 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x85, 0x59, 0x91, 0x05, 0x73, 0x56, 0x5b, 0xc9, 0x71, 0xe8, 0x71, 0x58, 0xd8, 0xdc, 0x04,
	0x14, 0xc4, 0xc2, 0xca, 0x42, 0x0b, 0xbb, 0x3d, 0xb0, 0xcf, 0x9e, 0x63, 0x2e, 0xb8, 0x9b, 0x89,
	0x99, 0xc9, 0xee, 0xcf, 0x15, 0xfc, 0x25, 0x42, 0xb2, 0x5a, 0x5e, 0x35, 0x33, 0xcc, 0x7c, 0xef,
	0xbd, 0x51, 0x67, 0x16, 0x05, 0x42, 0x24, 0xa1, 0x66, 0x65, 0x1c, 0x0c, 0x28, 0x46, 0x0e, 0xce,
	0x5b, 0x06, 0xc6, 0x38, 0xba, 0x3d, 0x02, 0xa3, 0x67, 0x8a, 0xcb, 0x4b, 0x4b, 0x64, 0x7b, 0xd4,
	0xf9, 0xb6, 0x4b, 0x1f, 0x7a, 0x8a, 0x26, 0x04, 0x8c, 0x5c, 0xe8, 0xe5, 0xbd, 0x75, 0x72, 0x48,
	0x1d, 0xec, 0x69, 0xd0, 0xc3, 0xe4, 0xe4, 0x93, 0x26, 0x6d, 0x69, 0x9b, 0x97, 0xdb, 0xd1, 0xf4,
	0xee, 0xdd, 0x08, 0x45, 0xd6, 0xff, 0xed, 0xcc, 0x9d, 0x17, 0xfd, 0x32, 0x6d, 0x9e, 0x95, 0x7a,
	0x41, 0x69, 0xf1, 0x2b, 0x21, 0x4b, 0xf3, 0xa0, 0x4e, 0xbd, 0x19, 0xf0, 0xa2, 0x5a, 0x57, 0x37,
	0x8b, 0xdb, 0x15, 0x94, 0x08, 0xf0, 0x17, 0x01, 0x76, 0x12, 0x9d, 0xb7, 0x6f, 0xa6, 0x4f, 0xf8,
	0x54, 0xff, 0x7c, 0x5f, 0x9d, 0xac, 0xab, 0x36, 0x13, 0x9b, 0x57, 0xb5, 0xc8, 0x3a, 0x1c, 0xc8,
	0x33, 0x36, 0x8f, 0xaa, 0x2e, 0x36, 0xb3, 0xd4, 0x35, 0x1c, 0xfb, 0x15, 0x76, 0xb9, 0xb4, 0x33,
	0xd3, 0xd5, 0xd9, 0xf0, 0xee, 0x37, 0x00, 0x00, 0xff, 0xff, 0x02, 0xa7, 0xca, 0xbe, 0x2c, 0x01,
	0x00, 0x00,
}
