// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: unary_call.proto

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

func (this *UnaryCallRequest) Validate() error {
	if nil == this.Device {
		return go_proto_validators.FieldError("Device", fmt.Errorf("message must exist"))
	}
	if this.Device != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Device); err != nil {
			return go_proto_validators.FieldError("Device", err)
		}
	}
	if nil == this.Value {
		return go_proto_validators.FieldError("Value", fmt.Errorf("message must exist"))
	}
	if this.Value != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
func (this *UnaryCallResponse) Validate() error {
	if this.Device != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Device); err != nil {
			return go_proto_validators.FieldError("Device", err)
		}
	}
	if this.Value != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
