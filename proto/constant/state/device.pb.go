// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: device.proto

package state

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

type DeviceState int32

const (
	DeviceState_DEVICE_STATE_UNKNOWN DeviceState = 0
	DeviceState_DEVICE_STATE_ONLINE  DeviceState = 1
	DeviceState_DEVICE_STATE_OFFLINE DeviceState = 2
)

// Enum value maps for DeviceState.
var (
	DeviceState_name = map[int32]string{
		0: "DEVICE_STATE_UNKNOWN",
		1: "DEVICE_STATE_ONLINE",
		2: "DEVICE_STATE_OFFLINE",
	}
	DeviceState_value = map[string]int32{
		"DEVICE_STATE_UNKNOWN": 0,
		"DEVICE_STATE_ONLINE":  1,
		"DEVICE_STATE_OFFLINE": 2,
	}
)

func (x DeviceState) Enum() *DeviceState {
	p := new(DeviceState)
	*p = x
	return p
}

func (x DeviceState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceState) Descriptor() protoreflect.EnumDescriptor {
	return file_device_proto_enumTypes[0].Descriptor()
}

func (DeviceState) Type() protoreflect.EnumType {
	return &file_device_proto_enumTypes[0]
}

func (x DeviceState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceState.Descriptor instead.
func (DeviceState) EnumDescriptor() ([]byte, []int) {
	return file_device_proto_rawDescGZIP(), []int{0}
}

var File_device_proto protoreflect.FileDescriptor

var file_device_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c,
	0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2a, 0x5a, 0x0a, 0x0b,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x44,
	0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x18,
	0x0a, 0x14, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f,
	0x46, 0x46, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x02, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_device_proto_rawDescOnce sync.Once
	file_device_proto_rawDescData = file_device_proto_rawDesc
)

func file_device_proto_rawDescGZIP() []byte {
	file_device_proto_rawDescOnce.Do(func() {
		file_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_device_proto_rawDescData)
	})
	return file_device_proto_rawDescData
}

var file_device_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_device_proto_goTypes = []interface{}{
	(DeviceState)(0), // 0: ai.metathings.constant.state.DeviceState
}
var file_device_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_device_proto_init() }
func file_device_proto_init() {
	if File_device_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_device_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_device_proto_goTypes,
		DependencyIndexes: file_device_proto_depIdxs,
		EnumInfos:         file_device_proto_enumTypes,
	}.Build()
	File_device_proto = out.File
	file_device_proto_rawDesc = nil
	file_device_proto_goTypes = nil
	file_device_proto_depIdxs = nil
}
