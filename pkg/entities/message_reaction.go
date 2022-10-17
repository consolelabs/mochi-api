package entities

import (
	"encoding/json"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) AddMessageReaction(req request.MessageReactionRequest) error {
	cfg, err := e.repo.GuildConfigReactionRole.GetByMessageID(req.GuildID, req.MessageID)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": req.GuildID, "messageID": req.MessageID}).
			Info("[e.AddMessageReaction] this message is not reaction role config for guild")
		return nil
	}

	roles := []response.Role{}
	if err := json.Unmarshal([]byte(cfg.ReactionRoles), &roles); err != nil {
		e.log.Fields(logger.Fields{"reactionRoles": cfg.ReactionRoles}).
			Error(err, "[e.AddMessageReaction] failed to unmarshal reaction roles")
		return err
	}

	for _, role := range roles {
		if role.Reaction == req.Reaction {
			if err := e.repo.MessageReaction.Create(model.MessageReaction{
				MessageID: req.MessageID,
				GuildID:   req.GuildID,
				UserID:    req.UserID,
				Reaction:  req.Reaction,
			}); err != nil {
				e.log.Fields(logger.Fields{
					"messageID": req.MessageID,
					"guildID":   req.GuildID,
					"userID":    req.UserID,
					"reaction":  req.Reaction,
				}).Error(err, "[e.AddMessageReaction] failed to create message reaction")
				return err
			}

			if err := e.AddGuildMemberRole(req.GuildID, req.UserID, role.ID); err != nil {
				e.log.Fields(logger.Fields{
					"guildID": req.GuildID,
					"userID":  req.UserID,
					"roleID":  role.ID,
				}).Error(err, "[e.AddMessageReaction] failed to add guild member role")
				return err
			}
		}
	}

	return nil
}

func (e *Entity) RemoveMessageReaction(req request.MessageReactionRequest) error {
	cfg, err := e.repo.GuildConfigReactionRole.GetByMessageID(req.GuildID, req.MessageID)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": req.GuildID, "messageID": req.MessageID}).
			Info("[e.RemoveMessageReaction] this message is not reaction role config for guild")
		return nil
	}

	roles := []response.Role{}
	if err := json.Unmarshal([]byte(cfg.ReactionRoles), &roles); err != nil {
		e.log.Fields(logger.Fields{"reactionRoles": cfg.ReactionRoles}).
			Error(err, "[e.RemoveMessageReaction] failed to unmarshal reaction roles")
		return err
	}

	// remove reaction roles
	for _, role := range roles {
		if role.Reaction == req.Reaction {
			if err := e.repo.MessageReaction.Delete(req.MessageID, req.UserID, req.Reaction); err != nil {
				e.log.Fields(logger.Fields{
					"messageID": req.MessageID,
					"userID":    req.UserID,
					"reaction":  req.Reaction,
				}).Error(err, "[e.RemoveMessageReaction] failed to delete message reaction")
				return err
			}

			if err := e.RemoveGuildMemberRole(req.GuildID, req.UserID, role.ID); err != nil {
				e.log.Fields(logger.Fields{
					"guildID": req.GuildID,
					"userID":  req.UserID,
					"roleID":  role.ID,
				}).Error(err, "[e.RemoveMessageReaction] failed to remove guild member role")
				return err
			}
		}
	}

	return nil
}
