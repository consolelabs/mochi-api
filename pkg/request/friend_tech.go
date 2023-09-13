package request

type TrackFriendTechKeyRequest struct {
	ProfileId       string `json:"profile_id,omitempty" binding:"required"`
	KeyAddress      string `json:"key_address,omitempty" binding:"required"`
	IncreaseAlertAt int    `json:"increase_alert_at,omitempty"`
	DecreaseAlertAt int    `json:"decrease_alert_at,omitempty"`
}

type UpdateFriendTechKeyTrackRequest struct {
	IncreaseAlertAt int `json:"increase_alert_at,omitempty"`
	DecreaseAlertAt int `json:"decrease_alert_at,omitempty"`
}
