// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete_application_credential.proto

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

func (this *DeleteApplicationCredentialRequest) Validate() error {
	if nil == this.UserId {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf("message must exist"))
	}
	if this.UserId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UserId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UserId", err)
		}
	}
	if nil == this.ApplicationCredentialId {
		return github_com_mwitkow_go_proto_validators.FieldError("ApplicationCredentialId", fmt.Errorf("message must exist"))
	}
	if this.ApplicationCredentialId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ApplicationCredentialId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ApplicationCredentialId", err)
		}
	}
	return nil
}
