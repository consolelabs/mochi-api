package entities

import (
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) GetReactionRoles(guildID string) (*response.RoleReactionsResponse, error) {
	guilds, err := e.repo.ReactionRoleConfig.Gets(guildID)
	if err != nil {
		return nil, err
	}

	var res response.RoleReactionsResponse
	res.Data = make([]*response.RoleReactionResponse, 0)
	for _, g := range guilds {
		res.Data = append(res.Data, &response.RoleReactionResponse{
			GuildID:       g.GuildID,
			ChannelID:     g.ChannelID,
			Title:         g.Title,
			TitleUrl:      g.TitleUrl,
			ThumbnailUrl:  g.ThumbnailUrl,
			Description:   g.Description,
			FooterImage:   g.FooterImage,
			FooterMessage: g.FooterMessage,
			MessageID:     g.MessageID,
			ReactionRoles: g.ReactionRoles,
		})
	}

	return &res, nil
}
