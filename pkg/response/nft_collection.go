package response

type IndexerNFTResponse struct {
	CollectionAddress string      `json:"collection_address"`
	CollectionName    string      `json:"collection_name"`
	Symbol            string      `json:"symbol"`
	ChainID           string      `json:"chain_id"`
	TradingVolume     TradeVolume `json:"trading_volume"`
}
type TradeVolume struct {
	BuyNumber  int `json:"nr_of_buy"`
	SellNumber int `json:"nr_of_sell"`
}
