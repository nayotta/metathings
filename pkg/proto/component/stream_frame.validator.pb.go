// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stream_frame.proto

package component

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/mwitkow/go-proto-validators"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *OpStreamCallConfig) Validate() error {
	if this.Session != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Session); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Session", err)
		}
	}
	if this.Method != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Method); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Method", err)
		}
	}
	if this.Ack != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Ack); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Ack", err)
		}
	}
	return nil
}
func (this *StreamCallConfig) Validate() error {
	return nil
}
func (this *OpStreamCallAck) Validate() error {
	if this.Value != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
func (this *StreamCallAck) Validate() error {
	return nil
}
func (this *OpStreamCallExit) Validate() error {
	return nil
}
func (this *StreamCallExit) Validate() error {
	return nil
}
func (this *OpUnaryCallValue) Validate() error {
	if nil == this.Session {
		return github_com_mwitkow_go_proto_validators.FieldError("Session", fmt.Errorf("message must exist"))
	}
	if this.Session != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Session); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Session", err)
		}
	}
	if nil == this.Method {
		return github_com_mwitkow_go_proto_validators.FieldError("Method", fmt.Errorf("message must exist"))
	}
	if this.Method != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Method); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Method", err)
		}
	}
	if nil == this.Value {
		return github_com_mwitkow_go_proto_validators.FieldError("Value", fmt.Errorf("message must exist"))
	}
	if this.Value != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
func (this *UnaryCallValue) Validate() error {
	if this.Value != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
func (this *OpStreamCallValue) Validate() error {
	if oneOfNester, ok := this.GetUnion().(*OpStreamCallValue_Value); ok {
		if oneOfNester.Value != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Value); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Value", err)
			}
		}
	}
	if oneOfNester, ok := this.GetUnion().(*OpStreamCallValue_Config); ok {
		if oneOfNester.Config != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Config); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Config", err)
			}
		}
	}
	if oneOfNester, ok := this.GetUnion().(*OpStreamCallValue_Ack); ok {
		if oneOfNester.Ack != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Ack); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Ack", err)
			}
		}
	}
	if oneOfNester, ok := this.GetUnion().(*OpStreamCallValue_Exit); ok {
		if oneOfNester.Exit != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Exit); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Exit", err)
			}
		}
	}
	return nil
}
func (this *StreamCallValue) Validate() error {
	if oneOfNester, ok := this.GetUnion().(*StreamCallValue_Value); ok {
		if oneOfNester.Value != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Value); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Value", err)
			}
		}
	}
	if oneOfNester, ok := this.GetUnion().(*StreamCallValue_Config); ok {
		if oneOfNester.Config != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Config); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Config", err)
			}
		}
	}
	if oneOfNester, ok := this.GetUnion().(*StreamCallValue_Ack); ok {
		if oneOfNester.Ack != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Ack); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Ack", err)
			}
		}
	}
	if oneOfNester, ok := this.GetUnion().(*StreamCallValue_Exit); ok {
		if oneOfNester.Exit != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Exit); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Exit", err)
			}
		}
	}
	return nil
}
func (this *ErrorValue) Validate() error {
	return nil
}
func (this *UpStreamFrame) Validate() error {
	if oneOfNester, ok := this.GetUnion().(*UpStreamFrame_UnaryCall); ok {
		if oneOfNester.UnaryCall != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.UnaryCall); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("UnaryCall", err)
			}
		}
	}
	if oneOfNester, ok := this.GetUnion().(*UpStreamFrame_StreamCall); ok {
		if oneOfNester.StreamCall != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.StreamCall); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StreamCall", err)
			}
		}
	}
	if oneOfNester, ok := this.GetUnion().(*UpStreamFrame_Error); ok {
		if oneOfNester.Error != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Error); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Error", err)
			}
		}
	}
	return nil
}
func (this *DownStreamFrame) Validate() error {
	if oneOfNester, ok := this.GetUnion().(*DownStreamFrame_UnaryCall); ok {
		if oneOfNester.UnaryCall != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.UnaryCall); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("UnaryCall", err)
			}
		}
	}
	if oneOfNester, ok := this.GetUnion().(*DownStreamFrame_StreamCall); ok {
		if oneOfNester.StreamCall != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.StreamCall); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StreamCall", err)
			}
		}
	}
	return nil
}
