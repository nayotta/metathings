// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: get_group.proto

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
var _get_group_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetGroupRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetGroupRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetGroup() == nil {
		return GetGroupRequestValidationError{
			field:  "Group",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetGroup()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetGroupRequestValidationError{
				field:  "Group",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetGroupRequestValidationError is the validation error returned by
// GetGroupRequest.Validate if the designated constraints aren't met.
type GetGroupRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetGroupRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetGroupRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetGroupRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetGroupRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetGroupRequestValidationError) ErrorName() string { return "GetGroupRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetGroupRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetGroupRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetGroupRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetGroupRequestValidationError{}

// Validate checks the field values on GetGroupResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetGroupResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetGroup()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetGroupResponseValidationError{
				field:  "Group",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetGroupResponseValidationError is the validation error returned by
// GetGroupResponse.Validate if the designated constraints aren't met.
type GetGroupResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetGroupResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetGroupResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetGroupResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetGroupResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetGroupResponseValidationError) ErrorName() string { return "GetGroupResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetGroupResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetGroupResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetGroupResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetGroupResponseValidationError{}