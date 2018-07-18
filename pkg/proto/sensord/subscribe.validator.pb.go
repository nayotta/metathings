// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: subscribe.proto

package sensord

import fmt "fmt"
import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"
import _ "github.com/nayotta/metathings/pkg/proto/sensor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *SubscribeByIdRequest) Validate() error {
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
func (this *SubscribeRequest) Validate() error {
	if nil == this.Session {
		return go_proto_validators.FieldError("Session", fmt.Errorf("message must exist"))
	}
	if this.Session != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Session); err != nil {
			return go_proto_validators.FieldError("Session", err)
		}
	}
	if oneOfNester, ok := this.GetPayload().(*SubscribeRequest_Id); ok {
		if oneOfNester.Id != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.Id); err != nil {
				return go_proto_validators.FieldError("Id", err)
			}
		}
	}
	return nil
}
func (this *SubscribeRequests) Validate() error {
	for _, item := range this.Requests {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Requests", err)
			}
		}
	}
	return nil
}
func (this *SubscribeResponse) Validate() error {
	if this.Data != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *SubscribeResponses) Validate() error {
	for _, item := range this.Responses {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Responses", err)
			}
		}
	}
	return nil
}
