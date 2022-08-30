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
func (pg *pg) GetUserNFTBalancesByGuild(nftCollectionIds []string, guildID string) ([]model.UserNFTBalancesByGuild, error) {
	var res []model.UserNFTBalancesByGuild
	rows, err := pg.db.Raw(`
	SELECT
		user_discord_id,
		SUM(list_balance.balance) AS total_balance
	FROM (
		SELECT
			user_discord_id,
			user_address,
			nft_collection_id,
			balance
		FROM
			user_nft_balances
			INNER JOIN user_wallets ON user_nft_balances.user_address = user_wallets.address
		WHERE
			user_wallets.guild_id = ?
			AND user_nft_balances.balance > 0 AND user_nft_balances.nft_collection_id IN ?
			) AS list_balance
	GROUP BY
		user_discord_id;
	`, guildID, nftCollectionIds).Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		tmp := model.UserNFTBalancesByGuild{}
		if err := rows.Scan(&tmp.UserDiscordId, &tmp.TotalBalance); err != nil {
			return nil, err
		}
		res = append(res, tmp)
	}
	return res, nil
}
