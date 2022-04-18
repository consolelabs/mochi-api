package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GuildUser struct {
	ID        uuid.NullUUID  `json:"id" gorm:"default:uuid_generate_v4()"`
	GuildID   string         `json:"guild_id" gorm:"unique_index:idx_guild_user_guild_id_user_id"`
	UserID    string         `json:"user_id" gorm:"uique_index:idx_guild_user_guild_id_user_id"`
	Nickname  JSONNullString `json:"nickname"`
	InvitedBy JSONNullInt64  `json:"invited_by"`
	Roles     []GuildRole    `json:"roles" gorm:"many2many:guild_user_role;foreignKey:UserID;joinForeignKey:UserID;References:ID;joinReferences:RoleID"`
}

type GuildUserRole struct {
	ID      uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()"`
	GuildID int64         `json:"guild_id"`
	UserID  int64         `json:"user_id"`
	RoleID  int64         `json:"role_id"`
}

func (u *GuildUser) BeforeCreate(tx *gorm.DB) (err error) {
	cols := []clause.Column{}
	colsNames := []string{}
	for _, field := range tx.Statement.Schema.PrimaryFields {
		cols = append(cols, clause.Column{Name: field.DBName})
		colsNames = append(colsNames, field.DBName)
	}

	tx.Statement.AddClause(clause.OnConflict{
		Columns:   cols,
		DoNothing: true,
	})

	tx.Statement.AddClause(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "guild_id"},
			{Name: "user_id"},
		},
		DoUpdates: []clause.Assignment{
			{Column: clause.Column{Name: "invited_by"}, Value: u.InvitedBy},
		},
	})

	return nil
}
