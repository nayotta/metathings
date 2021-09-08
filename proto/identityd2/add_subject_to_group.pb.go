// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: add_subject_to_group.proto

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

type AddSubjectToGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group   *OpGroup  `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Subject *OpEntity `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *AddSubjectToGroupRequest) Reset() {
	*x = AddSubjectToGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_add_subject_to_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSubjectToGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSubjectToGroupRequest) ProtoMessage() {}

func (x *AddSubjectToGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_add_subject_to_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSubjectToGroupRequest.ProtoReflect.Descriptor instead.
func (*AddSubjectToGroupRequest) Descriptor() ([]byte, []int) {
	return file_add_subject_to_group_proto_rawDescGZIP(), []int{0}
}

func (x *AddSubjectToGroupRequest) GetGroup() *OpGroup {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *AddSubjectToGroupRequest) GetSubject() *OpEntity {
	if x != nil {
		return x.Subject
	}
	return nil
}

var File_add_subject_to_group_proto protoreflect.FileDescriptor

var file_add_subject_to_group_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x64, 0x64, 0x5f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x6f,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x61, 0x69,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x64, 0x32, 0x1a, 0x0b,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x18,
	0x41, 0x64, 0x64, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x64, 0x32, 0x2e, 0x4f, 0x70, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x44, 0x0a, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x69, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x64, 0x32, 0x2e, 0x4f, 0x70,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42,
	0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61,
	0x79, 0x6f, 0x74, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x64,
	0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_add_subject_to_group_proto_rawDescOnce sync.Once
	file_add_subject_to_group_proto_rawDescData = file_add_subject_to_group_proto_rawDesc
)

func file_add_subject_to_group_proto_rawDescGZIP() []byte {
	file_add_subject_to_group_proto_rawDescOnce.Do(func() {
		file_add_subject_to_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_add_subject_to_group_proto_rawDescData)
	})
	return file_add_subject_to_group_proto_rawDescData
}

var file_add_subject_to_group_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_add_subject_to_group_proto_goTypes = []interface{}{
	(*AddSubjectToGroupRequest)(nil), // 0: ai.metathings.service.identityd2.AddSubjectToGroupRequest
	(*OpGroup)(nil),                  // 1: ai.metathings.service.identityd2.OpGroup
	(*OpEntity)(nil),                 // 2: ai.metathings.service.identityd2.OpEntity
}
var file_add_subject_to_group_proto_depIdxs = []int32{
	1, // 0: ai.metathings.service.identityd2.AddSubjectToGroupRequest.group:type_name -> ai.metathings.service.identityd2.OpGroup
	2, // 1: ai.metathings.service.identityd2.AddSubjectToGroupRequest.subject:type_name -> ai.metathings.service.identityd2.OpEntity
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_add_subject_to_group_proto_init() }
func file_add_subject_to_group_proto_init() {
	if File_add_subject_to_group_proto != nil {
		return
	}
	file_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_add_subject_to_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSubjectToGroupRequest); i {
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
			RawDescriptor: file_add_subject_to_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_add_subject_to_group_proto_goTypes,
		DependencyIndexes: file_add_subject_to_group_proto_depIdxs,
		MessageInfos:      file_add_subject_to_group_proto_msgTypes,
	}.Build()
	File_add_subject_to_group_proto = out.File
	file_add_subject_to_group_proto_rawDesc = nil
	file_add_subject_to_group_proto_goTypes = nil
	file_add_subject_to_group_proto_depIdxs = nil
}
