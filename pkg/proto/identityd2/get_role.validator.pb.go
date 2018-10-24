// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_role.proto

package identityd2

import fmt "fmt"
import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GetRoleRequest) Validate() error {
	if nil == this.Role {
		return go_proto_validators.FieldError("Role", fmt.Errorf("message must exist"))
	}
	if this.Role != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Role); err != nil {
			return go_proto_validators.FieldError("Role", err)
		}
	}
	return nil
}
func (this *GetRoleResponse) Validate() error {
	if this.Role != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Role); err != nil {
			return go_proto_validators.FieldError("Role", err)
		}
	}
	return nil
}
