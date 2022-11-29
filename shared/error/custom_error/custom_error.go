package customErrors

import (
	"errors"
	"fmt"
	"io"

	"github.com/infranyx/go-grpc-template/shared/error/contracts"
)

// https://github.com/pkg/errors/issues/75
type customError struct {
	internalCode int
	message      string
	err          error
	details      []ErrorDetail
}

type ErrorDetail struct {
	id   string
	desc string
}

type CustomError interface {
	error
	contracts.Wrapper
	contracts.Causer
	contracts.Formatter
	IsCustomError() bool
	Message() string
	Code() int
	Details() []ErrorDetail
}

func NewCustomError(err error, internalCode int, message string, details []ErrorDetail) CustomError {
	m := &customError{
		internalCode: internalCode,
		err:          err,
		message:      message,
		details:      details,
	}

	return m
}

func (e *customError) IsCustomError() bool {
	return true
}

func (e *customError) Error() string {
	if e.err != nil {
		return e.message + ": " + e.err.Error()
	}

	return e.message
}

func (e *customError) Message() string {
	return e.message
}

func (e *customError) Code() int {
	return e.internalCode
}

func (e *customError) Details() []ErrorDetail {
	return e.details
}

func (e *customError) Cause() error {
	return e.err
}

func (e *customError) Unwrap() error {
	return e.err
}

func (e *customError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v\n", e.Cause())
			io.WriteString(s, e.message)
			return
		}
		fallthrough
	case 's', 'q':
		io.WriteString(s, e.Error())
	}
}

func GetCustomError(err error) CustomError {
	if IsCustomError(err) {
		var internalErr CustomError
		errors.As(err, &internalErr)

		return internalErr
	}

	return nil
}

func IsCustomError(err error) bool {
	var customErr CustomError

	_, ok := err.(CustomError)
	if ok {
		return true
	}

	if errors.As(err, &customErr) {
		return customErr.IsCustomError()
	}

	return false
}
