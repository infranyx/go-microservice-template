package customError

import (
	"github.com/pkg/errors"
)

type internalServerError struct {
	CustomError
}

func (ie *internalServerError) IsInternalServerError() bool {
	return true
}

type InternalServerError interface {
	CustomError
	IsInternalServerError() bool
}

func IsInternalServerError(err error) bool {
	var internalErr InternalServerError

	if errors.As(err, &internalErr) {
		return internalErr.IsInternalServerError()
	}

	return false
}

func NewInternalServerError(message string, code int, details map[string]string) error {
	ie := &internalServerError{
		CustomError: NewCustomError(nil, code, message, details),
	}

	return ie
}

func NewInternalServerErrorWrap(err error, message string, code int, details map[string]string) error {
	ie := &internalServerError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(ie)

	return stackErr
}
