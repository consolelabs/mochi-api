package userdefiactivitylog

import (
	"database/sql"
	"time"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{
		db: db,
	}
}

func (pg *pg) CreateTransferLogs(req request.TransferRequest, tokenID int, amountEach, totalAmount float64) error {
	log := &model.UserDefiActivityLog{
		Type:      model.DefiType(req.TransferType),
		Role:      model.SENDER,
		TokenID:   tokenID,
		UserID:    req.Sender,
		Amount:    totalAmount,
		CreatedAt: time.Now(),
	}
	if req.GuildID != "" {
		log.GuildID = model.JSONNullString{NullString: sql.NullString{String: req.GuildID, Valid: true}}
	}

	tx := pg.db.Begin()
	if err := tx.Create(log).Error; err != nil {
		tx.Rollback()
		return err
	}
	// withdraw then no recipients
	if req.TransferType == request.WITHDRAW {
		return tx.Commit().Error
	}

	for _, recipient := range req.Recipients {
		log.ID.UUID = uuid.New()
		log.UserID = recipient
		log.Role = model.RECIPIENT
		log.Amount = amountEach
		err := tx.Create(log).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
