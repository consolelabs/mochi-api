package user_nft_balance

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/defipod/mochi/pkg/consts"
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

func (pg *pg) Upsert(b model.UserNFTBalance) error {
	tx := pg.db.Begin()

	updates := []string{"balance", "updated_at", "profile_id"}
	if b.StakingNekos != 0 {
		updates = append(updates, "staking_nekos")
	}

	if err := tx.Clauses(clause.OnConflict{
		OnConstraint: "user_nft_balances_collection_id_address",
		DoUpdates:    clause.AssignmentColumns(updates),
	}).Create(&b).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// user nft balances for all collections in 1 guild
func (pg *pg) GetUserNFTBalancesByUserInGuild(guildID string) ([]model.UserAddressNFTBalancesByGuild, error) {
	var res []model.UserAddressNFTBalancesByGuild
	rows, err := pg.db.Raw(`
		SELECT DISTINCT ON (temp.user_address, nft_collection_id)
				user_address, total_balance, staking_neko, nft_collection_id
			FROM (
				SELECT
					final_balance.user_address AS user_address,
					sum(final_balance.balance) AS total_balance,
					final_config.id AS group_id,
					final_config.nft_collection_id AS nft_collection_id,
					final_config.role_id AS role_id,
					final_config.number_of_tokens AS number_of_tokens,
					final_balance.staking_nekos AS staking_neko
				FROM (
					SELECT
						bals.user_address,
						bals.nft_collection_id,
						bals.balance,
						bals.staking_nekos
					FROM
						user_nft_balances AS bals) AS final_balance
					INNER JOIN (
						SELECT
							a.id, config.nft_collection_id, a.guild_id, a.role_id, a.number_of_tokens
						FROM
							guild_config_nft_roles AS config
							INNER JOIN guild_config_group_nft_roles AS a ON config.group_id = a.id
						WHERE
							guild_id = ?) AS final_config ON final_config.nft_collection_id = final_balance.nft_collection_id
					GROUP BY
						user_address,
						group_id,
						final_config.nft_collection_id,
						role_id,
						number_of_tokens,
						staking_nekos) AS temp
			ORDER BY
				temp.user_address,
				temp.nft_collection_id,
				temp.number_of_tokens DESC
	`, guildID).Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		tmp := model.UserAddressNFTBalancesByGuild{}
		if err := rows.Scan(&tmp.UserAddress, &tmp.TotalBalance, &tmp.StakingNeko, &tmp.NftCollectionID); err != nil {
			return nil, err
		}
		res = append(res, tmp)
	}

	return res, nil
}

func (pg *pg) List(q ListQuery) ([]model.UserNFTBalance, error) {
	var res []model.UserNFTBalance
	db := pg.db
	if q.ProfileID != "" {
		db = db.Where("profile_id = ?", q.ProfileID)
	}
	if q.CollectionID != "" {
		db = db.Where("nft_collection_id = ?", q.CollectionID)
	}
	return res, db.Find(&res).Error
}

func (pg *pg) TotalBalance(collectionID string) (total int, err error) {
	return total, pg.db.Table("user_nft_balances").
		Select("COALESCE(SUM(balance + staking_nekos),0)").
		Where("nft_collection_id = ?", collectionID).
		Scan(&total).Error
}

func (pg *pg) IsExists(collectionID, userAddress string) (bool, error) {
	var count int64
	err := pg.db.Table("user_nft_balances").
		Where("nft_collection_id = ? AND user_address ILIKE ?", collectionID, userAddress).
		Count(&count).Error
	return count > 0, err
}

func (pg *pg) GetPodTownUserNFTBalances(collectionAddresses []string) ([]model.PodTownUserNFTBalance, error) {
	podTownGuildId := "882287783169896468"
	var data []model.PodTownUserNFTBalance
	return data, pg.db.Table("user_nft_balances b").
		Select(`
			b.user_address,
			b.profile_id, 
			sum(b.balance + staking_nekos) filter (where c.symbol ilike 'NEKO') as neko,
			sum(b.balance) filter (where c.symbol ilike 'RABBY') as rabby,
			sum(b.balance) filter (where c.symbol ilike 'FUKURO') as fukuro,
			max(aa.platform_identifier) discord_id,
			max(ugs.total_count) as gm
		`).
		Joins("LEFT JOIN nft_collections c ON b.nft_collection_id = c.id").
		Joins("LEFT JOIN associated_accounts aa ON b.profile_id = aa.profile_id and aa.platform = ?", consts.PlatformDiscord).
		Joins("LEFT JOIN discord_user_gm_streaks ugs ON aa.platform_identifier = ugs.discord_id AND ugs.guild_id = ?", podTownGuildId).
		Where("c.address IN ?", collectionAddresses).
		Group("b.user_address, b.profile_id").
		Find(&data).Error
}
