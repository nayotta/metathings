// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: remove_subject_from_group.proto

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

func (this *RemoveSubjectFromGroupRequest) Validate() error {
	if nil == this.Group {
		return github_com_mwitkow_go_proto_validators.FieldError("Group", fmt.Errorf("message must exist"))
	}
	if this.Group != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Group); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Group", err)
		}
	}
	if nil == this.Subject {
		return github_com_mwitkow_go_proto_validators.FieldError("Subject", fmt.Errorf("message must exist"))
	}
	if this.Subject != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Subject); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Subject", err)
		}
	}
	return nil
}
