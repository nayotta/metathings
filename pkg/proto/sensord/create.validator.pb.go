// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create.proto

/*
Package sensord is a generated protocol buffer package.

It is generated from these files:
	create.proto
	delete.proto
	get.proto
	list.proto
	list_for_user.proto
	patch.proto
	publish.proto
	sensor.proto
	service.proto
	subscribe.proto

It has these top-level messages:
	CreateRequest
	CreateResponse
	DeleteRequest
	GetRequest
	GetResponse
	ListRequest
	ListResponse
	ListForUserRequest
	ListForUserResponse
	PatchRequest
	PatchResponse
	PublishRequest
	PublishRequests
	PublishResponse
	PublishResponses
	Sensor
	SensorData
	SubscribeByIdRequest
	UnsubscribeByIdRequest
	SubscribeByUserIdRequest
	UnsubscribeByUserIdRequest
	SubscribeByCoreIdRequest
	UnsubscribeByCoreIdRequest
	SubscribeRequest
	SubscribeRequests
	SubscribeResponse
	SubscribeResponses
*/
package sensord

import fmt "fmt"
import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CreateRequest) Validate() error {
	if nil == this.Name {
		return go_proto_validators.FieldError("Name", fmt.Errorf("message must exist"))
	}
	if this.Name != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return go_proto_validators.FieldError("Name", err)
		}
	}
	if nil == this.CoreId {
		return go_proto_validators.FieldError("CoreId", fmt.Errorf("message must exist"))
	}
	if this.CoreId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.CoreId); err != nil {
			return go_proto_validators.FieldError("CoreId", err)
		}
	}
	if nil == this.EntityName {
		return go_proto_validators.FieldError("EntityName", fmt.Errorf("message must exist"))
	}
	if this.EntityName != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.EntityName); err != nil {
			return go_proto_validators.FieldError("EntityName", err)
		}
	}
	if nil == this.ApplicationCredentialId {
		return go_proto_validators.FieldError("ApplicationCredentialId", fmt.Errorf("message must exist"))
	}
	if this.ApplicationCredentialId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.ApplicationCredentialId); err != nil {
			return go_proto_validators.FieldError("ApplicationCredentialId", err)
		}
	}
	return nil
}
func (this *CreateResponse) Validate() error {
	if this.Sensor != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Sensor); err != nil {
			return go_proto_validators.FieldError("Sensor", err)
		}
	}
	return nil
}
