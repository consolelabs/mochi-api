package guild_config_xp_role

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) *pg {
	return &pg{
		db: db,
	}
}

func (pg *pg) Create(config *model.GuildConfigXPRole) error {
	return pg.db.Create(config).Error
}

func (pg *pg) Get(id int) (model *model.GuildConfigXPRole, err error) {
	return model, pg.db.First(&model, id).Error
}

func (pg *pg) ListByGuildID(guildID string) ([]model.GuildConfigXPRole, error) {
	var configs []model.GuildConfigXPRole
	return configs, pg.db.Where("guild_id = ?", guildID).Order("required_xp asc").Find(&configs).Error
}

func (pg *pg) GetByRoleID(guildID, roleID string) (*model.GuildConfigXPRole, error) {
	config := &model.GuildConfigXPRole{}
	return config, pg.db.Where("guild_id = ? AND role_id = ?", guildID, roleID).First(config).Error
}

func (pg *pg) Update(config *model.GuildConfigXPRole) error {
	return pg.db.Save(config).Error
}

func (pg *pg) Delete(id int) error {
	return pg.db.Delete(&model.GuildConfigXPRole{}, "id = ?", id).Error
}

func (pg *pg) GetMemberCurrentRoles(guildID string) ([]model.MemberXPRole, error) {
	var urs []model.MemberXPRole

	rows, err := pg.db.Raw(`
	SELECT DISTINCT ON (guild_user_xps.user_id)
		guild_user_xps.user_id,
		config.role_id
	FROM
		guild_config_xp_roles AS config
		INNER JOIN guild_user_xps ON guild_user_xps.guild_id = config.guild_id
	WHERE
		guild_user_xps.total_xp >= config.required_xp
		AND config.guild_id = ?
	ORDER BY
		guild_user_xps.user_id,
		config.required_xp DESC
	`, guildID).Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		ur := model.MemberXPRole{}
		if err := rows.Scan(&ur.UserID, &ur.RoleID); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		urs = append(urs, ur)
	}

	return urs, nil
}
