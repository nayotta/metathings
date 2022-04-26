// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: sync_device_firmware_descriptor.proto

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

type SyncDeviceFirmwareDescriptorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device *OpDevice `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *SyncDeviceFirmwareDescriptorRequest) Reset() {
	*x = SyncDeviceFirmwareDescriptorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_device_firmware_descriptor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncDeviceFirmwareDescriptorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncDeviceFirmwareDescriptorRequest) ProtoMessage() {}

func (x *SyncDeviceFirmwareDescriptorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sync_device_firmware_descriptor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncDeviceFirmwareDescriptorRequest.ProtoReflect.Descriptor instead.
func (*SyncDeviceFirmwareDescriptorRequest) Descriptor() ([]byte, []int) {
	return file_sync_device_firmware_descriptor_proto_rawDescGZIP(), []int{0}
}

func (x *SyncDeviceFirmwareDescriptorRequest) GetDevice() *OpDevice {
	if x != nil {
		return x.Device
	}
	return nil
}

var File_sync_device_firmware_descriptor_proto protoreflect.FileDescriptor

var file_sync_device_firmware_descriptor_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x66, 0x69,
	0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x64, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x23, 0x53, 0x79, 0x6e, 0x63, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x06, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x69, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x64, 0x2e, 0x4f, 0x70, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x79, 0x6f, 0x74, 0x74,
	0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_sync_device_firmware_descriptor_proto_rawDescOnce sync.Once
	file_sync_device_firmware_descriptor_proto_rawDescData = file_sync_device_firmware_descriptor_proto_rawDesc
)

func file_sync_device_firmware_descriptor_proto_rawDescGZIP() []byte {
	file_sync_device_firmware_descriptor_proto_rawDescOnce.Do(func() {
		file_sync_device_firmware_descriptor_proto_rawDescData = protoimpl.X.CompressGZIP(file_sync_device_firmware_descriptor_proto_rawDescData)
	})
	return file_sync_device_firmware_descriptor_proto_rawDescData
}

var file_sync_device_firmware_descriptor_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sync_device_firmware_descriptor_proto_goTypes = []interface{}{
	(*SyncDeviceFirmwareDescriptorRequest)(nil), // 0: ai.metathings.service.deviced.SyncDeviceFirmwareDescriptorRequest
	(*OpDevice)(nil), // 1: ai.metathings.service.deviced.OpDevice
}
var file_sync_device_firmware_descriptor_proto_depIdxs = []int32{
	1, // 0: ai.metathings.service.deviced.SyncDeviceFirmwareDescriptorRequest.device:type_name -> ai.metathings.service.deviced.OpDevice
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sync_device_firmware_descriptor_proto_init() }
func file_sync_device_firmware_descriptor_proto_init() {
	if File_sync_device_firmware_descriptor_proto != nil {
		return
	}
	file_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sync_device_firmware_descriptor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncDeviceFirmwareDescriptorRequest); i {
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
			RawDescriptor: file_sync_device_firmware_descriptor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sync_device_firmware_descriptor_proto_goTypes,
		DependencyIndexes: file_sync_device_firmware_descriptor_proto_depIdxs,
		MessageInfos:      file_sync_device_firmware_descriptor_proto_msgTypes,
	}.Build()
	File_sync_device_firmware_descriptor_proto = out.File
	file_sync_device_firmware_descriptor_proto_rawDesc = nil
	file_sync_device_firmware_descriptor_proto_goTypes = nil
	file_sync_device_firmware_descriptor_proto_depIdxs = nil
}
