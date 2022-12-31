package healthCheckHttp

import (
	"net/http"

	"github.com/labstack/echo/v4"

	healthCheckDomain "github.com/infranyx/go-microservice-template/internal/health_check/domain"
)

type controller struct {
	useCase healthCheckDomain.HealthCheckUseCase
}

func NewController(useCase healthCheckDomain.HealthCheckUseCase) healthCheckDomain.HttpController {
	return &controller{
		useCase: useCase,
	}
}

func (c controller) Check(eCtx echo.Context) error {
	healthResult := c.useCase.Check()

	httpStatus := http.StatusOK
	if !healthResult.Status {
		httpStatus = http.StatusInternalServerError
	}

	return eCtx.JSON(httpStatus, healthResult)
}
