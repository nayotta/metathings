// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_cores.proto

package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ListCoresRequest struct {
}

func (m *ListCoresRequest) Reset()                    { *m = ListCoresRequest{} }
func (m *ListCoresRequest) String() string            { return proto.CompactTextString(m) }
func (*ListCoresRequest) ProtoMessage()               {}
func (*ListCoresRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type ListCoresResponse struct {
	Cores []*Core `protobuf:"bytes,1,rep,name=cores" json:"cores,omitempty"`
}

func (m *ListCoresResponse) Reset()                    { *m = ListCoresResponse{} }
func (m *ListCoresResponse) String() string            { return proto.CompactTextString(m) }
func (*ListCoresResponse) ProtoMessage()               {}
func (*ListCoresResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *ListCoresResponse) GetCores() []*Core {
	if m != nil {
		return m.Cores
	}
	return nil
}

func init() {
	proto.RegisterType((*ListCoresRequest)(nil), "ai.metathings.service.core.ListCoresRequest")
	proto.RegisterType((*ListCoresResponse)(nil), "ai.metathings.service.core.ListCoresResponse")
}

func init() { proto.RegisterFile("list_cores.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8e, 0xc1, 0xca, 0x82, 0x50,
	0x10, 0x85, 0xf9, 0xf9, 0xa9, 0xc5, 0x6d, 0x63, 0xae, 0xc2, 0x95, 0xb8, 0x6a, 0xe3, 0x08, 0x05,
	0xbe, 0x40, 0xcb, 0x5a, 0xf9, 0x02, 0x71, 0xb5, 0x41, 0x87, 0xd4, 0xb1, 0x3b, 0xa3, 0xbe, 0x7e,
	0x5c, 0x85, 0x68, 0xd3, 0xee, 0x30, 0x9c, 0x6f, 0xce, 0x67, 0x82, 0x96, 0x44, 0xef, 0x15, 0x3b,
	0x14, 0x18, 0x1c, 0x2b, 0x87, 0x91, 0x25, 0xe8, 0x50, 0xad, 0x36, 0xd4, 0xd7, 0x02, 0x82, 0x6e,
	0xa2, 0x0a, 0xc1, 0x57, 0xa2, 0xbc, 0x26, 0x6d, 0xc6, 0x12, 0x2a, 0xee, 0xb2, 0x6e, 0x26, 0x7d,
	0xf2, 0x9c, 0xd5, 0x9c, 0x2e, 0x60, 0x3a, 0xd9, 0x96, 0x1e, 0x56, 0xd9, 0x49, 0xf6, 0x89, 0xeb,
	0xcf, 0xc8, 0x78, 0x7a, 0xcd, 0x49, 0x68, 0x82, 0x1b, 0x89, 0x5e, 0xfc, 0x64, 0x81, 0xaf, 0x11,
	0x45, 0x93, 0xab, 0xd9, 0x7f, 0xdd, 0x64, 0xe0, 0x5e, 0x30, 0xcc, 0xcd, 0x66, 0xf1, 0x3a, 0xfc,
	0xc5, 0xff, 0xc7, 0xdd, 0x29, 0x86, 0xdf, 0x62, 0xe0, 0xc9, 0x62, 0xad, 0x97, 0xdb, 0x65, 0xe7,
	0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xce, 0x2c, 0x74, 0xdb, 0x00, 0x00, 0x00,
}