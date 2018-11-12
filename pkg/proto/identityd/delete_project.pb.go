// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delete_project.proto

package identityd

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

type DeleteProjectRequest struct {
	ProjectId            *wrappers.StringValue `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DeleteProjectRequest) Reset()         { *m = DeleteProjectRequest{} }
func (m *DeleteProjectRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteProjectRequest) ProtoMessage()    {}
func (*DeleteProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7162c34926f205b9, []int{0}
}

func (m *DeleteProjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProjectRequest.Unmarshal(m, b)
}
func (m *DeleteProjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProjectRequest.Marshal(b, m, deterministic)
}
func (m *DeleteProjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProjectRequest.Merge(m, src)
}
func (m *DeleteProjectRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteProjectRequest.Size(m)
}
func (m *DeleteProjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProjectRequest proto.InternalMessageInfo

func (m *DeleteProjectRequest) GetProjectId() *wrappers.StringValue {
	if m != nil {
		return m.ProjectId
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteProjectRequest)(nil), "ai.metathings.service.identityd.DeleteProjectRequest")
}

func init() { proto.RegisterFile("delete_project.proto", fileDescriptor_7162c34926f205b9) }

var fileDescriptor_7162c34926f205b9 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xce, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x06, 0x60, 0xd6, 0x43, 0xc1, 0x78, 0x2b, 0x3d, 0x48, 0x11, 0x5b, 0x3c, 0x79, 0xe9, 0x04,
	0x14, 0x7c, 0x00, 0xf5, 0xe2, 0x4d, 0x2a, 0x78, 0xf1, 0x50, 0xb2, 0x9b, 0x31, 0x1d, 0xcd, 0xee,
	0xc4, 0x64, 0xd2, 0xc5, 0xa7, 0x15, 0x7c, 0x12, 0x21, 0x59, 0x7b, 0x1b, 0x98, 0xff, 0xff, 0xf9,
	0xd4, 0xc2, 0xa2, 0x47, 0xc1, 0x5d, 0x88, 0xfc, 0x81, 0x9d, 0x40, 0x88, 0x2c, 0x3c, 0x5f, 0x19,
	0x82, 0x1e, 0xc5, 0xc8, 0x9e, 0x06, 0x97, 0x20, 0x61, 0x3c, 0x50, 0x87, 0x40, 0x16, 0x07, 0x21,
	0xf9, 0xb6, 0xcb, 0x4b, 0xc7, 0xec, 0x3c, 0xea, 0x12, 0x6f, 0xf3, 0xbb, 0x1e, 0xa3, 0x09, 0x01,
	0x63, 0xaa, 0x03, 0xcb, 0x3b, 0x47, 0xb2, 0xcf, 0x2d, 0x74, 0xdc, 0xeb, 0x7e, 0x24, 0xf9, 0xe4,
	0x51, 0x3b, 0xde, 0x94, 0xe7, 0xe6, 0x60, 0x3c, 0x59, 0x23, 0x1c, 0x93, 0x3e, 0x9e, 0xb5, 0x77,
	0xf5, 0xa6, 0x16, 0x8f, 0x05, 0xf4, 0x5c, 0x3d, 0x5b, 0xfc, 0xca, 0x98, 0x64, 0xfe, 0xa0, 0xd4,
	0x24, 0xdc, 0x91, 0x3d, 0x6f, 0xd6, 0xcd, 0xf5, 0xd9, 0xcd, 0x05, 0x54, 0x04, 0xfc, 0x23, 0xe0,
	0x45, 0x22, 0x0d, 0xee, 0xd5, 0xf8, 0x8c, 0xf7, 0xb3, 0xdf, 0x9f, 0xd5, 0xc9, 0xba, 0xd9, 0x9e,
	0x4e, 0xbd, 0x27, 0xdb, 0xce, 0x4a, 0xf0, 0xf6, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x2d, 0x2e,
	0xb8, 0xf4, 0x00, 0x00, 0x00,
}
