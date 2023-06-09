package model

type IntegrationMessage struct {
	ApiKey    string `json:"api_key,omitempty"`
	ApiSecret string `json:"api_secret,omitempty"`
	ProfileId string `json:"profile_id,omitempty"`
}

type KafkaIntegrateMessage struct {
	Type string             `json:"type,omitempty"`
	Data KafkaIntegrateData `json:"data,omitempty"`
}

type KafkaIntegrateData struct {
	IntegrationMessage *IntegrationMessage `json:"integration_message,omitempty"`
}
