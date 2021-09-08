// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: stream_call.proto

package component

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type StreamCallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//	*StreamCallRequest_Config
	//	*StreamCallRequest_Data
	Request isStreamCallRequest_Request `protobuf_oneof:"request"`
}

func (x *StreamCallRequest) Reset() {
	*x = StreamCallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_call_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamCallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamCallRequest) ProtoMessage() {}

func (x *StreamCallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stream_call_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamCallRequest.ProtoReflect.Descriptor instead.
func (*StreamCallRequest) Descriptor() ([]byte, []int) {
	return file_stream_call_proto_rawDescGZIP(), []int{0}
}

func (m *StreamCallRequest) GetRequest() isStreamCallRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *StreamCallRequest) GetConfig() *StreamCallConfigRequest {
	if x, ok := x.GetRequest().(*StreamCallRequest_Config); ok {
		return x.Config
	}
	return nil
}

func (x *StreamCallRequest) GetData() *StreamCallDataRequest {
	if x, ok := x.GetRequest().(*StreamCallRequest_Data); ok {
		return x.Data
	}
	return nil
}

type isStreamCallRequest_Request interface {
	isStreamCallRequest_Request()
}

type StreamCallRequest_Config struct {
	Config *StreamCallConfigRequest `protobuf:"bytes,1,opt,name=config,proto3,oneof"`
}

type StreamCallRequest_Data struct {
	Data *StreamCallDataRequest `protobuf:"bytes,21,opt,name=data,proto3,oneof"`
}

func (*StreamCallRequest_Config) isStreamCallRequest_Request() {}

func (*StreamCallRequest_Data) isStreamCallRequest_Request() {}

type StreamCallConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *StreamCallConfigRequest) Reset() {
	*x = StreamCallConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_call_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamCallConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamCallConfigRequest) ProtoMessage() {}

func (x *StreamCallConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stream_call_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamCallConfigRequest.ProtoReflect.Descriptor instead.
func (*StreamCallConfigRequest) Descriptor() ([]byte, []int) {
	return file_stream_call_proto_rawDescGZIP(), []int{1}
}

func (x *StreamCallConfigRequest) GetMethod() *wrapperspb.StringValue {
	if x != nil {
		return x.Method
	}
	return nil
}

type StreamCallDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *anypb.Any `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *StreamCallDataRequest) Reset() {
	*x = StreamCallDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_call_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamCallDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamCallDataRequest) ProtoMessage() {}

func (x *StreamCallDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stream_call_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamCallDataRequest.ProtoReflect.Descriptor instead.
func (*StreamCallDataRequest) Descriptor() ([]byte, []int) {
	return file_stream_call_proto_rawDescGZIP(), []int{2}
}

func (x *StreamCallDataRequest) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

type StreamCallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*StreamCallResponse_Config
	//	*StreamCallResponse_Data
	Response isStreamCallResponse_Response `protobuf_oneof:"response"`
}

func (x *StreamCallResponse) Reset() {
	*x = StreamCallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_call_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamCallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamCallResponse) ProtoMessage() {}

func (x *StreamCallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stream_call_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamCallResponse.ProtoReflect.Descriptor instead.
func (*StreamCallResponse) Descriptor() ([]byte, []int) {
	return file_stream_call_proto_rawDescGZIP(), []int{3}
}

func (m *StreamCallResponse) GetResponse() isStreamCallResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *StreamCallResponse) GetConfig() *StreamCallConfigResponse {
	if x, ok := x.GetResponse().(*StreamCallResponse_Config); ok {
		return x.Config
	}
	return nil
}

func (x *StreamCallResponse) GetData() *StreamCallDataResponse {
	if x, ok := x.GetResponse().(*StreamCallResponse_Data); ok {
		return x.Data
	}
	return nil
}

type isStreamCallResponse_Response interface {
	isStreamCallResponse_Response()
}

type StreamCallResponse_Config struct {
	Config *StreamCallConfigResponse `protobuf:"bytes,1,opt,name=config,proto3,oneof"`
}

type StreamCallResponse_Data struct {
	Data *StreamCallDataResponse `protobuf:"bytes,21,opt,name=data,proto3,oneof"`
}

func (*StreamCallResponse_Config) isStreamCallResponse_Response() {}

func (*StreamCallResponse_Data) isStreamCallResponse_Response() {}

type StreamCallConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *StreamCallConfigResponse) Reset() {
	*x = StreamCallConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_call_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamCallConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamCallConfigResponse) ProtoMessage() {}

func (x *StreamCallConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stream_call_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamCallConfigResponse.ProtoReflect.Descriptor instead.
func (*StreamCallConfigResponse) Descriptor() ([]byte, []int) {
	return file_stream_call_proto_rawDescGZIP(), []int{4}
}

func (x *StreamCallConfigResponse) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type StreamCallDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *anypb.Any `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *StreamCallDataResponse) Reset() {
	*x = StreamCallDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_call_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamCallDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamCallDataResponse) ProtoMessage() {}

func (x *StreamCallDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stream_call_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamCallDataResponse.ProtoReflect.Descriptor instead.
func (*StreamCallDataResponse) Descriptor() ([]byte, []int) {
	return file_stream_call_proto_rawDescGZIP(), []int{5}
}

func (x *StreamCallDataResponse) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_stream_call_proto protoreflect.FileDescriptor

var file_stream_call_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x61,
	0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x00, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x44, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x61, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42,
	0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4f, 0x0a, 0x17, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x43, 0x0a, 0x15, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x61, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0xb4, 0x01, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x45, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x43, 0x61, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x0a, 0x18, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x44, 0x0a, 0x16, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x61, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6e, 0x61, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stream_call_proto_rawDescOnce sync.Once
	file_stream_call_proto_rawDescData = file_stream_call_proto_rawDesc
)

func file_stream_call_proto_rawDescGZIP() []byte {
	file_stream_call_proto_rawDescOnce.Do(func() {
		file_stream_call_proto_rawDescData = protoimpl.X.CompressGZIP(file_stream_call_proto_rawDescData)
	})
	return file_stream_call_proto_rawDescData
}

var file_stream_call_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_stream_call_proto_goTypes = []interface{}{
	(*StreamCallRequest)(nil),        // 0: ai.metathings.component.StreamCallRequest
	(*StreamCallConfigRequest)(nil),  // 1: ai.metathings.component.StreamCallConfigRequest
	(*StreamCallDataRequest)(nil),    // 2: ai.metathings.component.StreamCallDataRequest
	(*StreamCallResponse)(nil),       // 3: ai.metathings.component.StreamCallResponse
	(*StreamCallConfigResponse)(nil), // 4: ai.metathings.component.StreamCallConfigResponse
	(*StreamCallDataResponse)(nil),   // 5: ai.metathings.component.StreamCallDataResponse
	(*wrapperspb.StringValue)(nil),   // 6: google.protobuf.StringValue
	(*anypb.Any)(nil),                // 7: google.protobuf.Any
}
var file_stream_call_proto_depIdxs = []int32{
	1, // 0: ai.metathings.component.StreamCallRequest.config:type_name -> ai.metathings.component.StreamCallConfigRequest
	2, // 1: ai.metathings.component.StreamCallRequest.data:type_name -> ai.metathings.component.StreamCallDataRequest
	6, // 2: ai.metathings.component.StreamCallConfigRequest.method:type_name -> google.protobuf.StringValue
	7, // 3: ai.metathings.component.StreamCallDataRequest.value:type_name -> google.protobuf.Any
	4, // 4: ai.metathings.component.StreamCallResponse.config:type_name -> ai.metathings.component.StreamCallConfigResponse
	5, // 5: ai.metathings.component.StreamCallResponse.data:type_name -> ai.metathings.component.StreamCallDataResponse
	7, // 6: ai.metathings.component.StreamCallDataResponse.value:type_name -> google.protobuf.Any
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_stream_call_proto_init() }
func file_stream_call_proto_init() {
	if File_stream_call_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stream_call_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamCallRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stream_call_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamCallConfigRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stream_call_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamCallDataRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stream_call_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamCallResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stream_call_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamCallConfigResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stream_call_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamCallDataResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_stream_call_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*StreamCallRequest_Config)(nil),
		(*StreamCallRequest_Data)(nil),
	}
	file_stream_call_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*StreamCallResponse_Config)(nil),
		(*StreamCallResponse_Data)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stream_call_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stream_call_proto_goTypes,
		DependencyIndexes: file_stream_call_proto_depIdxs,
		MessageInfos:      file_stream_call_proto_msgTypes,
	}.Build()
	File_stream_call_proto = out.File
	file_stream_call_proto_rawDesc = nil
	file_stream_call_proto_goTypes = nil
	file_stream_call_proto_depIdxs = nil
}
