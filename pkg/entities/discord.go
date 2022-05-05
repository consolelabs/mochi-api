package entities

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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

func (e *Entity) CountGuildChannels(guildID string) (int, int, int, int, int, int, error) {
	nr_of_channels, nr_of_text_channels, nr_of_voice_channels, nr_of_stage_channels, nr_of_categories, nr_of_announcement_channels := 0, 0, 0, 0, 0, 0
	guildChannels, err := e.discord.GuildChannels(guildID)

	if err != nil {
		return 0, 0, 0, 0, 0, 0, err
	}

	if len(guildChannels) == 0 {
		return 0, 0, 0, 0, 0, 0, err
	}

	for _, channel := range guildChannels {
		if channel.Type == 0 {
			nr_of_text_channels = nr_of_text_channels + 1
		} else if channel.Type == 2 {
			nr_of_voice_channels = nr_of_voice_channels + 1
		} else if channel.Type == 4 {
			nr_of_categories = nr_of_categories + 1
		} else if channel.Type == 5 {
			nr_of_announcement_channels = nr_of_announcement_channels + 1
		} else if channel.Type == 13 {
			nr_of_stage_channels = nr_of_stage_channels + 1
		}
	}

	nr_of_channels = len(guildChannels) - nr_of_categories
	return nr_of_channels, nr_of_text_channels, nr_of_voice_channels, nr_of_stage_channels, nr_of_categories, nr_of_announcement_channels, nil
}

func (e *Entity) CountGuildEmojis(guildID string) (int, int, int, error) {
	nr_of_emojis, nr_of_static_emojis, nr_of_animated_emojis := 0, 0, 0

	guildEmojis, err := e.discord.GuildEmojis(guildID)
	if err != nil {
		return 0, 0, 0, nil
	}

	if len(guildEmojis) == 0 {
		return 0, 0, 0, nil
	}

	nr_of_emojis = len(guildEmojis)

	for _, emoji := range guildEmojis {
		if emoji.Animated {
			nr_of_animated_emojis = nr_of_animated_emojis + 1
		} else if !emoji.Animated {
			nr_of_static_emojis = nr_of_static_emojis + 1
		}
	}

	return nr_of_emojis, nr_of_static_emojis, nr_of_animated_emojis, nil
}

func (e *Entity) CountGuildStickers(guildID string) (int, int, int, error) {
	nr_of_stickers, nr_of_standard_stickers, nr_of_guild_stickers := 0, 0, 0
	url := "https://discord.com/api/v9/guilds/" + guildID + "/stickers"

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, 0, 0, err
	}

	request.Header.Set("Authorization", e.discord.Token)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return 0, 0, 0, err
	}
	defer resp.Body.Close()

	var guildStickers []response.DiscordGuildSticker
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, 0, err
	}
	if err := json.Unmarshal(respBody, &guildStickers); err != nil {
		return 0, 0, 0, err
	}

	if len(guildStickers) == 0 {
		return 0, 0, 0, nil
	}

	for _, sticker := range guildStickers {
		if sticker.Type == 1 {
			nr_of_standard_stickers = nr_of_standard_stickers + 1
		} else if sticker.Type == 2 {
			nr_of_guild_stickers = nr_of_guild_stickers + 1
		}
	}
	nr_of_stickers = len(guildStickers)
	return nr_of_stickers, nr_of_standard_stickers, nr_of_guild_stickers, nil
}

func (e *Entity) CountGuildRoles(guildID string) (int, error) {
	guildRoles, err := e.discord.GuildRoles(guildID)
	if err != nil {
		return 0, err
	}

	return len(guildRoles), nil
}

func (e *Entity) CountGuildMembers(guildID string) (int, int, int, error) {
	nr_of_members, nr_of_user, nr_of_bots := 0, 0, 0
	members := make([]response.DiscordGuildUser, 0)

	after := ""
	limit := 1000
	for {
		guildMembers, err := e.discord.GuildMembers(guildID, after, limit)
		if err != nil {
			return 0, 0, 0, err
		}

		for _, member := range guildMembers {
			members = append(members, response.DiscordGuildUser{
				User: &response.DiscordUser{
					ID:       member.User.ID,
					Username: member.User.Username,
					Bot:      member.User.Bot,
				},
			})
		}

		if len(guildMembers) < limit {
			break
		}

		after = guildMembers[len(guildMembers)-1].User.ID
	}

	for _, member := range members {
		if member.User.Bot {
			nr_of_bots = nr_of_bots + 1
		} else if !member.User.Bot {
			nr_of_user = nr_of_user + 1
		}
	}
	nr_of_members = len(members)

	return nr_of_members, nr_of_user, nr_of_bots, nil
}
