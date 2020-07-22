// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_timer.proto

package evaluatord

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

func (this *CreateTimerRequest) Validate() error {
	if nil == this.Timer {
		return github_com_mwitkow_go_proto_validators.FieldError("Timer", fmt.Errorf("message must exist"))
	}
	if this.Timer != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Timer); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Timer", err)
		}
	}
	return nil
}
func (this *CreateTimerResponse) Validate() error {
	if this.Timer != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Timer); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Timer", err)
		}
	}
	return nil
}
