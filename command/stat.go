package main

import (
	"os"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jipark0716/discordTypecast/model"
)

func main() {
	guildsID, _ := model.GuildsID()

	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		fmt.Println(err)
		return
	}

	guilds := []*discordgo.Guild{}

	for _, id := range guildsID {
		guild, err := discord.Guild(id)
		if err == nil {
			guilds = append(guilds, guild)
		}
	}

	for i, guild := range guilds {
		fmt.Printf("%d : %s\n", i, guild.Name)
	}

	inputInt := -1
	i := 0
	for {
		_, err = fmt.Scanln(&inputInt)
		if err != nil || inputInt < 0 || inputInt >= len(guilds) {
			i += 1
			if i == 3 {
				println("입력 횟수를 초과해서 종료합니다.")
				return
			} else {
				println("입력을 확인해주세요.")
			}
		} else {
			break
		}
	}

	guild := guilds[inputInt]

	cnt, _ := model.ChatCount(guild.ID)

	for _, memberCount := range cnt {
		member, err := discord.GuildMember(guild.ID, memberCount.MemberID)
		if err == nil {
			name := member.Nick
			if name == "" {
				name = member.User.Username
			}
			fmt.Printf("%s : %d / %s\n", name, memberCount.Count, member.User.ID)
		}
	}
	return
}
