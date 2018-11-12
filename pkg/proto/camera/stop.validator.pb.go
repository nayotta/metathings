// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stop.proto

package camera

import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *StopResponse) Validate() error {
	if this.Camera != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Camera); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Camera", err)
		}
	}
	return nil
}
