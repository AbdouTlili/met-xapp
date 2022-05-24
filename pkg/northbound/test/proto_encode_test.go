package test

import (
	"testing"

	"github.com/AbdouTlili/met-xapp/pkg/northbound"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func Test_EncodeDecodeParameter(t *testing.T) {

	param := northbound.CreateParameter("test", "test")

	protoBytes, err := proto.Marshal(param)

	assert.NoError(t, err)

	decodedParam := &northbound.Parameter{}
	err = proto.Unmarshal(protoBytes, decodedParam)

	assert.NoError(t, err)
	assert.Equal(t, decodedParam.Name, param.Name)
	assert.Equal(t, param.Value, decodedParam.Value)
}

func Test_EncodeDecodePayload(t *testing.T) {

	param := northbound.CreateParameter("test", "test")

	params := make([]*northbound.Parameter, 0)

	params = append(params, param)

	pld := northbound.CreatePayload(3.3, params)

	protoBytes, err := proto.Marshal(pld)
	assert.NoError(t, err)

	// decoding

	tmp_pld := northbound.Payload{}

	err = proto.Unmarshal(protoBytes, &tmp_pld)
	assert.NoError(t, err)

	assert.Equal(t, tmp_pld.Params[0].Name, pld.Params[0].Name)
	assert.Equal(t, tmp_pld.Params[0].Value, pld.Params[0].Value)
	assert.Equal(t, tmp_pld.Value, pld.Value)

}

func Test_KpiMessageEncodeDecode(t *testing.T) {

	param := northbound.CreateParameter("test", "test")

	params := make([]*northbound.Parameter, 0)

	params = append(params, param)

	pld := northbound.CreatePayload(3.3, params)

	kpi := northbound.CreateKpiMessage(northbound.Kpi_RAN,
		15,
		"R1",
		3.66,
		55,
		"QCI",
		"None",
		pld)

	protoBytes, err := proto.Marshal(kpi)
	assert.NoError(t, err)

	// decoding

	tmp_kpi := northbound.Kpi{}

	err = proto.Unmarshal(protoBytes, &tmp_kpi)
	assert.NoError(t, err)

	assert.Equal(t, tmp_kpi.Payload.Params[0].Name, kpi.Payload.Params[0].Name)
	assert.Equal(t, tmp_kpi.Payload.Params[0].Value, kpi.Payload.Params[0].Value)
	assert.Equal(t, tmp_kpi.Nssid, kpi.Nssid)
	assert.Equal(t, tmp_kpi.Id, kpi.Id)
	assert.Equal(t, tmp_kpi.Metric, kpi.Metric)
	assert.Equal(t, tmp_kpi.Region, kpi.Region)
	assert.Equal(t, tmp_kpi.Timestamp, kpi.Timestamp)
	assert.Equal(t, tmp_kpi.Unit, kpi.Unit)
}
