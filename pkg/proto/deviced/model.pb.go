// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

package deviced

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	kind "github.com/nayotta/metathings/pkg/proto/constant/kind"
	state "github.com/nayotta/metathings/pkg/proto/constant/state"
	_ "github.com/nayotta/metathings/pkg/proto/identityd2"
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

type Device struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Kind                 kind.DeviceKind   `protobuf:"varint,2,opt,name=kind,proto3,enum=ai.metathings.constant.kind.DeviceKind" json:"kind,omitempty"`
	State                state.DeviceState `protobuf:"varint,3,opt,name=state,proto3,enum=ai.metathings.constant.state.DeviceState" json:"state,omitempty"`
	Name                 string            `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Alias                string            `protobuf:"bytes,5,opt,name=alias,proto3" json:"alias,omitempty"`
	Modules              []*Module         `protobuf:"bytes,6,rep,name=modules,proto3" json:"modules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{0}
}

func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (m *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(m, src)
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

func (m *Device) GetKind() kind.DeviceKind {
	if m != nil {
		return m.Kind
	}
	return kind.DeviceKind_DEVICE_KIND_UNKNOWN
}

func (m *Device) GetState() state.DeviceState {
	if m != nil {
		return m.State
	}
	return state.DeviceState_DEVICE_STATE_UNKNOWN
}

func (m *Device) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Device) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *Device) GetModules() []*Module {
	if m != nil {
		return m.Modules
	}
	return nil
}

type OpDevice struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Kind                 kind.DeviceKind       `protobuf:"varint,2,opt,name=kind,proto3,enum=ai.metathings.constant.kind.DeviceKind" json:"kind,omitempty"`
	State                state.DeviceState     `protobuf:"varint,3,opt,name=state,proto3,enum=ai.metathings.constant.state.DeviceState" json:"state,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Alias                *wrappers.StringValue `protobuf:"bytes,5,opt,name=alias,proto3" json:"alias,omitempty"`
	Modules              []*OpModule           `protobuf:"bytes,6,rep,name=modules,proto3" json:"modules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OpDevice) Reset()         { *m = OpDevice{} }
func (m *OpDevice) String() string { return proto.CompactTextString(m) }
func (*OpDevice) ProtoMessage()    {}
func (*OpDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{1}
}

func (m *OpDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpDevice.Unmarshal(m, b)
}
func (m *OpDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpDevice.Marshal(b, m, deterministic)
}
func (m *OpDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpDevice.Merge(m, src)
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

func (m *OpDevice) GetKind() kind.DeviceKind {
	if m != nil {
		return m.Kind
	}
	return kind.DeviceKind_DEVICE_KIND_UNKNOWN
}

func (m *OpDevice) GetState() state.DeviceState {
	if m != nil {
		return m.State
	}
	return state.DeviceState_DEVICE_STATE_UNKNOWN
}

func (m *OpDevice) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *OpDevice) GetAlias() *wrappers.StringValue {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *OpDevice) GetModules() []*OpModule {
	if m != nil {
		return m.Modules
	}
	return nil
}

type Module struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	State                state.ModuleState `protobuf:"varint,2,opt,name=state,proto3,enum=ai.metathings.constant.state.ModuleState" json:"state,omitempty"`
	DeviceId             string            `protobuf:"bytes,3,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Endpoint             string            `protobuf:"bytes,4,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Name                 string            `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Alias                string            `protobuf:"bytes,6,opt,name=alias,proto3" json:"alias,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Module) Reset()         { *m = Module{} }
func (m *Module) String() string { return proto.CompactTextString(m) }
func (*Module) ProtoMessage()    {}
func (*Module) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{2}
}

func (m *Module) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Module.Unmarshal(m, b)
}
func (m *Module) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Module.Marshal(b, m, deterministic)
}
func (m *Module) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Module.Merge(m, src)
}
func (m *Module) XXX_Size() int {
	return xxx_messageInfo_Module.Size(m)
}
func (m *Module) XXX_DiscardUnknown() {
	xxx_messageInfo_Module.DiscardUnknown(m)
}

var xxx_messageInfo_Module proto.InternalMessageInfo

func (m *Module) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Module) GetState() state.ModuleState {
	if m != nil {
		return m.State
	}
	return state.ModuleState_MODULE_STATE_UNKNOWN
}

func (m *Module) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *Module) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *Module) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Module) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

type OpModule struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	State                state.ModuleState     `protobuf:"varint,2,opt,name=state,proto3,enum=ai.metathings.constant.state.ModuleState" json:"state,omitempty"`
	DeviceId             *wrappers.StringValue `protobuf:"bytes,3,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Endpoint             *wrappers.StringValue `protobuf:"bytes,4,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Alias                *wrappers.StringValue `protobuf:"bytes,6,opt,name=alias,proto3" json:"alias,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OpModule) Reset()         { *m = OpModule{} }
func (m *OpModule) String() string { return proto.CompactTextString(m) }
func (*OpModule) ProtoMessage()    {}
func (*OpModule) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{3}
}

func (m *OpModule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpModule.Unmarshal(m, b)
}
func (m *OpModule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpModule.Marshal(b, m, deterministic)
}
func (m *OpModule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpModule.Merge(m, src)
}
func (m *OpModule) XXX_Size() int {
	return xxx_messageInfo_OpModule.Size(m)
}
func (m *OpModule) XXX_DiscardUnknown() {
	xxx_messageInfo_OpModule.DiscardUnknown(m)
}

var xxx_messageInfo_OpModule proto.InternalMessageInfo

func (m *OpModule) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *OpModule) GetState() state.ModuleState {
	if m != nil {
		return m.State
	}
	return state.ModuleState_MODULE_STATE_UNKNOWN
}

func (m *OpModule) GetDeviceId() *wrappers.StringValue {
	if m != nil {
		return m.DeviceId
	}
	return nil
}

func (m *OpModule) GetEndpoint() *wrappers.StringValue {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *OpModule) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *OpModule) GetAlias() *wrappers.StringValue {
	if m != nil {
		return m.Alias
	}
	return nil
}

type ErrorValue struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ServiceName          string   `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	MethodName           string   `protobuf:"bytes,3,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	Context              string   `protobuf:"bytes,4,opt,name=context,proto3" json:"context,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorValue) Reset()         { *m = ErrorValue{} }
func (m *ErrorValue) String() string { return proto.CompactTextString(m) }
func (*ErrorValue) ProtoMessage()    {}
func (*ErrorValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{4}
}

func (m *ErrorValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorValue.Unmarshal(m, b)
}
func (m *ErrorValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorValue.Marshal(b, m, deterministic)
}
func (m *ErrorValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorValue.Merge(m, src)
}
func (m *ErrorValue) XXX_Size() int {
	return xxx_messageInfo_ErrorValue.Size(m)
}
func (m *ErrorValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorValue.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorValue proto.InternalMessageInfo

func (m *ErrorValue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ErrorValue) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *ErrorValue) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

func (m *ErrorValue) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

type OpUnaryCallValue struct {
	ModuleName           *wrappers.StringValue `protobuf:"bytes,1,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	ComponentName        *wrappers.StringValue `protobuf:"bytes,2,opt,name=component_name,json=componentName,proto3" json:"component_name,omitempty"`
	MethodName           *wrappers.StringValue `protobuf:"bytes,3,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	Value                *any.Any              `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OpUnaryCallValue) Reset()         { *m = OpUnaryCallValue{} }
func (m *OpUnaryCallValue) String() string { return proto.CompactTextString(m) }
func (*OpUnaryCallValue) ProtoMessage()    {}
func (*OpUnaryCallValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{5}
}

func (m *OpUnaryCallValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpUnaryCallValue.Unmarshal(m, b)
}
func (m *OpUnaryCallValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpUnaryCallValue.Marshal(b, m, deterministic)
}
func (m *OpUnaryCallValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpUnaryCallValue.Merge(m, src)
}
func (m *OpUnaryCallValue) XXX_Size() int {
	return xxx_messageInfo_OpUnaryCallValue.Size(m)
}
func (m *OpUnaryCallValue) XXX_DiscardUnknown() {
	xxx_messageInfo_OpUnaryCallValue.DiscardUnknown(m)
}

var xxx_messageInfo_OpUnaryCallValue proto.InternalMessageInfo

func (m *OpUnaryCallValue) GetModuleName() *wrappers.StringValue {
	if m != nil {
		return m.ModuleName
	}
	return nil
}

func (m *OpUnaryCallValue) GetComponentName() *wrappers.StringValue {
	if m != nil {
		return m.ComponentName
	}
	return nil
}

func (m *OpUnaryCallValue) GetMethodName() *wrappers.StringValue {
	if m != nil {
		return m.MethodName
	}
	return nil
}

func (m *OpUnaryCallValue) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

type UnaryCallValue struct {
	ModuleName           string   `protobuf:"bytes,1,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	ComponentName        string   `protobuf:"bytes,2,opt,name=component_name,json=componentName,proto3" json:"component_name,omitempty"`
	MethodName           string   `protobuf:"bytes,3,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	Value                *any.Any `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnaryCallValue) Reset()         { *m = UnaryCallValue{} }
func (m *UnaryCallValue) String() string { return proto.CompactTextString(m) }
func (*UnaryCallValue) ProtoMessage()    {}
func (*UnaryCallValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{6}
}

func (m *UnaryCallValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnaryCallValue.Unmarshal(m, b)
}
func (m *UnaryCallValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnaryCallValue.Marshal(b, m, deterministic)
}
func (m *UnaryCallValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnaryCallValue.Merge(m, src)
}
func (m *UnaryCallValue) XXX_Size() int {
	return xxx_messageInfo_UnaryCallValue.Size(m)
}
func (m *UnaryCallValue) XXX_DiscardUnknown() {
	xxx_messageInfo_UnaryCallValue.DiscardUnknown(m)
}

var xxx_messageInfo_UnaryCallValue proto.InternalMessageInfo

func (m *UnaryCallValue) GetModuleName() string {
	if m != nil {
		return m.ModuleName
	}
	return ""
}

func (m *UnaryCallValue) GetComponentName() string {
	if m != nil {
		return m.ComponentName
	}
	return ""
}

func (m *UnaryCallValue) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

func (m *UnaryCallValue) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

type OpStreamCallValue struct {
	// Types that are valid to be assigned to Union:
	//	*OpStreamCallValue_Config
	//	*OpStreamCallValue_Data
	Union                isOpStreamCallValue_Union `protobuf_oneof:"union"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *OpStreamCallValue) Reset()         { *m = OpStreamCallValue{} }
func (m *OpStreamCallValue) String() string { return proto.CompactTextString(m) }
func (*OpStreamCallValue) ProtoMessage()    {}
func (*OpStreamCallValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{7}
}

func (m *OpStreamCallValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpStreamCallValue.Unmarshal(m, b)
}
func (m *OpStreamCallValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpStreamCallValue.Marshal(b, m, deterministic)
}
func (m *OpStreamCallValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpStreamCallValue.Merge(m, src)
}
func (m *OpStreamCallValue) XXX_Size() int {
	return xxx_messageInfo_OpStreamCallValue.Size(m)
}
func (m *OpStreamCallValue) XXX_DiscardUnknown() {
	xxx_messageInfo_OpStreamCallValue.DiscardUnknown(m)
}

var xxx_messageInfo_OpStreamCallValue proto.InternalMessageInfo

type isOpStreamCallValue_Union interface {
	isOpStreamCallValue_Union()
}

type OpStreamCallValue_Config struct {
	Config *OpStreamCallConfig `protobuf:"bytes,1,opt,name=config,proto3,oneof"`
}

type OpStreamCallValue_Data struct {
	Data *OpStreamCallData `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

func (*OpStreamCallValue_Config) isOpStreamCallValue_Union() {}

func (*OpStreamCallValue_Data) isOpStreamCallValue_Union() {}

func (m *OpStreamCallValue) GetUnion() isOpStreamCallValue_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (m *OpStreamCallValue) GetConfig() *OpStreamCallConfig {
	if x, ok := m.GetUnion().(*OpStreamCallValue_Config); ok {
		return x.Config
	}
	return nil
}

func (m *OpStreamCallValue) GetData() *OpStreamCallData {
	if x, ok := m.GetUnion().(*OpStreamCallValue_Data); ok {
		return x.Data
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*OpStreamCallValue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _OpStreamCallValue_OneofMarshaler, _OpStreamCallValue_OneofUnmarshaler, _OpStreamCallValue_OneofSizer, []interface{}{
		(*OpStreamCallValue_Config)(nil),
		(*OpStreamCallValue_Data)(nil),
	}
}

func _OpStreamCallValue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*OpStreamCallValue)
	// union
	switch x := m.Union.(type) {
	case *OpStreamCallValue_Config:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Config); err != nil {
			return err
		}
	case *OpStreamCallValue_Data:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Data); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("OpStreamCallValue.Union has unexpected type %T", x)
	}
	return nil
}

func _OpStreamCallValue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*OpStreamCallValue)
	switch tag {
	case 1: // union.config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(OpStreamCallConfig)
		err := b.DecodeMessage(msg)
		m.Union = &OpStreamCallValue_Config{msg}
		return true, err
	case 2: // union.data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(OpStreamCallData)
		err := b.DecodeMessage(msg)
		m.Union = &OpStreamCallValue_Data{msg}
		return true, err
	default:
		return false, nil
	}
}

func _OpStreamCallValue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*OpStreamCallValue)
	// union
	switch x := m.Union.(type) {
	case *OpStreamCallValue_Config:
		s := proto.Size(x.Config)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *OpStreamCallValue_Data:
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

type StreamCallValue struct {
	// Types that are valid to be assigned to Union:
	//	*StreamCallValue_Config
	//	*StreamCallValue_Data
	Union                isStreamCallValue_Union `protobuf_oneof:"union"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *StreamCallValue) Reset()         { *m = StreamCallValue{} }
func (m *StreamCallValue) String() string { return proto.CompactTextString(m) }
func (*StreamCallValue) ProtoMessage()    {}
func (*StreamCallValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{8}
}

func (m *StreamCallValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamCallValue.Unmarshal(m, b)
}
func (m *StreamCallValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamCallValue.Marshal(b, m, deterministic)
}
func (m *StreamCallValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamCallValue.Merge(m, src)
}
func (m *StreamCallValue) XXX_Size() int {
	return xxx_messageInfo_StreamCallValue.Size(m)
}
func (m *StreamCallValue) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamCallValue.DiscardUnknown(m)
}

var xxx_messageInfo_StreamCallValue proto.InternalMessageInfo

type isStreamCallValue_Union interface {
	isStreamCallValue_Union()
}

type StreamCallValue_Config struct {
	Config *StreamCallConfig `protobuf:"bytes,1,opt,name=config,proto3,oneof"`
}

type StreamCallValue_Data struct {
	Data *StreamCallData `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

func (*StreamCallValue_Config) isStreamCallValue_Union() {}

func (*StreamCallValue_Data) isStreamCallValue_Union() {}

func (m *StreamCallValue) GetUnion() isStreamCallValue_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (m *StreamCallValue) GetConfig() *StreamCallConfig {
	if x, ok := m.GetUnion().(*StreamCallValue_Config); ok {
		return x.Config
	}
	return nil
}

func (m *StreamCallValue) GetData() *StreamCallData {
	if x, ok := m.GetUnion().(*StreamCallValue_Data); ok {
		return x.Data
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StreamCallValue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StreamCallValue_OneofMarshaler, _StreamCallValue_OneofUnmarshaler, _StreamCallValue_OneofSizer, []interface{}{
		(*StreamCallValue_Config)(nil),
		(*StreamCallValue_Data)(nil),
	}
}

func _StreamCallValue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StreamCallValue)
	// union
	switch x := m.Union.(type) {
	case *StreamCallValue_Config:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Config); err != nil {
			return err
		}
	case *StreamCallValue_Data:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Data); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StreamCallValue.Union has unexpected type %T", x)
	}
	return nil
}

func _StreamCallValue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StreamCallValue)
	switch tag {
	case 1: // union.config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamCallConfig)
		err := b.DecodeMessage(msg)
		m.Union = &StreamCallValue_Config{msg}
		return true, err
	case 2: // union.data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamCallData)
		err := b.DecodeMessage(msg)
		m.Union = &StreamCallValue_Data{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StreamCallValue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StreamCallValue)
	// union
	switch x := m.Union.(type) {
	case *StreamCallValue_Config:
		s := proto.Size(x.Config)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamCallValue_Data:
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

type OpStreamCallConfig struct {
	ModuleName           *wrappers.StringValue `protobuf:"bytes,1,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	ComponentName        *wrappers.StringValue `protobuf:"bytes,2,opt,name=component_name,json=componentName,proto3" json:"component_name,omitempty"`
	MethodName           *wrappers.StringValue `protobuf:"bytes,3,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OpStreamCallConfig) Reset()         { *m = OpStreamCallConfig{} }
func (m *OpStreamCallConfig) String() string { return proto.CompactTextString(m) }
func (*OpStreamCallConfig) ProtoMessage()    {}
func (*OpStreamCallConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{9}
}

func (m *OpStreamCallConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpStreamCallConfig.Unmarshal(m, b)
}
func (m *OpStreamCallConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpStreamCallConfig.Marshal(b, m, deterministic)
}
func (m *OpStreamCallConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpStreamCallConfig.Merge(m, src)
}
func (m *OpStreamCallConfig) XXX_Size() int {
	return xxx_messageInfo_OpStreamCallConfig.Size(m)
}
func (m *OpStreamCallConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OpStreamCallConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OpStreamCallConfig proto.InternalMessageInfo

func (m *OpStreamCallConfig) GetModuleName() *wrappers.StringValue {
	if m != nil {
		return m.ModuleName
	}
	return nil
}

func (m *OpStreamCallConfig) GetComponentName() *wrappers.StringValue {
	if m != nil {
		return m.ComponentName
	}
	return nil
}

func (m *OpStreamCallConfig) GetMethodName() *wrappers.StringValue {
	if m != nil {
		return m.MethodName
	}
	return nil
}

type StreamCallConfig struct {
	ModuleName           string   `protobuf:"bytes,1,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	ComponentName        string   `protobuf:"bytes,2,opt,name=component_name,json=componentName,proto3" json:"component_name,omitempty"`
	MethodName           string   `protobuf:"bytes,3,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamCallConfig) Reset()         { *m = StreamCallConfig{} }
func (m *StreamCallConfig) String() string { return proto.CompactTextString(m) }
func (*StreamCallConfig) ProtoMessage()    {}
func (*StreamCallConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{10}
}

func (m *StreamCallConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamCallConfig.Unmarshal(m, b)
}
func (m *StreamCallConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamCallConfig.Marshal(b, m, deterministic)
}
func (m *StreamCallConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamCallConfig.Merge(m, src)
}
func (m *StreamCallConfig) XXX_Size() int {
	return xxx_messageInfo_StreamCallConfig.Size(m)
}
func (m *StreamCallConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamCallConfig.DiscardUnknown(m)
}

var xxx_messageInfo_StreamCallConfig proto.InternalMessageInfo

func (m *StreamCallConfig) GetModuleName() string {
	if m != nil {
		return m.ModuleName
	}
	return ""
}

func (m *StreamCallConfig) GetComponentName() string {
	if m != nil {
		return m.ComponentName
	}
	return ""
}

func (m *StreamCallConfig) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

type OpStreamCallData struct {
	Value                *any.Any `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpStreamCallData) Reset()         { *m = OpStreamCallData{} }
func (m *OpStreamCallData) String() string { return proto.CompactTextString(m) }
func (*OpStreamCallData) ProtoMessage()    {}
func (*OpStreamCallData) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{11}
}

func (m *OpStreamCallData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpStreamCallData.Unmarshal(m, b)
}
func (m *OpStreamCallData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpStreamCallData.Marshal(b, m, deterministic)
}
func (m *OpStreamCallData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpStreamCallData.Merge(m, src)
}
func (m *OpStreamCallData) XXX_Size() int {
	return xxx_messageInfo_OpStreamCallData.Size(m)
}
func (m *OpStreamCallData) XXX_DiscardUnknown() {
	xxx_messageInfo_OpStreamCallData.DiscardUnknown(m)
}

var xxx_messageInfo_OpStreamCallData proto.InternalMessageInfo

func (m *OpStreamCallData) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

type StreamCallData struct {
	Value                *any.Any `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamCallData) Reset()         { *m = StreamCallData{} }
func (m *StreamCallData) String() string { return proto.CompactTextString(m) }
func (*StreamCallData) ProtoMessage()    {}
func (*StreamCallData) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{12}
}

func (m *StreamCallData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamCallData.Unmarshal(m, b)
}
func (m *StreamCallData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamCallData.Marshal(b, m, deterministic)
}
func (m *StreamCallData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamCallData.Merge(m, src)
}
func (m *StreamCallData) XXX_Size() int {
	return xxx_messageInfo_StreamCallData.Size(m)
}
func (m *StreamCallData) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamCallData.DiscardUnknown(m)
}

var xxx_messageInfo_StreamCallData proto.InternalMessageInfo

func (m *StreamCallData) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Device)(nil), "ai.metathings.service.deviced.Device")
	proto.RegisterType((*OpDevice)(nil), "ai.metathings.service.deviced.OpDevice")
	proto.RegisterType((*Module)(nil), "ai.metathings.service.deviced.Module")
	proto.RegisterType((*OpModule)(nil), "ai.metathings.service.deviced.OpModule")
	proto.RegisterType((*ErrorValue)(nil), "ai.metathings.service.deviced.ErrorValue")
	proto.RegisterType((*OpUnaryCallValue)(nil), "ai.metathings.service.deviced.OpUnaryCallValue")
	proto.RegisterType((*UnaryCallValue)(nil), "ai.metathings.service.deviced.UnaryCallValue")
	proto.RegisterType((*OpStreamCallValue)(nil), "ai.metathings.service.deviced.OpStreamCallValue")
	proto.RegisterType((*StreamCallValue)(nil), "ai.metathings.service.deviced.StreamCallValue")
	proto.RegisterType((*OpStreamCallConfig)(nil), "ai.metathings.service.deviced.OpStreamCallConfig")
	proto.RegisterType((*StreamCallConfig)(nil), "ai.metathings.service.deviced.StreamCallConfig")
	proto.RegisterType((*OpStreamCallData)(nil), "ai.metathings.service.deviced.OpStreamCallData")
	proto.RegisterType((*StreamCallData)(nil), "ai.metathings.service.deviced.StreamCallData")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_4c16552f9fdb66d8) }

var fileDescriptor_4c16552f9fdb66d8 = []byte{
	// 755 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x55, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xad, 0xd3, 0xfc, 0x34, 0x37, 0xdf, 0x17, 0x8a, 0xd5, 0x45, 0x08, 0x3f, 0x2d, 0x96, 0xaa,
	0x16, 0x04, 0x76, 0x09, 0x1b, 0x10, 0xb4, 0x55, 0x9b, 0x56, 0x6a, 0x55, 0x41, 0x24, 0x57, 0xb0,
	0x60, 0x53, 0x4d, 0x33, 0xd3, 0x74, 0x54, 0x7b, 0xc6, 0x72, 0x26, 0x85, 0x08, 0x09, 0x89, 0x87,
	0x61, 0xc3, 0x8a, 0x07, 0xe0, 0x31, 0x90, 0x58, 0xf2, 0x2a, 0xc8, 0x33, 0x13, 0x3b, 0x71, 0x4a,
	0x9d, 0x14, 0x09, 0x21, 0x56, 0xb5, 0x67, 0xee, 0x39, 0xf7, 0xdc, 0x73, 0x4f, 0x1d, 0xa8, 0xf8,
	0x1c, 0x13, 0xcf, 0x0e, 0x42, 0x2e, 0xb8, 0x79, 0x1b, 0x51, 0xdb, 0x27, 0x02, 0x89, 0x53, 0xca,
	0x3a, 0x5d, 0xbb, 0x4b, 0xc2, 0x73, 0xda, 0x26, 0x36, 0x26, 0xd1, 0x1f, 0x5c, 0xbf, 0xd3, 0xe1,
	0xbc, 0xe3, 0x11, 0x47, 0x16, 0x1f, 0xf7, 0x4e, 0x9c, 0xb7, 0x21, 0x0a, 0x02, 0x12, 0x76, 0x15,
	0xbc, 0x7e, 0x23, 0x7d, 0x8f, 0x58, 0x5f, 0x5f, 0x6d, 0x74, 0xa8, 0x38, 0xed, 0x1d, 0xdb, 0x6d,
	0xee, 0x3b, 0x0c, 0xf5, 0xb9, 0x10, 0xc8, 0x49, 0x3a, 0x39, 0xc1, 0x59, 0x47, 0xa1, 0x1c, 0x8a,
	0x09, 0x13, 0x54, 0xf4, 0x71, 0xc3, 0x19, 0x52, 0x56, 0x6f, 0x4e, 0x8a, 0x6f, 0x73, 0xd6, 0x15,
	0x88, 0x09, 0xa7, 0x2b, 0x90, 0x20, 0x8e, 0x92, 0xae, 0x49, 0xb6, 0xa7, 0x26, 0x39, 0xa3, 0x0c,
	0x8f, 0x72, 0x5c, 0x55, 0x88, 0xcf, 0x71, 0xcf, 0xd3, 0x24, 0xd6, 0xc7, 0x1c, 0x14, 0x77, 0x24,
	0xab, 0x59, 0x85, 0x1c, 0xc5, 0x35, 0x63, 0xc9, 0x58, 0x2d, 0xbb, 0x39, 0x8a, 0xcd, 0x67, 0x90,
	0x8f, 0x9a, 0xd6, 0x72, 0x4b, 0xc6, 0x6a, 0xb5, 0xb1, 0x62, 0x8f, 0x6e, 0x64, 0x40, 0x6a, 0x47,
	0x35, 0xb6, 0xa2, 0x38, 0xa0, 0x0c, 0xbb, 0x12, 0x64, 0x6e, 0x42, 0x41, 0x76, 0xab, 0xcd, 0x4a,
	0xf4, 0xbd, 0x5f, 0xa1, 0x65, 0x91, 0x86, 0x1f, 0x46, 0xcf, 0xae, 0xc2, 0x99, 0x26, 0xe4, 0x19,
	0xf2, 0x49, 0x2d, 0x2f, 0xf5, 0xc8, 0x67, 0x73, 0x01, 0x0a, 0xc8, 0xa3, 0xa8, 0x5b, 0x2b, 0xc8,
	0x43, 0xf5, 0x62, 0x6e, 0x42, 0x49, 0x8d, 0xd4, 0xad, 0x15, 0x97, 0x66, 0x57, 0x2b, 0x8d, 0x65,
	0xfb, 0xd2, 0xf0, 0xd8, 0x2f, 0x64, 0xb5, 0x3b, 0x40, 0x59, 0x3f, 0x72, 0x30, 0xd7, 0x0a, 0xb4,
	0x0b, 0x0f, 0x62, 0x17, 0x2a, 0x8d, 0x5b, 0xb6, 0x8a, 0x91, 0x3d, 0x88, 0x91, 0x7d, 0x28, 0x42,
	0xca, 0x3a, 0xaf, 0x91, 0xd7, 0x23, 0x7f, 0x81, 0x47, 0x6b, 0x43, 0x1e, 0x65, 0xa9, 0x55, 0x0e,
	0x36, 0x86, 0x1d, 0xcc, 0x82, 0x68, 0x7f, 0xb7, 0xd2, 0xfe, 0xae, 0x64, 0xf8, 0xdb, 0x0a, 0xd2,
	0x0e, 0x7f, 0x35, 0xa0, 0xa8, 0xce, 0xc6, 0x52, 0x16, 0x9b, 0x90, 0x9b, 0xc4, 0x04, 0x45, 0x32,
	0x62, 0xc2, 0x4d, 0x28, 0xab, 0xc6, 0x47, 0x14, 0x4b, 0x27, 0xcb, 0xee, 0x9c, 0x3a, 0xd8, 0xc7,
	0x66, 0x1d, 0xe6, 0x08, 0xc3, 0x01, 0xa7, 0x4c, 0xe8, 0x24, 0xc5, 0xef, 0x71, 0xc2, 0x0a, 0x17,
	0x25, 0xac, 0x38, 0x94, 0x30, 0xeb, 0x9b, 0x0c, 0x88, 0x1e, 0x60, 0xba, 0x80, 0xfc, 0xf6, 0x78,
	0x4f, 0xd3, 0xe3, 0x65, 0x75, 0x4d, 0x86, 0x7f, 0x92, 0x1a, 0x3e, 0x13, 0x19, 0x5b, 0xb3, 0x36,
	0x64, 0xcd, 0x94, 0xc1, 0x2a, 0x4e, 0x1c, 0x2c, 0xeb, 0x03, 0xc0, 0x6e, 0x18, 0xf2, 0x50, 0x1e,
	0xc6, 0xeb, 0x30, 0x86, 0xd6, 0x71, 0x17, 0xfe, 0xd3, 0xe1, 0x3a, 0x92, 0x77, 0x39, 0x79, 0x57,
	0xd1, 0x67, 0x2f, 0xa3, 0x92, 0x45, 0xa8, 0xf8, 0x44, 0x9c, 0x72, 0xac, 0x2a, 0x54, 0x00, 0x40,
	0x1d, 0xc9, 0x82, 0x1a, 0x94, 0xda, 0x9c, 0x09, 0xf2, 0x6e, 0x90, 0x80, 0xc1, 0x6b, 0xf4, 0xed,
	0x9b, 0x6f, 0x05, 0xaf, 0x18, 0x0a, 0xfb, 0x4d, 0xe4, 0x79, 0x4a, 0xc6, 0xba, 0xfc, 0x1d, 0xea,
	0x79, 0xba, 0xe3, 0x24, 0x7b, 0x06, 0x05, 0x90, 0xdd, 0x9a, 0x50, 0x6d, 0x73, 0x3f, 0xe0, 0x8c,
	0x30, 0x91, 0x68, 0xce, 0x62, 0xf8, 0x3f, 0xc6, 0x48, 0x92, 0xf5, 0xf1, 0x99, 0xb2, 0x35, 0x24,
	0x13, 0xdf, 0x87, 0xc2, 0x79, 0x74, 0xa8, 0x97, 0xbe, 0x30, 0x06, 0xdc, 0x62, 0x7d, 0x57, 0x95,
	0x58, 0x9f, 0x0c, 0xa8, 0xa6, 0x1c, 0x58, 0x1c, 0x77, 0xa0, 0x3c, 0x32, 0xe3, 0xf2, 0x85, 0x33,
	0x96, 0xd3, 0x53, 0x64, 0x6e, 0x66, 0x1a, 0x9d, 0x5f, 0x0c, 0xb8, 0xde, 0x0a, 0x0e, 0x45, 0x48,
	0x90, 0x9f, 0x48, 0x3d, 0x80, 0x62, 0x9b, 0xb3, 0x13, 0xda, 0xd1, 0x7b, 0x7a, 0x94, 0xf9, 0x65,
	0x4a, 0x18, 0x9a, 0x12, 0xb8, 0x37, 0xe3, 0x6a, 0x0a, 0x73, 0x17, 0xf2, 0x18, 0x09, 0xa4, 0x17,
	0xe6, 0x4c, 0x41, 0xb5, 0x83, 0x04, 0xda, 0x9b, 0x71, 0x25, 0x7c, 0xbb, 0x04, 0x85, 0x1e, 0xa3,
	0x9c, 0x59, 0x9f, 0x0d, 0xb8, 0x96, 0x16, 0xbc, 0x9f, 0x12, 0x9c, 0xd5, 0xe5, 0x12, 0xb9, 0xcd,
	0x11, 0xb9, 0x0f, 0x27, 0x26, 0xba, 0x58, 0xec, 0x77, 0x03, 0xcc, 0x71, 0x77, 0xfe, 0x81, 0xff,
	0x06, 0xeb, 0x3d, 0xcc, 0x8f, 0x8d, 0xf5, 0xa7, 0x22, 0x6e, 0x6d, 0x44, 0x5f, 0x98, 0x51, 0xef,
	0x93, 0xd8, 0x1b, 0xd9, 0xb1, 0x7f, 0x0e, 0xd5, 0xab, 0xa3, 0xb7, 0xcb, 0x6f, 0x4a, 0x7a, 0xff,
	0xc7, 0x45, 0x79, 0xff, 0xf8, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x30, 0x98, 0x3c, 0x65,
	0x0b, 0x00, 0x00,
}
