package customError

import (
	"github.com/pkg/errors"
)

type marshalingError struct {
	CustomError
}

func (me *marshalingError) IsMarshalingError() bool {
	return true
}

type MarshalingError interface {
	CustomError
	IsMarshalingError() bool
}

func IsMarshalingError(err error) bool {
	var me MarshalingError

	if errors.As(err, &me) {
		return me.IsMarshalingError()
	}

	return false
}

func NewMarshalingError(message string, code int, details map[string]string) error {
	me := &marshalingError{
		CustomError: NewCustomError(nil, code, message, details),
	}

	return me
}

func NewMarshalingErrorWrap(err error, message string, code int, details map[string]string) error {
	me := &marshalingError{
		CustomError: NewCustomError(err, code, message, details),
	}
	stackErr := errors.WithStack(me)

	return stackErr
}
