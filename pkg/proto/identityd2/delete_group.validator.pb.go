// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete_group.proto

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

func (this *DeleteGroupRequest) Validate() error {
	if nil == this.Group {
		return go_proto_validators.FieldError("Group", fmt.Errorf("message must exist"))
	}
	if this.Group != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Group); err != nil {
			return go_proto_validators.FieldError("Group", err)
		}
	}
	return nil
}
