package discord

import (
	"os"
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestDiscord_NotifyNewGuild(t *testing.T) {
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		t.Fatal(err)
	}

	mochiLogChannelID := os.Getenv("MOCHI_LOG_CHANNEL_ID")

	hnhGuild := "895659000996200508"

	type fields struct {
		session           *discordgo.Session
		mochiLogChannelID string
	}
	type args struct {
		guildID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			fields: fields{
				session:           discord,
				mochiLogChannelID: mochiLogChannelID,
			},
			args: args{
				guildID: hnhGuild,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Discord{
				session:           tt.fields.session,
				mochiLogChannelID: tt.fields.mochiLogChannelID,
			}
			if err := d.NotifyNewGuild(tt.args.guildID); (err != nil) != tt.wantErr {
				t.Errorf("Discord.NotifyNewGuild() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
