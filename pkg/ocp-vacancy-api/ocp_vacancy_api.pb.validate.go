// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/ocp-vacancy-api/ocp_vacancy_api.proto

package ocp_vacancy_api

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

// Validate checks the field values on VacancyV1 with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *VacancyV1) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Link

	// no validation rules for Status

	return nil
}

// VacancyV1ValidationError is the validation error returned by
// VacancyV1.Validate if the designated constraints aren't met.
type VacancyV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VacancyV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VacancyV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VacancyV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VacancyV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VacancyV1ValidationError) ErrorName() string { return "VacancyV1ValidationError" }

// Error satisfies the builtin error interface
func (e VacancyV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVacancyV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VacancyV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VacancyV1ValidationError{}

// Validate checks the field values on CreateVacancyV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateVacancyV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Link

	// no validation rules for Status

	return nil
}

// CreateVacancyV1RequestValidationError is the validation error returned by
// CreateVacancyV1Request.Validate if the designated constraints aren't met.
type CreateVacancyV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateVacancyV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateVacancyV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateVacancyV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateVacancyV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateVacancyV1RequestValidationError) ErrorName() string {
	return "CreateVacancyV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateVacancyV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateVacancyV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateVacancyV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateVacancyV1RequestValidationError{}

// Validate checks the field values on CreateVacancyV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateVacancyV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// CreateVacancyV1ResponseValidationError is the validation error returned by
// CreateVacancyV1Response.Validate if the designated constraints aren't met.
type CreateVacancyV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateVacancyV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateVacancyV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateVacancyV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateVacancyV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateVacancyV1ResponseValidationError) ErrorName() string {
	return "CreateVacancyV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateVacancyV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateVacancyV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateVacancyV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateVacancyV1ResponseValidationError{}

// Validate checks the field values on UpdateVacancyV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateVacancyV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Link

	// no validation rules for Status

	return nil
}

// UpdateVacancyV1RequestValidationError is the validation error returned by
// UpdateVacancyV1Request.Validate if the designated constraints aren't met.
type UpdateVacancyV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateVacancyV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateVacancyV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateVacancyV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateVacancyV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateVacancyV1RequestValidationError) ErrorName() string {
	return "UpdateVacancyV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateVacancyV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateVacancyV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateVacancyV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateVacancyV1RequestValidationError{}

// Validate checks the field values on UpdateVacancyV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateVacancyV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetVacancy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateVacancyV1ResponseValidationError{
				field:  "Vacancy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateVacancyV1ResponseValidationError is the validation error returned by
// UpdateVacancyV1Response.Validate if the designated constraints aren't met.
type UpdateVacancyV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateVacancyV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateVacancyV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateVacancyV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateVacancyV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateVacancyV1ResponseValidationError) ErrorName() string {
	return "UpdateVacancyV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateVacancyV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateVacancyV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateVacancyV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateVacancyV1ResponseValidationError{}

// Validate checks the field values on DescribeVacancyV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeVacancyV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// DescribeVacancyV1RequestValidationError is the validation error returned by
// DescribeVacancyV1Request.Validate if the designated constraints aren't met.
type DescribeVacancyV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeVacancyV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeVacancyV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeVacancyV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeVacancyV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeVacancyV1RequestValidationError) ErrorName() string {
	return "DescribeVacancyV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeVacancyV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeVacancyV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeVacancyV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeVacancyV1RequestValidationError{}

// Validate checks the field values on DescribeVacancyV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeVacancyV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetVacancy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeVacancyV1ResponseValidationError{
				field:  "Vacancy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeVacancyV1ResponseValidationError is the validation error returned by
// DescribeVacancyV1Response.Validate if the designated constraints aren't met.
type DescribeVacancyV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeVacancyV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeVacancyV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeVacancyV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeVacancyV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeVacancyV1ResponseValidationError) ErrorName() string {
	return "DescribeVacancyV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeVacancyV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeVacancyV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeVacancyV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeVacancyV1ResponseValidationError{}

// Validate checks the field values on ListVacanciesV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListVacanciesV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Limit

	// no validation rules for Offset

	return nil
}

// ListVacanciesV1RequestValidationError is the validation error returned by
// ListVacanciesV1Request.Validate if the designated constraints aren't met.
type ListVacanciesV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListVacanciesV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListVacanciesV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListVacanciesV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListVacanciesV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListVacanciesV1RequestValidationError) ErrorName() string {
	return "ListVacanciesV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListVacanciesV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListVacanciesV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListVacanciesV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListVacanciesV1RequestValidationError{}

// Validate checks the field values on ListVacanciesV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListVacanciesV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetVacancies() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListVacanciesV1ResponseValidationError{
					field:  fmt.Sprintf("Vacancies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListVacanciesV1ResponseValidationError is the validation error returned by
// ListVacanciesV1Response.Validate if the designated constraints aren't met.
type ListVacanciesV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListVacanciesV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListVacanciesV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListVacanciesV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListVacanciesV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListVacanciesV1ResponseValidationError) ErrorName() string {
	return "ListVacanciesV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListVacanciesV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListVacanciesV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListVacanciesV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListVacanciesV1ResponseValidationError{}

// Validate checks the field values on RemoveVacancyV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveVacancyV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// RemoveVacancyV1RequestValidationError is the validation error returned by
// RemoveVacancyV1Request.Validate if the designated constraints aren't met.
type RemoveVacancyV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveVacancyV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveVacancyV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveVacancyV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveVacancyV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveVacancyV1RequestValidationError) ErrorName() string {
	return "RemoveVacancyV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveVacancyV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveVacancyV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveVacancyV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveVacancyV1RequestValidationError{}

// Validate checks the field values on RemoveVacancyV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveVacancyV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Found

	return nil
}

// RemoveVacancyV1ResponseValidationError is the validation error returned by
// RemoveVacancyV1Response.Validate if the designated constraints aren't met.
type RemoveVacancyV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveVacancyV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveVacancyV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveVacancyV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveVacancyV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveVacancyV1ResponseValidationError) ErrorName() string {
	return "RemoveVacancyV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveVacancyV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveVacancyV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveVacancyV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveVacancyV1ResponseValidationError{}

// Validate checks the field values on MultiCreateVacanciesV1Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MultiCreateVacanciesV1Request) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetVacancies() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MultiCreateVacanciesV1RequestValidationError{
					field:  fmt.Sprintf("Vacancies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MultiCreateVacanciesV1RequestValidationError is the validation error
// returned by MultiCreateVacanciesV1Request.Validate if the designated
// constraints aren't met.
type MultiCreateVacanciesV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateVacanciesV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateVacanciesV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateVacanciesV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateVacanciesV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateVacanciesV1RequestValidationError) ErrorName() string {
	return "MultiCreateVacanciesV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateVacanciesV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateVacanciesV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateVacanciesV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateVacanciesV1RequestValidationError{}

// Validate checks the field values on MultiCreateVacanciesV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MultiCreateVacanciesV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Added

	return nil
}

// MultiCreateVacanciesV1ResponseValidationError is the validation error
// returned by MultiCreateVacanciesV1Response.Validate if the designated
// constraints aren't met.
type MultiCreateVacanciesV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateVacanciesV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateVacanciesV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateVacanciesV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateVacanciesV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateVacanciesV1ResponseValidationError) ErrorName() string {
	return "MultiCreateVacanciesV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateVacanciesV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateVacanciesV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateVacanciesV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateVacanciesV1ResponseValidationError{}
