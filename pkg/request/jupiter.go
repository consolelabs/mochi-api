package request

type QuoteResponse struct {
	InputMint            string      `json:"inputMint"`
	InAmount             string      `json:"inAmount"`
	OutputMint           string      `json:"outputMint"`
	OutAmount            string      `json:"outAmount"`
	OtherAmountThreshold string      `json:"otherAmountThreshold"`
	SwapMode             string      `json:"swapMode"`
	SlippageBps          int32       `json:"slippageBps"`
	PlatformFee          interface{} `json:"platformFee"`
	PriceImpactPct       string      `json:"priceImpactPct"`
}

type JupiterBuildSwapRouteRequest struct {
	QuoteResponse                 interface{} `json:"quoteResponse"`
	UserPublicKey                 string      `json:"userPublicKey"`
	WrapAndUnwrapSol              bool        `json:"wrapAndUnwrapSol"`
	ComputeUnitPriceMicroLamports string      `json:"computeUnitPriceMicroLamports"`
}
