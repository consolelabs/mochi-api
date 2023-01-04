package onchaintipbottransaction

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	List(q ListQuery) ([]model.OnchainTipBotTransaction, error)
	GetOne(ID int) (*model.OnchainTipBotTransaction, error)
	UpsertMany([]*model.OnchainTipBotTransaction) error
}
