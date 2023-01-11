package entities

import (
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/common/math"
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
	guidelineChannel, err := e.createGuidelineChannel(req.GuildID, req.ChannelID)
	if err != nil {
		return nil, err
	}

	messageTempl, err := e.repo.DaoGuidelineMessages.GetByAuthority(model.Admin)
	if err != nil {
		e.log.Fields(logger.Fields{
			"authority": model.Admin,
		}).Error(err, "[entity.CreateProposalChannelConfig] e.repo.DaoGuidelineMessages.GetByAuthority failed")
		return nil, err
	}
	if err := e.sendGuidelineMessage(guidelineChannel.ID, messageTempl.Message); err != nil {
		return nil, err
	}

	guildConfigDaoProposal := model.GuildConfigDaoProposal{
		GuildId:            req.GuildID,
		GuidelineChannelId: guidelineChannel.ID,
		ProposalChannelId:  req.ChannelID,
		Authority:          model.Admin,
		RequiredAmount:     "0",
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
	requiredAmount := big.NewInt(0)
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
		requiredAmount = big.NewInt(int64(req.RequiredAmount))
	case model.CryptoToken:
		token, err := e.repo.Token.GetByAddress(req.Address, chainIdNumber)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, errors.ErrInvalidTokenContract
			}
		}
		symbol = token.Symbol
		// convert decimal token
		requiredAmount = big.NewInt(1).Mul(big.NewInt(int64(req.RequiredAmount)), math.BigPow(10, int64(token.Decimals)))
	}

	guidelineChannel, err := e.createGuidelineChannel(req.GuildID, req.ChannelID)
	if err != nil {
		return nil, err
	}

	messageTempl, err := e.repo.DaoGuidelineMessages.GetByAuthority(model.TokenHolder)
	if err != nil {
		e.log.Fields(logger.Fields{
			"authority": model.TokenHolder,
		}).Error(err, "[entity.CreateProposalChannelConfig] e.repo.DaoGuidelineMessages.GetByAuthority failed")
		return nil, err
	}
	msgDescription := fmt.Sprintf(messageTempl.Message, req.RequiredAmount, symbol)
	if err := e.sendGuidelineMessage(guidelineChannel.ID, msgDescription); err != nil {
		return nil, err
	}

	guildConfigDaoProposal := model.GuildConfigDaoProposal{
		GuildId:            req.GuildID,
		GuidelineChannelId: guidelineChannel.ID,
		ProposalChannelId:  req.ChannelID,
		Authority:          model.TokenHolder,
		Type:               &req.Type,
		RequiredAmount:     requiredAmount.Text(10),
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

func (e *Entity) createGuidelineChannel(guildID, proposalChannelID string) (guidelineChannel *discordgo.Channel, err error) {
	proposalChannel, err := e.svc.Discord.Channel(proposalChannelID)
	if err != nil {
		e.log.Fields(logger.Fields{
			"proposalChannelID": proposalChannelID,
		}).Error(err, "[entity.CreateProposalChannelConfig] svc.Discord.Channel failed")
		return nil, errors.ErrInvalidDiscordChannelID
	}

	guidelineChannelCreateData := discordgo.GuildChannelCreateData{
		Name:                 fmt.Sprintf("guideline - %s", proposalChannel.Name),
		Type:                 proposalChannel.Type,
		PermissionOverwrites: proposalChannel.PermissionOverwrites,
		ParentID:             proposalChannel.ParentID,
	}
	guidelineChannel, err = e.svc.Discord.CreateChannel(guildID, guidelineChannelCreateData)
	if err != nil {
		e.log.Fields(logger.Fields{
			"guildID":    guildID,
			"createData": guidelineChannelCreateData,
		}).Error(err, "[entity.CreateProposalChannelConfig] svc.Discord.CreateChannel failed")
		return nil, err
	}

	return guidelineChannel, nil
}

func (e *Entity) sendGuidelineMessage(guidelineChannelID, description string) error {
	msgSend := discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       "<:mail:1058304339237666866> Create a proposal",
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
							Name: "mailsend",
							ID:   "1058304343293567056",
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
