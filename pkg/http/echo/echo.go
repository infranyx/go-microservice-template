package httpEcho

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/infranyx/go-grpc-template/pkg/constant"
	loggerConst "github.com/infranyx/go-grpc-template/pkg/constant/logger"
	echoErrorHandler "github.com/infranyx/go-grpc-template/pkg/http/echo/handlers/error_handler"
	"go.uber.org/zap"

	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type echoHttpServer struct {
	echo   *echo.Echo
	config *EchoHttpConfig
}

type EchoHttpServer interface {
	RunHttpServer(ctx context.Context, configEcho func(echoServer *echo.Echo)) error
	GracefulShutdown(ctx context.Context) error
	GetEchoInstance() *echo.Echo
	SetupDefaultMiddlewares()
	AddMiddlewares(middlewares ...echo.MiddlewareFunc)
	GetBasePath() string
}

type EchoHttpConfig struct {
	Port        int
	Development bool
	BasePath    string
}

func NewEchoHttpServer(config *EchoHttpConfig) *echoHttpServer {
	return &echoHttpServer{echo: echo.New(), config: config}
}

func (s *echoHttpServer) RunHttpServer(ctx context.Context, configEcho func(echo *echo.Echo)) error {
	s.echo.Server.ReadTimeout = constant.ReadTimeout
	s.echo.Server.WriteTimeout = constant.WriteTimeout
	s.echo.Server.MaxHeaderBytes = constant.MaxHeaderBytes

	if configEcho != nil {
		configEcho(s.echo)
	}

	go func() {
		for {
			<-ctx.Done()
			logger.Zap.Sugar().Infof("Http server is shutting down PORT: %d", s.config.Port)
			if err := s.GracefulShutdown(ctx); err != nil {
				logger.Zap.Sugar().Warnf("(Shutdown) err: {%v}", err)
			}
			return
		}
	}()

	logger.Zap.Sugar().Infof("[echoServer.RunHttpServer] Echo server is listening on: %d", s.config.Port)
	return s.echo.Start(fmt.Sprintf(":%d", s.config.Port))
}

func (s *echoHttpServer) AddMiddlewares(middlewares ...echo.MiddlewareFunc) {
	if len(middlewares) > 0 {
		s.echo.Use(middlewares...)
	}
}

func (s *echoHttpServer) GracefulShutdown(ctx context.Context) error {
	err := s.echo.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *echoHttpServer) SetupDefaultMiddlewares() {
	// handling internal echo middlewares logs with our log provider
	s.echo.HideBanner = false
	s.echo.Pre(middleware.RemoveTrailingSlash())
	s.echo.HTTPErrorHandler = echoErrorHandler.ErrorHandler

	s.echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogError:     false,
		LogMethod:    true,
		LogRequestID: true,
		LogURI:       true,
		LogStatus:    true,
		LogLatency:   true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			t := time.Now()
			logger.Zap.Info(
				"Incoming Request",
				zap.String(loggerConst.TYPE, loggerConst.HTTP),
				zap.String(loggerConst.METHOD, v.Method),
				zap.String(loggerConst.REQUEST_ID, v.RequestID),
				zap.String(loggerConst.URI, v.URI),
				zap.String(loggerConst.STATUS, http.StatusText(v.Status)),
				zap.Duration(loggerConst.LATENCY, v.Latency),
				zap.Time(loggerConst.TIME, t),
			)
			return nil
		},
	}))
	s.echo.Use(middleware.BodyLimit(constant.BodyLimit))
	s.echo.Use(middleware.RequestID())
	s.echo.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: constant.GzipLevel,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
}

func (s *echoHttpServer) GetEchoInstance() *echo.Echo {
	return s.echo
}

func (s *echoHttpServer) GetBasePath() string {
	return s.config.BasePath
}
