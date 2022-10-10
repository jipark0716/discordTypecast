package services

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Discord struct {
	*discordgo.Session
}

var discordInstance *Discord

func GetDiscordInstance() *Discord {

	if discordInstance == nil {
		var err error
		config := GetConfigInstance()
		discordInstance = new(Discord)
		discordInstance.Session, err = discordgo.New(config.Get("DISCORD_TOKEN").(string))

		if err != nil {
			log.Fatal(err)
		}
	}

	return discordInstance
}

func (d *Discord) Start() (err error) {
	err = d.Open()
	// d.AddHandler()
	return err
}

func (d *Discord) ApplicationCommands() ([]*discordgo.ApplicationCommand, error) {
	return d.Session.ApplicationCommands(
		GetConfigInstance().Get("DISCORD_ID").(string),
		"",
	)
}

func (d *Discord) ApplicationCommandCreate(cmd *discordgo.ApplicationCommand) (*discordgo.ApplicationCommand, error) {
	cmd.ApplicationID = GetConfigInstance().Get("DISCORD_ID").(string)
	return d.Session.ApplicationCommandCreate(
		GetConfigInstance().Get("DISCORD_ID").(string),
		"",
		cmd,
	)
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
	d.Start()
	return
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
