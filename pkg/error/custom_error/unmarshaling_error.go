package customError

import (
	"github.com/pkg/errors"
)

type unMarshalingError struct {
	CustomError
}

func (e *unMarshalingError) IsUnMarshalingError() bool {
	return true
}

type UnMarshalingError interface {
	CustomError
	IsUnMarshalingError() bool
}

func IsUnMarshalingError(e error) bool {
	var unMarshalingError UnMarshalingError

	if errors.As(e, &unMarshalingError) {
		return unMarshalingError.IsUnMarshalingError()
	}

	return false
}

func NewUnMarshalingError(message string, code int, details map[string]string) error {
	e := &unMarshalingError{
		CustomError: NewCustomError(nil, code, message, details),
	}

	return e
}

func NewUnMarshalingErrorWrap(err error, message string, code int, details map[string]string) error {
	e := &unMarshalingError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(e)

	return stackErr
}
