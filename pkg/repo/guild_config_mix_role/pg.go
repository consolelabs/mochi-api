package guild_config_mix_role

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) *pg {
	return &pg{
		db: db,
	}
}

func (pg *pg) Create(config *model.GuildConfigMixRole) error {
	return pg.db.Create(config).Error
}

func (pg *pg) Get(id int) (model *model.GuildConfigMixRole, err error) {
	return model, pg.db.First(&model, id).Error
}

func (pg *pg) ListByGuildID(guildID string) ([]model.GuildConfigMixRole, error) {
	var configs []model.GuildConfigMixRole
	return configs, pg.db.
		Where("guild_id = ?", guildID).
		Preload("TokenRequirement.Token").
		Preload("NFTRequirement.NFTCollection").
		Joins("LEFT JOIN mix_role_nft_requirements as nft_req on nft_req.id = guild_config_mix_roles.nft_requirement_id").
		Joins("LEFT JOIN mix_role_token_requirements as token_req on token_req.id = guild_config_mix_roles.token_requirement_id").
		Order("COALESCE(nft_req.required_amount, 0), COALESCE(token_req.required_amount, 0), required_level ASC").
		Find(&configs).Error
}

func (pg *pg) GetByRoleID(guildID, roleID string) (*model.GuildConfigMixRole, error) {
	config := &model.GuildConfigMixRole{}
	return config, pg.db.Where("guild_id = ? AND role_id = ?", guildID, roleID).First(config).Error
}

func (pg *pg) Update(config *model.GuildConfigMixRole) error {
	return pg.db.Save(config).Error
}

func (pg *pg) Delete(id int) error {
	return pg.db.Delete(&model.GuildConfigMixRole{}, "id = ?", id).Error
}

func (pg *pg) GetMemberCurrentRoles(guildID string) ([]model.MemberMixRole, error) {
	var urs []model.MemberMixRole
	rows, err := pg.db.Raw(`
	SELECT DISTINCT ON (user_xp.user_id)
	user_xp.user_id,
	config.role_id
	FROM
		guild_config_mix_roles AS config
		INNER JOIN guild_user_xps AS user_xp ON user_xp.guild_id = config.guild_id
		LEFT JOIN mix_role_nft_requirements AS nft_req ON nft_req.id = config.nft_requirement_id
		LEFT JOIN mix_role_token_requirements AS token_req ON token_req.id = config.token_requirement_id
		LEFT JOIN (
			SELECT
				*
			FROM
				user_token_balances
				INNER JOIN user_wallets ON user_wallets.address = user_token_balances.user_address) AS user_token_wallet ON user_token_wallet.guild_id = config.guild_id
		LEFT JOIN (
			SELECT
				*
			FROM
				user_nft_balances
				INNER JOIN user_wallets ON user_wallets.address = user_nft_balances.user_address) AS user_nft_wallet ON user_nft_wallet.guild_id = config.guild_id
	WHERE
		user_xp. "level" >= config.required_level
		AND(
			CASE WHEN token_requirement_id IS NOT NULL THEN
				token_req.token_id = user_token_wallet.token_id
				AND user_token_wallet.balance >= token_req.required_amount
			ELSE
				TRUE
			END)
		AND(
			CASE WHEN nft_requirement_id IS NOT NULL THEN
				nft_req.nft_collection_id = user_nft_wallet.nft_collection_id
				AND user_nft_wallet.balance >= nft_req.required_amount
			ELSE
				TRUE
			END)
		AND config.guild_id = ?
	ORDER BY
		user_xp.user_id,
		config.required_level DESC,
		COALESCE(nft_req.required_amount, 0) DESC,
		COALESCE(token_req.required_amount, 0) DESC;
	`, guildID).Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		ur := model.MemberMixRole{}
		if err := rows.Scan(&ur.UserDiscordID, &ur.RoleID); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		urs = append(urs, ur)
	}

	return urs, nil
}
