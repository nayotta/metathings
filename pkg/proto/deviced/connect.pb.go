// Code generated by protoc-gen-go. DO NOT EDIT.
// source: connect.proto

package deviced

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/mwitkow/go-proto-validators"
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

type ConnectMessageKind int32

const (
	ConnectMessageKind_CONNECT_MESSAGE_KIND_UNKNOWN ConnectMessageKind = 0
	ConnectMessageKind_CONNECT_MESSAGE_KIND_SYSTEM  ConnectMessageKind = 1
	ConnectMessageKind_CONNECT_MESSAGE_KIND_USER    ConnectMessageKind = 2
)

var ConnectMessageKind_name = map[int32]string{
	0: "CONNECT_MESSAGE_KIND_UNKNOWN",
	1: "CONNECT_MESSAGE_KIND_SYSTEM",
	2: "CONNECT_MESSAGE_KIND_USER",
}

var ConnectMessageKind_value = map[string]int32{
	"CONNECT_MESSAGE_KIND_UNKNOWN": 0,
	"CONNECT_MESSAGE_KIND_SYSTEM":  1,
	"CONNECT_MESSAGE_KIND_USER":    2,
}

func (x ConnectMessageKind) String() string {
	return proto.EnumName(ConnectMessageKind_name, int32(x))
}

func (ConnectMessageKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_778b7e3040344da6, []int{0}
}

type ConnectResponse struct {
	SessionId int64              `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Kind      ConnectMessageKind `protobuf:"varint,2,opt,name=kind,proto3,enum=ai.metathings.service.deviced.ConnectMessageKind" json:"kind,omitempty"`
	// Types that are valid to be assigned to Union:
	//	*ConnectResponse_UnaryCall
	//	*ConnectResponse_StreamCall
	//	*ConnectResponse_Err
	Union                isConnectResponse_Union `protobuf_oneof:"union"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConnectResponse) Reset()         { *m = ConnectResponse{} }
func (m *ConnectResponse) String() string { return proto.CompactTextString(m) }
func (*ConnectResponse) ProtoMessage()    {}
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_778b7e3040344da6, []int{0}
}

func (m *ConnectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectResponse.Unmarshal(m, b)
}
func (m *ConnectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectResponse.Marshal(b, m, deterministic)
}
func (m *ConnectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectResponse.Merge(m, src)
}
func (m *ConnectResponse) XXX_Size() int {
	return xxx_messageInfo_ConnectResponse.Size(m)
}
func (m *ConnectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectResponse proto.InternalMessageInfo

func (m *ConnectResponse) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *ConnectResponse) GetKind() ConnectMessageKind {
	if m != nil {
		return m.Kind
	}
	return ConnectMessageKind_CONNECT_MESSAGE_KIND_UNKNOWN
}

type isConnectResponse_Union interface {
	isConnectResponse_Union()
}

type ConnectResponse_UnaryCall struct {
	UnaryCall *UnaryCallValue `protobuf:"bytes,3,opt,name=unary_call,json=unaryCall,proto3,oneof"`
}

type ConnectResponse_StreamCall struct {
	StreamCall *StreamCallValue `protobuf:"bytes,4,opt,name=stream_call,json=streamCall,proto3,oneof"`
}

type ConnectResponse_Err struct {
	Err *ErrorValue `protobuf:"bytes,9,opt,name=err,proto3,oneof"`
}

func (*ConnectResponse_UnaryCall) isConnectResponse_Union() {}

func (*ConnectResponse_StreamCall) isConnectResponse_Union() {}

func (*ConnectResponse_Err) isConnectResponse_Union() {}

func (m *ConnectResponse) GetUnion() isConnectResponse_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (m *ConnectResponse) GetUnaryCall() *UnaryCallValue {
	if x, ok := m.GetUnion().(*ConnectResponse_UnaryCall); ok {
		return x.UnaryCall
	}
	return nil
}

func (m *ConnectResponse) GetStreamCall() *StreamCallValue {
	if x, ok := m.GetUnion().(*ConnectResponse_StreamCall); ok {
		return x.StreamCall
	}
	return nil
}

func (m *ConnectResponse) GetErr() *ErrorValue {
	if x, ok := m.GetUnion().(*ConnectResponse_Err); ok {
		return x.Err
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ConnectResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ConnectResponse_UnaryCall)(nil),
		(*ConnectResponse_StreamCall)(nil),
		(*ConnectResponse_Err)(nil),
	}
}

type ConnectRequest struct {
	SessionId *wrappers.Int64Value `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Kind      ConnectMessageKind   `protobuf:"varint,2,opt,name=kind,proto3,enum=ai.metathings.service.deviced.ConnectMessageKind" json:"kind,omitempty"`
	// Types that are valid to be assigned to Union:
	//	*ConnectRequest_UnaryCall
	//	*ConnectRequest_StreamCall
	Union                isConnectRequest_Union `protobuf_oneof:"union"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ConnectRequest) Reset()         { *m = ConnectRequest{} }
func (m *ConnectRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectRequest) ProtoMessage()    {}
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_778b7e3040344da6, []int{1}
}

func (m *ConnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectRequest.Unmarshal(m, b)
}
func (m *ConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectRequest.Marshal(b, m, deterministic)
}
func (m *ConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectRequest.Merge(m, src)
}
func (m *ConnectRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectRequest.Size(m)
}
func (m *ConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectRequest proto.InternalMessageInfo

func (m *ConnectRequest) GetSessionId() *wrappers.Int64Value {
	if m != nil {
		return m.SessionId
	}
	return nil
}

func (m *ConnectRequest) GetKind() ConnectMessageKind {
	if m != nil {
		return m.Kind
	}
	return ConnectMessageKind_CONNECT_MESSAGE_KIND_UNKNOWN
}

type isConnectRequest_Union interface {
	isConnectRequest_Union()
}

type ConnectRequest_UnaryCall struct {
	UnaryCall *OpUnaryCallValue `protobuf:"bytes,3,opt,name=unary_call,json=unaryCall,proto3,oneof"`
}

type ConnectRequest_StreamCall struct {
	StreamCall *OpStreamCallValue `protobuf:"bytes,4,opt,name=stream_call,json=streamCall,proto3,oneof"`
}

func (*ConnectRequest_UnaryCall) isConnectRequest_Union() {}

func (*ConnectRequest_StreamCall) isConnectRequest_Union() {}

func (m *ConnectRequest) GetUnion() isConnectRequest_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (m *ConnectRequest) GetUnaryCall() *OpUnaryCallValue {
	if x, ok := m.GetUnion().(*ConnectRequest_UnaryCall); ok {
		return x.UnaryCall
	}
	return nil
}

func (m *ConnectRequest) GetStreamCall() *OpStreamCallValue {
	if x, ok := m.GetUnion().(*ConnectRequest_StreamCall); ok {
		return x.StreamCall
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ConnectRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ConnectRequest_UnaryCall)(nil),
		(*ConnectRequest_StreamCall)(nil),
	}
}

func init() {
	proto.RegisterEnum("ai.metathings.service.deviced.ConnectMessageKind", ConnectMessageKind_name, ConnectMessageKind_value)
	proto.RegisterType((*ConnectResponse)(nil), "ai.metathings.service.deviced.ConnectResponse")
	proto.RegisterType((*ConnectRequest)(nil), "ai.metathings.service.deviced.ConnectRequest")
}

func init() { proto.RegisterFile("connect.proto", fileDescriptor_778b7e3040344da6) }

var fileDescriptor_778b7e3040344da6 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x97, 0x74, 0x0c, 0xd5, 0x15, 0xa3, 0xf2, 0xa9, 0x6c, 0x94, 0x45, 0x3b, 0x15, 0xa4,
	0x3a, 0x50, 0xd0, 0x6e, 0x1c, 0x68, 0x89, 0xa0, 0xaa, 0x9a, 0x42, 0xb2, 0x81, 0x38, 0x45, 0x6e,
	0xf2, 0xc8, 0xac, 0x25, 0x76, 0xb0, 0x9d, 0x56, 0x7c, 0x4e, 0x3e, 0x00, 0xd2, 0x3e, 0x09, 0x6a,
	0xd2, 0x46, 0x93, 0x1a, 0x35, 0x07, 0x4e, 0x89, 0x9f, 0xdf, 0xff, 0x27, 0xbf, 0xf7, 0x43, 0x4f,
	0x42, 0xc1, 0x39, 0x84, 0x9a, 0x64, 0x52, 0x68, 0x81, 0xfb, 0x94, 0x91, 0x14, 0x34, 0xd5, 0xb7,
	0x8c, 0xc7, 0x8a, 0x28, 0x90, 0x2b, 0x16, 0x02, 0x89, 0x60, 0xf3, 0x89, 0xce, 0x5e, 0xc4, 0x42,
	0xc4, 0x09, 0xd8, 0x45, 0xf3, 0x32, 0xff, 0x69, 0xaf, 0x25, 0xcd, 0x32, 0x90, 0xaa, 0x8c, 0x9f,
	0x5d, 0xc5, 0x4c, 0xdf, 0xe6, 0x4b, 0x12, 0x8a, 0xd4, 0x4e, 0xd7, 0x4c, 0xdf, 0x89, 0xb5, 0x1d,
	0x8b, 0x61, 0x71, 0x39, 0x5c, 0xd1, 0x84, 0x45, 0x54, 0x0b, 0xa9, 0xec, 0xea, 0x77, 0x9b, 0xeb,
	0xa4, 0x22, 0x82, 0xa4, 0x3c, 0x5c, 0xde, 0x9b, 0xe8, 0xe9, 0xa4, 0x7c, 0x95, 0x07, 0x2a, 0x13,
	0x5c, 0x01, 0xee, 0x23, 0xa4, 0x40, 0x29, 0x26, 0x78, 0xc0, 0xa2, 0x9e, 0x61, 0x19, 0x83, 0x96,
	0xd7, 0xde, 0x56, 0xa6, 0x11, 0x76, 0xd0, 0xf1, 0x1d, 0xe3, 0x51, 0xcf, 0xb4, 0x8c, 0xc1, 0xe9,
	0xe8, 0x0d, 0x39, 0x38, 0x05, 0xd9, 0xc2, 0xe7, 0xa0, 0x14, 0x8d, 0x61, 0xc6, 0x78, 0xe4, 0x15,
	0x71, 0xec, 0x22, 0x94, 0x73, 0x2a, 0x7f, 0x07, 0x21, 0x4d, 0x92, 0x5e, 0xcb, 0x32, 0x06, 0x9d,
	0xd1, 0xb0, 0x01, 0x76, 0xb3, 0x09, 0x4c, 0x68, 0x92, 0x7c, 0xa3, 0x49, 0x0e, 0x9f, 0x8f, 0xbc,
	0x76, 0xbe, 0xab, 0xe0, 0xaf, 0xa8, 0xa3, 0xb4, 0x04, 0x9a, 0x96, 0xc0, 0xe3, 0x02, 0x48, 0x1a,
	0x80, 0x7e, 0x91, 0x78, 0x48, 0x44, 0xaa, 0x2a, 0xe1, 0xf7, 0xa8, 0x05, 0x52, 0xf6, 0xda, 0x05,
	0xea, 0x65, 0x03, 0xca, 0x91, 0x52, 0xc8, 0x1d, 0x65, 0x93, 0x1b, 0x3f, 0x46, 0x8f, 0x72, 0xce,
	0x04, 0xbf, 0xfc, 0x63, 0xa2, 0xd3, 0x6a, 0xc9, 0xbf, 0x72, 0x50, 0x1a, 0x8f, 0xf7, 0x76, 0xdc,
	0x19, 0x9d, 0x93, 0xd2, 0x38, 0xd9, 0x19, 0x27, 0x53, 0xae, 0xaf, 0xde, 0x15, 0xcc, 0xf1, 0xc9,
	0xfd, 0xdf, 0x0b, 0xd3, 0x32, 0x1e, 0x8a, 0x98, 0xff, 0xa7, 0x88, 0x8a, 0x59, 0x0a, 0xf9, 0x52,
	0x23, 0xc4, 0x6e, 0x80, 0x2e, 0xb2, 0x43, 0x4a, 0xfc, 0x3a, 0x25, 0xaf, 0x1b, 0x91, 0x07, 0xa5,
	0x54, 0x5b, 0x7d, 0xb5, 0x42, 0x78, 0x7f, 0x26, 0x6c, 0xa1, 0xe7, 0x93, 0x85, 0xeb, 0x3a, 0x93,
	0xeb, 0x60, 0xee, 0xf8, 0xfe, 0x87, 0x4f, 0x4e, 0x30, 0x9b, 0xba, 0x1f, 0x83, 0x1b, 0x77, 0xe6,
	0x2e, 0xbe, 0xbb, 0xdd, 0x23, 0x7c, 0x81, 0xce, 0x6b, 0x3b, 0xfc, 0x1f, 0xfe, 0xb5, 0x33, 0xef,
	0x1a, 0xb8, 0x8f, 0x9e, 0xd5, 0x23, 0x7c, 0xc7, 0xeb, 0x9a, 0xcb, 0x93, 0x42, 0xcf, 0xdb, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x25, 0x2d, 0xe8, 0x30, 0xce, 0x03, 0x00, 0x00,
}
