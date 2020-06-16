// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: remove_configs_from_timer.proto

package evaluatord

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "github.com/nayotta/metathings/pkg/proto/deviced"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *RemoveConfigsFromTimerRequest) Validate() error {
	if nil == this.Timer {
		return github_com_mwitkow_go_proto_validators.FieldError("Timer", fmt.Errorf("message must exist"))
	}
	if this.Timer != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Timer); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Timer", err)
		}
	}
	for _, item := range this.Configs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Configs", err)
			}
		}
	}
	return nil
}
