package entities

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) Test() {
	// e.repo.GuildConfigDaoProposal.GetById(1)

	// e.repo.DaoProposal.GetById(1)

	// e.repo.DaoVoteOption.GetById(1)

	// e.repo.DaoVote.GetById(1)

	// e.repo.DaoProposalVoteOption.GetById(1)
}

func (e *Entity) CreateProposalChannelConfig(req request.CreateProposalChannelConfig) (*model.GuildConfigDaoProposal, error) {
	switch req.Authority {
	case model.Admin:
		return e.createConfigDaoProposalWithAdminAuthority(req)
	case model.TokenHolder:
		return e.createConfigDaoProposalWithTokenHolderAuthority(req)
	}
	return nil, errors.ErrInternalError
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

func (e *Entity) createConfigDaoProposalWithAdminAuthority(req request.CreateProposalChannelConfig) (*model.GuildConfigDaoProposal, error) {
	guidelineChannel, proposalChannel, err := e.createDAOVotingChannels(req.GuildID, req.ChannelID)
	if err != nil {
		return nil, err
	}

	msgDescription := `Only Administrators have the authority to post proposals.
	The proposal should include this information:
	- Proposal Title
	- Proposal Description (maximum 2000 words)
	- Vote option
	- Vote duration

	To create a proposal, click **<:transaction:933341692667506718>.**,`
	if err := e.sendGuidelineMessage(guidelineChannel.ID, msgDescription); err != nil {
		return nil, err
	}

	guildConfigDaoProposal := model.GuildConfigDaoProposal{
		GuildId:            req.GuildID,
		GuidelineChannelId: guidelineChannel.ID,
		ProposalChannelId:  proposalChannel.ID,
		Authority:          model.Admin,
	}
	config, err := e.repo.GuildConfigDaoProposal.Create(guildConfigDaoProposal)
	if err != nil {
		e.log.Fields(logger.Fields{
			"config": guildConfigDaoProposal,
		}).Error(err, "[entity.CreateProposalChannelConfig] e.repo.GuildConfigDaoProposal.Create failed")
		return nil, err
	}
	return config, err
}

func (e *Entity) createConfigDaoProposalWithTokenHolderAuthority(req request.CreateProposalChannelConfig) (*model.GuildConfigDaoProposal, error) {
	chainID := util.ConvertInputToChainId(req.Chain)
	if chainID == "" {
		e.log.Fields(logger.Fields{
			"chain": req.Chain,
		}).Error(nil, "[entity.CreateProposalChannelConfig] util.ConvertInputToChainId failed")
		return nil, errors.ErrInvalidChain
	}
	chainIdNumber, err := strconv.Atoi(chainID)
	if err != nil {
		e.log.Fields(logger.Fields{
			"chainID": chainID,
		}).Error(err, "[entity.GetGuildConfigDaoProposalByGuildID] strconv.Atoi failed")
		return nil, err
	}
	symbol := ""
	var requiredAmount int64 = 0
	if req.Type == "" {
		e.log.Fields(logger.Fields{
			"type": req.Type,
		}).Error(nil, "[entity.CreateProposalChannelConfig] proposal type is empty")
		return nil, errors.ErrInvalidProposalType
	}
	switch req.Type {
	case model.NFT:
		collection, err := e.repo.NFTCollection.GetByAddressChainId(req.Address, chainID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, errors.ErrInvalidTokenContract
			}
			e.log.Fields(logger.Fields{
				"address": req.Address,
				"chainID": chainID,
			}).Error(err, "[entity.GetGuildConfigDaoProposalByGuildID] e.repo.GuildConfigDaoProposal.GetByGuildId failed")
			return nil, err
		}
		symbol = collection.Symbol
		requiredAmount = int64(req.RequiredAmount)
	case model.CryptoToken:
		token, err := e.repo.Token.GetByAddress(req.Address, chainIdNumber)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, errors.ErrInvalidTokenContract
			}
		}
		symbol = token.Symbol
		// convert decimal token
		requiredAmount = int64(req.RequiredAmount * float64(10^token.Decimals))
	}

	guidelineChannel, proposalChannel, err := e.createDAOVotingChannels(req.GuildID, req.ChannelID)
	if err != nil {
		return nil, err
	}

	msgDescription := fmt.Sprintf(`You have to connect your wallet and own %v %s to post a proposal.
	The proposal should include this information:
	- Proposal Title
	- Proposal Description (maximum 2000 words)
	- Vote option
	- Vote duration

	To create a proposal, click **<:transaction:933341692667506718>.**,`, req.RequiredAmount, symbol)
	if err := e.sendGuidelineMessage(guidelineChannel.ID, msgDescription); err != nil {
		return nil, err
	}

	guildConfigDaoProposal := model.GuildConfigDaoProposal{
		GuildId:            req.GuildID,
		GuidelineChannelId: guidelineChannel.ID,
		ProposalChannelId:  proposalChannel.ID,
		Authority:          model.TokenHolder,
		Type:               &req.Type,
		RequiredAmount:     requiredAmount,
		ChainID:            int64(chainIdNumber),
		Symbol:             symbol,
		Address:            req.Address,
	}
	config, err := e.repo.GuildConfigDaoProposal.Create(guildConfigDaoProposal)
	if err != nil {
		e.log.Fields(logger.Fields{
			"config": guildConfigDaoProposal,
		}).Error(err, "[entity.CreateProposalChannelConfig] e.repo.GuildConfigDaoProposal.Create failed")
		return nil, err
	}

	return config, nil
}

func (e *Entity) createDAOVotingChannels(guildID, sampleChannelID string) (guidelineChannel, proposalChannel *discordgo.Channel, err error) {
	sampleChannel, err := e.svc.Discord.Channel(sampleChannelID)
	if err != nil {
		e.log.Fields(logger.Fields{
			"channelID": sampleChannelID,
		}).Error(err, "[entity.CreateProposalChannelConfig] svc.Discord.Channel failed")
		return nil, nil, errors.ErrInvalidDiscordChannelID
	}

	guidelineChannelCreateData := discordgo.GuildChannelCreateData{
		Name:                 fmt.Sprintf("guideline - %s", sampleChannel.Name),
		Type:                 sampleChannel.Type,
		PermissionOverwrites: sampleChannel.PermissionOverwrites,
		ParentID:             sampleChannel.ParentID,
	}
	guidelineChannel, err = e.svc.Discord.CreateChannel(guildID, guidelineChannelCreateData)
	if err != nil {
		e.log.Fields(logger.Fields{
			"guildID":    guildID,
			"createData": guidelineChannelCreateData,
		}).Error(err, "[entity.CreateProposalChannelConfig] svc.Discord.CreateChannel failed")
		return nil, nil, err
	}

	proposalChannelCreateData := discordgo.GuildChannelCreateData{
		Name:                 "Proposals",
		Type:                 sampleChannel.Type,
		PermissionOverwrites: sampleChannel.PermissionOverwrites,
		ParentID:             sampleChannel.ParentID,
	}
	proposalChannel, err = e.svc.Discord.CreateChannel(guildID, proposalChannelCreateData)
	if err != nil {
		e.log.Fields(logger.Fields{
			"guildID":    guildID,
			"createData": proposalChannelCreateData,
		}).Error(err, "[entity.CreateProposalChannelConfig] svc.Discord.CreateChannel failed")
		return nil, nil, err
	}

	return guidelineChannel, proposalChannel, nil
}

func (e *Entity) sendGuidelineMessage(guidelineChannelID, description string) error {
	msgSend := discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       "<:transaction:933341692667506718> Create a proposal",
				Description: description,
				Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
				Type:        discordgo.EmbedTypeArticle,
			},
		},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "Create a proposal",
						Style:    discordgo.PrimaryButton,
						Disabled: false,
						CustomID: "create-proposal",
						Emoji: discordgo.ComponentEmoji{
							Name: "transaction",
							ID:   "933341692667506718",
						},
					},
				},
			},
		},
	}
	if err := e.svc.Discord.SendMessage(guidelineChannelID, msgSend); err != nil {
		e.log.Fields(logger.Fields{
			"guidelineChannelID": guidelineChannelID,
		}).Error(err, "[entity.CreateProposalChannelConfig] e.svc.Discord.SendMessage failed")
		return err
	}
	return nil
}
