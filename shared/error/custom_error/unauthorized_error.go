package customErrors

import (
	"github.com/pkg/errors"
)

func NewUnAuthorizedError(message string, code int, details []ErrorDetail) error {
	ue := &unauthorizedError{
		CustomError: NewCustomError(nil, code, message, details),
	}
	stackErr := errors.WithStack(ue)

	return stackErr
}

func NewUnAuthorizedErrorWrap(err error, message string, code int, details []ErrorDetail) error {
	ue := &unauthorizedError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(ue)

	return stackErr
}

type unauthorizedError struct {
	CustomError
}

type UnauthorizedError interface {
	CustomError
	IsUnAuthorizedError() bool
}

func (u *unauthorizedError) IsUnAuthorizedError() bool {
	return true
}

func IsUnAuthorizedError(err error) bool {
	var unauthorizedError UnauthorizedError

	if errors.As(err, &unauthorizedError) {
		return unauthorizedError.IsUnAuthorizedError()
	}

	return false
}
