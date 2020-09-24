// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: put_object_streaming.proto

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
var _put_object_streaming_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on PutObjectStreamingRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PutObjectStreamingRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PutObjectStreamingRequestValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.Request.(type) {

	case *PutObjectStreamingRequest_Metadata_:

		if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PutObjectStreamingRequestValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *PutObjectStreamingRequest_Chunks:

		if v, ok := interface{}(m.GetChunks()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PutObjectStreamingRequestValidationError{
					field:  "Chunks",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *PutObjectStreamingRequest_Ack_:

		if v, ok := interface{}(m.GetAck()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PutObjectStreamingRequestValidationError{
					field:  "Ack",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PutObjectStreamingRequestValidationError is the validation error returned by
// PutObjectStreamingRequest.Validate if the designated constraints aren't met.
type PutObjectStreamingRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PutObjectStreamingRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PutObjectStreamingRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PutObjectStreamingRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PutObjectStreamingRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PutObjectStreamingRequestValidationError) ErrorName() string {
	return "PutObjectStreamingRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PutObjectStreamingRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPutObjectStreamingRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PutObjectStreamingRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PutObjectStreamingRequestValidationError{}

// Validate checks the field values on PutObjectStreamingResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PutObjectStreamingResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	switch m.Response.(type) {

	case *PutObjectStreamingResponse_Chunks:

		if v, ok := interface{}(m.GetChunks()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PutObjectStreamingResponseValidationError{
					field:  "Chunks",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *PutObjectStreamingResponse_Ack_:

		if v, ok := interface{}(m.GetAck()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PutObjectStreamingResponseValidationError{
					field:  "Ack",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PutObjectStreamingResponseValidationError is the validation error returned
// by PutObjectStreamingResponse.Validate if the designated constraints aren't met.
type PutObjectStreamingResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PutObjectStreamingResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PutObjectStreamingResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PutObjectStreamingResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PutObjectStreamingResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PutObjectStreamingResponseValidationError) ErrorName() string {
	return "PutObjectStreamingResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PutObjectStreamingResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPutObjectStreamingResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PutObjectStreamingResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PutObjectStreamingResponseValidationError{}

// Validate checks the field values on PutObjectStreamingRequest_Ack with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PutObjectStreamingRequest_Ack) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// PutObjectStreamingRequest_AckValidationError is the validation error
// returned by PutObjectStreamingRequest_Ack.Validate if the designated
// constraints aren't met.
type PutObjectStreamingRequest_AckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PutObjectStreamingRequest_AckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PutObjectStreamingRequest_AckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PutObjectStreamingRequest_AckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PutObjectStreamingRequest_AckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PutObjectStreamingRequest_AckValidationError) ErrorName() string {
	return "PutObjectStreamingRequest_AckValidationError"
}

// Error satisfies the builtin error interface
func (e PutObjectStreamingRequest_AckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPutObjectStreamingRequest_Ack.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PutObjectStreamingRequest_AckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PutObjectStreamingRequest_AckValidationError{}

// Validate checks the field values on PutObjectStreamingRequest_Metadata with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *PutObjectStreamingRequest_Metadata) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetObject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PutObjectStreamingRequest_MetadataValidationError{
				field:  "Object",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSha1()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PutObjectStreamingRequest_MetadataValidationError{
				field:  "Sha1",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PutObjectStreamingRequest_MetadataValidationError is the validation error
// returned by PutObjectStreamingRequest_Metadata.Validate if the designated
// constraints aren't met.
type PutObjectStreamingRequest_MetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PutObjectStreamingRequest_MetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PutObjectStreamingRequest_MetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PutObjectStreamingRequest_MetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PutObjectStreamingRequest_MetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PutObjectStreamingRequest_MetadataValidationError) ErrorName() string {
	return "PutObjectStreamingRequest_MetadataValidationError"
}

// Error satisfies the builtin error interface
func (e PutObjectStreamingRequest_MetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPutObjectStreamingRequest_Metadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PutObjectStreamingRequest_MetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PutObjectStreamingRequest_MetadataValidationError{}

// Validate checks the field values on PutObjectStreamingResponse_Ack with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PutObjectStreamingResponse_Ack) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// PutObjectStreamingResponse_AckValidationError is the validation error
// returned by PutObjectStreamingResponse_Ack.Validate if the designated
// constraints aren't met.
type PutObjectStreamingResponse_AckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PutObjectStreamingResponse_AckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PutObjectStreamingResponse_AckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PutObjectStreamingResponse_AckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PutObjectStreamingResponse_AckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PutObjectStreamingResponse_AckValidationError) ErrorName() string {
	return "PutObjectStreamingResponse_AckValidationError"
}

// Error satisfies the builtin error interface
func (e PutObjectStreamingResponse_AckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPutObjectStreamingResponse_Ack.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PutObjectStreamingResponse_AckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PutObjectStreamingResponse_AckValidationError{}
