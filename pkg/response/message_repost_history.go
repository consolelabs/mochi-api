package response

type RepostReactionEventResponse struct {
	Data RepostReactionEventData `json:"data"`
}

type RepostReactionEventData struct {
	Status          string `json:"status"`
	RepostChannelID string `json:"repost_channel_id"`
	RepostMessageID string `json:"repost_message_id"`
}
