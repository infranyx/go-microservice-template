package customError

import (
	"github.com/pkg/errors"
)

func NewNotFoundError(message string, code int, details map[string]string) error {
	ne := &notFoundError{
		CustomError: NewCustomError(nil, code, message, details),
	}
	stackErr := errors.WithStack(ne)

	return stackErr
}

func NewNotFoundErrorWrap(err error, message string, code int, details map[string]string) error {
	ne := &notFoundError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(ne)

	return stackErr
}

type notFoundError struct {
	CustomError
}

type NotFoundError interface {
	CustomError
	IsNotFoundError() bool
}

func (n *notFoundError) IsNotFoundError() bool {
	return true
}

func IsNotFoundError(err error) bool {
	var notFoundError NotFoundError

	if errors.As(err, &notFoundError) {
		return notFoundError.IsNotFoundError()
	}

	return false
}
