package customError

import (
	"github.com/pkg/errors"
)

func NewMarshalingError(message string, code int, details map[string]string) error {
	me := &marshalingError{
		CustomError: NewCustomError(nil, code, message, details),
	}
	// stackErr := error.WithStack(ne)

	return me
}

func NewMarshalingErrorWrap(err error, message string, code int, details map[string]string) error {
	me := &marshalingError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(me)

	return stackErr
}

type marshalingError struct {
	CustomError
}

type MarshalingError interface {
	CustomError
	IsMarshalingError() bool
}

func (m *marshalingError) IsMarshalingError() bool {
	return true
}

func IsMarshalingError(err error) bool {
	var me MarshalingError

	if errors.As(err, &me) {
		return me.IsMarshalingError()
	}

	return false
}
