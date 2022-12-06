package customError

import (
	"github.com/pkg/errors"
)

type forbiddenError struct {
	CustomError
}

func (fe *forbiddenError) IsForbiddenError() bool {
	return true
}

type ForbiddenError interface {
	CustomError
	IsForbiddenError() bool
}

func IsForbiddenError(err error) bool {
	var forbiddenError ForbiddenError

	if errors.As(err, &forbiddenError) {
		return forbiddenError.IsForbiddenError()
	}

	return false
}

func NewForbiddenError(message string, code int, details map[string]string) error {
	ne := &forbiddenError{
		CustomError: NewCustomError(nil, code, message, details),
	}

	return ne
}

func NewForbiddenErrorWrap(err error, message string, code int, details map[string]string) error {
	ne := &forbiddenError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(ne)

	return stackErr
}
