package customError

import (
	"github.com/pkg/errors"
)

type unauthorizedError struct {
	CustomError
}

func (e *unauthorizedError) IsUnAuthorizedError() bool {
	return true
}

type UnauthorizedError interface {
	CustomError
	IsUnAuthorizedError() bool
}

func IsUnAuthorizedError(e error) bool {
	var unauthorizedError UnauthorizedError

	if errors.As(e, &unauthorizedError) {
		return unauthorizedError.IsUnAuthorizedError()
	}

	return false
}

func NewUnAuthorizedError(message string, code int, details map[string]string) error {
	e := &unauthorizedError{
		CustomError: NewCustomError(nil, code, message, details),
	}

	return e
}

func NewUnAuthorizedErrorWrap(err error, message string, code int, details map[string]string) error {
	e := &unauthorizedError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(e)

	return stackErr
}
