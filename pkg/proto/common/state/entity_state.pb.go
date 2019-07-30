// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entity_state.proto

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EntityState int32

const (
	EntityState_ENTITY_STATE_UNKNOWN EntityState = 0
	EntityState_ENTITY_STATE_ONLINE  EntityState = 1
	EntityState_ENTITY_STATE_OFFLIEN EntityState = 2
)

var EntityState_name = map[int32]string{
	0: "ENTITY_STATE_UNKNOWN",
	1: "ENTITY_STATE_ONLINE",
	2: "ENTITY_STATE_OFFLIEN",
}

var EntityState_value = map[string]int32{
	"ENTITY_STATE_UNKNOWN": 0,
	"ENTITY_STATE_ONLINE":  1,
	"ENTITY_STATE_OFFLIEN": 2,
}

func (x EntityState) String() string {
	return proto.EnumName(EntityState_name, int32(x))
}

func (EntityState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_888bca0db16bdd3e, []int{0}
}

func init() {
	proto.RegisterEnum("ai.metathings.state.EntityState", EntityState_name, EntityState_value)
}

func init() { proto.RegisterFile("entity_state.proto", fileDescriptor_888bca0db16bdd3e) }

var fileDescriptor_888bca0db16bdd3e = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xcd, 0x2b, 0xc9,
	0x2c, 0xa9, 0x8c, 0x2f, 0x2e, 0x49, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x4e, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x03, 0x4b,
	0x69, 0x45, 0x71, 0x71, 0xbb, 0x82, 0x95, 0x06, 0x83, 0xb8, 0x42, 0x12, 0x5c, 0x22, 0xae, 0x7e,
	0x21, 0x9e, 0x21, 0x91, 0xf1, 0xc1, 0x21, 0x8e, 0x21, 0xae, 0xf1, 0xa1, 0x7e, 0xde, 0x7e, 0xfe,
	0xe1, 0x7e, 0x02, 0x0c, 0x42, 0xe2, 0x5c, 0xc2, 0x28, 0x32, 0xfe, 0x7e, 0x3e, 0x9e, 0x7e, 0xae,
	0x02, 0x8c, 0x18, 0x5a, 0xfc, 0xdd, 0xdc, 0x7c, 0x3c, 0x5d, 0xfd, 0x04, 0x98, 0x92, 0xd8, 0xc0,
	0xf6, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x6a, 0x7b, 0xc1, 0x8d, 0x00, 0x00, 0x00,
}
