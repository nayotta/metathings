// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: list_groups_for_object.proto

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

type ListGroupsForObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object *OpEntity `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *ListGroupsForObjectRequest) Reset() {
	*x = ListGroupsForObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_groups_for_object_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGroupsForObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupsForObjectRequest) ProtoMessage() {}

func (x *ListGroupsForObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_list_groups_for_object_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupsForObjectRequest.ProtoReflect.Descriptor instead.
func (*ListGroupsForObjectRequest) Descriptor() ([]byte, []int) {
	return file_list_groups_for_object_proto_rawDescGZIP(), []int{0}
}

func (x *ListGroupsForObjectRequest) GetObject() *OpEntity {
	if x != nil {
		return x.Object
	}
	return nil
}

type ListGroupsForObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups []*Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *ListGroupsForObjectResponse) Reset() {
	*x = ListGroupsForObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_groups_for_object_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGroupsForObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupsForObjectResponse) ProtoMessage() {}

func (x *ListGroupsForObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_list_groups_for_object_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupsForObjectResponse.ProtoReflect.Descriptor instead.
func (*ListGroupsForObjectResponse) Descriptor() ([]byte, []int) {
	return file_list_groups_for_object_proto_rawDescGZIP(), []int{1}
}

func (x *ListGroupsForObjectResponse) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

var File_list_groups_for_object_proto protoreflect.FileDescriptor

var file_list_groups_for_object_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x5f, 0x66, 0x6f,
	0x72, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20,
	0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x64, 0x32,
	0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a,
	0x1a, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x46, 0x6f, 0x72, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x69,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x64, 0x32, 0x2e, 0x4f,
	0x70, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22,
	0x5e, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x46, 0x6f, 0x72,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f,
	0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x64,
	0x32, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x42,
	0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61,
	0x79, 0x6f, 0x74, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x64,
	0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_list_groups_for_object_proto_rawDescOnce sync.Once
	file_list_groups_for_object_proto_rawDescData = file_list_groups_for_object_proto_rawDesc
)

func file_list_groups_for_object_proto_rawDescGZIP() []byte {
	file_list_groups_for_object_proto_rawDescOnce.Do(func() {
		file_list_groups_for_object_proto_rawDescData = protoimpl.X.CompressGZIP(file_list_groups_for_object_proto_rawDescData)
	})
	return file_list_groups_for_object_proto_rawDescData
}

var file_list_groups_for_object_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_list_groups_for_object_proto_goTypes = []interface{}{
	(*ListGroupsForObjectRequest)(nil),  // 0: ai.metathings.service.identityd2.ListGroupsForObjectRequest
	(*ListGroupsForObjectResponse)(nil), // 1: ai.metathings.service.identityd2.ListGroupsForObjectResponse
	(*OpEntity)(nil),                    // 2: ai.metathings.service.identityd2.OpEntity
	(*Group)(nil),                       // 3: ai.metathings.service.identityd2.Group
}
var file_list_groups_for_object_proto_depIdxs = []int32{
	2, // 0: ai.metathings.service.identityd2.ListGroupsForObjectRequest.object:type_name -> ai.metathings.service.identityd2.OpEntity
	3, // 1: ai.metathings.service.identityd2.ListGroupsForObjectResponse.groups:type_name -> ai.metathings.service.identityd2.Group
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_list_groups_for_object_proto_init() }
func file_list_groups_for_object_proto_init() {
	if File_list_groups_for_object_proto != nil {
		return
	}
	file_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_list_groups_for_object_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGroupsForObjectRequest); i {
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
		file_list_groups_for_object_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGroupsForObjectResponse); i {
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
			RawDescriptor: file_list_groups_for_object_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_list_groups_for_object_proto_goTypes,
		DependencyIndexes: file_list_groups_for_object_proto_depIdxs,
		MessageInfos:      file_list_groups_for_object_proto_msgTypes,
	}.Build()
	File_list_groups_for_object_proto = out.File
	file_list_groups_for_object_proto_rawDesc = nil
	file_list_groups_for_object_proto_goTypes = nil
	file_list_groups_for_object_proto_depIdxs = nil
}
