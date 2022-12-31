package customError

import (
	"github.com/pkg/errors"
)

type internalServerError struct {
	CustomError
}

func (e *internalServerError) IsInternalServerError() bool {
	return true
}

type InternalServerError interface {
	CustomError
	IsInternalServerError() bool
}

func IsInternalServerError(e error) bool {
	var internalError InternalServerError

	if errors.As(e, &internalError) {
		return internalError.IsInternalServerError()
	}

	return false
}

func NewInternalServerError(message string, code int, details map[string]string) error {
	e := &internalServerError{
		CustomError: NewCustomError(nil, code, message, details),
	}

	return e
}

func NewInternalServerErrorWrap(err error, message string, code int, details map[string]string) error {
	e := &internalServerError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(e)

	return stackErr
}
