package entities

import (
	"fmt"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) CreateGuildUserIfNotExists(guildID, userID, nickname string) error {
	guildUser := &model.GuildUser{
		GuildID:   guildID,
		UserID:    userID,
		Nickname:  nickname,
		InvitedBy: "",
	}

	if err := e.repo.GuildUsers.FirstOrCreate(guildUser); err != nil {
		return fmt.Errorf("failed to create if not exists guild user: %w", err)
	}

	return nil
}

func (e *Entity) SendUserXP(req request.SendUserXPRequest) error {
	amountXP := req.Amount / len(req.Recipients)
	if req.Each {
		amountXP = req.Amount
	}
	records := []model.GuildUserActivityLog{}
	for _, recipient := range req.Recipients {
		records = append(records, model.GuildUserActivityLog{
			GuildID:      req.GuildID,
			UserID:       recipient,
			ActivityName: "sendXP",
			EarnedXP:     amountXP,
		})
	}
	return e.repo.GuildUserActivityLog.CreateBatch(records)
}
