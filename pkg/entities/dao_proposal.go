package entities

import (
	"strconv"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	errs "github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) CreateDaoProposal(req *request.CreateDaoProposalRequest) (*model.DaoProposal, error) {
	config, err := e.repo.GuildConfigDaoProposal.GetByGuildId(req.GuildId)
	if err != nil {
		e.log.Fields(logger.Fields{"GuildId": req.GuildId}).Error(err, "[entities.CreateDaoProposal][repo.GuildConfigDaoProposal.GetByGuildID] - failed to get guild config")
		return nil, err
	}

	daoProposal, err := e.repo.DaoProposal.Create(&model.DaoProposal{
		GuildId:                  req.GuildId,
		GuildConfigDaoProposalId: config.Id,
		VotingChannelId:          req.VotingChannelId,
		CreatorId:                req.CreatorId,
		Title:                    req.Title,
		Description:              req.Description,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entities.CreateDaoProposal][repo.DaoProposal.Create()] - failed to create dao proposal")
		return nil, err
	}

	proposalVoteOption := model.DaoProposalVoteOption{
		ProposalId:     daoProposal.Id,
		Address:        config.Address,
		ChainId:        config.ChainID,
		Symbol:         config.Symbol,
		RequiredAmount: config.RequiredAmount,
	}
	if req.VoteOption != nil {
		proposalVoteOption.VoteOptionId = req.VoteOption.Id
		proposalVoteOption.Address = req.VoteOption.Address
		proposalVoteOption.ChainId = req.VoteOption.ChainId
		proposalVoteOption.Symbol = req.VoteOption.Symbol
		proposalVoteOption.RequiredAmount = req.VoteOption.RequiredAmount
	}

	_, err = e.repo.DaoProposalVoteOption.Create(&proposalVoteOption)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entities.CreateDaoProposal][repo.DaoProposalVoteOption.Create()] - failed to create dao proposal vote option")
		return nil, err
	}

	discussionChannel, err := e.svc.Discord.CreateDiscussionChannelForProposal(req.GuildId, req.Title)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entities.CreateDaoProposal][svc.Discord.CreateDiscussionChannelForProposal()] - failed to create discussion channel")
		return nil, err
	}

	err = e.repo.DaoProposal.UpdateDiscussionChannel(daoProposal.Id, discussionChannel)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entities.CreateDaoProposal][repo.DaoProposal.UpdateDiscussionChannel()] - failed to update discussion channel")
		return nil, err
	}

	daoProposal.DiscussionChannelId = discussionChannel

	return daoProposal, nil
}

func (e *Entity) DeleteDaoProposal(proposalId string) error {
	proposalIdNumber, err := strconv.ParseInt(proposalId, 10, 64)
	if err != nil {
		e.log.Error(err, "[Entity][DeleteDaoProposal] proposal_id is invalid")
		return errs.ErrInvalidProposalID
	}

	proposal, err := e.repo.DaoProposal.GetById(proposalIdNumber)
	if err != nil {
		e.log.Fields(logger.Fields{"proposalId": proposalId}).Error(err, "[entities.DeleteDaoProposal][repo.DaoProposal.GetById] - failed to get DAO proposal")
		return err
	}

	err = e.repo.DaoProposalVoteOption.DeleteAllByProposalID(proposal.Id)
	if err != nil {
		e.log.Fields(logger.Fields{"proposalId": proposalId}).Error(err, "[entities.DeleteDaoProposal][repo.DaoProposalVoteOption.DeleteAllByProposalID()] - failed to clear dao proposal vote options")
		return err
	}

	err = e.repo.DaoProposal.DeleteById(proposal.Id)
	if err != nil {
		e.log.Fields(logger.Fields{"proposalId": proposalId}).Error(err, "[entities.DeleteDaoProposal][repo.DaoProposal.DeleteById()] - failed to delete dao proposal")
		return err
	}

	err = e.svc.Discord.DeleteChannel(proposal.DiscussionChannelId)
	if err != nil {
		e.log.Fields(logger.Fields{"proposalId": proposalId}).Error(err, "[entities.DeleteDaoProposal][svc.Discord.DeleteChannel()] - failed to delete discussion channel")
	}

	return nil
}
