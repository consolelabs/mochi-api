package mochipay

type CreateTokenRequest struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Decimal     int64  `json:"decimal"`
	ChainId     string `json:"chain_id"`
	Address     string `json:"address"`
	Icon        string `json:"icon"`
	CoinGeckoId string `json:"coin_gecko_id"`
}

type GetTokenRequest struct {
	Symbol  string `json:"symbol"`
	ChainId string `json:"chain_id"`
}

type CreateBatchTokenRequest struct {
	Tokens []CreateTokenRequest `json:"tokens"`
}

type TokenProperties struct {
	// can add more if want
	ChainId string `json:"chain_id"`
	Address string `json:"address"`
}

type Wallet struct {
	ProfileGlobalId string `json:"profile_global_id"`
}

type TransferV2Request struct {
	From     *Wallet                `json:"from"`
	Tos      []*Wallet              `json:"tos"`
	Amount   []string               `json:"amount"`
	TokenId  string                 `json:"token_id"`
	Platform string                 `json:"platform"`
	Action   string                 `json:"action" binding:"required"`
	Metadata map[string]interface{} `json:"metadata"`
}

type ApplicationBaseHeaderRequest struct {
	Application string `header:"x-application"`
	Message     string `header:"x-message"`
	Signature   string `header:"x-signature"`
}

type ApplicationTransferRequest struct {
	AppId    string                       `json:"appId"`
	Metadata ApplicationTransferMetadata  `json:"metadata"`
	Header   ApplicationBaseHeaderRequest `json:"header"`
}

type ApplicationTransferMetadata struct {
	RecipientIds []string `json:"recipient_ids"`
	Amounts      []string `json:"amounts"`
	TokenId      string   `json:"token_id" `
	References   string   `json:"references"`
	Description  string   `json:"description"`
	Platform     string   `json:"platform"`
}
