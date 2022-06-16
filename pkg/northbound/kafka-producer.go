package northbound

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

type BrokerClient struct {
	Topic    string
	Producer *kafka.Producer
}

var log = logging.GetLogger()

func NewBrokerClient(kafkaTopic string) (BrokerClient, error) {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		return BrokerClient{}, err
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	return BrokerClient{Topic: kafkaTopic, Producer: p}, nil
}

func (b *BrokerClient) Publish(messageBytes []byte) error {
	err := b.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &b.Topic, Partition: kafka.PartitionAny},
		Value:          messageBytes,
	}, nil)
	return err
}
