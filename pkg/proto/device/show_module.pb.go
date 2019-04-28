// Code generated by protoc-gen-go. DO NOT EDIT.
// source: show_module.proto

package ai_metathings_service_device

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	deviced "github.com/nayotta/metathings/pkg/proto/deviced"
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

type ShowModuleResponse struct {
	Module               *deviced.Module `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ShowModuleResponse) Reset()         { *m = ShowModuleResponse{} }
func (m *ShowModuleResponse) String() string { return proto.CompactTextString(m) }
func (*ShowModuleResponse) ProtoMessage()    {}
func (*ShowModuleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_950c15bb855660dd, []int{0}
}

func (m *ShowModuleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowModuleResponse.Unmarshal(m, b)
}
func (m *ShowModuleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowModuleResponse.Marshal(b, m, deterministic)
}
func (m *ShowModuleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowModuleResponse.Merge(m, src)
}
func (m *ShowModuleResponse) XXX_Size() int {
	return xxx_messageInfo_ShowModuleResponse.Size(m)
}
func (m *ShowModuleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowModuleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShowModuleResponse proto.InternalMessageInfo

func (m *ShowModuleResponse) GetModule() *deviced.Module {
	if m != nil {
		return m.Module
	}
	return nil
}

func init() {
	proto.RegisterType((*ShowModuleResponse)(nil), "ai.metathings.service.device.ShowModuleResponse")
}

func init() { proto.RegisterFile("show_module.proto", fileDescriptor_950c15bb855660dd) }

var fileDescriptor_950c15bb855660dd = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xce, 0xc8, 0x2f,
	0x8f, 0xcf, 0xcd, 0x4f, 0x29, 0xcd, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x49,
	0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d,
	0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4b, 0x49, 0x05, 0x51, 0x52, 0xd6, 0xe9, 0x99, 0x25, 0x19, 0xa5,
	0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x79, 0x89, 0x95, 0xf9, 0x25, 0x25, 0x89, 0xfa, 0x08, 0xd5,
	0xfa, 0x05, 0xd9, 0xe9, 0xfa, 0x60, 0x63, 0xf4, 0x21, 0xea, 0x53, 0xf4, 0x73, 0xf3, 0x53, 0x52,
	0x73, 0x20, 0x46, 0x2b, 0x05, 0x73, 0x09, 0x05, 0x67, 0xe4, 0x97, 0xfb, 0x82, 0xad, 0x0b, 0x4a,
	0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0xb2, 0xe5, 0x62, 0x83, 0x38, 0x40, 0x82, 0x51, 0x81,
	0x51, 0x83, 0xdb, 0x48, 0x55, 0x0f, 0x9f, 0x0b, 0x52, 0xf4, 0xa0, 0xda, 0xa1, 0x9a, 0x92, 0xd8,
	0xc0, 0x66, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x34, 0xa9, 0xe1, 0x7e, 0xcb, 0x00, 0x00,
	0x00,
}
