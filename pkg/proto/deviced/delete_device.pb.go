// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delete_device.proto

package deviced

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

type DeleteDeviceRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DeleteDeviceRequest) Reset()         { *m = DeleteDeviceRequest{} }
func (m *DeleteDeviceRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDeviceRequest) ProtoMessage()    {}
func (*DeleteDeviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_delete_device_4dae2197e82dfd62, []int{0}
}
func (m *DeleteDeviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDeviceRequest.Unmarshal(m, b)
}
func (m *DeleteDeviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDeviceRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteDeviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDeviceRequest.Merge(dst, src)
}
func (m *DeleteDeviceRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteDeviceRequest.Size(m)
}
func (m *DeleteDeviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDeviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDeviceRequest proto.InternalMessageInfo

func (m *DeleteDeviceRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteDeviceRequest)(nil), "ai.metathings.service.deviced.DeleteDeviceRequest")
}

func init() { proto.RegisterFile("delete_device.proto", fileDescriptor_delete_device_4dae2197e82dfd62) }

var fileDescriptor_delete_device_4dae2197e82dfd62 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0xcd, 0x4a, 0xc5, 0x30,
	0x10, 0x46, 0x69, 0x17, 0x5d, 0xd4, 0x5d, 0xbb, 0x91, 0xe2, 0x4f, 0x71, 0xe5, 0xa6, 0x09, 0xa8,
	0xf8, 0x00, 0xd2, 0x9d, 0xbb, 0x0a, 0x6e, 0x25, 0x6d, 0xc6, 0x74, 0x30, 0xed, 0xd4, 0x64, 0xd2,
	0x3e, 0xae, 0xe0, 0x93, 0x5c, 0x48, 0xb8, 0x77, 0x37, 0xf0, 0x9d, 0xc3, 0x99, 0xb2, 0xd6, 0x60,
	0x81, 0xe1, 0x4b, 0xc3, 0x8e, 0x13, 0x88, 0xcd, 0x11, 0x53, 0x75, 0xab, 0x50, 0x2c, 0xc0, 0x8a,
	0x67, 0x5c, 0x8d, 0x17, 0x1e, 0x5c, 0x1c, 0x13, 0xa3, 0x9b, 0x3b, 0x43, 0x64, 0x2c, 0xc8, 0x08,
	0x8f, 0xe1, 0x5b, 0x1e, 0x4e, 0x6d, 0x1b, 0x38, 0x9f, 0xf4, 0xe6, 0xd5, 0x20, 0xcf, 0x61, 0x14,
	0x13, 0x2d, 0x72, 0x39, 0x90, 0x7f, 0xe8, 0x90, 0x86, 0xba, 0x38, 0x76, 0xbb, 0xb2, 0xa8, 0x15,
	0x93, 0xf3, 0xf2, 0x72, 0x26, 0xef, 0xe1, 0xbd, 0xac, 0xfb, 0xf8, 0x4d, 0x1f, 0x43, 0x03, 0xfc,
	0x06, 0xf0, 0x5c, 0xbd, 0x94, 0x39, 0xea, 0xeb, 0xac, 0xcd, 0x1e, 0xaf, 0x9e, 0x6e, 0x44, 0x6a,
	0x8b, 0x73, 0x5b, 0x7c, 0xb0, 0xc3, 0xd5, 0x7c, 0x2a, 0x1b, 0xe0, 0xad, 0xf8, 0xff, 0xbb, 0xcf,
	0xdb, 0x6c, 0xc8, 0x51, 0x8f, 0x45, 0x24, 0x9e, 0x4f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x99, 0x72,
	0x40, 0x7d, 0xe1, 0x00, 0x00, 0x00,
}