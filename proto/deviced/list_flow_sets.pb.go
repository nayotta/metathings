// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: list_flow_sets.proto

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

type ListFlowSetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlowSet *OpFlowSet `protobuf:"bytes,1,opt,name=flow_set,json=flowSet,proto3" json:"flow_set,omitempty"`
}

func (x *ListFlowSetsRequest) Reset() {
	*x = ListFlowSetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_flow_sets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFlowSetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFlowSetsRequest) ProtoMessage() {}

func (x *ListFlowSetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_list_flow_sets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFlowSetsRequest.ProtoReflect.Descriptor instead.
func (*ListFlowSetsRequest) Descriptor() ([]byte, []int) {
	return file_list_flow_sets_proto_rawDescGZIP(), []int{0}
}

func (x *ListFlowSetsRequest) GetFlowSet() *OpFlowSet {
	if x != nil {
		return x.FlowSet
	}
	return nil
}

type ListFlowSetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlowSets []*FlowSet `protobuf:"bytes,1,rep,name=flow_sets,json=flowSets,proto3" json:"flow_sets,omitempty"`
}

func (x *ListFlowSetsResponse) Reset() {
	*x = ListFlowSetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_flow_sets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFlowSetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFlowSetsResponse) ProtoMessage() {}

func (x *ListFlowSetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_list_flow_sets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFlowSetsResponse.ProtoReflect.Descriptor instead.
func (*ListFlowSetsResponse) Descriptor() ([]byte, []int) {
	return file_list_flow_sets_proto_rawDescGZIP(), []int{1}
}

func (x *ListFlowSetsResponse) GetFlowSets() []*FlowSet {
	if x != nil {
		return x.FlowSets
	}
	return nil
}

var File_list_flow_sets_proto protoreflect.FileDescriptor

var file_list_flow_sets_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x64, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x53, 0x65,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x08, 0x66, 0x6c, 0x6f,
	0x77, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x69,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x64, 0x2e, 0x4f, 0x70, 0x46, 0x6c,
	0x6f, 0x77, 0x53, 0x65, 0x74, 0x52, 0x07, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x74, 0x22, 0x5b,
	0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x09, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73,
	0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x69, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x64, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x53, 0x65,
	0x74, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x74, 0x73, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x79, 0x6f, 0x74, 0x74,
	0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_list_flow_sets_proto_rawDescOnce sync.Once
	file_list_flow_sets_proto_rawDescData = file_list_flow_sets_proto_rawDesc
)

func file_list_flow_sets_proto_rawDescGZIP() []byte {
	file_list_flow_sets_proto_rawDescOnce.Do(func() {
		file_list_flow_sets_proto_rawDescData = protoimpl.X.CompressGZIP(file_list_flow_sets_proto_rawDescData)
	})
	return file_list_flow_sets_proto_rawDescData
}

var file_list_flow_sets_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_list_flow_sets_proto_goTypes = []interface{}{
	(*ListFlowSetsRequest)(nil),  // 0: ai.metathings.service.deviced.ListFlowSetsRequest
	(*ListFlowSetsResponse)(nil), // 1: ai.metathings.service.deviced.ListFlowSetsResponse
	(*OpFlowSet)(nil),            // 2: ai.metathings.service.deviced.OpFlowSet
	(*FlowSet)(nil),              // 3: ai.metathings.service.deviced.FlowSet
}
var file_list_flow_sets_proto_depIdxs = []int32{
	2, // 0: ai.metathings.service.deviced.ListFlowSetsRequest.flow_set:type_name -> ai.metathings.service.deviced.OpFlowSet
	3, // 1: ai.metathings.service.deviced.ListFlowSetsResponse.flow_sets:type_name -> ai.metathings.service.deviced.FlowSet
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_list_flow_sets_proto_init() }
func file_list_flow_sets_proto_init() {
	if File_list_flow_sets_proto != nil {
		return
	}
	file_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_list_flow_sets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFlowSetsRequest); i {
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
		file_list_flow_sets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFlowSetsResponse); i {
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
			RawDescriptor: file_list_flow_sets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_list_flow_sets_proto_goTypes,
		DependencyIndexes: file_list_flow_sets_proto_depIdxs,
		MessageInfos:      file_list_flow_sets_proto_msgTypes,
	}.Build()
	File_list_flow_sets_proto = out.File
	file_list_flow_sets_proto_rawDesc = nil
	file_list_flow_sets_proto_goTypes = nil
	file_list_flow_sets_proto_depIdxs = nil
}
