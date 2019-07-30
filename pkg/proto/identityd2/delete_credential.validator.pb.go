// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete_credential.proto

package identityd2

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

func (this *DeleteCredentialRequest) Validate() error {
	if nil == this.Credential {
		return github_com_mwitkow_go_proto_validators.FieldError("Credential", fmt.Errorf("message must exist"))
	}
	if this.Credential != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Credential); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Credential", err)
		}
	}
	return nil
}
