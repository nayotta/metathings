// Code generated by protoc-gen-go. DO NOT EDIT.
// source: put_object_streaming.proto

package deviced

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type PutObjectStreamingRequest struct {
	Id *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Request:
	//	*PutObjectStreamingRequest_Metadata_
	//	*PutObjectStreamingRequest_Chunks
	//	*PutObjectStreamingRequest_Ack_
	Request              isPutObjectStreamingRequest_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *PutObjectStreamingRequest) Reset()         { *m = PutObjectStreamingRequest{} }
func (m *PutObjectStreamingRequest) String() string { return proto.CompactTextString(m) }
func (*PutObjectStreamingRequest) ProtoMessage()    {}
func (*PutObjectStreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_662424e73095d172, []int{0}
}

func (m *PutObjectStreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutObjectStreamingRequest.Unmarshal(m, b)
}
func (m *PutObjectStreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutObjectStreamingRequest.Marshal(b, m, deterministic)
}
func (m *PutObjectStreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutObjectStreamingRequest.Merge(m, src)
}
func (m *PutObjectStreamingRequest) XXX_Size() int {
	return xxx_messageInfo_PutObjectStreamingRequest.Size(m)
}
func (m *PutObjectStreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutObjectStreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutObjectStreamingRequest proto.InternalMessageInfo

func (m *PutObjectStreamingRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

type isPutObjectStreamingRequest_Request interface {
	isPutObjectStreamingRequest_Request()
}

type PutObjectStreamingRequest_Metadata_ struct {
	Metadata *PutObjectStreamingRequest_Metadata `protobuf:"bytes,2,opt,name=metadata,proto3,oneof"`
}

type PutObjectStreamingRequest_Chunks struct {
	Chunks *OpObjectChunks `protobuf:"bytes,3,opt,name=chunks,proto3,oneof"`
}

type PutObjectStreamingRequest_Ack_ struct {
	Ack *PutObjectStreamingRequest_Ack `protobuf:"bytes,4,opt,name=ack,proto3,oneof"`
}

func (*PutObjectStreamingRequest_Metadata_) isPutObjectStreamingRequest_Request() {}

func (*PutObjectStreamingRequest_Chunks) isPutObjectStreamingRequest_Request() {}

func (*PutObjectStreamingRequest_Ack_) isPutObjectStreamingRequest_Request() {}

func (m *PutObjectStreamingRequest) GetRequest() isPutObjectStreamingRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *PutObjectStreamingRequest) GetMetadata() *PutObjectStreamingRequest_Metadata {
	if x, ok := m.GetRequest().(*PutObjectStreamingRequest_Metadata_); ok {
		return x.Metadata
	}
	return nil
}

func (m *PutObjectStreamingRequest) GetChunks() *OpObjectChunks {
	if x, ok := m.GetRequest().(*PutObjectStreamingRequest_Chunks); ok {
		return x.Chunks
	}
	return nil
}

func (m *PutObjectStreamingRequest) GetAck() *PutObjectStreamingRequest_Ack {
	if x, ok := m.GetRequest().(*PutObjectStreamingRequest_Ack_); ok {
		return x.Ack
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PutObjectStreamingRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PutObjectStreamingRequest_Metadata_)(nil),
		(*PutObjectStreamingRequest_Chunks)(nil),
		(*PutObjectStreamingRequest_Ack_)(nil),
	}
}

type PutObjectStreamingRequest_Ack struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutObjectStreamingRequest_Ack) Reset()         { *m = PutObjectStreamingRequest_Ack{} }
func (m *PutObjectStreamingRequest_Ack) String() string { return proto.CompactTextString(m) }
func (*PutObjectStreamingRequest_Ack) ProtoMessage()    {}
func (*PutObjectStreamingRequest_Ack) Descriptor() ([]byte, []int) {
	return fileDescriptor_662424e73095d172, []int{0, 0}
}

func (m *PutObjectStreamingRequest_Ack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutObjectStreamingRequest_Ack.Unmarshal(m, b)
}
func (m *PutObjectStreamingRequest_Ack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutObjectStreamingRequest_Ack.Marshal(b, m, deterministic)
}
func (m *PutObjectStreamingRequest_Ack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutObjectStreamingRequest_Ack.Merge(m, src)
}
func (m *PutObjectStreamingRequest_Ack) XXX_Size() int {
	return xxx_messageInfo_PutObjectStreamingRequest_Ack.Size(m)
}
func (m *PutObjectStreamingRequest_Ack) XXX_DiscardUnknown() {
	xxx_messageInfo_PutObjectStreamingRequest_Ack.DiscardUnknown(m)
}

var xxx_messageInfo_PutObjectStreamingRequest_Ack proto.InternalMessageInfo

type PutObjectStreamingRequest_Metadata struct {
	Object               *OpObject             `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	Sha1                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=sha1,proto3" json:"sha1,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PutObjectStreamingRequest_Metadata) Reset()         { *m = PutObjectStreamingRequest_Metadata{} }
func (m *PutObjectStreamingRequest_Metadata) String() string { return proto.CompactTextString(m) }
func (*PutObjectStreamingRequest_Metadata) ProtoMessage()    {}
func (*PutObjectStreamingRequest_Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_662424e73095d172, []int{0, 1}
}

func (m *PutObjectStreamingRequest_Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutObjectStreamingRequest_Metadata.Unmarshal(m, b)
}
func (m *PutObjectStreamingRequest_Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutObjectStreamingRequest_Metadata.Marshal(b, m, deterministic)
}
func (m *PutObjectStreamingRequest_Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutObjectStreamingRequest_Metadata.Merge(m, src)
}
func (m *PutObjectStreamingRequest_Metadata) XXX_Size() int {
	return xxx_messageInfo_PutObjectStreamingRequest_Metadata.Size(m)
}
func (m *PutObjectStreamingRequest_Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_PutObjectStreamingRequest_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_PutObjectStreamingRequest_Metadata proto.InternalMessageInfo

func (m *PutObjectStreamingRequest_Metadata) GetObject() *OpObject {
	if m != nil {
		return m.Object
	}
	return nil
}

func (m *PutObjectStreamingRequest_Metadata) GetSha1() *wrappers.StringValue {
	if m != nil {
		return m.Sha1
	}
	return nil
}

type PutObjectStreamingResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Response:
	//	*PutObjectStreamingResponse_Chunks
	//	*PutObjectStreamingResponse_Ack_
	Response             isPutObjectStreamingResponse_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *PutObjectStreamingResponse) Reset()         { *m = PutObjectStreamingResponse{} }
func (m *PutObjectStreamingResponse) String() string { return proto.CompactTextString(m) }
func (*PutObjectStreamingResponse) ProtoMessage()    {}
func (*PutObjectStreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_662424e73095d172, []int{1}
}

func (m *PutObjectStreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutObjectStreamingResponse.Unmarshal(m, b)
}
func (m *PutObjectStreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutObjectStreamingResponse.Marshal(b, m, deterministic)
}
func (m *PutObjectStreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutObjectStreamingResponse.Merge(m, src)
}
func (m *PutObjectStreamingResponse) XXX_Size() int {
	return xxx_messageInfo_PutObjectStreamingResponse.Size(m)
}
func (m *PutObjectStreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutObjectStreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutObjectStreamingResponse proto.InternalMessageInfo

func (m *PutObjectStreamingResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type isPutObjectStreamingResponse_Response interface {
	isPutObjectStreamingResponse_Response()
}

type PutObjectStreamingResponse_Chunks struct {
	Chunks *ObjectChunks `protobuf:"bytes,2,opt,name=chunks,proto3,oneof"`
}

type PutObjectStreamingResponse_Ack_ struct {
	Ack *PutObjectStreamingResponse_Ack `protobuf:"bytes,3,opt,name=ack,proto3,oneof"`
}

func (*PutObjectStreamingResponse_Chunks) isPutObjectStreamingResponse_Response() {}

func (*PutObjectStreamingResponse_Ack_) isPutObjectStreamingResponse_Response() {}

func (m *PutObjectStreamingResponse) GetResponse() isPutObjectStreamingResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *PutObjectStreamingResponse) GetChunks() *ObjectChunks {
	if x, ok := m.GetResponse().(*PutObjectStreamingResponse_Chunks); ok {
		return x.Chunks
	}
	return nil
}

func (m *PutObjectStreamingResponse) GetAck() *PutObjectStreamingResponse_Ack {
	if x, ok := m.GetResponse().(*PutObjectStreamingResponse_Ack_); ok {
		return x.Ack
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PutObjectStreamingResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PutObjectStreamingResponse_Chunks)(nil),
		(*PutObjectStreamingResponse_Ack_)(nil),
	}
}

type PutObjectStreamingResponse_Ack struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutObjectStreamingResponse_Ack) Reset()         { *m = PutObjectStreamingResponse_Ack{} }
func (m *PutObjectStreamingResponse_Ack) String() string { return proto.CompactTextString(m) }
func (*PutObjectStreamingResponse_Ack) ProtoMessage()    {}
func (*PutObjectStreamingResponse_Ack) Descriptor() ([]byte, []int) {
	return fileDescriptor_662424e73095d172, []int{1, 0}
}

func (m *PutObjectStreamingResponse_Ack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutObjectStreamingResponse_Ack.Unmarshal(m, b)
}
func (m *PutObjectStreamingResponse_Ack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutObjectStreamingResponse_Ack.Marshal(b, m, deterministic)
}
func (m *PutObjectStreamingResponse_Ack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutObjectStreamingResponse_Ack.Merge(m, src)
}
func (m *PutObjectStreamingResponse_Ack) XXX_Size() int {
	return xxx_messageInfo_PutObjectStreamingResponse_Ack.Size(m)
}
func (m *PutObjectStreamingResponse_Ack) XXX_DiscardUnknown() {
	xxx_messageInfo_PutObjectStreamingResponse_Ack.DiscardUnknown(m)
}

var xxx_messageInfo_PutObjectStreamingResponse_Ack proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PutObjectStreamingRequest)(nil), "ai.metathings.service.deviced.PutObjectStreamingRequest")
	proto.RegisterType((*PutObjectStreamingRequest_Ack)(nil), "ai.metathings.service.deviced.PutObjectStreamingRequest.Ack")
	proto.RegisterType((*PutObjectStreamingRequest_Metadata)(nil), "ai.metathings.service.deviced.PutObjectStreamingRequest.Metadata")
	proto.RegisterType((*PutObjectStreamingResponse)(nil), "ai.metathings.service.deviced.PutObjectStreamingResponse")
	proto.RegisterType((*PutObjectStreamingResponse_Ack)(nil), "ai.metathings.service.deviced.PutObjectStreamingResponse.Ack")
}

func init() { proto.RegisterFile("put_object_streaming.proto", fileDescriptor_662424e73095d172) }

var fileDescriptor_662424e73095d172 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0x86, 0x81, 0x22, 0x1f, 0x43, 0xe2, 0x61, 0x4f, 0xb5, 0x51, 0x63, 0xb8, 0x68, 0xa2, 0x2e,
	0x7e, 0x5c, 0x35, 0x06, 0x8c, 0x91, 0x8b, 0x01, 0x4b, 0xe2, 0xc1, 0x0b, 0x59, 0xda, 0xb1, 0xd4,
	0x42, 0xb7, 0xee, 0x6e, 0xf5, 0xe4, 0xdf, 0x34, 0xf1, 0xdf, 0x18, 0x76, 0x17, 0x2e, 0x22, 0x18,
	0x4e, 0x4d, 0xb3, 0xf3, 0x3e, 0xf3, 0xe6, 0x19, 0xf0, 0xb2, 0x5c, 0x0d, 0xf9, 0xe8, 0x15, 0x03,
	0x35, 0x94, 0x4a, 0x20, 0x9b, 0xc6, 0x69, 0x44, 0x33, 0xc1, 0x15, 0x27, 0x7b, 0x2c, 0xa6, 0x53,
	0x54, 0x4c, 0x8d, 0xe3, 0x34, 0x92, 0x54, 0xa2, 0x78, 0x8f, 0x03, 0xa4, 0x21, 0xce, 0x3e, 0xa1,
	0xb7, 0x1f, 0x71, 0x1e, 0x4d, 0xb0, 0xa5, 0x87, 0x47, 0xf9, 0x4b, 0xeb, 0x43, 0xb0, 0x2c, 0x43,
	0x21, 0x4d, 0xdc, 0x6b, 0x4c, 0x79, 0x88, 0x13, 0xf3, 0xd3, 0xfc, 0x72, 0x60, 0xa7, 0x9f, 0xab,
	0x9e, 0xde, 0x34, 0x98, 0x2f, 0xf2, 0xf1, 0x2d, 0x47, 0xa9, 0xc8, 0x09, 0x94, 0xe2, 0xd0, 0x2d,
	0x1e, 0x14, 0x8f, 0x1a, 0x17, 0xbb, 0xd4, 0x70, 0xe9, 0x9c, 0x4b, 0x07, 0x4a, 0xc4, 0x69, 0xf4,
	0xc4, 0x26, 0x39, 0xfa, 0xa5, 0x38, 0x24, 0x43, 0xa8, 0xcd, 0x6a, 0x85, 0x4c, 0x31, 0xb7, 0xa4,
	0x33, 0x6d, 0xba, 0xb2, 0x2a, 0xfd, 0x73, 0x33, 0x7d, 0xb0, 0xa0, 0x6e, 0xc1, 0x5f, 0x40, 0xc9,
	0x3d, 0x54, 0x82, 0x71, 0x9e, 0x26, 0xd2, 0x75, 0x34, 0xfe, 0x74, 0x0d, 0xbe, 0x97, 0x19, 0xfa,
	0xad, 0x0e, 0x75, 0x0b, 0xbe, 0x8d, 0x93, 0x3e, 0x38, 0x2c, 0x48, 0xdc, 0xb2, 0xa6, 0x5c, 0x6d,
	0x5c, 0xb2, 0x1d, 0x24, 0xdd, 0x82, 0x3f, 0x43, 0x79, 0x5b, 0xe0, 0xb4, 0x83, 0xc4, 0xfb, 0x84,
	0xda, 0xbc, 0x39, 0xb9, 0x81, 0x8a, 0x39, 0xa0, 0x15, 0x78, 0xf8, 0xcf, 0xb6, 0xbe, 0x8d, 0x91,
	0x33, 0x28, 0xcb, 0x31, 0x3b, 0xb7, 0x2e, 0x57, 0xfb, 0xd7, 0x93, 0x9d, 0x3a, 0x54, 0x85, 0xe9,
	0xd6, 0xfc, 0x2e, 0x82, 0xb7, 0xac, 0xb9, 0xcc, 0x78, 0x2a, 0x91, 0x6c, 0x2f, 0x2e, 0x5b, 0xd7,
	0xb7, 0xbb, 0x5b, 0xa8, 0x35, 0xdb, 0x8e, 0xd7, 0x95, 0x5d, 0x2e, 0xf6, 0xd1, 0x88, 0x35, 0xe7,
	0xb9, 0xde, 0x40, 0xac, 0xa9, 0xf7, 0xdb, 0x6c, 0x07, 0xa0, 0x26, 0xec, 0x6b, 0xa7, 0xfe, 0x5c,
	0xb5, 0x8c, 0x51, 0x45, 0xdb, 0xb8, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x87, 0x12, 0xe2,
	0x30, 0x03, 0x00, 0x00,
}
