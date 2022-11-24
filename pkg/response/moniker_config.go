package response

import "github.com/defipod/mochi/pkg/model"

type MonikerConfigData struct {
	Moniker model.MonikerConfig `json:"moniker"`
	Value   float64             `json:"value"`
}

type MonikerConfigResponse struct {
	Data []MonikerConfigData `json:"data"`
}
