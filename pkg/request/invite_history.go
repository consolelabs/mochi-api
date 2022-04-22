package request

type CreateInviteHistoryRequest struct {
	GuildID string `json:"guild_id"`
	Inviter string `json:"inviter"`
	Invitee string `json:"invitee"`
	Type    string `json:"type"`
}
