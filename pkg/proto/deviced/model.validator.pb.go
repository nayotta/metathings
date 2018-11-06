// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: model.proto

package deviced

import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/golang/protobuf/ptypes/any"
import _ "github.com/nayotta/metathings/pkg/proto/identityd2"
import _ "github.com/nayotta/metathings/pkg/proto/constant/state"
import _ "github.com/nayotta/metathings/pkg/proto/constant/kind"
import _ "github.com/nayotta/metathings/pkg/proto/constant/state"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Device) Validate() error {
	if this.Entity != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Entity); err != nil {
			return go_proto_validators.FieldError("Entity", err)
		}
	}
	for _, item := range this.Modules {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Modules", err)
			}
		}
	}
	return nil
}
func (this *OpDevice) Validate() error {
	if this.Id != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return go_proto_validators.FieldError("Id", err)
		}
	}
	if this.Entity != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Entity); err != nil {
			return go_proto_validators.FieldError("Entity", err)
		}
	}
	if this.Name != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return go_proto_validators.FieldError("Name", err)
		}
	}
	if this.Alias != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Alias); err != nil {
			return go_proto_validators.FieldError("Alias", err)
		}
	}
	for _, item := range this.Modules {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Modules", err)
			}
		}
	}
	return nil
}
func (this *Module) Validate() error {
	if this.Entity != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Entity); err != nil {
			return go_proto_validators.FieldError("Entity", err)
		}
	}
	return nil
}
func (this *OpModule) Validate() error {
	if this.Id != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return go_proto_validators.FieldError("Id", err)
		}
	}
	if this.Entity != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Entity); err != nil {
			return go_proto_validators.FieldError("Entity", err)
		}
	}
	if this.DeviceId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.DeviceId); err != nil {
			return go_proto_validators.FieldError("DeviceId", err)
		}
	}
	if this.Endpoint != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Endpoint); err != nil {
			return go_proto_validators.FieldError("Endpoint", err)
		}
	}
	if this.Name != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return go_proto_validators.FieldError("Name", err)
		}
	}
	if this.Alias != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Alias); err != nil {
			return go_proto_validators.FieldError("Alias", err)
		}
	}
	return nil
}
func (this *ErrorValue) Validate() error {
	return nil
}
func (this *OpUnaryCallValue) Validate() error {
	if this.ModuleName != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.ModuleName); err != nil {
			return go_proto_validators.FieldError("ModuleName", err)
		}
	}
	if this.ComponentName != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.ComponentName); err != nil {
			return go_proto_validators.FieldError("ComponentName", err)
		}
	}
	if this.MethodName != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.MethodName); err != nil {
			return go_proto_validators.FieldError("MethodName", err)
		}
	}
	if this.Value != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
func (this *UnaryCallValue) Validate() error {
	if this.Value != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
func (this *OpStreamCallValue) Validate() error {
	if oneOfNester, ok := this.GetUnion().(*OpStreamCallValue_Config); ok {
		if oneOfNester.Config != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.Config); err != nil {
				return go_proto_validators.FieldError("Config", err)
			}
		}
	}
	if oneOfNester, ok := this.GetUnion().(*OpStreamCallValue_Data); ok {
		if oneOfNester.Data != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.Data); err != nil {
				return go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *StreamCallValue) Validate() error {
	if oneOfNester, ok := this.GetUnion().(*StreamCallValue_Config); ok {
		if oneOfNester.Config != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.Config); err != nil {
				return go_proto_validators.FieldError("Config", err)
			}
		}
	}
	if oneOfNester, ok := this.GetUnion().(*StreamCallValue_Data); ok {
		if oneOfNester.Data != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.Data); err != nil {
				return go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *OpStreamCallConfig) Validate() error {
	if this.ModuleName != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.ModuleName); err != nil {
			return go_proto_validators.FieldError("ModuleName", err)
		}
	}
	if this.ComponentName != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.ComponentName); err != nil {
			return go_proto_validators.FieldError("ComponentName", err)
		}
	}
	if this.MethodName != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.MethodName); err != nil {
			return go_proto_validators.FieldError("MethodName", err)
		}
	}
	return nil
}
func (this *StreamCallConfig) Validate() error {
	return nil
}
func (this *OpStreamCallData) Validate() error {
	if this.Value != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
func (this *StreamCallData) Validate() error {
	if this.Value != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}