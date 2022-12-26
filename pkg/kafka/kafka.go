package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/riferrei/srclient"
)

type Kafka struct {
	schemaRegistryClient *srclient.SchemaRegistryClient
	brokers              string
	producer             *producer
}

type producer struct {
	producer *kafka.Producer
	ready    bool
	termChan chan bool
}

func New(brokers string) *Kafka {
	kafka := &Kafka{
		brokers:  brokers,
		producer: &producer{},
	}

	return kafka
}
