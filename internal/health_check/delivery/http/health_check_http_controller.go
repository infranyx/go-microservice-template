package healthCheckHttp

import (
	healthCheckDomain "github.com/infranyx/go-grpc-template/internal/health_check/domain"
	"github.com/labstack/echo/v4"
)

type healthCheckHttpController struct {
	healthCheckUC healthCheckDomain.HealthCheckUseCase
}

func NewHealthCheckHttpController(uc healthCheckDomain.HealthCheckUseCase) healthCheckDomain.HealthCheckHttpController {
	return &healthCheckHttpController{
		healthCheckUC: uc,
	}
}

func (hc healthCheckHttpController) Check(c echo.Context) error {
	return nil
}
