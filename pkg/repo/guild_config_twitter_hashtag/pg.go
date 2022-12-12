package guild_config_twitter_hashtag

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) UpsertOne(hashtag *model.GuildConfigTwitterHashtag) error {
	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "guild_id"}},
		UpdateAll: true,
	}).Create(&hashtag).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
func (pg *pg) DeleteByGuildID(guildID string) error {
	hashtag := model.GuildConfigTwitterHashtag{
		GuildID: guildID,
	}
	return pg.db.Table("guild_config_twitter_hashtags").Where(fmt.Sprintf("guild_id='%s'", guildID)).Delete(hashtag).Error
}

func (pg *pg) GetByGuildID(guildID string) (*model.GuildConfigTwitterHashtag, error) {
	hashtag := model.GuildConfigTwitterHashtag{}
	err := pg.db.Table("guild_config_twitter_hashtags").Where(fmt.Sprintf("guild_id='%s'", guildID)).First(&hashtag)
	if err.Error != nil {
		return nil, err.Error
	}
	return &hashtag, nil
}

func (pg *pg) GetAll() ([]model.GuildConfigTwitterHashtag, error) {
	hashtags := []model.GuildConfigTwitterHashtag{}
	err := pg.db.Table("guild_config_twitter_hashtags").Find(&hashtags)
	if err.Error != nil {
		return nil, err.Error
	}
	return hashtags, nil
}
