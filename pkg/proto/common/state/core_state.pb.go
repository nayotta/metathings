// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core_state.proto

package ai_metathings_state

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

type CoreState int32

const (
	CoreState_CORE_STATE_UNKNOWN CoreState = 0
	CoreState_CORE_STATE_ONLINE  CoreState = 1
	CoreState_CORE_STATE_OFFLINE CoreState = 2
)

var CoreState_name = map[int32]string{
	0: "CORE_STATE_UNKNOWN",
	1: "CORE_STATE_ONLINE",
	2: "CORE_STATE_OFFLINE",
}

var CoreState_value = map[string]int32{
	"CORE_STATE_UNKNOWN": 0,
	"CORE_STATE_ONLINE":  1,
	"CORE_STATE_OFFLINE": 2,
}

func (x CoreState) String() string {
	return proto.EnumName(CoreState_name, int32(x))
}

func (CoreState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f6483212e034594c, []int{0}
}

func init() {
	proto.RegisterEnum("ai.metathings.state.CoreState", CoreState_name, CoreState_value)
}

func init() { proto.RegisterFile("core_state.proto", fileDescriptor_f6483212e034594c) }

var fileDescriptor_f6483212e034594c = []byte{
	// 122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xce, 0x2f, 0x4a,
	0x8d, 0x2f, 0x2e, 0x49, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4e, 0xcc,
	0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x03, 0x4b, 0x69, 0x05,
	0x71, 0x71, 0x3a, 0xe7, 0x17, 0xa5, 0x06, 0x83, 0x38, 0x42, 0x62, 0x5c, 0x42, 0xce, 0xfe, 0x41,
	0xae, 0xf1, 0xc1, 0x21, 0x8e, 0x21, 0xae, 0xf1, 0xa1, 0x7e, 0xde, 0x7e, 0xfe, 0xe1, 0x7e, 0x02,
	0x0c, 0x42, 0xa2, 0x5c, 0x82, 0x48, 0xe2, 0xfe, 0x7e, 0x3e, 0x9e, 0x7e, 0xae, 0x02, 0x8c, 0x68,
	0xca, 0xfd, 0xdd, 0xdc, 0xc0, 0xe2, 0x4c, 0x49, 0x6c, 0x60, 0xfb, 0x8c, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x3e, 0x82, 0xaf, 0xe9, 0x83, 0x00, 0x00, 0x00,
}
