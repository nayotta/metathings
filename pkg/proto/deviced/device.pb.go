// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device.proto

package deviced

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import state "github.com/nayotta/metathings/pkg/proto/constant/state"
import _type "github.com/nayotta/metathings/pkg/proto/constant/type"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Device struct {
	Id                   string             `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type                 _type.DeviceType   `protobuf:"varint,2,opt,name=type,enum=ai.metathings.constant.type.DeviceType" json:"type,omitempty"`
	Name                 string             `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	ProjectId            string             `protobuf:"bytes,4,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	Owners               []string           `protobuf:"bytes,5,rep,name=owners" json:"owners,omitempty"`
	Groups               []string           `protobuf:"bytes,6,rep,name=groups" json:"groups,omitempty"`
	State                state.DeviceState  `protobuf:"varint,7,opt,name=state,enum=ai.metathings.constant.state.DeviceState" json:"state,omitempty"`
	Modules              map[string]*Module `protobuf:"bytes,8,rep,name=modules" json:"modules,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_7a93cbc6f16bec76, []int{0}
}
func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (dst *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(dst, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Device) GetType() _type.DeviceType {
	if m != nil {
		return m.Type
	}
	return _type.DeviceType_DEVICE_TYPE_UNKNOWN
}

func (m *Device) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Device) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *Device) GetOwners() []string {
	if m != nil {
		return m.Owners
	}
	return nil
}

func (m *Device) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *Device) GetState() state.DeviceState {
	if m != nil {
		return m.State
	}
	return state.DeviceState_DEVICE_STATE_UNKNOWN
}

func (m *Device) GetModules() map[string]*Module {
	if m != nil {
		return m.Modules
	}
	return nil
}

type OpDevice struct {
	Id                   *wrappers.StringValue   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type                 _type.DeviceType        `protobuf:"varint,2,opt,name=type,enum=ai.metathings.constant.type.DeviceType" json:"type,omitempty"`
	Name                 *wrappers.StringValue   `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	ProjectId            *wrappers.StringValue   `protobuf:"bytes,4,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	Owners               []*wrappers.StringValue `protobuf:"bytes,5,rep,name=owners" json:"owners,omitempty"`
	Groups               []*wrappers.StringValue `protobuf:"bytes,6,rep,name=groups" json:"groups,omitempty"`
	State                state.DeviceState       `protobuf:"varint,7,opt,name=state,enum=ai.metathings.constant.state.DeviceState" json:"state,omitempty"`
	Modules              map[string]*OpModule    `protobuf:"bytes,8,rep,name=modules" json:"modules,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *OpDevice) Reset()         { *m = OpDevice{} }
func (m *OpDevice) String() string { return proto.CompactTextString(m) }
func (*OpDevice) ProtoMessage()    {}
func (*OpDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_7a93cbc6f16bec76, []int{1}
}
func (m *OpDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpDevice.Unmarshal(m, b)
}
func (m *OpDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpDevice.Marshal(b, m, deterministic)
}
func (dst *OpDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpDevice.Merge(dst, src)
}
func (m *OpDevice) XXX_Size() int {
	return xxx_messageInfo_OpDevice.Size(m)
}
func (m *OpDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_OpDevice.DiscardUnknown(m)
}

var xxx_messageInfo_OpDevice proto.InternalMessageInfo

func (m *OpDevice) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *OpDevice) GetType() _type.DeviceType {
	if m != nil {
		return m.Type
	}
	return _type.DeviceType_DEVICE_TYPE_UNKNOWN
}

func (m *OpDevice) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *OpDevice) GetProjectId() *wrappers.StringValue {
	if m != nil {
		return m.ProjectId
	}
	return nil
}

func (m *OpDevice) GetOwners() []*wrappers.StringValue {
	if m != nil {
		return m.Owners
	}
	return nil
}

func (m *OpDevice) GetGroups() []*wrappers.StringValue {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *OpDevice) GetState() state.DeviceState {
	if m != nil {
		return m.State
	}
	return state.DeviceState_DEVICE_STATE_UNKNOWN
}

func (m *OpDevice) GetModules() map[string]*OpModule {
	if m != nil {
		return m.Modules
	}
	return nil
}

func init() {
	proto.RegisterType((*Device)(nil), "ai.metathings.service.deviced.Device")
	proto.RegisterMapType((map[string]*Module)(nil), "ai.metathings.service.deviced.Device.ModulesEntry")
	proto.RegisterType((*OpDevice)(nil), "ai.metathings.service.deviced.OpDevice")
	proto.RegisterMapType((map[string]*OpModule)(nil), "ai.metathings.service.deviced.OpDevice.ModulesEntry")
}

func init() { proto.RegisterFile("device.proto", fileDescriptor_device_7a93cbc6f16bec76) }

var fileDescriptor_device_7a93cbc6f16bec76 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x5d, 0x8b, 0xd4, 0x30,
	0x14, 0x65, 0xa6, 0x33, 0xb3, 0x3b, 0x99, 0x61, 0x91, 0x3c, 0x48, 0x18, 0x5c, 0x19, 0x16, 0x64,
	0x2b, 0x48, 0x22, 0x75, 0x1e, 0xc4, 0x41, 0x04, 0x3f, 0x1e, 0x04, 0x3f, 0xa0, 0x2b, 0xbe, 0x4a,
	0xa6, 0x8d, 0xdd, 0xb8, 0xd3, 0x24, 0x24, 0xe9, 0x2e, 0xfd, 0x4b, 0xfe, 0x10, 0x7f, 0x97, 0x34,
	0x69, 0xb5, 0xdd, 0x65, 0x0d, 0xa2, 0x6f, 0xb9, 0x1f, 0xe7, 0xdc, 0x9b, 0x7b, 0x0e, 0x58, 0xe6,
	0xec, 0x92, 0x67, 0x0c, 0x2b, 0x2d, 0xad, 0x84, 0xc7, 0x94, 0xe3, 0x92, 0x59, 0x6a, 0xcf, 0xb9,
	0x28, 0x0c, 0x36, 0x4c, 0xbb, 0xa2, 0xef, 0xc9, 0x57, 0xf7, 0x0b, 0x29, 0x8b, 0x3d, 0x23, 0xae,
	0x79, 0x57, 0x7d, 0x25, 0x57, 0x9a, 0x2a, 0xc5, 0xb4, 0xf1, 0xf0, 0xd5, 0xab, 0x82, 0xdb, 0xf3,
	0x6a, 0x87, 0x33, 0x59, 0x12, 0x41, 0x6b, 0x69, 0x2d, 0x25, 0xbf, 0xe9, 0x88, 0xba, 0x28, 0x3c,
	0x94, 0x64, 0x52, 0x18, 0x4b, 0x85, 0x25, 0xc6, 0x52, 0xcb, 0x48, 0x7f, 0x87, 0xd5, 0xcb, 0xbf,
	0x26, 0xb1, 0xb5, 0xba, 0xc6, 0xb1, 0x2c, 0x65, 0x5e, 0xed, 0xdb, 0xe8, 0xe4, 0x7b, 0x04, 0x66,
	0xaf, 0x5d, 0x19, 0x1e, 0x81, 0x31, 0xcf, 0xd1, 0x68, 0x3d, 0x8a, 0xe7, 0xe9, 0x98, 0xe7, 0x70,
	0x0b, 0x26, 0x0d, 0x1a, 0x8d, 0xd7, 0xa3, 0xf8, 0x28, 0x39, 0xc5, 0xc3, 0xff, 0x77, 0x13, 0x70,
	0xd3, 0x83, 0x3d, 0xc5, 0xa7, 0x5a, 0xb1, 0xd4, 0x81, 0x20, 0x04, 0x13, 0x41, 0x4b, 0x86, 0x22,
	0x47, 0xe7, 0xde, 0xf0, 0x18, 0x00, 0xa5, 0xe5, 0x37, 0x96, 0xd9, 0x2f, 0x3c, 0x47, 0x13, 0x57,
	0x99, 0xb7, 0x99, 0xb7, 0x39, 0xbc, 0x0b, 0x66, 0xf2, 0x4a, 0x30, 0x6d, 0xd0, 0x74, 0x1d, 0xc5,
	0xf3, 0xb4, 0x8d, 0x9a, 0x7c, 0xa1, 0x65, 0xa5, 0x0c, 0x9a, 0xf9, 0xbc, 0x8f, 0xe0, 0x0b, 0x30,
	0x75, 0x27, 0x42, 0x07, 0x6e, 0xc1, 0x87, 0xb7, 0x2d, 0xe8, 0x9a, 0xda, 0x0d, 0xcf, 0x9a, 0x77,
	0xea, 0x71, 0xf0, 0x1d, 0x38, 0xf0, 0xb7, 0x30, 0xe8, 0x70, 0x1d, 0xc5, 0x8b, 0x24, 0xc1, 0x7f,
	0xd4, 0xb8, 0xe5, 0xc0, 0xef, 0x3d, 0xe8, 0x8d, 0xb0, 0xba, 0x4e, 0x3b, 0x8a, 0x15, 0x05, 0xcb,
	0x7e, 0x01, 0xde, 0x01, 0xd1, 0x05, 0xab, 0xdb, 0x7b, 0x36, 0x4f, 0xb8, 0x05, 0xd3, 0x4b, 0xba,
	0xaf, 0xfc, 0x45, 0x17, 0xc9, 0x83, 0xc0, 0x34, 0xcf, 0x96, 0x7a, 0xcc, 0xb3, 0xf1, 0xd3, 0xd1,
	0xc9, 0x8f, 0x09, 0x38, 0xfc, 0xa8, 0x5a, 0xb9, 0x1e, 0xfd, 0x92, 0x6b, 0x91, 0xdc, 0xc3, 0xde,
	0x7d, 0xb8, 0x73, 0x1f, 0x3e, 0xb3, 0x9a, 0x8b, 0xe2, 0x73, 0x03, 0xfd, 0x77, 0x31, 0x1f, 0xf7,
	0xc4, 0x0c, 0x0d, 0xf3, 0x52, 0x6f, 0x6f, 0x48, 0x1d, 0xc2, 0xf5, 0x8c, 0xb0, 0x19, 0x18, 0x21,
	0x04, 0xec, 0x6c, 0xb2, 0x19, 0xd8, 0x24, 0x88, 0xfa, 0x5f, 0x26, 0xfa, 0x70, 0xdd, 0x44, 0x9b,
	0x80, 0xac, 0x9d, 0x80, 0xb7, 0xd8, 0x28, 0x0b, 0xda, 0xe8, 0xf9, 0xd0, 0x46, 0xa7, 0xc1, 0x79,
	0x37, 0x8c, 0xb4, 0x9b, 0xb9, 0x9b, 0x3c, 0xf9, 0x19, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xca, 0xe3,
	0xd4, 0xe2, 0x04, 0x00, 0x00,
}
