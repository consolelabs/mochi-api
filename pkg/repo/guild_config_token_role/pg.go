package guild_config_token_role

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

func (pg *pg) Create(config *model.GuildConfigTokenRole) error {
	return pg.db.Create(config).Error
}

func (pg *pg) Get(id int) (model *model.GuildConfigTokenRole, err error) {
	return model, pg.db.First(&model, id).Error
}

func (pg *pg) ListByGuildID(guildID string) ([]model.GuildConfigTokenRole, error) {
	var configs []model.GuildConfigTokenRole
	return configs, pg.db.Preload("Token.Chain").Where("guild_id = ?", guildID).Order("token_id, required_amount asc").Find(&configs).Error
}

func (pg *pg) Update(config *model.GuildConfigTokenRole) error {
	return pg.db.Save(config).Error
}

func (pg *pg) Delete(id int) error {
	return pg.db.Delete(&model.GuildConfigTokenRole{}, "id = ?", id).Error
}

func (pg *pg) ListAllTokenConfigs(guildID string) ([]model.Token, error) {
	var tokens []model.Token
	query := pg.db.Table("guild_config_token_roles").
		Distinct("guild_config_token_roles.token_id").
		Select("guild_config_token_roles.token_id, tokens.name, tokens.address, tokens.decimals, tokens.chain_id").
		Joins("INNER JOIN tokens ON guild_config_token_roles.token_id = tokens.id").
		Group("tokens.id, guild_config_token_roles.id")

	if guildID != "" {
		query = query.Where("guild_config_token_roles.guild_id = ?", guildID)
	}
	rows, err := query.Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		tmp := model.Token{}
		if err := rows.Scan(&tmp.ID, &tmp.Name, &tmp.Address, &tmp.Decimals, &tmp.ChainID); err != nil {
			return nil, err
		}
		tokens = append(tokens, tmp)
	}

	return tokens, nil
}

func (pg *pg) ListConfigGuildIds() ([]string, error) {
	var guildIds []string
	rows, err := pg.db.Table("guild_config_token_roles").
		Distinct("guild_id").
		Select("guild_id").
		Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var tmp string
		if err := rows.Scan(&tmp); err != nil {
			return nil, err
		}
		guildIds = append(guildIds, tmp)
	}

	return guildIds, nil
}
