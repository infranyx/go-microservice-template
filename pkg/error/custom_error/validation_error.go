package customError

import (
	"github.com/pkg/errors"
)

type validationError struct {
	BadRequestError
}

func (e *validationError) IsValidationError() bool {
	return true
}

type ValidationError interface {
	BadRequestError
	IsValidationError() bool
}

func IsValidationError(e error) bool {
	var validationError ValidationError

	if errors.As(e, &validationError) {
		return validationError.IsValidationError()
	}

	return false
}

func NewValidationError(message string, code int, details map[string]string) error {
	e := NewBadRequestError(message, code, details)
	ve := &validationError{
		BadRequestError: &badRequestError{
			CustomError: AsCustomError(e),
		},
	}

	return ve
}

func NewValidationErrorWrap(err error, message string, code int, details map[string]string) error {
	e := NewBadRequestErrorWrap(err, message, code, details)
	ve := &validationError{
		BadRequestError: &badRequestError{
			CustomError: AsCustomError(e),
		},
	}

	stackErr := errors.WithStack(ve)

	return stackErr
}
