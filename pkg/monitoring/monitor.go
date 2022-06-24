// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package monitoring

import (
	"context"
	"encoding/binary"

	// "encoding/binary"
	"math"
	"strconv"

	"github.com/AbdouTlili/met-xapp/pkg/broker"
	"github.com/AbdouTlili/met-xapp/pkg/northbound"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"google.golang.org/protobuf/proto"

	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"

	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger()

// Monitor indication monitor
type Monitor struct {
	streamReader             broker.StreamReader
	northboundBrokerProducer northbound.BrokerClient
}

func NewMonitor(streamReader broker.StreamReader, nbBrokerProducer northbound.BrokerClient) *Monitor {

	return &Monitor{
		streamReader:             streamReader,
		northboundBrokerProducer: nbBrokerProducer,
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

	for _, mr := range indMsgFormat1.MeasData.Value {
		// fmt.Printf("\nueid is : %v, uetag is : %v, records are : \n", mr.UeId, mr.UeTag)
		for i, mri := range mr.MeasRecordItemList.Value {
			if i == len(mr.MeasRecordItemList.Value)-1 {
				break
			}
			// creating the Labels in a map

			labels := make(map[string]string)

			labels["gnb_ue_ngap_id"] = strconv.FormatInt(mr.UeId.GetValue()-1, 10)
			labels["amf_ue_ngap_id"] = mr.MeasRecordItemList.Value[len(mr.MeasRecordItemList.Value)-1].GetValue()
			labels["ue_tag"] = "None"

			tmpKpi := KpiRecObj{
				Kpi:       indHdrFormat1.MeasInfoList.Value[i].GetValue(),
				Slice_id:  "None",
				Source:    "RAN",
				Timestamp: indHdrFormat1.ColletStartTime.GetValue(),
				Unit:      "None",
				Value:     mri.GetValue(),
				Labels:    labels,
			}

			// log.Info(binary.BigEndian.Uint64(indHdrFormat1.ColletStartTime.GetValue()))
			log.Infof("------ : %#v", indHdrFormat1.ColletStartTime.GetValue())

			// log.Info(tmpKpi)

			protoPayload, err := CreateKpiMessagefromObj(tmpKpi)

			if err != nil {
				return err
			}

			log.Infof("KPI object created")

			// result := tpl.String()
			// fmt.Printf("%v", result)

			log.Info("// Attempt to publish a message to the queue")

			log.Infof("the payload content : %#v \n", protoPayload)

			// tmpKpiDecode := northbound.Kpi{}
			// if err := proto.Unmarshal(protoPayload, &tmpKpiDecode); err != nil {
			// 	log.Warnf("\n error Unmarshalling KpiBytes : %v", err)

			// }

			// log.Infof("the decodede payload : %v ", protoPayload)

			if err := m.northboundBrokerProducer.Publish(protoPayload); err != nil {
				log.Warn("error : ", err)
				return err
			}

			log.Info("message sent to the broker")
		}
	}

	return nil

}

func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

type KpiRecObj struct {
	Kpi       string            `json:"kpi"`
	Slice_id  string            `json:"slice_id"`
	Source    string            `json:"source"`
	Timestamp []byte            `json:"timestamp"`
	Unit      string            `json:"unit"`
	Value     string            `json:"value"`
	Labels    map[string]string `json:"labels"`
}

// this served a a teplate for the needed JSON message but no longer used by It still here for the sake of calrity
// so someone who does not know what the structure of the json looks like can see this
const tmp = `"{"kpi": "{{.Kpi}}","slice_id": "{{.Slice_id}}","source": "RAN","timestamp": "{{.Timestamp}}","unit": "{{.Unit}}","value": {{.Value}},"labels": [{"amf_ue_ngap_id":"{{.Amf_ue_ngap_id}}"},{"gNB_ue_ngap_id":"{{.GNB_ue_ngap_id}}"},{"ue_tag":"{{.Ue_tag}}"}]}"`

func CreateKpiMessagefromObj(kpiObj KpiRecObj) ([]byte, error) {

	gnb_ue_ngap_id := northbound.CreateParameter("gnb_ue_ngap_id", kpiObj.Labels["gnb_ue_ngap_id"])

	amf_ue_ngap_id := northbound.CreateParameter("amf_ue_ngap_id", kpiObj.Labels["amf_ue_ngap_id"])

	ue_tag := northbound.CreateParameter("ue_tag", kpiObj.Labels["ue_tag"])

	value, err := strconv.ParseFloat(kpiObj.Value, 64)

	params := make([]*northbound.Parameter, 0)
	params = append(params, gnb_ue_ngap_id, amf_ue_ngap_id, ue_tag)

	payload := northbound.CreatePayload(value, params)
	kpi := northbound.CreateKpiMessage(northbound.Kpi_RAN, // nssmf
		0,     //id
		"RAN", //region
		float64(binary.BigEndian.Uint32([]byte(kpiObj.Timestamp))), //timestamp
		0,          //nssid
		kpiObj.Kpi, //metric
		"None",     //unit
		payload)    //payload

	// log.Infof("\n ========== KPI object created : %#v", kpi)

	protoBytes, err := proto.Marshal(kpi)
	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	}

	// log.Infof("\n ========== KPI object bytes : %#v", protoBytes)
	return protoBytes, nil
}
