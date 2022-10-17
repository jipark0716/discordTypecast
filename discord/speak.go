package discord

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/jipark0716/discordTypecast/ent"
	"github.com/jipark0716/discordTypecast/handler"
)

func (d *Discord) OnSpeakTts(event *discordgo.InteractionCreate) {
	setting, err := d.UserRepository.FindByUserId(event.Member.User.ID)
	if err != nil {
		handler.Report(err)
		return
	}

	if setting == nil {
		d.InteractionRespond(
			event.Interaction,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "설정된 목소리가 없습니다.",
				},
			},
		)
		return
	}

	err = d.InteractionRespond(
		event.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "말하는중",
			},
		},
	)
	if err != nil {
		handler.Report(err)
		return
	}

	defer func() {
		err := d.InteractionResponseDelete(event.Interaction)
		if err != nil {
			handler.Report(err)
		}
	}()

	if err != nil {
		handler.Report((err))
	}

	channel, err := d.Channel(event.ChannelID)
	if err != nil {
		handler.Report(err)
		return
	}

	if channel.Type != discordgo.ChannelTypeGuildVoice {
		return
	}

	var content string
	for _, option := range event.ApplicationCommandData().Options {
		if option.Name == "text" {
			content = option.Value.(string)
		}
	}

	if content == "" {
		handler.Report(err)
		return
	}

	tyepcastMessage, err := d.TypecastMessageRepository.Create(setting, event.GuildID, channel.ID, content)

	if err != nil {
		handler.Report(err)
		return
	}

	d.Speak(tyepcastMessage)
}

func (d *Discord) Speak(tyepcastMessage *ent.TypecastMessage) {
	audio, err := d.Typecast.Do(tyepcastMessage)
	if err != nil {
		handler.Report(err)
		return
	}

	audioFilePath := fmt.Sprintf(
		"./queue/%s/%d",
		tyepcastMessage.GuildID,
		tyepcastMessage.ID,
	)

	err = ioutil.WriteFile(
		audioFilePath,
		audio,
		755,
	)
	if err != nil {
		handler.Report(err)
		return
	}

	voiceConnection, err := d.ChannelVoiceJoin(tyepcastMessage.GuildID, tyepcastMessage.ChannelID, false, true)
	if err != nil {
		handler.Report(err)
	}

	voiceConnection.PlayAudioFile(audioFilePath, make(chan bool))
	os.Remove(audioFilePath)
	tyepcastMessage.
		Update().
		SetSendAt(time.Now()).
		SetStatus(1).
		Save(context.Background())

	go d.RefreshIdleConnect(voiceConnection)
}
