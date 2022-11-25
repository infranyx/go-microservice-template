package customErrors

import (
	"github.com/pkg/errors"
)

func NewUnAuthorizedError(message string, code int) error {
	ue := &unauthorizedError{
		CustomError: NewCustomError(nil, code, message),
	}
	stackErr := errors.WithStack(ue)

	return stackErr
}

func NewUnAuthorizedErrorWrap(err error, code int, message string) error {
	ue := &unauthorizedError{
		CustomError: NewCustomError(err, code, message),
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
	//us, ok := grpc_errors.Cause(err).(UnauthorizedError)
	if errors.As(err, &unauthorizedError) {
		return unauthorizedError.IsUnAuthorizedError()
	}

	return false
}
