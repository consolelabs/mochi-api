package message_repost_history

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByMessageID(guildID, messageID string) (model.MessageRepostHistory, error)
	CreateIfNotExist(config model.MessageRepostHistory) error
}
