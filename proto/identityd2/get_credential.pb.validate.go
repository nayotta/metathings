// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: get_credential.proto

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
var _get_credential_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetCredentialRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetCredentialRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetCredential() == nil {
		return GetCredentialRequestValidationError{
			field:  "Credential",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetCredential()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetCredentialRequestValidationError{
				field:  "Credential",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetCredentialRequestValidationError is the validation error returned by
// GetCredentialRequest.Validate if the designated constraints aren't met.
type GetCredentialRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCredentialRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCredentialRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCredentialRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCredentialRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCredentialRequestValidationError) ErrorName() string {
	return "GetCredentialRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetCredentialRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCredentialRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCredentialRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCredentialRequestValidationError{}

// Validate checks the field values on GetCredentialResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetCredentialResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCredential()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetCredentialResponseValidationError{
				field:  "Credential",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetCredentialResponseValidationError is the validation error returned by
// GetCredentialResponse.Validate if the designated constraints aren't met.
type GetCredentialResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCredentialResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCredentialResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCredentialResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCredentialResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCredentialResponseValidationError) ErrorName() string {
	return "GetCredentialResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetCredentialResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCredentialResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCredentialResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCredentialResponseValidationError{}