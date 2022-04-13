package response_converter

import (
	"strconv"

	"github.com/defipod/mochi/pkg/model"
)

type GetGuildsResponse struct {
	Data []*Guilds `json:"data"`
}

type Guilds struct {
	GuildId         string   `json:"guild_id"`
	Name            string   `json:"name"`
	BotScopes       []string `json:"bot_scopes"`
	VerifyChannelId string   `json:"verify_channel_id"`
}

func ConvertGetGuildsResponse(discordGuilds []*model.DiscordGuild) *GetGuildsResponse {
	var guilds []*Guilds
	for _, g := range discordGuilds {
		guilds = append(guilds, &Guilds{
			GuildId:         strconv.FormatInt(g.ID, 10),
			Name:            g.Name,
			BotScopes:       g.BotScopes,
			VerifyChannelId: "0",
		})
	}
	return &GetGuildsResponse{
		Data: guilds,
	}
}
