package model

type UserSubmittedAd struct {
	ID           int    `json:"id"`
	CreatorId    string `json:"creator_id"`
	AdChannelId  string `json:"ad_channel_id"`
	Status       string `json:"string"`
	Introduction string `json:"introduction"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Image        string `json:"image"`
	IsPodtownAd  bool   `json:"is_podtown_ad"`
}
