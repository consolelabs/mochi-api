package entities

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) NewGuildScheduledEvents(guildID string) error {
	events, err := e.discord.GuildScheduledEvents(guildID, false)
	if err != nil {
		return fmt.Errorf("failed to get scheduled events: %w", err)
	}

	for _, event := range events {
		err = e.repo.GuildScheduledEvent.UpsertOne(&model.GuildScheduledEvent{
			GuildID: event.GuildID,
			EventID: event.ID,
			Status:  event.Status,
		})
		if err != nil {
			e.log.Errorf(err, "failed to upsert guild scheduled event %s", event.ID)
			continue
		}
	}

	return nil
}

func (e *Entity) UpdateGuildScheduledEventStatus(guildID string) error {

	events, err := e.repo.GuildScheduledEvent.ListUncompleteByGuildID(guildID)
	if err != nil {
		return fmt.Errorf("failed to get scheduled events: %w", err)
	}

	for _, event := range events {
		dEvent, err := e.discord.GuildScheduledEvent(event.GuildID, event.EventID, false)
		if err != nil {
			e.log.Errorf(err, "failed to get scheduled event %s", event.EventID)
			continue
		}

		if dEvent.Status == event.Status {
			continue
		}

		err = e.repo.GuildScheduledEvent.UpsertOne(&model.GuildScheduledEvent{
			GuildID: dEvent.GuildID,
			EventID: dEvent.ID,
			Status:  dEvent.Status,
		})
		if err != nil {
			e.log.Errorf(err, "failed to upsert guild scheduled event %s", event.EventID)
			continue
		}

		if dEvent.Status == discordgo.GuildScheduledEventStatusCompleted {
			err = e.HandleCompletedEvent(dEvent)
			if err != nil {
				e.log.Errorf(err, "failed to handle completed event %s", event.EventID)
				continue
			}
		}
	}

	return nil
}

func (e *Entity) HandleCompletedEvent(event *discordgo.GuildScheduledEvent) error {

	_, err := e.HandleUserActivities(&request.HandleUserActivityRequest{
		GuildID:   event.GuildID,
		UserID:    event.CreatorID,
		Action:    "event_host",
		Timestamp: event.ScheduledStartTime,
	})
	if err != nil {
		return fmt.Errorf("failed to handle event host: %v", err.Error())
	}

	var afterID string

	for {
		users, err := e.discord.GuildScheduledEventUsers(event.GuildID, event.ID, 100, false, "", afterID)
		if err != nil {
			return fmt.Errorf("failed to get scheduled event users after ID %s: %v", afterID, err.Error())
		}

		for _, user := range users {
			if user.User.ID == event.CreatorID {
				continue
			}
			_, err := e.HandleUserActivities(&request.HandleUserActivityRequest{
				GuildID:   event.GuildID,
				UserID:    user.User.ID,
				Action:    "event_participant",
				Timestamp: event.ScheduledStartTime,
			})
			if err != nil {
				e.log.Errorf(err, "failed to handle event participant: %v guild %s event %s", user.User.ID, event.GuildID, event.ID)
				continue
			}
		}

		if len(users) < 100 {
			break
		}

		afterID = users[len(users)-1].User.ID
	}

	return nil
}
