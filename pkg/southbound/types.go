package southbound

type client struct {
	Kpi       string            `json:"kpi"`
	Slice_id  string            `json:"slice_id"`
	Source    string            `json:"source"`
	Timestamp string            `json:"timestamp"`
	Unit      string            `json:"unit"`
	Value     string            `json:"value"`
	Labels    map[string]string `json:"labels"`
}

// templsate := `{
//     "kpi": "",
//     "slice_id": "None",
//     "source": "RAN",
//     "timestamp": "1652270699",
//     "unit": "None",
//     "value": 11,
//       "labels": [{"amf_ue_ngap_id":"558745"},{"gNB_ue_ngap_id":"0"},{"ue_tag":"0"}]
// }`
