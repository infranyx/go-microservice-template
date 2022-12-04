package customErrors

import (
	"github.com/pkg/errors"
)

func NewApplicationError(message string, code int, details []ErrorDetail) error {
	ae := &applicationError{
		CustomError: NewCustomError(nil, code, message, details),
	}
	// stackErr := error.WithStack(ae)

	return ae
}

func NewApplicationErrorWrap(err error, message string, code int, details []ErrorDetail) error {
	ae := &applicationError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(ae)

	return stackErr
}

type applicationError struct {
	CustomError
}

type ApplicationError interface {
	CustomError
	IsApplicationError() bool
}

func (a *applicationError) IsApplicationError() bool {
	return true
}

func IsApplicationError(err error) bool {
	var applicationError ApplicationError

	if errors.As(err, &applicationError) {
		return applicationError.IsApplicationError()
	}

	return false
}
