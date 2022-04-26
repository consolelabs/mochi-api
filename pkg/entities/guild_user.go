package entities

import (
	"fmt"

	"github.com/defipod/mochi/pkg/model"
)

func (e *Entity) CreateGuildUserIfNotExists(guildID, userID, nickname string) error {
	guildUser := &model.GuildUser{
		GuildID:   guildID,
		UserID:    userID,
		Nickname:  nickname,
		InvitedBy: "",
	}

	if err := e.repo.GuildUsers.FirstOrCreate(guildUser); err != nil {
		return fmt.Errorf("failed to create if not exists guild user: %w", err)
	}

	return nil
}
