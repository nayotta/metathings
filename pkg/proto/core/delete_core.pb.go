// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delete_core.proto

package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DeleteCoreRequest struct {
	Id *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteCoreRequest) Reset()                    { *m = DeleteCoreRequest{} }
func (m *DeleteCoreRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteCoreRequest) ProtoMessage()               {}
func (*DeleteCoreRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *DeleteCoreRequest) GetId() *google_protobuf.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteCoreRequest)(nil), "ai.metathings.service.core.DeleteCoreRequest")
}

func init() { proto.RegisterFile("delete_core.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0x86, 0x49, 0x8a, 0x2b, 0x62, 0x75, 0x57, 0x49, 0x10, 0x3d, 0xac, 0x6c, 0x32, 0x0b, 0x2a,
	0x3e, 0x80, 0xda, 0xd8, 0x46, 0xb0, 0x95, 0x4d, 0x76, 0xdc, 0x0c, 0x6e, 0x32, 0x71, 0x76, 0x36,
	0x79, 0x5c, 0xc1, 0x27, 0x11, 0x37, 0x78, 0xdd, 0xc0, 0x7c, 0x1f, 0xdf, 0x5f, 0xed, 0x1d, 0x06,
	0x54, 0x7c, 0xef, 0x59, 0x10, 0x66, 0x61, 0xe5, 0x43, 0x6d, 0x09, 0x46, 0x54, 0xab, 0x03, 0x4d,
	0x3e, 0x42, 0x44, 0x59, 0xa8, 0x47, 0xf8, 0x23, 0xea, 0x4b, 0xcf, 0xec, 0x03, 0x9a, 0x4c, 0x76,
	0xe9, 0xc3, 0xac, 0x62, 0xe7, 0x19, 0x25, 0x6e, 0x6e, 0xfd, 0xe0, 0x49, 0x87, 0xd4, 0x41, 0xcf,
	0xa3, 0x19, 0x57, 0xd2, 0x4f, 0x5e, 0x8d, 0xe7, 0x26, 0x3f, 0x9b, 0xc5, 0x06, 0x72, 0x56, 0x59,
	0xa2, 0x39, 0x9d, 0x9b, 0x77, 0xfd, 0x52, 0xed, 0x9f, 0xf3, 0x90, 0x27, 0x16, 0x6c, 0xf1, 0x2b,
	0x61, 0xd4, 0xc3, 0x7d, 0x55, 0x92, 0x3b, 0x2f, 0x8e, 0xc5, 0xcd, 0xd9, 0xed, 0x05, 0x6c, 0x65,
	0xf8, 0x2f, 0xc3, 0xab, 0x0a, 0x4d, 0xfe, 0xcd, 0x86, 0x84, 0x8f, 0xbb, 0x9f, 0xef, 0xab, 0xf2,
	0x58, 0xb4, 0x25, 0xb9, 0x6e, 0x97, 0x89, 0xbb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x55, 0xce,
	0xc2, 0x96, 0xda, 0x00, 0x00, 0x00,
}