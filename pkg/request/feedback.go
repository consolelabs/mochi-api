package request

type UserFeedbackRequest struct {
	DiscordID string `json:"discord_id"`
	MessageID string `json:"message_id"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	Command   string `json:"command"`
	Feedback  string `json:"feedback"`
	ProfileID string `json:"profile_id"`
}

type UpdateUserFeedbackRequest struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type GetUserFeedbackRequest struct {
	Page      int64  `form:"page,default=0"`
	Size      int64  `form:"size,default=5"`
	DiscordId string `form:"discord_id"`
	ProfileId string `form:"profile_id"`
	Command   string `form:"command"`
	Status    string `form:"status"`
}
