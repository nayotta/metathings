// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: list_objects.proto

package device

import (
	proto "github.com/golang/protobuf/proto"
	deviced "github.com/nayotta/metathings/proto/deviced"
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

type ListObjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object    *deviced.OpObject      `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	Recursive *wrapperspb.BoolValue  `protobuf:"bytes,2,opt,name=recursive,proto3" json:"recursive,omitempty"`
	Depth     *wrapperspb.Int32Value `protobuf:"bytes,3,opt,name=depth,proto3" json:"depth,omitempty"`
}

func (x *ListObjectsRequest) Reset() {
	*x = ListObjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_objects_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListObjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListObjectsRequest) ProtoMessage() {}

func (x *ListObjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_list_objects_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListObjectsRequest.ProtoReflect.Descriptor instead.
func (*ListObjectsRequest) Descriptor() ([]byte, []int) {
	return file_list_objects_proto_rawDescGZIP(), []int{0}
}

func (x *ListObjectsRequest) GetObject() *deviced.OpObject {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *ListObjectsRequest) GetRecursive() *wrapperspb.BoolValue {
	if x != nil {
		return x.Recursive
	}
	return nil
}

func (x *ListObjectsRequest) GetDepth() *wrapperspb.Int32Value {
	if x != nil {
		return x.Depth
	}
	return nil
}

type ListObjectsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Objects []*deviced.Object `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *ListObjectsResponse) Reset() {
	*x = ListObjectsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_objects_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListObjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListObjectsResponse) ProtoMessage() {}

func (x *ListObjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_list_objects_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListObjectsResponse.ProtoReflect.Descriptor instead.
func (*ListObjectsResponse) Descriptor() ([]byte, []int) {
	return file_list_objects_proto_rawDescGZIP(), []int{1}
}

func (x *ListObjectsResponse) GetObjects() []*deviced.Object {
	if x != nil {
		return x.Objects
	}
	return nil
}

var File_list_objects_proto protoreflect.FileDescriptor

var file_list_objects_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69,
	0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e,
	0x61, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x64, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3f, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x64, 0x2e, 0x4f, 0x70, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x12, 0x31, 0x0a,
	0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68,
	0x22, 0x56, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x69, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x64, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_list_objects_proto_rawDescOnce sync.Once
	file_list_objects_proto_rawDescData = file_list_objects_proto_rawDesc
)

func file_list_objects_proto_rawDescGZIP() []byte {
	file_list_objects_proto_rawDescOnce.Do(func() {
		file_list_objects_proto_rawDescData = protoimpl.X.CompressGZIP(file_list_objects_proto_rawDescData)
	})
	return file_list_objects_proto_rawDescData
}

var file_list_objects_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_list_objects_proto_goTypes = []interface{}{
	(*ListObjectsRequest)(nil),    // 0: ai.metathings.service.device.ListObjectsRequest
	(*ListObjectsResponse)(nil),   // 1: ai.metathings.service.device.ListObjectsResponse
	(*deviced.OpObject)(nil),      // 2: ai.metathings.service.deviced.OpObject
	(*wrapperspb.BoolValue)(nil),  // 3: google.protobuf.BoolValue
	(*wrapperspb.Int32Value)(nil), // 4: google.protobuf.Int32Value
	(*deviced.Object)(nil),        // 5: ai.metathings.service.deviced.Object
}
var file_list_objects_proto_depIdxs = []int32{
	2, // 0: ai.metathings.service.device.ListObjectsRequest.object:type_name -> ai.metathings.service.deviced.OpObject
	3, // 1: ai.metathings.service.device.ListObjectsRequest.recursive:type_name -> google.protobuf.BoolValue
	4, // 2: ai.metathings.service.device.ListObjectsRequest.depth:type_name -> google.protobuf.Int32Value
	5, // 3: ai.metathings.service.device.ListObjectsResponse.objects:type_name -> ai.metathings.service.deviced.Object
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_list_objects_proto_init() }
func file_list_objects_proto_init() {
	if File_list_objects_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_list_objects_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListObjectsRequest); i {
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
		file_list_objects_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListObjectsResponse); i {
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
			RawDescriptor: file_list_objects_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_list_objects_proto_goTypes,
		DependencyIndexes: file_list_objects_proto_depIdxs,
		MessageInfos:      file_list_objects_proto_msgTypes,
	}.Build()
	File_list_objects_proto = out.File
	file_list_objects_proto_rawDesc = nil
	file_list_objects_proto_goTypes = nil
	file_list_objects_proto_depIdxs = nil
}
