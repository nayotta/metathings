// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: get_descriptor.proto

package deviced

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
var _get_descriptor_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetDescriptorRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetDescriptorRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetDescriptor_() == nil {
		return GetDescriptorRequestValidationError{
			field:  "Descriptor_",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetDescriptor_()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetDescriptorRequestValidationError{
				field:  "Descriptor_",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetDescriptorRequestValidationError is the validation error returned by
// GetDescriptorRequest.Validate if the designated constraints aren't met.
type GetDescriptorRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDescriptorRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDescriptorRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDescriptorRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDescriptorRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDescriptorRequestValidationError) ErrorName() string {
	return "GetDescriptorRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetDescriptorRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDescriptorRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDescriptorRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDescriptorRequestValidationError{}

// Validate checks the field values on GetDescriptorResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetDescriptorResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetDescriptor_()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetDescriptorResponseValidationError{
				field:  "Descriptor_",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetDescriptorResponseValidationError is the validation error returned by
// GetDescriptorResponse.Validate if the designated constraints aren't met.
type GetDescriptorResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDescriptorResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDescriptorResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDescriptorResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDescriptorResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDescriptorResponseValidationError) ErrorName() string {
	return "GetDescriptorResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetDescriptorResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDescriptorResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDescriptorResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDescriptorResponseValidationError{}