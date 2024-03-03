package store

import "fmt"

type ErrQueryFailed struct {
	wrapped error
}

func NewErrQueryFailed(err error) *ErrQueryFailed {
	return &ErrQueryFailed{
		wrapped: err,
	}
}

func (e *ErrQueryFailed) Error() string {
	return fmt.Sprintf("failed to execute query: %v", e.wrapped)
}

func (e *ErrQueryFailed) Unwrap() error {
	return e.wrapped
}
