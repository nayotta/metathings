// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: pull_frame_from_flow.proto

package deviced

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type PullFrameFromFlowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to Request:
	//	*PullFrameFromFlowRequest_Config_
	Request isPullFrameFromFlowRequest_Request `protobuf_oneof:"request"`
}

func (x *PullFrameFromFlowRequest) Reset() {
	*x = PullFrameFromFlowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pull_frame_from_flow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullFrameFromFlowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullFrameFromFlowRequest) ProtoMessage() {}

func (x *PullFrameFromFlowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pull_frame_from_flow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullFrameFromFlowRequest.ProtoReflect.Descriptor instead.
func (*PullFrameFromFlowRequest) Descriptor() ([]byte, []int) {
	return file_pull_frame_from_flow_proto_rawDescGZIP(), []int{0}
}

func (x *PullFrameFromFlowRequest) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (m *PullFrameFromFlowRequest) GetRequest() isPullFrameFromFlowRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *PullFrameFromFlowRequest) GetConfig() *PullFrameFromFlowRequest_Config {
	if x, ok := x.GetRequest().(*PullFrameFromFlowRequest_Config_); ok {
		return x.Config
	}
	return nil
}

type isPullFrameFromFlowRequest_Request interface {
	isPullFrameFromFlowRequest_Request()
}

type PullFrameFromFlowRequest_Config_ struct {
	Config *PullFrameFromFlowRequest_Config `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

func (*PullFrameFromFlowRequest_Config_) isPullFrameFromFlowRequest_Request() {}

type PullFrameFromFlowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stream alive check when id is ffffffffffffffffffffffffffffffff,
	// drop it in receive side.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to Response:
	//	*PullFrameFromFlowResponse_Ack_
	//	*PullFrameFromFlowResponse_Pack_
	Response isPullFrameFromFlowResponse_Response `protobuf_oneof:"response"`
}

func (x *PullFrameFromFlowResponse) Reset() {
	*x = PullFrameFromFlowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pull_frame_from_flow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullFrameFromFlowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullFrameFromFlowResponse) ProtoMessage() {}

func (x *PullFrameFromFlowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pull_frame_from_flow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullFrameFromFlowResponse.ProtoReflect.Descriptor instead.
func (*PullFrameFromFlowResponse) Descriptor() ([]byte, []int) {
	return file_pull_frame_from_flow_proto_rawDescGZIP(), []int{1}
}

func (x *PullFrameFromFlowResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (m *PullFrameFromFlowResponse) GetResponse() isPullFrameFromFlowResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *PullFrameFromFlowResponse) GetAck() *PullFrameFromFlowResponse_Ack {
	if x, ok := x.GetResponse().(*PullFrameFromFlowResponse_Ack_); ok {
		return x.Ack
	}
	return nil
}

func (x *PullFrameFromFlowResponse) GetPack() *PullFrameFromFlowResponse_Pack {
	if x, ok := x.GetResponse().(*PullFrameFromFlowResponse_Pack_); ok {
		return x.Pack
	}
	return nil
}

type isPullFrameFromFlowResponse_Response interface {
	isPullFrameFromFlowResponse_Response()
}

type PullFrameFromFlowResponse_Ack_ struct {
	Ack *PullFrameFromFlowResponse_Ack `protobuf:"bytes,2,opt,name=ack,proto3,oneof"`
}

type PullFrameFromFlowResponse_Pack_ struct {
	Pack *PullFrameFromFlowResponse_Pack `protobuf:"bytes,3,opt,name=pack,proto3,oneof"`
}

func (*PullFrameFromFlowResponse_Ack_) isPullFrameFromFlowResponse_Response() {}

func (*PullFrameFromFlowResponse_Pack_) isPullFrameFromFlowResponse_Response() {}

type PullFrameFromFlowRequest_Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device    *OpDevice             `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	ConfigAck *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=config_ack,json=configAck,proto3" json:"config_ack,omitempty"`
}

func (x *PullFrameFromFlowRequest_Config) Reset() {
	*x = PullFrameFromFlowRequest_Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pull_frame_from_flow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullFrameFromFlowRequest_Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullFrameFromFlowRequest_Config) ProtoMessage() {}

func (x *PullFrameFromFlowRequest_Config) ProtoReflect() protoreflect.Message {
	mi := &file_pull_frame_from_flow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullFrameFromFlowRequest_Config.ProtoReflect.Descriptor instead.
func (*PullFrameFromFlowRequest_Config) Descriptor() ([]byte, []int) {
	return file_pull_frame_from_flow_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PullFrameFromFlowRequest_Config) GetDevice() *OpDevice {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *PullFrameFromFlowRequest_Config) GetConfigAck() *wrapperspb.BoolValue {
	if x != nil {
		return x.ConfigAck
	}
	return nil
}

type PullFrameFromFlowResponse_Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PullFrameFromFlowResponse_Ack) Reset() {
	*x = PullFrameFromFlowResponse_Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pull_frame_from_flow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullFrameFromFlowResponse_Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullFrameFromFlowResponse_Ack) ProtoMessage() {}

func (x *PullFrameFromFlowResponse_Ack) ProtoReflect() protoreflect.Message {
	mi := &file_pull_frame_from_flow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullFrameFromFlowResponse_Ack.ProtoReflect.Descriptor instead.
func (*PullFrameFromFlowResponse_Ack) Descriptor() ([]byte, []int) {
	return file_pull_frame_from_flow_proto_rawDescGZIP(), []int{1, 0}
}

type PullFrameFromFlowResponse_Pack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flow   *Flow    `protobuf:"bytes,1,opt,name=flow,proto3" json:"flow,omitempty"`
	Frames []*Frame `protobuf:"bytes,2,rep,name=frames,proto3" json:"frames,omitempty"`
}

func (x *PullFrameFromFlowResponse_Pack) Reset() {
	*x = PullFrameFromFlowResponse_Pack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pull_frame_from_flow_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullFrameFromFlowResponse_Pack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullFrameFromFlowResponse_Pack) ProtoMessage() {}

func (x *PullFrameFromFlowResponse_Pack) ProtoReflect() protoreflect.Message {
	mi := &file_pull_frame_from_flow_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullFrameFromFlowResponse_Pack.ProtoReflect.Descriptor instead.
func (*PullFrameFromFlowResponse_Pack) Descriptor() ([]byte, []int) {
	return file_pull_frame_from_flow_proto_rawDescGZIP(), []int{1, 1}
}

func (x *PullFrameFromFlowResponse_Pack) GetFlow() *Flow {
	if x != nil {
		return x.Flow
	}
	return nil
}

func (x *PullFrameFromFlowResponse_Pack) GetFrames() []*Frame {
	if x != nil {
		return x.Frames
	}
	return nil
}

var File_pull_frame_from_flow_proto protoreflect.FileDescriptor

var file_pull_frame_from_flow_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x61, 0x69,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x64, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x02, 0x0a, 0x18, 0x50, 0x75, 0x6c,
	0x6c, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x58, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69,
	0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x64, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x84, 0x01,
	0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3f, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x69, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x64, 0x2e, 0x4f, 0x70, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x5f, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x41, 0x63, 0x6b, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0xe4, 0x02, 0x0a, 0x19, 0x50, 0x75, 0x6c, 0x6c, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x50, 0x0a,
	0x03, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x61, 0x69, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x64, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x12,
	0x53, 0x0a, 0x04, 0x70, 0x61, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e,
	0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x64, 0x2e, 0x50, 0x75,
	0x6c, 0x6c, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x46, 0x6c, 0x6f, 0x77, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x04,
	0x70, 0x61, 0x63, 0x6b, 0x1a, 0x05, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x1a, 0x7d, 0x0a, 0x04, 0x50,
	0x61, 0x63, 0x6b, 0x12, 0x37, 0x0a, 0x04, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x64, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x04, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x3c, 0x0a, 0x06,
	0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61,
	0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x64, 0x2e, 0x46, 0x72, 0x61,
	0x6d, 0x65, 0x52, 0x06, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pull_frame_from_flow_proto_rawDescOnce sync.Once
	file_pull_frame_from_flow_proto_rawDescData = file_pull_frame_from_flow_proto_rawDesc
)

func file_pull_frame_from_flow_proto_rawDescGZIP() []byte {
	file_pull_frame_from_flow_proto_rawDescOnce.Do(func() {
		file_pull_frame_from_flow_proto_rawDescData = protoimpl.X.CompressGZIP(file_pull_frame_from_flow_proto_rawDescData)
	})
	return file_pull_frame_from_flow_proto_rawDescData
}

var file_pull_frame_from_flow_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pull_frame_from_flow_proto_goTypes = []interface{}{
	(*PullFrameFromFlowRequest)(nil),        // 0: ai.metathings.service.deviced.PullFrameFromFlowRequest
	(*PullFrameFromFlowResponse)(nil),       // 1: ai.metathings.service.deviced.PullFrameFromFlowResponse
	(*PullFrameFromFlowRequest_Config)(nil), // 2: ai.metathings.service.deviced.PullFrameFromFlowRequest.Config
	(*PullFrameFromFlowResponse_Ack)(nil),   // 3: ai.metathings.service.deviced.PullFrameFromFlowResponse.Ack
	(*PullFrameFromFlowResponse_Pack)(nil),  // 4: ai.metathings.service.deviced.PullFrameFromFlowResponse.Pack
	(*wrapperspb.StringValue)(nil),          // 5: google.protobuf.StringValue
	(*OpDevice)(nil),                        // 6: ai.metathings.service.deviced.OpDevice
	(*wrapperspb.BoolValue)(nil),            // 7: google.protobuf.BoolValue
	(*Flow)(nil),                            // 8: ai.metathings.service.deviced.Flow
	(*Frame)(nil),                           // 9: ai.metathings.service.deviced.Frame
}
var file_pull_frame_from_flow_proto_depIdxs = []int32{
	5, // 0: ai.metathings.service.deviced.PullFrameFromFlowRequest.id:type_name -> google.protobuf.StringValue
	2, // 1: ai.metathings.service.deviced.PullFrameFromFlowRequest.config:type_name -> ai.metathings.service.deviced.PullFrameFromFlowRequest.Config
	3, // 2: ai.metathings.service.deviced.PullFrameFromFlowResponse.ack:type_name -> ai.metathings.service.deviced.PullFrameFromFlowResponse.Ack
	4, // 3: ai.metathings.service.deviced.PullFrameFromFlowResponse.pack:type_name -> ai.metathings.service.deviced.PullFrameFromFlowResponse.Pack
	6, // 4: ai.metathings.service.deviced.PullFrameFromFlowRequest.Config.device:type_name -> ai.metathings.service.deviced.OpDevice
	7, // 5: ai.metathings.service.deviced.PullFrameFromFlowRequest.Config.config_ack:type_name -> google.protobuf.BoolValue
	8, // 6: ai.metathings.service.deviced.PullFrameFromFlowResponse.Pack.flow:type_name -> ai.metathings.service.deviced.Flow
	9, // 7: ai.metathings.service.deviced.PullFrameFromFlowResponse.Pack.frames:type_name -> ai.metathings.service.deviced.Frame
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_pull_frame_from_flow_proto_init() }
func file_pull_frame_from_flow_proto_init() {
	if File_pull_frame_from_flow_proto != nil {
		return
	}
	file_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pull_frame_from_flow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullFrameFromFlowRequest); i {
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
		file_pull_frame_from_flow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullFrameFromFlowResponse); i {
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
		file_pull_frame_from_flow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullFrameFromFlowRequest_Config); i {
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
		file_pull_frame_from_flow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullFrameFromFlowResponse_Ack); i {
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
		file_pull_frame_from_flow_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullFrameFromFlowResponse_Pack); i {
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
	file_pull_frame_from_flow_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PullFrameFromFlowRequest_Config_)(nil),
	}
	file_pull_frame_from_flow_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*PullFrameFromFlowResponse_Ack_)(nil),
		(*PullFrameFromFlowResponse_Pack_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pull_frame_from_flow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pull_frame_from_flow_proto_goTypes,
		DependencyIndexes: file_pull_frame_from_flow_proto_depIdxs,
		MessageInfos:      file_pull_frame_from_flow_proto_msgTypes,
	}.Build()
	File_pull_frame_from_flow_proto = out.File
	file_pull_frame_from_flow_proto_rawDesc = nil
	file_pull_frame_from_flow_proto_goTypes = nil
	file_pull_frame_from_flow_proto_depIdxs = nil
}
