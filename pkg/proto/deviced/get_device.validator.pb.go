// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_device.proto

package deviced

import fmt "fmt"
import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GetDeviceRequest) Validate() error {
	if nil == this.Id {
		return go_proto_validators.FieldError("Id", fmt.Errorf("message must exist"))
	}
	if this.Id != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return go_proto_validators.FieldError("Id", err)
		}
	}
	return nil
}
func (this *GetDeviceResponse) Validate() error {
	if this.Device != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Device); err != nil {
			return go_proto_validators.FieldError("Device", err)
		}
	}
	return nil
}
