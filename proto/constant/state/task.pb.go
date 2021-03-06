// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

package constant

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
	TaskState_TASK_STATE_CREATED TaskState = 10
	TaskState_TASK_STATE_PENDING TaskState = 20
	TaskState_TASK_STATE_RUNNING TaskState = 30
	TaskState_TASK_STATE_DONE    TaskState = 40
	TaskState_TASK_STATE_ERROR   TaskState = 50
)

var TaskState_name = map[int32]string{
	0:  "TASK_STATE_UNKNOWN",
	10: "TASK_STATE_CREATED",
	20: "TASK_STATE_PENDING",
	30: "TASK_STATE_RUNNING",
	40: "TASK_STATE_DONE",
	50: "TASK_STATE_ERROR",
}

var TaskState_value = map[string]int32{
	"TASK_STATE_UNKNOWN": 0,
	"TASK_STATE_CREATED": 10,
	"TASK_STATE_PENDING": 20,
	"TASK_STATE_RUNNING": 30,
	"TASK_STATE_DONE":    40,
	"TASK_STATE_ERROR":   50,
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
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2c, 0xce,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x49, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c,
	0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x4b, 0xce, 0xcf, 0x2b, 0x2e, 0x49, 0xcc, 0x2b, 0xd1, 0x2b,
	0x2e, 0x49, 0x2c, 0x49, 0xd5, 0x9a, 0xc6, 0xc8, 0xc5, 0x19, 0x92, 0x58, 0x9c, 0x1d, 0x0c, 0xe2,
	0x09, 0x89, 0x71, 0x09, 0x85, 0x38, 0x06, 0x7b, 0xc7, 0x07, 0x87, 0x38, 0x86, 0xb8, 0xc6, 0x87,
	0xfa, 0x79, 0xfb, 0xf9, 0x87, 0xfb, 0x09, 0x30, 0xa0, 0x89, 0x3b, 0x07, 0xb9, 0x3a, 0x86, 0xb8,
	0xba, 0x08, 0x70, 0xa1, 0x89, 0x07, 0xb8, 0xfa, 0xb9, 0x78, 0xfa, 0xb9, 0x0b, 0x88, 0xa0, 0x89,
	0x07, 0x85, 0xfa, 0xf9, 0x81, 0xc4, 0xe5, 0x84, 0x84, 0xb9, 0xf8, 0x91, 0xc4, 0x5d, 0xfc, 0xfd,
	0x5c, 0x05, 0x34, 0x84, 0x44, 0xb8, 0x04, 0x90, 0x04, 0x5d, 0x83, 0x82, 0xfc, 0x83, 0x04, 0x8c,
	0x9c, 0xb8, 0xa2, 0x38, 0x60, 0x4e, 0x4d, 0x62, 0x03, 0xfb, 0xc4, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0xf6, 0xdf, 0x52, 0x24, 0xd7, 0x00, 0x00, 0x00,
}
