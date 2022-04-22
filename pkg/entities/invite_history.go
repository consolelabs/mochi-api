package entities

import (
	"fmt"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) CreateInviteHistory(req request.CreateInviteHistoryRequest) error {
	inviteHistory := &model.InviteHistory{
		GuildID:   req.GuildID,
		UserID:    req.Invitee,
		InvitedBy: req.Inviter,
		Type:      req.Type,
	}

	if err := e.repo.InviteHistories.Create(inviteHistory); err != nil {
		return fmt.Errorf("failed to create invite history: %w", err)
	}

	if err := e.repo.GuildUsers.Update(req.GuildID, req.Invitee, "invited_by", req.Inviter); err != nil {
		return fmt.Errorf("failed to update guild user: %w", err)
	}

	return nil
}

func (e *Entity) CountInviteHistoriesByGuildUser(guildID, userID string) (int64, error) {
	count, err := e.repo.GuildUsers.CountByGuildUser(guildID, userID)
	if err != nil {
		return 0, fmt.Errorf("failed to count invite histories: %w", err)
	}

	return count, nil
}
