package conversation_repost_histories

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Upsert(model model.ConversationRepostHistories) error
	Update(model *model.ConversationRepostHistories) error
	GetByGuildAndChannel(guildID, channelID string) (*model.ConversationRepostHistories, error)
}
