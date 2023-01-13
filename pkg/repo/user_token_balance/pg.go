package user_token_balance

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{
		db: db,
	}
}

func (pg *pg) Upsert(balance model.UserTokenBalance) error {
	tx := pg.db.Begin()

	if err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "user_address"},
			{Name: "chain_type"},
			{Name: "token_id"},
		},
		UpdateAll: true,
	}).Create(&balance).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// user token balances for all collections in 1 guild
func (pg *pg) GetUserTokenBalancesByUserInGuild(guildID string) ([]model.MemberTokenRole, error) {
	var res []model.MemberTokenRole
	rows, err := pg.db.Raw(`
	SELECT DISTINCT ON (temp.user_discord_id)
		user_discord_id, role_id
	FROM (
		SELECT
			final_balance.user_discord_id AS user_discord_id,
			final_balance.balance AS balance,
			final_config.id AS config_id,
			final_config.role_id AS role_id,
			final_config.required_amount AS required_amount
		FROM (
			SELECT
				bals.user_address,
				bals.token_id,
				wall.user_discord_id,
				bals.balance
			FROM
				user_token_balances AS bals
				INNER JOIN user_wallets AS wall ON wall.address = bals.user_address
			WHERE
				wall.guild_id = ?) AS final_balance
			INNER JOIN (
				SELECT
					config.id, config.token_id, config.guild_id, config.role_id, config.required_amount
				FROM
					guild_config_token_roles AS config
				WHERE
					guild_id = ?) AS final_config ON final_config.token_id = final_balance.token_id
			GROUP BY
				user_discord_id,
				config_id,
				balance,
				role_id,
				required_amount) AS temp
	WHERE
		COALESCE(temp.balance, 0) >= temp.required_amount
	ORDER BY
		temp.user_discord_id,
		temp.required_amount DESC;
	`, guildID, guildID).Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		tmp := model.MemberTokenRole{}
		if err := rows.Scan(&tmp.UserDiscordID, &tmp.RoleID); err != nil {
			return nil, err
		}
		res = append(res, tmp)
	}

	return res, nil
}
