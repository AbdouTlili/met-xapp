// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"

	"github.com/AbdouTlili/met-xapp/pkg/northbound"
	"github.com/AbdouTlili/met-xapp/pkg/southbound"

	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger()

func main() {
	// {

	// 	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "172.21.16.115:9092"})
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	defer p.Close()

	// 	// Delivery report handler for produced messages
	// 	go func() {
	// 		for e := range p.Events() {
	// 			switch ev := e.(type) {
	// 			case *kafka.Message:
	// 				if ev.TopicPartition.Error != nil {
	// 					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
	// 				} else {
	// 					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
	// 				}
	// 			}
	// 		}
	// 	}()

	// 	// Produce messages to topic (asynchronously)
	// 	topic := "MyTopic"
	// 	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
	// 		p.Produce(&kafka.Message{
	// 			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	// 			Value:          []byte(word),
	// 		}, nil)
	// 	}
	// 	p.Flush(15 * 1000)
	// }

	// Wait for message deliveries before shutting down

	caPath := flag.String("caPath", "", "path to CA certificate")
	keyPath := flag.String("keyPath", "", "path to client private key")
	certPath := flag.String("certPath", "", "path to client certificate")
	e2tEndpoint := flag.String("e2tEndpoint", "onos-e2t:5150", "E2T service endpoint")
	ricActionID := flag.Int("ricActionID", 10, "RIC Action ID in E2 message")
	configPath := flag.String("configPath", "/etc/onos/config/config.json", "path to config.json file")
	grpcPort := flag.Int("grpcPort", 5150, "grpc Port number")
	smName := flag.String("smName", "e2sm_met", "Service model name in RAN function description")
	smVersion := flag.String("smVersion", "v1", "Service model version in RAN function description")

	log.Info(e2tEndpoint, ricActionID, configPath, grpcPort, smName, smVersion)

	flag.Parse()

	_, err := certs.HandleCertPaths(*caPath, *keyPath, *certPath, true)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Starting met-xapp")

	ready := make(chan bool)

	nbManager, err := northbound.NewBrokerClient("ran")

	if err != nil {
		log.Warn(err)
		panic(err)
	}

	subManager, err := southbound.NewManager(nbManager)

	if err != nil {
		log.Warn(err)
		panic(err)
	}

	go nbManager.Start()

	go subManager.Start()

	<-ready

}
