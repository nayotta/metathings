// Code generated by protoc-gen-go. DO NOT EDIT.
// source: show.proto

package switcher

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ShowResponse struct {
	Switcher             *Switcher `protobuf:"bytes,1,opt,name=switcher" json:"switcher,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ShowResponse) Reset()         { *m = ShowResponse{} }
func (m *ShowResponse) String() string { return proto.CompactTextString(m) }
func (*ShowResponse) ProtoMessage()    {}
func (*ShowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_show_28f2d19110d4bd6e, []int{0}
}
func (m *ShowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowResponse.Unmarshal(m, b)
}
func (m *ShowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowResponse.Marshal(b, m, deterministic)
}
func (dst *ShowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowResponse.Merge(dst, src)
}
func (m *ShowResponse) XXX_Size() int {
	return xxx_messageInfo_ShowResponse.Size(m)
}
func (m *ShowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShowResponse proto.InternalMessageInfo

func (m *ShowResponse) GetSwitcher() *Switcher {
	if m != nil {
		return m.Switcher
	}
	return nil
}

func init() {
	proto.RegisterType((*ShowResponse)(nil), "ai.metathings.service.switcher.ShowResponse")
}

func init() { proto.RegisterFile("show.proto", fileDescriptor_show_28f2d19110d4bd6e) }

var fileDescriptor_show_28f2d19110d4bd6e = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xce, 0xc8, 0x2f,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4b, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c,
	0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x2e,
	0xcf, 0x2c, 0x49, 0xce, 0x48, 0x2d, 0x92, 0xe2, 0x83, 0xb1, 0x20, 0xea, 0x95, 0x42, 0xb8, 0x78,
	0x82, 0x33, 0xf2, 0xcb, 0x83, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x5c, 0xb8, 0x38,
	0x60, 0x2a, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x34, 0xf4, 0xf0, 0x1b, 0xa9, 0x17, 0x0c,
	0x65, 0x04, 0xc1, 0x75, 0x26, 0xb1, 0x81, 0x0d, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb7,
	0x41, 0xb3, 0xfe, 0x9a, 0x00, 0x00, 0x00,
}
