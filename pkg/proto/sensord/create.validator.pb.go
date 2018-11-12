// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create.proto

package sensord

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

func (this *CreateRequest) Validate() error {
	if nil == this.Name {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf("message must exist"))
	}
	if this.Name != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Name", err)
		}
	}
	if nil == this.CoreId {
		return github_com_mwitkow_go_proto_validators.FieldError("CoreId", fmt.Errorf("message must exist"))
	}
	if this.CoreId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CoreId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CoreId", err)
		}
	}
	if nil == this.EntityName {
		return github_com_mwitkow_go_proto_validators.FieldError("EntityName", fmt.Errorf("message must exist"))
	}
	if this.EntityName != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EntityName); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EntityName", err)
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
func (this *CreateResponse) Validate() error {
	if this.Sensor != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Sensor); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Sensor", err)
		}
	}
	return nil
}
