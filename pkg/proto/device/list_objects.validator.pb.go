// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_objects.proto

package ai_metathings_service_device

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/nayotta/metathings/pkg/proto/deviced"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ListObjectsRequest) Validate() error {
	if nil == this.Object {
		return github_com_mwitkow_go_proto_validators.FieldError("Object", fmt.Errorf("message must exist"))
	}
	if this.Object != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Object); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Object", err)
		}
	}
	return nil
}
func (this *ListObjectsResponse) Validate() error {
	for _, item := range this.Objects {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Objects", err)
			}
		}
	}
	return nil
}
