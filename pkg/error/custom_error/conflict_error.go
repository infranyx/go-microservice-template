package customError

import (
	"github.com/pkg/errors"
)

type conflictError struct {
	CustomError
}

func (e *conflictError) IsConflictError() bool {
	return true
}

type ConflictError interface {
	CustomError
	IsConflictError() bool
}

func IsConflictError(e error) bool {
	var conflictError ConflictError

	if errors.As(e, &conflictError) {
		return conflictError.IsConflictError()
	}

	return false
}

func NewConflictError(message string, code int, details map[string]string) error {
	e := &conflictError{
		CustomError: NewCustomError(nil, code, message, details),
	}

	return e
}

func NewConflictErrorWrap(err error, message string, code int, details map[string]string) error {
	e := &conflictError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(e)

	return stackErr
}
