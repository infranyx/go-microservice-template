package articleKafka

import (
	"github.com/segmentio/kafka-go"
)

type articleProducer struct {
	createWriter *kafka.Writer
}
