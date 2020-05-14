// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

package ai_metathings_constant_state

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

type TaskState int32

const (
	TaskState_TASK_STATE_UNKNOWN TaskState = 0
	TaskState_TASK_STATE_CREATED TaskState = 1
	TaskState_TASK_STATE_RUNNING TaskState = 2
	TaskState_TASK_STATE_DONE    TaskState = 3
	TaskState_TASK_STATE_ERROR   TaskState = 4
)

var TaskState_name = map[int32]string{
	0: "TASK_STATE_UNKNOWN",
	1: "TASK_STATE_CREATED",
	2: "TASK_STATE_RUNNING",
	3: "TASK_STATE_DONE",
	4: "TASK_STATE_ERROR",
}

var TaskState_value = map[string]int32{
	"TASK_STATE_UNKNOWN": 0,
	"TASK_STATE_CREATED": 1,
	"TASK_STATE_RUNNING": 2,
	"TASK_STATE_DONE":    3,
	"TASK_STATE_ERROR":   4,
}

func (x TaskState) String() string {
	return proto.EnumName(TaskState_name, int32(x))
}

func (TaskState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{0}
}

func init() {
	proto.RegisterEnum("ai.metathings.constant.state.TaskState", TaskState_name, TaskState_value)
}

func init() { proto.RegisterFile("task.proto", fileDescriptor_ce5d8dd45b4a91ff) }

var fileDescriptor_ce5d8dd45b4a91ff = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2c, 0xce,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x49, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c,
	0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x4b, 0xce, 0xcf, 0x2b, 0x2e, 0x49, 0xcc, 0x2b, 0xd1, 0x2b,
	0x2e, 0x49, 0x2c, 0x49, 0xd5, 0xaa, 0xe3, 0xe2, 0x0c, 0x49, 0x2c, 0xce, 0x0e, 0x06, 0x71, 0x84,
	0xc4, 0xb8, 0x84, 0x42, 0x1c, 0x83, 0xbd, 0xe3, 0x83, 0x43, 0x1c, 0x43, 0x5c, 0xe3, 0x43, 0xfd,
	0xbc, 0xfd, 0xfc, 0xc3, 0xfd, 0x04, 0x18, 0xd0, 0xc4, 0x9d, 0x83, 0x5c, 0x1d, 0x43, 0x5c, 0x5d,
	0x04, 0x18, 0xd1, 0xc4, 0x83, 0x42, 0xfd, 0xfc, 0x3c, 0xfd, 0xdc, 0x05, 0x98, 0x84, 0x84, 0xb9,
	0xf8, 0x91, 0xc4, 0x5d, 0xfc, 0xfd, 0x5c, 0x05, 0x98, 0x85, 0x44, 0xb8, 0x04, 0x90, 0x04, 0x5d,
	0x83, 0x82, 0xfc, 0x83, 0x04, 0x58, 0x92, 0xd8, 0xc0, 0x8e, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x0d, 0x92, 0x9d, 0xe0, 0xb2, 0x00, 0x00, 0x00,
}
