package request

type CreateTwitterSaleConfigRequest struct {
	Address     string `json:"address"`
	Marketplace string `json:"marketplace"`
	ChainID     int    `json:"chain_id"`
}
