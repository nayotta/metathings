// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: delete_flow_set.proto

package deviced

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

type DeleteFlowSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlowSet *OpFlowSet `protobuf:"bytes,1,opt,name=flow_set,json=flowSet,proto3" json:"flow_set,omitempty"`
}

func (x *DeleteFlowSetRequest) Reset() {
	*x = DeleteFlowSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delete_flow_set_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFlowSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFlowSetRequest) ProtoMessage() {}

func (x *DeleteFlowSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_delete_flow_set_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFlowSetRequest.ProtoReflect.Descriptor instead.
func (*DeleteFlowSetRequest) Descriptor() ([]byte, []int) {
	return file_delete_flow_set_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteFlowSetRequest) GetFlowSet() *OpFlowSet {
	if x != nil {
		return x.FlowSet
	}
	return nil
}

var File_delete_flow_set_proto protoreflect.FileDescriptor

var file_delete_flow_set_proto_rawDesc = []byte{
	0x0a, 0x15, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x64, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6c, 0x6f,
	0x77, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x08, 0x66,
	0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x64, 0x2e, 0x4f, 0x70,
	0x46, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x74, 0x52, 0x07, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x74,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e,
	0x61, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x64, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_delete_flow_set_proto_rawDescOnce sync.Once
	file_delete_flow_set_proto_rawDescData = file_delete_flow_set_proto_rawDesc
)

func file_delete_flow_set_proto_rawDescGZIP() []byte {
	file_delete_flow_set_proto_rawDescOnce.Do(func() {
		file_delete_flow_set_proto_rawDescData = protoimpl.X.CompressGZIP(file_delete_flow_set_proto_rawDescData)
	})
	return file_delete_flow_set_proto_rawDescData
}

var file_delete_flow_set_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_delete_flow_set_proto_goTypes = []interface{}{
	(*DeleteFlowSetRequest)(nil), // 0: ai.metathings.service.deviced.DeleteFlowSetRequest
	(*OpFlowSet)(nil),            // 1: ai.metathings.service.deviced.OpFlowSet
}
var file_delete_flow_set_proto_depIdxs = []int32{
	1, // 0: ai.metathings.service.deviced.DeleteFlowSetRequest.flow_set:type_name -> ai.metathings.service.deviced.OpFlowSet
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_delete_flow_set_proto_init() }
func file_delete_flow_set_proto_init() {
	if File_delete_flow_set_proto != nil {
		return
	}
	file_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_delete_flow_set_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFlowSetRequest); i {
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
			RawDescriptor: file_delete_flow_set_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_delete_flow_set_proto_goTypes,
		DependencyIndexes: file_delete_flow_set_proto_depIdxs,
		MessageInfos:      file_delete_flow_set_proto_msgTypes,
	}.Build()
	File_delete_flow_set_proto = out.File
	file_delete_flow_set_proto_rawDesc = nil
	file_delete_flow_set_proto_goTypes = nil
	file_delete_flow_set_proto_depIdxs = nil
}
