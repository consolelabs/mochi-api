package util

import (
	"strings"

	"github.com/defipod/mochi/pkg/model"
)

func IsMatchConfig(content string, sticker string, configs []model.GuildConfigGmGn) (bool, bool, bool) {
	isMatchMsg, isMatchEmoji, isMatchSticker := false, false, false
	for _, config := range configs {
		if strings.EqualFold(config.Msg, content) {
			isMatchMsg = true
		}
		if strings.EqualFold(config.Emoji, content) {
			isMatchEmoji = true
		}
		if config.Sticker == sticker {
			isMatchSticker = true
		}
	}

	return isMatchMsg, isMatchEmoji, isMatchSticker
}

func IsMatchChannel(channelId string, configs []model.GuildConfigGmGn) bool {
	for _, config := range configs {
		if channelId == config.ChannelID {
			return true
		}
	}
	return false
}
