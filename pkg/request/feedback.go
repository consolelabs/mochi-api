package request

type UserFeedbackRequest struct {
	DiscordID string `json:"discord_id"`
	MessageID string `json:"message_id"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	Command   string `json:"command"`
	Feedback  string `json:"feedback"`
}

type UpdateUserFeedbackRequest struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}
