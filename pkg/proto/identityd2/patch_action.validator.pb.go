// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: patch_action.proto

package identityd2

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

func (this *PatchActionRequest) Validate() error {
	if nil == this.Action {
		return github_com_mwitkow_go_proto_validators.FieldError("Action", fmt.Errorf("message must exist"))
	}
	if this.Action != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Action); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Action", err)
		}
	}
	return nil
}
func (this *PatchActionResponse) Validate() error {
	if this.Action != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Action); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Action", err)
		}
	}
	return nil
}