package customErrors

import (
	"github.com/pkg/errors"
)

func NewInternalServerError(message string, code int, details []ErrorDetail) error {
	ie := &internalServerError{
		CustomError: NewCustomError(nil, code, message, details),
	}
	// stackErr := errors.WithStack(br)

	return ie
}

func NewInternalServerErrorWrap(err error, message string, code int, details []ErrorDetail) error {
	ie := &internalServerError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(ie)

	return stackErr
}

type internalServerError struct {
	CustomError
}

type InternalServerError interface {
	CustomError
	IsInternalServerError() bool
}

func (i *internalServerError) IsInternalServerError() bool {
	return true
}

func IsInternalServerError(err error) bool {
	var internalErr InternalServerError

	if errors.As(err, &internalErr) {
		return internalErr.IsInternalServerError()
	}

	return false
}
