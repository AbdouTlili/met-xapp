package northbound

import (
	"time"

	"github.com/onosproject/onos-lib-go/pkg/logging"

	"github.com/streadway/amqp"
)

type BrokerClient struct {
	amqpServerUrl    string
	connectionBroker *amqp.Connection
	channelBroker    *amqp.Channel
}

var log = logging.GetLogger()

func NewBrokerClient(amqpServerUrl string) (BrokerClient, error) {

	connectionBroker, err := amqp.Dial(amqpServerUrl)
	if err != nil {
		return BrokerClient{}, err
	}

	channelBroker, err := connectionBroker.Channel()
	if err != nil {
		return BrokerClient{}, err
	}

	_, err = channelBroker.QueueDeclare(
		"onos-queue1", // the queue name
		true,          // durable
		false,         // autodelete
		false,         // exclusive
		false,         // nowait
		nil,           // other args
	)

	if err != nil {
		return BrokerClient{}, err
	}

	return BrokerClient{
		amqpServerUrl:    amqpServerUrl,
		connectionBroker: connectionBroker,
		channelBroker:    channelBroker,
	}, nil
}

func (b *BrokerClient) Start() {

	log.Warn("Broker client Started")

	ticker := time.NewTicker(5 * time.Second)

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				if err := b.Publish(); err != nil {
					log.Warn(err)
				}

			}
		}
	}()

	<-done

}

func (*BrokerClient) Stop() {

}

func (b *BrokerClient) Publish() error {

	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("this is a test message"),
	}

	// Attempt to publish a message to the queue

	if err := b.channelBroker.Publish("", "onos-queue1", false, false, message); err != nil {
		return err
	}

	return nil
}
