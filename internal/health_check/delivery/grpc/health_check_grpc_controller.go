package healthCheckGrpc

import (
	"context"
	"fmt"
	healthCheckDomain "github.com/infranyx/go-grpc-template/internal/health_check/domain"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
)

type healthCheckGrpcController struct {
	healthCheckUseCase healthCheckDomain.HealthCheckUseCase
	kafkaHealthCheckUc healthCheckDomain.KafkaHealthCheckUseCase
}

func NewHealthCheckGrpcController(uc healthCheckDomain.HealthCheckUseCase, kuc healthCheckDomain.KafkaHealthCheckUseCase) healthCheckDomain.HealthCheckGrpcController {
	return &healthCheckGrpcController{
		healthCheckUseCase: uc,
		kafkaHealthCheckUc: kuc,
	}
}

func (hc *healthCheckGrpcController) Check(ctx context.Context, request *grpcHealthV1.HealthCheckRequest) (*grpcHealthV1.HealthCheckResponse, error) {
	var healthStatus bool

	switch request.Service {
	case "kafka":
		healthStatus = hc.kafkaHealthCheckUc.PingCheck()
		fmt.Println("kafka")
	default:
		healthStatus = hc.healthCheckUseCase.Check().Status
		fmt.Println("default")
	}

	grpcStatus := grpcHealthV1.HealthCheckResponse_SERVING

	if !healthStatus {
		grpcStatus = grpcHealthV1.HealthCheckResponse_NOT_SERVING
	}

	return &grpcHealthV1.HealthCheckResponse{
		Status: grpcStatus,
	}, nil
}

func (hc *healthCheckGrpcController) Watch(request *grpcHealthV1.HealthCheckRequest, server grpcHealthV1.Health_WatchServer) error {
	return nil
}
