// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_flow_set.proto

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

func (this *GetFlowSetRequest) Validate() error {
	if nil == this.FlowSet {
		return github_com_mwitkow_go_proto_validators.FieldError("FlowSet", fmt.Errorf("message must exist"))
	}
	if this.FlowSet != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.FlowSet); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("FlowSet", err)
		}
	}
	return nil
}
func (this *GetFlowSetResponse) Validate() error {
	if this.FlowSet != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.FlowSet); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("FlowSet", err)
		}
	}
	return nil
}
