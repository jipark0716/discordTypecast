package services

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/jipark0716/discordTypecast/handler"
)

const (
	VoiceActorsListPage string = "voice-actors-list-page-"
	VoiceActorSelect    string = "voice-actor-select"
)

func (d *Discord) OnExecuteVoiceCommand(event *discordgo.InteractionCreate) {
	if len(event.ApplicationCommandData().Options) == 0 {
		d.ReplyActorList(event.Interaction, 1)
	}
}

func (d *Discord) OnClickVoiceListPaginateButton(event *discordgo.InteractionCreate) {
	page, err := strconv.Atoi(strings.Replace(event.MessageComponentData().CustomID, VoiceActorsListPage, "", 1))
	if err != nil {
		handler.Report(err)
	}
	d.ReplyActorList(event.Interaction, page)
}

func (d *Discord) ReplyActorList(interaction *discordgo.Interaction, page int) {
	pageLength := 25
	actors, err := d.Typecast.GetActors()
	if err != nil {
		handler.Report(err)
		return
	}

	lastpage := (len(actors) + pageLength - 1) / pageLength

	voiceSelect := discordgo.SelectMenu{
		CustomID: VoiceActorSelect,
	}
	startIndex := ((page - 1) * pageLength)
	endIndex := (page * pageLength)
	if endIndex > len(actors) {
		endIndex = len(actors) - 1
	}
	for _, actor := range actors[startIndex:endIndex] {
		voiceSelect.Options = append(voiceSelect.Options, discordgo.SelectMenuOption{
			Label: actor.Name.Ko,
			Value: actor.ActorId,
		})
	}

	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("%d 페이지", page),
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{voiceSelect},
				},
				d.GetPaginateButton(page, lastpage, VoiceActorsListPage),
			},
		},
	}

	if interaction.Message != nil {
		response.Type = discordgo.InteractionResponseUpdateMessage
	}

	err = d.InteractionRespond(
		interaction,
		response,
	)

	if err != nil {
		handler.Report(err)
	}
}

func (d *Discord) GetPaginateButton(page int, lastpage int, idPrefix string) (actionsRow discordgo.ActionsRow) {
	actionsRow.Components = append(
		actionsRow.Components,
		discordgo.Button{
			Style:    discordgo.PrimaryButton,
			Label:    "prev",
			CustomID: fmt.Sprintf("%s%d", idPrefix, page-1),
			Disabled: page <= 1,
		},
		discordgo.Button{
			Style:    discordgo.PrimaryButton,
			Label:    "next",
			CustomID: fmt.Sprintf("%s%d", idPrefix, page+1),
			Disabled: lastpage <= page,
		},
	)
	return
}

func (d *Discord) OnVoiceActorSlect(event *discordgo.InteractionCreate) {
	values := event.Interaction.MessageComponentData().Values
	if len(values) < 1 {
		handler.Report(fmt.Errorf("not enough values actor select"))
		return
	}

	err := d.UserRepository.SaveActor(event.Member.User.ID, values[0])
	if err != nil {
		handler.Report(err)
	}

	err = d.InteractionRespond(
		event.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseDeferredMessageUpdate,
		},
	)

	if err != nil {
		handler.Report(err)
	}
}
