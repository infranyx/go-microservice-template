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
	healthResult := hc.healthCheckUC.Check()

	httpStatus := http.StatusOK
	if !healthResult.Status {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, healthResult)
}
