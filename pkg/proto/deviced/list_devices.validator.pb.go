// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_devices.proto

package deviced

import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"
import _ "github.com/nayotta/metathings/pkg/proto/constant/state"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ListDevicesRequest) Validate() error {
	if this.Name != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return go_proto_validators.FieldError("Name", err)
		}
	}
	if this.ProjectId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.ProjectId); err != nil {
			return go_proto_validators.FieldError("ProjectId", err)
		}
	}
	for _, item := range this.Owners {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Owners", err)
			}
		}
	}
	for _, item := range this.Groups {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Groups", err)
			}
		}
	}
	return nil
}
func (this *ListDevicesResponse) Validate() error {
	for _, item := range this.Devices {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Devices", err)
			}
		}
	}
	return nil
}
