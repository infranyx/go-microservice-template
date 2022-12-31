package customError

import (
	"github.com/pkg/errors"
)

type marshalingError struct {
	CustomError
}

func (e *marshalingError) IsMarshalingError() bool {
	return true
}

type MarshalingError interface {
	CustomError
	IsMarshalingError() bool
}

func IsMarshalingError(e error) bool {
	var marshalingError MarshalingError

	if errors.As(e, &marshalingError) {
		return marshalingError.IsMarshalingError()
	}

	return false
}

func NewMarshalingError(message string, code int, details map[string]string) error {
	e := &marshalingError{
		CustomError: NewCustomError(nil, code, message, details),
	}

	return e
}

func NewMarshalingErrorWrap(err error, message string, code int, details map[string]string) error {
	e := &marshalingError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(e)

	return stackErr
}
