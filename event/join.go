package event

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
)

func OnGuildMemeberAdd(discord *discordgo.Session, event *discordgo.GuildMemberAdd) {
	println("asd")
	guildName := event.GuildID
	Guild, err := discord.Guild(event.GuildID)
	if err != nil {
		guildName = Guild.Name
	}
	fmt.Printf("%#v\n%#v\n%#v\n", guildName, event.User.Username, event.Member.Nick)
}

func OnGuildMemeberRemove(discord *discordgo.Session, event *discordgo.GuildMemberRemove) {
	println("remove")
}
