package response

type Role struct {
	ID       string `json:"id"`
	Reaction string `json:"reaction"`
}

type RoleReactionResponse struct {
	MessageID string `json:"message_id"`
	ChannelID string `json:"channel_id"`
	GuildID   string `json:"guild_id"`
	Role      Role   `json:"role"`
}

type RoleReactionConfigResponse struct {
	MessageID string `json:"message_id"`
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
	Roles     []Role `json:"roles"`
	Success   bool   `json:"success"`
}

type RoleReactionByMessage struct {
	MessageID string `json:"message_id"`
	ChannelID string `json:"channel_id"`
	Roles     []Role `json:"roles"`
}

type ListRoleReactionResponse struct {
	GuildID string                  `json:"guild_id"`
	Configs []RoleReactionByMessage `json:"configs"`
	Success bool                    `json:"success"`
}

type DataListRoleReactionResponse struct {
	*PaginationResponse `json:",omitempty"`
	Data                ListRoleReactionResponse `json:"data,omitempty"`
}

type DataFilterConfigByReaction struct {
	*PaginationResponse `json:",omitempty"`
	Data                *RoleReactionResponse `json:"data"`
}
