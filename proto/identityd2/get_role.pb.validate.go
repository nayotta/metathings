// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: get_role.proto

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
var _get_role_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetRoleRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetRoleRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetRole() == nil {
		return GetRoleRequestValidationError{
			field:  "Role",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetRole()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetRoleRequestValidationError{
				field:  "Role",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetRoleRequestValidationError is the validation error returned by
// GetRoleRequest.Validate if the designated constraints aren't met.
type GetRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRoleRequestValidationError) ErrorName() string { return "GetRoleRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRoleRequestValidationError{}

// Validate checks the field values on GetRoleResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetRoleResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRole()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetRoleResponseValidationError{
				field:  "Role",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetRoleResponseValidationError is the validation error returned by
// GetRoleResponse.Validate if the designated constraints aren't met.
type GetRoleResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRoleResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRoleResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRoleResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRoleResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRoleResponseValidationError) ErrorName() string { return "GetRoleResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetRoleResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRoleResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRoleResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRoleResponseValidationError{}
