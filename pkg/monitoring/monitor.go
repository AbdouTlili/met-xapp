// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package monitoring

import (
	"context"

	"github.com/AbdouTlili/met-xapp/pkg/broker"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"google.golang.org/protobuf/proto"

	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"

	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger()

// Monitor indication monitor
type Monitor struct {
	streamReader broker.StreamReader
}

func NewMonitor(streamReader broker.StreamReader) *Monitor {

	return &Monitor{
		streamReader: streamReader,
	}
}

// Start start monitoring of indication messages for a given subscription ID
func (m *Monitor) Start(ctx context.Context) error {
	log.Info("Monitor started")
	errCh := make(chan error)
	go func() {
		for {
			indMsg, err := m.streamReader.Recv(ctx)
			if err != nil {
				errCh <- err
			}
			log.Infof("\n ----------we received an indication message !----- %#v", indMsg)
			err = m.processIndication(ctx, indMsg)
			if err != nil {
				errCh <- err
			}
		}
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (m *Monitor) processIndication(ctx context.Context, indication e2api.Indication) error {
	err := m.processIndicationFormat1(ctx, indication)
	if err != nil {
		log.Warn(err)
		return err
	}

	return nil
}

func (m *Monitor) processIndicationFormat1(ctx context.Context, indication e2api.Indication) error {
	indHeader := e2smmet.E2SmMetIndicationHeader{}
	err := proto.Unmarshal(indication.Header, &indHeader)
	if err != nil {
		log.Warn(err)
		return err
	}
	log.Info(indHeader.IndicationHeaderFormats)

	indMessage := e2smmet.E2SmMetIndicationMessage{}
	err = proto.Unmarshal(indication.Payload, &indMessage)
	if err != nil {
		log.Warn(err)
		return err
	}

	indHdrFormat1 := indHeader.GetIndicationHeaderFormats().GetIndicationHeaderFormat1()
	indMsgFormat1 := indMessage.GetIndicationMessageFormats().GetIndicationMessageFormat1()
	log.Info("Received indication header format 1 %v:", indHdrFormat1.MeasInfoList.Value[0].Value)
	log.Info("Received indication message format 1:-- %v", indMsgFormat1.SubscriptId)

	return nil

}
