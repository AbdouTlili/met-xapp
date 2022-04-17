// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"flag"

	"github.com/abdoutlili/met-xapp/pkg/rnib"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger()

func main() {
	caPath := flag.String("caPath", "", "path to CA certificate")
	keyPath := flag.String("keyPath", "", "path to client private key")
	certPath := flag.String("certPath", "", "path to client certificate")
	// e2tEndpoint := flag.String("e2tEndpoint", "onos-e2t:5150", "E2T service endpoint")
	// ricActionID := flag.Int("ricActionID", 10, "RIC Action ID in E2 message")
	// configPath := flag.String("configPath", "/etc/onos/config/config.json", "path to config.json file")
	// grpcPort := flag.Int("grpcPort", 5150, "grpc Port number")
	// smName := flag.String("smName", "met-xapp", "Service model name in RAN function description")
	// smVersion := flag.String("smVersion", "v1", "Service model version in RAN function description")

	flag.Parse()

	_, err := certs.HandleCertPaths(*caPath, *keyPath, *certPath, true)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Starting met-xapp")
	// cfg := manager.Config{
	// 	CAPath:      *caPath,
	// 	KeyPath:     *keyPath,
	// 	CertPath:    *certPath,
	// 	E2tEndpoint: *e2tEndpoint,
	// 	GRPCPort:    *grpcPort,
	// 	RicActionID: int32(*ricActionID),
	// 	ConfigPath:  *configPath,
	// 	SMName:      *smName,
	// 	SMVersion:   *smVersion,
	// }

	rnibClient, err := rnib.NewClient()
	ctx, _ := context.WithCancel(context.Background())
	watchE2Connections(ctx, rnibClient)

}

func watchE2Connections(ctx context.Context, rnibClient rnib.Client) error {
	ch := make(chan topoapi.Event)
	err := rnibClient.WatchE2Connections(ctx, ch)
	if err != nil {
		log.Warn(err)
		return err
	}

	// creates a new subscription whenever there is a new E2 node connected and supports KPM service model
	for topoEvent := range ch {
		log.Debugf("Received topo event: %v", topoEvent)

		if topoEvent.Type == topoapi.EventType_ADDED || topoEvent.Type == topoapi.EventType_NONE {
			relation := topoEvent.Object.Obj.(*topoapi.Object_Relation)
			e2NodeID := relation.Relation.TgtEntityID
			if !rnibClient.HasKPMRanFunction(ctx, e2NodeID, "1.3.6.1.4.1.53148.1.2.2.97") {
				continue
			}
		}

	}
	return nil
}
