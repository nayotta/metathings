// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stream.proto

package servo

import fmt "fmt"
import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *StreamPingRequest) Validate() error {
	if nil == this.Timestamp {
		return go_proto_validators.FieldError("Timestamp", fmt.Errorf("message must exist"))
	}
	if this.Timestamp != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Timestamp); err != nil {
			return go_proto_validators.FieldError("Timestamp", err)
		}
	}
	return nil
}
func (this *StreamPingResponse) Validate() error {
	if this.Timestamp != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Timestamp); err != nil {
			return go_proto_validators.FieldError("Timestamp", err)
		}
	}
	return nil
}
func (this *StreamSetAngleRequest) Validate() error {
	if nil == this.Name {
		return go_proto_validators.FieldError("Name", fmt.Errorf("message must exist"))
	}
	if this.Name != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return go_proto_validators.FieldError("Name", err)
		}
	}
	if nil == this.Angle {
		return go_proto_validators.FieldError("Angle", fmt.Errorf("message must exist"))
	}
	if this.Angle != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Angle); err != nil {
			return go_proto_validators.FieldError("Angle", err)
		}
	}
	return nil
}
func (this *StreamRequest) Validate() error {
	if nil == this.Session {
		return go_proto_validators.FieldError("Session", fmt.Errorf("message must exist"))
	}
	if this.Session != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Session); err != nil {
			return go_proto_validators.FieldError("Session", err)
		}
	}
	if oneOfNester, ok := this.GetPayload().(*StreamRequest_Ping); ok {
		if oneOfNester.Ping != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.Ping); err != nil {
				return go_proto_validators.FieldError("Ping", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*StreamRequest_SetAngle); ok {
		if oneOfNester.SetAngle != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.SetAngle); err != nil {
				return go_proto_validators.FieldError("SetAngle", err)
			}
		}
	}
	return nil
}
func (this *StreamRequests) Validate() error {
	for _, item := range this.Requests {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Requests", err)
			}
		}
	}
	return nil
}
func (this *StreamResponse) Validate() error {
	if oneOfNester, ok := this.GetPayload().(*StreamResponse_Ping); ok {
		if oneOfNester.Ping != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.Ping); err != nil {
				return go_proto_validators.FieldError("Ping", err)
			}
		}
	}
	return nil
}
