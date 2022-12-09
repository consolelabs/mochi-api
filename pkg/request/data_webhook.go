package request

type NotifyCompleteNftIntegrationRequest struct {
	CollectionAddress string `json:"collection_address"`
	ChainID           int64  `json:"chain_id"`
}

type NotifyCompleteNftSyncRequest struct {
	CollectionAddress string `json:"collection_address"`
	ChainID           int64  `json:"chain_id"`
}
