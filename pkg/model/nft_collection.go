package model

type NFTCollection struct {
	Address   string `json:"address"`
	Name      string `json:"name"`
	Symbol    string `json:"symbol"`
	ChainID   string `json:"chain_id"`
	ERCFormat string `json:"erc_format"`
}
