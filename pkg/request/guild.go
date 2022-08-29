package request

type CreateGuildRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateGuildRequest struct {
	GlobalXP   string `json:"global_xp"`
	LogChannel string `json:"log_channel"`
}
