package kafka

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
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
				run = false

			case e := <-p.Events():
				switch ev := e.(type) {
				case *kafka.Message:
					// Message delivery report
					m := ev
					if m.TopicPartition.Error != nil {
						continue
					}

				case kafka.Error:
					e := ev
					if e.IsFatal() {
						run = false
					}

				default:
					// Other events, such as rebalances, etc.
				}
			}
		}
	}()

	k.producer.producer = p
	k.producer.ready = true

	<-k.producer.termChan
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

func (q *Kafka) ProduceNotification(topic string, value []byte) error {
	if !q.producer.ready {
		return fmt.Errorf("producer is not ready")
	}

	q.producer.producer.ProduceChannel() <- &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: value,
		Key:   []byte(strconv.Itoa(rand.Intn(100000))),
	}

	q.producer.producer.Flush(15 * 1000)

	return nil
}
