package customErrors

import (
	"errors"
	// "github.com/infranyx/go-grpc-template/shared/error/contracts"
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

func GetCustomError(err error) CustomError {
	var customErr CustomError
	if errors.As(err, &customErr) {
		return customErr
	}
	return nil
}

func (e *customError) IsCustomError() bool {
	return true
}

func IsCustomError(err error) bool {
	var customErr CustomError
	// _, ok := err.(CustomError)
	// if ok {
	// 	return true
	// }
	if errors.As(err, &customErr) {
		return customErr.IsCustomError()
	}
	return false
}
