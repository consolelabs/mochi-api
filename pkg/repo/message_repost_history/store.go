package message_repost_history

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

type Store interface {
	GetByMessageID(guildID, messageID string) (*model.MessageRepostHistory, error)
	Upsert(record model.MessageRepostHistory) error
	EditMessageRepost(req *request.EditMessageRepostRequest) error
}
