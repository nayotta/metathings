// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: add_devices_to_firmware_hub.proto

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

func (this *AddDevicesToFirmwareHubRequest) Validate() error {
	if nil == this.FirmwareHub {
		return github_com_mwitkow_go_proto_validators.FieldError("FirmwareHub", fmt.Errorf("message must exist"))
	}
	if this.FirmwareHub != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.FirmwareHub); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("FirmwareHub", err)
		}
	}
	for _, item := range this.Devices {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Devices", err)
			}
		}
	}
	return nil
}
