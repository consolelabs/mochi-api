package request

type InsertUserAd struct {
	CreatorId    string `json:"creator_id"`
	Introduction string `json:"introduction"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Image        string `json:"image"`
	IsPodtownAd  bool   `json:"is_podtown_ad"`
}

type InitAdSubmission struct {
	GuildId   string `json:"guild_id"`
	ChannelId string `json:"channel_id"`
}

type DeleteUserAd struct {
	ID int `json:"id"`
}
type UpdateUserAd struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}
