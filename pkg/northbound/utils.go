package northbound

// here we   ignore the unit sinc it is duplicated and already exists in the mail KPM message
func CreateKpiPayload(name string, value string) Payload {
	return Payload{
		Params: make([]*Parameter, 0),
	}
}

func CreateParameter(name string, value string) Parameter {
	return Parameter{
		Name:  name,
		Value: value,
	}
}

func AddParameterToPayload(payload Payload, parameter Parameter) {

}
