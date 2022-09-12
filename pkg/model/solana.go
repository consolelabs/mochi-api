package model

type SolanaCollectionMetadata struct {
	OffChainData SolanaOffchainData `json:"off_chain_data"`
}

type SolanaOffchainData struct {
	Image  string `json:"image"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
