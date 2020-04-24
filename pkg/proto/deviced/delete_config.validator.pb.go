// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete_config.proto

package deviced

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *DeleteConfigRequest) Validate() error {
	if nil == this.Config {
		return github_com_mwitkow_go_proto_validators.FieldError("Config", fmt.Errorf("message must exist"))
	}
	if this.Config != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Config); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Config", err)
		}
	}
	return nil
}
