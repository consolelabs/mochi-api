package invite_histories

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

func (pg *pg) Create(invite *model.InviteHistory) error {
	return pg.db.Create(invite).Error
}

func (pg *pg) CountByInviter(inviterID int64) (int64, error) {
	var count int64
	err := pg.db.Model(&model.InviteHistory{}).Where("inviter_id = ?", inviterID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
