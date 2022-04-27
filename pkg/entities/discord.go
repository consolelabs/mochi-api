package entities

import (
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) GetGuildUsersFromDiscord(guildID string) ([]response.DiscordGuildUser, error) {
	members := make([]response.DiscordGuildUser, 0)

	after := ""
	limit := 1000
	for {
		guildMembers, err := e.discord.GuildMembers(guildID, after, limit)
		if err != nil {
			return nil, err
		}

		for _, member := range guildMembers {
			// ignore bots
			if member.User.Bot {
				continue
			}

			members = append(members, response.DiscordGuildUser{
				User: &response.DiscordUser{
					ID:       member.User.ID,
					Username: member.User.Username,
				},
				GuildID:  guildID,
				Nickname: member.Nick,
			})
		}

		if len(guildMembers) < limit {
			break
		}

		after = guildMembers[len(guildMembers)-1].User.ID
	}

	return members, nil
}
