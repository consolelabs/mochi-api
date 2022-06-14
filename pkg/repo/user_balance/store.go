package userbalance

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetOne(string, int) (*model.UserBalance, error)
	GetUserBalances(string) ([]model.UserBalance, error)
}
