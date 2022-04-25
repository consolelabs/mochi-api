package invite_histories

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
)

type Store interface {
	Create(invite *model.InviteHistory) error
	CountByInviter(inviterID int64) (int64, error)
	GetInvitesLeaderboard(guildID string) ([]response.LeaderboardRecord, error)
}
