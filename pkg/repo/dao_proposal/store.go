package dao_proposal

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetById(id int64) (model *model.DaoProposal, err error) {
	return model, pg.db.First(&model, id).Error
}

func (pg *pg) GetAllByCreatorId(userId string) (models *[]model.DaoProposal, err error) {
	return models, pg.db.Where("creator_id = ?", userId).Find(&models).Error
}
func (pg *pg) GetUsageStatsWithPaging(page int, size int) (models *[]response.ProposalCount, total int64, err error) {
	return models, total, pg.db.Table("dao_proposal").
		Count(&total).
		Select("guild_id, discord_guilds.name as guild_name, COUNT(guild_id)").
		Joins("JOIN discord_guilds on discord_guilds.id = guild_id").
		Group("guild_id, discord_guilds.name").
		Offset(size * page).
		Limit(size).
		Order("count DESC").
		Scan(&models).Error
}

func (pg *pg) GetAllByGuildId(guildId string) (models *[]model.DaoProposal, err error) {
	return models, pg.db.Where("guild_id = ?", guildId).Find(&models).Error
}
func (pg *pg) GetByCreatorIdAndProposalId(proposal int64, userId string) (models []model.DaoProposalWithView, err error) {
	rows, err := pg.db.Table("dao_proposal").Joins("join view_dao_proposal ON id = ? AND creator_id = ? AND view_dao_proposal.proposal_id = ?", proposal, userId, proposal).Rows()
	for rows.Next() {
		tmp := model.DaoProposalWithView{}
		if err := rows.Scan(&tmp.Id, &tmp.GuildId, &tmp.GuildConfigDaoProposalId, &tmp.VotingChannelId, &tmp.DiscussionChannelId, &tmp.CreatorId, &tmp.Title, &tmp.Description, &tmp.CreatedAt, &tmp.UpdatedAt, &tmp.ClosedAt, &tmp.Sum, &tmp.Choice, &tmp.ProposalID, &tmp.GuildId); err != nil {
			return nil, err
		}
		models = append(models, tmp)
	}
	return models, err
}

func (pg *pg) Create(model *model.DaoProposal) (*model.DaoProposal, error) {
	return model, pg.db.Create(&model).Error
}

func (pg *pg) UpdateDiscussionChannel(id int64, discussionChannelId string) error {
	return pg.db.Model(&model.DaoProposal{}).Where("id = ?", id).Update("discussion_channel_id", discussionChannelId).Error
}

func (pg *pg) DeleteById(id int64) error {
	return pg.db.Delete(&model.DaoProposal{}, id).Error
}
