// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue_module_token.proto

package ai_metathings_service_device

import fmt "fmt"
import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"
import _ "github.com/nayotta/metathings/pkg/proto/identityd2"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *IssueModuleTokenRequest) Validate() error {
	if nil == this.Credential {
		return github_com_mwitkow_go_proto_validators.FieldError("Credential", fmt.Errorf("message must exist"))
	}
	if this.Credential != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Credential); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Credential", err)
		}
	}
	if nil == this.Timestamp {
		return github_com_mwitkow_go_proto_validators.FieldError("Timestamp", fmt.Errorf("message must exist"))
	}
	if this.Timestamp != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Timestamp); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Timestamp", err)
		}
	}
	if nil == this.Nonce {
		return github_com_mwitkow_go_proto_validators.FieldError("Nonce", fmt.Errorf("message must exist"))
	}
	if this.Nonce != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Nonce); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Nonce", err)
		}
	}
	if nil == this.Hmac {
		return github_com_mwitkow_go_proto_validators.FieldError("Hmac", fmt.Errorf("message must exist"))
	}
	if this.Hmac != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Hmac); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Hmac", err)
		}
	}
	return nil
}
func (this *IssueModuleTokenResponse) Validate() error {
	if this.Token != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Token); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Token", err)
		}
	}
	return nil
}