package response

import "time"

type BluemoveCollectionResponse struct {
	Data BluemoveCollection `json:"data"`
}
type BluemoveCollection struct {
	Address    string    `json:"address"`
	Name       string    `json:"name"`
	Symbol     string    `json:"symbol"`
	ChainID    string    `json:"chain_id"`
	ERCFormat  string    `json:"erc_format"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
	Image      string    `json:"image"`
	Author     string    `json:"author"`
}
