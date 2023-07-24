package userfeedback

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/util"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{
		db: db,
	}
}

func (pg *pg) CreateOne(feedback *model.UserFeedback) (*model.UserFeedback, error) {
	return feedback, pg.db.Create(&feedback).Error
}

// All response should be sorted by status and create time
// Only returns data in the past 2 months
func (pg *pg) GetAll(page int, size int) ([]model.UserFeedback, int64, error) {
	var fb = []model.UserFeedback{}
	var count int64
	return fb, count, pg.db.
		Table("user_feedbacks").
		Where("created_at > (CURRENT_DATE - INTERVAL '60 days')").
		Count(&count).
		Limit(size).
		Offset(size * page).
		Order("status ASC").
		Order("created_at DESC").
		Find(&fb).Error
}
func (pg *pg) GetAllByStatus(status string, page int, size int) ([]model.UserFeedback, int64, error) {
	var fb = []model.UserFeedback{}
	var count int64
	return fb, count, pg.db.
		Table("user_feedbacks").
		Where("status=? AND created_at > (CURRENT_DATE - INTERVAL '60 days')", status).
		Count(&count).
		Limit(size).
		Offset(size * page).
		Order("created_at DESC").
		Find(&fb).Error
}
func (pg *pg) GetAllByCommand(command string, page int, size int) ([]model.UserFeedback, int64, error) {
	var fb = []model.UserFeedback{}
	var count int64
	return fb, count, pg.db.
		Table("user_feedbacks").
		Where("command=? AND created_at > (CURRENT_DATE - INTERVAL '60 days')", command).
		Count(&count).
		Limit(size).
		Offset(size * page).
		Order("status ASC").
		Order("created_at DESC").
		Find(&fb).Error
}
func (pg *pg) GetAllByDiscordID(id string, page int, size int) ([]model.UserFeedback, int64, error) {
	var fb = []model.UserFeedback{}
	var count int64
	return fb, count, pg.db.
		Table("user_feedbacks").
		Where("discord_id=? AND created_at > (CURRENT_DATE - INTERVAL '60 days')", id).
		Count(&count).
		Limit(size).
		Offset(size * page).
		Order("status ASC").
		Order("created_at DESC").
		Find(&fb).Error
}
func (pg *pg) UpdateStatusByID(id string, status string) (*model.UserFeedback, error) {
	fb := model.UserFeedback{ID: util.GetNullUUID(id)}
	return &fb, pg.db.Table("user_feedbacks").Model(&fb).Clauses(clause.Returning{}).Where("id=?", id).Updates(map[string]interface{}{"status": status, fmt.Sprintf("%s_at", status): time.Now().UTC()}).Error
}

func (pg *pg) List(q FeedbackQuery) (feedbacks []model.UserFeedback, total int64, err error) {
	db := pg.db.Table("user_feedbacks")
	if q.ProfileID != "" {
		db = db.Where("profile_id = ?", q.ProfileID)
	}

	if q.DiscordId != "" {
		db = db.Where("discord_id = ?", q.DiscordId)
	}

	if q.Status != "" {
		db = db.Where("status = ?", q.Status)
	}

	if q.Command != "" {
		db = db.Where("command = ?", q.Command)
	}

	if q.Sort != "" {
		db = db.Order(q.Sort)
	}

	db.Count(&total)
	if q.Offset != 0 {
		db = db.Offset(int(q.Offset))
	}
	if q.Limit != 0 {
		db = db.Limit(int(q.Limit))
	}

	return feedbacks, total, db.Find(&feedbacks).Error
}
