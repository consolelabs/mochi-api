package krystal

type BuildStakeTxReq struct {
	ChainID      int        `json:"chainID"`
	EarningType  string     `json:"earningType"`
	ExtraData    *ExtraData `json:"extraData"`
	Platform     string     `json:"platform"`
	TokenAddress string     `json:"tokenAddress"`
	TokenAmount  string     `json:"tokenAmount"`
	UserAddress  string     `json:"userAddress"`
}

type BuildUnstakeTxReq struct {
	ChainID      int        `json:"chainID"`
	EarningType  string     `json:"earningType"`
	ExtraData    *ExtraData `json:"extraData"`
	Platform     string     `json:"platform"`
	TokenAddress string     `json:"tokenAddress"`
	TokenAmount  string     `json:"tokenAmount"`
	UserAddress  string     `json:"userAddress"`
}

type Ankr struct {
	UseTokenC bool `json:"useTokenC"`
}

type Lido struct {
	NftTokenID string `json:"nftTokenID"`
}

type ExtraData struct {
	Ankr *Ankr `json:"ankr"`
	Lido *Lido `json:"lido"`
}
