package request

type UserUpvoteProcessorRequest struct {
	App    string        `json:"dapp"`
	Action string        `json:"action"`
	Data   UserDiscordID `json:"data"`
}

type UserDiscordID struct {
	UserID string `json:"user_discord_id"`
}
