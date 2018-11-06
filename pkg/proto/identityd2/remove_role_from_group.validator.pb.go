// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: remove_role_from_group.proto

package identityd2

import fmt "fmt"
import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *RemoveRoleFromGroupRequest) Validate() error {
	if nil == this.Group {
		return github_com_mwitkow_go_proto_validators.FieldError("Group", fmt.Errorf("message must exist"))
	}
	if this.Group != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Group); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Group", err)
		}
	}
	if nil == this.Role {
		return github_com_mwitkow_go_proto_validators.FieldError("Role", fmt.Errorf("message must exist"))
	}
	if this.Role != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Role); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Role", err)
		}
	}
	return nil
}
