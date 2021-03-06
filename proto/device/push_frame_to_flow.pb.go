// Code generated by protoc-gen-go. DO NOT EDIT.
// source: push_frame_to_flow.proto

package device

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type PushFrameToFlowRequest struct {
	Id *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Request:
	//	*PushFrameToFlowRequest_Config_
	//	*PushFrameToFlowRequest_Ping_
	//	*PushFrameToFlowRequest_Frame
	Request              isPushFrameToFlowRequest_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *PushFrameToFlowRequest) Reset()         { *m = PushFrameToFlowRequest{} }
func (m *PushFrameToFlowRequest) String() string { return proto.CompactTextString(m) }
func (*PushFrameToFlowRequest) ProtoMessage()    {}
func (*PushFrameToFlowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d9d865a9a5db986, []int{0}
}

func (m *PushFrameToFlowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushFrameToFlowRequest.Unmarshal(m, b)
}
func (m *PushFrameToFlowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushFrameToFlowRequest.Marshal(b, m, deterministic)
}
func (m *PushFrameToFlowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushFrameToFlowRequest.Merge(m, src)
}
func (m *PushFrameToFlowRequest) XXX_Size() int {
	return xxx_messageInfo_PushFrameToFlowRequest.Size(m)
}
func (m *PushFrameToFlowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushFrameToFlowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushFrameToFlowRequest proto.InternalMessageInfo

func (m *PushFrameToFlowRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

type isPushFrameToFlowRequest_Request interface {
	isPushFrameToFlowRequest_Request()
}

type PushFrameToFlowRequest_Config_ struct {
	Config *PushFrameToFlowRequest_Config `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type PushFrameToFlowRequest_Ping_ struct {
	Ping *PushFrameToFlowRequest_Ping `protobuf:"bytes,3,opt,name=ping,proto3,oneof"`
}

type PushFrameToFlowRequest_Frame struct {
	Frame *deviced.OpFrame `protobuf:"bytes,4,opt,name=frame,proto3,oneof"`
}

func (*PushFrameToFlowRequest_Config_) isPushFrameToFlowRequest_Request() {}

func (*PushFrameToFlowRequest_Ping_) isPushFrameToFlowRequest_Request() {}

func (*PushFrameToFlowRequest_Frame) isPushFrameToFlowRequest_Request() {}

func (m *PushFrameToFlowRequest) GetRequest() isPushFrameToFlowRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *PushFrameToFlowRequest) GetConfig() *PushFrameToFlowRequest_Config {
	if x, ok := m.GetRequest().(*PushFrameToFlowRequest_Config_); ok {
		return x.Config
	}
	return nil
}

func (m *PushFrameToFlowRequest) GetPing() *PushFrameToFlowRequest_Ping {
	if x, ok := m.GetRequest().(*PushFrameToFlowRequest_Ping_); ok {
		return x.Ping
	}
	return nil
}

func (m *PushFrameToFlowRequest) GetFrame() *deviced.OpFrame {
	if x, ok := m.GetRequest().(*PushFrameToFlowRequest_Frame); ok {
		return x.Frame
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PushFrameToFlowRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PushFrameToFlowRequest_Config_)(nil),
		(*PushFrameToFlowRequest_Ping_)(nil),
		(*PushFrameToFlowRequest_Frame)(nil),
	}
}

type PushFrameToFlowRequest_Config struct {
	Flow                 *deviced.OpFlow     `protobuf:"bytes,1,opt,name=flow,proto3" json:"flow,omitempty"`
	ConfigAck            *wrappers.BoolValue `protobuf:"bytes,2,opt,name=config_ack,json=configAck,proto3" json:"config_ack,omitempty"`
	PushAck              *wrappers.BoolValue `protobuf:"bytes,3,opt,name=push_ack,json=pushAck,proto3" json:"push_ack,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PushFrameToFlowRequest_Config) Reset()         { *m = PushFrameToFlowRequest_Config{} }
func (m *PushFrameToFlowRequest_Config) String() string { return proto.CompactTextString(m) }
func (*PushFrameToFlowRequest_Config) ProtoMessage()    {}
func (*PushFrameToFlowRequest_Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d9d865a9a5db986, []int{0, 0}
}

func (m *PushFrameToFlowRequest_Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushFrameToFlowRequest_Config.Unmarshal(m, b)
}
func (m *PushFrameToFlowRequest_Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushFrameToFlowRequest_Config.Marshal(b, m, deterministic)
}
func (m *PushFrameToFlowRequest_Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushFrameToFlowRequest_Config.Merge(m, src)
}
func (m *PushFrameToFlowRequest_Config) XXX_Size() int {
	return xxx_messageInfo_PushFrameToFlowRequest_Config.Size(m)
}
func (m *PushFrameToFlowRequest_Config) XXX_DiscardUnknown() {
	xxx_messageInfo_PushFrameToFlowRequest_Config.DiscardUnknown(m)
}

var xxx_messageInfo_PushFrameToFlowRequest_Config proto.InternalMessageInfo

func (m *PushFrameToFlowRequest_Config) GetFlow() *deviced.OpFlow {
	if m != nil {
		return m.Flow
	}
	return nil
}

func (m *PushFrameToFlowRequest_Config) GetConfigAck() *wrappers.BoolValue {
	if m != nil {
		return m.ConfigAck
	}
	return nil
}

func (m *PushFrameToFlowRequest_Config) GetPushAck() *wrappers.BoolValue {
	if m != nil {
		return m.PushAck
	}
	return nil
}

type PushFrameToFlowRequest_Ping struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushFrameToFlowRequest_Ping) Reset()         { *m = PushFrameToFlowRequest_Ping{} }
func (m *PushFrameToFlowRequest_Ping) String() string { return proto.CompactTextString(m) }
func (*PushFrameToFlowRequest_Ping) ProtoMessage()    {}
func (*PushFrameToFlowRequest_Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d9d865a9a5db986, []int{0, 1}
}

func (m *PushFrameToFlowRequest_Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushFrameToFlowRequest_Ping.Unmarshal(m, b)
}
func (m *PushFrameToFlowRequest_Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushFrameToFlowRequest_Ping.Marshal(b, m, deterministic)
}
func (m *PushFrameToFlowRequest_Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushFrameToFlowRequest_Ping.Merge(m, src)
}
func (m *PushFrameToFlowRequest_Ping) XXX_Size() int {
	return xxx_messageInfo_PushFrameToFlowRequest_Ping.Size(m)
}
func (m *PushFrameToFlowRequest_Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_PushFrameToFlowRequest_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_PushFrameToFlowRequest_Ping proto.InternalMessageInfo

type PushFrameToFlowResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Response:
	//	*PushFrameToFlowResponse_Config_
	//	*PushFrameToFlowResponse_Pong_
	//	*PushFrameToFlowResponse_Ack_
	Response             isPushFrameToFlowResponse_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *PushFrameToFlowResponse) Reset()         { *m = PushFrameToFlowResponse{} }
func (m *PushFrameToFlowResponse) String() string { return proto.CompactTextString(m) }
func (*PushFrameToFlowResponse) ProtoMessage()    {}
func (*PushFrameToFlowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d9d865a9a5db986, []int{1}
}

func (m *PushFrameToFlowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushFrameToFlowResponse.Unmarshal(m, b)
}
func (m *PushFrameToFlowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushFrameToFlowResponse.Marshal(b, m, deterministic)
}
func (m *PushFrameToFlowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushFrameToFlowResponse.Merge(m, src)
}
func (m *PushFrameToFlowResponse) XXX_Size() int {
	return xxx_messageInfo_PushFrameToFlowResponse.Size(m)
}
func (m *PushFrameToFlowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PushFrameToFlowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PushFrameToFlowResponse proto.InternalMessageInfo

func (m *PushFrameToFlowResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type isPushFrameToFlowResponse_Response interface {
	isPushFrameToFlowResponse_Response()
}

type PushFrameToFlowResponse_Config_ struct {
	Config *PushFrameToFlowResponse_Config `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type PushFrameToFlowResponse_Pong_ struct {
	Pong *PushFrameToFlowResponse_Pong `protobuf:"bytes,3,opt,name=pong,proto3,oneof"`
}

type PushFrameToFlowResponse_Ack_ struct {
	Ack *PushFrameToFlowResponse_Ack `protobuf:"bytes,4,opt,name=ack,proto3,oneof"`
}

func (*PushFrameToFlowResponse_Config_) isPushFrameToFlowResponse_Response() {}

func (*PushFrameToFlowResponse_Pong_) isPushFrameToFlowResponse_Response() {}

func (*PushFrameToFlowResponse_Ack_) isPushFrameToFlowResponse_Response() {}

func (m *PushFrameToFlowResponse) GetResponse() isPushFrameToFlowResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *PushFrameToFlowResponse) GetConfig() *PushFrameToFlowResponse_Config {
	if x, ok := m.GetResponse().(*PushFrameToFlowResponse_Config_); ok {
		return x.Config
	}
	return nil
}

func (m *PushFrameToFlowResponse) GetPong() *PushFrameToFlowResponse_Pong {
	if x, ok := m.GetResponse().(*PushFrameToFlowResponse_Pong_); ok {
		return x.Pong
	}
	return nil
}

func (m *PushFrameToFlowResponse) GetAck() *PushFrameToFlowResponse_Ack {
	if x, ok := m.GetResponse().(*PushFrameToFlowResponse_Ack_); ok {
		return x.Ack
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PushFrameToFlowResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PushFrameToFlowResponse_Config_)(nil),
		(*PushFrameToFlowResponse_Pong_)(nil),
		(*PushFrameToFlowResponse_Ack_)(nil),
	}
}

type PushFrameToFlowResponse_Config struct {
	Session              string   `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushFrameToFlowResponse_Config) Reset()         { *m = PushFrameToFlowResponse_Config{} }
func (m *PushFrameToFlowResponse_Config) String() string { return proto.CompactTextString(m) }
func (*PushFrameToFlowResponse_Config) ProtoMessage()    {}
func (*PushFrameToFlowResponse_Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d9d865a9a5db986, []int{1, 0}
}

func (m *PushFrameToFlowResponse_Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushFrameToFlowResponse_Config.Unmarshal(m, b)
}
func (m *PushFrameToFlowResponse_Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushFrameToFlowResponse_Config.Marshal(b, m, deterministic)
}
func (m *PushFrameToFlowResponse_Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushFrameToFlowResponse_Config.Merge(m, src)
}
func (m *PushFrameToFlowResponse_Config) XXX_Size() int {
	return xxx_messageInfo_PushFrameToFlowResponse_Config.Size(m)
}
func (m *PushFrameToFlowResponse_Config) XXX_DiscardUnknown() {
	xxx_messageInfo_PushFrameToFlowResponse_Config.DiscardUnknown(m)
}

var xxx_messageInfo_PushFrameToFlowResponse_Config proto.InternalMessageInfo

func (m *PushFrameToFlowResponse_Config) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

type PushFrameToFlowResponse_Ack struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushFrameToFlowResponse_Ack) Reset()         { *m = PushFrameToFlowResponse_Ack{} }
func (m *PushFrameToFlowResponse_Ack) String() string { return proto.CompactTextString(m) }
func (*PushFrameToFlowResponse_Ack) ProtoMessage()    {}
func (*PushFrameToFlowResponse_Ack) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d9d865a9a5db986, []int{1, 1}
}

func (m *PushFrameToFlowResponse_Ack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushFrameToFlowResponse_Ack.Unmarshal(m, b)
}
func (m *PushFrameToFlowResponse_Ack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushFrameToFlowResponse_Ack.Marshal(b, m, deterministic)
}
func (m *PushFrameToFlowResponse_Ack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushFrameToFlowResponse_Ack.Merge(m, src)
}
func (m *PushFrameToFlowResponse_Ack) XXX_Size() int {
	return xxx_messageInfo_PushFrameToFlowResponse_Ack.Size(m)
}
func (m *PushFrameToFlowResponse_Ack) XXX_DiscardUnknown() {
	xxx_messageInfo_PushFrameToFlowResponse_Ack.DiscardUnknown(m)
}

var xxx_messageInfo_PushFrameToFlowResponse_Ack proto.InternalMessageInfo

type PushFrameToFlowResponse_Pong struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushFrameToFlowResponse_Pong) Reset()         { *m = PushFrameToFlowResponse_Pong{} }
func (m *PushFrameToFlowResponse_Pong) String() string { return proto.CompactTextString(m) }
func (*PushFrameToFlowResponse_Pong) ProtoMessage()    {}
func (*PushFrameToFlowResponse_Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d9d865a9a5db986, []int{1, 2}
}

func (m *PushFrameToFlowResponse_Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushFrameToFlowResponse_Pong.Unmarshal(m, b)
}
func (m *PushFrameToFlowResponse_Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushFrameToFlowResponse_Pong.Marshal(b, m, deterministic)
}
func (m *PushFrameToFlowResponse_Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushFrameToFlowResponse_Pong.Merge(m, src)
}
func (m *PushFrameToFlowResponse_Pong) XXX_Size() int {
	return xxx_messageInfo_PushFrameToFlowResponse_Pong.Size(m)
}
func (m *PushFrameToFlowResponse_Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_PushFrameToFlowResponse_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_PushFrameToFlowResponse_Pong proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PushFrameToFlowRequest)(nil), "ai.metathings.service.device.PushFrameToFlowRequest")
	proto.RegisterType((*PushFrameToFlowRequest_Config)(nil), "ai.metathings.service.device.PushFrameToFlowRequest.Config")
	proto.RegisterType((*PushFrameToFlowRequest_Ping)(nil), "ai.metathings.service.device.PushFrameToFlowRequest.Ping")
	proto.RegisterType((*PushFrameToFlowResponse)(nil), "ai.metathings.service.device.PushFrameToFlowResponse")
	proto.RegisterType((*PushFrameToFlowResponse_Config)(nil), "ai.metathings.service.device.PushFrameToFlowResponse.Config")
	proto.RegisterType((*PushFrameToFlowResponse_Ack)(nil), "ai.metathings.service.device.PushFrameToFlowResponse.Ack")
	proto.RegisterType((*PushFrameToFlowResponse_Pong)(nil), "ai.metathings.service.device.PushFrameToFlowResponse.Pong")
}

func init() { proto.RegisterFile("push_frame_to_flow.proto", fileDescriptor_6d9d865a9a5db986) }

var fileDescriptor_6d9d865a9a5db986 = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x86, 0xdb, 0x34, 0x4d, 0xdb, 0x23, 0x88, 0xcc, 0x85, 0x1b, 0xc2, 0x22, 0x52, 0x50, 0xbc,
	0x9a, 0x80, 0xa2, 0xb2, 0x2a, 0x42, 0xb3, 0xb0, 0xf4, 0x46, 0xb6, 0x44, 0xdd, 0x0b, 0x6f, 0xca,
	0x34, 0x99, 0xa6, 0x43, 0xd3, 0x39, 0x31, 0x33, 0xd9, 0xe2, 0x2b, 0x08, 0x3e, 0x93, 0xcf, 0xe2,
	0x6b, 0x78, 0x25, 0x33, 0x93, 0x56, 0xb1, 0xb2, 0x8b, 0xbd, 0x4a, 0x42, 0xce, 0xff, 0x9d, 0x73,
	0xfe, 0xff, 0x40, 0x58, 0x35, 0x6a, 0x35, 0x5f, 0xd6, 0x6c, 0xc3, 0xe7, 0x1a, 0xe7, 0xcb, 0x12,
	0xb7, 0xb4, 0xaa, 0x51, 0x23, 0x39, 0x65, 0x82, 0x6e, 0xb8, 0x66, 0x7a, 0x25, 0x64, 0xa1, 0xa8,
	0xe2, 0xf5, 0xb5, 0xc8, 0x38, 0xcd, 0xb9, 0x79, 0x44, 0x0f, 0x0a, 0xc4, 0xa2, 0xe4, 0xb1, 0xad,
	0x5d, 0x34, 0xcb, 0x78, 0x5b, 0xb3, 0xaa, 0xe2, 0xb5, 0x72, 0xea, 0xe8, 0xe4, 0x9a, 0x95, 0x22,
	0x67, 0x9a, 0xc7, 0xbb, 0x97, 0xf6, 0xc7, 0xcb, 0x42, 0xe8, 0x55, 0xb3, 0xa0, 0x19, 0x6e, 0x62,
	0xc9, 0xbe, 0xa0, 0xd6, 0x2c, 0xfe, 0xdd, 0xc6, 0xf1, 0x62, 0xd7, 0x24, 0x8f, 0x37, 0x98, 0xf3,
	0xd2, 0x09, 0xc7, 0xdf, 0x7c, 0xb8, 0x3f, 0x6b, 0xd4, 0xea, 0xc2, 0xcc, 0xfa, 0x01, 0x2f, 0x4a,
	0xdc, 0xa6, 0xfc, 0x73, 0xc3, 0x95, 0x26, 0x2f, 0xc0, 0x13, 0x79, 0xd8, 0x7d, 0xd8, 0x7d, 0x72,
	0xe7, 0xe9, 0x29, 0x75, 0x93, 0xd1, 0xdd, 0x64, 0xf4, 0xbd, 0xae, 0x85, 0x2c, 0xae, 0x58, 0xd9,
	0xf0, 0x64, 0xf8, 0x33, 0xe9, 0x7f, 0xed, 0x7a, 0xf7, 0xba, 0xa9, 0x27, 0x72, 0xf2, 0x11, 0x82,
	0x0c, 0xe5, 0x52, 0x14, 0xa1, 0x67, 0xb5, 0xaf, 0xe9, 0x4d, 0x3b, 0xd3, 0x7f, 0x77, 0xa7, 0xe7,
	0x16, 0x31, 0xed, 0xa4, 0x2d, 0x8c, 0x5c, 0x82, 0x5f, 0x09, 0x59, 0x84, 0x3d, 0x0b, 0x3d, 0x3b,
	0x0a, 0x3a, 0x13, 0xd2, 0x20, 0x2d, 0x88, 0xbc, 0x85, 0xbe, 0x4d, 0x28, 0xf4, 0x2d, 0xf1, 0xf1,
	0x8d, 0xc4, 0x9c, 0x5e, 0x56, 0x16, 0x38, 0xed, 0xa4, 0x4e, 0x16, 0x7d, 0xef, 0x42, 0xe0, 0xa6,
	0x24, 0xe7, 0xe0, 0x9b, 0x8c, 0x5b, 0xb3, 0x1e, 0xdd, 0x4e, 0x2a, 0x71, 0xfb, 0x87, 0x6b, 0x56,
	0x4c, 0xce, 0x00, 0xdc, 0xaa, 0x73, 0x96, 0xad, 0x5b, 0xef, 0xa2, 0x03, 0xdf, 0x13, 0xc4, 0xd2,
	0xba, 0x9e, 0x8e, 0x5c, 0xf5, 0x24, 0x5b, 0x93, 0xe7, 0x30, 0xb4, 0x17, 0x67, 0x84, 0xbd, 0x5b,
	0x85, 0x03, 0x53, 0x3b, 0xc9, 0xd6, 0x51, 0x00, 0xbe, 0x71, 0x24, 0x19, 0xc1, 0xa0, 0x76, 0x0e,
	0x8d, 0x7f, 0x78, 0x70, 0x72, 0x60, 0x9e, 0xaa, 0x50, 0x2a, 0x4e, 0xee, 0xee, 0x0f, 0x62, 0x64,
	0x83, 0xbe, 0xfa, 0x2b, 0xe8, 0x37, 0xff, 0x99, 0x89, 0xc3, 0x1e, 0x26, 0x3d, 0x03, 0xbf, 0xc2,
	0x7d, 0xd2, 0xaf, 0x8e, 0xa3, 0xce, 0xb0, 0x8d, 0x1a, 0x65, 0x41, 0xde, 0x41, 0xcf, 0x58, 0xe3,
	0x1f, 0x75, 0x3a, 0x2d, 0x70, 0x92, 0xad, 0xa7, 0x9d, 0xd4, 0x70, 0xa2, 0xf1, 0x3e, 0xf8, 0x10,
	0x06, 0x8a, 0x2b, 0x25, 0x50, 0xb6, 0xbe, 0xec, 0x3e, 0xa3, 0x3e, 0xf4, 0x76, 0x16, 0xa3, 0x2c,
	0x12, 0x80, 0x61, 0xdd, 0x92, 0x92, 0xe1, 0xa7, 0xc0, 0xf5, 0x5a, 0x04, 0x36, 0x9d, 0x67, 0xbf,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x8c, 0xed, 0x5c, 0x30, 0x04, 0x00, 0x00,
}
