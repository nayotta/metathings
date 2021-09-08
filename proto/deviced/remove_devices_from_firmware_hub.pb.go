// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: remove_devices_from_firmware_hub.proto

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

type RemoveDevicesFromFirmwareHubRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirmwareHub *OpFirmwareHub `protobuf:"bytes,1,opt,name=firmware_hub,json=firmwareHub,proto3" json:"firmware_hub,omitempty"`
	Devices     []*OpDevice    `protobuf:"bytes,2,rep,name=devices,proto3" json:"devices,omitempty"`
}

func (x *RemoveDevicesFromFirmwareHubRequest) Reset() {
	*x = RemoveDevicesFromFirmwareHubRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remove_devices_from_firmware_hub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveDevicesFromFirmwareHubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveDevicesFromFirmwareHubRequest) ProtoMessage() {}

func (x *RemoveDevicesFromFirmwareHubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remove_devices_from_firmware_hub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveDevicesFromFirmwareHubRequest.ProtoReflect.Descriptor instead.
func (*RemoveDevicesFromFirmwareHubRequest) Descriptor() ([]byte, []int) {
	return file_remove_devices_from_firmware_hub_proto_rawDescGZIP(), []int{0}
}

func (x *RemoveDevicesFromFirmwareHubRequest) GetFirmwareHub() *OpFirmwareHub {
	if x != nil {
		return x.FirmwareHub
	}
	return nil
}

func (x *RemoveDevicesFromFirmwareHubRequest) GetDevices() []*OpDevice {
	if x != nil {
		return x.Devices
	}
	return nil
}

var File_remove_devices_from_firmware_hub_proto protoreflect.FileDescriptor

var file_remove_devices_from_firmware_hub_proto_rawDesc = []byte{
	0x0a, 0x26, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x68,
	0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x64, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x23, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61,
	0x72, 0x65, 0x48, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4f, 0x0a, 0x0c,
	0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x68, 0x75, 0x62, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x64, 0x2e, 0x4f, 0x70, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x48, 0x75, 0x62,
	0x52, 0x0b, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x48, 0x75, 0x62, 0x12, 0x41, 0x0a,
	0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x61, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x64, 0x2e, 0x4f,
	0x70, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e,
	0x61, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x64, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_remove_devices_from_firmware_hub_proto_rawDescOnce sync.Once
	file_remove_devices_from_firmware_hub_proto_rawDescData = file_remove_devices_from_firmware_hub_proto_rawDesc
)

func file_remove_devices_from_firmware_hub_proto_rawDescGZIP() []byte {
	file_remove_devices_from_firmware_hub_proto_rawDescOnce.Do(func() {
		file_remove_devices_from_firmware_hub_proto_rawDescData = protoimpl.X.CompressGZIP(file_remove_devices_from_firmware_hub_proto_rawDescData)
	})
	return file_remove_devices_from_firmware_hub_proto_rawDescData
}

var file_remove_devices_from_firmware_hub_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_remove_devices_from_firmware_hub_proto_goTypes = []interface{}{
	(*RemoveDevicesFromFirmwareHubRequest)(nil), // 0: ai.metathings.service.deviced.RemoveDevicesFromFirmwareHubRequest
	(*OpFirmwareHub)(nil),                       // 1: ai.metathings.service.deviced.OpFirmwareHub
	(*OpDevice)(nil),                            // 2: ai.metathings.service.deviced.OpDevice
}
var file_remove_devices_from_firmware_hub_proto_depIdxs = []int32{
	1, // 0: ai.metathings.service.deviced.RemoveDevicesFromFirmwareHubRequest.firmware_hub:type_name -> ai.metathings.service.deviced.OpFirmwareHub
	2, // 1: ai.metathings.service.deviced.RemoveDevicesFromFirmwareHubRequest.devices:type_name -> ai.metathings.service.deviced.OpDevice
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_remove_devices_from_firmware_hub_proto_init() }
func file_remove_devices_from_firmware_hub_proto_init() {
	if File_remove_devices_from_firmware_hub_proto != nil {
		return
	}
	file_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_remove_devices_from_firmware_hub_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveDevicesFromFirmwareHubRequest); i {
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
			RawDescriptor: file_remove_devices_from_firmware_hub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_remove_devices_from_firmware_hub_proto_goTypes,
		DependencyIndexes: file_remove_devices_from_firmware_hub_proto_depIdxs,
		MessageInfos:      file_remove_devices_from_firmware_hub_proto_msgTypes,
	}.Build()
	File_remove_devices_from_firmware_hub_proto = out.File
	file_remove_devices_from_firmware_hub_proto_rawDesc = nil
	file_remove_devices_from_firmware_hub_proto_goTypes = nil
	file_remove_devices_from_firmware_hub_proto_depIdxs = nil
}
