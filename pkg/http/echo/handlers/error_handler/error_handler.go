package echoErrorHandler

import (
	"net/http"

	"github.com/getsentry/sentry-go"

	sentryEcho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	loggerConstant "github.com/infranyx/go-grpc-template/pkg/constant/logger"
	customErrors "github.com/infranyx/go-grpc-template/pkg/error/custom_error"
	errorCodes "github.com/infranyx/go-grpc-template/pkg/error/error_codes"
	errorUtils "github.com/infranyx/go-grpc-template/pkg/error/error_utils"
	httpError "github.com/infranyx/go-grpc-template/pkg/error/http"
	"github.com/infranyx/go-grpc-template/pkg/logger"
)

func ErrorHandler(err error, c echo.Context) {
	// default echo errors
	ehe, ok := err.(*echo.HTTPError)

	if ok {
		switch ehe.Code {
		case http.StatusNotFound:
			notFoundErrorCode := errorCodes.InternalErrorCodes.NotFoundError
			err = customErrors.NewNotFoundError(notFoundErrorCode.Msg, notFoundErrorCode.Code, nil)
		case http.StatusMethodNotAllowed:
			methodNotAllowedErrorCode := errorCodes.InternalErrorCodes.MethodNotAllowedError
			err = customErrors.NewMethodNotAllowedError(methodNotAllowedErrorCode.Msg, methodNotAllowedErrorCode.Code, nil)
		default:
			internalServerErrorCode := errorCodes.InternalErrorCodes.InternalServerError
			err = customErrors.NewInternalServerError(internalServerErrorCode.Msg, internalServerErrorCode.Code, nil)
		}
	}
	// system errors
	he := httpError.ParseError(err)

	if customErrors.IsInternalServerError(err) {
		hub := sentryEcho.GetHubFromContext(c)
		if hub != nil {
			hub.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetLevel(sentry.LevelError)
				scope.SetContext("HttpErr", map[string]interface{}{
					loggerConstant.TYPE:        loggerConstant.HTTP,
					loggerConstant.TITILE:      he.GetTitle(),
					loggerConstant.CODE:        he.GetCode(),
					loggerConstant.STATUS:      http.StatusText(he.GetStatus()),
					loggerConstant.TIME:        he.GetTimestamp(),
					loggerConstant.DETAILS:     he.GetDetails(),
					loggerConstant.STACK_TRACE: errorUtils.RootStackTrace(err),
				})
			})
			hub.CaptureException(err)
		}
	}

	if !c.Response().Committed {
		if _, err := he.WriteTo(c.Response()); err != nil {
			logger.Zap.Sugar().Error(`error while writing http error response: %v`, err)
		}
		logger.Zap.Error(
			err.Error(),
			zap.String(loggerConstant.TYPE, loggerConstant.HTTP),
			zap.String(loggerConstant.TITILE, he.GetTitle()),
			zap.Int(loggerConstant.CODE, he.GetCode()),
			zap.String(loggerConstant.STATUS, http.StatusText(he.GetStatus())),
			zap.Time(loggerConstant.TIME, he.GetTimestamp()),
			zap.Any(loggerConstant.DETAILS, he.GetDetails()),
			zap.String(loggerConstant.STACK_TRACE, errorUtils.RootStackTrace(err)),
		)
	}
}
