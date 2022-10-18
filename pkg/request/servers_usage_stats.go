package request

type UsageInformation struct {
	UserID          string `json:"user_id"`
	GuildID         string `json:"guild_id"`
	Command         string `json:"command"`
	Args            string `json:"args"`
	Success         bool   `json:"success"`
	ExecutionTimeMs int    `json:"execution_time_ms"`
}
