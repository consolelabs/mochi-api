package entities

import (
	"fmt"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) CreateInviteHistory(req request.CreateInviteHistoryRequest) error {
	inviteHistory := &model.InviteHistory{
		GuildID:   req.GuildID,
		UserID:    req.Invitee,
		InvitedBy: req.Inviter,
		Type:      req.Type,
	}
	err := e.repo.InviteHistories.Create(inviteHistory)
	if err != nil {
		e.log.Error(err, "[entity.CreateInviteHistory] repo.InviteHistories.Create() failed")
	}
	return err
}

func (e *Entity) CountInviteHistoriesByGuildUser(guildID, userID string) (int64, error) {
	count, err := e.repo.GuildUsers.CountByGuildUser(guildID, userID)
	if err != nil {
		return 0, fmt.Errorf("failed to count invite histories: %w", err)
	}

	return count, nil
}

func (e *Entity) GetInvitesLeaderboard(guildID string) ([]response.UserInvitesAggregation, error) {
	resp, err := e.repo.InviteHistories.GetInvitesLeaderboard(guildID)
	if err != nil {
		return nil, fmt.Errorf("failed to get invites leaderboard: %w", err)
	}

	return resp, err
}

func (e *Entity) GetUserInvitesAggregation(guildID, userID string) (*response.UserInvitesAggregation, error) {
	resp, err := e.repo.InviteHistories.GetUserInvitesAggregation(guildID, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user invites aggregation: %w", err)
	}

	return resp, err
}
