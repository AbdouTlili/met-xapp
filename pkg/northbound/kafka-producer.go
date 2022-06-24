package northbound

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

type BrokerClient struct {
	Topic    string
	Producer *kafka.Producer
}

var log = logging.GetLogger()

func NewBrokerClient(kafkaTopic string) (BrokerClient, error) {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "172.21.16.114:30007"})
	if err != nil {
		return BrokerClient{}, err
	}

	// defer p.Close()

	return BrokerClient{Topic: kafkaTopic, Producer: p}, nil
}

func (b *BrokerClient) Start() {

	log.Info("Northbound Broker Started")

	forever := make(chan bool)
	// Delivery report handler for produced messages
	go func() {
		log.Infof("Delivery report handler for produced messages started \n")
		for e := range b.Producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Warnf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					log.Infof("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	<-forever
}

func (b *BrokerClient) Publish(messageBytes []byte) error {
	err := b.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &b.Topic, Partition: kafka.PartitionAny},
		Value:          messageBytes,
	}, nil)
	return err
}
