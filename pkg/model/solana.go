package model

type SolanaCollectionMetadata struct {
	OffChainData SolanaOffchainData `json:"off_chain_data"`
	Data         SolanaData         `json:"data"`
}

type SolanaOffchainData struct {
	Image  string `json:"image"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type SolanaData struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
