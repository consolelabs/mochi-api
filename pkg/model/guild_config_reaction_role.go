package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type GuildConfigReactionRole struct {
	ID               uuid.NullUUID     `json:"id" gorm:"default:uuid_generate_v4()"`
	GuildID          string            `json:"guild_id"`
	ChannelID        string            `json:"user_id"`
	Author           JSONNullString    `json:"author"`
	AuthorAvatar     JSONNullString    `json:"author_avatar"`
	Title            JSONNullString    `json:"title"`
	TitleUrl         JSONNullString    `json:"title_url"`
	HeaderImage      JSONNullString    `json:"header_image"`
	Message          string            `json:"message"`
	FooterImage      string            `json:"footer_image"`
	FooterMessage    string            `json:"footer_message"`
	FooterImageSmall string            `json:"footer_image_small"`
	ReactionRoles    ReactionRoleArray `json:"reaction_roles"`
}

type ReactionRole struct {
	Reaction string `json:"reaction"`
	RoleID   int64  `json:"role_id"`
	RoleName string `json:"role_name"`
}

type ReactionRoleArray []ReactionRole

func (v ReactionRole) Value() (driver.Value, error) {
	if v.Reaction == "" && v.RoleID == 0 {
		return nil, nil
	}

	raw, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return string(raw), nil
}

func (j *ReactionRole) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch t := value.(type) {
	case []uint8:
		jsonData := value.([]uint8)
		if string(jsonData) == "null" {
			return nil
		}

		return json.Unmarshal(t, j)
	default:
		return fmt.Errorf("could not scan type %T into ReactionRole", t)
	}
}

func (c ReactionRoleArray) Value() (driver.Value, error) {
	if c == nil {
		return nil, nil
	}

	raw, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return string(raw), nil
}

func (c *ReactionRoleArray) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch t := value.(type) {
	case []uint8:
		return json.Unmarshal(value.([]uint8), c)
	default:
		return fmt.Errorf("Could not scan type %T into ReactionRoleArray", t)
	}
}
