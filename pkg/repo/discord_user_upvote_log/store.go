package discord_user_upvote_log

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	UpsertOne(log model.DiscordUserUpvoteLog) error
	GetByDiscordID(discordID string) ([]model.DiscordUserUpvoteLog, error)
}
