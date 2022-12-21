package healthCheckUseCase

import (
	"context"

	healthCheckDomain "github.com/infranyx/go-grpc-template/internal/health_check/domain"
)

type healthCheckUseCase struct{}

func NewHealthCheckUseCase() healthCheckDomain.HealthCheckUseCase {
	return &healthCheckUseCase{}
}

func (hu *healthCheckUseCase) Check(ctx context.Context) (*healthCheckDomain.HealthCheckResult, error) {
	return nil, nil
}
