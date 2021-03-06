// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: issue_token_by_password.proto

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
var _issue_token_by_password_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on IssueTokenByPasswordRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *IssueTokenByPasswordRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetEntity() == nil {
		return IssueTokenByPasswordRequestValidationError{
			field:  "Entity",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetEntity()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IssueTokenByPasswordRequestValidationError{
				field:  "Entity",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// IssueTokenByPasswordRequestValidationError is the validation error returned
// by IssueTokenByPasswordRequest.Validate if the designated constraints
// aren't met.
type IssueTokenByPasswordRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IssueTokenByPasswordRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IssueTokenByPasswordRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IssueTokenByPasswordRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IssueTokenByPasswordRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IssueTokenByPasswordRequestValidationError) ErrorName() string {
	return "IssueTokenByPasswordRequestValidationError"
}

// Error satisfies the builtin error interface
func (e IssueTokenByPasswordRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIssueTokenByPasswordRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IssueTokenByPasswordRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IssueTokenByPasswordRequestValidationError{}

// Validate checks the field values on IssueTokenByPasswordResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *IssueTokenByPasswordResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetToken()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IssueTokenByPasswordResponseValidationError{
				field:  "Token",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// IssueTokenByPasswordResponseValidationError is the validation error returned
// by IssueTokenByPasswordResponse.Validate if the designated constraints
// aren't met.
type IssueTokenByPasswordResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IssueTokenByPasswordResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IssueTokenByPasswordResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IssueTokenByPasswordResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IssueTokenByPasswordResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IssueTokenByPasswordResponseValidationError) ErrorName() string {
	return "IssueTokenByPasswordResponseValidationError"
}

// Error satisfies the builtin error interface
func (e IssueTokenByPasswordResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIssueTokenByPasswordResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IssueTokenByPasswordResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IssueTokenByPasswordResponseValidationError{}
