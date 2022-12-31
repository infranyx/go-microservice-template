package customError

import (
	"github.com/pkg/errors"
)

type forbiddenError struct {
	CustomError
}

func (e *forbiddenError) IsForbiddenError() bool {
	return true
}

type ForbiddenError interface {
	CustomError
	IsForbiddenError() bool
}

func IsForbiddenError(e error) bool {
	var forbiddenError ForbiddenError

	if errors.As(e, &forbiddenError) {
		return forbiddenError.IsForbiddenError()
	}

	return false
}

func NewForbiddenError(message string, code int, details map[string]string) error {
	e := &forbiddenError{
		CustomError: NewCustomError(nil, code, message, details),
	}

	return e
}

func NewForbiddenErrorWrap(err error, message string, code int, details map[string]string) error {
	e := &forbiddenError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(e)

	return stackErr
}
