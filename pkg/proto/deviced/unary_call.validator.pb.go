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
	if this.DeviceId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.DeviceId); err != nil {
			return go_proto_validators.FieldError("DeviceId", err)
		}
	}
	if nil == this.Payload {
		return go_proto_validators.FieldError("Payload", fmt.Errorf("message must exist"))
	}
	if this.Payload != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Payload); err != nil {
			return go_proto_validators.FieldError("Payload", err)
		}
	}
	return nil
}
func (this *UnaryCallResponse) Validate() error {
	if this.Payload != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Payload); err != nil {
			return go_proto_validators.FieldError("Payload", err)
		}
	}
	return nil
}