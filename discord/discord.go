package discord

import (
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"github.com/jipark0716/discordTypecast/config"
	"github.com/jipark0716/discordTypecast/repositories"
)

type Discord struct {
	*discordgo.Session
	Typecast                  repositories.Typecast
	UserRepository            repositories.UserRepository
	TypecastMessageRepository repositories.TypecastMessageRepository
	VoiceConnections          map[string]*VoiceConnection
	ApplicationId             string
}

type VoiceConnection struct {
	*discordgo.VoiceConnection
	idleConnect chan (bool)
	Mutex       *sync.Mutex
}

func NewDiscord(
	typecast repositories.Typecast,
	userRepository repositories.UserRepository,
	messageRepository repositories.TypecastMessageRepository,
) (discord Discord, err error) {

	discord = Discord{}
	discord.UserRepository = userRepository
	discord.VoiceConnections = make(map[string]*VoiceConnection)
	discord.Typecast = typecast
	discord.TypecastMessageRepository = messageRepository
	discord.ApplicationId = config.Get("DISCORD_ID").(string)
	discord.Session, err = discordgo.New(config.Get("DISCORD_TOKEN").(string))
	return
}

func (d *Discord) ApplicationCommands() ([]*discordgo.ApplicationCommand, error) {
	return d.Session.ApplicationCommands(
		d.ApplicationId,
		"",
	)
}

func (d *Discord) ApplicationCommandCreate(cmd *discordgo.ApplicationCommand) (*discordgo.ApplicationCommand, error) {
	cmd.ApplicationID = d.ApplicationId
	return d.Session.ApplicationCommandCreate(
		d.ApplicationId,
		"",
		cmd,
	)
}

func (d *Discord) CreateCommand() (err error) {
	_, err = d.CreateChangeVoiceCommand()
	if err != nil {
		return
	}

	_, err = d.CreateSpeakCommand()
	return
}

func (d *Discord) CreateSpeakCommand() (*discordgo.ApplicationCommand, error) {
	option := &discordgo.ApplicationCommandOption{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "text",
		Required:    true,
		Description: "할말",
	}

	return d.ApplicationCommandCreate(&discordgo.ApplicationCommand{
		Name:        "speak",
		Description: "tts 실행",
		Options:     []*discordgo.ApplicationCommandOption{option},
	})
}

func (d *Discord) CreateChangeVoiceCommand() (*discordgo.ApplicationCommand, error) {
	option := &discordgo.ApplicationCommandOption{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "actor",
		Description: "지원하는 목록 확인하려면 help",
	}
	return d.ApplicationCommandCreate(&discordgo.ApplicationCommand{
		Name:        "voice",
		Description: "tts 목소리 변경",
		Options:     []*discordgo.ApplicationCommandOption{option},
	})
}

func (d *Discord) Serve() (err error) {
	d.AddHandler(d.OnInteractionCreate)
	d.AddHandler(d.OnCreateMessage)
	d.Open()
	return
}

func (d *Discord) OnCreateMessage(s *discordgo.Session, event *discordgo.MessageCreate) {
}

func (d *Discord) OnInteractionCreate(s *discordgo.Session, event *discordgo.InteractionCreate) {
	switch event.Type {
	case discordgo.InteractionApplicationCommand:
		d.OnInteractionApplicationCommand(event)
	case discordgo.InteractionMessageComponent:
		d.OnInteractionMessageComponent(event)
	}
}

func (d *Discord) OnInteractionApplicationCommand(event *discordgo.InteractionCreate) {
	switch event.ApplicationCommandData().Name {
	case "voice":
		d.OnExecuteVoiceCommand(event)
	case "speak":
		d.OnSpeakTts(event)
	}
}

func (d *Discord) OnInteractionMessageComponent(event *discordgo.InteractionCreate) {
	switch CustomID := event.MessageComponentData().CustomID; {
	case strings.HasPrefix(CustomID, VoiceActorsListPage):
		d.OnClickVoiceListPaginateButton(event)
	case CustomID == VoiceActorSelect:
		d.OnVoiceActorSlect(event)
	}
}

func (d *Discord) Close() {
	d.Session.Close()
}

func (d *Discord) ChannelVoiceJoin(gID, cID string, mute, deaf bool) (voiceConnection *VoiceConnection, err error) {
	voiceConnection, _ = d.VoiceConnections[gID]
	if voiceConnection == nil {
		var voice *discordgo.VoiceConnection
		voice, err = d.Session.ChannelVoiceJoin(gID, cID, mute, deaf)
		voiceConnection = &VoiceConnection{
			VoiceConnection: voice,
			Mutex:           &sync.Mutex{},
			idleConnect:     make(chan bool),
		}
		d.VoiceConnections[gID] = voiceConnection
	}
	go d.DisconnectIdleConnect(voiceConnection)
	return
}

func (d *Discord) RefreshIdleConnect(vc *VoiceConnection) {
	vc.idleConnect <- true
	go d.DisconnectIdleConnect(vc)
}

func (d *Discord) DisconnectIdleConnect(vc *VoiceConnection) {
	select {
	case <-vc.idleConnect:
		break
	case <-time.After(time.Second * 300):
		d.Disconnect(vc)
	}
}

func (d *Discord) Disconnect(vc *VoiceConnection) {
	delete(d.VoiceConnections, vc.GuildID)
	vc.Disconnect()
}

func (vc *VoiceConnection) PlayAudioFile(filename string, stop <-chan bool) {
	vc.Mutex.Lock()
	defer vc.Mutex.Unlock()

	dgvoice.PlayAudioFile(vc.VoiceConnection, filename, stop)
}
