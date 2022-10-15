package services

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/jipark0716/discordTypecast/config"
	"github.com/jipark0716/discordTypecast/repositories"
)

type Discord struct {
	*discordgo.Session
	Typecast       repositories.Typecast
	UserRepository repositories.UserRepository
	ApplicationId  string
}

func NewDiscord(
	typecast repositories.Typecast,
	userRepository repositories.UserRepository,
) (discord Discord, err error) {

	discord = Discord{}
	discord.UserRepository = userRepository
	discord.Typecast = typecast
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
