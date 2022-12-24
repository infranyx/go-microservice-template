package healthCheckGrpc

import (
	"context"
	healthCheckDomain "github.com/infranyx/go-grpc-template/internal/health_check/domain"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
)

type healthCheckGrpcController struct {
	healthCheckUseCase healthCheckDomain.HealthCheckUseCase
}

func NewHealthCheckGrpcController(uc healthCheckDomain.HealthCheckUseCase) healthCheckDomain.HealthCheckGrpcController {
	return &healthCheckGrpcController{
		healthCheckUseCase: uc,
	}
}

func (hc *healthCheckGrpcController) Check(ctx context.Context, request *grpcHealthV1.HealthCheckRequest) (*grpcHealthV1.HealthCheckResponse, error) {
	healthResult := hc.healthCheckUseCase.Check()

	grpcStatus := grpcHealthV1.HealthCheckResponse_SERVING

	if !healthResult.Status {
		grpcStatus = grpcHealthV1.HealthCheckResponse_NOT_SERVING
	}

	return &grpcHealthV1.HealthCheckResponse{
		Status: grpcStatus,
	}, nil
}

func (hc *healthCheckGrpcController) Watch(request *grpcHealthV1.HealthCheckRequest, server grpcHealthV1.Health_WatchServer) error {
	return nil
}
