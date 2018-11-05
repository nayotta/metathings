// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service.proto

/*
Package policyd is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	NewEnforcerRequest
	NewEnforcerReply
	NewAdapterRequest
	NewAdapterReply
	EnforceRequest
	EnforceBucketRequest
	BoolReply
	EmptyRequest
	EmptyReply
	PolicyRequest
	SimpleGetRequest
	ArrayReply
	FilteredPolicyRequest
	UserRoleRequest
	PermissionRequest
	Array2DReply
*/
package policyd

import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *NewEnforcerRequest) Validate() error {
	return nil
}
func (this *NewEnforcerReply) Validate() error {
	return nil
}
func (this *NewAdapterRequest) Validate() error {
	return nil
}
func (this *NewAdapterReply) Validate() error {
	return nil
}
func (this *EnforceRequest) Validate() error {
	return nil
}
func (this *EnforceBucketRequest) Validate() error {
	for _, item := range this.Requests {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Requests", err)
			}
		}
	}
	return nil
}
func (this *BoolReply) Validate() error {
	return nil
}
func (this *EmptyRequest) Validate() error {
	return nil
}
func (this *EmptyReply) Validate() error {
	return nil
}
func (this *PolicyRequest) Validate() error {
	return nil
}
func (this *SimpleGetRequest) Validate() error {
	return nil
}
func (this *ArrayReply) Validate() error {
	return nil
}
func (this *FilteredPolicyRequest) Validate() error {
	return nil
}
func (this *UserRoleRequest) Validate() error {
	return nil
}
func (this *PermissionRequest) Validate() error {
	return nil
}
func (this *Array2DReply) Validate() error {
	for _, item := range this.D2 {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("D2", err)
			}
		}
	}
	return nil
}
func (this *Array2DReplyD) Validate() error {
	return nil
}