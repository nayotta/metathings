// Code generated by protoc-gen-go. DO NOT EDIT.
// source: put_object_streaming.proto

package ai_metathings_service_device

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PutObjectStreamingRequest struct {
	Id *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Request:
	//	*PutObjectStreamingRequest_Metadata_
	//	*PutObjectStreamingRequest_Chunks
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
	Chunks *deviced.OpObjectChunks `protobuf:"bytes,3,opt,name=chunks,proto3,oneof"`
}

func (*PutObjectStreamingRequest_Metadata_) isPutObjectStreamingRequest_Request() {}

func (*PutObjectStreamingRequest_Chunks) isPutObjectStreamingRequest_Request() {}

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

func (m *PutObjectStreamingRequest) GetChunks() *deviced.OpObjectChunks {
	if x, ok := m.GetRequest().(*PutObjectStreamingRequest_Chunks); ok {
		return x.Chunks
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PutObjectStreamingRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PutObjectStreamingRequest_Metadata_)(nil),
		(*PutObjectStreamingRequest_Chunks)(nil),
	}
}

type PutObjectStreamingRequest_Metadata struct {
	Object               *deviced.OpObject     `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	Sha1                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=sha1,proto3" json:"sha1,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PutObjectStreamingRequest_Metadata) Reset()         { *m = PutObjectStreamingRequest_Metadata{} }
func (m *PutObjectStreamingRequest_Metadata) String() string { return proto.CompactTextString(m) }
func (*PutObjectStreamingRequest_Metadata) ProtoMessage()    {}
func (*PutObjectStreamingRequest_Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_662424e73095d172, []int{0, 0}
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

func (m *PutObjectStreamingRequest_Metadata) GetObject() *deviced.OpObject {
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
	Chunks *deviced.ObjectChunks `protobuf:"bytes,2,opt,name=chunks,proto3,oneof"`
}

func (*PutObjectStreamingResponse_Chunks) isPutObjectStreamingResponse_Response() {}

func (m *PutObjectStreamingResponse) GetResponse() isPutObjectStreamingResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *PutObjectStreamingResponse) GetChunks() *deviced.ObjectChunks {
	if x, ok := m.GetResponse().(*PutObjectStreamingResponse_Chunks); ok {
		return x.Chunks
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PutObjectStreamingResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PutObjectStreamingResponse_Chunks)(nil),
	}
}

func init() {
	proto.RegisterType((*PutObjectStreamingRequest)(nil), "ai.metathings.service.device.PutObjectStreamingRequest")
	proto.RegisterType((*PutObjectStreamingRequest_Metadata)(nil), "ai.metathings.service.device.PutObjectStreamingRequest.Metadata")
	proto.RegisterType((*PutObjectStreamingResponse)(nil), "ai.metathings.service.device.PutObjectStreamingResponse")
}

func init() { proto.RegisterFile("put_object_streaming.proto", fileDescriptor_662424e73095d172) }

var fileDescriptor_662424e73095d172 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbd, 0x4f, 0x32, 0x31,
	0x1c, 0xc7, 0xe1, 0x9e, 0x27, 0x08, 0x35, 0x71, 0xb8, 0x09, 0x2f, 0xc4, 0x18, 0x16, 0x4d, 0xd4,
	0xd6, 0x97, 0xd1, 0x41, 0x83, 0x31, 0xb2, 0x18, 0xcc, 0x91, 0x38, 0x4a, 0x7a, 0xd7, 0xda, 0xab,
	0x70, 0xd7, 0xda, 0xfe, 0xaa, 0x71, 0x30, 0xfe, 0xe3, 0x0e, 0x86, 0xf6, 0xc0, 0x05, 0xd1, 0xa9,
	0xcb, 0xef, 0xfb, 0xf9, 0xbe, 0x14, 0x25, 0xda, 0xc1, 0x44, 0x65, 0x4f, 0x3c, 0x87, 0x89, 0x05,
	0xc3, 0x69, 0x29, 0x2b, 0x81, 0xb5, 0x51, 0xa0, 0xe2, 0x1e, 0x95, 0xb8, 0xe4, 0x40, 0xa1, 0x90,
	0x95, 0xb0, 0xd8, 0x72, 0xf3, 0x22, 0x73, 0x8e, 0x19, 0x9f, 0x3f, 0xc9, 0x8e, 0x50, 0x4a, 0xcc,
	0x38, 0xf1, 0xb7, 0x99, 0x7b, 0x24, 0xaf, 0x86, 0x6a, 0xcd, 0x8d, 0x0d, 0xea, 0xe4, 0x5c, 0x48,
	0x28, 0x5c, 0x86, 0x73, 0x55, 0x92, 0x8a, 0xbe, 0x29, 0x00, 0x4a, 0xbe, 0x69, 0x44, 0x4f, 0x45,
	0x90, 0x92, 0xc0, 0x63, 0xa4, 0x54, 0x8c, 0xcf, 0x82, 0xb8, 0xff, 0x19, 0xa1, 0xed, 0x3b, 0x07,
	0x23, 0x1f, 0x6c, 0xbc, 0xc8, 0x95, 0xf2, 0x67, 0xc7, 0x2d, 0xc4, 0x87, 0x28, 0x92, 0xac, 0xdb,
	0xdc, 0x6d, 0xee, 0x6f, 0x9e, 0xf6, 0x70, 0xc8, 0x81, 0x17, 0x39, 0xf0, 0x18, 0x8c, 0xac, 0xc4,
	0x3d, 0x9d, 0x39, 0x9e, 0x46, 0x92, 0xc5, 0x0f, 0xa8, 0x3d, 0xf7, 0x65, 0x14, 0x68, 0x37, 0xf2,
	0x9a, 0x4b, 0xbc, 0xae, 0x19, 0xfe, 0xd1, 0x18, 0xdf, 0xd6, 0x9c, 0x61, 0x23, 0x5d, 0x32, 0xe3,
	0x1b, 0xd4, 0xca, 0x0b, 0x57, 0x4d, 0x6d, 0xf7, 0x9f, 0xa7, 0x1f, 0xad, 0xa5, 0x33, 0x3c, 0xd2,
	0x81, 0x7e, 0xe5, 0x45, 0xc3, 0x46, 0x5a, 0xcb, 0x93, 0x77, 0xd4, 0x5e, 0x18, 0xc4, 0x17, 0xa8,
	0x15, 0x7e, 0xa5, 0xae, 0xb9, 0xf7, 0x47, 0x68, 0x5a, 0xcb, 0xe2, 0x63, 0xf4, 0xdf, 0x16, 0xf4,
	0xa4, 0x6e, 0xbc, 0x7e, 0x25, 0x7f, 0x39, 0xe8, 0xa0, 0x0d, 0x13, 0x7a, 0xf6, 0x3f, 0x50, 0xb2,
	0x6a, 0x04, 0xab, 0x55, 0x65, 0x79, 0xbc, 0xb5, 0x9c, 0xbf, 0xe3, 0x07, 0xbe, 0x5e, 0x0e, 0x10,
	0xcc, 0x0e, 0x7e, 0xcb, 0xba, 0xb2, 0xfe, 0x00, 0xa1, 0xb6, 0xa9, 0x2d, 0xb2, 0x96, 0xcf, 0x79,
	0xf6, 0x15, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x50, 0xce, 0x26, 0x9f, 0x02, 0x00, 0x00,
}
