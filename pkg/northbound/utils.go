package northbound

func CreateParameter(name string, value string) *Parameter {
	return &Parameter{
		Name:  name,
		Value: value,
	}
}

// here we   ignore the unit sinc it is duplicated and already exists in the mail KPM message
func CreatePayload(value float64, params []*Parameter) *Payload {
	return &Payload{
		Value:  value,
		Params: params,
	}
}

func AddParameterToPayload(payload *Payload, parameter *Parameter) {
	payload.Params = append(payload.Params, parameter)
}

func CreateKpiMessage(nssmf Kpi_NSSMF,
	id int64,
	region string,
	timestamp float64,
	nssid int64,
	metric string,
	unit string,
	payload *Payload) *Kpi {
	return &Kpi{Nssmf: nssmf,
		Id:        id,
		Region:    region,
		Timestamp: timestamp,
		Nssid:     nssid,
		Metric:    metric,
		Unit:      unit,
		Payload:   payload,
	}
}
