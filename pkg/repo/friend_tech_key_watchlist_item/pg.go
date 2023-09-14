package friend_tech_key_watchlist_item

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) Create(item model.FriendTechKeyWatchlistItem) (*model.FriendTechKeyWatchlistItem, error) {
	return &item, pg.db.Create(&item).Error
}

func (pg *pg) Update(item model.FriendTechKeyWatchlistItem) error {
	return pg.db.
		Table("friend_tech_key_watchlist_items").
		Where(
			"id = ?", item.Id,
		).Updates(
		map[string]interface{}{
			"increase_alert_at": item.IncreaseAlertAt,
			"decrease_alert_at": item.DecreaseAlertAt,
			"updated_at":        item.UpdatedAt,
		},
	).Error
}

func (pg *pg) Delete(id int) error {
	return pg.db.Delete(&model.FriendTechKeyWatchlistItem{}, "id = ?", id).Error
}

func (pg *pg) DeleteByAddressAndProfileId(address string, profileId string) error {
	return pg.db.Delete(&model.FriendTechKeyWatchlistItem{}, "key_address = ? AND profile_id = ?", address, profileId).Error
}

func (pg *pg) List(filter model.ListFriendTechKeysFilter) ([]model.FriendTechKeyWatchlistItem, error) {
	var items []model.FriendTechKeyWatchlistItem

	db := pg.db
	if filter.ProfileId != "" {
		db = db.Where("profile_id = ?", filter.ProfileId)
	}

	return items, db.Find(&items).Error
}

func (pg *pg) Exist(id int, address string, profileId string) (bool, error) {
	var count int64
	return count > 0, pg.db.Model(&model.FriendTechKeyWatchlistItem{}).Where("id = ? OR (key_address = ? AND profile_id = ?)", id, address, profileId).Count(&count).Error
}

func (pg *pg) Get(id int) (*model.FriendTechKeyWatchlistItem, error) {
	var item model.FriendTechKeyWatchlistItem
	return &item, pg.db.First(&item, "id = ?", id).Error
}
