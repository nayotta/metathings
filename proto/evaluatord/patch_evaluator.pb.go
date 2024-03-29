// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: patch_evaluator.proto

package evaluatord

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

type PatchEvaluatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Evaluator *OpEvaluator `protobuf:"bytes,1,opt,name=evaluator,proto3" json:"evaluator,omitempty"`
}

func (x *PatchEvaluatorRequest) Reset() {
	*x = PatchEvaluatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patch_evaluator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchEvaluatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchEvaluatorRequest) ProtoMessage() {}

func (x *PatchEvaluatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_patch_evaluator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchEvaluatorRequest.ProtoReflect.Descriptor instead.
func (*PatchEvaluatorRequest) Descriptor() ([]byte, []int) {
	return file_patch_evaluator_proto_rawDescGZIP(), []int{0}
}

func (x *PatchEvaluatorRequest) GetEvaluator() *OpEvaluator {
	if x != nil {
		return x.Evaluator
	}
	return nil
}

type PatchEvaluatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Evaluator *Evaluator `protobuf:"bytes,1,opt,name=evaluator,proto3" json:"evaluator,omitempty"`
}

func (x *PatchEvaluatorResponse) Reset() {
	*x = PatchEvaluatorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patch_evaluator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchEvaluatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchEvaluatorResponse) ProtoMessage() {}

func (x *PatchEvaluatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_patch_evaluator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchEvaluatorResponse.ProtoReflect.Descriptor instead.
func (*PatchEvaluatorResponse) Descriptor() ([]byte, []int) {
	return file_patch_evaluator_proto_rawDescGZIP(), []int{1}
}

func (x *PatchEvaluatorResponse) GetEvaluator() *Evaluator {
	if x != nil {
		return x.Evaluator
	}
	return nil
}

var File_patch_evaluator_proto protoreflect.FileDescriptor

var file_patch_evaluator_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x65,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x64, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x15, 0x50, 0x61, 0x74, 0x63, 0x68, 0x45,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x4b, 0x0a, 0x09, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x65, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x6f, 0x72, 0x64, 0x2e, 0x4f, 0x70, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x09, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x63, 0x0a, 0x16,
	0x50, 0x61, 0x74, 0x63, 0x68, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x09, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x69, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x64, 0x2e, 0x45, 0x76, 0x61,
	0x6c, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x6f,
	0x72, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6e, 0x61, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x6f, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_patch_evaluator_proto_rawDescOnce sync.Once
	file_patch_evaluator_proto_rawDescData = file_patch_evaluator_proto_rawDesc
)

func file_patch_evaluator_proto_rawDescGZIP() []byte {
	file_patch_evaluator_proto_rawDescOnce.Do(func() {
		file_patch_evaluator_proto_rawDescData = protoimpl.X.CompressGZIP(file_patch_evaluator_proto_rawDescData)
	})
	return file_patch_evaluator_proto_rawDescData
}

var file_patch_evaluator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_patch_evaluator_proto_goTypes = []interface{}{
	(*PatchEvaluatorRequest)(nil),  // 0: ai.metathings.service.evaluatord.PatchEvaluatorRequest
	(*PatchEvaluatorResponse)(nil), // 1: ai.metathings.service.evaluatord.PatchEvaluatorResponse
	(*OpEvaluator)(nil),            // 2: ai.metathings.service.evaluatord.OpEvaluator
	(*Evaluator)(nil),              // 3: ai.metathings.service.evaluatord.Evaluator
}
var file_patch_evaluator_proto_depIdxs = []int32{
	2, // 0: ai.metathings.service.evaluatord.PatchEvaluatorRequest.evaluator:type_name -> ai.metathings.service.evaluatord.OpEvaluator
	3, // 1: ai.metathings.service.evaluatord.PatchEvaluatorResponse.evaluator:type_name -> ai.metathings.service.evaluatord.Evaluator
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_patch_evaluator_proto_init() }
func file_patch_evaluator_proto_init() {
	if File_patch_evaluator_proto != nil {
		return
	}
	file_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_patch_evaluator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchEvaluatorRequest); i {
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
		file_patch_evaluator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchEvaluatorResponse); i {
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
			RawDescriptor: file_patch_evaluator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_patch_evaluator_proto_goTypes,
		DependencyIndexes: file_patch_evaluator_proto_depIdxs,
		MessageInfos:      file_patch_evaluator_proto_msgTypes,
	}.Build()
	File_patch_evaluator_proto = out.File
	file_patch_evaluator_proto_rawDesc = nil
	file_patch_evaluator_proto_goTypes = nil
	file_patch_evaluator_proto_depIdxs = nil
}
