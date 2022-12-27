package customError

import (
	"github.com/pkg/errors"
)

type methodNotAllowedError struct {
	CustomError
}

func (e *methodNotAllowedError) IsMethodNotAllowedError() bool {
	return true
}

type MethodNotAllowedError interface {
	CustomError
	IsMethodNotAllowedError() bool
}

func IsMethodNotAllowedError(e error) bool {
	var methodNotAllowedError MethodNotAllowedError

	if errors.As(e, &methodNotAllowedError) {
		return methodNotAllowedError.IsMethodNotAllowedError()
	}

	return false
}

func NewMethodNotAllowedError(message string, code int, details map[string]string) error {
	e := &methodNotAllowedError{
		CustomError: NewCustomError(nil, code, message, details),
	}

	return e
}

func NewMethodNotAllowedWrap(err error, message string, code int, details map[string]string) error {
	e := &methodNotAllowedError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(e)

	return stackErr
}
