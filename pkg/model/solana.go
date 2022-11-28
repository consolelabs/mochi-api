package model

type SolanaCollectionMetadata struct {
	Success bool        `json:"success"`
	Data    SolscanData `json:"data"`
}

type SolscanData struct {
	Success bool                 `json:"success"`
	Data    SolanaCollectionInfo `json:"data"`
}

type SolanaCollectionInfo struct {
	Avatar       string `json:"avatar"`
	Collection   string `json:"collection"`
	Symbol       string `json:"symbol"`
	CollectionId string `json:"collectionId"`
}
