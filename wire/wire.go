package wire

import (
	"github.com/jipark0716/discordTypecast/database"
	"github.com/jipark0716/discordTypecast/discord"
	"github.com/jipark0716/discordTypecast/repositories"
)

func NewDiscordService() (d discord.Discord, err error) {
	typecastService := repositories.NewTypecast()
	connection, err := database.GetConnection()
	if err != nil {
		return
	}
	userRepository := repositories.NewUserRepository(connection)
	messageRepository := repositories.NewTypecastMessageRepository(connection)
	return discord.NewDiscord(typecastService, userRepository, messageRepository)
}
