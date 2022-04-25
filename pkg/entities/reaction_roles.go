package entities

import (
	"encoding/json"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) GetReactionRoleByMessageID(guildID, messageID, reaction string) (*response.RoleReactionResponse, error) {
	config, err := e.repo.ReactionRoleConfig.GetByMessageID(guildID, messageID)
	if err != nil {
		return nil, err
	}

	var roles []response.Role
	err = json.Unmarshal([]byte(config.ReactionRoles), &roles)
	if err != nil {
		return nil, err
	}

	var filteredRole response.Role
	for _, r := range roles {
		if r.Reaction == reaction {
			filteredRole = r
		}
	}

	var res = response.RoleReactionResponse{
		GuildID:   config.GuildID,
		ChannelID: config.ChannelID,
		MessageID: config.MessageID,
		Role:      filteredRole,
	}
	return &res, nil
}
