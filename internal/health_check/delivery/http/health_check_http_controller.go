package healthCheckHttp

import (
	healthCheckDomain "github.com/infranyx/go-grpc-template/internal/health_check/domain"
	"github.com/labstack/echo/v4"
	"net/http"
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
	healthResult, _ := hc.healthCheckUC.Check(c.Request().Context())
	return c.JSON(http.StatusOK, healthResult)
}
