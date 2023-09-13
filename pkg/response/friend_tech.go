package response

import "time"

type FriendTechKeysResponse struct {
	Data []FriendTechKey `json:"data"`
}

type FriendTechKey struct {
	ID              int64     `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	Address         string    `json:"address"`
	TwitterUsername string    `json:"twitterUsername"`
	TwitterPfpUrl   string    `json:"twitterPfpUrl"`
	ProfileChecked  bool      `json:"profileChecked"`
	Price           float64   `json:"price"`
	Supply          int       `json:"supply"`
	Holders         int       `json:"holders"`
}
