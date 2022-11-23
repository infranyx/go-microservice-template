package customErrors

import (
	"github.com/pkg/errors"
)

func NewConflictError(message string, code int) error {
	ce := &conflictError{
		CustomError: NewCustomError(nil, code, message),
	}
	stackErr := errors.WithStack(ce)

	return stackErr
}

func NewConflictErrorWrap(err error, code int, message string) error {
	ce := &conflictError{
		CustomError: NewCustomError(err, code, message),
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
	//us, ok := grpc_errors.Cause(err).(ConflictError)
	if errors.As(err, &conflictError) {
		return conflictError.IsConflictError()
	}

	return false
}
