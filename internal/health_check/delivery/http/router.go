package healthCheckHttp

import (
	healthCheckDomain "github.com/infranyx/go-grpc-template/internal/health_check/domain"
	"github.com/labstack/echo/v4"
)

type Router struct {
	healthCheckCtrl healthCheckDomain.HealthCheckHttpController
}

func NewHealthCheckAPI(hc healthCheckDomain.HealthCheckHttpController) *Router {
	return &Router{
		healthCheckCtrl: hc,
	}
}

func (r *Router) Register(e *echo.Group) {
	e.GET("/health", r.healthCheckCtrl.Check)
}
