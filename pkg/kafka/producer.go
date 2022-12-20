package kafka

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	// "github.com/defipod/mochi/pkg/logger"
)

func (k *Kafka) RunProducer() error {
	config := kafka.ConfigMap{
		"bootstrap.servers":  k.brokers,
		"enable.idempotence": true,
		"acks":               "all",
	}

	p, err := kafka.NewProducer(&config)
	if err != nil {
		return fmt.Errorf("failed to create producer: %s", err)
	}

	go func() {
		run := true
		for run {
			select {
			case <-k.producer.termChan:
				// logger.Logger.Info("producer received termination signal")
				run = false

			case e := <-p.Events():
				switch ev := e.(type) {
				case *kafka.Message:
					// Message delivery report
					m := ev
					if m.TopicPartition.Error != nil {
						// logger.Logger.Error(m.TopicPartition.Error, "failed to deliver message")
						continue
					}

				case kafka.Error:
					e := ev
					if e.IsFatal() {
						// logger.Logger.Error(e, "fatal error event")
						run = false
					}
					// logger.Logger.Error(e, "error event")

				default:
					// Other events, such as rebalances, etc.
				}
			}
		}
	}()

	k.producer.producer = p
	k.producer.ready = true

	// logger.L.Info("producer is ready")
	<-k.producer.termChan
	// logger.L.Info("closing producer")
	p.Close()

	fatalErr := p.GetFatalError()
	if fatalErr != nil {
		return fmt.Errorf("fatal error: %s", fatalErr)
	}

	return nil
}

func (k *Kafka) ProducerReady() bool {
	return k.producer.ready
}

func (k *Kafka) CloseProducer() error {
	if k.producer == nil {
		return nil
	}

	k.producer.termChan <- true
	return nil
}

func (k *Kafka) Produce(topic, key string, value []byte) error {
	// wait for producer to be ready
	if !k.producer.ready {
		// logger.L.Info("producer is not ready, waiting")
		for {
			if k.producer.ready {
				break
			}
			time.Sleep(time.Millisecond * 100)
		}
	}

	// * SEND MESSAGE *
	k.producer.producer.ProduceChannel() <- &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: value,
		Key:   []byte(key),
	}

	k.producer.producer.Flush(15 * 1000)

	return nil
}
