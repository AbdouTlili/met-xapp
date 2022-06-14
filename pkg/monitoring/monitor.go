// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package monitoring

import (
	"bytes"
	"context"
	"encoding/binary"
	"html/template"
	"math"

	"github.com/AbdouTlili/met-xapp/pkg/broker"
	"github.com/AbdouTlili/met-xapp/pkg/northbound"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/streadway/amqp"
	"google.golang.org/protobuf/proto"

	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"

	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger()

// Monitor indication monitor
type Monitor struct {
	streamReader           broker.StreamReader
	northboundBrokerWriter northbound.BrokerClient
}

func NewMonitor(streamReader broker.StreamReader, nbBrokerWriter northbound.BrokerClient) *Monitor {

	return &Monitor{
		streamReader:           streamReader,
		northboundBrokerWriter: nbBrokerWriter,
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
	// log.Info(indHeader.IndicationHeaderFormats)

	indMessage := e2smmet.E2SmMetIndicationMessage{}
	err = proto.Unmarshal(indication.Payload, &indMessage)
	if err != nil {
		log.Warn(err)
		return err
	}

	indHdrFormat1 := indHeader.GetIndicationHeaderFormats().GetIndicationHeaderFormat1()
	indMsgFormat1 := indMessage.GetIndicationMessageFormats().GetIndicationMessageFormat1()

	// log.Info("\nReceived indication header format 1 %v:", indHdrFormat1.MeasInfoList.Value[0].Value)

	// log.Info("\nReceived indication header format 1 %v:", indHdrFormat1.ColletStartTime.Value)
	// log.Info("\nReceived indication message format 1:-- %v", indMsgFormat1.SubscriptId)

	// northbound.Kpi{
	// 	Nssmf:     northbound.Kpi_RAN,
	// 	Id:        indHdrFormat1.MetNodeId.Value,
	// 	Region:    "None",
	// 	Timestamp: Float64frombytes(indHdrFormat1.ColletStartTime.Value),

	// }

	t, err := template.New("").Parse(tmp)
	if err != nil {
		return err
	}

	for _, mr := range indMsgFormat1.MeasData.Value {
		// fmt.Printf("\nueid is : %v, uetag is : %v, records are : \n", mr.UeId, mr.UeTag)
		for i, mri := range mr.MeasRecordItemList.Value {
			if i == len(mr.MeasRecordItemList.Value)-1 {
				break
			}
			tmpKpi := KpiRec{
				Kpi:            indHdrFormat1.MeasInfoList.Value[i].GetValue(),
				Slice_id:       "None",
				Source:         "RAN",
				Timestamp:      indHdrFormat1.ColletStartTime.String(),
				Unit:           "None",
				Value:          mri.GetValue(),
				Amf_ue_ngap_id: mr.MeasRecordItemList.Value[len(mr.MeasRecordItemList.Value)-1].GetValue(),
				Ue_tag:         "None",
				GNB_ue_ngap_id: mr.UeId.String(),
			}
			log.Info(tmpKpi)

			log.Info("KPI object created")

			var tpl bytes.Buffer
			err = t.Execute(&tpl, tmpKpi)
			if err != nil {
				return err
			}

			// result := tpl.String()
			// fmt.Printf("%v", result)

			message := amqp.Publishing{
				Body: tpl.Bytes(),
			}

			// Attempt to publish a message to the queue

			if err := m.northboundBrokerWriter.ChannelBroker.Publish("", "onos-queue1", false, false, message); err != nil {
				return err
			}

		}
	}

	// log.Info("Received indication header format 1 %v:", indHdrFormat1.MeasInfoList.Value[0].Value)
	// log.Info("Received indication message format 1:-- %v", indMsgFormat1.SubscriptId)
	// log.Info("\nMET  1 :-- %v", indMsgFormat1.MeasData.Value[0].GetMeasRecordItem()[0].GetInteger())
	// log.Info("\n\nMET  2 :-- %v", indMsgFormat1.MeasData.Value[0].GetMeasRecordItem()[1].GetInteger())

	return nil

}

func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

type KpiRec struct {
	Kpi            string `json:"kpi"`
	Slice_id       string `json:"slice_id"`
	Source         string `json:"source"`
	Timestamp      string `json:"timestamp"`
	Unit           string `json:"unit"`
	Value          string `json:"value"`
	Amf_ue_ngap_id string `json:"Amf_ue_ngap_id"`
	GNB_ue_ngap_id string `json:"GNB_ue_ngap_id"`
	Ue_tag         string `json:"Ue_tag"`
}

const tmp = `"{"kpi": "{{.Kpi}}","slice_id": "{{.Slice_id}}","source": "RAN","timestamp": "{{.Timestamp}}","unit": "{{.Unit}}","value": {{.Value}},"labels": [{"amf_ue_ngap_id":"{{.Amf_ue_ngap_id}}"},{"gNB_ue_ngap_id":"{{.GNB_ue_ngap_id}}"},{"ue_tag":"{{.Ue_tag}}"}]}"`
