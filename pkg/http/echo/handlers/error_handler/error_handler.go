package echoErrorHandler

import (
	"github.com/getsentry/sentry-go"
	"net/http"

	sentryecho "github.com/getsentry/sentry-go/echo"
	errorConst "github.com/infranyx/go-grpc-template/pkg/constant/error"
	loggerConst "github.com/infranyx/go-grpc-template/pkg/constant/logger"
	customErrors "github.com/infranyx/go-grpc-template/pkg/error/custom_error"
	errorUtils "github.com/infranyx/go-grpc-template/pkg/error/error_utils"
	httpError "github.com/infranyx/go-grpc-template/pkg/error/http"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func ErrorHandler(err error, c echo.Context) {
	// default echo errors
	ehe, ok := err.(*echo.HTTPError)
	if ok {
		switch ehe.Code {
		case http.StatusNotFound:
			err = customErrors.NewNotFoundError(errorConst.ErrInfo.NotFoundErr.Msg, errorConst.ErrInfo.NotFoundErr.Code, nil)
		default:
			err = customErrors.NewInternalServerError(errorConst.ErrInfo.InternalServerErr.Msg, errorConst.ErrInfo.InternalServerErr.Code, nil)
		}
	}
	// system errors
	he := httpError.ParseError(err)

	if customErrors.IsInternalServerError(err) {
		hub := sentryecho.GetHubFromContext(c)
		if hub != nil {
			hub.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetLevel(sentry.LevelError)
				scope.SetContext("grpcErr", map[string]interface{}{
					loggerConst.TYPE:        loggerConst.HTTP,
					loggerConst.TITILE:      he.GetTitle(),
					loggerConst.CODE:        he.GetCode(),
					loggerConst.STATUS:      http.StatusText(he.GetStatus()),
					loggerConst.TIME:        he.GetTimestamp(),
					loggerConst.DETAILS:     he.GetDetails(),
					loggerConst.STACK_TRACE: errorUtils.RootStackTrace(err),
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
			zap.String(loggerConst.TYPE, loggerConst.HTTP),
			zap.String(loggerConst.TITILE, he.GetTitle()),
			zap.Int(loggerConst.CODE, he.GetCode()),
			zap.String(loggerConst.STATUS, http.StatusText(he.GetStatus())),
			zap.Time(loggerConst.TIME, he.GetTimestamp()),
			zap.Any(loggerConst.DETAILS, he.GetDetails()),
			zap.String(loggerConst.STACK_TRACE, errorUtils.RootStackTrace(err)),
		)
	}
}
