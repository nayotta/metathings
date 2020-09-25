// Code generated by protoc-gen-go. DO NOT EDIT.
// source: show_module_firmware_descriptor.proto

package device

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	deviced "github.com/nayotta/metathings/proto/deviced"
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

type ShowModuleFirmwareDescriptorResponse struct {
	FirmwareDescriptor   *deviced.FirmwareDescriptor `protobuf:"bytes,1,opt,name=firmware_descriptor,json=firmwareDescriptor,proto3" json:"firmware_descriptor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ShowModuleFirmwareDescriptorResponse) Reset()         { *m = ShowModuleFirmwareDescriptorResponse{} }
func (m *ShowModuleFirmwareDescriptorResponse) String() string { return proto.CompactTextString(m) }
func (*ShowModuleFirmwareDescriptorResponse) ProtoMessage()    {}
func (*ShowModuleFirmwareDescriptorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d5c1470cee9d487, []int{0}
}

func (m *ShowModuleFirmwareDescriptorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowModuleFirmwareDescriptorResponse.Unmarshal(m, b)
}
func (m *ShowModuleFirmwareDescriptorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowModuleFirmwareDescriptorResponse.Marshal(b, m, deterministic)
}
func (m *ShowModuleFirmwareDescriptorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowModuleFirmwareDescriptorResponse.Merge(m, src)
}
func (m *ShowModuleFirmwareDescriptorResponse) XXX_Size() int {
	return xxx_messageInfo_ShowModuleFirmwareDescriptorResponse.Size(m)
}
func (m *ShowModuleFirmwareDescriptorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowModuleFirmwareDescriptorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShowModuleFirmwareDescriptorResponse proto.InternalMessageInfo

func (m *ShowModuleFirmwareDescriptorResponse) GetFirmwareDescriptor() *deviced.FirmwareDescriptor {
	if m != nil {
		return m.FirmwareDescriptor
	}
	return nil
}

func init() {
	proto.RegisterType((*ShowModuleFirmwareDescriptorResponse)(nil), "ai.metathings.service.device.ShowModuleFirmwareDescriptorResponse")
}

func init() {
	proto.RegisterFile("show_module_firmware_descriptor.proto", fileDescriptor_0d5c1470cee9d487)
}

var fileDescriptor_0d5c1470cee9d487 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0xce, 0xc8, 0x2f,
	0x8f, 0xcf, 0xcd, 0x4f, 0x29, 0xcd, 0x49, 0x8d, 0x4f, 0xcb, 0x2c, 0xca, 0x2d, 0x4f, 0x2c, 0x4a,
	0x8d, 0x4f, 0x49, 0x2d, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0x49, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f,
	0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4b, 0x49, 0x05, 0x51, 0x52, 0xe6, 0xe9,
	0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x79, 0x89, 0x95, 0xf9, 0x25, 0x25,
	0x89, 0xfa, 0x08, 0xd5, 0xfa, 0x60, 0x23, 0xf4, 0x21, 0x6a, 0x53, 0xf4, 0x73, 0xf3, 0x53, 0x52,
	0x73, 0x20, 0xc6, 0x2a, 0x75, 0x31, 0x72, 0xa9, 0x04, 0x67, 0xe4, 0x97, 0xfb, 0x82, 0xed, 0x77,
	0x83, 0x5a, 0xef, 0x02, 0xb7, 0x3d, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x28, 0x89,
	0x4b, 0x18, 0x8b, 0xe3, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x0c, 0xf5, 0xf0, 0xb9, 0x2e,
	0x45, 0x0f, 0x8b, 0xb9, 0x42, 0x69, 0x18, 0x62, 0x4e, 0x1c, 0x51, 0x6c, 0x10, 0x1d, 0x49, 0x6c,
	0x60, 0xd7, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x96, 0x21, 0x9f, 0x1d, 0x01, 0x00,
	0x00,
}
