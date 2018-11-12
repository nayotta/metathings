// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: application_credential.proto

package identityd

import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ApplicationCredential) Validate() error {
	for _, item := range this.Roles {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Roles", err)
			}
		}
	}
	if this.ExpiresAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ExpiresAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ExpiresAt", err)
		}
	}
	return nil
}
func (this *ApplicationCredential__Role) Validate() error {
	return nil
}
