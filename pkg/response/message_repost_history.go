package response

type RepostReactionEventResponse struct {
	Data RepostReactionEventData `json:"data"`
}

type RepostReactionEventData struct {
	Status               string `json:"status"`
	RepostChannelID      string `json:"repost_channel_id"`
	RepostMessageID      string `json:"repost_message_id"`
	ReactionType         string `json:"reaction_type"`
	OriginStartMessageID string `json:"origin_start_message_id"`
	OriginStopMessageID  string `json:"origin_stop_message_id"`
}
