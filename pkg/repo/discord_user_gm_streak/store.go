package discord_user_gm_streak

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	UpsertOne(streak model.DiscordUserGMStreak) error
	UpsertBatch(streaks []model.DiscordUserGMStreak) error
	GetByDiscordIDGuildID(discordID, guildID string) (*model.DiscordUserGMStreak, error)
	GetAll() ([]model.DiscordUserGMStreak, error)
}
