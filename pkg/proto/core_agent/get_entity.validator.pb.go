// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_entity.proto

package core_agent

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

func (this *GetEntityRequest) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	return nil
}
func (this *GetEntityResponse) Validate() error {
	if this.Entity != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Entity); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Entity", err)
		}
	}
	return nil
}
