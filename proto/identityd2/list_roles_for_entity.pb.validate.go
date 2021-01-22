// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: list_roles_for_entity.proto

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
var _list_roles_for_entity_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ListRolesForEntityRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListRolesForEntityRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRole()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListRolesForEntityRequestValidationError{
				field:  "Role",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ListRolesForEntityRequestValidationError is the validation error returned by
// ListRolesForEntityRequest.Validate if the designated constraints aren't met.
type ListRolesForEntityRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRolesForEntityRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRolesForEntityRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRolesForEntityRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRolesForEntityRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRolesForEntityRequestValidationError) ErrorName() string {
	return "ListRolesForEntityRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListRolesForEntityRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRolesForEntityRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRolesForEntityRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRolesForEntityRequestValidationError{}

// Validate checks the field values on ListRolesForEntityResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListRolesForEntityResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRoles() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListRolesForEntityResponseValidationError{
					field:  fmt.Sprintf("Roles[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListRolesForEntityResponseValidationError is the validation error returned
// by ListRolesForEntityResponse.Validate if the designated constraints aren't met.
type ListRolesForEntityResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRolesForEntityResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRolesForEntityResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRolesForEntityResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRolesForEntityResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRolesForEntityResponseValidationError) ErrorName() string {
	return "ListRolesForEntityResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListRolesForEntityResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRolesForEntityResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRolesForEntityResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRolesForEntityResponseValidationError{}