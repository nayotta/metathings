// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_users_in_group.proto

package identityd

import fmt "fmt"
import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ListUsersInGroupRequest) Validate() error {
	if nil == this.GroupId {
		return github_com_mwitkow_go_proto_validators.FieldError("GroupId", fmt.Errorf("message must exist"))
	}
	if this.GroupId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.GroupId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("GroupId", err)
		}
	}
	return nil
}
func (this *ListUsersInGroupResponse) Validate() error {
	for _, item := range this.Users {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Users", err)
			}
		}
	}
	return nil
}
