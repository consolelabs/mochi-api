package guildconfiggroupnftrole

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) *pg {
	return &pg{
		db: db,
	}
}

func (pg *pg) Create(config model.GuildConfigGroupNFTRole) (*model.GuildConfigGroupNFTRole, error) {
	return &config, pg.db.Table("guild_config_group_nft_roles").Create(&config).Error
}

func (pg *pg) ListByGuildID(guildID string) ([]model.GuildConfigGroupNFTRole, error) {
	var configs []model.GuildConfigGroupNFTRole
	return configs, pg.db.Where("guild_id = ?", guildID).Preload("GuildConfigNFTRole").Find(&configs).Error
}

func (pg *pg) Delete(id string) error {
	return pg.db.Delete(&model.GuildConfigGroupNFTRole{}, "id = ?", id).Error
}

func (pg *pg) GetByRoleID(guildID, roleID string) (*model.GuildConfigGroupNFTRole, error) {
	config := &model.GuildConfigGroupNFTRole{}
	return config, pg.db.Where("guild_id = ? AND role_id = ?", guildID, roleID).First(config).Error
}
