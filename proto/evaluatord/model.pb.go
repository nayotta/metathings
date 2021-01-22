// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

package evaluatord

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	state "github.com/nayotta/metathings/proto/constant/state"
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

type Resource struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 string   `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{0}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Resource) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type OpResource struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 *wrappers.StringValue `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OpResource) Reset()         { *m = OpResource{} }
func (m *OpResource) String() string { return proto.CompactTextString(m) }
func (*OpResource) ProtoMessage()    {}
func (*OpResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{1}
}

func (m *OpResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpResource.Unmarshal(m, b)
}
func (m *OpResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpResource.Marshal(b, m, deterministic)
}
func (m *OpResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpResource.Merge(m, src)
}
func (m *OpResource) XXX_Size() int {
	return xxx_messageInfo_OpResource.Size(m)
}
func (m *OpResource) XXX_DiscardUnknown() {
	xxx_messageInfo_OpResource.DiscardUnknown(m)
}

var xxx_messageInfo_OpResource proto.InternalMessageInfo

func (m *OpResource) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *OpResource) GetType() *wrappers.StringValue {
	if m != nil {
		return m.Type
	}
	return nil
}

type Evaluator struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Alias                string          `protobuf:"bytes,24,opt,name=alias,proto3" json:"alias,omitempty"`
	Description          string          `protobuf:"bytes,25,opt,name=description,proto3" json:"description,omitempty"`
	Sources              []*Resource     `protobuf:"bytes,8,rep,name=sources,proto3" json:"sources,omitempty"`
	Operator             *Operator       `protobuf:"bytes,10,opt,name=operator,proto3" json:"operator,omitempty"`
	Config               *_struct.Struct `protobuf:"bytes,11,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Evaluator) Reset()         { *m = Evaluator{} }
func (m *Evaluator) String() string { return proto.CompactTextString(m) }
func (*Evaluator) ProtoMessage()    {}
func (*Evaluator) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{2}
}

func (m *Evaluator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Evaluator.Unmarshal(m, b)
}
func (m *Evaluator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Evaluator.Marshal(b, m, deterministic)
}
func (m *Evaluator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Evaluator.Merge(m, src)
}
func (m *Evaluator) XXX_Size() int {
	return xxx_messageInfo_Evaluator.Size(m)
}
func (m *Evaluator) XXX_DiscardUnknown() {
	xxx_messageInfo_Evaluator.DiscardUnknown(m)
}

var xxx_messageInfo_Evaluator proto.InternalMessageInfo

func (m *Evaluator) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Evaluator) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *Evaluator) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Evaluator) GetSources() []*Resource {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *Evaluator) GetOperator() *Operator {
	if m != nil {
		return m.Operator
	}
	return nil
}

func (m *Evaluator) GetConfig() *_struct.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

type OpEvaluator struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Alias                *wrappers.StringValue `protobuf:"bytes,24,opt,name=alias,proto3" json:"alias,omitempty"`
	Description          *wrappers.StringValue `protobuf:"bytes,25,opt,name=description,proto3" json:"description,omitempty"`
	Sources              []*OpResource         `protobuf:"bytes,8,rep,name=sources,proto3" json:"sources,omitempty"`
	Operator             *OpOperator           `protobuf:"bytes,10,opt,name=operator,proto3" json:"operator,omitempty"`
	Config               *_struct.Struct       `protobuf:"bytes,11,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OpEvaluator) Reset()         { *m = OpEvaluator{} }
func (m *OpEvaluator) String() string { return proto.CompactTextString(m) }
func (*OpEvaluator) ProtoMessage()    {}
func (*OpEvaluator) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{3}
}

func (m *OpEvaluator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpEvaluator.Unmarshal(m, b)
}
func (m *OpEvaluator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpEvaluator.Marshal(b, m, deterministic)
}
func (m *OpEvaluator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpEvaluator.Merge(m, src)
}
func (m *OpEvaluator) XXX_Size() int {
	return xxx_messageInfo_OpEvaluator.Size(m)
}
func (m *OpEvaluator) XXX_DiscardUnknown() {
	xxx_messageInfo_OpEvaluator.DiscardUnknown(m)
}

var xxx_messageInfo_OpEvaluator proto.InternalMessageInfo

func (m *OpEvaluator) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *OpEvaluator) GetAlias() *wrappers.StringValue {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *OpEvaluator) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *OpEvaluator) GetSources() []*OpResource {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *OpEvaluator) GetOperator() *OpOperator {
	if m != nil {
		return m.Operator
	}
	return nil
}

func (m *OpEvaluator) GetConfig() *_struct.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

type Operator struct {
	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Alias       string `protobuf:"bytes,24,opt,name=alias,proto3" json:"alias,omitempty"`
	Description string `protobuf:"bytes,25,opt,name=description,proto3" json:"description,omitempty"`
	Driver      string `protobuf:"bytes,8,opt,name=driver,proto3" json:"driver,omitempty"`
	// Types that are valid to be assigned to Descriptor_:
	//	*Operator_Lua
	Descriptor_          isOperator_Descriptor_ `protobuf_oneof:"descriptor"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Operator) Reset()         { *m = Operator{} }
func (m *Operator) String() string { return proto.CompactTextString(m) }
func (*Operator) ProtoMessage()    {}
func (*Operator) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{4}
}

func (m *Operator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operator.Unmarshal(m, b)
}
func (m *Operator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operator.Marshal(b, m, deterministic)
}
func (m *Operator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operator.Merge(m, src)
}
func (m *Operator) XXX_Size() int {
	return xxx_messageInfo_Operator.Size(m)
}
func (m *Operator) XXX_DiscardUnknown() {
	xxx_messageInfo_Operator.DiscardUnknown(m)
}

var xxx_messageInfo_Operator proto.InternalMessageInfo

func (m *Operator) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Operator) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *Operator) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Operator) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

type isOperator_Descriptor_ interface {
	isOperator_Descriptor_()
}

type Operator_Lua struct {
	Lua *LuaDescriptor `protobuf:"bytes,32,opt,name=lua,proto3,oneof"`
}

func (*Operator_Lua) isOperator_Descriptor_() {}

func (m *Operator) GetDescriptor_() isOperator_Descriptor_ {
	if m != nil {
		return m.Descriptor_
	}
	return nil
}

func (m *Operator) GetLua() *LuaDescriptor {
	if x, ok := m.GetDescriptor_().(*Operator_Lua); ok {
		return x.Lua
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Operator) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Operator_Lua)(nil),
	}
}

type OpOperator struct {
	Id          *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Alias       *wrappers.StringValue `protobuf:"bytes,24,opt,name=alias,proto3" json:"alias,omitempty"`
	Description *wrappers.StringValue `protobuf:"bytes,25,opt,name=description,proto3" json:"description,omitempty"`
	Driver      *wrappers.StringValue `protobuf:"bytes,8,opt,name=driver,proto3" json:"driver,omitempty"`
	// Types that are valid to be assigned to Descriptor_:
	//	*OpOperator_Lua
	Descriptor_          isOpOperator_Descriptor_ `protobuf_oneof:"descriptor"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *OpOperator) Reset()         { *m = OpOperator{} }
func (m *OpOperator) String() string { return proto.CompactTextString(m) }
func (*OpOperator) ProtoMessage()    {}
func (*OpOperator) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{5}
}

func (m *OpOperator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpOperator.Unmarshal(m, b)
}
func (m *OpOperator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpOperator.Marshal(b, m, deterministic)
}
func (m *OpOperator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpOperator.Merge(m, src)
}
func (m *OpOperator) XXX_Size() int {
	return xxx_messageInfo_OpOperator.Size(m)
}
func (m *OpOperator) XXX_DiscardUnknown() {
	xxx_messageInfo_OpOperator.DiscardUnknown(m)
}

var xxx_messageInfo_OpOperator proto.InternalMessageInfo

func (m *OpOperator) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *OpOperator) GetAlias() *wrappers.StringValue {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *OpOperator) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *OpOperator) GetDriver() *wrappers.StringValue {
	if m != nil {
		return m.Driver
	}
	return nil
}

type isOpOperator_Descriptor_ interface {
	isOpOperator_Descriptor_()
}

type OpOperator_Lua struct {
	Lua *OpLuaDescriptor `protobuf:"bytes,32,opt,name=lua,proto3,oneof"`
}

func (*OpOperator_Lua) isOpOperator_Descriptor_() {}

func (m *OpOperator) GetDescriptor_() isOpOperator_Descriptor_ {
	if m != nil {
		return m.Descriptor_
	}
	return nil
}

func (m *OpOperator) GetLua() *OpLuaDescriptor {
	if x, ok := m.GetDescriptor_().(*OpOperator_Lua); ok {
		return x.Lua
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*OpOperator) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*OpOperator_Lua)(nil),
	}
}

type LuaDescriptor struct {
	Code                 string   `protobuf:"bytes,8,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LuaDescriptor) Reset()         { *m = LuaDescriptor{} }
func (m *LuaDescriptor) String() string { return proto.CompactTextString(m) }
func (*LuaDescriptor) ProtoMessage()    {}
func (*LuaDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{6}
}

func (m *LuaDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LuaDescriptor.Unmarshal(m, b)
}
func (m *LuaDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LuaDescriptor.Marshal(b, m, deterministic)
}
func (m *LuaDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LuaDescriptor.Merge(m, src)
}
func (m *LuaDescriptor) XXX_Size() int {
	return xxx_messageInfo_LuaDescriptor.Size(m)
}
func (m *LuaDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_LuaDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_LuaDescriptor proto.InternalMessageInfo

func (m *LuaDescriptor) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type OpLuaDescriptor struct {
	Code                 *wrappers.StringValue `protobuf:"bytes,8,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OpLuaDescriptor) Reset()         { *m = OpLuaDescriptor{} }
func (m *OpLuaDescriptor) String() string { return proto.CompactTextString(m) }
func (*OpLuaDescriptor) ProtoMessage()    {}
func (*OpLuaDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{7}
}

func (m *OpLuaDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpLuaDescriptor.Unmarshal(m, b)
}
func (m *OpLuaDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpLuaDescriptor.Marshal(b, m, deterministic)
}
func (m *OpLuaDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpLuaDescriptor.Merge(m, src)
}
func (m *OpLuaDescriptor) XXX_Size() int {
	return xxx_messageInfo_OpLuaDescriptor.Size(m)
}
func (m *OpLuaDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_OpLuaDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_OpLuaDescriptor proto.InternalMessageInfo

func (m *OpLuaDescriptor) GetCode() *wrappers.StringValue {
	if m != nil {
		return m.Code
	}
	return nil
}

type OpTaskState struct {
	At                   *timestamp.Timestamp `protobuf:"bytes,1,opt,name=at,proto3" json:"at,omitempty"`
	State                state.TaskState      `protobuf:"varint,2,opt,name=state,proto3,enum=ai.metathings.constant.state.TaskState" json:"state,omitempty"`
	Tags                 *_struct.Struct      `protobuf:"bytes,8,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OpTaskState) Reset()         { *m = OpTaskState{} }
func (m *OpTaskState) String() string { return proto.CompactTextString(m) }
func (*OpTaskState) ProtoMessage()    {}
func (*OpTaskState) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{8}
}

func (m *OpTaskState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpTaskState.Unmarshal(m, b)
}
func (m *OpTaskState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpTaskState.Marshal(b, m, deterministic)
}
func (m *OpTaskState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpTaskState.Merge(m, src)
}
func (m *OpTaskState) XXX_Size() int {
	return xxx_messageInfo_OpTaskState.Size(m)
}
func (m *OpTaskState) XXX_DiscardUnknown() {
	xxx_messageInfo_OpTaskState.DiscardUnknown(m)
}

var xxx_messageInfo_OpTaskState proto.InternalMessageInfo

func (m *OpTaskState) GetAt() *timestamp.Timestamp {
	if m != nil {
		return m.At
	}
	return nil
}

func (m *OpTaskState) GetState() state.TaskState {
	if m != nil {
		return m.State
	}
	return state.TaskState_TASK_STATE_UNKNOWN
}

func (m *OpTaskState) GetTags() *_struct.Struct {
	if m != nil {
		return m.Tags
	}
	return nil
}

type TaskState struct {
	At                   *timestamp.Timestamp `protobuf:"bytes,1,opt,name=at,proto3" json:"at,omitempty"`
	State                state.TaskState      `protobuf:"varint,2,opt,name=state,proto3,enum=ai.metathings.constant.state.TaskState" json:"state,omitempty"`
	Tags                 *_struct.Struct      `protobuf:"bytes,8,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TaskState) Reset()         { *m = TaskState{} }
func (m *TaskState) String() string { return proto.CompactTextString(m) }
func (*TaskState) ProtoMessage()    {}
func (*TaskState) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{9}
}

func (m *TaskState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskState.Unmarshal(m, b)
}
func (m *TaskState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskState.Marshal(b, m, deterministic)
}
func (m *TaskState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskState.Merge(m, src)
}
func (m *TaskState) XXX_Size() int {
	return xxx_messageInfo_TaskState.Size(m)
}
func (m *TaskState) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskState.DiscardUnknown(m)
}

var xxx_messageInfo_TaskState proto.InternalMessageInfo

func (m *TaskState) GetAt() *timestamp.Timestamp {
	if m != nil {
		return m.At
	}
	return nil
}

func (m *TaskState) GetState() state.TaskState {
	if m != nil {
		return m.State
	}
	return state.TaskState_TASK_STATE_UNKNOWN
}

func (m *TaskState) GetTags() *_struct.Struct {
	if m != nil {
		return m.Tags
	}
	return nil
}

type OpTask struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp  `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp  `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CurrentState         *OpTaskState          `protobuf:"bytes,4,opt,name=current_state,json=currentState,proto3" json:"current_state,omitempty"`
	Source               *OpResource           `protobuf:"bytes,8,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OpTask) Reset()         { *m = OpTask{} }
func (m *OpTask) String() string { return proto.CompactTextString(m) }
func (*OpTask) ProtoMessage()    {}
func (*OpTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{10}
}

func (m *OpTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpTask.Unmarshal(m, b)
}
func (m *OpTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpTask.Marshal(b, m, deterministic)
}
func (m *OpTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpTask.Merge(m, src)
}
func (m *OpTask) XXX_Size() int {
	return xxx_messageInfo_OpTask.Size(m)
}
func (m *OpTask) XXX_DiscardUnknown() {
	xxx_messageInfo_OpTask.DiscardUnknown(m)
}

var xxx_messageInfo_OpTask proto.InternalMessageInfo

func (m *OpTask) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *OpTask) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *OpTask) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *OpTask) GetCurrentState() *OpTaskState {
	if m != nil {
		return m.CurrentState
	}
	return nil
}

func (m *OpTask) GetSource() *OpResource {
	if m != nil {
		return m.Source
	}
	return nil
}

type Task struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CurrentState         *TaskState           `protobuf:"bytes,4,opt,name=current_state,json=currentState,proto3" json:"current_state,omitempty"`
	Source               *Resource            `protobuf:"bytes,8,opt,name=source,proto3" json:"source,omitempty"`
	States               []*TaskState         `protobuf:"bytes,9,rep,name=states,proto3" json:"states,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{11}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Task) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Task) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Task) GetCurrentState() *TaskState {
	if m != nil {
		return m.CurrentState
	}
	return nil
}

func (m *Task) GetSource() *Resource {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Task) GetStates() []*TaskState {
	if m != nil {
		return m.States
	}
	return nil
}

type OpTimer struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Alias                *wrappers.StringValue `protobuf:"bytes,24,opt,name=alias,proto3" json:"alias,omitempty"`
	Description          *wrappers.StringValue `protobuf:"bytes,25,opt,name=description,proto3" json:"description,omitempty"`
	Schedule             *wrappers.StringValue `protobuf:"bytes,8,opt,name=schedule,proto3" json:"schedule,omitempty"`
	Timezone             *wrappers.StringValue `protobuf:"bytes,9,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Enabled              *wrappers.BoolValue   `protobuf:"bytes,10,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Configs              []*deviced.OpConfig   `protobuf:"bytes,11,rep,name=configs,proto3" json:"configs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OpTimer) Reset()         { *m = OpTimer{} }
func (m *OpTimer) String() string { return proto.CompactTextString(m) }
func (*OpTimer) ProtoMessage()    {}
func (*OpTimer) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{12}
}

func (m *OpTimer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpTimer.Unmarshal(m, b)
}
func (m *OpTimer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpTimer.Marshal(b, m, deterministic)
}
func (m *OpTimer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpTimer.Merge(m, src)
}
func (m *OpTimer) XXX_Size() int {
	return xxx_messageInfo_OpTimer.Size(m)
}
func (m *OpTimer) XXX_DiscardUnknown() {
	xxx_messageInfo_OpTimer.DiscardUnknown(m)
}

var xxx_messageInfo_OpTimer proto.InternalMessageInfo

func (m *OpTimer) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *OpTimer) GetAlias() *wrappers.StringValue {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *OpTimer) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *OpTimer) GetSchedule() *wrappers.StringValue {
	if m != nil {
		return m.Schedule
	}
	return nil
}

func (m *OpTimer) GetTimezone() *wrappers.StringValue {
	if m != nil {
		return m.Timezone
	}
	return nil
}

func (m *OpTimer) GetEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.Enabled
	}
	return nil
}

func (m *OpTimer) GetConfigs() []*deviced.OpConfig {
	if m != nil {
		return m.Configs
	}
	return nil
}

type Timer struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Alias                string            `protobuf:"bytes,24,opt,name=alias,proto3" json:"alias,omitempty"`
	Description          string            `protobuf:"bytes,25,opt,name=description,proto3" json:"description,omitempty"`
	Schedule             string            `protobuf:"bytes,8,opt,name=schedule,proto3" json:"schedule,omitempty"`
	Timezone             string            `protobuf:"bytes,9,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Enabled              bool              `protobuf:"varint,10,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Configs              []*deviced.Config `protobuf:"bytes,11,rep,name=configs,proto3" json:"configs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Timer) Reset()         { *m = Timer{} }
func (m *Timer) String() string { return proto.CompactTextString(m) }
func (*Timer) ProtoMessage()    {}
func (*Timer) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{13}
}

func (m *Timer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timer.Unmarshal(m, b)
}
func (m *Timer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timer.Marshal(b, m, deterministic)
}
func (m *Timer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timer.Merge(m, src)
}
func (m *Timer) XXX_Size() int {
	return xxx_messageInfo_Timer.Size(m)
}
func (m *Timer) XXX_DiscardUnknown() {
	xxx_messageInfo_Timer.DiscardUnknown(m)
}

var xxx_messageInfo_Timer proto.InternalMessageInfo

func (m *Timer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Timer) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *Timer) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Timer) GetSchedule() string {
	if m != nil {
		return m.Schedule
	}
	return ""
}

func (m *Timer) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *Timer) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Timer) GetConfigs() []*deviced.Config {
	if m != nil {
		return m.Configs
	}
	return nil
}

func init() {
	proto.RegisterType((*Resource)(nil), "ai.metathings.service.evaluatord.Resource")
	proto.RegisterType((*OpResource)(nil), "ai.metathings.service.evaluatord.OpResource")
	proto.RegisterType((*Evaluator)(nil), "ai.metathings.service.evaluatord.Evaluator")
	proto.RegisterType((*OpEvaluator)(nil), "ai.metathings.service.evaluatord.OpEvaluator")
	proto.RegisterType((*Operator)(nil), "ai.metathings.service.evaluatord.Operator")
	proto.RegisterType((*OpOperator)(nil), "ai.metathings.service.evaluatord.OpOperator")
	proto.RegisterType((*LuaDescriptor)(nil), "ai.metathings.service.evaluatord.LuaDescriptor")
	proto.RegisterType((*OpLuaDescriptor)(nil), "ai.metathings.service.evaluatord.OpLuaDescriptor")
	proto.RegisterType((*OpTaskState)(nil), "ai.metathings.service.evaluatord.OpTaskState")
	proto.RegisterType((*TaskState)(nil), "ai.metathings.service.evaluatord.TaskState")
	proto.RegisterType((*OpTask)(nil), "ai.metathings.service.evaluatord.OpTask")
	proto.RegisterType((*Task)(nil), "ai.metathings.service.evaluatord.Task")
	proto.RegisterType((*OpTimer)(nil), "ai.metathings.service.evaluatord.OpTimer")
	proto.RegisterType((*Timer)(nil), "ai.metathings.service.evaluatord.Timer")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_4c16552f9fdb66d8) }

var fileDescriptor_4c16552f9fdb66d8 = []byte{
	// 841 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x96, 0x41, 0x4f, 0xe3, 0x46,
	0x14, 0xc7, 0x1b, 0x27, 0x24, 0xf1, 0x0b, 0x50, 0x69, 0x54, 0xb5, 0x6e, 0x84, 0xda, 0xc8, 0x55,
	0x05, 0x02, 0x6a, 0xb7, 0x29, 0x52, 0xdb, 0x03, 0xad, 0x48, 0x00, 0x71, 0xa8, 0x94, 0xca, 0xa0,
	0x1e, 0x7a, 0x41, 0x13, 0x7b, 0x08, 0x16, 0x8e, 0xc7, 0x9a, 0x19, 0x53, 0xd1, 0x8f, 0xd1, 0xd3,
	0xde, 0x57, 0xfb, 0x21, 0xf6, 0xb0, 0xd7, 0xfd, 0x20, 0x7b, 0xd9, 0xfb, 0x7e, 0x82, 0x95, 0xc7,
	0x63, 0x27, 0x71, 0x60, 0xed, 0x48, 0x48, 0xcb, 0xde, 0x12, 0xcf, 0xfb, 0x3f, 0xbf, 0xff, 0xef,
	0x8d, 0xdf, 0x0c, 0x74, 0xa6, 0xd4, 0x23, 0x81, 0x15, 0x31, 0x2a, 0x28, 0xea, 0x61, 0xdf, 0x9a,
	0x12, 0x81, 0xc5, 0xb5, 0x1f, 0x4e, 0xb8, 0xc5, 0x09, 0xbb, 0xf5, 0x5d, 0x62, 0x91, 0x5b, 0x1c,
	0xc4, 0x58, 0x50, 0xe6, 0x75, 0xbf, 0x99, 0x50, 0x3a, 0x09, 0x88, 0x2d, 0xe3, 0xc7, 0xf1, 0x95,
	0xfd, 0x2f, 0xc3, 0x51, 0x44, 0x18, 0x4f, 0x33, 0x74, 0xb7, 0x8a, 0xeb, 0x5c, 0xb0, 0xd8, 0x15,
	0x6a, 0xf5, 0xdb, 0xe2, 0xaa, 0xf0, 0xa7, 0x84, 0x0b, 0x3c, 0x8d, 0x54, 0xc0, 0xe1, 0xc4, 0x17,
	0xd7, 0xf1, 0xd8, 0x72, 0xe9, 0xd4, 0x0e, 0xf1, 0x1d, 0x15, 0x02, 0xdb, 0xb3, 0x82, 0x52, 0x9d,
	0xed, 0xd2, 0x90, 0x0b, 0x1c, 0x0a, 0x9b, 0x0b, 0x2c, 0x88, 0x2d, 0x30, 0xbf, 0x51, 0xf2, 0x5f,
	0xaa, 0xc8, 0x3d, 0x92, 0x98, 0xf2, 0xec, 0x39, 0xe3, 0xa6, 0x05, 0x6d, 0x87, 0x70, 0x1a, 0x33,
	0x97, 0xa0, 0x4d, 0xd0, 0x7c, 0xcf, 0xa8, 0xf5, 0x6a, 0x3b, 0xba, 0xa3, 0xf9, 0x1e, 0x42, 0xd0,
	0x10, 0x77, 0x11, 0x31, 0xda, 0xf2, 0x89, 0xfc, 0x6d, 0x06, 0x00, 0xa3, 0x28, 0x57, 0xec, 0xe7,
	0x8a, 0x4e, 0x7f, 0xcb, 0x4a, 0x3d, 0x5a, 0x99, 0x47, 0xeb, 0x5c, 0x30, 0x3f, 0x9c, 0xfc, 0x8d,
	0x83, 0x98, 0xc8, 0x7c, 0x3f, 0xce, 0xe5, 0x2b, 0x8b, 0x4f, 0xdf, 0xf6, 0xbf, 0x06, 0xfa, 0x49,
	0xd6, 0x83, 0xa5, 0xfa, 0xbe, 0x80, 0x35, 0x1c, 0xf8, 0x98, 0x1b, 0x86, 0x7c, 0x94, 0xfe, 0x41,
	0x3d, 0xe8, 0x78, 0x84, 0xbb, 0xcc, 0x8f, 0x84, 0x4f, 0x43, 0xe3, 0x6b, 0xb9, 0x36, 0xff, 0x08,
	0x1d, 0x43, 0x2b, 0xad, 0x9f, 0x1b, 0xed, 0x5e, 0x7d, 0xa7, 0xd3, 0xdf, 0xb5, 0xca, 0xda, 0x6f,
	0x65, 0x96, 0x9d, 0x4c, 0x8a, 0x4e, 0xa1, 0x4d, 0x23, 0xc2, 0x92, 0x65, 0x03, 0xa4, 0xa3, 0x0a,
	0x69, 0x46, 0x4a, 0xe1, 0xe4, 0x5a, 0x64, 0x43, 0xd3, 0xa5, 0xe1, 0x95, 0x3f, 0x31, 0x3a, 0x32,
	0xcb, 0x57, 0xf7, 0x71, 0x89, 0x5d, 0xe1, 0xa8, 0x30, 0xf3, 0x9d, 0x06, 0x9d, 0x51, 0x34, 0xc3,
	0xb2, 0x5a, 0x13, 0xfa, 0xf3, 0xd0, 0xca, 0x04, 0x0a, 0xe9, 0xef, 0xcb, 0x48, 0xcb, 0x94, 0x0b,
	0xc0, 0x4f, 0x8b, 0xc0, 0xf7, 0xab, 0x90, 0x5a, 0x46, 0x7e, 0xb6, 0x84, 0xbc, 0x52, 0xa2, 0xc7,
	0x80, 0xfe, 0xaa, 0x06, 0xed, 0x2c, 0xcf, 0xa3, 0x6d, 0xc4, 0x2f, 0xa1, 0xe9, 0x31, 0xff, 0x96,
	0x30, 0xf5, 0x89, 0xa9, 0x7f, 0x68, 0x08, 0xf5, 0x20, 0xc6, 0x46, 0x4f, 0x96, 0x66, 0x97, 0x5b,
	0xfc, 0x33, 0xc6, 0xc7, 0x2a, 0x2d, 0x65, 0x67, 0x9f, 0x39, 0x89, 0x7a, 0xb0, 0x0e, 0xe0, 0xe5,
	0x0f, 0xcd, 0x97, 0x5a, 0xf2, 0xe1, 0xe6, 0x0e, 0x9e, 0xfe, 0x9e, 0x39, 0x58, 0x60, 0x53, 0x26,
	0xcd, 0xc8, 0x9d, 0xcc, 0x93, 0xfb, 0xa9, 0xca, 0xe6, 0xa8, 0xc0, 0xee, 0x3b, 0xd8, 0x58, 0x88,
	0x4a, 0x06, 0xa3, 0x4b, 0xbd, 0x7c, 0x30, 0x26, 0xbf, 0xcd, 0x21, 0x7c, 0x5e, 0x48, 0x96, 0xcc,
	0xbb, 0x3c, 0xac, 0x74, 0xde, 0xc9, 0x24, 0x2f, 0x6a, 0xc9, 0xa7, 0x7d, 0x81, 0xf9, 0xcd, 0x79,
	0x32, 0xe1, 0xd1, 0x2e, 0x68, 0x58, 0xa8, 0x36, 0x75, 0x97, 0xf4, 0x17, 0xd9, 0x19, 0xe2, 0x68,
	0x58, 0xa0, 0x43, 0x58, 0x93, 0xc7, 0x82, 0xa1, 0xf5, 0x6a, 0x3b, 0x9b, 0xfd, 0xed, 0x82, 0xf9,
	0xec, 0xec, 0xb0, 0x64, 0x90, 0x95, 0xbf, 0xc3, 0x49, 0x55, 0x68, 0x0f, 0x1a, 0x02, 0x4f, 0xb8,
	0x2a, 0xf6, 0xc1, 0xef, 0x41, 0x06, 0x99, 0xcf, 0x6b, 0xa0, 0x3f, 0xfd, 0x2a, 0x5f, 0x6b, 0xd0,
	0x4c, 0x69, 0xae, 0xb8, 0xdf, 0x7f, 0x03, 0x70, 0x19, 0xc1, 0x82, 0x78, 0x97, 0x58, 0xc8, 0x4a,
	0x3f, 0x6c, 0x4c, 0x57, 0xd1, 0x47, 0x22, 0x91, 0xc6, 0x91, 0x97, 0x49, 0xeb, 0xe5, 0x52, 0x15,
	0x7d, 0x24, 0x90, 0x03, 0x1b, 0x6e, 0xcc, 0x18, 0x09, 0xc5, 0x65, 0x8a, 0xa8, 0x21, 0xd5, 0x3f,
	0x54, 0xd9, 0xc5, 0x33, 0x50, 0xeb, 0x2a, 0x47, 0xda, 0x9a, 0x63, 0x68, 0xa6, 0xc3, 0x53, 0x11,
	0x5b, 0x6d, 0xf0, 0x2a, 0xad, 0xf9, 0x56, 0x83, 0x86, 0xc4, 0x58, 0x1c, 0x7c, 0x1f, 0x07, 0xd4,
	0x5f, 0xf7, 0x83, 0xda, 0x2b, 0xf7, 0xf6, 0x10, 0xa6, 0x41, 0x01, 0xd3, 0x2a, 0x17, 0x02, 0xa5,
	0x44, 0x43, 0x68, 0xca, 0x6a, 0xb8, 0xa1, 0xcb, 0x33, 0x6e, 0xa5, 0x72, 0x94, 0xd4, 0x7c, 0x56,
	0x87, 0xd6, 0x28, 0x4a, 0x5c, 0x7f, 0x0a, 0x33, 0xfa, 0x57, 0x68, 0x73, 0xf7, 0x9a, 0x78, 0x71,
	0x50, 0x6d, 0xc8, 0xe5, 0xd1, 0x89, 0x32, 0xb9, 0x01, 0xff, 0x47, 0x43, 0x62, 0xe8, 0x55, 0x94,
	0x59, 0x34, 0x3a, 0x80, 0x16, 0x09, 0xf1, 0x38, 0x20, 0x9e, 0xba, 0x02, 0x2c, 0x6f, 0x9a, 0x01,
	0xa5, 0x41, 0x2a, 0xcb, 0x42, 0xd1, 0x11, 0xb4, 0xd2, 0x83, 0x9c, 0x1b, 0x1d, 0xd9, 0x9d, 0xed,
	0x07, 0xba, 0xa3, 0xee, 0xc8, 0xd6, 0x28, 0x1a, 0xca, 0x78, 0x27, 0xd3, 0x99, 0x6f, 0x6a, 0xb0,
	0x96, 0x36, 0xe6, 0xb1, 0x8e, 0xff, 0x6e, 0x01, 0x9f, 0x3e, 0x07, 0xa8, 0x5b, 0x00, 0xa4, 0xcf,
	0x21, 0x30, 0x16, 0x11, 0xb4, 0x67, 0x36, 0xff, 0x28, 0xda, 0xfc, 0xbe, 0xc4, 0x66, 0xc1, 0xe4,
	0x60, 0xfd, 0x1f, 0x98, 0xed, 0xcf, 0x71, 0x53, 0x22, 0xfd, 0xf9, 0x7d, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xca, 0x13, 0xdf, 0x3a, 0x2b, 0x0d, 0x00, 0x00,
}