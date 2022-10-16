package discord

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"github.com/jipark0716/discordTypecast/ent"
	"github.com/jipark0716/discordTypecast/handler"
)

func (d *Discord) OnSpeakTts(event *discordgo.InteractionCreate) {
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

	setting, err := d.UserRepository.FindByUserId(event.Member.User.ID)
	if err != nil {
		handler.Report(err)
		return
	}

	d.Speak(
		channel,
		content,
		event.ID,
		setting,
	)
}

func (d *Discord) Speak(channel *discordgo.Channel, text string, messageID string, setting *ent.UserTypecastSetting) {
	audio, err := d.Typecast.Do(setting, text)
	if err != nil {
		handler.Report(err)
		return
	}

	audioFilePath := fmt.Sprintf(
		"./queue/%s/%s",
		channel.GuildID,
		messageID,
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

	d.DispathVoice(channel, audioFilePath)
}

func (d *Discord) DispathVoice(channel *discordgo.Channel, audioFilePath string) {
	voiceConnection, err := d.ChannelVoiceJoin(channel.GuildID, channel.ID, false, true)
	if err != nil {
		handler.Report(err)
	}

	dgvoice.PlayAudioFile(voiceConnection, audioFilePath, make(chan bool))
	os.Remove(audioFilePath)
}
