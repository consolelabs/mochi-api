package guild_config_dao_proposal

import (
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

func (pg *pg) GetById(id int64) (model *model.GuildConfigDaoProposal, err error) {
	return model, pg.db.First(&model, id).Error
}

func (pg *pg) GetByGuildId(guildId string) (model *model.GuildConfigDaoProposal, err error) {
	return model, pg.db.Where("guild_id = ?", guildId).First(&model).Error
}
func (pg *pg) DeleteById(id string) (*model.GuildConfigDaoProposal, error) {
	var config = model.GuildConfigDaoProposal{}
	return &config, pg.db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&config, id).Error
}
