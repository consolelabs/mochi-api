package guild_config_repost_reaction

import (
	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	GetByGuildID(guildID string) ([]model.GuildConfigRepostReaction, error)
	GetByReaction(guildID string, reaction string) (model.GuildConfigRepostReaction, error)
	GetByReactionStartOrStop(guildID, emoji string) (model.GuildConfigRepostReaction, error)
	GetByReactionConversationStartOrStop(guildID, emoji string) (model.GuildConfigRepostReaction, error)
	GetByRepostChannelID(guildID string, channelID string) (model.GuildConfigRepostReaction, error)
	UpsertOne(config model.GuildConfigRepostReaction) error
	DeleteOne(guildID string, emoji string) error
}
