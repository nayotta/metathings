// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pull_frame_from_flow_set.proto

package deviced

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *PullFrameFromFlowSetRequest) Validate() error {
	if nil == this.Id {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf("message must exist"))
	}
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	if oneOfNester, ok := this.GetRequest().(*PullFrameFromFlowSetRequest_Config_); ok {
		if oneOfNester.Config != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Config); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Config", err)
			}
		}
	}
	return nil
}
func (this *PullFrameFromFlowSetRequest_Config) Validate() error {
	if this.FlowSet != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.FlowSet); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("FlowSet", err)
		}
	}
	if this.ConfigAck != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ConfigAck); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ConfigAck", err)
		}
	}
	return nil
}
func (this *PullFrameFromFlowSetResponse) Validate() error {
	if oneOfNester, ok := this.GetResponse().(*PullFrameFromFlowSetResponse_Ack_); ok {
		if oneOfNester.Ack != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Ack); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Ack", err)
			}
		}
	}
	if oneOfNester, ok := this.GetResponse().(*PullFrameFromFlowSetResponse_Pack_); ok {
		if oneOfNester.Pack != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Pack); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Pack", err)
			}
		}
	}
	return nil
}
func (this *PullFrameFromFlowSetResponse_Ack) Validate() error {
	return nil
}
func (this *PullFrameFromFlowSetResponse_Pack) Validate() error {
	if this.Device != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Device); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Device", err)
		}
	}
	for _, item := range this.Frames {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Frames", err)
			}
		}
	}
	return nil
}