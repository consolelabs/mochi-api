package discordbottransaction

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(param model.DiscordBotTransaction) (*model.DiscordBotTransaction, error)
	Delete(id string) error
}
