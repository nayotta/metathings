// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sync_device_firmware_descriptor.proto

package deviced

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *SyncDeviceFirmwareDescriptorRequest) Validate() error {
	if nil == this.Device {
		return github_com_mwitkow_go_proto_validators.FieldError("Device", fmt.Errorf("message must exist"))
	}
	if this.Device != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Device); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Device", err)
		}
	}
	return nil
}
