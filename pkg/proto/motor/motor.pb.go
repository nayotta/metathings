// Code generated by protoc-gen-go. DO NOT EDIT.
// source: motor.proto

package motor

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

type MotorState int32

const (
	MotorState_MOTOR_STATE_UNKNOWN MotorState = 0
	MotorState_MOTOR_STATE_ON      MotorState = 1
	MotorState_MOTOR_STATE_OFF     MotorState = 2
)

var MotorState_name = map[int32]string{
	0: "MOTOR_STATE_UNKNOWN",
	1: "MOTOR_STATE_ON",
	2: "MOTOR_STATE_OFF",
}

var MotorState_value = map[string]int32{
	"MOTOR_STATE_UNKNOWN": 0,
	"MOTOR_STATE_ON":      1,
	"MOTOR_STATE_OFF":     2,
}

func (x MotorState) String() string {
	return proto.EnumName(MotorState_name, int32(x))
}

func (MotorState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7024f82ae1d6a6dd, []int{0}
}

type MotorDirection int32

const (
	MotorDirection_MTOOR_DIRECTION_UNKNOWN  MotorDirection = 0
	MotorDirection_MOTOR_DIRECTION_FORWARD  MotorDirection = 1
	MotorDirection_MOTOR_DIRECTION_BACKWARD MotorDirection = 2
)

var MotorDirection_name = map[int32]string{
	0: "MTOOR_DIRECTION_UNKNOWN",
	1: "MOTOR_DIRECTION_FORWARD",
	2: "MOTOR_DIRECTION_BACKWARD",
}

var MotorDirection_value = map[string]int32{
	"MTOOR_DIRECTION_UNKNOWN":  0,
	"MOTOR_DIRECTION_FORWARD":  1,
	"MOTOR_DIRECTION_BACKWARD": 2,
}

func (x MotorDirection) String() string {
	return proto.EnumName(MotorDirection_name, int32(x))
}

func (MotorDirection) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7024f82ae1d6a6dd, []int{1}
}

type Motor struct {
	Name      string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	State     MotorState     `protobuf:"varint,2,opt,name=state,proto3,enum=ai.metathings.service.motor.MotorState" json:"state,omitempty"`
	Direction MotorDirection `protobuf:"varint,3,opt,name=direction,proto3,enum=ai.metathings.service.motor.MotorDirection" json:"direction,omitempty"`
	// speed range from 0 to 1
	Speed                float32  `protobuf:"fixed32,4,opt,name=speed,proto3" json:"speed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Motor) Reset()         { *m = Motor{} }
func (m *Motor) String() string { return proto.CompactTextString(m) }
func (*Motor) ProtoMessage()    {}
func (*Motor) Descriptor() ([]byte, []int) {
	return fileDescriptor_7024f82ae1d6a6dd, []int{0}
}

func (m *Motor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Motor.Unmarshal(m, b)
}
func (m *Motor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Motor.Marshal(b, m, deterministic)
}
func (m *Motor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Motor.Merge(m, src)
}
func (m *Motor) XXX_Size() int {
	return xxx_messageInfo_Motor.Size(m)
}
func (m *Motor) XXX_DiscardUnknown() {
	xxx_messageInfo_Motor.DiscardUnknown(m)
}

var xxx_messageInfo_Motor proto.InternalMessageInfo

func (m *Motor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Motor) GetState() MotorState {
	if m != nil {
		return m.State
	}
	return MotorState_MOTOR_STATE_UNKNOWN
}

func (m *Motor) GetDirection() MotorDirection {
	if m != nil {
		return m.Direction
	}
	return MotorDirection_MTOOR_DIRECTION_UNKNOWN
}

func (m *Motor) GetSpeed() float32 {
	if m != nil {
		return m.Speed
	}
	return 0
}

func init() {
	proto.RegisterEnum("ai.metathings.service.motor.MotorState", MotorState_name, MotorState_value)
	proto.RegisterEnum("ai.metathings.service.motor.MotorDirection", MotorDirection_name, MotorDirection_value)
	proto.RegisterType((*Motor)(nil), "ai.metathings.service.motor.Motor")
}

func init() { proto.RegisterFile("motor.proto", fileDescriptor_7024f82ae1d6a6dd) }

var fileDescriptor_7024f82ae1d6a6dd = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x4d, 0x5d, 0x85, 0x3d, 0xa1, 0x96, 0x37, 0x61, 0x85, 0x79, 0x28, 0x5e, 0x2c, 0x13,
	0x72, 0xd0, 0xb3, 0x87, 0xba, 0xae, 0x50, 0x46, 0x13, 0xc8, 0x2a, 0x3b, 0x96, 0xba, 0x05, 0x97,
	0x43, 0x9b, 0xd1, 0x06, 0xbf, 0x9c, 0x5f, 0x4e, 0x9a, 0x21, 0xdd, 0x3c, 0xe8, 0xed, 0x25, 0xbf,
	0xff, 0xff, 0x97, 0xc0, 0x83, 0xeb, 0x5a, 0x1b, 0xdd, 0xd2, 0x43, 0xab, 0x8d, 0xc6, 0x59, 0xa5,
	0x68, 0x2d, 0x4d, 0x65, 0xf6, 0xaa, 0xf9, 0xe8, 0x68, 0x27, 0xdb, 0x4f, 0xb5, 0x95, 0xd4, 0x46,
	0xee, 0xbf, 0x08, 0xb8, 0x79, 0x3f, 0x21, 0xc2, 0xa8, 0xa9, 0x6a, 0x19, 0x90, 0x90, 0x44, 0x63,
	0x61, 0x67, 0x7c, 0x01, 0xb7, 0x33, 0x95, 0x91, 0x81, 0x13, 0x92, 0xc8, 0x7b, 0x7a, 0xa0, 0x7f,
	0xa8, 0xa8, 0xd5, 0xac, 0xfb, 0xb8, 0x38, 0xb6, 0x30, 0x83, 0xf1, 0x4e, 0xb5, 0x72, 0x6b, 0x94,
	0x6e, 0x82, 0x4b, 0xab, 0x78, 0xfc, 0x5f, 0x91, 0xfc, 0x54, 0xc4, 0xd0, 0xc6, 0x5b, 0x70, 0xbb,
	0x83, 0x94, 0xbb, 0x60, 0x14, 0x92, 0xc8, 0x11, 0xc7, 0xc3, 0x9c, 0x01, 0x0c, 0xaf, 0xe2, 0x14,
	0x26, 0x39, 0x2f, 0xb8, 0x28, 0xd7, 0x45, 0x5c, 0x2c, 0xcb, 0x37, 0xb6, 0x62, 0x7c, 0xc3, 0xfc,
	0x0b, 0x44, 0xf0, 0x4e, 0x01, 0x67, 0x3e, 0xc1, 0x09, 0xdc, 0x9c, 0xdd, 0xa5, 0xa9, 0xef, 0xcc,
	0xf7, 0xe0, 0x9d, 0x7f, 0x01, 0x67, 0x30, 0xcd, 0x0b, 0xce, 0x45, 0x99, 0x64, 0x62, 0xb9, 0x28,
	0x32, 0xce, 0x4e, 0xbc, 0x3d, 0xb4, 0x8e, 0x01, 0xa6, 0x5c, 0x6c, 0x62, 0x91, 0xf8, 0x04, 0xef,
	0x20, 0xf8, 0x0d, 0x5f, 0xe3, 0xc5, 0xca, 0x52, 0xe7, 0xfd, 0xca, 0xee, 0xe6, 0xf9, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0x5e, 0x92, 0x29, 0x52, 0xaa, 0x01, 0x00, 0x00,
}
