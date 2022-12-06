package customError

import (
	"github.com/pkg/errors"
)

type validationError struct {
	BadRequestError
}

func (ve *validationError) IsValidationError() bool {
	return true
}

type ValidationError interface {
	BadRequestError
	IsValidationError() bool
}

func IsValidationError(err error) bool {
	var validationError ValidationError

	if errors.As(err, &validationError) {
		return validationError.IsValidationError()
	}

	return false
}

func NewValidationError(message string, code int, details map[string]string) error {
	be := NewBadRequestError(message, code, details)
	customErr := GetCustomError(be)
	ve := &validationError{
		BadRequestError: customErr.(BadRequestError),
	}

	return ve
}

func NewValidationErrorWrap(err error, message string, code int, details map[string]string) error {
	be := NewBadRequestErrorWrap(err, message, code, details)
	customErr := GetCustomError(be)
	ve := &validationError{
		BadRequestError: customErr.(BadRequestError),
	}
	stackErr := errors.WithStack(ve)

	return stackErr
}
