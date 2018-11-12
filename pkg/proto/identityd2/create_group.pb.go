// Code generated by protoc-gen-go. DO NOT EDIT.
// source: create_group.proto

package identityd2

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

type CreateGroupRequest struct {
	Id                   *wrappers.StringValue            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Domain               *OpDomain                        `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Name                 *wrappers.StringValue            `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Alias                *wrappers.StringValue            `protobuf:"bytes,6,opt,name=alias,proto3" json:"alias,omitempty"`
	Description          *wrappers.StringValue            `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Extra                map[string]*wrappers.StringValue `protobuf:"bytes,8,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *CreateGroupRequest) Reset()         { *m = CreateGroupRequest{} }
func (m *CreateGroupRequest) String() string { return proto.CompactTextString(m) }
func (*CreateGroupRequest) ProtoMessage()    {}
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b54aa6a17b3be7d3, []int{0}
}

func (m *CreateGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGroupRequest.Unmarshal(m, b)
}
func (m *CreateGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGroupRequest.Marshal(b, m, deterministic)
}
func (m *CreateGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGroupRequest.Merge(m, src)
}
func (m *CreateGroupRequest) XXX_Size() int {
	return xxx_messageInfo_CreateGroupRequest.Size(m)
}
func (m *CreateGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGroupRequest proto.InternalMessageInfo

func (m *CreateGroupRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *CreateGroupRequest) GetDomain() *OpDomain {
	if m != nil {
		return m.Domain
	}
	return nil
}

func (m *CreateGroupRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CreateGroupRequest) GetAlias() *wrappers.StringValue {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *CreateGroupRequest) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *CreateGroupRequest) GetExtra() map[string]*wrappers.StringValue {
	if m != nil {
		return m.Extra
	}
	return nil
}

type CreateGroupResponse struct {
	Group                *Group   `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateGroupResponse) Reset()         { *m = CreateGroupResponse{} }
func (m *CreateGroupResponse) String() string { return proto.CompactTextString(m) }
func (*CreateGroupResponse) ProtoMessage()    {}
func (*CreateGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b54aa6a17b3be7d3, []int{1}
}

func (m *CreateGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGroupResponse.Unmarshal(m, b)
}
func (m *CreateGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGroupResponse.Marshal(b, m, deterministic)
}
func (m *CreateGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGroupResponse.Merge(m, src)
}
func (m *CreateGroupResponse) XXX_Size() int {
	return xxx_messageInfo_CreateGroupResponse.Size(m)
}
func (m *CreateGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGroupResponse proto.InternalMessageInfo

func (m *CreateGroupResponse) GetGroup() *Group {
	if m != nil {
		return m.Group
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateGroupRequest)(nil), "ai.metathings.service.identityd2.CreateGroupRequest")
	proto.RegisterMapType((map[string]*wrappers.StringValue)(nil), "ai.metathings.service.identityd2.CreateGroupRequest.ExtraEntry")
	proto.RegisterType((*CreateGroupResponse)(nil), "ai.metathings.service.identityd2.CreateGroupResponse")
}

func init() { proto.RegisterFile("create_group.proto", fileDescriptor_b54aa6a17b3be7d3) }

var fileDescriptor_b54aa6a17b3be7d3 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x51, 0xab, 0xd3, 0x30,
	0x14, 0xc7, 0x69, 0xef, 0x6d, 0xbd, 0xa6, 0x2f, 0x12, 0x5f, 0xca, 0x10, 0x2d, 0xf7, 0xc5, 0x21,
	0x2e, 0x85, 0x0a, 0x32, 0x06, 0x2a, 0xa8, 0x43, 0xf0, 0x45, 0xa8, 0xba, 0x57, 0xc9, 0x9a, 0x63,
	0x77, 0x58, 0x9b, 0xd4, 0x24, 0xdd, 0xdc, 0xbb, 0xdf, 0x53, 0xf0, 0x93, 0x48, 0x93, 0x3a, 0x15,
	0xc1, 0xea, 0x5b, 0xda, 0x9e, 0xdf, 0xbf, 0xbf, 0x73, 0x72, 0x08, 0xad, 0x34, 0x70, 0x0b, 0x1f,
	0x6a, 0xad, 0xfa, 0x8e, 0x75, 0x5a, 0x59, 0x45, 0x33, 0x8e, 0xac, 0x05, 0xcb, 0xed, 0x0e, 0x65,
	0x6d, 0x98, 0x01, 0x7d, 0xc0, 0x0a, 0x18, 0x0a, 0x90, 0x16, 0xed, 0x49, 0x14, 0xb3, 0xbb, 0xb5,
	0x52, 0x75, 0x03, 0xb9, 0xab, 0xdf, 0xf6, 0x1f, 0xf3, 0xa3, 0xe6, 0x5d, 0x07, 0xda, 0xf8, 0x84,
	0xd9, 0xe3, 0x1a, 0xed, 0xae, 0xdf, 0xb2, 0x4a, 0xb5, 0x79, 0x7b, 0x44, 0xbb, 0x57, 0xc7, 0xbc,
	0x56, 0x0b, 0xf7, 0x71, 0x71, 0xe0, 0x0d, 0x0a, 0x6e, 0x95, 0x36, 0xf9, 0xf9, 0x38, 0x72, 0x49,
	0xab, 0x04, 0x34, 0xfe, 0xe1, 0xfa, 0xcb, 0x25, 0xa1, 0x2f, 0x9c, 0xdd, 0xab, 0x41, 0xae, 0x84,
	0x4f, 0x3d, 0x18, 0x4b, 0x1f, 0x92, 0x10, 0x45, 0x1a, 0x64, 0xc1, 0x3c, 0x29, 0xee, 0x30, 0x2f,
	0xc2, 0x7e, 0x88, 0xb0, 0xb7, 0x56, 0xa3, 0xac, 0x37, 0xbc, 0xe9, 0xa1, 0x0c, 0x51, 0xd0, 0xd7,
	0x24, 0x16, 0xaa, 0xe5, 0x28, 0xd3, 0xd0, 0x11, 0x0f, 0xd8, 0x54, 0x73, 0xec, 0x4d, 0xf7, 0xd2,
	0x11, 0xcf, 0xe3, 0x6f, 0x5f, 0xef, 0x85, 0x59, 0x50, 0x8e, 0x09, 0x74, 0x49, 0x2e, 0x25, 0x6f,
	0x21, 0x8d, 0xa6, 0xff, 0x7d, 0x66, 0x1d, 0x41, 0x57, 0x24, 0xe2, 0x0d, 0x72, 0x93, 0xc6, 0xff,
	0x81, 0x7a, 0x84, 0x3e, 0x25, 0x89, 0x00, 0x53, 0x69, 0xec, 0x2c, 0x2a, 0x99, 0xde, 0xf8, 0x87,
	0xc6, 0x7f, 0x05, 0xe8, 0x7b, 0x12, 0xc1, 0x67, 0xab, 0x79, 0x7a, 0x95, 0x5d, 0xcc, 0x93, 0xe2,
	0xd9, 0xf4, 0x00, 0xfe, 0x1c, 0x3a, 0x5b, 0x0f, 0x09, 0x6b, 0x69, 0xf5, 0xa9, 0xf4, 0x69, 0xb3,
	0x0d, 0x21, 0x3f, 0x5f, 0xd2, 0x5b, 0xe4, 0x62, 0x0f, 0x27, 0x77, 0x2b, 0x37, 0xcb, 0xe1, 0x48,
	0x0b, 0x12, 0x1d, 0x06, 0x99, 0x71, 0xee, 0x7f, 0x17, 0xf6, 0xa5, 0xab, 0x70, 0x19, 0x94, 0x91,
	0x56, 0x0d, 0x98, 0xf2, 0xca, 0xd9, 0x20, 0x98, 0xeb, 0x77, 0xe4, 0xf6, 0x6f, 0x42, 0xa6, 0x53,
	0xd2, 0x00, 0x7d, 0x42, 0x22, 0xb7, 0xb3, 0xe3, 0x26, 0xdc, 0x9f, 0x6e, 0xcb, 0xf3, 0x9e, 0xda,
	0xc6, 0xce, 0xe3, 0xd1, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x07, 0x6f, 0xe1, 0x0c, 0x00, 0x03,
	0x00, 0x00,
}
