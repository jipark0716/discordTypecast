package main

import (
	"fmt"
	"os"
	"log"

	"github.com/bwmarrin/discordgo"
	eventHandler "github.com/jipark0716/discordTypecast/event"
)

func main() {
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		fmt.Println(err)
		return
	}

	discord.AddHandler(eventHandler.OnCreateMessage)
	discord.AddHandler(eventHandler.OnGuildMemeberAdd)
	discord.AddHandler(eventHandler.OnGuildMemeberRemove)

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("start server")

	select{}
}
