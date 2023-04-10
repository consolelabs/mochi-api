package entities

import (
	"time"

	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/repo/user_tag"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) UpsertUserTag(req request.UpsertUserTag) (*model.UserTag, error) {
	tagme, err := e.repo.UserTag.UpsertOne(model.UserTag{
		UserId:          req.UserId,
		GuildId:         req.GuildId,
		MentionUsername: req.MentionUsername,
		MentionRole:     req.MentionRole,
		UpdatedAt:       time.Now(),
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.UpsertUserTag] e.repo.UserTag.UpsertOne failed")
		return nil, err
	}
	return tagme, nil
}

func (e *Entity) GetUserTag(userID, guildID string) (*model.UserTag, error) {
	tag, err := e.repo.UserTag.GetOne(user_tag.GetOneQuery{
		GuildID: guildID,
		UserID:  userID,
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		e.log.Fields(logger.Fields{"guildID": guildID, "userID": userID}).Error(err, "[entity.GetUserTag] e.repo.UserTag.GetOne failed")
		return nil, err
	}
	return tag, nil
}
