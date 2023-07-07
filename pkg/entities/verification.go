package entities

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/service/mochiprofile"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) NewGuildConfigWalletVerificationMessage(req model.GuildConfigWalletVerificationMessage) (*model.GuildConfigWalletVerificationMessage, error) {

	_, err := e.repo.DiscordGuilds.GetByID(req.GuildID)
	if err != nil {
		return nil, fmt.Errorf("failed to get discord guild: %v", err.Error())
	}

	_, err = e.repo.GuildConfigWalletVerificationMessage.GetOne(req.GuildID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("failed to get guild config verification: %v", err.Error())
	}

	verificationMsg := &discordgo.MessageSend{
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "Verify",
						Style:    discordgo.PrimaryButton,
						CustomID: "mochi_verify",
					},
				},
			},
		},
	}

	switch {
	case req.EmbeddedMessage != nil:
		var embeddedMsg discordgo.MessageEmbed

		err = json.Unmarshal([]byte(req.EmbeddedMessage), &embeddedMsg)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal embedded message %v: %v", req.EmbeddedMessage, err.Error())
		}

		verificationMsg.Embed = &embeddedMsg

	case req.Content != "":
		verificationMsg.Content = req.Content

	default:
		verificationMsg.Embed = &discordgo.MessageEmbed{
			Title:       ":robot: Verification required",
			Description: "Verify your wallet. This is a read-only connection. Do not share your private keys. We will never ask for your seed phrase. We will never DM you.",
			Color:       15240072,
		}
	}

	m, err := e.discord.ChannelMessageSendComplex(req.VerifyChannelID, verificationMsg)
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %v", err.Error())
	}

	req.DiscordMessageID = m.ID

	if err := e.repo.GuildConfigWalletVerificationMessage.UpsertOne(req); err != nil {
		return nil, fmt.Errorf("failed to upsert guild config verification: %v", err.Error())
	}

	return &req, nil
}

func (e *Entity) GetGuildConfigWalletVerificationMessage(guildId string) (*model.GuildConfigWalletVerificationMessage, error) {
	_, err := e.repo.DiscordGuilds.GetByID(guildId)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": guildId}).Error(err, "[e.repo.DiscordGuilds.GetByID] - failed to get guild by id")
		return nil, err
	}

	res, err := e.repo.GuildConfigWalletVerificationMessage.GetOne(guildId)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": guildId}).Error(err, "[e.repo.GuildConfigWalletVerificationMessage.GetOne] - failed to get config by guild id")
		return nil, err
	}

	return res, nil
}

func (e *Entity) UpdateGuildConfigWalletVerificationMessage(req model.GuildConfigWalletVerificationMessage) (*model.GuildConfigWalletVerificationMessage, error) {

	_, err := e.repo.DiscordGuilds.GetByID(req.GuildID)
	if err != nil {
		return nil, fmt.Errorf("failed to get discord guild: %v", err.Error())
	}

	verificationMsg, err := e.repo.GuildConfigWalletVerificationMessage.GetOne(req.GuildID)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild config verification: %v", err.Error())
	}

	var embeddedMsg discordgo.MessageEmbed

	if req.EmbeddedMessage == nil && req.Content == "" {
		embeddedMsg = discordgo.MessageEmbed{
			Title:       ":robot: Verification required",
			Description: "Verify your wallet. This is a read-only connection. Do not share your private keys. We will never ask for your seed phrase. We will never DM you.",
			Color:       15240072,
		}
	} else {
		err = json.Unmarshal([]byte(req.EmbeddedMessage), &embeddedMsg)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal embedded message %v: %v", req.EmbeddedMessage, err.Error())
		}
	}

	components := []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "Verify",
					Style:    discordgo.PrimaryButton,
					CustomID: "mochi_verify",
				},
			},
		},
	}

	switch {
	case req.VerifyChannelID != verificationMsg.VerifyChannelID:
		if verificationMsg.DiscordMessageID != "" {
			err = e.discord.ChannelMessageDelete(verificationMsg.VerifyChannelID, verificationMsg.DiscordMessageID)
			if err != nil {
				return nil, fmt.Errorf("failed to delete discord message: %v", err.Error())
			}
		}
		fallthrough

	case verificationMsg.DiscordMessageID == "":

		m, err := e.discord.ChannelMessageSendComplex(req.VerifyChannelID, &discordgo.MessageSend{
			Content:    req.Content,
			Embed:      &embeddedMsg,
			Components: components,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to send new verification message: %v", err.Error())
		}
		req.DiscordMessageID = m.ID

	default:
		_, err = e.discord.ChannelMessageEditComplex(&discordgo.MessageEdit{
			ID:         verificationMsg.DiscordMessageID,
			Channel:    verificationMsg.VerifyChannelID,
			Content:    &req.Content,
			Embed:      &embeddedMsg,
			Components: components,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to edit verification message: %v", err.Error())
		}
		req.DiscordMessageID = verificationMsg.DiscordMessageID
	}

	if err := e.repo.GuildConfigWalletVerificationMessage.UpsertOne(req); err != nil {
		return nil, fmt.Errorf("failed to upsert guild config verification: %v", err.Error())
	}

	return &req, nil
}

func (e *Entity) DeleteGuildConfigWalletVerificationMessage(guildID string) error {

	verificationMsg, err := e.repo.GuildConfigWalletVerificationMessage.GetOne(guildID)
	if err != nil {
		return fmt.Errorf("failed to get guild config verification message: %v", err.Error())
	}

	if verificationMsg.DiscordMessageID != "" {
		err = e.discord.ChannelMessageDelete(verificationMsg.VerifyChannelID, verificationMsg.DiscordMessageID)
		// case user deleted channel
		if err != nil && !strings.Contains(err.Error(), "Not Found") {
			return fmt.Errorf("failed to delete discord message: %v", err.Error())
		}
	}

	if err := e.repo.GuildConfigWalletVerificationMessage.DeleteOne(guildID); err != nil {
		return fmt.Errorf("failed to delete guild config verification: %v", err.Error())
	}

	return nil
}

func (e *Entity) AssignVerifiedRole(userDiscordID, guildID string) error {
	guildConfigVerification, err := e.repo.GuildConfigWalletVerificationMessage.GetOne(guildID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e.log.Fields(logger.Fields{
				"guild_id": guildID,
			}).Info("[entities.AssignVerifiedRole] Guild does not have config verification")
			return nil
		}
		e.log.Fields(logger.Fields{
			"guild_id": guildID,
		}).Error(err, "[entities.AssignVerifiedRole] Failed to get guild config verification")
		return err
	}

	// Skip assign if guild does not have config verified role
	if guildConfigVerification.VerifyRoleID == "" {
		e.log.Fields(logger.Fields{
			"guild_id": guildID,
		}).Info("[entities.AssignVerifiedRole] Guild does not have config verification channel")
		return nil
	}

	// Get user mochi profile
	profile, err := e.svc.MochiProfile.GetByDiscordID(userDiscordID, true)
	if err != nil {
		e.log.Fields(logger.Fields{
			"guild_id":        guildID,
			"user_discord_id": userDiscordID,
		}).Error(err, "[entities.AssignVerifiedRole] Failed to get user profile")
		return err
	}

	shouldAssignRole := false
	for _, acc := range profile.AssociatedAccounts {
		if acc.Platform == mochiprofile.PlatformEVM || acc.Platform == mochiprofile.PlatformSol {
			shouldAssignRole = true
			break
		}
	}

	if !shouldAssignRole {
		err = fmt.Errorf("user does not have any verified wallet")
		e.log.Fields(logger.Fields{
			"discord_id": userDiscordID,
			"guild_id":   guildID,
			"profile_id": profile.ID,
		}).Error(err, "[entities.AssignVerifiedRole] user does not have any verified wallet")
		return err
	}

	// assign role to user
	err = e.discord.GuildMemberRoleAdd(guildID, userDiscordID, guildConfigVerification.VerifyRoleID)
	if err != nil {
		// allow acceptable error like bot not have access to assign role
		if util.IsAcceptableErr(err) {
			e.log.Fields(logger.Fields{
				"guild_id":        guildID,
				"user_discord_id": userDiscordID,
				"verify_role_id":  guildConfigVerification.VerifyRoleID,
			}).Infof("[entities.VerifyWalletAssignRole] Acceptable errors: %v", err)
			return nil
		}
		e.log.Fields(logger.Fields{
			"guild_id":        guildID,
			"user_discord_id": userDiscordID,
			"verify_role_id":  guildConfigVerification.VerifyRoleID,
		}).Error(err, "[entities.VerifyWalletAssignRole] Failed to assign role to user")
		return err
	}
	return nil
}
