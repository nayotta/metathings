// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: model.proto

package evaluatord

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

	constant "github.com/nayotta/metathings/proto/constant/state"
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

	_ = constant.TaskState(0)

	_ = constant.TaskState(0)
)

// define the regex for a UUID once up-front
var _model_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Resource with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Resource) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Type

	return nil
}

// ResourceValidationError is the validation error returned by
// Resource.Validate if the designated constraints aren't met.
type ResourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResourceValidationError) ErrorName() string { return "ResourceValidationError" }

// Error satisfies the builtin error interface
func (e ResourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResourceValidationError{}

// Validate checks the field values on OpResource with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *OpResource) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpResourceValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetType()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpResourceValidationError{
				field:  "Type",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// OpResourceValidationError is the validation error returned by
// OpResource.Validate if the designated constraints aren't met.
type OpResourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpResourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpResourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpResourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpResourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpResourceValidationError) ErrorName() string { return "OpResourceValidationError" }

// Error satisfies the builtin error interface
func (e OpResourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpResource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpResourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpResourceValidationError{}

// Validate checks the field values on Evaluator with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Evaluator) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Alias

	// no validation rules for Description

	for idx, item := range m.GetSources() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EvaluatorValidationError{
					field:  fmt.Sprintf("Sources[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetOperator()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EvaluatorValidationError{
				field:  "Operator",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EvaluatorValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// EvaluatorValidationError is the validation error returned by
// Evaluator.Validate if the designated constraints aren't met.
type EvaluatorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EvaluatorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EvaluatorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EvaluatorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EvaluatorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EvaluatorValidationError) ErrorName() string { return "EvaluatorValidationError" }

// Error satisfies the builtin error interface
func (e EvaluatorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEvaluator.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EvaluatorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EvaluatorValidationError{}

// Validate checks the field values on OpEvaluator with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *OpEvaluator) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpEvaluatorValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAlias()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpEvaluatorValidationError{
				field:  "Alias",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDescription()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpEvaluatorValidationError{
				field:  "Description",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetSources() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OpEvaluatorValidationError{
					field:  fmt.Sprintf("Sources[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetOperator()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpEvaluatorValidationError{
				field:  "Operator",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpEvaluatorValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// OpEvaluatorValidationError is the validation error returned by
// OpEvaluator.Validate if the designated constraints aren't met.
type OpEvaluatorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpEvaluatorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpEvaluatorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpEvaluatorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpEvaluatorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpEvaluatorValidationError) ErrorName() string { return "OpEvaluatorValidationError" }

// Error satisfies the builtin error interface
func (e OpEvaluatorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpEvaluator.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpEvaluatorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpEvaluatorValidationError{}

// Validate checks the field values on Operator with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Operator) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Alias

	// no validation rules for Description

	// no validation rules for Driver

	switch m.Descriptor_.(type) {

	case *Operator_Lua:

		if v, ok := interface{}(m.GetLua()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OperatorValidationError{
					field:  "Lua",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// OperatorValidationError is the validation error returned by
// Operator.Validate if the designated constraints aren't met.
type OperatorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OperatorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OperatorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OperatorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OperatorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OperatorValidationError) ErrorName() string { return "OperatorValidationError" }

// Error satisfies the builtin error interface
func (e OperatorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOperator.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OperatorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OperatorValidationError{}

// Validate checks the field values on OpOperator with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *OpOperator) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpOperatorValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAlias()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpOperatorValidationError{
				field:  "Alias",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDescription()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpOperatorValidationError{
				field:  "Description",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDriver()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpOperatorValidationError{
				field:  "Driver",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.Descriptor_.(type) {

	case *OpOperator_Lua:

		if v, ok := interface{}(m.GetLua()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OpOperatorValidationError{
					field:  "Lua",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// OpOperatorValidationError is the validation error returned by
// OpOperator.Validate if the designated constraints aren't met.
type OpOperatorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpOperatorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpOperatorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpOperatorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpOperatorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpOperatorValidationError) ErrorName() string { return "OpOperatorValidationError" }

// Error satisfies the builtin error interface
func (e OpOperatorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpOperator.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpOperatorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpOperatorValidationError{}

// Validate checks the field values on LuaDescriptor with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LuaDescriptor) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	return nil
}

// LuaDescriptorValidationError is the validation error returned by
// LuaDescriptor.Validate if the designated constraints aren't met.
type LuaDescriptorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LuaDescriptorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LuaDescriptorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LuaDescriptorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LuaDescriptorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LuaDescriptorValidationError) ErrorName() string { return "LuaDescriptorValidationError" }

// Error satisfies the builtin error interface
func (e LuaDescriptorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLuaDescriptor.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LuaDescriptorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LuaDescriptorValidationError{}

// Validate checks the field values on OpLuaDescriptor with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *OpLuaDescriptor) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpLuaDescriptorValidationError{
				field:  "Code",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// OpLuaDescriptorValidationError is the validation error returned by
// OpLuaDescriptor.Validate if the designated constraints aren't met.
type OpLuaDescriptorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpLuaDescriptorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpLuaDescriptorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpLuaDescriptorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpLuaDescriptorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpLuaDescriptorValidationError) ErrorName() string { return "OpLuaDescriptorValidationError" }

// Error satisfies the builtin error interface
func (e OpLuaDescriptorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpLuaDescriptor.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpLuaDescriptorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpLuaDescriptorValidationError{}

// Validate checks the field values on OpTaskState with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *OpTaskState) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpTaskStateValidationError{
				field:  "At",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for State

	if v, ok := interface{}(m.GetTags()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpTaskStateValidationError{
				field:  "Tags",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// OpTaskStateValidationError is the validation error returned by
// OpTaskState.Validate if the designated constraints aren't met.
type OpTaskStateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpTaskStateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpTaskStateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpTaskStateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpTaskStateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpTaskStateValidationError) ErrorName() string { return "OpTaskStateValidationError" }

// Error satisfies the builtin error interface
func (e OpTaskStateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpTaskState.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpTaskStateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpTaskStateValidationError{}

// Validate checks the field values on TaskState with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *TaskState) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskStateValidationError{
				field:  "At",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for State

	if v, ok := interface{}(m.GetTags()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskStateValidationError{
				field:  "Tags",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// TaskStateValidationError is the validation error returned by
// TaskState.Validate if the designated constraints aren't met.
type TaskStateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskStateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskStateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskStateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskStateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskStateValidationError) ErrorName() string { return "TaskStateValidationError" }

// Error satisfies the builtin error interface
func (e TaskStateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTaskState.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskStateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskStateValidationError{}

// Validate checks the field values on OpTask with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *OpTask) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpTaskValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpTaskValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpTaskValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCurrentState()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpTaskValidationError{
				field:  "CurrentState",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpTaskValidationError{
				field:  "Source",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// OpTaskValidationError is the validation error returned by OpTask.Validate if
// the designated constraints aren't met.
type OpTaskValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpTaskValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpTaskValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpTaskValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpTaskValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpTaskValidationError) ErrorName() string { return "OpTaskValidationError" }

// Error satisfies the builtin error interface
func (e OpTaskValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpTask.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpTaskValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpTaskValidationError{}

// Validate checks the field values on Task with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Task) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCurrentState()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskValidationError{
				field:  "CurrentState",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskValidationError{
				field:  "Source",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetStates() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TaskValidationError{
					field:  fmt.Sprintf("States[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// TaskValidationError is the validation error returned by Task.Validate if the
// designated constraints aren't met.
type TaskValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskValidationError) ErrorName() string { return "TaskValidationError" }

// Error satisfies the builtin error interface
func (e TaskValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTask.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskValidationError{}

// Validate checks the field values on OpTimer with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *OpTimer) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpTimerValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAlias()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpTimerValidationError{
				field:  "Alias",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDescription()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpTimerValidationError{
				field:  "Description",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSchedule()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpTimerValidationError{
				field:  "Schedule",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTimezone()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpTimerValidationError{
				field:  "Timezone",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetEnabled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpTimerValidationError{
				field:  "Enabled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetConfigs() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OpTimerValidationError{
					field:  fmt.Sprintf("Configs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// OpTimerValidationError is the validation error returned by OpTimer.Validate
// if the designated constraints aren't met.
type OpTimerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpTimerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpTimerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpTimerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpTimerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpTimerValidationError) ErrorName() string { return "OpTimerValidationError" }

// Error satisfies the builtin error interface
func (e OpTimerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpTimer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpTimerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpTimerValidationError{}

// Validate checks the field values on Timer with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Timer) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Alias

	// no validation rules for Description

	// no validation rules for Schedule

	// no validation rules for Timezone

	// no validation rules for Enabled

	for idx, item := range m.GetConfigs() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TimerValidationError{
					field:  fmt.Sprintf("Configs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// TimerValidationError is the validation error returned by Timer.Validate if
// the designated constraints aren't met.
type TimerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TimerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TimerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TimerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TimerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TimerValidationError) ErrorName() string { return "TimerValidationError" }

// Error satisfies the builtin error interface
func (e TimerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTimer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TimerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TimerValidationError{}
