package customErrors

import (
	"github.com/pkg/errors"
)

func NewConflictError(message string, code int, details []ErrorDetail) error {
	ce := &conflictError{
		CustomError: NewCustomError(nil, code, message, details),
	}
	// stackErr := error.WithStack(ce)

	return ce
}

func NewConflictErrorWrap(err error, message string, code int, details []ErrorDetail) error {
	ce := &conflictError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(ce)

	return stackErr
}

type conflictError struct {
	CustomError
}

type ConflictError interface {
	CustomError
	IsConflictError() bool
}

func (c *conflictError) IsConflictError() bool {
	return true
}

func IsConflictError(err error) bool {
	var conflictError ConflictError

	if errors.As(err, &conflictError) {
		return conflictError.IsConflictError()
	}

	return false
}
