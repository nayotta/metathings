// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_object_content.proto

package ai_metathings_service_device

import fmt "fmt"
import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"
import _ "github.com/nayotta/metathings/pkg/proto/deviced"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GetObjectContentRequest) Validate() error {
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
func (this *GetObjectContentResponse) Validate() error {
	return nil
}