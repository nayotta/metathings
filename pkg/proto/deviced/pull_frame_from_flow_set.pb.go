// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pull_frame_from_flow_set.proto

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

type PullFrameFromFlowSetRequest struct {
	Id *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Request:
	//	*PullFrameFromFlowSetRequest_Config_
	Request              isPullFrameFromFlowSetRequest_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *PullFrameFromFlowSetRequest) Reset()         { *m = PullFrameFromFlowSetRequest{} }
func (m *PullFrameFromFlowSetRequest) String() string { return proto.CompactTextString(m) }
func (*PullFrameFromFlowSetRequest) ProtoMessage()    {}
func (*PullFrameFromFlowSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_07ba570db6f59c09, []int{0}
}

func (m *PullFrameFromFlowSetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullFrameFromFlowSetRequest.Unmarshal(m, b)
}
func (m *PullFrameFromFlowSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullFrameFromFlowSetRequest.Marshal(b, m, deterministic)
}
func (m *PullFrameFromFlowSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullFrameFromFlowSetRequest.Merge(m, src)
}
func (m *PullFrameFromFlowSetRequest) XXX_Size() int {
	return xxx_messageInfo_PullFrameFromFlowSetRequest.Size(m)
}
func (m *PullFrameFromFlowSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PullFrameFromFlowSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PullFrameFromFlowSetRequest proto.InternalMessageInfo

func (m *PullFrameFromFlowSetRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

type isPullFrameFromFlowSetRequest_Request interface {
	isPullFrameFromFlowSetRequest_Request()
}

type PullFrameFromFlowSetRequest_Config_ struct {
	Config *PullFrameFromFlowSetRequest_Config `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

func (*PullFrameFromFlowSetRequest_Config_) isPullFrameFromFlowSetRequest_Request() {}

func (m *PullFrameFromFlowSetRequest) GetRequest() isPullFrameFromFlowSetRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *PullFrameFromFlowSetRequest) GetConfig() *PullFrameFromFlowSetRequest_Config {
	if x, ok := m.GetRequest().(*PullFrameFromFlowSetRequest_Config_); ok {
		return x.Config
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PullFrameFromFlowSetRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PullFrameFromFlowSetRequest_Config_)(nil),
	}
}

type PullFrameFromFlowSetRequest_Config struct {
	FlowSet              *OpFlowSet          `protobuf:"bytes,1,opt,name=flow_set,json=flowSet,proto3" json:"flow_set,omitempty"`
	ConfigAck            *wrappers.BoolValue `protobuf:"bytes,2,opt,name=config_ack,json=configAck,proto3" json:"config_ack,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PullFrameFromFlowSetRequest_Config) Reset()         { *m = PullFrameFromFlowSetRequest_Config{} }
func (m *PullFrameFromFlowSetRequest_Config) String() string { return proto.CompactTextString(m) }
func (*PullFrameFromFlowSetRequest_Config) ProtoMessage()    {}
func (*PullFrameFromFlowSetRequest_Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_07ba570db6f59c09, []int{0, 0}
}

func (m *PullFrameFromFlowSetRequest_Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullFrameFromFlowSetRequest_Config.Unmarshal(m, b)
}
func (m *PullFrameFromFlowSetRequest_Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullFrameFromFlowSetRequest_Config.Marshal(b, m, deterministic)
}
func (m *PullFrameFromFlowSetRequest_Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullFrameFromFlowSetRequest_Config.Merge(m, src)
}
func (m *PullFrameFromFlowSetRequest_Config) XXX_Size() int {
	return xxx_messageInfo_PullFrameFromFlowSetRequest_Config.Size(m)
}
func (m *PullFrameFromFlowSetRequest_Config) XXX_DiscardUnknown() {
	xxx_messageInfo_PullFrameFromFlowSetRequest_Config.DiscardUnknown(m)
}

var xxx_messageInfo_PullFrameFromFlowSetRequest_Config proto.InternalMessageInfo

func (m *PullFrameFromFlowSetRequest_Config) GetFlowSet() *OpFlowSet {
	if m != nil {
		return m.FlowSet
	}
	return nil
}

func (m *PullFrameFromFlowSetRequest_Config) GetConfigAck() *wrappers.BoolValue {
	if m != nil {
		return m.ConfigAck
	}
	return nil
}

type PullFrameFromFlowSetResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Response:
	//	*PullFrameFromFlowSetResponse_Ack_
	//	*PullFrameFromFlowSetResponse_Pack_
	Response             isPullFrameFromFlowSetResponse_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *PullFrameFromFlowSetResponse) Reset()         { *m = PullFrameFromFlowSetResponse{} }
func (m *PullFrameFromFlowSetResponse) String() string { return proto.CompactTextString(m) }
func (*PullFrameFromFlowSetResponse) ProtoMessage()    {}
func (*PullFrameFromFlowSetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_07ba570db6f59c09, []int{1}
}

func (m *PullFrameFromFlowSetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullFrameFromFlowSetResponse.Unmarshal(m, b)
}
func (m *PullFrameFromFlowSetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullFrameFromFlowSetResponse.Marshal(b, m, deterministic)
}
func (m *PullFrameFromFlowSetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullFrameFromFlowSetResponse.Merge(m, src)
}
func (m *PullFrameFromFlowSetResponse) XXX_Size() int {
	return xxx_messageInfo_PullFrameFromFlowSetResponse.Size(m)
}
func (m *PullFrameFromFlowSetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PullFrameFromFlowSetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PullFrameFromFlowSetResponse proto.InternalMessageInfo

func (m *PullFrameFromFlowSetResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type isPullFrameFromFlowSetResponse_Response interface {
	isPullFrameFromFlowSetResponse_Response()
}

type PullFrameFromFlowSetResponse_Ack_ struct {
	Ack *PullFrameFromFlowSetResponse_Ack `protobuf:"bytes,2,opt,name=ack,proto3,oneof"`
}

type PullFrameFromFlowSetResponse_Pack_ struct {
	Pack *PullFrameFromFlowSetResponse_Pack `protobuf:"bytes,3,opt,name=pack,proto3,oneof"`
}

func (*PullFrameFromFlowSetResponse_Ack_) isPullFrameFromFlowSetResponse_Response() {}

func (*PullFrameFromFlowSetResponse_Pack_) isPullFrameFromFlowSetResponse_Response() {}

func (m *PullFrameFromFlowSetResponse) GetResponse() isPullFrameFromFlowSetResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *PullFrameFromFlowSetResponse) GetAck() *PullFrameFromFlowSetResponse_Ack {
	if x, ok := m.GetResponse().(*PullFrameFromFlowSetResponse_Ack_); ok {
		return x.Ack
	}
	return nil
}

func (m *PullFrameFromFlowSetResponse) GetPack() *PullFrameFromFlowSetResponse_Pack {
	if x, ok := m.GetResponse().(*PullFrameFromFlowSetResponse_Pack_); ok {
		return x.Pack
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PullFrameFromFlowSetResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PullFrameFromFlowSetResponse_Ack_)(nil),
		(*PullFrameFromFlowSetResponse_Pack_)(nil),
	}
}

type PullFrameFromFlowSetResponse_Ack struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PullFrameFromFlowSetResponse_Ack) Reset()         { *m = PullFrameFromFlowSetResponse_Ack{} }
func (m *PullFrameFromFlowSetResponse_Ack) String() string { return proto.CompactTextString(m) }
func (*PullFrameFromFlowSetResponse_Ack) ProtoMessage()    {}
func (*PullFrameFromFlowSetResponse_Ack) Descriptor() ([]byte, []int) {
	return fileDescriptor_07ba570db6f59c09, []int{1, 0}
}

func (m *PullFrameFromFlowSetResponse_Ack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullFrameFromFlowSetResponse_Ack.Unmarshal(m, b)
}
func (m *PullFrameFromFlowSetResponse_Ack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullFrameFromFlowSetResponse_Ack.Marshal(b, m, deterministic)
}
func (m *PullFrameFromFlowSetResponse_Ack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullFrameFromFlowSetResponse_Ack.Merge(m, src)
}
func (m *PullFrameFromFlowSetResponse_Ack) XXX_Size() int {
	return xxx_messageInfo_PullFrameFromFlowSetResponse_Ack.Size(m)
}
func (m *PullFrameFromFlowSetResponse_Ack) XXX_DiscardUnknown() {
	xxx_messageInfo_PullFrameFromFlowSetResponse_Ack.DiscardUnknown(m)
}

var xxx_messageInfo_PullFrameFromFlowSetResponse_Ack proto.InternalMessageInfo

type PullFrameFromFlowSetResponse_Pack struct {
	Device               *Device  `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	Frames               []*Frame `protobuf:"bytes,2,rep,name=frames,proto3" json:"frames,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PullFrameFromFlowSetResponse_Pack) Reset()         { *m = PullFrameFromFlowSetResponse_Pack{} }
func (m *PullFrameFromFlowSetResponse_Pack) String() string { return proto.CompactTextString(m) }
func (*PullFrameFromFlowSetResponse_Pack) ProtoMessage()    {}
func (*PullFrameFromFlowSetResponse_Pack) Descriptor() ([]byte, []int) {
	return fileDescriptor_07ba570db6f59c09, []int{1, 1}
}

func (m *PullFrameFromFlowSetResponse_Pack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullFrameFromFlowSetResponse_Pack.Unmarshal(m, b)
}
func (m *PullFrameFromFlowSetResponse_Pack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullFrameFromFlowSetResponse_Pack.Marshal(b, m, deterministic)
}
func (m *PullFrameFromFlowSetResponse_Pack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullFrameFromFlowSetResponse_Pack.Merge(m, src)
}
func (m *PullFrameFromFlowSetResponse_Pack) XXX_Size() int {
	return xxx_messageInfo_PullFrameFromFlowSetResponse_Pack.Size(m)
}
func (m *PullFrameFromFlowSetResponse_Pack) XXX_DiscardUnknown() {
	xxx_messageInfo_PullFrameFromFlowSetResponse_Pack.DiscardUnknown(m)
}

var xxx_messageInfo_PullFrameFromFlowSetResponse_Pack proto.InternalMessageInfo

func (m *PullFrameFromFlowSetResponse_Pack) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *PullFrameFromFlowSetResponse_Pack) GetFrames() []*Frame {
	if m != nil {
		return m.Frames
	}
	return nil
}

func init() {
	proto.RegisterType((*PullFrameFromFlowSetRequest)(nil), "ai.metathings.service.deviced.PullFrameFromFlowSetRequest")
	proto.RegisterType((*PullFrameFromFlowSetRequest_Config)(nil), "ai.metathings.service.deviced.PullFrameFromFlowSetRequest.Config")
	proto.RegisterType((*PullFrameFromFlowSetResponse)(nil), "ai.metathings.service.deviced.PullFrameFromFlowSetResponse")
	proto.RegisterType((*PullFrameFromFlowSetResponse_Ack)(nil), "ai.metathings.service.deviced.PullFrameFromFlowSetResponse.Ack")
	proto.RegisterType((*PullFrameFromFlowSetResponse_Pack)(nil), "ai.metathings.service.deviced.PullFrameFromFlowSetResponse.Pack")
}

func init() { proto.RegisterFile("pull_frame_from_flow_set.proto", fileDescriptor_07ba570db6f59c09) }

var fileDescriptor_07ba570db6f59c09 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xc1, 0x6a, 0xd4, 0x40,
	0x18, 0xc7, 0x9b, 0x6c, 0x4d, 0xbb, 0xdf, 0x82, 0x87, 0x39, 0x2d, 0xb1, 0xd6, 0xa5, 0x28, 0xec,
	0xa5, 0x13, 0xa8, 0x22, 0x08, 0x8a, 0xee, 0x56, 0x96, 0xde, 0x2c, 0x59, 0xe8, 0xc5, 0x43, 0x98,
	0x4d, 0x26, 0xe9, 0x90, 0x49, 0xbe, 0x38, 0x33, 0x69, 0x1e, 0xc0, 0x8b, 0x4f, 0xe6, 0x6b, 0x08,
	0x3e, 0x83, 0x0f, 0x20, 0x99, 0xc9, 0xd6, 0x83, 0xb2, 0x0b, 0xf6, 0x34, 0x93, 0xc9, 0xfc, 0x7f,
	0xdf, 0xef, 0xfb, 0x42, 0xe0, 0xb4, 0x69, 0xa5, 0x4c, 0x72, 0xc5, 0x2a, 0x9e, 0xe4, 0x0a, 0xab,
	0x24, 0x97, 0xd8, 0x25, 0x9a, 0x1b, 0xda, 0x28, 0x34, 0x48, 0x9e, 0x32, 0x41, 0x2b, 0x6e, 0x98,
	0xb9, 0x15, 0x75, 0xa1, 0xa9, 0xe6, 0xea, 0x4e, 0xa4, 0x9c, 0x66, 0xbc, 0x5f, 0xb2, 0xf0, 0xb4,
	0x40, 0x2c, 0x24, 0x8f, 0xec, 0xe5, 0x4d, 0x9b, 0x47, 0x9d, 0x62, 0x4d, 0xc3, 0x95, 0x76, 0xf1,
	0xf0, 0x75, 0x21, 0xcc, 0x6d, 0xbb, 0xa1, 0x29, 0x56, 0x51, 0xd5, 0x09, 0x53, 0x62, 0x17, 0x15,
	0x78, 0x6e, 0x5f, 0x9e, 0xdf, 0x31, 0x29, 0x32, 0x66, 0x50, 0xe9, 0xe8, 0x7e, 0x3b, 0xe4, 0x26,
	0x15, 0x66, 0x5c, 0xba, 0x87, 0xb3, 0xef, 0x3e, 0x3c, 0xb9, 0x6e, 0xa5, 0x5c, 0xf5, 0x96, 0x2b,
	0x85, 0xd5, 0x4a, 0x62, 0xb7, 0xe6, 0x26, 0xe6, 0x5f, 0x5a, 0xae, 0x0d, 0x79, 0x05, 0xbe, 0xc8,
	0xa6, 0xde, 0xcc, 0x9b, 0x4f, 0x2e, 0x4e, 0xa8, 0x33, 0xa2, 0x5b, 0x23, 0xba, 0x36, 0x4a, 0xd4,
	0xc5, 0x0d, 0x93, 0x2d, 0x5f, 0x06, 0x3f, 0x7f, 0x3c, 0xf3, 0x67, 0x5e, 0xec, 0x8b, 0x8c, 0x7c,
	0x86, 0x20, 0xc5, 0x3a, 0x17, 0xc5, 0xd4, 0xb7, 0xc9, 0x05, 0xdd, 0xd9, 0x2a, 0xdd, 0x61, 0x40,
	0x2f, 0x2d, 0xe8, 0xea, 0x20, 0x1e, 0x90, 0xe1, 0x37, 0x0f, 0x02, 0x77, 0x48, 0x2e, 0xe1, 0x78,
	0x3b, 0xd3, 0xc1, 0x71, 0xbe, 0xa7, 0xd2, 0xa7, 0x66, 0x8b, 0x3f, 0xca, 0xdd, 0x86, 0xbc, 0x01,
	0x70, 0xe4, 0x84, 0xa5, 0xe5, 0x20, 0x1c, 0xfe, 0xd5, 0xea, 0x12, 0x51, 0xda, 0x46, 0xe3, 0xb1,
	0xbb, 0xbd, 0x48, 0xcb, 0xe5, 0x18, 0x8e, 0x94, 0xd3, 0x3c, 0xfb, 0xe5, 0xc3, 0xc9, 0xbf, 0xdb,
	0xd0, 0x0d, 0xd6, 0x9a, 0x93, 0xc7, 0xf7, 0x93, 0x1c, 0xdb, 0x19, 0xad, 0x61, 0xf4, 0xa7, 0xde,
	0xfb, 0xff, 0x1a, 0x90, 0x23, 0xd3, 0x45, 0x5a, 0x5e, 0x1d, 0xc4, 0x3d, 0x8d, 0xdc, 0xc0, 0x61,
	0xd3, 0x53, 0x47, 0x96, 0xfa, 0xe1, 0x21, 0xd4, 0x6b, 0x66, 0xb1, 0x96, 0x17, 0x3e, 0x82, 0xd1,
	0x22, 0x2d, 0xc3, 0xaf, 0x1e, 0x1c, 0xf6, 0xe7, 0xe4, 0x1d, 0x04, 0x0e, 0x32, 0x8c, 0xfd, 0xc5,
	0x9e, 0x4a, 0x1f, 0xed, 0x1a, 0x0f, 0x21, 0xf2, 0x16, 0x02, 0xfb, 0x5b, 0xe8, 0xa9, 0x3f, 0x1b,
	0xcd, 0x27, 0x17, 0xcf, 0xf7, 0xc4, 0xad, 0x64, 0x3c, 0x64, 0x96, 0x00, 0xc7, 0x6a, 0xb0, 0xdc,
	0x04, 0xf6, 0x03, 0xbd, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x0b, 0x9d, 0xa9, 0x6c, 0x03,
	0x00, 0x00,
}
