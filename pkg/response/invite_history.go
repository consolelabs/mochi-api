package response

type HandleInviteHistoryResponse struct {
	InviterID     string `json:"inviter_id"`
	InviteeID     string `json:"invitee_id"`
	InvitesAmount int    `json:"invites_amount"`
	IsVanity      bool   `json:"is_vanity"`
	IsBot         bool   `json:"is_bot"`
}

type GetInvitesLeaderboardResponse struct {
	Data []UserInvitesAggregation `json:"data"`
}

type UserInvitesAggregation struct {
	InviterID string `json:"inviter_id"`
	Regular   int    `json:"regular"`
	Fake      int    `json:"fake"`
	Left      int    `json:"left"`
}
