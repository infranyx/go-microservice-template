package customErrors

import (
	"github.com/pkg/errors"
)

func NewForbiddenError(message string, code int) error {
	ne := &forbiddenError{
		CustomError: NewCustomError(nil, code, message),
	}
	stackErr := errors.WithStack(ne)

	return stackErr
}

func NewForbiddenErrorWrap(err error, code int, message string) error {
	ne := &forbiddenError{
		CustomError: NewCustomError(err, code, message),
	}
	stackErr := errors.WithStack(ne)

	return stackErr
}

type forbiddenError struct {
	CustomError
}

type ForbiddenError interface {
	CustomError
	IsForbiddenError() bool
}

func (f *forbiddenError) IsForbiddenError() bool {
	return true
}

func IsForbiddenError(err error) bool {
	var forbiddenError ForbiddenError
	//us, ok := grpc_errors.Cause(err).(ForbiddenError)
	if errors.As(err, &forbiddenError) {
		return forbiddenError.IsForbiddenError()
	}

	return false
}
