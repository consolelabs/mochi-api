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
func (pg *pg) GetAll() ([]model.UserFeedback, error) {
	var fb = []model.UserFeedback{}
	return fb, pg.db.Find(&fb).Error
}
func (pg *pg) GetAllByStatus(status string) ([]model.UserFeedback, error) {
	var fb = []model.UserFeedback{}
	return fb, pg.db.Where("status=?", status).Find(&fb).Error
}
func (pg *pg) GetAllByCommand(command string) ([]model.UserFeedback, error) {
	var fb = []model.UserFeedback{}
	return fb, pg.db.Where("command=?", command).Find(&fb).Error
}
func (pg *pg) GetAllByDiscordID(id string) ([]model.UserFeedback, error) {
	var fb = []model.UserFeedback{}
	return fb, pg.db.Where("discord_id=?", id).Find(&fb).Error
}
func (pg *pg) UpdateStatusByID(id string, status string) (*model.UserFeedback, error) {
	fb := model.UserFeedback{ID: util.GetNullUUID(id)}
	return &fb, pg.db.Table("user_feedbacks").Model(&fb).Clauses(clause.Returning{}).Where("id=?", id).Updates(map[string]interface{}{"status": status, fmt.Sprintf("%s_at", status): time.Now().UTC()}).Error
}