package util

import (
	"strconv"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/model"
)

func CreateChannelName(guildStat *model.DiscordGuildStat, countType string) string {
	switch countType {
	case consts.Members:
		nrOfMembersStr := strconv.Itoa(guildStat.NrOfMembers)
		return "Members - " + nrOfMembersStr
	case consts.Users:
		nrOfUsersStr := strconv.Itoa(guildStat.NrOfUsers)
		return "Users - " + nrOfUsersStr
	case consts.Bots:
		nrOfBotsStr := strconv.Itoa(guildStat.NrOfBots)
		return "Bots - " + nrOfBotsStr
	case consts.Channels:
		nrOfChannelsStr := strconv.Itoa(guildStat.NrOfChannels)
		return "Channels - " + nrOfChannelsStr
	case consts.TextChannels:
		nrOfTextChannelsStr := strconv.Itoa(guildStat.NrOfTextChannels)
		return "Text Channels - " + nrOfTextChannelsStr
	case consts.VoiceChannels:
		nrOfVoiceChannelsStr := strconv.Itoa(guildStat.NrOfVoiceChannels)
		return "Voice Channels - " + nrOfVoiceChannelsStr
	case consts.StageChannels:
		nrOfStageChannelsStr := strconv.Itoa(guildStat.NrOfStageChannels)
		return "Stage Channels - " + nrOfStageChannelsStr
	case consts.Categories:
		nrOfCategoriesStr := strconv.Itoa(guildStat.NrOfCategories)
		return "Categories - " + nrOfCategoriesStr
	case consts.AnnouncementChannels:
		nrOfAnnouncementChannelsStr := strconv.Itoa(guildStat.NrOfAnnouncementChannels)
		return "Announcement Channels - " + nrOfAnnouncementChannelsStr
	case consts.Emojis:
		nrOfEmojisStr := strconv.Itoa(guildStat.NrOfEmojis)
		return "Emojis - " + nrOfEmojisStr
	case consts.StaticEmojis:
		nrOfStaticEmojisStr := strconv.Itoa(guildStat.NrOfStaticEmojis)
		return "Static Emojis - " + nrOfStaticEmojisStr
	case consts.AnimatedEmojis:
		nrOfAnimatedEmojisStr := strconv.Itoa(guildStat.NrOfAnimatedEmojis)
		return "Animated Emojis - " + nrOfAnimatedEmojisStr
	case consts.Stickers:
		nrOfStickersStr := strconv.Itoa(guildStat.NrOfStickers)
		return "Stickers - " + nrOfStickersStr
	case consts.CustomStickers:
		nrOfCustomStickersStr := strconv.Itoa(guildStat.NrOfCustomStickers)
		return "Custom Stickers - " + nrOfCustomStickersStr
	case consts.ServerStickers:
		nrOfServerStickersStr := strconv.Itoa(guildStat.NrOfServerStickers)
		return "Server Stickers - " + nrOfServerStickersStr
	case consts.Roles:
		nrOfRolesStr := strconv.Itoa(guildStat.NrOfRoles)
		return "Roles - " + nrOfRolesStr
	// case consts.HighestTicker:
	// 	symbol := coinData[0]
	// 	interval := coinData[1]
	// 	highest := coinData[2]
	// 	if i, _ := strconv.Atoi(interval); i > 1 {
	// 		return "Top ticker - " + symbol + " - " + interval + " days - " + highest
	// 	}
	// 	return "Top ticker - " + symbol + " - " + interval + " day - " + highest
	default:
		nrOfMembersStr := strconv.Itoa(guildStat.NrOfMembers)
		return "Members - " + nrOfMembersStr
	}
}
