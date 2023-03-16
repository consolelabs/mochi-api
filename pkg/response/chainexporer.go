package response

type ChainExplorerGasTracker struct {
	Status  string           `json:"status"`
	Message string           `json:"message"`
	Result  GasTrackerResult `json:"result"`
}

type GasTrackerResult struct {
	LastBlock       string `json:"LastBlock"`
	SafeGasPrice    string `json:"SafeGasPrice"`
	ProposeGasPrice string `json:"ProposeGasPrice"`
	FastGasPrice    string `json:"FastGasPrice"`
	SuggestBaseFee  string `json:"SuggestBaseFee"`
	GasUsedRatio    string `json:"GasUsedRatio"`
}
