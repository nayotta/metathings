// Code generated by protoc-gen-go. DO NOT EDIT.
// source: switcher.proto

package switcher

import (
	fmt "fmt"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SwitcherState int32

const (
	SwitcherState_SWITCHER_STATE_UNKNOWN SwitcherState = 0
	SwitcherState_SWITCHER_STATE_ON      SwitcherState = 1
	SwitcherState_SWITCHER_STATE_OFF     SwitcherState = 2
)

var SwitcherState_name = map[int32]string{
	0: "SWITCHER_STATE_UNKNOWN",
	1: "SWITCHER_STATE_ON",
	2: "SWITCHER_STATE_OFF",
}

var SwitcherState_value = map[string]int32{
	"SWITCHER_STATE_UNKNOWN": 0,
	"SWITCHER_STATE_ON":      1,
	"SWITCHER_STATE_OFF":     2,
}

func (x SwitcherState) String() string {
	return proto.EnumName(SwitcherState_name, int32(x))
}

func (SwitcherState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fac3a916beaffe96, []int{0}
}

type Switcher struct {
	State                SwitcherState `protobuf:"varint,1,opt,name=state,proto3,enum=ai.metathings.service.switcher.SwitcherState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Switcher) Reset()         { *m = Switcher{} }
func (m *Switcher) String() string { return proto.CompactTextString(m) }
func (*Switcher) ProtoMessage()    {}
func (*Switcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_fac3a916beaffe96, []int{0}
}

func (m *Switcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Switcher.Unmarshal(m, b)
}
func (m *Switcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Switcher.Marshal(b, m, deterministic)
}
func (m *Switcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Switcher.Merge(m, src)
}
func (m *Switcher) XXX_Size() int {
	return xxx_messageInfo_Switcher.Size(m)
}
func (m *Switcher) XXX_DiscardUnknown() {
	xxx_messageInfo_Switcher.DiscardUnknown(m)
}

var xxx_messageInfo_Switcher proto.InternalMessageInfo

func (m *Switcher) GetState() SwitcherState {
	if m != nil {
		return m.State
	}
	return SwitcherState_SWITCHER_STATE_UNKNOWN
}

func init() {
	proto.RegisterEnum("ai.metathings.service.switcher.SwitcherState", SwitcherState_name, SwitcherState_value)
	proto.RegisterType((*Switcher)(nil), "ai.metathings.service.switcher.Switcher")
}

func init() { proto.RegisterFile("switcher.proto", fileDescriptor_fac3a916beaffe96) }

var fileDescriptor_fac3a916beaffe96 = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x2e, 0xcf, 0x2c,
	0x49, 0xce, 0x48, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4b, 0xcc, 0xd4, 0xcb,
	0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c,
	0x4e, 0xd5, 0x83, 0xa9, 0x52, 0xf2, 0xe7, 0xe2, 0x08, 0x86, 0xb2, 0x85, 0x9c, 0xb9, 0x58, 0x8b,
	0x4b, 0x12, 0x4b, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0xf8, 0x8c, 0x74, 0xf5, 0xf0, 0xeb, 0xd5,
	0x83, 0x69, 0x0c, 0x06, 0x69, 0x0a, 0x82, 0xe8, 0xd5, 0x8a, 0xe2, 0xe2, 0x45, 0x11, 0x17, 0x92,
	0xe2, 0x12, 0x0b, 0x0e, 0xf7, 0x0c, 0x71, 0xf6, 0x70, 0x0d, 0x8a, 0x0f, 0x0e, 0x71, 0x0c, 0x71,
	0x8d, 0x0f, 0xf5, 0xf3, 0xf6, 0xf3, 0x0f, 0xf7, 0x13, 0x60, 0x10, 0x12, 0xe5, 0x12, 0x44, 0x93,
	0xf3, 0xf7, 0x13, 0x60, 0x14, 0x12, 0xe3, 0x12, 0x42, 0x17, 0x76, 0x73, 0x13, 0x60, 0x4a, 0x62,
	0x03, 0xfb, 0xc9, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x82, 0x77, 0xb0, 0xe5, 0x00, 0x00,
	0x00,
}
