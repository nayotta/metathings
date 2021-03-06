// Code generated by protoc-gen-go. DO NOT EDIT.
// source: query_storage_by_device.proto

package evaluatord

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	pagination "github.com/nayotta/metathings/proto/common/option/pagination"
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

type QueryStorageByDeviceRequest struct {
	Device               *deviced.OpDevice              `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	Source               *OpResource                    `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Measurement          *wrappers.StringValue          `protobuf:"bytes,3,opt,name=measurement,proto3" json:"measurement,omitempty"`
	RangeFrom            *timestamp.Timestamp           `protobuf:"bytes,4,opt,name=range_from,json=rangeFrom,proto3" json:"range_from,omitempty"`
	RangeTo              *timestamp.Timestamp           `protobuf:"bytes,5,opt,name=range_to,json=rangeTo,proto3" json:"range_to,omitempty"`
	QueryString          *wrappers.StringValue          `protobuf:"bytes,6,opt,name=query_string,json=queryString,proto3" json:"query_string,omitempty"`
	Pagination           *pagination.OpPaginationOption `protobuf:"bytes,33,opt,name=pagination,proto3" json:"pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *QueryStorageByDeviceRequest) Reset()         { *m = QueryStorageByDeviceRequest{} }
func (m *QueryStorageByDeviceRequest) String() string { return proto.CompactTextString(m) }
func (*QueryStorageByDeviceRequest) ProtoMessage()    {}
func (*QueryStorageByDeviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d25223cb37921b9, []int{0}
}

func (m *QueryStorageByDeviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryStorageByDeviceRequest.Unmarshal(m, b)
}
func (m *QueryStorageByDeviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryStorageByDeviceRequest.Marshal(b, m, deterministic)
}
func (m *QueryStorageByDeviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStorageByDeviceRequest.Merge(m, src)
}
func (m *QueryStorageByDeviceRequest) XXX_Size() int {
	return xxx_messageInfo_QueryStorageByDeviceRequest.Size(m)
}
func (m *QueryStorageByDeviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStorageByDeviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStorageByDeviceRequest proto.InternalMessageInfo

func (m *QueryStorageByDeviceRequest) GetDevice() *deviced.OpDevice {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *QueryStorageByDeviceRequest) GetSource() *OpResource {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *QueryStorageByDeviceRequest) GetMeasurement() *wrappers.StringValue {
	if m != nil {
		return m.Measurement
	}
	return nil
}

func (m *QueryStorageByDeviceRequest) GetRangeFrom() *timestamp.Timestamp {
	if m != nil {
		return m.RangeFrom
	}
	return nil
}

func (m *QueryStorageByDeviceRequest) GetRangeTo() *timestamp.Timestamp {
	if m != nil {
		return m.RangeTo
	}
	return nil
}

func (m *QueryStorageByDeviceRequest) GetQueryString() *wrappers.StringValue {
	if m != nil {
		return m.QueryString
	}
	return nil
}

func (m *QueryStorageByDeviceRequest) GetPagination() *pagination.OpPaginationOption {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryStorageByDeviceResponse struct {
	Frames               []*deviced.Frame             `protobuf:"bytes,2,rep,name=frames,proto3" json:"frames,omitempty"`
	Pagination           *pagination.PaginationOption `protobuf:"bytes,33,opt,name=pagination,proto3" json:"pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *QueryStorageByDeviceResponse) Reset()         { *m = QueryStorageByDeviceResponse{} }
func (m *QueryStorageByDeviceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryStorageByDeviceResponse) ProtoMessage()    {}
func (*QueryStorageByDeviceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d25223cb37921b9, []int{1}
}

func (m *QueryStorageByDeviceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryStorageByDeviceResponse.Unmarshal(m, b)
}
func (m *QueryStorageByDeviceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryStorageByDeviceResponse.Marshal(b, m, deterministic)
}
func (m *QueryStorageByDeviceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStorageByDeviceResponse.Merge(m, src)
}
func (m *QueryStorageByDeviceResponse) XXX_Size() int {
	return xxx_messageInfo_QueryStorageByDeviceResponse.Size(m)
}
func (m *QueryStorageByDeviceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStorageByDeviceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStorageByDeviceResponse proto.InternalMessageInfo

func (m *QueryStorageByDeviceResponse) GetFrames() []*deviced.Frame {
	if m != nil {
		return m.Frames
	}
	return nil
}

func (m *QueryStorageByDeviceResponse) GetPagination() *pagination.PaginationOption {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryStorageByDeviceRequest)(nil), "ai.metathings.service.evaluatord.QueryStorageByDeviceRequest")
	proto.RegisterType((*QueryStorageByDeviceResponse)(nil), "ai.metathings.service.evaluatord.QueryStorageByDeviceResponse")
}

func init() { proto.RegisterFile("query_storage_by_device.proto", fileDescriptor_1d25223cb37921b9) }

var fileDescriptor_1d25223cb37921b9 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x6e, 0x13, 0x3d,
	0x14, 0xc5, 0x35, 0x69, 0x3b, 0x5f, 0x3f, 0xa7, 0x0b, 0xe4, 0x0d, 0xa3, 0x50, 0x20, 0x54, 0x48,
	0x74, 0x81, 0x6c, 0xa9, 0x08, 0xf1, 0x47, 0xac, 0x22, 0x54, 0xd1, 0x05, 0x04, 0xa6, 0x15, 0x42,
	0xdd, 0x44, 0x4e, 0xe6, 0x66, 0x6a, 0x69, 0xec, 0xeb, 0xd8, 0x9e, 0xa0, 0xbc, 0x02, 0xaf, 0xc0,
	0x7b, 0xf0, 0x70, 0xac, 0x50, 0xec, 0x09, 0x19, 0x85, 0x88, 0x94, 0x9d, 0x67, 0xee, 0x3d, 0x3f,
	0x1d, 0xdd, 0x73, 0xc8, 0xfd, 0x59, 0x0d, 0x76, 0x31, 0x72, 0x1e, 0xad, 0x28, 0x61, 0x34, 0x5e,
	0x8c, 0x0a, 0x98, 0xcb, 0x09, 0x30, 0x63, 0xd1, 0x23, 0xed, 0x0b, 0xc9, 0x14, 0x78, 0xe1, 0x6f,
	0xa4, 0x2e, 0x1d, 0x73, 0x60, 0xc3, 0x10, 0xe6, 0xa2, 0xaa, 0x85, 0x47, 0x5b, 0xf4, 0x1e, 0x94,
	0x88, 0x65, 0x05, 0x3c, 0xec, 0x8f, 0xeb, 0x29, 0xff, 0x6a, 0x85, 0x31, 0x60, 0x5d, 0x24, 0xf4,
	0x1e, 0x6e, 0xce, 0xbd, 0x54, 0xe0, 0xbc, 0x50, 0xa6, 0x59, 0xb8, 0x3b, 0x17, 0x95, 0x2c, 0x84,
	0x07, 0xbe, 0x7a, 0x34, 0x83, 0x17, 0xa5, 0xf4, 0x37, 0xf5, 0x98, 0x4d, 0x50, 0x71, 0x2d, 0x16,
	0xe8, 0xbd, 0xe0, 0x6b, 0x2f, 0x11, 0xc8, 0xa3, 0xdb, 0x82, 0x2b, 0x2c, 0xa0, 0x6a, 0x84, 0xef,
	0x6f, 0x23, 0x9c, 0xa0, 0x52, 0xa8, 0x39, 0x1a, 0x2f, 0x51, 0x73, 0x23, 0x4a, 0xa9, 0xc5, 0xc6,
	0xb3, 0xc1, 0x75, 0x5b, 0xec, 0x93, 0xef, 0xfb, 0xe4, 0xde, 0xa7, 0xe5, 0xc9, 0x2e, 0xe3, 0xc5,
	0x06, 0x8b, 0xb7, 0xc1, 0x41, 0x0e, 0xb3, 0x1a, 0x9c, 0xa7, 0x17, 0x24, 0x8d, 0x96, 0xb2, 0xa4,
	0x9f, 0x9c, 0x76, 0xcf, 0x9e, 0xb0, 0xed, 0x17, 0x6c, 0x7c, 0xb3, 0xa1, 0x89, 0xfa, 0xc1, 0xe1,
	0xcf, 0xc1, 0xc1, 0xb7, 0xa4, 0x73, 0x27, 0xc9, 0x1b, 0x00, 0xfd, 0x40, 0x52, 0x87, 0xb5, 0x9d,
	0x40, 0xd6, 0x09, 0xa8, 0xa7, 0x6c, 0x57, 0x18, 0x6c, 0x68, 0x72, 0x88, 0x9a, 0x36, 0x2f, 0xfe,
	0xa1, 0xef, 0x48, 0x57, 0x81, 0x70, 0xb5, 0x05, 0x05, 0xda, 0x67, 0x7b, 0x01, 0x7a, 0xcc, 0x62,
	0x3e, 0x6c, 0x95, 0x0f, 0xbb, 0xf4, 0x56, 0xea, 0xf2, 0xb3, 0xa8, 0xea, 0x36, 0xa4, 0x2d, 0xa5,
	0xaf, 0x08, 0xb1, 0x42, 0x97, 0x30, 0x9a, 0x5a, 0x54, 0xd9, 0x7e, 0x00, 0xf5, 0xfe, 0x00, 0x5d,
	0xad, 0x82, 0xce, 0xff, 0x0f, 0xdb, 0xe7, 0x16, 0x15, 0x7d, 0x4e, 0x0e, 0xa3, 0xd4, 0x63, 0x76,
	0xb0, 0x53, 0xf8, 0x5f, 0xd8, 0xbd, 0x42, 0x7a, 0x41, 0x8e, 0x56, 0x45, 0x5d, 0xba, 0xcb, 0xd2,
	0x7f, 0x33, 0x3f, 0x8b, 0x89, 0x2d, 0x67, 0xf4, 0x9a, 0x90, 0x75, 0xc4, 0xd9, 0xa3, 0x00, 0x7a,
	0xbd, 0x71, 0xda, 0x58, 0x0e, 0x16, 0xcb, 0xc1, 0x5a, 0x8d, 0x18, 0x9a, 0x8f, 0xbf, 0x3f, 0x86,
	0x61, 0x9a, 0xb7, 0x68, 0x27, 0x3f, 0x12, 0x72, 0xbc, 0xbd, 0x1d, 0xce, 0xa0, 0x76, 0x40, 0xdf,
	0x90, 0x74, 0x6a, 0x85, 0x02, 0x97, 0x75, 0xfa, 0x7b, 0xa7, 0xdd, 0xb3, 0xc7, 0x3b, 0xea, 0x71,
	0xbe, 0x5c, 0xce, 0x1b, 0x0d, 0xfd, 0xb2, 0xc5, 0xfa, 0xcb, 0xdb, 0x5a, 0xff, 0x9b, 0xf1, 0xc1,
	0xd1, 0x35, 0x59, 0xd7, 0x68, 0x9c, 0x86, 0x7b, 0x3e, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0xdf,
	0x51, 0x63, 0x47, 0x1d, 0x04, 0x00, 0x00,
}
