// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: validate_token.proto

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
var _validate_token_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ValidateTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ValidateTokenRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetToken() == nil {
		return ValidateTokenRequestValidationError{
			field:  "Token",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetToken()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ValidateTokenRequestValidationError{
				field:  "Token",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ValidateTokenRequestValidationError is the validation error returned by
// ValidateTokenRequest.Validate if the designated constraints aren't met.
type ValidateTokenRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValidateTokenRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValidateTokenRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValidateTokenRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValidateTokenRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValidateTokenRequestValidationError) ErrorName() string {
	return "ValidateTokenRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ValidateTokenRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValidateTokenRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValidateTokenRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValidateTokenRequestValidationError{}

// Validate checks the field values on ValidateTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ValidateTokenResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetToken()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ValidateTokenResponseValidationError{
				field:  "Token",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ValidateTokenResponseValidationError is the validation error returned by
// ValidateTokenResponse.Validate if the designated constraints aren't met.
type ValidateTokenResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValidateTokenResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValidateTokenResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValidateTokenResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValidateTokenResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValidateTokenResponseValidationError) ErrorName() string {
	return "ValidateTokenResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ValidateTokenResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValidateTokenResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValidateTokenResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValidateTokenResponseValidationError{}
