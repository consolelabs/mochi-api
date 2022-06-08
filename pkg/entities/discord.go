package entities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
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
	log := logger.NewLogrusLogger()
	log.Infof("Counting guild channels. GuildId: %v", guildID)
	nr_of_channels, nr_of_text_channels, nr_of_voice_channels, nr_of_stage_channels, nr_of_categories, nr_of_announcement_channels := 0, 0, 0, 0, 0, 0
	guildChannels, err := e.discord.GuildChannels(guildID)
	if err != nil {
		log.Error(err, "failed to get guild channels")
		return 0, 0, 0, 0, 0, 0, err
	}
	if len(guildChannels) == 0 {
		log.Info("Members not exist in current guild")
		return 0, 0, 0, 0, 0, 0, err
	}
	for _, channel := range guildChannels {
		// https://discord.com/developers/docs/resources/channel#channel-object-channel-types
		// Refer to discord doc: 0 -> text channel, 2 -> voice channel, 4 -> category, 5 -> announcement channel, 13 -> stage channel
		switch channel.Type {
		case consts.TextChannel:
			nr_of_text_channels = nr_of_text_channels + 1
		case consts.VoiceChannel:
			nr_of_voice_channels = nr_of_voice_channels + 1
		case consts.Category:
			nr_of_categories = nr_of_categories + 1
		case consts.AnnouncementChannel:
			nr_of_announcement_channels = nr_of_announcement_channels + 1
		case consts.StageChannel:
			nr_of_stage_channels = nr_of_stage_channels + 1
		default:
			fmt.Println("still not handle this case")
		}
	}
	nr_of_channels = len(guildChannels) - nr_of_categories
	return nr_of_channels, nr_of_text_channels, nr_of_voice_channels, nr_of_stage_channels, nr_of_categories, nr_of_announcement_channels, nil
}

func (e *Entity) CountGuildEmojis(guildID string) (int, int, int, error) {
	log := logger.NewLogrusLogger()
	log.Infof("Counting guild emojis. GuildId: %v", guildID)
	nr_of_emojis, nr_of_static_emojis, nr_of_animated_emojis := 0, 0, 0
	guildEmojis, err := e.discord.GuildEmojis(guildID)
	if err != nil {
		log.Error(err, "failed to get guild emojis")
		return 0, 0, 0, nil
	}
	if len(guildEmojis) == 0 {
		log.Info("Emojis not exist in current guild")
		return 0, 0, 0, nil
	}
	nr_of_emojis = len(guildEmojis)
	for _, emoji := range guildEmojis {
		// https://discord.com/developers/docs/resources/emoji#list-guild-emojis
		// Refer to discord doc: true is animated, false is static
		switch emoji.Animated {
		case true:
			nr_of_animated_emojis = nr_of_animated_emojis + 1
		case false:
			nr_of_static_emojis = nr_of_static_emojis + 1
		default:
			nr_of_static_emojis = nr_of_static_emojis + 1
		}
	}
	return nr_of_emojis, nr_of_static_emojis, nr_of_animated_emojis, nil
}

func (e *Entity) CountGuildStickers(guildID string) (int, int, int, error) {
	log := logger.NewLogrusLogger()
	log.Infof("Counting guild stickers. GuildId: %v", guildID)
	nr_of_stickers, nr_of_custom_stickers, nr_of_server_stickers := 0, 0, 0
	url := "https://discord.com/api/v9/guilds/" + guildID + "/stickers"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Error(err, "failed to set up request for guild stickers")
		return 0, 0, 0, err
	}
	request.Header.Set("Authorization", e.discord.Token)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Error(err, "failed to get guild stickers")
		return 0, 0, 0, err
	}
	defer resp.Body.Close()
	var guildStickers []response.DiscordGuildSticker
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err, "failed to read guild stickers response")
		return 0, 0, 0, err
	}
	if err := json.Unmarshal(respBody, &guildStickers); err != nil {
		log.Error(err, "failed to parse guild stickers response")
		return 0, 0, 0, err
	}
	if len(guildStickers) == 0 {
		log.Info("Stickers not exist in current guild")
		return 0, 0, 0, nil
	}
	// https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-types
	// Refer to discord doc: custom sticker is 1, server sticker is 2
	for _, sticker := range guildStickers {
		switch sticker.Type {
		case consts.CustomSticker:
			nr_of_custom_stickers = nr_of_custom_stickers + 1
		case consts.ServerSticker:
			nr_of_server_stickers = nr_of_server_stickers + 1
		default:
			nr_of_custom_stickers = nr_of_custom_stickers + 1
		}
	}
	nr_of_stickers = len(guildStickers)
	return nr_of_stickers, nr_of_custom_stickers, nr_of_server_stickers, nil
}

func (e *Entity) CountGuildRoles(guildID string) (int, error) {
	log := logger.NewLogrusLogger()
	log.Infof("Counting guild roles. GuildId: %v", guildID)
	guildRoles, err := e.discord.GuildRoles(guildID)
	if err != nil {
		log.Error(err, "failed to get guild roles")
		return 0, err
	}

	return len(guildRoles), nil
}

func (e *Entity) CountGuildMembers(guildID string) (int, int, int, error) {
	log := logger.NewLogrusLogger()
	log.Infof("Counting guild members. GuildId: %v", guildID)
	nr_of_members, nr_of_user, nr_of_bots := 0, 0, 0
	members := make([]response.DiscordGuildUser, 0)
	after := ""
	limit := 1000
	for {
		guildMembers, err := e.discord.GuildMembers(guildID, after, limit)
		if err != nil {
			log.Error(err, "failed to get guild members")
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

func (e *Entity) CreateGuildChannel(guildID string, countType string, coinData ...string) error {
	log := logger.NewLogrusLogger()
	err := e.UpdateOneGuildStats(guildID)
	if err != nil {
		log.Error(err, "failed to get guild stats from discord")
		return err
	}
	guildStat, err := e.GetByGuildID(guildID)
	if err != nil {
		log.Error(err, "failed to get guild stats from database")
		return err
	}

	// create channel count stat
	channelName := util.CreateChannelName(guildStat, countType, coinData...)
	createdChannel, err := e.discord.GuildChannelCreate(guildID, channelName, 2)
	if err != nil {
		log.Error(err, "failed to create discord channel")
		return err
	}

	_, err = e.discord.ChannelEditComplex(createdChannel.ID, &discordgo.ChannelEdit{
		PermissionOverwrites: []*discordgo.PermissionOverwrite{
			{
				ID:    guildID,
				Type:  0,
				Allow: 0,
				Deny:  1048576,
			},
		},
	})
	if err != nil {
		log.Error(err, "failed to update discord channel permission")
		return err
	}

	// store channel_id to db
	err = e.CreateDiscordGuildStatChannel(model.DiscordGuildStatChannel{
		GuildID:   guildID,
		ChannelID: createdChannel.ID,
		CountType: countType,
	})
	if err != nil {
		log.Error(err, "failed to store channel id")
		return err
	}
	return nil
}

func (e *Entity) UpdateOneGuildStats(guildID string) error {
	log := logger.NewLogrusLogger()
	log.Infof("Starting count stats for . GuildId: %v", guildID)
	nr_of_members, nr_of_user, nr_of_bots, err := e.CountGuildMembers(guildID)
	if err != nil {
		return err
	}
	nr_of_channels, nr_of_text_channels, nr_of_voice_channels, nr_of_stage_channels, nr_of_categories, nr_of_announcement_channels, err := e.CountGuildChannels(guildID)
	if err != nil {
		return err
	}
	nr_of_emojis, nr_of_static_emojis, nr_of_animated_emojis, err := e.CountGuildEmojis(guildID)
	if err != nil {
		return err
	}
	nr_of_stickers, nr_of_custom_stickers, nr_of_server_stickers, err := e.CountGuildStickers(guildID)
	if err != nil {
		return err
	}
	nr_of_roles, err := e.CountGuildRoles(guildID)
	if err != nil {
		return err
	}

	// update stats to database
	err = e.UpdateGuildStats(model.DiscordGuildStat{
		GuildID:     guildID,
		NrOfMembers: nr_of_members,
		NrOfUsers:   nr_of_user,
		NrOfBots:    nr_of_bots,

		NrOfChannels:             nr_of_channels,
		NrOfTextChannels:         nr_of_text_channels,
		NrOfVoiceChannels:        nr_of_voice_channels,
		NrOfStageChannels:        nr_of_stage_channels,
		NrOfCategories:           nr_of_categories,
		NrOfAnnouncementChannels: nr_of_announcement_channels,

		NrOfEmojis:         nr_of_emojis,
		NrOfStaticEmojis:   nr_of_static_emojis,
		NrOfAnimatedEmojis: nr_of_animated_emojis,

		NrOfStickers:       nr_of_stickers,
		NrOfCustomStickers: nr_of_custom_stickers,
		NrOfServerStickers: nr_of_server_stickers,
		NrOfRoles:          nr_of_roles,
	})
	if err != nil {
		return err
	}
	return nil
}

func (e *Entity) EditGuildChannel(guildID string, statChannel model.DiscordGuildStatChannel) error {
	log := logger.NewLogrusLogger()
	log.Infof("Edit stats channel for guild. GuildId: %v", guildID)
	guildStat, err := e.GetByGuildID(guildID)
	if err != nil {
		log.Error(err, "failed to get guild stats from database")
		return err
	}

	newChannelName := util.CreateChannelName(guildStat, statChannel.CountType)
	_, err = e.discord.ChannelEdit(statChannel.ChannelID, newChannelName)
	if err != nil {
		log.Error(err, "failed to edit channel name")
		return err
	}
	return nil
}

func (e *Entity) GetGuildChannel(channelID string) (*discordgo.Channel, error) {
	log := logger.NewLogrusLogger()
	log.Infof("Get channel info from discord for guild. ChannelId: %v", channelID)
	channel, err := e.discord.Channel(channelID)
	return channel, err
}

func (e *Entity) GetGuildById(guildID string) (*discordgo.Guild, error) {
	guild, err := e.discord.Guild(guildID)
	if err != nil {
		return nil, err
	}
	return guild, nil
}
