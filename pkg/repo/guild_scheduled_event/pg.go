package guild_scheduled_event

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) UpsertOne(config *model.GuildScheduledEvent) error {
	tx := pg.db.Begin()

	// update on conflict
	err := tx.Table("guild_scheduled_events").Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "guild_id"},
			{Name: "event_id"},
		},
		DoUpdates: clause.AssignmentColumns([]string{"status"}),
	}).Create(&config).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (pg *pg) ListUncompleteByGuildID(guildID string) ([]model.GuildScheduledEvent, error) {
	var events []model.GuildScheduledEvent

	return events, pg.db.Where("guild_id = ? and status in (?)", guildID, []discordgo.GuildScheduledEventStatus{
		discordgo.GuildScheduledEventStatusScheduled,
		discordgo.GuildScheduledEventStatusActive,
	}).Find(&events).Error
}
