package kanban

import (
	"fmt"
	"net/http"
)

type HTTPError interface {
	Error() string
	StatusCode() int
}

func RenderError(w http.ResponseWriter, e error) {
	if h, ok := e.(HTTPError); ok {
		http.Error(w, h.Error(), h.StatusCode())
		return
	}

	http.Error(w, e.Error(), http.StatusInternalServerError)
}

type ErrInvalidFields struct {
	FieldErrors map[string][]error
}

func (e *ErrInvalidFields) Error() string {
	return "invalid fields"
}

func (e *ErrInvalidFields) Unwrap() []error {
	var errs []error
	for _, v := range e.FieldErrors {
		errs = append(errs, v...)
	}
	return errs
}

func (e *ErrInvalidFields) StatusCode() int {
	return http.StatusBadRequest
}

type ErrRenderFailed struct {
	wrapped error
}

func (e *ErrRenderFailed) Error() string {
	return fmt.Sprintf("failed to render: %v", e.wrapped)
}

func (e *ErrRenderFailed) Unwrap() error {
	return e.wrapped
}

func (e *ErrRenderFailed) StatusCode() int {
	return http.StatusInternalServerError
}

type ErrQueryFailed struct {
	wrapped error
}

func (e *ErrQueryFailed) Error() string {
	return fmt.Sprintf("failed to execute query: %v", e.wrapped)
}

func (e *ErrQueryFailed) Unwrap() error {
	return e.wrapped
}

func (e *ErrQueryFailed) StatusCode() int {
	return http.StatusInternalServerError
}

type ErrGenerateIDFailed struct {
	wrapped error
}

func (e *ErrGenerateIDFailed) Error() string {
	return fmt.Sprintf("failed to generate id: %v", e.wrapped)
}

func (e *ErrGenerateIDFailed) Unwrap() error {
	return e.wrapped
}

func (e *ErrGenerateIDFailed) StatusCode() int {
	return http.StatusInternalServerError
}

type ErrInvalidID struct {
	wrapped error
}

func (e *ErrInvalidID) Error() string {
	return fmt.Sprintf("invalid id: %v", e.wrapped)
}

func (e *ErrInvalidID) Unwrap() error {
	return e.wrapped
}

func (e *ErrInvalidID) StatusCode() int {
	return http.StatusInternalServerError
}

type ErrInvalidRequestData struct {
	wrapped error
}

func (e *ErrInvalidRequestData) Error() string {
	return fmt.Sprintf("invalid request: %v", e.wrapped)
}

func (e *ErrInvalidRequestData) Unwrap() error {
	return e.wrapped
}

func (e *ErrInvalidRequestData) StatusCode() int {
	return http.StatusInternalServerError
}

type ErrFieldRequired struct {
	Field string
}

func (e *ErrFieldRequired) Error() string {
	return fmt.Sprintf("%q is required", e.Field)
}

type ErrFieldTooLong struct {
	Field string
	Max   int
}

func (e *ErrFieldTooLong) Error() string {
	return fmt.Sprintf("%q length must be less than %d", e.Field, e.Max)
}
