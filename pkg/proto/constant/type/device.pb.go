// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device.proto

package ai_metathings_constant_type

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DeviceType int32

const (
	DeviceType_DEVICE_TYPE_UNKNOWN  DeviceType = 0
	DeviceType_DEVICE_TYPE_SIMPLE   DeviceType = 1
	DeviceType_DEVICE_TYPE_ADVANCED DeviceType = 2
)

var DeviceType_name = map[int32]string{
	0: "DEVICE_TYPE_UNKNOWN",
	1: "DEVICE_TYPE_SIMPLE",
	2: "DEVICE_TYPE_ADVANCED",
}

var DeviceType_value = map[string]int32{
	"DEVICE_TYPE_UNKNOWN":  0,
	"DEVICE_TYPE_SIMPLE":   1,
	"DEVICE_TYPE_ADVANCED": 2,
}

func (x DeviceType) String() string {
	return proto.EnumName(DeviceType_name, int32(x))
}

func (DeviceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{0}
}

func init() {
	proto.RegisterEnum("ai.metathings.constant.type.DeviceType", DeviceType_name, DeviceType_value)
}

func init() { proto.RegisterFile("device.proto", fileDescriptor_870276a56ac00da5) }

var fileDescriptor_870276a56ac00da5 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x49, 0x2d, 0xcb,
	0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4e, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d,
	0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x4b, 0xce, 0xcf, 0x2b, 0x2e, 0x49, 0xcc, 0x2b,
	0xd1, 0x2b, 0xa9, 0x2c, 0x48, 0xd5, 0x0a, 0xe7, 0xe2, 0x72, 0x01, 0x2b, 0x0e, 0xa9, 0x2c, 0x48,
	0x15, 0x12, 0xe7, 0x12, 0x76, 0x71, 0x0d, 0xf3, 0x74, 0x76, 0x8d, 0x0f, 0x89, 0x0c, 0x70, 0x8d,
	0x0f, 0xf5, 0xf3, 0xf6, 0xf3, 0x0f, 0xf7, 0x13, 0x60, 0x10, 0x12, 0xe3, 0x12, 0x42, 0x96, 0x08,
	0xf6, 0xf4, 0x0d, 0xf0, 0x71, 0x15, 0x60, 0x14, 0x92, 0xe0, 0x12, 0x41, 0x16, 0x77, 0x74, 0x09,
	0x73, 0xf4, 0x73, 0x76, 0x75, 0x11, 0x60, 0x4a, 0x62, 0x03, 0x5b, 0x6e, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x4a, 0xe9, 0x71, 0x3b, 0x8c, 0x00, 0x00, 0x00,
}
