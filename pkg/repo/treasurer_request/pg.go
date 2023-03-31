package treasurerrequest

import (
	"time"

	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) Create(treasurerReq *model.TreasurerRequest) (*model.TreasurerRequest, error) {
	return treasurerReq, pg.db.Create(treasurerReq).Error
}

func (pg *pg) GetById(id int64) (treasurerReq *model.TreasurerRequest, err error) {
	return treasurerReq, pg.db.Where("id = ? and deleted_at is null", id).First(&treasurerReq).Error
}

func (p *pg) Delete(treasurerReq *model.TreasurerRequest) error {
	return p.db.Model(&model.TreasurerRequest{}).Where("guild_id = ? and vault_id = ? and user_discord_id = ?", treasurerReq.GuildId, treasurerReq.VaultId, treasurerReq.UserDiscordId).Update("deleted_at", time.Now()).Error
}
