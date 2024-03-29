package entities

import (
	"errors"
	"strconv"

	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	errs "github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) CreateDaoVote(req request.CreateDaoVoteRequest) error {
	return e.repo.DaoVote.CreateDaoVote(&model.DaoVote{
		UserId:     req.UserID,
		ProposalId: req.ProposalID,
		Choice:     req.Choice,
		Point:      1,
	})
}

func (e *Entity) UpdateDaoVote(voteId string, req request.UpdateDaoVoteRequest) (*model.DaoVote, error) {
	if err := req.Choice.IsValid(); err != nil {
		e.log.Fields(logger.Fields{
			"choice": req.Choice,
		}).Error(err, "[Entity][UpdateDaoVote] invalid vote choice")
		return nil, err
	}

	voteIdNum, err := strconv.ParseInt(voteId, 10, 64)
	if err != nil {
		e.log.Fields(logger.Fields{
			"voteId": voteId,
		}).Error(err, "[Entity][UpdateDaoVote] convert voteId to int64 failed")
		return nil, errs.ErrInvalidVoteID
	}

	vote, err := e.repo.DaoVote.GetById(voteIdNum)
	if err != nil {
		e.log.Fields(logger.Fields{
			"voteId": voteIdNum,
		}).Error(err, "[Entity][UpdateDaoVote] repo.DaoVote.GetById failed")
		if err == gorm.ErrRecordNotFound {
			return nil, errs.ErrRecordNotFound
		}
		return nil, err
	}

	if vote.UserId != req.UserID {
		e.log.Fields(logger.Fields{
			"userID": req.UserID,
		}).Error(err, "[Entity][UpdateDaoVote] the update user is not the creator")
		return nil, errs.ErrInvalidDiscordUserID
	}

	vote.Choice = req.Choice
	updatedVote, err := e.repo.DaoVote.Update(vote)
	if err != nil {
		e.log.Fields(logger.Fields{
			"vote": vote,
		}).Error(err, "[Entity][UpdateDaoVote] repo.DaoVote.Update failed")
		return nil, err
	}
	return updatedVote, nil
}

func (e *Entity) GetAllDaoProposalByUserId(userId string) (*[]model.DaoProposal, error) {
	proposals, err := e.repo.DaoProposal.GetAllByCreatorId(userId)
	if err != nil {
		e.log.Error(err, "[Entity][GetAllDaoProposalByUserId] DaoProposal.GetAllByCreatorId failed")
		return nil, err
	}
	return proposals, nil
}

func (e *Entity) GetAllDaoProposalByGuild(guildId string) (*[]model.DaoProposal, error) {
	return e.repo.DaoProposal.GetAllByGuildId(guildId)
}

func (e *Entity) GetDaoVotesByUserId(userId string) (*[]model.DaoVote, error) {
	votes, err := e.repo.DaoVote.GetByUserId(userId)
	if err != nil {
		e.log.Error(err, "[Entity][GetDaoVotesByUserId] repo.DaoVote.GetByUserId failed")
		return nil, err
	}
	return votes, nil
}

func (e *Entity) GetDaoProposalVotes(proposalId, discordId string) (*response.GetAllDaoProposalVotes, error) {
	pId, err := strconv.Atoi(proposalId)
	if err != nil {
		e.log.Error(err, "[Entity][GetDaoVotesByProposalId] proposal id invalid")
		return nil, err
	}
	// get proposal
	proposal, err := e.repo.DaoProposal.GetByCreatorIdAndProposalId(int64(pId), discordId)
	if err != nil {
		e.log.Error(err, "[Entity][GetDaoVotesByProposalId] repo.DaoProposal.GetById failed")
		return nil, err
	}
	if len(proposal) == 0 {
		e.log.Info("[Entity][GetDaoVotesByProposalId] creator id not matched")
		return nil, errors.New("creator id not matched")
	}

	// get proposal's total points
	var points []model.DaoProposalVoteCount
	for _, p := range proposal {
		v := model.DaoProposalVoteCount{
			Sum:        p.Sum,
			Choice:     p.Choice,
			ProposalID: p.ProposalID,
			GuildId:    p.GuildId,
		}
		points = append(points, v)
	}

	// get proposal's votes
	votes, err := e.repo.DaoVote.GetByProposalId(int64(pId))
	if err != nil {
		e.log.Error(err, "[Entity][GetDaoVotesByProposalId] repo.DaoVote.GetByProposalId failed")
		return nil, err
	}

	return &response.GetAllDaoProposalVotes{
		Proposal: &response.GetDaoProposalData{
			Id:                       proposal[0].Id,
			GuildId:                  proposal[0].GuildId,
			GuildConfigDaoProposalId: proposal[0].GuildConfigDaoProposalId,
			VotingChannelId:          proposal[0].VotingChannelId,
			DiscussionChannelId:      proposal[0].DiscussionChannelId,
			CreatorId:                proposal[0].CreatorId,
			Title:                    proposal[0].Title,
			Points:                   &points,
			Description:              proposal[0].Description,
			CreatedAt:                proposal[0].CreatedAt,
			UpdatedAt:                proposal[0].UpdatedAt,
			ClosedAt:                 proposal[0].ClosedAt,
		},
		Votes: votes,
	}, nil
}

func (e *Entity) GetDaoProposalVoteOfUser(proposalId, userId string) (*model.DaoVote, error) {
	proposalIdNumber, err := strconv.ParseInt(proposalId, 10, 64)
	if err != nil {
		e.log.Error(err, "[Entity][GetDaoProposalVoteOfUser] proposal id invalid")
		return nil, errs.ErrInvalidProposalID
	}
	vote, err := e.repo.DaoVote.GetByUserAndProposalID(proposalIdNumber, userId)
	if err != nil {
		e.log.Error(err, "[Entity][GetDaoProposalVoteOfUser] repo.DaoVote.GetByUserAndProposalID failed")
		if err == gorm.ErrRecordNotFound {
			return nil, errs.ErrRecordNotFound
		}
		return nil, err
	}
	return vote, nil
}
