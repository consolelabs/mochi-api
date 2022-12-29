package dao_vote_option

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetById(id int64) (*model.DaoVoteOption, error)
}
