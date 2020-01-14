// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: put_object_streaming.proto

package deviced

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *PutObjectStreamingRequest) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	if oneOfNester, ok := this.GetRequest().(*PutObjectStreamingRequest_Metadata_); ok {
		if oneOfNester.Metadata != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Metadata); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Metadata", err)
			}
		}
	}
	if oneOfNester, ok := this.GetRequest().(*PutObjectStreamingRequest_Chunks); ok {
		if oneOfNester.Chunks != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Chunks); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Chunks", err)
			}
		}
	}
	if oneOfNester, ok := this.GetRequest().(*PutObjectStreamingRequest_Ack_); ok {
		if oneOfNester.Ack != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Ack); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Ack", err)
			}
		}
	}
	return nil
}
func (this *PutObjectStreamingRequest_Ack) Validate() error {
	return nil
}
func (this *PutObjectStreamingRequest_Metadata) Validate() error {
	if this.Object != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Object); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Object", err)
		}
	}
	if this.Sha1 != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Sha1); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Sha1", err)
		}
	}
	return nil
}
func (this *PutObjectStreamingResponse) Validate() error {
	if oneOfNester, ok := this.GetResponse().(*PutObjectStreamingResponse_Chunks); ok {
		if oneOfNester.Chunks != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Chunks); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Chunks", err)
			}
		}
	}
	if oneOfNester, ok := this.GetResponse().(*PutObjectStreamingResponse_Ack_); ok {
		if oneOfNester.Ack != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Ack); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Ack", err)
			}
		}
	}
	return nil
}
func (this *PutObjectStreamingResponse_Ack) Validate() error {
	return nil
}
