package customError

import (
	"github.com/pkg/errors"
)

type badRequestError struct {
	CustomError
}

func (br *badRequestError) IsBadRequestError() bool {
	return true
}

type BadRequestError interface {
	CustomError
	IsBadRequestError() bool
}

func IsBadRequestError(err error) bool {
	var badRequestError BadRequestError

	if errors.As(err, &badRequestError) {
		return badRequestError.IsBadRequestError()
	}

	return false
}

func NewBadRequestError(message string, code int, details map[string]string) error {
	br := &badRequestError{
		CustomError: NewCustomError(nil, code, message, details),
	}

	return br
}

func NewBadRequestErrorWrap(err error, message string, code int, details map[string]string) error {
	br := &badRequestError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(br)

	return stackErr
}
