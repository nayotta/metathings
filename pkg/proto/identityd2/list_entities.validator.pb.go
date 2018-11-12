// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_entities.proto

package identityd2

import fmt "fmt"
import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ListEntitiesRequest) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	if nil == this.Domain {
		return github_com_mwitkow_go_proto_validators.FieldError("Domain", fmt.Errorf("message must exist"))
	}
	if this.Domain != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Domain); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Domain", err)
		}
	}
	for _, item := range this.Groups {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Groups", err)
			}
		}
	}
	for _, item := range this.Roles {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Roles", err)
			}
		}
	}
	if nil == this.Name {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf("message must exist"))
	}
	if this.Name != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Name", err)
		}
	}
	if nil == this.Alias {
		return github_com_mwitkow_go_proto_validators.FieldError("Alias", fmt.Errorf("message must exist"))
	}
	if this.Alias != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Alias); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Alias", err)
		}
	}
	return nil
}
func (this *ListEntitiesResponse) Validate() error {
	for _, item := range this.Entities {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Entities", err)
			}
		}
	}
	return nil
}
