package job

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/service"
)

type checkInvalidateEmoji struct {
	entity  *entities.Entity
	log     logger.Logger
	service *service.Service
	config  config.Config
}

// NewCheckInvalidateEmoji returns a new job that checks for invalid emojis
func NewCheckInvalidateEmoji(e *entities.Entity, l logger.Logger, s *service.Service, cfg config.Config) Job {
	return &checkInvalidateEmoji{
		entity:  e,
		log:     l,
		service: s,
		config:  cfg,
	}
}

func (j *checkInvalidateEmoji) Run() error {
	emojis, err := j.service.Discord.GetGuildEmojis()
	if err != nil {
		j.log.Error(err, "failed to get guild emojis")
		return err
	}

	dbEmojis, _, err := j.entity.GetListEmojis(request.GetListEmojiRequest{Size: 10000, Page: 0, IsQueryAll: true})
	if err != nil {
		j.log.Error(err, "failed to get db emojis")
		return err
	}

	// Find missing emojis
	invalidateEmojis := make([]*model.EmojiData, 0)
	for _, emoji := range dbEmojis {
		found := false

		for _, e := range emojis {
			if strings.Contains(emoji.Emoji, e.ID) {
				found = true

				if !e.Available {
					invalidateEmojis = append(invalidateEmojis, emoji)
				}
				break
			}
		}

		if !found {
			invalidateEmojis = append(invalidateEmojis, emoji)
		}
	}

	// Send message to product tracking channel
	if len(invalidateEmojis) > 0 {
		content := ""
		for i := range invalidateEmojis {
			e := invalidateEmojis[i]

			content += fmt.Sprintf("・%s\n", e.Code)
		}

		msg := discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Invalid Emojis",
				Description: content,
				Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
			},
		}

		err = j.service.Discord.SendMessage(
			j.config.MochiProductTrackingChannelID,
			msg,
		)

		if err != nil {
			j.log.Error(err, "failed to send message to product tracking channel")
			return err
		}
	}

	return nil
}
