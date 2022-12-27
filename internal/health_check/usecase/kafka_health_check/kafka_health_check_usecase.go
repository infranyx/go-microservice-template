package kafkaHealthCheckUseCase

import (
	healthCheckDomain "github.com/infranyx/go-grpc-template/internal/health_check/domain"
	"github.com/infranyx/go-grpc-template/pkg/config"
	"github.com/segmentio/kafka-go"
)

type useCase struct{}

func NewUseCase() healthCheckDomain.KafkaHealthCheckUseCase {
	return &useCase{}
}

func (uc *useCase) Check() bool {
	brokers := kafka.TCP(config.BaseConfig.Kafka.ClientBrokers...)

	conn, err := kafka.Dial(brokers.Network(), brokers.String())
	if err != nil {
		return false
	}

	_ = conn.Close()

	return true
}
