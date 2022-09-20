package discord_user_upvote_streak

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	UpsertOne(streak model.DiscordUserUpvoteStreak) error
	UpsertBatch(streak []model.DiscordUserUpvoteStreak) error
	GetByDiscordID(discordID string) (*model.DiscordUserUpvoteStreak, error)
	GetAll() ([]model.DiscordUserUpvoteStreak, error)
	GetTopByStreak() ([]model.DiscordUserUpvoteStreak, error)
	GetTopByTotal() ([]model.DiscordUserUpvoteStreak, error)
	GetGuildTopByStreak(guildId string) ([]model.DiscordUserUpvoteStreak, error)
	GetGuildTopByTotal(guildId string) ([]model.DiscordUserUpvoteStreak, error)
}
