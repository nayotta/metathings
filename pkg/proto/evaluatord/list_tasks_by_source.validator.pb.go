// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_tasks_by_source.proto

package evaluatord

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ListTasksBySourceRequest) Validate() error {
	if nil == this.Source {
		return github_com_mwitkow_go_proto_validators.FieldError("Source", fmt.Errorf("message must exist"))
	}
	if this.Source != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Source); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Source", err)
		}
	}
	if this.RangeFrom != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RangeFrom); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RangeFrom", err)
		}
	}
	if this.RangeTo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RangeTo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RangeTo", err)
		}
	}
	return nil
}
func (this *ListTasksBySourceResponse) Validate() error {
	for _, item := range this.Tasks {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Tasks", err)
			}
		}
	}
	return nil
}