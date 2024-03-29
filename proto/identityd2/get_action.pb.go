// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: get_action.proto

package identityd2

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type GetActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action *OpAction `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *GetActionRequest) Reset() {
	*x = GetActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActionRequest) ProtoMessage() {}

func (x *GetActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_get_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActionRequest.ProtoReflect.Descriptor instead.
func (*GetActionRequest) Descriptor() ([]byte, []int) {
	return file_get_action_proto_rawDescGZIP(), []int{0}
}

func (x *GetActionRequest) GetAction() *OpAction {
	if x != nil {
		return x.Action
	}
	return nil
}

type GetActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action *Action `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *GetActionResponse) Reset() {
	*x = GetActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_action_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActionResponse) ProtoMessage() {}

func (x *GetActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_get_action_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActionResponse.ProtoReflect.Descriptor instead.
func (*GetActionResponse) Descriptor() ([]byte, []int) {
	return file_get_action_proto_rawDescGZIP(), []int{1}
}

func (x *GetActionResponse) GetAction() *Action {
	if x != nil {
		return x.Action
	}
	return nil
}

var File_get_action_proto protoreflect.FileDescriptor

var file_get_action_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x20, 0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x64, 0x32, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x56, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x64, 0x32, 0x2e, 0x4f, 0x70, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x55, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x64,
	0x32, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e,
	0x61, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x64, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_get_action_proto_rawDescOnce sync.Once
	file_get_action_proto_rawDescData = file_get_action_proto_rawDesc
)

func file_get_action_proto_rawDescGZIP() []byte {
	file_get_action_proto_rawDescOnce.Do(func() {
		file_get_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_get_action_proto_rawDescData)
	})
	return file_get_action_proto_rawDescData
}

var file_get_action_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_get_action_proto_goTypes = []interface{}{
	(*GetActionRequest)(nil),  // 0: ai.metathings.service.identityd2.GetActionRequest
	(*GetActionResponse)(nil), // 1: ai.metathings.service.identityd2.GetActionResponse
	(*OpAction)(nil),          // 2: ai.metathings.service.identityd2.OpAction
	(*Action)(nil),            // 3: ai.metathings.service.identityd2.Action
}
var file_get_action_proto_depIdxs = []int32{
	2, // 0: ai.metathings.service.identityd2.GetActionRequest.action:type_name -> ai.metathings.service.identityd2.OpAction
	3, // 1: ai.metathings.service.identityd2.GetActionResponse.action:type_name -> ai.metathings.service.identityd2.Action
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_get_action_proto_init() }
func file_get_action_proto_init() {
	if File_get_action_proto != nil {
		return
	}
	file_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_get_action_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetActionRequest); i {
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
		file_get_action_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetActionResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_get_action_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_get_action_proto_goTypes,
		DependencyIndexes: file_get_action_proto_depIdxs,
		MessageInfos:      file_get_action_proto_msgTypes,
	}.Build()
	File_get_action_proto = out.File
	file_get_action_proto_rawDesc = nil
	file_get_action_proto_goTypes = nil
	file_get_action_proto_depIdxs = nil
}
