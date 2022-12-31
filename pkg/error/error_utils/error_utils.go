package errorUtils

import (
	"context"
	"fmt"
	"strings"

	validator "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	errorContract "github.com/infranyx/go-microservice-template/pkg/error/contracts"
	customError "github.com/infranyx/go-microservice-template/pkg/error/custom_error"
	"github.com/infranyx/go-microservice-template/pkg/logger"
)

// CheckErrorMessages checks for specific messages contains in the error
func CheckErrorMessages(err error, messages ...string) bool {
	for _, message := range messages {
		if strings.Contains(strings.TrimSpace(strings.ToLower(err.Error())), strings.TrimSpace(strings.ToLower(message))) {
			return true
		}
	}
	return false
}

// RootStackTrace returns root stack trace with a string contains just stack trace levels for the given error
func RootStackTrace(err error) string {
	var stackStr string
	for {
		st, ok := err.(errorContract.StackTracer)
		if ok {
			stackStr = fmt.Sprintf("%+v\n", st.StackTrace())

			if !ok {
				break
			}
		}
		err = errors.Unwrap(err)
		if err == nil {
			break
		}
	}

	return stackStr
}

func ValidationErrorHandler(err error) (map[string]string, error) {
	var customErr validator.Errors
	if errors.As(err, &customErr) {
		details := make(map[string]string)
		for k, v := range customErr {
			details[k] = v.Error()
		}
		return details, nil
	}
	// TODO : get internal error from constant.
	return nil, customError.NewInternalServerErrorWrap(err, "internal error", 8585, nil)
}

type HandlerFunc func() error
type WrappedFunc func()

func HandlerErrorWrapper(ctx context.Context, f HandlerFunc) WrappedFunc { // must return without error
	return func() {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					logger.Zap.Sugar().Errorf("%v", r)
					return
				}
				logger.Zap.Error(err.Error(), zap.Error(err))
			}
		}()
		e := f()
		if e != nil {
			fmt.Println(e)
		}
	}
}
