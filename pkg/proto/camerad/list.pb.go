// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list.proto

package camerad

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/mwitkow/go-proto-validators"
	camera "github.com/nayotta/metathings/pkg/proto/camera"
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

type ListRequest struct {
	Name                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CoreId               *wrappers.StringValue `protobuf:"bytes,2,opt,name=core_id,json=coreId,proto3" json:"core_id,omitempty"`
	EntityName           *wrappers.StringValue `protobuf:"bytes,3,opt,name=entity_name,json=entityName,proto3" json:"entity_name,omitempty"`
	OwnerId              *wrappers.StringValue `protobuf:"bytes,4,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	State                camera.CameraState    `protobuf:"varint,5,opt,name=state,proto3,enum=ai.metathings.service.camera.CameraState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{0}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ListRequest) GetCoreId() *wrappers.StringValue {
	if m != nil {
		return m.CoreId
	}
	return nil
}

func (m *ListRequest) GetEntityName() *wrappers.StringValue {
	if m != nil {
		return m.EntityName
	}
	return nil
}

func (m *ListRequest) GetOwnerId() *wrappers.StringValue {
	if m != nil {
		return m.OwnerId
	}
	return nil
}

func (m *ListRequest) GetState() camera.CameraState {
	if m != nil {
		return m.State
	}
	return camera.CameraState_CAMERA_STATE_UNKNOWN
}

type ListResponse struct {
	Cameras              []*Camera `protobuf:"bytes,1,rep,name=cameras,proto3" json:"cameras,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{1}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetCameras() []*Camera {
	if m != nil {
		return m.Cameras
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "ai.metathings.service.camerad.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "ai.metathings.service.camerad.ListResponse")
}

func init() { proto.RegisterFile("list.proto", fileDescriptor_af793ce248ee1bf0) }

var fileDescriptor_af793ce248ee1bf0 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x3f, 0x6b, 0xf3, 0x30,
	0x10, 0xc6, 0x71, 0xfe, 0xbe, 0xc8, 0xe1, 0x1d, 0x3c, 0x99, 0xd0, 0x96, 0x10, 0x28, 0xa4, 0x43,
	0xa4, 0x92, 0xd2, 0x76, 0x28, 0x25, 0x43, 0xa7, 0x40, 0x69, 0xc1, 0x81, 0xae, 0x41, 0xb1, 0xaf,
	0x8e, 0x88, 0x2d, 0xb9, 0xd2, 0x39, 0x26, 0x9f, 0xaa, 0x5f, 0xb1, 0x58, 0x72, 0xda, 0x4c, 0x21,
	0xd3, 0x49, 0xe8, 0xf9, 0xdd, 0xf3, 0x9c, 0x8e, 0x90, 0x4c, 0x18, 0xa4, 0x85, 0x56, 0xa8, 0x82,
	0x4b, 0x2e, 0x68, 0x0e, 0xc8, 0x71, 0x23, 0x64, 0x6a, 0xa8, 0x01, 0xbd, 0x13, 0x31, 0xd0, 0x98,
	0xe7, 0xa0, 0x79, 0x32, 0xbc, 0x4a, 0x95, 0x4a, 0x33, 0x60, 0x56, 0xbc, 0x2e, 0x3f, 0x59, 0xa5,
	0x79, 0x51, 0x80, 0x36, 0x0e, 0x1f, 0x3e, 0xa4, 0x02, 0x37, 0xe5, 0x9a, 0xc6, 0x2a, 0x67, 0x79,
	0x25, 0x70, 0xab, 0x2a, 0x96, 0xaa, 0xa9, 0x7d, 0x9c, 0xee, 0x78, 0x26, 0x12, 0x8e, 0x4a, 0x1b,
	0xf6, 0x7b, 0x6c, 0xb8, 0xa7, 0x23, 0x4e, 0xf2, 0xbd, 0x42, 0xe4, 0xec, 0x2f, 0x06, 0x2b, 0xb6,
	0xa9, 0xb3, 0x64, 0x2e, 0x48, 0x53, 0x1a, 0x78, 0x70, 0x7c, 0x1b, 0x7f, 0xb7, 0x88, 0xff, 0x2a,
	0x0c, 0x46, 0xf0, 0x55, 0x82, 0xc1, 0xe0, 0x96, 0x74, 0x24, 0xcf, 0x21, 0xf4, 0x46, 0xde, 0xc4,
	0x9f, 0x5d, 0x50, 0x37, 0x01, 0x3d, 0x4c, 0x40, 0x97, 0xa8, 0x85, 0x4c, 0x3f, 0x78, 0x56, 0x42,
	0x64, 0x95, 0xc1, 0x3d, 0xe9, 0xc7, 0x4a, 0xc3, 0x4a, 0x24, 0x61, 0xeb, 0x0c, 0xa8, 0x57, 0x8b,
	0x17, 0x49, 0xf0, 0x4c, 0x7c, 0x90, 0x28, 0x70, 0xbf, 0xb2, 0x7e, 0xed, 0x33, 0x50, 0xe2, 0x80,
	0xb7, 0xda, 0xf5, 0x91, 0xfc, 0x53, 0x95, 0x04, 0x5d, 0xdb, 0x76, 0xce, 0x60, 0xfb, 0x56, 0xbd,
	0x48, 0x82, 0x39, 0xe9, 0x1a, 0xe4, 0x08, 0x61, 0x77, 0xe4, 0x4d, 0xfe, 0xcf, 0x6e, 0xe8, 0xa9,
	0x15, 0xd2, 0x17, 0x5b, 0x96, 0x35, 0x10, 0x39, 0x6e, 0xfc, 0x4e, 0x06, 0xee, 0xc3, 0x4c, 0xa1,
	0xa4, 0x81, 0x60, 0x4e, 0xfa, 0x4e, 0x6c, 0x42, 0x6f, 0xd4, 0x9e, 0xf8, 0xb3, 0xeb, 0x93, 0x2d,
	0x93, 0xa6, 0x67, 0x74, 0xa0, 0xd6, 0x3d, 0x1b, 0xf8, 0xee, 0x27, 0x00, 0x00, 0xff, 0xff, 0xcb,
	0x03, 0x1d, 0x6d, 0x59, 0x02, 0x00, 0x00,
}
