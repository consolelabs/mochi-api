package model

type CommandCooldown int

const (
	NoCooldown CommandCooldown = iota
	UserCooldown
	ServerCooldown
)

type GuildCustomCommand struct {
	ID                  string          `json:"id"`
	GuildID             string          `json:"guild_id"`
	Description         string          `json:"description"`
	Actions             JSON            `json:"actions"`
	Cooldown            CommandCooldown `json:"cooldown,omitempty"`
	CooldownDuration    int             `json:"cooldown_duration,omitempty"`
	Enabled             bool            `json:"enabled"`
	RolesPermissions    JSON            `json:"roles_permissions,omitempty"`
	ChannelsPermissions JSON            `json:"channels_permissions,omitempty"`
}
