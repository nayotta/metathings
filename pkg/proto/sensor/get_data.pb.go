// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_data.proto

package sensor

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

type GetDataRequest struct {
	Name                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetDataRequest) Reset()         { *m = GetDataRequest{} }
func (m *GetDataRequest) String() string { return proto.CompactTextString(m) }
func (*GetDataRequest) ProtoMessage()    {}
func (*GetDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_get_data_8912bf3d33fbc83c, []int{0}
}
func (m *GetDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataRequest.Unmarshal(m, b)
}
func (m *GetDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataRequest.Marshal(b, m, deterministic)
}
func (dst *GetDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataRequest.Merge(dst, src)
}
func (m *GetDataRequest) XXX_Size() int {
	return xxx_messageInfo_GetDataRequest.Size(m)
}
func (m *GetDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataRequest proto.InternalMessageInfo

func (m *GetDataRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

type GetDataResponse struct {
	Data                 *SensorData `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetDataResponse) Reset()         { *m = GetDataResponse{} }
func (m *GetDataResponse) String() string { return proto.CompactTextString(m) }
func (*GetDataResponse) ProtoMessage()    {}
func (*GetDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_get_data_8912bf3d33fbc83c, []int{1}
}
func (m *GetDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataResponse.Unmarshal(m, b)
}
func (m *GetDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataResponse.Marshal(b, m, deterministic)
}
func (dst *GetDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataResponse.Merge(dst, src)
}
func (m *GetDataResponse) XXX_Size() int {
	return xxx_messageInfo_GetDataResponse.Size(m)
}
func (m *GetDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataResponse proto.InternalMessageInfo

func (m *GetDataResponse) GetData() *SensorData {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetDataRequest)(nil), "ai.metathings.service.sensor.GetDataRequest")
	proto.RegisterType((*GetDataResponse)(nil), "ai.metathings.service.sensor.GetDataResponse")
}

func init() { proto.RegisterFile("get_data.proto", fileDescriptor_get_data_8912bf3d33fbc83c) }

var fileDescriptor_get_data_8912bf3d33fbc83c = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x41, 0x4b, 0x03, 0x41,
	0x0c, 0x85, 0x59, 0x29, 0x3d, 0x8c, 0x52, 0x61, 0x4f, 0x52, 0x8a, 0x96, 0x9e, 0x7a, 0x69, 0x06,
	0x14, 0xc4, 0x83, 0x27, 0x11, 0x04, 0x2f, 0x42, 0x0b, 0x5e, 0x25, 0xdb, 0xc6, 0xe9, 0xe0, 0xee,
	0x64, 0x9d, 0x64, 0xbb, 0x3f, 0x57, 0xf0, 0x97, 0x88, 0x33, 0x75, 0x8f, 0x9e, 0x92, 0x90, 0x7c,
	0x2f, 0xef, 0x99, 0x89, 0x23, 0x7d, 0xdb, 0xa1, 0x22, 0xb4, 0x91, 0x95, 0xcb, 0x19, 0x7a, 0x68,
	0x48, 0x51, 0xf7, 0x3e, 0x38, 0x01, 0xa1, 0x78, 0xf0, 0x5b, 0x02, 0xa1, 0x20, 0x1c, 0xa7, 0x97,
	0x8e, 0xd9, 0xd5, 0x64, 0xd3, 0x6d, 0xd5, 0xbd, 0xdb, 0x3e, 0x62, 0xdb, 0x52, 0x94, 0x4c, 0x4f,
	0x6f, 0x9d, 0xd7, 0x7d, 0x57, 0xc1, 0x96, 0x1b, 0xdb, 0xf4, 0x5e, 0x3f, 0xb8, 0xb7, 0x8e, 0x57,
	0x69, 0xb9, 0x3a, 0x60, 0xed, 0x77, 0xa8, 0x1c, 0xc5, 0x0e, 0xed, 0x91, 0x3b, 0xcb, 0xfa, 0x79,
	0x5a, 0x3c, 0x9b, 0xc9, 0x13, 0xe9, 0x23, 0x2a, 0xae, 0xe9, 0xb3, 0x23, 0xd1, 0xf2, 0xce, 0x8c,
	0x02, 0x36, 0x74, 0x51, 0xcc, 0x8b, 0xe5, 0xe9, 0xf5, 0x0c, 0xb2, 0x0d, 0xf8, 0xb3, 0x01, 0x1b,
	0x8d, 0x3e, 0xb8, 0x57, 0xac, 0x3b, 0x7a, 0x18, 0x7f, 0x7f, 0x5d, 0x9d, 0xcc, 0x8b, 0x75, 0x22,
	0x16, 0x2f, 0xe6, 0x7c, 0xd0, 0x92, 0x96, 0x83, 0x50, 0x79, 0x6f, 0x46, 0xbf, 0x81, 0x8f, 0x62,
	0x4b, 0xf8, 0x2f, 0x31, 0x6c, 0x52, 0x49, 0x7c, 0xa2, 0xaa, 0x71, 0x7a, 0x7a, 0xf3, 0x13, 0x00,
	0x00, 0xff, 0xff, 0x6d, 0xad, 0x04, 0x8b, 0x39, 0x01, 0x00, 0x00,
}