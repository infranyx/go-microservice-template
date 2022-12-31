package httpError

import (
	"net/http"
	"time"

	errorConstant "github.com/infranyx/go-microservice-template/pkg/constant/error"
)

func NewHttpValidationError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errorConstant.ErrValidationFailedTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusBadRequest,
		Timestamp: time.Now(),
	}
}

func NewHttpConflictError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errorConstant.ErrConflictTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusConflict,
		Timestamp: time.Now(),
	}
}

func NewHttpBadRequestError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errorConstant.ErrBadRequestTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusBadRequest,
		Timestamp: time.Now(),
	}
}

func NewHttpNotFoundError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errorConstant.ErrNotFoundTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusNotFound,
		Timestamp: time.Now(),
	}
}

func NewHttpUnAuthorizedError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errorConstant.ErrUnauthorizedTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusUnauthorized,
		Timestamp: time.Now(),
	}
}

func NewHttpForbiddenError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errorConstant.ErrForbiddenTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusForbidden,
		Timestamp: time.Now(),
	}
}

func NewHttpInternalServerError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errorConstant.ErrInternalServerErrorTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusInternalServerError,
		Timestamp: time.Now(),
	}
}

func NewHttpDomainError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errorConstant.ErrDomainTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusBadRequest,
		Timestamp: time.Now(),
	}
}

func NewHttpApplicationError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errorConstant.ErrApplicationTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusInternalServerError,
		Timestamp: time.Now(),
	}
}

func NewHttpApiError(code int, message string, details map[string]string) HttpErr {
	return &httpErr{
		Title:     errorConstant.ErrApiTitle,
		Code:      code,
		Msg:       message,
		Details:   details,
		Status:    http.StatusInternalServerError,
		Timestamp: time.Now(),
	}
}
