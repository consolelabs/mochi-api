package request

type NotifyCompleteNftIntegrationRequest struct {
	CollectionAddress string `json:"collection_address"`
	ChainID           int64  `json:"chain_id"`
}

type NotifyCompleteNftSyncRequest struct {
	CollectionAddress string `json:"collection_address"`
	ChainID           int64  `json:"chain_id"`
}

type NotifySaleMarketplaceRequest struct {
	Event       string     `json:"event"`
	TokenId     string     `json:"token_id"`
	Address     string     `json:"address"`
	Marketplace string     `json:"marketplace"`
	Transaction string     `json:"transaction"`
	From        string     `json:"from"`
	To          string     `json:"to"`
	Price       TokenModel `json:"price"`
	Hodl        int        `json:"hold"`
	LastPrice   TokenModel `json:"last_price"`
	ChainId     int64      `json:"chain_id"`
}

type TokenModel struct {
	Token  TokenInfoModel `json:"token"`
	Amount string         `json:"amount"`
}
type TokenInfoModel struct {
	Symbol   string `json:"symbol"`
	IsNative bool   `json:"is_native"`
	Address  string `json:"address"`
	Decimal  int    `json:"decimal"`
}
