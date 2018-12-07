// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stream_call.proto

package component

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StreamCallRequest struct {
	// Types that are valid to be assigned to Request:
	//	*StreamCallRequest_Config
	//	*StreamCallRequest_Data
	Request              isStreamCallRequest_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *StreamCallRequest) Reset()         { *m = StreamCallRequest{} }
func (m *StreamCallRequest) String() string { return proto.CompactTextString(m) }
func (*StreamCallRequest) ProtoMessage()    {}
func (*StreamCallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_call_2c4480fd8af519df, []int{0}
}
func (m *StreamCallRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamCallRequest.Unmarshal(m, b)
}
func (m *StreamCallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamCallRequest.Marshal(b, m, deterministic)
}
func (dst *StreamCallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamCallRequest.Merge(dst, src)
}
func (m *StreamCallRequest) XXX_Size() int {
	return xxx_messageInfo_StreamCallRequest.Size(m)
}
func (m *StreamCallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamCallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamCallRequest proto.InternalMessageInfo

type isStreamCallRequest_Request interface {
	isStreamCallRequest_Request()
}

type StreamCallRequest_Config struct {
	Config *StreamCallConfigRequest `protobuf:"bytes,1,opt,name=config,oneof"`
}
type StreamCallRequest_Data struct {
	Data *StreamCallDataRequest `protobuf:"bytes,21,opt,name=data,oneof"`
}

func (*StreamCallRequest_Config) isStreamCallRequest_Request() {}
func (*StreamCallRequest_Data) isStreamCallRequest_Request()   {}

func (m *StreamCallRequest) GetRequest() isStreamCallRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *StreamCallRequest) GetConfig() *StreamCallConfigRequest {
	if x, ok := m.GetRequest().(*StreamCallRequest_Config); ok {
		return x.Config
	}
	return nil
}

func (m *StreamCallRequest) GetData() *StreamCallDataRequest {
	if x, ok := m.GetRequest().(*StreamCallRequest_Data); ok {
		return x.Data
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StreamCallRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StreamCallRequest_OneofMarshaler, _StreamCallRequest_OneofUnmarshaler, _StreamCallRequest_OneofSizer, []interface{}{
		(*StreamCallRequest_Config)(nil),
		(*StreamCallRequest_Data)(nil),
	}
}

func _StreamCallRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StreamCallRequest)
	// request
	switch x := m.Request.(type) {
	case *StreamCallRequest_Config:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Config); err != nil {
			return err
		}
	case *StreamCallRequest_Data:
		b.EncodeVarint(21<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Data); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StreamCallRequest.Request has unexpected type %T", x)
	}
	return nil
}

func _StreamCallRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StreamCallRequest)
	switch tag {
	case 1: // request.config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamCallConfigRequest)
		err := b.DecodeMessage(msg)
		m.Request = &StreamCallRequest_Config{msg}
		return true, err
	case 21: // request.data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamCallDataRequest)
		err := b.DecodeMessage(msg)
		m.Request = &StreamCallRequest_Data{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StreamCallRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StreamCallRequest)
	// request
	switch x := m.Request.(type) {
	case *StreamCallRequest_Config:
		s := proto.Size(x.Config)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamCallRequest_Data:
		s := proto.Size(x.Data)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type StreamCallConfigRequest struct {
	Method               *wrappers.StringValue `protobuf:"bytes,1,opt,name=method" json:"method,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *StreamCallConfigRequest) Reset()         { *m = StreamCallConfigRequest{} }
func (m *StreamCallConfigRequest) String() string { return proto.CompactTextString(m) }
func (*StreamCallConfigRequest) ProtoMessage()    {}
func (*StreamCallConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_call_2c4480fd8af519df, []int{1}
}
func (m *StreamCallConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamCallConfigRequest.Unmarshal(m, b)
}
func (m *StreamCallConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamCallConfigRequest.Marshal(b, m, deterministic)
}
func (dst *StreamCallConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamCallConfigRequest.Merge(dst, src)
}
func (m *StreamCallConfigRequest) XXX_Size() int {
	return xxx_messageInfo_StreamCallConfigRequest.Size(m)
}
func (m *StreamCallConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamCallConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamCallConfigRequest proto.InternalMessageInfo

func (m *StreamCallConfigRequest) GetMethod() *wrappers.StringValue {
	if m != nil {
		return m.Method
	}
	return nil
}

type StreamCallDataRequest struct {
	Value                *any.Any `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamCallDataRequest) Reset()         { *m = StreamCallDataRequest{} }
func (m *StreamCallDataRequest) String() string { return proto.CompactTextString(m) }
func (*StreamCallDataRequest) ProtoMessage()    {}
func (*StreamCallDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_call_2c4480fd8af519df, []int{2}
}
func (m *StreamCallDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamCallDataRequest.Unmarshal(m, b)
}
func (m *StreamCallDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamCallDataRequest.Marshal(b, m, deterministic)
}
func (dst *StreamCallDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamCallDataRequest.Merge(dst, src)
}
func (m *StreamCallDataRequest) XXX_Size() int {
	return xxx_messageInfo_StreamCallDataRequest.Size(m)
}
func (m *StreamCallDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamCallDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamCallDataRequest proto.InternalMessageInfo

func (m *StreamCallDataRequest) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

type StreamCallResponse struct {
	// Types that are valid to be assigned to Response:
	//	*StreamCallResponse_Config
	//	*StreamCallResponse_Data
	Response             isStreamCallResponse_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *StreamCallResponse) Reset()         { *m = StreamCallResponse{} }
func (m *StreamCallResponse) String() string { return proto.CompactTextString(m) }
func (*StreamCallResponse) ProtoMessage()    {}
func (*StreamCallResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_call_2c4480fd8af519df, []int{3}
}
func (m *StreamCallResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamCallResponse.Unmarshal(m, b)
}
func (m *StreamCallResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamCallResponse.Marshal(b, m, deterministic)
}
func (dst *StreamCallResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamCallResponse.Merge(dst, src)
}
func (m *StreamCallResponse) XXX_Size() int {
	return xxx_messageInfo_StreamCallResponse.Size(m)
}
func (m *StreamCallResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamCallResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamCallResponse proto.InternalMessageInfo

type isStreamCallResponse_Response interface {
	isStreamCallResponse_Response()
}

type StreamCallResponse_Config struct {
	Config *StreamCallConfigResponse `protobuf:"bytes,1,opt,name=config,oneof"`
}
type StreamCallResponse_Data struct {
	Data *StreamCallDataResponse `protobuf:"bytes,21,opt,name=data,oneof"`
}

func (*StreamCallResponse_Config) isStreamCallResponse_Response() {}
func (*StreamCallResponse_Data) isStreamCallResponse_Response()   {}

func (m *StreamCallResponse) GetResponse() isStreamCallResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *StreamCallResponse) GetConfig() *StreamCallConfigResponse {
	if x, ok := m.GetResponse().(*StreamCallResponse_Config); ok {
		return x.Config
	}
	return nil
}

func (m *StreamCallResponse) GetData() *StreamCallDataResponse {
	if x, ok := m.GetResponse().(*StreamCallResponse_Data); ok {
		return x.Data
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StreamCallResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StreamCallResponse_OneofMarshaler, _StreamCallResponse_OneofUnmarshaler, _StreamCallResponse_OneofSizer, []interface{}{
		(*StreamCallResponse_Config)(nil),
		(*StreamCallResponse_Data)(nil),
	}
}

func _StreamCallResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StreamCallResponse)
	// response
	switch x := m.Response.(type) {
	case *StreamCallResponse_Config:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Config); err != nil {
			return err
		}
	case *StreamCallResponse_Data:
		b.EncodeVarint(21<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Data); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StreamCallResponse.Response has unexpected type %T", x)
	}
	return nil
}

func _StreamCallResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StreamCallResponse)
	switch tag {
	case 1: // response.config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamCallConfigResponse)
		err := b.DecodeMessage(msg)
		m.Response = &StreamCallResponse_Config{msg}
		return true, err
	case 21: // response.data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamCallDataResponse)
		err := b.DecodeMessage(msg)
		m.Response = &StreamCallResponse_Data{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StreamCallResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StreamCallResponse)
	// response
	switch x := m.Response.(type) {
	case *StreamCallResponse_Config:
		s := proto.Size(x.Config)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamCallResponse_Data:
		s := proto.Size(x.Data)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type StreamCallConfigResponse struct {
	Method               string   `protobuf:"bytes,1,opt,name=method" json:"method,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamCallConfigResponse) Reset()         { *m = StreamCallConfigResponse{} }
func (m *StreamCallConfigResponse) String() string { return proto.CompactTextString(m) }
func (*StreamCallConfigResponse) ProtoMessage()    {}
func (*StreamCallConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_call_2c4480fd8af519df, []int{4}
}
func (m *StreamCallConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamCallConfigResponse.Unmarshal(m, b)
}
func (m *StreamCallConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamCallConfigResponse.Marshal(b, m, deterministic)
}
func (dst *StreamCallConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamCallConfigResponse.Merge(dst, src)
}
func (m *StreamCallConfigResponse) XXX_Size() int {
	return xxx_messageInfo_StreamCallConfigResponse.Size(m)
}
func (m *StreamCallConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamCallConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamCallConfigResponse proto.InternalMessageInfo

func (m *StreamCallConfigResponse) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

type StreamCallDataResponse struct {
	Value                *any.Any `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamCallDataResponse) Reset()         { *m = StreamCallDataResponse{} }
func (m *StreamCallDataResponse) String() string { return proto.CompactTextString(m) }
func (*StreamCallDataResponse) ProtoMessage()    {}
func (*StreamCallDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_call_2c4480fd8af519df, []int{5}
}
func (m *StreamCallDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamCallDataResponse.Unmarshal(m, b)
}
func (m *StreamCallDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamCallDataResponse.Marshal(b, m, deterministic)
}
func (dst *StreamCallDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamCallDataResponse.Merge(dst, src)
}
func (m *StreamCallDataResponse) XXX_Size() int {
	return xxx_messageInfo_StreamCallDataResponse.Size(m)
}
func (m *StreamCallDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamCallDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamCallDataResponse proto.InternalMessageInfo

func (m *StreamCallDataResponse) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamCallRequest)(nil), "ai.metathings.component.StreamCallRequest")
	proto.RegisterType((*StreamCallConfigRequest)(nil), "ai.metathings.component.StreamCallConfigRequest")
	proto.RegisterType((*StreamCallDataRequest)(nil), "ai.metathings.component.StreamCallDataRequest")
	proto.RegisterType((*StreamCallResponse)(nil), "ai.metathings.component.StreamCallResponse")
	proto.RegisterType((*StreamCallConfigResponse)(nil), "ai.metathings.component.StreamCallConfigResponse")
	proto.RegisterType((*StreamCallDataResponse)(nil), "ai.metathings.component.StreamCallDataResponse")
}

func init() { proto.RegisterFile("stream_call.proto", fileDescriptor_stream_call_2c4480fd8af519df) }

var fileDescriptor_stream_call_2c4480fd8af519df = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x4d, 0x4f, 0xea, 0x40,
	0x14, 0x86, 0x6f, 0x6f, 0xee, 0xed, 0xbd, 0x1c, 0x57, 0x4c, 0xe4, 0x43, 0x62, 0x94, 0x74, 0x65,
	0x4c, 0x98, 0x2a, 0x26, 0xae, 0xdc, 0x08, 0x98, 0x18, 0xd9, 0x61, 0xa2, 0x4b, 0x73, 0x80, 0xa1,
	0x34, 0xb6, 0x33, 0xb5, 0x33, 0x85, 0xf0, 0xaf, 0x5c, 0xf8, 0x7f, 0x4c, 0xfc, 0x25, 0x86, 0x99,
	0xb1, 0x7c, 0x08, 0x09, 0xee, 0x26, 0xe7, 0xcc, 0xfb, 0x9c, 0x37, 0x0f, 0x14, 0xa5, 0x4a, 0x19,
	0xc6, 0x4f, 0x03, 0x8c, 0x22, 0x9a, 0xa4, 0x42, 0x09, 0x52, 0xc1, 0x90, 0xc6, 0x4c, 0xa1, 0x1a,
	0x87, 0x3c, 0x90, 0x74, 0x20, 0xe2, 0x44, 0x70, 0xc6, 0x55, 0xed, 0x28, 0x10, 0x22, 0x88, 0x98,
	0xaf, 0xbf, 0xf5, 0xb3, 0x91, 0x3f, 0x4d, 0x31, 0x49, 0x58, 0x2a, 0x4d, 0xb0, 0x76, 0xb0, 0xbe,
	0x47, 0x3e, 0xb3, 0xab, 0xcb, 0x20, 0x54, 0xe3, 0xac, 0x3f, 0x87, 0xf9, 0xf1, 0x34, 0x54, 0xcf,
	0x62, 0xea, 0x07, 0xa2, 0xa1, 0x97, 0x8d, 0x09, 0x46, 0xe1, 0x10, 0x95, 0x48, 0xa5, 0x9f, 0x3f,
	0x4d, 0xce, 0x7b, 0x75, 0xa0, 0x78, 0xaf, 0x1b, 0xb6, 0x31, 0x8a, 0x7a, 0xec, 0x25, 0x63, 0x52,
	0x91, 0x3b, 0x70, 0x07, 0x82, 0x8f, 0xc2, 0xa0, 0xea, 0xd4, 0x9d, 0x93, 0xbd, 0xe6, 0x19, 0xdd,
	0x52, 0x99, 0x2e, 0xb2, 0x6d, 0x1d, 0xb0, 0x84, 0xdb, 0x5f, 0x3d, 0x4b, 0x20, 0x1d, 0xf8, 0x33,
	0x44, 0x85, 0xd5, 0x92, 0x26, 0xd1, 0x1d, 0x48, 0x1d, 0x54, 0xb8, 0xe0, 0xe8, 0x74, 0xab, 0x00,
	0xff, 0x52, 0x33, 0xf2, 0x1e, 0xa1, 0xb2, 0xe5, 0x2a, 0xb9, 0x02, 0x37, 0x66, 0x6a, 0x2c, 0x86,
	0xb6, 0xf7, 0x21, 0x35, 0xc6, 0xe8, 0x97, 0xb1, 0xf9, 0x95, 0x90, 0x07, 0x0f, 0x18, 0x65, 0xac,
	0xe5, 0x7e, 0xbc, 0x1f, 0xff, 0xae, 0x3b, 0x3d, 0x9b, 0xf1, 0xba, 0x50, 0xda, 0x58, 0x82, 0x34,
	0xe1, 0xef, 0x64, 0x9e, 0xb0, 0xd4, 0xfd, 0x6f, 0xd4, 0x6b, 0x3e, 0xcb, 0x69, 0xe6, 0xab, 0xf7,
	0xe6, 0x00, 0x59, 0x16, 0x2b, 0x13, 0xc1, 0x25, 0x23, 0xdd, 0x35, 0xb3, 0xe7, 0x3f, 0x30, 0x6b,
	0x10, 0x4b, 0x6a, 0x6f, 0x56, 0xd4, 0xfa, 0x3b, 0xab, 0xcd, 0x41, 0xc6, 0x2d, 0xc0, 0xff, 0xd4,
	0xce, 0xbc, 0x26, 0x54, 0xb7, 0x1d, 0x26, 0xe5, 0x15, 0xbb, 0x85, 0xdc, 0x5b, 0x07, 0xca, 0x9b,
	0x2f, 0x90, 0xd3, 0x1d, 0xc4, 0x59, 0x61, 0x7d, 0x57, 0x0f, 0x2f, 0x3e, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xa8, 0x8d, 0xeb, 0x68, 0x31, 0x03, 0x00, 0x00,
}