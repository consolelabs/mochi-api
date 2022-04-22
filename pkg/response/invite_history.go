package response

type HandleInviteHistoryResponse struct {
	InviterID     string `json:"inviter_id"`
	InviteeID     string `json:"invitee_id"`
	InvitesAmount int    `json:"invites_amount"`
	IsVanity      bool   `json:"is_vanity"`
	IsBot         bool   `json:"is_bot"`
}
