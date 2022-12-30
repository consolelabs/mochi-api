package entities

import (
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"gorm.io/gorm"
)

func (e *Entity) Test() {
	// e.repo.GuildConfigDaoProposal.GetById(1)

	// e.repo.DaoProposal.GetById(1)

	// e.repo.DaoVoteOption.GetById(1)

	// e.repo.DaoVote.GetById(1)

	// e.repo.DaoProposalVoteOption.GetById(1)
}

func (e *Entity) GetGuildConfigDaoProposalByGuildID(guildId string) (*model.GuildConfigDaoProposal, error) {
	config, err := e.repo.GuildConfigDaoProposal.GetByGuildId(guildId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		e.log.Fields(logger.Fields{"guildId": guildId}).Error(err, "[entity.GetGuildConfigDaoProposalByGuildID] e.repo.GuildConfigDaoProposal.GetByGuildId failed")
		return nil, err
	}
	return config, nil
}

func (e *Entity) DeleteGuildConfigDaoProposalByGuildID(req *request.DeleteGuildConfigDaoProposal) error {
	// delete from db
	channel, err := e.repo.GuildConfigDaoProposal.DeleteById(req.ID)
	if err != nil {
		e.log.Fields(logger.Fields{"id": req.ID}).Error(err, "[entity.DeleteGuildConfigDaoProposalByGuildID] e.repo.GuildConfigDaoProposal.DeleteByGuildId failed")
		return err
	}

	// delete guideline channel from discord client
	err = e.svc.Discord.DeleteChannel(channel.GuidelineChannelId)
	if err != nil {
		e.log.Fields(logger.Fields{"guild_id": channel.GuildId}).Error(err, "[entity.DeleteGuildConfigDaoProposalByGuildID] failed to delete guideline channel")
	}
	return nil
}
