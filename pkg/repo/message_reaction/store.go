package message_reaction

import (
	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	Create(record model.MessageReaction) error
	GetByMessageID(messageID string) ([]model.MessageReaction, error)
	Delete(messageID string, userID string, reaction string) error
	DeleteByMessageID(messageID string) error
}
