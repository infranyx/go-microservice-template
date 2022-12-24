package healthCheckGrpc

import (
	"context"
	healthCheckDomain "github.com/infranyx/go-grpc-template/internal/health_check/domain"
	"google.golang.org/grpc/codes"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

type healthCheckGrpcController struct {
	healthCheckUseCase   healthCheckDomain.HealthCheckUseCase
	pgHealthCheckUc      healthCheckDomain.PgHealthCheckUseCase
	kafkaHealthCheckUc   healthCheckDomain.KafkaHealthCheckUseCase
	tempDirHealthCheckUc healthCheckDomain.TmpDirHealthCheckUseCase
}

func NewHealthCheckGrpcController(uc healthCheckDomain.HealthCheckUseCase, pguc healthCheckDomain.PgHealthCheckUseCase, kuc healthCheckDomain.KafkaHealthCheckUseCase, tmpDirUc healthCheckDomain.TmpDirHealthCheckUseCase) healthCheckDomain.HealthCheckGrpcController {
	return &healthCheckGrpcController{
		healthCheckUseCase:   uc,
		pgHealthCheckUc:      pguc,
		kafkaHealthCheckUc:   kuc,
		tempDirHealthCheckUc: tmpDirUc,
	}
}

func (hc *healthCheckGrpcController) Check(ctx context.Context, request *grpcHealthV1.HealthCheckRequest) (*grpcHealthV1.HealthCheckResponse, error) {
	var healthStatus bool

	switch request.Service {
	case "":
		healthStatus = hc.healthCheckUseCase.Check().Status
	case "kafka":
		healthStatus = hc.kafkaHealthCheckUc.PingCheck()
	case "postgres":
		healthStatus = hc.pgHealthCheckUc.PingCheck()
	case "writable-tmp-dir":
		healthStatus = hc.tempDirHealthCheckUc.PingCheck()
	default:
		return &grpcHealthV1.HealthCheckResponse{
			Status: grpcHealthV1.HealthCheckResponse_UNKNOWN,
		}, nil
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
	return status.Error(codes.Unimplemented, "unimplemented")
}
