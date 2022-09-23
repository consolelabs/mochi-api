package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func isMissingPermissionsErr(msg string) bool {
	return strings.Contains(msg, "missing permissions") && strings.Contains(msg, "50013")
}

func isMissingAccessErr(msg string) bool {
	return strings.Contains(msg, "missing access") && strings.Contains(msg, "50001")
}

func isMemberNotFoundErr(msg string) bool {
	return strings.Contains(msg, "404 not found") && strings.Contains(msg, "10007")
}

func IsAcceptableErr(err error) bool {
	if err == nil {
		return false
	}
	msg := strings.ToLower(err.Error())
	return isMissingPermissionsErr(msg) || isMissingAccessErr(msg) || isMemberNotFoundErr(msg)
}

type upvoteMsg struct {
	Title       string
	Description string
	Image       string
}

func GenerateUpvoteMessage(discordID, source string) *upvoteMsg {
	sourceName, sourceUrl := UpvoteSourceNameAndUrl(source)
	presets := []upvoteMsg{
		{
			Title:       "Mochi appreciates it!",
			Description: fmt.Sprintf("<@%s> just voted for Mochi on [%s](%s).", discordID, sourceName, sourceUrl),
			Image:       "https://cdn.discordapp.com/attachments/1003381172178530494/1019165378213068840/unknown.png",
		},
		{
			Title:       "Wait, what?",
			Description: fmt.Sprintf("<@%s> voted for Mochi, is that all you needed to do to receive rewards?", discordID),
			Image:       "https://cdn.discordapp.com/attachments/1003381172178530494/1019165378447937556/unknown.png",
		},
		{
			Title:       "Promoted!",
			Description: fmt.Sprintf("Mochi got a vote and <@%s> can now use the `$wl` command to its fullest, win-win.", discordID),
			Image:       "https://cdn.discordapp.com/attachments/1003381172178530494/1019165378720583750/unknown.png",
		},
		{
			Title:       "Thank you!",
			Description: fmt.Sprintf("Thank you <@%s> for voting Mochi, Mochi truly is one of the greatest bots.", discordID),
			Image:       "https://cdn.discordapp.com/attachments/986854719999864863/1019183908681695282/obamamochi.jpg",
		},
		{
			Title:       "Imagine not voting for Mochi",
			Description: fmt.Sprintf("Fortunately <@%s> has redeemed themselves by voting on [%s](%s).", discordID, sourceName, sourceUrl),
			Image:       "https://cdn.discordapp.com/attachments/986854719999864863/1019184889725206528/unknown.png",
		},
		{
			Title:       "Trade offer alert!",
			Description: fmt.Sprintf("Happy to announce that <@%s> has closed a great deal on [%s](%s).", discordID, sourceName, sourceUrl),
			Image:       "https://cdn.discordapp.com/attachments/986854719999864863/1019188156584706048/trademochi.jpg",
		},
		{
			Title:       "Mochi is grateful",
			Description: fmt.Sprintf("Thank you <@%s> for the upvote, can Mochi have another one uwu?", discordID),
			Image:       "https://cdn.discordapp.com/attachments/986854719999864863/1019189320600530974/unknown.png",
		},
		{
			Title:       "You sure that is enough?",
			Description: fmt.Sprintf("Absolutely, an upvote is all <@%s> needs to enjoy new perks.", discordID),
			Image:       "https://cdn.discordapp.com/attachments/986854719999864863/1019190354018308146/onepls.jpg",
		},
	}
	ran := rand.Intn(len(presets) - 1)
	return &presets[ran]
}

// RetryRequest retry handler until it succeeds or is acceptable or reaches the limit of times
func RetryRequest(handler func() error, times int, interval time.Duration) error {
	err := handler()
	for i := 0; err != nil && !IsAcceptableErr(err) && i < times-1; i++ {
		time.Sleep(interval)
		err = handler()
	}
	return err
}
