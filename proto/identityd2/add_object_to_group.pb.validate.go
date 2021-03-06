// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: add_object_to_group.proto

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
var _add_object_to_group_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on AddObjectToGroupRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AddObjectToGroupRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetGroup() == nil {
		return AddObjectToGroupRequestValidationError{
			field:  "Group",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetGroup()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddObjectToGroupRequestValidationError{
				field:  "Group",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetObject() == nil {
		return AddObjectToGroupRequestValidationError{
			field:  "Object",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetObject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddObjectToGroupRequestValidationError{
				field:  "Object",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// AddObjectToGroupRequestValidationError is the validation error returned by
// AddObjectToGroupRequest.Validate if the designated constraints aren't met.
type AddObjectToGroupRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddObjectToGroupRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddObjectToGroupRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddObjectToGroupRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddObjectToGroupRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddObjectToGroupRequestValidationError) ErrorName() string {
	return "AddObjectToGroupRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddObjectToGroupRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddObjectToGroupRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddObjectToGroupRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddObjectToGroupRequestValidationError{}
