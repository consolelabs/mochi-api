package birdeye

type TokenPrice struct {
	Data struct {
		Value           float64 `json:"value"`
		UpdateUnixTime  int     `json:"updateUnixTime"`
		UpdateHumanTime string  `json:"updateHumanTime"`
	} `json:"data"`
	Success bool `json:"success"`
}
