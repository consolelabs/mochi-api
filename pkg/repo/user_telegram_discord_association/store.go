package usertelegramdiscordassociation

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetOneByTelegramUsername(telegramID string) (*model.UserTelegramDiscordAssociation, error)
	Upsert(*model.UserTelegramDiscordAssociation) error
}
