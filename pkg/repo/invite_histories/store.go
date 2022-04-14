package invite_histories

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(invite *model.InviteHistory) error
	CountByInviter(inviterID int64) (int64, error)
}
