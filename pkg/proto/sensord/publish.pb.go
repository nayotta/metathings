// Code generated by protoc-gen-go. DO NOT EDIT.
// source: publish.proto

package sensord

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PublishRequest struct {
	Session *wrappers.UInt64Value `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*PublishRequest_Data
	Payload              isPublishRequest_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *PublishRequest) Reset()         { *m = PublishRequest{} }
func (m *PublishRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()    {}
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_34180b7635741fb2, []int{0}
}

func (m *PublishRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRequest.Unmarshal(m, b)
}
func (m *PublishRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRequest.Marshal(b, m, deterministic)
}
func (m *PublishRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRequest.Merge(m, src)
}
func (m *PublishRequest) XXX_Size() int {
	return xxx_messageInfo_PublishRequest.Size(m)
}
func (m *PublishRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRequest proto.InternalMessageInfo

func (m *PublishRequest) GetSession() *wrappers.UInt64Value {
	if m != nil {
		return m.Session
	}
	return nil
}

type isPublishRequest_Payload interface {
	isPublishRequest_Payload()
}

type PublishRequest_Data struct {
	Data *SensorData `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

func (*PublishRequest_Data) isPublishRequest_Payload() {}

func (m *PublishRequest) GetPayload() isPublishRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PublishRequest) GetData() *SensorData {
	if x, ok := m.GetPayload().(*PublishRequest_Data); ok {
		return x.Data
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PublishRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PublishRequest_OneofMarshaler, _PublishRequest_OneofUnmarshaler, _PublishRequest_OneofSizer, []interface{}{
		(*PublishRequest_Data)(nil),
	}
}

func _PublishRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PublishRequest)
	// payload
	switch x := m.Payload.(type) {
	case *PublishRequest_Data:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Data); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("PublishRequest.Payload has unexpected type %T", x)
	}
	return nil
}

func _PublishRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PublishRequest)
	switch tag {
	case 2: // payload.data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SensorData)
		err := b.DecodeMessage(msg)
		m.Payload = &PublishRequest_Data{msg}
		return true, err
	default:
		return false, nil
	}
}

func _PublishRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PublishRequest)
	// payload
	switch x := m.Payload.(type) {
	case *PublishRequest_Data:
		s := proto.Size(x.Data)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type PublishRequests struct {
	Requests             []*PublishRequest `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PublishRequests) Reset()         { *m = PublishRequests{} }
func (m *PublishRequests) String() string { return proto.CompactTextString(m) }
func (*PublishRequests) ProtoMessage()    {}
func (*PublishRequests) Descriptor() ([]byte, []int) {
	return fileDescriptor_34180b7635741fb2, []int{1}
}

func (m *PublishRequests) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRequests.Unmarshal(m, b)
}
func (m *PublishRequests) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRequests.Marshal(b, m, deterministic)
}
func (m *PublishRequests) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRequests.Merge(m, src)
}
func (m *PublishRequests) XXX_Size() int {
	return xxx_messageInfo_PublishRequests.Size(m)
}
func (m *PublishRequests) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRequests.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRequests proto.InternalMessageInfo

func (m *PublishRequests) GetRequests() []*PublishRequest {
	if m != nil {
		return m.Requests
	}
	return nil
}

type PublishResponse struct {
	Session              uint64   `protobuf:"varint,1,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishResponse) Reset()         { *m = PublishResponse{} }
func (m *PublishResponse) String() string { return proto.CompactTextString(m) }
func (*PublishResponse) ProtoMessage()    {}
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_34180b7635741fb2, []int{2}
}

func (m *PublishResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishResponse.Unmarshal(m, b)
}
func (m *PublishResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishResponse.Marshal(b, m, deterministic)
}
func (m *PublishResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishResponse.Merge(m, src)
}
func (m *PublishResponse) XXX_Size() int {
	return xxx_messageInfo_PublishResponse.Size(m)
}
func (m *PublishResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PublishResponse proto.InternalMessageInfo

func (m *PublishResponse) GetSession() uint64 {
	if m != nil {
		return m.Session
	}
	return 0
}

type PublishResponses struct {
	Responses            []*PublishResponse `protobuf:"bytes,1,rep,name=responses,proto3" json:"responses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PublishResponses) Reset()         { *m = PublishResponses{} }
func (m *PublishResponses) String() string { return proto.CompactTextString(m) }
func (*PublishResponses) ProtoMessage()    {}
func (*PublishResponses) Descriptor() ([]byte, []int) {
	return fileDescriptor_34180b7635741fb2, []int{3}
}

func (m *PublishResponses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishResponses.Unmarshal(m, b)
}
func (m *PublishResponses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishResponses.Marshal(b, m, deterministic)
}
func (m *PublishResponses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishResponses.Merge(m, src)
}
func (m *PublishResponses) XXX_Size() int {
	return xxx_messageInfo_PublishResponses.Size(m)
}
func (m *PublishResponses) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishResponses.DiscardUnknown(m)
}

var xxx_messageInfo_PublishResponses proto.InternalMessageInfo

func (m *PublishResponses) GetResponses() []*PublishResponse {
	if m != nil {
		return m.Responses
	}
	return nil
}

func init() {
	proto.RegisterType((*PublishRequest)(nil), "ai.metathings.service.sensord.PublishRequest")
	proto.RegisterType((*PublishRequests)(nil), "ai.metathings.service.sensord.PublishRequests")
	proto.RegisterType((*PublishResponse)(nil), "ai.metathings.service.sensord.PublishResponse")
	proto.RegisterType((*PublishResponses)(nil), "ai.metathings.service.sensord.PublishResponses")
}

func init() { proto.RegisterFile("publish.proto", fileDescriptor_34180b7635741fb2) }

var fileDescriptor_34180b7635741fb2 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0xc7, 0x9f, 0xf4, 0x29, 0xad, 0xdd, 0xfa, 0x46, 0x4e, 0xa1, 0xf8, 0x52, 0x72, 0xaa, 0x48,
	0x37, 0x50, 0xa5, 0x47, 0x85, 0xe2, 0xc1, 0x82, 0x07, 0x89, 0xe8, 0xc9, 0x83, 0x93, 0x66, 0x4d,
	0x17, 0xd3, 0xcc, 0xba, 0xb3, 0x69, 0xf1, 0xbb, 0xf8, 0xdd, 0x04, 0x3f, 0x89, 0xb0, 0xdb, 0x56,
	0xe3, 0x41, 0x3d, 0x65, 0x86, 0xcc, 0x6f, 0xe6, 0xf7, 0x5f, 0xb6, 0xa5, 0xca, 0x24, 0x97, 0x34,
	0xe5, 0x4a, 0xa3, 0x41, 0x7f, 0x1f, 0x24, 0x9f, 0x09, 0x03, 0x66, 0x2a, 0x8b, 0x8c, 0x38, 0x09,
	0x3d, 0x97, 0x13, 0xc1, 0x49, 0x14, 0x84, 0x3a, 0xed, 0x1c, 0x64, 0x88, 0x59, 0x2e, 0x22, 0x3b,
	0x9c, 0x94, 0x8f, 0xd1, 0x42, 0x83, 0x52, 0x42, 0x93, 0xc3, 0x3b, 0xc3, 0x4c, 0x9a, 0x69, 0x99,
	0xf0, 0x09, 0xce, 0xa2, 0xd9, 0x42, 0x9a, 0x27, 0x5c, 0x44, 0x19, 0xf6, 0xed, 0xcf, 0xfe, 0x1c,
	0x72, 0x99, 0x82, 0x41, 0x4d, 0xd1, 0xba, 0x5c, 0x72, 0x9b, 0xee, 0x80, 0xeb, 0xc2, 0x57, 0x8f,
	0x6d, 0x5f, 0x3b, 0xad, 0x58, 0x3c, 0x97, 0x82, 0x8c, 0x7f, 0xc6, 0x9a, 0x24, 0x88, 0x24, 0x16,
	0x81, 0xd7, 0xf5, 0x7a, 0xed, 0xc1, 0x1e, 0x77, 0x2a, 0x7c, 0xa5, 0xc2, 0x6f, 0xc7, 0x85, 0x19,
	0x9e, 0xde, 0x41, 0x5e, 0x8a, 0x51, 0xe3, 0xfd, 0xed, 0xb0, 0xd6, 0xf5, 0xe2, 0x15, 0xe4, 0x9f,
	0xb3, 0x7a, 0x0a, 0x06, 0x82, 0x9a, 0x85, 0x8f, 0xf8, 0x8f, 0x31, 0xf9, 0x8d, 0xfd, 0x5e, 0x80,
	0x81, 0xcb, 0x7f, 0xb1, 0x05, 0x47, 0x2d, 0xd6, 0x54, 0xf0, 0x92, 0x23, 0xa4, 0xe1, 0x3d, 0xdb,
	0xa9, 0xda, 0x91, 0x3f, 0x66, 0x1b, 0x7a, 0x59, 0x07, 0x5e, 0xf7, 0x7f, 0xaf, 0x3d, 0xe8, 0xff,
	0x72, 0xa2, 0xba, 0x21, 0x5e, 0xe3, 0xe1, 0xf1, 0x97, 0xed, 0xa4, 0xb0, 0x20, 0xe1, 0x07, 0xd5,
	0xf0, 0xf5, 0x75, 0xac, 0xf0, 0x81, 0xed, 0x7e, 0x1b, 0x26, 0xff, 0x8a, 0xb5, 0xf4, 0xaa, 0x59,
	0xca, 0xf0, 0xbf, 0xca, 0x38, 0x2c, 0xfe, 0x5c, 0x90, 0x34, 0xec, 0xfb, 0x9e, 0x7c, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x64, 0x46, 0x81, 0xae, 0x28, 0x02, 0x00, 0x00,
}
