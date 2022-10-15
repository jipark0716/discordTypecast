package wire

import (
	"github.com/jipark0716/discordTypecast/database"
	"github.com/jipark0716/discordTypecast/repositories"
	"github.com/jipark0716/discordTypecast/services"
)

func NewDiscordService() (discord services.Discord, err error) {
	typecastService := repositories.NewTypecast()
	connection, err := database.GetConnection()
	if err != nil {
		return
	}
	userRepository := repositories.NewUserRepository(connection)
	return services.NewDiscord(typecastService, userRepository)
}
