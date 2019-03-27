// Code generated by protoc-gen-go. DO NOT EDIT.
// source: push_frame_to_flow.proto

package ai_metathings_service_device

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/mwitkow/go-proto-validators"
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

type PushFrameToFlowRequest struct {
	Id *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Request:
	//	*PushFrameToFlowRequest_Config_
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

type PushFrameToFlowRequest_Frame struct {
	Frame *deviced.OpFrame `protobuf:"bytes,3,opt,name=frame,proto3,oneof"`
}

func (*PushFrameToFlowRequest_Config_) isPushFrameToFlowRequest_Request() {}

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

func (m *PushFrameToFlowRequest) GetFrame() *deviced.OpFrame {
	if x, ok := m.GetRequest().(*PushFrameToFlowRequest_Frame); ok {
		return x.Frame
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PushFrameToFlowRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PushFrameToFlowRequest_OneofMarshaler, _PushFrameToFlowRequest_OneofUnmarshaler, _PushFrameToFlowRequest_OneofSizer, []interface{}{
		(*PushFrameToFlowRequest_Config_)(nil),
		(*PushFrameToFlowRequest_Frame)(nil),
	}
}

func _PushFrameToFlowRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PushFrameToFlowRequest)
	// request
	switch x := m.Request.(type) {
	case *PushFrameToFlowRequest_Config_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Config); err != nil {
			return err
		}
	case *PushFrameToFlowRequest_Frame:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Frame); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("PushFrameToFlowRequest.Request has unexpected type %T", x)
	}
	return nil
}

func _PushFrameToFlowRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PushFrameToFlowRequest)
	switch tag {
	case 2: // request.config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PushFrameToFlowRequest_Config)
		err := b.DecodeMessage(msg)
		m.Request = &PushFrameToFlowRequest_Config_{msg}
		return true, err
	case 3: // request.frame
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(deviced.OpFrame)
		err := b.DecodeMessage(msg)
		m.Request = &PushFrameToFlowRequest_Frame{msg}
		return true, err
	default:
		return false, nil
	}
}

func _PushFrameToFlowRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PushFrameToFlowRequest)
	// request
	switch x := m.Request.(type) {
	case *PushFrameToFlowRequest_Config_:
		s := proto.Size(x.Config)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PushFrameToFlowRequest_Frame:
		s := proto.Size(x.Frame)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
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

type PushFrameToFlowResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Response:
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

type PushFrameToFlowResponse_Ack_ struct {
	Ack *PushFrameToFlowResponse_Ack `protobuf:"bytes,2,opt,name=ack,proto3,oneof"`
}

func (*PushFrameToFlowResponse_Ack_) isPushFrameToFlowResponse_Response() {}

func (m *PushFrameToFlowResponse) GetResponse() isPushFrameToFlowResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *PushFrameToFlowResponse) GetAck() *PushFrameToFlowResponse_Ack {
	if x, ok := m.GetResponse().(*PushFrameToFlowResponse_Ack_); ok {
		return x.Ack
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PushFrameToFlowResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PushFrameToFlowResponse_OneofMarshaler, _PushFrameToFlowResponse_OneofUnmarshaler, _PushFrameToFlowResponse_OneofSizer, []interface{}{
		(*PushFrameToFlowResponse_Ack_)(nil),
	}
}

func _PushFrameToFlowResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PushFrameToFlowResponse)
	// response
	switch x := m.Response.(type) {
	case *PushFrameToFlowResponse_Ack_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ack); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("PushFrameToFlowResponse.Response has unexpected type %T", x)
	}
	return nil
}

func _PushFrameToFlowResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PushFrameToFlowResponse)
	switch tag {
	case 2: // response.ack
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PushFrameToFlowResponse_Ack)
		err := b.DecodeMessage(msg)
		m.Response = &PushFrameToFlowResponse_Ack_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _PushFrameToFlowResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PushFrameToFlowResponse)
	// response
	switch x := m.Response.(type) {
	case *PushFrameToFlowResponse_Ack_:
		s := proto.Size(x.Ack)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
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
	return fileDescriptor_6d9d865a9a5db986, []int{1, 0}
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

func init() {
	proto.RegisterType((*PushFrameToFlowRequest)(nil), "ai.metathings.service.device.PushFrameToFlowRequest")
	proto.RegisterType((*PushFrameToFlowRequest_Config)(nil), "ai.metathings.service.device.PushFrameToFlowRequest.Config")
	proto.RegisterType((*PushFrameToFlowResponse)(nil), "ai.metathings.service.device.PushFrameToFlowResponse")
	proto.RegisterType((*PushFrameToFlowResponse_Ack)(nil), "ai.metathings.service.device.PushFrameToFlowResponse.Ack")
}

func init() { proto.RegisterFile("push_frame_to_flow.proto", fileDescriptor_6d9d865a9a5db986) }

var fileDescriptor_6d9d865a9a5db986 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xdf, 0x8a, 0xd4, 0x30,
	0x14, 0xc6, 0xa7, 0x53, 0x77, 0x76, 0xe7, 0x08, 0x5e, 0xe4, 0x42, 0x87, 0xb2, 0xe8, 0xb2, 0xa0,
	0x78, 0xb3, 0x09, 0xf8, 0x0f, 0x96, 0x05, 0x61, 0x46, 0x58, 0xe6, 0x46, 0x94, 0xfa, 0xe7, 0x76,
	0xc8, 0xb4, 0x69, 0x1a, 0x9a, 0xf6, 0xd4, 0x24, 0xdd, 0xe2, 0x33, 0xf8, 0x32, 0xbe, 0x80, 0xcf,
	0x22, 0xf8, 0x24, 0xd2, 0xa4, 0xce, 0x0a, 0x2b, 0x3b, 0x78, 0xd5, 0x42, 0xce, 0xf7, 0xcb, 0x77,
	0xbe, 0x2f, 0xb0, 0x68, 0x3b, 0x5b, 0x6e, 0x0a, 0xc3, 0x6b, 0xb1, 0x71, 0xb8, 0x29, 0x34, 0xf6,
	0xb4, 0x35, 0xe8, 0x90, 0x1c, 0x73, 0x45, 0x6b, 0xe1, 0xb8, 0x2b, 0x55, 0x23, 0x2d, 0xb5, 0xc2,
	0x5c, 0xa9, 0x4c, 0xd0, 0x5c, 0x0c, 0x9f, 0xe4, 0xa1, 0x44, 0x94, 0x5a, 0x30, 0x3f, 0xbb, 0xed,
	0x0a, 0xd6, 0x1b, 0xde, 0xb6, 0xc2, 0xd8, 0xa0, 0x4e, 0x5e, 0x49, 0xe5, 0xca, 0x6e, 0x4b, 0x33,
	0xac, 0x59, 0xdd, 0x2b, 0x57, 0x61, 0xcf, 0x24, 0x9e, 0xf9, 0xc3, 0xb3, 0x2b, 0xae, 0x55, 0xce,
	0x1d, 0x1a, 0xcb, 0x76, 0xbf, 0xa3, 0xee, 0xe2, 0x2f, 0x5d, 0xc3, 0xbf, 0xa2, 0x73, 0x9c, 0x5d,
	0xbb, 0x60, 0x6d, 0x25, 0xc3, 0x95, 0x2c, 0xf8, 0xc8, 0x59, 0x8d, 0xb9, 0xd0, 0x41, 0x7c, 0xfa,
	0x3d, 0x86, 0xfb, 0xef, 0x3b, 0x5b, 0x5e, 0x0e, 0xeb, 0x7c, 0xc4, 0x4b, 0x8d, 0x7d, 0x2a, 0xbe,
	0x74, 0xc2, 0x3a, 0xf2, 0x02, 0xa6, 0x2a, 0x5f, 0x44, 0x27, 0xd1, 0xd3, 0xbb, 0xcf, 0x8e, 0x69,
	0x30, 0x4f, 0xff, 0x98, 0xa7, 0x1f, 0x9c, 0x51, 0x8d, 0xfc, 0xcc, 0x75, 0x27, 0x56, 0xb3, 0x5f,
	0x3f, 0x1f, 0x4d, 0x4f, 0xa2, 0x74, 0xaa, 0x72, 0xf2, 0x09, 0x66, 0x19, 0x36, 0x85, 0x92, 0x8b,
	0xa9, 0x57, 0x5e, 0xd0, 0xdb, 0x42, 0xa1, 0xff, 0xbe, 0x9b, 0xbe, 0xf1, 0x88, 0xf5, 0x24, 0x1d,
	0x61, 0xe4, 0x35, 0x1c, 0xf8, 0xc4, 0x17, 0xb1, 0xa7, 0x3e, 0xb9, 0x95, 0x9a, 0xd3, 0x77, 0xad,
	0x87, 0xae, 0x27, 0x69, 0x90, 0x25, 0x3f, 0x22, 0x98, 0x05, 0x28, 0x59, 0xc2, 0x9d, 0xa1, 0xb3,
	0x71, 0xb3, 0xc7, 0xfb, 0x49, 0x1a, 0xfb, 0xdd, 0x8a, 0x5e, 0x4a, 0xce, 0x01, 0x82, 0xaf, 0x0d,
	0xcf, 0xaa, 0x71, 0xd1, 0xe4, 0x46, 0x44, 0x2b, 0x44, 0xed, 0x03, 0x4a, 0xe7, 0x61, 0x7a, 0x99,
	0x55, 0xe4, 0x25, 0x1c, 0xf9, 0xf7, 0x33, 0x08, 0xe3, 0xbd, 0xc2, 0xc3, 0x61, 0x76, 0x99, 0x55,
	0xab, 0x39, 0x1c, 0x9a, 0x90, 0xcd, 0xe9, 0xb7, 0x08, 0x1e, 0xdc, 0x88, 0xcd, 0xb6, 0xd8, 0x58,
	0x41, 0xee, 0xed, 0x3a, 0x9b, 0xfb, 0x36, 0xde, 0x42, 0x7c, 0xed, 0xf0, 0xfc, 0x3f, 0xab, 0x08,
	0x4c, 0xba, 0xcc, 0xaa, 0xf5, 0x24, 0x1d, 0x38, 0xc9, 0x01, 0xc4, 0x83, 0x19, 0x80, 0x23, 0x33,
	0x9e, 0x6e, 0x67, 0xde, 0xf5, 0xf3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x88, 0x7c, 0xa5,
	0x16, 0x03, 0x00, 0x00,
}
