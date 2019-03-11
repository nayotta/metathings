// Code generated by protoc-gen-go. DO NOT EDIT.
// source: query_frames_from_flow.proto

package deviced

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/golang/protobuf/ptypes/wrappers"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type QueryFramesFromFlowRequest struct {
	Device               *OpDevice            `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	From                 *timestamp.Timestamp `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To                   *timestamp.Timestamp `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *QueryFramesFromFlowRequest) Reset()         { *m = QueryFramesFromFlowRequest{} }
func (m *QueryFramesFromFlowRequest) String() string { return proto.CompactTextString(m) }
func (*QueryFramesFromFlowRequest) ProtoMessage()    {}
func (*QueryFramesFromFlowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d61c91b09a6fd445, []int{0}
}

func (m *QueryFramesFromFlowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryFramesFromFlowRequest.Unmarshal(m, b)
}
func (m *QueryFramesFromFlowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryFramesFromFlowRequest.Marshal(b, m, deterministic)
}
func (m *QueryFramesFromFlowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFramesFromFlowRequest.Merge(m, src)
}
func (m *QueryFramesFromFlowRequest) XXX_Size() int {
	return xxx_messageInfo_QueryFramesFromFlowRequest.Size(m)
}
func (m *QueryFramesFromFlowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFramesFromFlowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFramesFromFlowRequest proto.InternalMessageInfo

func (m *QueryFramesFromFlowRequest) GetDevice() *OpDevice {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *QueryFramesFromFlowRequest) GetFrom() *timestamp.Timestamp {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *QueryFramesFromFlowRequest) GetTo() *timestamp.Timestamp {
	if m != nil {
		return m.To
	}
	return nil
}

type QueryFramesFromFlowResponse struct {
	Packs                []*QueryFramesFromFlowResponse_Pack `protobuf:"bytes,1,rep,name=packs,proto3" json:"packs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *QueryFramesFromFlowResponse) Reset()         { *m = QueryFramesFromFlowResponse{} }
func (m *QueryFramesFromFlowResponse) String() string { return proto.CompactTextString(m) }
func (*QueryFramesFromFlowResponse) ProtoMessage()    {}
func (*QueryFramesFromFlowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d61c91b09a6fd445, []int{1}
}

func (m *QueryFramesFromFlowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryFramesFromFlowResponse.Unmarshal(m, b)
}
func (m *QueryFramesFromFlowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryFramesFromFlowResponse.Marshal(b, m, deterministic)
}
func (m *QueryFramesFromFlowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFramesFromFlowResponse.Merge(m, src)
}
func (m *QueryFramesFromFlowResponse) XXX_Size() int {
	return xxx_messageInfo_QueryFramesFromFlowResponse.Size(m)
}
func (m *QueryFramesFromFlowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFramesFromFlowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFramesFromFlowResponse proto.InternalMessageInfo

func (m *QueryFramesFromFlowResponse) GetPacks() []*QueryFramesFromFlowResponse_Pack {
	if m != nil {
		return m.Packs
	}
	return nil
}

type QueryFramesFromFlowResponse_Pack struct {
	Flow                 *Flow    `protobuf:"bytes,1,opt,name=flow,proto3" json:"flow,omitempty"`
	Frames               []*Frame `protobuf:"bytes,2,rep,name=frames,proto3" json:"frames,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryFramesFromFlowResponse_Pack) Reset()         { *m = QueryFramesFromFlowResponse_Pack{} }
func (m *QueryFramesFromFlowResponse_Pack) String() string { return proto.CompactTextString(m) }
func (*QueryFramesFromFlowResponse_Pack) ProtoMessage()    {}
func (*QueryFramesFromFlowResponse_Pack) Descriptor() ([]byte, []int) {
	return fileDescriptor_d61c91b09a6fd445, []int{1, 0}
}

func (m *QueryFramesFromFlowResponse_Pack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryFramesFromFlowResponse_Pack.Unmarshal(m, b)
}
func (m *QueryFramesFromFlowResponse_Pack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryFramesFromFlowResponse_Pack.Marshal(b, m, deterministic)
}
func (m *QueryFramesFromFlowResponse_Pack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFramesFromFlowResponse_Pack.Merge(m, src)
}
func (m *QueryFramesFromFlowResponse_Pack) XXX_Size() int {
	return xxx_messageInfo_QueryFramesFromFlowResponse_Pack.Size(m)
}
func (m *QueryFramesFromFlowResponse_Pack) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFramesFromFlowResponse_Pack.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFramesFromFlowResponse_Pack proto.InternalMessageInfo

func (m *QueryFramesFromFlowResponse_Pack) GetFlow() *Flow {
	if m != nil {
		return m.Flow
	}
	return nil
}

func (m *QueryFramesFromFlowResponse_Pack) GetFrames() []*Frame {
	if m != nil {
		return m.Frames
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryFramesFromFlowRequest)(nil), "ai.metathings.service.deviced.QueryFramesFromFlowRequest")
	proto.RegisterType((*QueryFramesFromFlowResponse)(nil), "ai.metathings.service.deviced.QueryFramesFromFlowResponse")
	proto.RegisterType((*QueryFramesFromFlowResponse_Pack)(nil), "ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack")
}

func init() { proto.RegisterFile("query_frames_from_flow.proto", fileDescriptor_d61c91b09a6fd445) }

var fileDescriptor_d61c91b09a6fd445 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x69, 0x37, 0x77, 0xc8, 0x6e, 0x39, 0x95, 0xfa, 0x6f, 0x4c, 0xc1, 0x21, 0x2c, 0x85,
	0x09, 0x7a, 0x11, 0x76, 0x91, 0x5d, 0xd5, 0xa2, 0xe7, 0x91, 0xb5, 0x6f, 0xbb, 0xb0, 0x66, 0x6f,
	0x96, 0xa4, 0x2b, 0x1e, 0xfc, 0x60, 0x7e, 0x2e, 0xbf, 0x80, 0x34, 0xed, 0x3c, 0x88, 0xac, 0xa7,
	0x26, 0xbc, 0xcf, 0xef, 0xe1, 0x79, 0xde, 0x86, 0x9c, 0xed, 0x4a, 0xd0, 0x1f, 0xcb, 0x4c, 0x73,
	0x09, 0x66, 0x99, 0x69, 0x94, 0xcb, 0xac, 0xc0, 0x8a, 0x29, 0x8d, 0x16, 0xe9, 0x39, 0x17, 0x4c,
	0x82, 0xe5, 0x76, 0x2d, 0xb6, 0xb9, 0x61, 0x06, 0xf4, 0x5e, 0x24, 0xc0, 0x52, 0xa8, 0x3f, 0x69,
	0x78, 0x91, 0x23, 0xe6, 0x05, 0x44, 0x4e, 0xbc, 0x2a, 0xb3, 0xa8, 0xd2, 0x5c, 0x29, 0xd0, 0xa6,
	0xc1, 0xc3, 0xcb, 0xbf, 0x73, 0x2b, 0x24, 0x18, 0xcb, 0xa5, 0x6a, 0x05, 0xf7, 0xb9, 0xb0, 0xeb,
	0x72, 0xc5, 0x12, 0x94, 0x91, 0xac, 0x84, 0xdd, 0x60, 0x15, 0xe5, 0x38, 0x75, 0xc3, 0xe9, 0x9e,
	0x17, 0x22, 0xe5, 0x16, 0xb5, 0x89, 0x7e, 0x8f, 0x2d, 0x37, 0x94, 0x98, 0x42, 0xd1, 0x5c, 0xc6,
	0x5f, 0x1e, 0x09, 0x5f, 0xeb, 0x16, 0x0b, 0x57, 0x62, 0xa1, 0x51, 0x2e, 0x0a, 0xac, 0x62, 0xd8,
	0x95, 0x60, 0x2c, 0x9d, 0x93, 0x41, 0x93, 0x37, 0xf0, 0x46, 0xde, 0x64, 0x38, 0xbb, 0x61, 0x47,
	0x4b, 0xb1, 0x67, 0xf5, 0xe4, 0x4e, 0x71, 0x8b, 0x51, 0x46, 0xfa, 0xf5, 0x5e, 0x02, 0xdf, 0xe1,
	0x21, 0x6b, 0x4a, 0xb1, 0x43, 0x29, 0xf6, 0x76, 0x28, 0x15, 0x3b, 0x1d, 0xbd, 0x25, 0xbe, 0xc5,
	0xa0, 0xd7, 0xa9, 0xf6, 0x2d, 0x8e, 0xbf, 0x3d, 0x72, 0xfa, 0x6f, 0x76, 0xa3, 0x70, 0x6b, 0x80,
	0xbe, 0x93, 0x13, 0xc5, 0x93, 0x8d, 0x09, 0xbc, 0x51, 0x6f, 0x32, 0x9c, 0xcd, 0x3b, 0xb2, 0x1f,
	0xb1, 0x62, 0x2f, 0x3c, 0xd9, 0xc4, 0x8d, 0x5b, 0xf8, 0x49, 0xfa, 0xf5, 0x95, 0x3e, 0x90, 0x7e,
	0xfd, 0xb7, 0xdb, 0xcd, 0x5c, 0x75, 0xb8, 0x3b, 0x3b, 0x07, 0xd0, 0x47, 0x32, 0x68, 0x9e, 0x4c,
	0xe0, 0xbb, 0x60, 0xd7, 0x5d, 0x68, 0x2d, 0x8e, 0x5b, 0x66, 0x35, 0x70, 0xdb, 0xb8, 0xfb, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x5d, 0x3d, 0x52, 0x4e, 0x7d, 0x02, 0x00, 0x00,
}