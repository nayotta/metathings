// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: push_frame_to_flow.proto

package ai_metathings_service_device

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/nayotta/metathings/pkg/proto/deviced"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *PushFrameToFlowRequest) Validate() error {
	if nil == this.Id {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf("message must exist"))
	}
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	if oneOfNester, ok := this.GetRequest().(*PushFrameToFlowRequest_Config_); ok {
		if oneOfNester.Config != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Config); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Config", err)
			}
		}
	}
	if oneOfNester, ok := this.GetRequest().(*PushFrameToFlowRequest_Ping_); ok {
		if oneOfNester.Ping != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Ping); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Ping", err)
			}
		}
	}
	if oneOfNester, ok := this.GetRequest().(*PushFrameToFlowRequest_Frame); ok {
		if oneOfNester.Frame != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Frame); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Frame", err)
			}
		}
	}
	return nil
}
func (this *PushFrameToFlowRequest_Config) Validate() error {
	if nil == this.Flow {
		return github_com_mwitkow_go_proto_validators.FieldError("Flow", fmt.Errorf("message must exist"))
	}
	if this.Flow != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Flow); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Flow", err)
		}
	}
	if this.ConfigAck != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ConfigAck); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ConfigAck", err)
		}
	}
	if this.PushAck != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PushAck); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PushAck", err)
		}
	}
	return nil
}
func (this *PushFrameToFlowRequest_Ping) Validate() error {
	return nil
}
func (this *PushFrameToFlowResponse) Validate() error {
	if oneOfNester, ok := this.GetResponse().(*PushFrameToFlowResponse_Config_); ok {
		if oneOfNester.Config != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Config); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Config", err)
			}
		}
	}
	if oneOfNester, ok := this.GetResponse().(*PushFrameToFlowResponse_Pong_); ok {
		if oneOfNester.Pong != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Pong); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Pong", err)
			}
		}
	}
	if oneOfNester, ok := this.GetResponse().(*PushFrameToFlowResponse_Ack_); ok {
		if oneOfNester.Ack != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Ack); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Ack", err)
			}
		}
	}
	return nil
}
func (this *PushFrameToFlowResponse_Config) Validate() error {
	return nil
}
func (this *PushFrameToFlowResponse_Ack) Validate() error {
	return nil
}
func (this *PushFrameToFlowResponse_Pong) Validate() error {
	return nil
}
