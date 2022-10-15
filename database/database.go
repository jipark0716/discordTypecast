package database

import (
	"github.com/jipark0716/discordTypecast/config"
	"github.com/jipark0716/discordTypecast/ent"
)

func GetConnection() (*ent.Client, error) {
	return ent.Open("mysql", config.Get("DATABASE_CONNECTION_STRING").(string))
}
