// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: upload_descriptor.proto

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
var _upload_descriptor_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on UploadDescriptorRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UploadDescriptorRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetDescriptor_() == nil {
		return UploadDescriptorRequestValidationError{
			field:  "Descriptor_",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetDescriptor_()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UploadDescriptorRequestValidationError{
				field:  "Descriptor_",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UploadDescriptorRequestValidationError is the validation error returned by
// UploadDescriptorRequest.Validate if the designated constraints aren't met.
type UploadDescriptorRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UploadDescriptorRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UploadDescriptorRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UploadDescriptorRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UploadDescriptorRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UploadDescriptorRequestValidationError) ErrorName() string {
	return "UploadDescriptorRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UploadDescriptorRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUploadDescriptorRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UploadDescriptorRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UploadDescriptorRequestValidationError{}

// Validate checks the field values on UploadDescriptorResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UploadDescriptorResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetDescriptor_()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UploadDescriptorResponseValidationError{
				field:  "Descriptor_",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UploadDescriptorResponseValidationError is the validation error returned by
// UploadDescriptorResponse.Validate if the designated constraints aren't met.
type UploadDescriptorResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UploadDescriptorResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UploadDescriptorResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UploadDescriptorResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UploadDescriptorResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UploadDescriptorResponseValidationError) ErrorName() string {
	return "UploadDescriptorResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UploadDescriptorResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUploadDescriptorResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UploadDescriptorResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UploadDescriptorResponseValidationError{}