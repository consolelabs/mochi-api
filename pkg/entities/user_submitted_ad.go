package entities

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) GetAllAd() ([]model.UserSubmittedAd, int64, error) {
	return e.repo.UserSubmittedAd.GetAll()
}

func (e *Entity) GetAdById(id string) (*model.UserSubmittedAd, error) {
	if id == "random" {
		return e.GetOneRandomAd()
	}
	i, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id")
	}
	data, err := e.repo.UserSubmittedAd.GetById(i)
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return data, err
}

func (e *Entity) GetOneRandomAd() (*model.UserSubmittedAd, error) {
	ads, count, err := e.repo.UserSubmittedAd.GetAll()
	if err != nil {
		e.log.Errorf(err, "[entities.GetOneRandomAd] repo.UserSubmittedAd.GetAll failed")
		return nil, err
	}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(int(count))
	return &ads[randomIndex], nil
}

func (e *Entity) CreateAd(req request.InsertUserAd) error {
	// create channel in Mochi server
	channel, err := e.svc.Discord.CreateChannel(e.cfg.MochiGuildID, discordgo.GuildChannelCreateData{
		Name:     req.Name,
		ParentID: e.cfg.MochiAdDiscussionCategoryID,
	})
	if err != nil {
		e.log.Errorf(err, "[entities.CreateAd] e.svc.Discord.CreateChannel failed")
		return err
	}
	// add to db
	ad, err := e.repo.UserSubmittedAd.CreateOne(model.UserSubmittedAd{
		Status:       "pending",
		Introduction: req.Introduction,
		CreatorId:    req.CreatorId,
		AdChannelId:  channel.ID,
		Name:         req.Name,
		Description:  req.Description,
		Image:        req.Image,
		IsPodtownAd:  req.IsPodtownAd,
	})
	if err != nil {
		e.log.Errorf(err, "[entities.CreateAd] e.repo.UserSubmittedAd.CreateOne failed")
		// delete channel if any step fails
		delErr := e.svc.Discord.DeleteChannel(channel.ID)
		if delErr != nil {
			e.log.Errorf(err, "[entities.CreateAd] e.svc.Discord.DeleteChannel failed")
		}
		return err
	}
	// submit ad to Mochi channel
	if req.Image == "" {
		req.Image = "none"
	}
	err = e.svc.Discord.SendMessage(e.cfg.MochiAdDiscussionChannelID, discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       "Ads submission received",
				Description: fmt.Sprintf("<@%s> has submitted an ad with the following info\n\nBrief introduction\n```%s```\nName\n```%s```\nDescription\n```%s```\nImage\n```%s```", req.CreatorId, req.Introduction, req.Name, req.Description, req.Image),
			},
		},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{discordgo.Button{
					Label:    "Approve",
					CustomID: fmt.Sprintf("ad-approve-%d", ad.ID),
					Style:    discordgo.SuccessButton,
					Emoji: discordgo.ComponentEmoji{
						Name: "approve",
						ID:   "1013775501757780098",
					},
				},
					discordgo.Button{
						Label:    "Reject",
						CustomID: fmt.Sprintf("ad-reject-%d", ad.ID),
						Style:    discordgo.DangerButton,
					},
					discordgo.Button{
						Label: "Jump to channel",
						Style: discordgo.LinkButton,
						URL:   fmt.Sprintf("https://discord.com/channels/%s/%s", e.cfg.MochiGuildID, channel.ID),
					}},
			},
		},
	})
	if err != nil {
		e.log.Errorf(err, "[entities.CreateAd] e.svc.Discord.SendMessage failed")
		// delete channel if any step fails
		delErr := e.svc.Discord.DeleteChannel(channel.ID)
		if delErr != nil {
			e.log.Errorf(err, "[entities.CreateAd] e.svc.Discord.DeleteChannel failed")
		}
		return err
	}
	return nil
}

func (e *Entity) DeleteAdById(req request.DeleteUserAd) error {
	return e.repo.UserSubmittedAd.DeleteOne(req.ID)
}

func (e *Entity) UpdateAdById(req request.UpdateUserAd) error {
	return e.repo.UserSubmittedAd.UpdateStatus(req.ID, req.Status)
}

func (e *Entity) InitAdSubmission(req request.InitAdSubmission) error {
	err := e.svc.Discord.SendMessage(req.ChannelId, discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       "Create ads with Mochi",
				Description: "Welcome to Mochi Ads, where we enable interested partners to advertise via our Discord bot. Please be noted that we uphold high standards for honesty and ethics, and only accept ads that align with these values.\n\nSimply click on <:mailsend:1058304343293567056> to get started.",
			},
		},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "Create Ads",
						CustomID: fmt.Sprintf("ad-create-%s", req.GuildId),
						Style:    discordgo.PrimaryButton,
						Emoji: discordgo.ComponentEmoji{
							Name: "mailsend",
							ID:   "1058304343293567056",
						},
					},
				},
			},
		},
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[InitAdSubmission] - e.svc.Discord.SendMessage failed")
		return fmt.Errorf("failed to init ad submission")
	}
	return nil
}
