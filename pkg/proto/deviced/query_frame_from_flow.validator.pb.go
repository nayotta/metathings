// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: query_frame_from_flow.proto

package deviced

import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *QueryFrameFromFlowRequest) Validate() error {
	if this.Device != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Device); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Device", err)
		}
	}
	if this.From != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.From); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("From", err)
		}
	}
	if this.To != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.To); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("To", err)
		}
	}
	return nil
}
func (this *QueryFrameFromFlowResponse) Validate() error {
	for _, item := range this.Frames {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Frames", err)
			}
		}
	}
	return nil
}
