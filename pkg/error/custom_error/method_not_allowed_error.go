package customError

import (
	"github.com/pkg/errors"
)

type methodNotAllowedError struct {
	CustomError
}

func (na *methodNotAllowedError) IsMethodNotAllowedError() bool {
	return true
}

type MethodNotAllowedError interface {
	CustomError
	IsMethodNotAllowedError() bool
}

func IsMethodNotAllowedError(err error) bool {
	var methodNotAllowedError MethodNotAllowedError

	if errors.As(err, &methodNotAllowedError) {
		return methodNotAllowedError.IsMethodNotAllowedError()
	}

	return false
}

func NewMethodNotAllowedError(message string, code int, details map[string]string) error {
	ne := &methodNotAllowedError{
		CustomError: NewCustomError(nil, code, message, details),
	}

	return ne
}

func NewMethodNotAllowedWrap(err error, message string, code int, details map[string]string) error {
	ne := &methodNotAllowedError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(ne)

	return stackErr
}
