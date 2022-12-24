package kafkaHealthCheckUseCase

import (
	healthCheckDomain "github.com/infranyx/go-grpc-template/internal/health_check/domain"
	"github.com/infranyx/go-grpc-template/pkg/config"
	"github.com/segmentio/kafka-go"
)

type kafkaHealthCheck struct {
}

func NewKafkaHealthCheck() healthCheckDomain.KafkaHealthCheckUseCase {
	return &kafkaHealthCheck{}
}

func (kh *kafkaHealthCheck) PingCheck() bool {
	brokers := kafka.TCP(config.Conf.Kafka.ClientBrokers...)

	conn, err := kafka.Dial(brokers.Network(), brokers.String())
	if err != nil {
		return false
	}

	_ = conn.Close()

	return true
}
