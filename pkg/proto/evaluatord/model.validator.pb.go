// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: model.proto

package evaluatord

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/nayotta/metathings/pkg/proto/constant/state"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Resource) Validate() error {
	return nil
}
func (this *OpResource) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	if this.Type != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Type); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Type", err)
		}
	}
	return nil
}
func (this *Evaluator) Validate() error {
	for _, item := range this.Sources {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Sources", err)
			}
		}
	}
	if this.Operator != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Operator); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Operator", err)
		}
	}
	if this.Config != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Config); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Config", err)
		}
	}
	return nil
}
func (this *OpEvaluator) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	if this.Alias != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Alias); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Alias", err)
		}
	}
	if this.Description != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Description); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Description", err)
		}
	}
	for _, item := range this.Sources {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Sources", err)
			}
		}
	}
	if this.Operator != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Operator); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Operator", err)
		}
	}
	if this.Config != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Config); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Config", err)
		}
	}
	return nil
}
func (this *Operator) Validate() error {
	if oneOfNester, ok := this.GetDescriptor_().(*Operator_Lua); ok {
		if oneOfNester.Lua != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Lua); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Lua", err)
			}
		}
	}
	return nil
}
func (this *OpOperator) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	if this.Alias != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Alias); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Alias", err)
		}
	}
	if this.Description != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Description); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Description", err)
		}
	}
	if this.Driver != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Driver); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Driver", err)
		}
	}
	if oneOfNester, ok := this.GetDescriptor_().(*OpOperator_Lua); ok {
		if oneOfNester.Lua != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Lua); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Lua", err)
			}
		}
	}
	return nil
}
func (this *LuaDescriptor) Validate() error {
	return nil
}
func (this *OpLuaDescriptor) Validate() error {
	if this.Code != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Code); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Code", err)
		}
	}
	return nil
}
func (this *OpTask) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	if this.UpdatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdatedAt", err)
		}
	}
	if this.Source != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Source); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Source", err)
		}
	}
	return nil
}
func (this *Task) Validate() error {
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	if this.UpdatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdatedAt", err)
		}
	}
	if this.Source != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Source); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Source", err)
		}
	}
	for _, item := range this.StateTimeline {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StateTimeline", err)
			}
		}
	}
	return nil
}
func (this *Task_StateNode) Validate() error {
	if this.At != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.At); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("At", err)
		}
	}
	if this.Tags != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Tags); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Tags", err)
		}
	}
	return nil
}
func (this *OpTimer) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	if this.Alias != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Alias); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Alias", err)
		}
	}
	if this.Description != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Description); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Description", err)
		}
	}
	if this.Schedule != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Schedule); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Schedule", err)
		}
	}
	if this.Timezone != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Timezone); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Timezone", err)
		}
	}
	if this.Enabled != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Enabled); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Enabled", err)
		}
	}
	if this.Config != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Config); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Config", err)
		}
	}
	return nil
}
func (this *Timer) Validate() error {
	if this.Config != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Config); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Config", err)
		}
	}
	return nil
}
