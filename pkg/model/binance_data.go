package model

type IntegrationMessage struct {
	ApiKey    string `json:"apiKey,omitempty"`
	ApiSecret string `json:"apiSecret,omitempty"`
	ProfileId string `json:"profileId,omitempty"`
}

type KafkaIntegrateMessage struct {
	Type string             `json:"type,omitempty"`
	Data KafkaIntegrateData `json:"data,omitempty"`
}

type KafkaIntegrateData struct {
	IntegrationMessage *IntegrationMessage `json:"integrationMessage,omitempty"`
}
