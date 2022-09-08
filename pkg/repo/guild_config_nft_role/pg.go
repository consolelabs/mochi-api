package guild_config_nft_role

import (
	"fmt"

	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) *pg {
	return &pg{
		db: db,
	}
}

func (pg *pg) GetMemberCurrentRoles(guildID string) ([]model.MemberNFTRole, error) {
	var urs []model.MemberNFTRole

	rows, err := pg.db.Raw(`
	select DISTINCT ON (user_discord_id, guild_config_nft_roles.nft_collection_id, guild_config_nft_roles.token_id) user_discord_id, role_id from guild_config_nft_roles
inner join user_wallets on user_wallets.guild_id = guild_config_nft_roles.guild_id
inner join user_nft_balances on user_nft_balances.nft_collection_id = guild_config_nft_roles.nft_collection_id and user_nft_balances.user_address = user_wallets.address
where balance >= number_of_tokens and guild_config_nft_roles.guild_id = ?
order by user_discord_id, guild_config_nft_roles.nft_collection_id, guild_config_nft_roles.token_id, number_of_tokens desc
	`, guildID).Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		ur := model.MemberNFTRole{}
		if err := rows.Scan(&ur.UserID, &ur.RoleID); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		urs = append(urs, ur)
	}

	return urs, nil
}

func (pg *pg) ListByGuildID(guildID string) ([]model.GuildConfigNFTRole, error) {
	var configs []model.GuildConfigNFTRole
	return configs, pg.db.Where("guild_id = ?", guildID).Order("nft_collection_id, token_id, number_of_tokens asc").Find(&configs).Error
}

func (pg *pg) GetByRoleID(guildID, roleID string) (*model.GuildConfigNFTRole, error) {
	config := &model.GuildConfigNFTRole{}
	return config, pg.db.Where("guild_id = ? AND role_id = ?", guildID, roleID).First(config).Error
}

func (pg *pg) UpsertOne(config *model.GuildConfigNFTRole) error {
	tx := pg.db.Begin()

	if err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "guild_id"},
			{Name: "role_id"},
		},
		UpdateAll: true,
	}).Create(config).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (pg *pg) Update(config *model.GuildConfigNFTRole) error {
	return pg.db.Save(config).Error
}

func (pg *pg) Delete(id string) error {
	return pg.db.Delete(&model.GuildConfigNFTRole{}, "id = ?", id).Error
}

func (pg *pg) Create(config model.GuildConfigNFTRole) (*model.GuildConfigNFTRole, error) {
	return &config, pg.db.Table("guild_config_nft_roles").Create(&config).Error
}

func (pg *pg) DeleteByGroupId(groupNFTRoleId string) error {
	return pg.db.Delete(&model.GuildConfigNFTRole{}, "group_id = ?", groupNFTRoleId).Error
}

func (pg *pg) DeleteByIds(ids []string) error {
	return pg.db.Delete(&model.GuildConfigNFTRole{}, "id IN ?", ids).Error
}
