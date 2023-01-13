package user_token_balance

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Upsert(balance model.UserTokenBalance) error
	GetUserTokenBalancesByUserInGuild(guildID string) ([]model.MemberTokenRole, error)
}
