package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type handlerLog struct {
	Log            *logrus.Logger
	discordWebhook string
	fields         logrus.Fields
}

type logFields map[string]interface{}

func (h *handlerLog) WithFields(f logFields) *handlerLog {
	fields := logrus.Fields(f)
	for k, v := range h.fields {
		fields[k] = v
	}

	return &handlerLog{
		Log:            h.Log,
		discordWebhook: h.discordWebhook,
		fields:         fields,
	}
}

func (h *handlerLog) WithGinContext(c *gin.Context) *handlerLog {
	cc := c.Copy()

	url := cc.Request.URL.String()
	method := cc.Request.Method

	fields := logFields{
		"method": method,
		"url":    url,
	}

	for k, v := range h.fields {
		fields[k] = v
	}

	return &handlerLog{
		Log:            h.Log,
		discordWebhook: h.discordWebhook,
		fields:         logrus.Fields(fields),
	}
}

func newHandlerLog(l logger.Log, discordWebhook string) handlerLog {
	return handlerLog{
		Log:            l.GetLogger(),
		discordWebhook: discordWebhook,
		fields:         logrus.Fields{},
	}
}

func (l handlerLog) Errorf(format string, v ...interface{}) {
	l.Log.WithFields(l.fields).Errorf(format, v...)
	l.sendDiscordMessage(fmt.Sprintf(format, v...), l.fields)
}

func (l handlerLog) Error(v ...interface{}) {
	l.Log.WithFields(l.fields).Error(v...)
	l.sendDiscordMessage(fmt.Sprint(v...), l.fields)
}

func (l handlerLog) sendDiscordMessage(description string, fields logrus.Fields) {
	hostname, err := os.Hostname()
	if err != nil {
		l.Log.Errorf("Error getting hostname: %s", err)
		return
	}

	msg := fmt.Sprintf("```%s %s\n--------\n%s```", fields["method"], fields["url"], description)

	// Add fields
	discordMsg := discordgo.Message{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       hostname,
				Description: msg,
			},
		},
	}

	discordMsgPayload, err := json.Marshal(discordMsg)
	if err != nil {
		l.Log.Errorf("Error marshalling discord message: %s", err)
		return
	}

	res, err := http.Post(l.discordWebhook, "application/json", bytes.NewReader(discordMsgPayload))
	if err != nil {
		l.Log.Errorf("Error sending discord message: %s", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			l.Log.Errorf("Error reading response body: %s", err)
			return
		}

		l.Log.Errorf("Error sending discord message with status %v: %v", res.StatusCode, string(body))
		return
	}
}

func (l handlerLog) SendMessageToDiscord(description, token, amount string) {
	logFields := map[string]string{"token": token, "amount": amount}
	// Add fields
	msgFields := []*discordgo.MessageEmbedField{}
	for k, v := range logFields {
		msgFields = append(msgFields, &discordgo.MessageEmbedField{
			Name:  k,
			Value: v,
		})
	}
	discordMsg := discordgo.Message{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       "Test",
				Description: description,
				Fields:      msgFields,
			},
		},
	}

	discordMsgPayload, err := json.Marshal(discordMsg)
	if err != nil {
		l.Log.Errorf("Error marshalling discord message: %s", err)
		return
	}

	res, err := http.Post(l.discordWebhook, "application/json", bytes.NewReader(discordMsgPayload))
	if err != nil {
		l.Log.Errorf("Error sending discord message: %s", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			l.Log.Errorf("Error reading response body: %s", err)
			return
		}

		l.Log.Errorf("Error sending discord message with status %v: %v", res.StatusCode, string(body))
		return
	}
}
