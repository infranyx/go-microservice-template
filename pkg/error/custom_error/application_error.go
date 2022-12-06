package customError

import (
	"github.com/pkg/errors"
)

type applicationError struct {
	CustomError
}

func (ae *applicationError) IsApplicationError() bool {
	return true
}

type ApplicationError interface {
	CustomError
	IsApplicationError() bool
}

func IsApplicationError(err error) bool {
	var applicationError ApplicationError

	if errors.As(err, &applicationError) {
		return applicationError.IsApplicationError()
	}

	return false
}

func NewApplicationError(message string, code int, details map[string]string) error {
	ae := &applicationError{
		CustomError: NewCustomError(nil, code, message, details),
	}

	return ae
}

func NewApplicationErrorWrap(err error, message string, code int, details map[string]string) error {
	ae := &applicationError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(ae)

	return stackErr
}
