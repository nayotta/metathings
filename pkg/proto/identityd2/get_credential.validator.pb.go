// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_credential.proto

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

func (this *GetCredentialRequest) Validate() error {
	if nil == this.Credential {
		return github_com_mwitkow_go_proto_validators.FieldError("Credential", fmt.Errorf("message must exist"))
	}
	if this.Credential != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Credential); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Credential", err)
		}
	}
	return nil
}
func (this *GetCredentialResponse) Validate() error {
	if this.Credential != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Credential); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Credential", err)
		}
	}
	return nil
}
