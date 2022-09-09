package user_nft_balance

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{
		db: db,
	}
}

func (pg *pg) Upsert(balance model.UserNFTBalance) error {
	tx := pg.db.Begin()

	if err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "user_address"},
			{Name: "chain_type"},
			{Name: "nft_collection_id"},
			{Name: "token_id"},
		},
		UpdateAll: true,
	}).Create(&balance).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// user nft balances for all collections in 1 guild
func (pg *pg) GetUserNFTBalancesByUserInGuild(guildID string) ([]model.MemberNFTRole, error) {
	var res []model.MemberNFTRole
	rows, err := pg.db.Raw(`
	SELECT DISTINCT ON (temp.user_discord_id)
		user_discord_id, role_id
	FROM (
		SELECT
			final_balance.user_discord_id AS user_discord_id,
			sum(final_balance.balance) AS total_balance,
			final_config.id AS group_id,
			final_config.role_id AS role_id,
			final_config.number_of_tokens AS number_of_tokens
		FROM (
			SELECT
				bals.user_address,
				bals.nft_collection_id,
				wall.user_discord_id,
				bals.balance
			FROM
				user_nft_balances AS bals
				INNER JOIN user_wallets AS wall ON wall.address = bals.user_address
			WHERE
				wall.guild_id = ?) AS final_balance
			INNER JOIN (
				SELECT
					a.id, config.nft_collection_id, a.guild_id, a.role_id, a.number_of_tokens
				FROM
					guild_config_nft_roles AS config
					INNER JOIN guild_config_group_nft_roles AS a ON config.group_id = a.id
				WHERE
					guild_id = ?) AS final_config ON final_config.nft_collection_id = final_balance.nft_collection_id
			GROUP BY
				user_discord_id,
				group_id,
				role_id,
				number_of_tokens) AS temp
	WHERE
		temp.total_balance > temp.number_of_tokens
	ORDER BY
		temp.user_discord_id,
		temp.number_of_tokens DESC;
	`, guildID, guildID).Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		tmp := model.MemberNFTRole{}
		if err := rows.Scan(&tmp.UserDiscordID, &tmp.RoleID); err != nil {
			return nil, err
		}
		res = append(res, tmp)
	}

	return res, nil
}
