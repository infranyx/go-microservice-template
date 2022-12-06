package customError

import (
	"github.com/pkg/errors"
)

type conflictError struct {
	CustomError
}

func (ce *conflictError) IsConflictError() bool {
	return true
}

type ConflictError interface {
	CustomError
	IsConflictError() bool
}

func IsConflictError(err error) bool {
	var conflictError ConflictError

	if errors.As(err, &conflictError) {
		return conflictError.IsConflictError()
	}

	return false
}

func NewConflictError(message string, code int, details map[string]string) error {
	ce := &conflictError{
		CustomError: NewCustomError(nil, code, message, details),
	}

	return ce
}

func NewConflictErrorWrap(err error, message string, code int, details map[string]string) error {
	ce := &conflictError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(ce)

	return stackErr
}
