// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_groups_for_subject.proto

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

func (this *ListGroupsForSubjectRequest) Validate() error {
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
func (this *ListGroupsForSubjectResponse) Validate() error {
	for _, item := range this.Groups {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Groups", err)
			}
		}
	}
	return nil
}