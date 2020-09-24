// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: add_action_to_role.proto

package identityd2

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _add_action_to_role_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on AddActionToRoleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AddActionToRoleRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetAction() == nil {
		return AddActionToRoleRequestValidationError{
			field:  "Action",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetAction()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddActionToRoleRequestValidationError{
				field:  "Action",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetRole() == nil {
		return AddActionToRoleRequestValidationError{
			field:  "Role",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetRole()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddActionToRoleRequestValidationError{
				field:  "Role",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// AddActionToRoleRequestValidationError is the validation error returned by
// AddActionToRoleRequest.Validate if the designated constraints aren't met.
type AddActionToRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddActionToRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddActionToRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddActionToRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddActionToRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddActionToRoleRequestValidationError) ErrorName() string {
	return "AddActionToRoleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddActionToRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddActionToRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddActionToRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddActionToRoleRequestValidationError{}
