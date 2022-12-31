package customError

import (
	"github.com/pkg/errors"
)

type badRequestError struct {
	CustomError
}

func (e *badRequestError) IsBadRequestError() bool {
	return true
}

type BadRequestError interface {
	CustomError
	IsBadRequestError() bool
}

func IsBadRequestError(e error) bool {
	var badRequestError BadRequestError

	if errors.As(e, &badRequestError) {
		return badRequestError.IsBadRequestError()
	}

	return false
}

func NewBadRequestError(message string, code int, details map[string]string) error {
	e := &badRequestError{
		CustomError: NewCustomError(nil, code, message, details),
	}

	return e
}

func NewBadRequestErrorWrap(err error, message string, code int, details map[string]string) error {
	e := &badRequestError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(e)

	return stackErr
}
