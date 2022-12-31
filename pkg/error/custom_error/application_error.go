package customError

import (
	"github.com/pkg/errors"
)

type applicationError struct {
	CustomError
}

func (e *applicationError) IsApplicationError() bool {
	return true
}

type ApplicationError interface {
	CustomError
	IsApplicationError() bool
}

func IsApplicationError(e error) bool {
	var applicationError ApplicationError

	if errors.As(e, &applicationError) {
		return applicationError.IsApplicationError()
	}

	return false
}

func NewApplicationError(message string, code int, details map[string]string) error {
	e := &applicationError{
		CustomError: NewCustomError(nil, code, message, details),
	}

	return e
}

func NewApplicationErrorWrap(err error, message string, code int, details map[string]string) error {
	e := &applicationError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(e)

	return stackErr
}
