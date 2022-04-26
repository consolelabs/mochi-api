package request

type CreateUserRequest struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	GuildID   string `json:"guild_id"`
	InvitedBy string `json:"invited_by"`
}
