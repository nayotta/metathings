// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: show.proto

package switcher

import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ShowResponse) Validate() error {
	if this.Switcher != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Switcher); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Switcher", err)
		}
	}
	return nil
}
