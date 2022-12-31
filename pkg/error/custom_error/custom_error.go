package customError

import "github.com/pkg/errors"

type customError struct {
	internalCode int
	message      string
	err          error
	details      map[string]string
}

func (ce *customError) Error() string {
	if ce.err != nil {
		return ce.message + ": " + ce.err.Error()
	}

	return ce.message
}

func (ce *customError) Message() string {
	return ce.message
}

func (ce *customError) Code() int {
	return ce.internalCode
}

func (ce *customError) Details() map[string]string {
	return ce.details
}

func (ce *customError) IsCustomError() bool {
	return true
}

type CustomError interface {
	error
	IsCustomError() bool
	Message() string
	Code() int
	Details() map[string]string
}

func IsCustomError(err error) bool {
	var customErr CustomError
	if errors.As(err, &customErr) {
		return customErr.IsCustomError()
	}
	return false
}

func NewCustomError(err error, internalCode int, message string, details map[string]string) CustomError {
	ce := &customError{
		internalCode: internalCode,
		err:          err,
		message:      message,
		details:      details,
	}

	return ce
}

func AsCustomError(err error) CustomError {
	var customErr CustomError
	if errors.As(err, &customErr) {
		return customErr
	}
	return nil
}
