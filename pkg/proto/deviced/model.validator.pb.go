// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: model.proto

package deviced

import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"
import _ "github.com/golang/protobuf/ptypes/struct"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/nayotta/metathings/pkg/proto/constant/kind"
import _ "github.com/nayotta/metathings/pkg/proto/constant/state"
import _ "github.com/nayotta/metathings/pkg/proto/identityd2"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Device) Validate() error {
	for _, item := range this.Modules {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Modules", err)
			}
		}
	}
	if this.HeartbeatAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.HeartbeatAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("HeartbeatAt", err)
		}
	}
	for _, item := range this.Flows {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Flows", err)
			}
		}
	}
	return nil
}
func (this *OpDevice) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	if this.Name != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Name", err)
		}
	}
	if this.Alias != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Alias); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Alias", err)
		}
	}
	for _, item := range this.Modules {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Modules", err)
			}
		}
	}
	if this.HeartbeatAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.HeartbeatAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("HeartbeatAt", err)
		}
	}
	for _, item := range this.Flows {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Flows", err)
			}
		}
	}
	return nil
}
func (this *Module) Validate() error {
	if this.HeartbeatAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.HeartbeatAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("HeartbeatAt", err)
		}
	}
	return nil
}
func (this *OpModule) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	if this.DeviceId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeviceId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeviceId", err)
		}
	}
	if this.Endpoint != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Endpoint); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Endpoint", err)
		}
	}
	if this.Component != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Component); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Component", err)
		}
	}
	if this.Name != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Name", err)
		}
	}
	if this.Alias != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Alias); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Alias", err)
		}
	}
	if this.HeartbeatAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.HeartbeatAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("HeartbeatAt", err)
		}
	}
	return nil
}
func (this *Flow) Validate() error {
	return nil
}
func (this *OpFlow) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	if this.DeviceId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeviceId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeviceId", err)
		}
	}
	if this.Name != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Name", err)
		}
	}
	if this.Alias != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Alias); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Alias", err)
		}
	}
	return nil
}
func (this *Frame) Validate() error {
	if this.Ts != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Ts); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Ts", err)
		}
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *OpFrame) Validate() error {
	if this.Ts != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Ts); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Ts", err)
		}
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *Object) Validate() error {
	if this.Device != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Device); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Device", err)
		}
	}
	if this.LastModified != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LastModified); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LastModified", err)
		}
	}
	return nil
}
func (this *OpObject) Validate() error {
	if this.Device != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Device); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Device", err)
		}
	}
	if this.Prefix != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Prefix); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Prefix", err)
		}
	}
	if this.Name != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Name", err)
		}
	}
	if this.Length != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Length); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Length", err)
		}
	}
	if this.Etag != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Etag); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Etag", err)
		}
	}
	if this.LiastModified != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LiastModified); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LiastModified", err)
		}
	}
	return nil
}
func (this *ErrorValue) Validate() error {
	return nil
}
func (this *OpUnaryCallValue) Validate() error {
	if this.Name != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Name", err)
		}
	}
	if this.Component != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Component); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Component", err)
		}
	}
	if this.Method != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Method); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Method", err)
		}
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
	if oneOfNester, ok := this.GetUnion().(*OpStreamCallValue_ConfigAck); ok {
		if oneOfNester.ConfigAck != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.ConfigAck); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ConfigAck", err)
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
	if oneOfNester, ok := this.GetUnion().(*StreamCallValue_ConfigAck); ok {
		if oneOfNester.ConfigAck != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.ConfigAck); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ConfigAck", err)
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
func (this *OpStreamCallConfig) Validate() error {
	if this.Name != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Name", err)
		}
	}
	if this.Component != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Component); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Component", err)
		}
	}
	if this.Method != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Method); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Method", err)
		}
	}
	return nil
}
func (this *StreamCallConfig) Validate() error {
	return nil
}
func (this *OpStreamCallConfigAck) Validate() error {
	return nil
}
func (this *StreamCallConfigAck) Validate() error {
	return nil
}
func (this *OpStreamCallExit) Validate() error {
	return nil
}
func (this *StreamCallExit) Validate() error {
	return nil
}
