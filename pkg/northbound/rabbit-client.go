package northbound

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"

	"github.com/streadway/amqp"
)

type BrokerClient struct {
	amqpServerUrl    string
	connectionBroker *amqp.Connection
	ChannelBroker    *amqp.Channel
}

var log = logging.GetLogger()

func NewBrokerClient(amqpServerUrl string) (BrokerClient, error) {

	connectionBroker, err := amqp.Dial(amqpServerUrl)
	if err != nil {
		return BrokerClient{}, err
	}

	ChannelBroker, err := connectionBroker.Channel()
	if err != nil {
		return BrokerClient{}, err
	}

	_, err = ChannelBroker.QueueDeclare(
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
		ChannelBroker:    ChannelBroker,
	}, nil
}

func (b *BrokerClient) Start() {

	log.Info("Northbound Broker Started")

	// ticker := time.NewTicker(10 * time.Second)

	done := make(chan bool)

	// go func() {
	// 	for {
	// 		select {
	// 		case <-done:
	// 			return
	// 		case <-ticker.C:
	// 			if err := b.Publish(); err != nil {
	// 				log.Warn(err)
	// 			}
	// 			log.Info("a Kpi is sent to the broker")

	// 		}
	// 	}
	// }()

	<-done

}

func (*BrokerClient) Stop() {

}

func (b *BrokerClient) Publish() error {
	kpiBytes := CreateDummyKpiMessage()
	message := amqp.Publishing{
		Body: kpiBytes,
	}

	// Attempt to publish a message to the queue

	if err := b.ChannelBroker.Publish("", "onos-queue1", false, false, message); err != nil {
		return err
	}

	return nil
}
